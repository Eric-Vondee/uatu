package server

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"slices"
	"time"

	"github.com/Eric-Vondee/metron"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/uatu"
	"github.com/uatu/config"
	"github.com/uatu/internal/dex"
	feeds "github.com/uatu/internal/price_feed"
	redisstore "github.com/uatu/internal/storage/redis"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type quoteHandler struct {
	cfg        config.Config
	quoteRepo  uatu.QuoteRepository
	chainRepo  uatu.ChainRepository
	priceCache *redisstore.RedisService
}

const (
	basisPointsDenominator = 10000
	minOracleOutputBps     = 9800
	maxOracleOutputBps     = 10200
	maxOraclePriceAge      = 90 * time.Second
)

func newQuoteID() string {
	id := uuid.New()
	return id.String()
}

func Deadline() *big.Int {
	quoteTTL := 5 * time.Minute
	return big.NewInt(time.Now().Add(quoteTTL).Unix())
}

// CreateQuote prices a swap across every known pool for the pair and returns
// the best route.
//
// @Summary Create a swap quote
// @Description Resolves the pools for a token pair on the given chain, prices the swap against
// @Description each DEX's on-chain contracts concurrently, and returns the best output together
// @Description with the encoded Permit2 approval and swap calldata needed to execute it.
// @Description Quotes carry a 5-minute deadline and are persisted with a pending status.
// @Tags quotes
// @Accept json
// @Produce json
// @Param message body quoteRequest true "request body to create a swap quote"
// @Success 200 {object} APIResponse{data=uatu.Quote}
// @Failure 400 {object} APIResponse
// @Failure 500 {object} APIResponse
// @Router /quotes [post]
func (q *quoteHandler) CreateQuote(
	ctx context.Context,
	span trace.Span,
	logger *zap.Logger,
	w http.ResponseWriter,
	r *http.Request,
) (render.Renderer, error) {
	req := new(uatu.QuoteRequest)
	if err := render.Bind(r, req); err != nil {
		return nil, err
	}
	if err := metron.ValidateStruct(req); err != nil {
		return APIError{
			newAPIResponse(http.StatusBadRequest, err.Error(), nil),
		}, err
	}

	chain, err := q.chainRepo.GetBlockchain(ctx, uatu.QueryOptions{
		ChainID: req.ChainID,
	})
	if err != nil {
		logger.Error("Failed to get chain", zap.Error(err))
		return APIError{
			newAPIResponse(http.StatusInternalServerError, "an error occurred fetching chain", nil),
		}, err
	}

	tokenIn, tokenOut, err := getTokenInAndOut(chain.Tokens, req.TokenIn, req.TokenOut)
	if err != nil {
		return APIError{
			newAPIResponse(http.StatusBadRequest, err.Error(), nil),
		}, err
	}

	pools, err := q.chainRepo.GetPools(ctx, uatu.QueryOptions{
		TokenIn:  tokenIn.Address,
		TokenOut: tokenOut.Address,
		ChainID:  req.ChainID,
	})
	if err != nil {
		logger.Error("Failed to get pools", zap.Error(err))
		return APIError{
			newAPIResponse(http.StatusInternalServerError, "an error occurred fetching pools", nil),
		}, err
	}

	amountIn := ConvertDecimalToBigInt(req.Amount, tokenIn.Decimals)
	walletAddress := uatu.FormatEvmAddress(req.RecipientAddress)

	res, err := q.getBestOutput(
		ctx, amountIn, chain,
		tokenIn, tokenOut, walletAddress,
		pools,
	)
	if err != nil {
		logger.Error("Failed to get best route", zap.Error(err))
		return APIError{
			newAPIResponse(http.StatusInternalServerError, err.Error(), nil),
		}, err
	}

	quoteResponse, err := newQuote(res, chain, tokenIn, tokenOut, walletAddress)
	if err != nil {
		logger.Error("Failed to build quote", zap.Error(err))
		return newAPIResponse(
			http.StatusInternalServerError,
			"an error occurred building quote",
			nil,
		), err
	}
	if err := q.quoteRepo.Create(ctx, &quoteResponse.Quote); err != nil {
		logger.Error("Failed to create quote", zap.Error(err))
		return newAPIResponse(
				http.StatusInternalServerError,
				"an error occurred creating quote",
				nil,
			),
			err
	}
	return newAPIResponse(http.StatusOK, "Quote created successfully", quoteResponse), nil
}

