package server

import (
	"context"
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
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type GenericRequest struct{}

func (g GenericRequest) Bind(_ *http.Request) error { return nil }

type quoteHandler struct {
	cfg       config.Config
	quoteRepo uatu.QuoteRepository
	chainRepo uatu.ChainRepository
}

// quoteRequest is the body accepted by POST /quotes. Amounts are given in
// human units (not base units); token fields take checksummed or lowercase
// EVM addresses.
type quoteRequest struct {
	Amount           decimal.Decimal `json:"amount" validate:"required" swaggertype:"string"`
	Chain            string          `json:"chain" validate:"required"`
	ChainID          uint            `json:"chainId" validate:"required"`
	RecipientAddress string          `json:"recipientAddress" validate:"required"`
	TokenIn          string          `json:"tokenIn" validate:"required"`
	TokenOut         string          `json:"tokenOut" validate:"required"`
	GenericRequest
}

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
	req := new(quoteRequest)
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

	tokenIn, _ := getToken(chain.Tokens, uatu.FormatEvmAddress(req.TokenIn))
	tokenOut, _ := getToken(chain.Tokens, uatu.FormatEvmAddress(req.TokenOut))

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
	token0 := uatu.FormatEvmAddress(tokenIn.Address)
	token1 := uatu.FormatEvmAddress(tokenOut.Address)

	res, err := q.getBestOutput(
		ctx, amountIn, chain,
		token0, token1, walletAddress,
		pools,
	)
	if err != nil {
		logger.Error("Failed to get best route", zap.Error(err))
		return APIError{
			newAPIResponse(http.StatusInternalServerError, "an error occurred fetching quote", nil),
		}, err
	}

	quote, err := newQuote(res, chain, tokenIn, tokenOut, walletAddress)
	if err := q.quoteRepo.Create(ctx, quote); err != nil {
		logger.Error("Failed to create quote", zap.Error(err))
		return newAPIResponse(
				http.StatusInternalServerError,
				"an error occurred creating quote",
				nil,
			),
			err
	}
	return newAPIResponse(http.StatusOK, "Quote created successfully", quote), nil
}

func newQuote(
	res *uatu.IDexResponse,
	chain uatu.Chain,
	tokenIn, tokenOut uatu.Token,
	walletAddress common.Address,
) (*uatu.Quote, error) {
	quoteID := newQuoteID()
	deadline := Deadline()
	steps := []uatu.Actions{
		{
			Message: "Permit2 Token Approval",
			Amount:  res.AmountIn.String(),
			From:    walletAddress.String(),
			To:      res.Dex.Permit2Address,
			Data:    uatu.HexBytes(res.EncodedTokenApproval),
			ChainID: chain.ChainID,
		},
		{
			Message: fmt.Sprintf("Swap %s to %s", tokenIn.Symbol, tokenOut.Symbol),
			Amount:  res.AmountIn.String(),
			From:    walletAddress.String(),
			To:      res.Dex.UniversalRouterAddress,
			Data:    uatu.HexBytes(res.EncodedData),
			ChainID: chain.ChainID,
		},
	}
	quote := uatu.Quote{
		QuoteId:            quoteID,
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
		ExplorerUrl: chain.BlockExplorer,
		Steps:       steps,
		Status:      uatu.Pending,
		Deadline:    deadline,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	return &quote, nil
}

type quoteResult struct {
	response *uatu.IDexResponse
	err      error
}

func (q *quoteHandler) getBestOutput(
	ctx context.Context,
	amountIn *big.Int,
	chain uatu.Chain,
	tokenIn, tokenOut, walletAddress common.Address,
	pools []uatu.Pool,
) (*uatu.IDexResponse, error) {
	client, err := dex.Provider(q.cfg.GetRPC(chain.Slug))
	if err != nil {
		return nil, fmt.Errorf("could not connect to %s rpc: %w", chain.Slug, err)
	}

	dexBySlug := make(map[string]uatu.Dex, len(pools))
	for _, pool := range pools {
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
	}

	results := make(chan quoteResult, len(pools))
	for _, pool := range pools {
		go func(pool uatu.Pool, d uatu.Dex) {
			pairAddress := uatu.FormatEvmAddress(pool.PairAddress)
			dexRequest := uatu.IDexRequest{
				TokenIn:       tokenIn,
				TokenOut:      tokenOut,
				AmountIn:      amountIn,
				PairAddress:   pairAddress,
				PoolFee:       pool.PoolFee,
				WalletAddress: walletAddress,
				ChainId:       chain.ChainID,
				Dex:           d,
				PoolType:      pool.PoolType,
			}
			var (
				output *uatu.IDexResponse
				err    error
			)
			output, err = client.Swap(ctx, dexRequest)
			results <- quoteResult{response: output, err: err}
		}(pool, dexBySlug[pool.DexName])
	}
	var (
		best    *uatu.IDexResponse
		lastErr error
	)

	for range pools {
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

		if best == nil || r.response.AmountOut.Cmp(best.AmountOut) > 0 {
			best = r.response
		}
	}

	if best == nil {
		if lastErr != nil {
			return nil, fmt.Errorf("no supported route found: %w", lastErr)
		}
		return nil, fmt.Errorf("no supported route found")
	}
	if best.AmountOut == nil {
		return nil, fmt.Errorf("Best swap quote missing output amount")
	}
	return best, nil
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
