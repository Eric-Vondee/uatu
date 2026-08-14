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

func newQuoteID() string {
	id := uuid.New()
	return id.String()
}

func Deadline() *big.Int {
	quoteTTL := time.Minute
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
// @Param message body uatu.QuoteRequest true "request body to create a swap quote"
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
	executionTokenIn, wrapNativeInput, err := dex.WrappedNativeToken(chain, tokenIn)
	if err != nil {
		return APIError{
			newAPIResponse(http.StatusBadRequest, err.Error(), nil),
		}, err
	}
	executionTokenOut, unwrapNativeOutput, err := dex.WrappedNativeToken(chain, tokenOut)
	if err != nil {
		return APIError{
			newAPIResponse(http.StatusBadRequest, err.Error(), nil),
		}, err
	}

	pools, err := q.chainRepo.GetPools(ctx, uatu.QueryOptions{
		TokenIn:  executionTokenIn.Address,
		TokenOut: executionTokenOut.Address,
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

	res, err := dex.GetBestDexQuote(ctx, &dex.BestQuoteParams{
		AmountIn:           amountIn,
		Chain:              chain,
		TokenIn:            tokenIn,
		TokenOut:           tokenOut,
		ExecutionTokenIn:   executionTokenIn,
		ExecutionTokenOut:  executionTokenOut,
		WrapNativeInput:    wrapNativeInput,
		UnwrapNativeOutput: unwrapNativeOutput,
		WalletAddress:      walletAddress,
		Pools:              pools,
		RPCURL:             q.cfg.GetRPC(chain.Slug),
		PriceCache:         q.priceCache,
	})
	if err != nil {
		logger.Error("Failed to get best route", zap.Error(err))
		return APIError{
			newAPIResponse(http.StatusInternalServerError, err.Error(), nil),
		}, err
	}

	quoteResponse, err := q.newQuote(ctx, res, chain, tokenIn, tokenOut, walletAddress)
	if err != nil {
		logger.Error("Failed to build quote", zap.Error(err))
		return newAPIResponse(
			http.StatusInternalServerError,
			err.Error(),
			nil,
		), err
	}
	return newAPIResponse(http.StatusOK, "Quote created successfully", quoteResponse), nil
}

// GetQuotes returns one non-persisted quote option for each supported DEX.
//
// @Summary List DEX quote options
// @Description Prices a swap against every supported DEX and returns the valid routes ordered by output amount.
// @Tags quotes
// @Accept json
// @Produce json
// @Param message body uatu.QuoteRequest true "request body to compare DEX quotes"
// @Success 200 {object} APIResponse{data=[]uatu.RouteQuote}
// @Failure 400 {object} APIResponse
// @Failure 500 {object} APIResponse
// @Router /quotes/routes [post]
func (q *quoteHandler) GetQuotes(
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

	chain, err := q.chainRepo.GetBlockchain(ctx, uatu.QueryOptions{ChainID: req.ChainID})
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
	executionTokenIn, wrapNativeInput, err := dex.WrappedNativeToken(chain, tokenIn)
	if err != nil {
		return APIError{
			newAPIResponse(http.StatusBadRequest, err.Error(), nil),
		}, err
	}
	executionTokenOut, unwrapNativeOutput, err := dex.WrappedNativeToken(chain, tokenOut)
	if err != nil {
		return APIError{
			newAPIResponse(http.StatusBadRequest, err.Error(), nil),
		}, err
	}

	pools, err := q.chainRepo.GetPools(ctx, uatu.QueryOptions{
		TokenIn:  executionTokenIn.Address,
		TokenOut: executionTokenOut.Address,
		ChainID:  req.ChainID,
	})
	if err != nil {
		logger.Error("Failed to get pools", zap.Error(err))
		return APIError{
			newAPIResponse(http.StatusInternalServerError, "an error occurred fetching pools", nil),
		}, err
	}

	responses, err := dex.GetDexQuotes(ctx, &dex.BestQuoteParams{
		AmountIn:           ConvertDecimalToBigInt(req.Amount, tokenIn.Decimals),
		Chain:              chain,
		TokenIn:            tokenIn,
		TokenOut:           tokenOut,
		ExecutionTokenIn:   executionTokenIn,
		ExecutionTokenOut:  executionTokenOut,
		WrapNativeInput:    wrapNativeInput,
		UnwrapNativeOutput: unwrapNativeOutput,
		WalletAddress:      uatu.FormatEvmAddress(req.RecipientAddress),
		Pools:              pools,
		RPCURL:             q.cfg.GetRPC(chain.Slug),
		PriceCache:         q.priceCache,
	})
	if err != nil {
		logger.Error("Failed to get DEX quotes", zap.Error(err))
		return APIError{
			newAPIResponse(http.StatusInternalServerError, err.Error(), nil),
		}, err
	}

	quotes := make([]uatu.RouteQuote, 0, len(responses))
	deadline := Deadline()
	for _, response := range responses {
		quotes = append(quotes, uatu.RouteQuote{
			AmountIn:  response.AmountIn.String(),
			AmountOut: response.AmountOut.String(),
			Deadline:  deadline,
			TokenIn:   tokenIn,
			TokenOut:  tokenOut,
			Route:     response.Route,
		})
	}

	return newAPIResponse(http.StatusOK, "DEX quotes fetched successfully", quotes), nil
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

func (q *quoteHandler) newQuote(
	ctx context.Context,
	res *uatu.IDexResponse,
	chain uatu.Chain,
	tokenIn, tokenOut uatu.Token,
	walletAddress common.Address,
) (*uatu.QuoteResponse, error) {
	quoteID := newQuoteID()
	deadline := Deadline()
	amountIn := res.AmountIn.String()
	wallet := walletAddress.String()

	step := func(msg, to string, data []byte, value *big.Int) uatu.Actions {
		valueString := ""
		if value != nil {
			valueString = value.String()
		}
		return uatu.Actions{
			Message: msg,
			Amount:  amountIn,
			Value:   valueString,
			From:    wallet,
			To:      to,
			Data:    uatu.HexBytes(data),
			ChainID: chain.ChainID,
		}
	}

	steps := make([]uatu.Actions, 0, 3)
	if len(res.EncodedERC20Approval) != 0 {
		msg := fmt.Sprintf("%s token approval", tokenIn.Symbol)
		steps = append(steps, step(msg, tokenIn.Address, res.EncodedERC20Approval, nil))
	}

	if len(res.EncodedPermit2Approval) != 0 {
		msg := "Permit2 approval"
		steps = append(steps, step(msg, res.Dex.Permit2Address, res.EncodedPermit2Approval, nil))
	}

	steps = append(steps, step(
		fmt.Sprintf("Swap %s to %s", tokenIn.Symbol, tokenOut.Symbol),
		res.RouterAddress.String(),
		res.EncodedData,
		res.NativeValue,
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
		Route:       res.Route,
		Deadline:    deadline,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	if err := q.quoteRepo.Create(ctx, &quote); err != nil {
		return nil, fmt.Errorf("an error occurred creating quote")
	}
	return &uatu.QuoteResponse{Quote: quote}, nil
}