func getTokenInAndOut(
	tokens []uatu.Token,
	token0, token1 string,
) (uatu.Token, uatu.Token, error) {
	tokenIn, err := getToken(tokens, uatu.FormatEvmAddress(token0))
	if err != nil {
		return uatu.Token{}, uatu.Token{}, fmt.Errorf("tokenIn was not found on the selected chain")
	}
	tokenOut, err := getToken(tokens, uatu.FormatEvmAddress(token1))
	if err != nil {
		return uatu.Token{}, uatu.Token{}, fmt.Errorf("tokenOut was not found on the selected chain")
	}

	return tokenIn, tokenOut, nil
}

func newQuote(
	res *uatu.IDexResponse,
	chain uatu.Chain,
	tokenIn, tokenOut uatu.Token,
	walletAddress common.Address,
) (*uatu.QuoteResponse, error) {
	quoteID := newQuoteID()
	deadline := Deadline()
	amountIn := res.AmountIn.String()
	wallet := walletAddress.String()

	step := func(msg, to string, data []byte) uatu.Actions {
		return uatu.Actions{
			Message: msg,
			Amount:  amountIn,
			From:    wallet,
			To:      to,
			Data:    uatu.HexBytes(data),
			ChainID: chain.ChainID,
		}
	}

	steps := make([]uatu.Actions, 0, 3)
	if len(res.EncodedERC20Approval) != 0 {
		msg := fmt.Sprintf("%s token approval", tokenIn.Symbol)
		steps = append(steps, step(msg, tokenIn.Address, res.EncodedERC20Approval))
	}

	if len(res.EncodedPermit2Approval) != 0 {
		msg := "Permit2 approval"
		steps = append(steps, step(msg, res.Dex.Permit2Address, res.EncodedPermit2Approval))
	}

	steps = append(steps, step(
		fmt.Sprintf("Swap %s to %s", tokenIn.Symbol, tokenOut.Symbol),
		res.RouterAddress.String(),
		res.EncodedData,
	))
	quote := uatu.Quote{
		QuoteID:            quoteID,
		AmountIn:           res.AmountIn.String(),
		AmountOut:          res.AmountOut.String(),
		AmountInFloat:      decimal.NewFromFloat(0),
		AmountOutFloat:     decimal.NewFromFloat(0),
		OriginChainId:      chain.ChainID,
		OriginChain:        chain.Name,
		DestinationChainId: chain.ChainID,
		DestinationChain:   chain.Name,
		WalletAddress:      walletAddress.String(),
		RecipientAddress:   walletAddress.String(),
		TokenIn: uatu.Token{
			Address:  tokenIn.Address,
			Decimals: tokenIn.Decimals,
			Logo:     tokenIn.Logo,
			Slug:     tokenIn.Slug,
			Symbol:   tokenIn.Symbol,
			Name:     tokenIn.Name,
		},
		TokenOut: uatu.Token{
			Address:  tokenOut.Address,
			Decimals: tokenOut.Decimals,
			Logo:     tokenOut.Logo,
			Slug:     tokenOut.Slug,
			Symbol:   tokenOut.Symbol,
			Name:     tokenOut.Name,
		},
		PairAddress: res.PairAddress.String(),
		ExplorerUrl: chain.BlockExplorer,
		Steps:       steps,
		Status:      uatu.Pending,
		Deadline:    deadline,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	return &uatu.QuoteResponse{Quote: quote}, nil
}

type quoteResult struct {
	response *uatu.IDexResponse
	err      error
}

type quoteRoute struct {
	pool *uatu.Pool
	dex  uatu.Dex
}

func (q *quoteHandler) cachedOraclePrice(ctx context.Context, token uatu.Token) (oraclePrice, error) {
	if q.priceCache == nil {
		return oraclePrice{}, fmt.Errorf("price cache is not configured")
	}

	var response feeds.TokenFeedResponse
	if err := q.priceCache.Get(ctx, feeds.PriceCacheKey(token.Slug), &response); err != nil {
		if errors.Is(err, redisstore.ErrCacheMiss) {
			return oraclePrice{}, fmt.Errorf("price is unavailable")
		}
		return oraclePrice{}, err
	}
	answer, ok := new(big.Int).SetString(response.PriceAnswer, 10)
	if !ok || answer.Sign() <= 0 {
		return oraclePrice{}, fmt.Errorf("cached price is invalid")
	}
	if response.UpdatedAt <= 0 || response.FetchedAt <= 0 {
		return oraclePrice{}, fmt.Errorf("cached price has invalid timestamps")
	}
	updatedAt := time.Unix(response.UpdatedAt, 0)
	if updatedAt.After(time.Now().Add(time.Minute)) {
		return oraclePrice{}, fmt.Errorf("cached price has a future update time")
	}
	fetchedAt := time.Unix(response.FetchedAt, 0)
	if age := time.Since(fetchedAt); age < 0 || age > maxOraclePriceAge {
		return oraclePrice{}, fmt.Errorf("cached price is stale")
	}

	return oraclePrice{
		answer:    answer,
		decimals:  response.PriceDecimals,
		updatedAt: updatedAt,
		fetchedAt: fetchedAt,
	}, nil
}

func (q *quoteHandler) getBestOutput(
	ctx context.Context,
	amountIn *big.Int,
	chain uatu.Chain,
	tokenIn, tokenOut uatu.Token,
	walletAddress common.Address,
	pools []uatu.Pool,
) (*uatu.IDexResponse, error) {
	client, err := dex.Provider(q.cfg.GetRPC(chain.Slug))
	if err != nil {
		return nil, fmt.Errorf("could not connect to %s rpc: %w", chain.Slug, err)
	}

	priceIn, err := q.cachedOraclePrice(ctx, tokenIn)
	if err != nil {
		return nil, fmt.Errorf("could not load oracle price for %s: %w", tokenIn.Symbol, err)
	}
	priceOut, err := q.cachedOraclePrice(ctx, tokenOut)
	if err != nil {
		return nil, fmt.Errorf("could not load oracle price for %s: %w", tokenOut.Symbol, err)
	}

	tokenInAddress := uatu.FormatEvmAddress(tokenIn.Address)
	tokenOutAddress := uatu.FormatEvmAddress(tokenOut.Address)

	dexBySlug := make(map[string]uatu.Dex, len(pools)+1)
	routes := make([]quoteRoute, 0, len(pools)+1)
	for i := range pools {
		pool := pools[i]
		if _, ok := dexBySlug[pool.DexName]; ok {
			continue
		}
		d, err := q.chainRepo.GetDex(ctx, uatu.QueryOptions{
			ChainID: chain.ChainID,
			Slug:    pool.DexName,
		})
		if err != nil {
			return nil, fmt.Errorf("could not get dex %s on %s: %w", pool.DexName, chain.Slug, err)
		}
		dexBySlug[pool.DexName] = d
		routes = append(routes, quoteRoute{pool: &pool, dex: d})
	}

	for _, d := range chain.Dex {
		if d.Slug != dex.CowSlug {
			continue
		}
		if _, exists := dexBySlug[d.Slug]; !exists {
			dexBySlug[d.Slug] = d
			routes = append(routes, quoteRoute{dex: d})
		}
		break
	}

	results := make(chan quoteResult, len(routes))
	for _, route := range routes {
		go func(route quoteRoute) {
			var (
				poolFee     uint
				poolType    string
				pairAddress common.Address
			)
			if route.pool != nil {
				poolFee = route.pool.PoolFee
				poolType = route.pool.PoolType
				pairAddress = uatu.FormatEvmAddress(route.pool.PairAddress)
			}
			dexRequest := uatu.IDexRequest{
				TokenIn:       tokenInAddress,
				TokenOut:      tokenOutAddress,
				AmountIn:      amountIn,
				PairAddress:   pairAddress,
				PoolFee:       poolFee,
				WalletAddress: walletAddress,
				ChainId:       chain.ChainID,
				Dex:           route.dex,
				PoolType:      poolType,
			}
			var (
				output *uatu.IDexResponse
				err    error
			)
			switch route.dex.Slug {
			case "uniswap", "pancakeswap", "oku":
				output, err = client.Swap(ctx, dexRequest)
			case "aerodrome":
				output, err = client.SwapAerodrome(ctx, dexRequest)
			case dex.CowSlug:
				output, err = client.SwapCow(ctx, dexRequest)
			default:
				err = fmt.Errorf("unsupported dex %q", route.dex.Slug)
			}
			results <- quoteResult{response: output, err: err}
		}(route)
	}
	var (
		best    *uatu.IDexResponse
		lastErr error
	)

	for range routes {
		var r quoteResult
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case r = <-results:
		}

		switch {
		case r.err != nil:
			lastErr = r.err
			continue
		case r.response == nil || r.response.AmountOut == nil:
			continue
		}
		if err := guardSwapOutput(amountIn, r.response.AmountOut, tokenIn, tokenOut, priceIn, priceOut); err != nil {
			lastErr = err
			continue
		}

		if best == nil || r.response.AmountOut.Cmp(best.AmountOut) > 0 {
			best = r.response
		}
	}
	if best == nil {
		if lastErr != nil {
			return nil, fmt.Errorf("no acceptable route found: %w", lastErr)
		}
		return nil, fmt.Errorf("no acceptable route found")
	}
	if best.RegisterOrder != nil {
		if err := best.RegisterOrder(ctx); err != nil {
			return nil, fmt.Errorf("could not register selected DEX order: %w", err)
		}
	}
	return best, nil
}

type oraclePrice struct {
	answer    *big.Int
	decimals  uint8
	updatedAt time.Time
	fetchedAt time.Time
}

func guardSwapOutput(
	amountIn, amountOut *big.Int,
	tokenIn, tokenOut uatu.Token,
	priceIn, priceOut oraclePrice,
) error {
	if amountIn == nil || amountIn.Sign() <= 0 {
		return fmt.Errorf("swap input amount must be positive")
	}
	if amountOut == nil || amountOut.Sign() <= 0 {
		zap.L().Warn("swap quote returned no output",
			zap.String("tokenIn", tokenIn.Symbol),
			zap.String("tokenOut", tokenOut.Symbol),
			zap.String("amountIn", amountIn.String()),
			zap.String("amountOut", amountString(amountOut)),
		)
		return fmt.Errorf("this swap would result in no payout")
	}
	if priceIn.answer == nil || priceIn.answer.Sign() <= 0 || priceOut.answer == nil || priceOut.answer.Sign() <= 0 {
		return fmt.Errorf("oracle prices must be positive")
	}

	numerator, denominator := oracleOutputFraction(amountIn, tokenIn, tokenOut, priceIn, priceOut)
	expectedOut := new(big.Int).Quo(new(big.Int).Set(numerator), denominator)
	if expectedOut.Sign() <= 0 {
		return fmt.Errorf("oracle expected output is too small")
	}

	amountOutScaled := new(big.Int).Mul(amountOut, denominator)
	amountOutScaled.Mul(amountOutScaled, big.NewInt(basisPointsDenominator))
	minScaled := new(big.Int).Mul(numerator, big.NewInt(minOracleOutputBps))
	maxScaled := new(big.Int).Mul(numerator, big.NewInt(maxOracleOutputBps))
	if amountOutScaled.Cmp(minScaled) < 0 || amountOutScaled.Cmp(maxScaled) > 0 {
		zap.L().Warn("swap quote outside oracle bounds",
			zap.String("tokenIn", tokenIn.Symbol),
			zap.String("tokenOut", tokenOut.Symbol),
			zap.String("amountIn", amountIn.String()),
			zap.String("amountOut", amountOut.String()),
			zap.String("expectedOut", expectedOut.String()),
		)
		return fmt.Errorf("swap output failed")
	}
	return nil
}

func oracleOutputFraction(
	amountIn *big.Int,
	tokenIn, tokenOut uatu.Token,
	priceIn, priceOut oraclePrice,
) (*big.Int, *big.Int) {
	base := big.NewInt(10)
	inTokenScale := new(big.Int).Exp(base, big.NewInt(int64(tokenIn.Decimals)), nil)
	outTokenScale := new(big.Int).Exp(base, big.NewInt(int64(tokenOut.Decimals)), nil)
	inPriceScale := new(big.Int).Exp(base, big.NewInt(int64(priceIn.decimals)), nil)
	outPriceScale := new(big.Int).Exp(base, big.NewInt(int64(priceOut.decimals)), nil)

	numerator := new(big.Int).Mul(amountIn, priceIn.answer)
	numerator.Mul(numerator, outTokenScale)
	numerator.Mul(numerator, outPriceScale)
	denominator := new(big.Int).Mul(priceOut.answer, inTokenScale)
	denominator.Mul(denominator, inPriceScale)
	return numerator, denominator
}

func amountString(amount *big.Int) string {
	if amount == nil {
		return "<nil>"
	}
	return amount.String()
}

func ConvertDecimalToBigInt(amount decimal.Decimal, decimals uint8) *big.Int {
	return amount.Shift(int32(decimals)).BigInt()
}

func getToken(tokens []uatu.Token, tokenAddress common.Address) (uatu.Token, error) {
	idx := slices.IndexFunc(tokens, func(t uatu.Token) bool {
		return uatu.FormatEvmAddress(t.Address) == tokenAddress
	})
	if idx == -1 {
		return uatu.Token{}, fmt.Errorf("token not found")
	}
	return tokens[idx], nil
}
