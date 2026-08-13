package server

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"github.com/uatu"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type chainHandler struct {
	chainRepo uatu.ChainRepository
}

// GetBlockchains lists every chain in the catalogue.
//
// @Summary List supported blockchains
// @Description Returns every chain the catalogue knows about, each with its embedded
// @Description token and DEX definitions.
// @Tags catalogue
// @Produce json
// @Success 200 {object} APIResponse{data=[]uatu.Chain}
// @Failure 500 {object} APIResponse
// @Router /blockchains [get]
func (c *chainHandler) GetBlockchains(
	ctx context.Context,
	span trace.Span,
	logger *zap.Logger,
	w http.ResponseWriter,
	r *http.Request,
) (render.Renderer, error) {
	chains, err := c.chainRepo.GetBlockchains(ctx)
	if err != nil {
		logger.Error("Failed to get blockchains", zap.Error(err))
		return APIError{
			newAPIResponse(http.StatusInternalServerError, "an error occurred fetching blockchains", nil),
		}, err
	}
	return newAPIResponse(http.StatusOK, "Blockchains fetched successfully", chains), nil
}

// GetTokens lists known tokens, optionally narrowed to one chain.
//
// @Summary List tokens
// @Description Returns the known tokens. Omit chainId to list tokens across every chain.
// @Tags catalogue
// @Produce json
// @Param chainId query int false "restrict results to this EVM chain ID"
// @Success 200 {object} APIResponse{data=[]uatu.Token}
// @Failure 400 {object} APIResponse
// @Failure 500 {object} APIResponse
// @Router /blockchains/tokens [get]
func (c *chainHandler) GetTokens(
	ctx context.Context,
	span trace.Span,
	logger *zap.Logger,
	w http.ResponseWriter,
	r *http.Request,
) (render.Renderer, error) {
	opts := uatu.QueryOptions{}

	if raw := r.URL.Query().Get("chainId"); raw != "" {
		chainID, err := strconv.ParseUint(raw, 10, 64)
		if err != nil {
			return APIError{
				newAPIResponse(http.StatusBadRequest, "invalid chain id", nil),
			}, err
		}
		opts.ChainID = uint(chainID)
	}

	tokens, err := c.chainRepo.GetTokens(ctx, opts)
	if err != nil {
		logger.Error("Failed to get tokens", zap.Error(err))
		return APIError{
			newAPIResponse(http.StatusInternalServerError, "an error occurred fetching tokens", nil),
		}, err
	}
	return newAPIResponse(http.StatusOK, "Tokens fetched successfully", tokens), nil
}

// GetPools lists the liquidity pools discovered for a chain.
//
// @Summary List pools for a chain
// @Description Returns the liquidity pools discovered by the seeder for the given chain,
// @Description optionally filtered to a single DEX. Unlike the token endpoint, chainId
// @Description is required here.
// @Tags catalogue
// @Produce json
// @Param chainId query int true "EVM chain ID"
// @Param dex query string false "restrict results to this DEX slug"
// @Success 200 {object} APIResponse{data=[]uatu.Pool}
// @Failure 400 {object} APIResponse
// @Failure 500 {object} APIResponse
// @Router /blockchains/pools [get]
func (c *chainHandler) GetPools(
	ctx context.Context,
	span trace.Span,
	logger *zap.Logger,
	w http.ResponseWriter,
	r *http.Request,
) (render.Renderer, error) {
	chainID, err := strconv.ParseUint(r.URL.Query().Get("chainId"), 10, 64)
	if err != nil {
		return APIError{
			newAPIResponse(http.StatusBadRequest, "invalid chain id", nil),
		}, err
	}
	pools, err := c.chainRepo.GetPools(ctx, uatu.QueryOptions{
		ChainID: uint(chainID),
		DexName: r.URL.Query().Get("dex"),
	})
	if err != nil {
		logger.Error("Failed to get pools", zap.Error(err), zap.Uint64("chainID", chainID))
		return APIError{
			newAPIResponse(http.StatusInternalServerError, "an error occurred fetching pools", nil),
		}, err
	}
	return newAPIResponse(http.StatusOK, "Pools fetched successfully", pools), nil
}

// GetDex returns one DEX's contract addresses on a chain.
//
// @Summary Get a DEX on a chain
// @Description Returns the DEX seeded under the given slug on the given chain, including
// @Description the router, factory, quoter and settlement addresses used to price and
// @Description encode a swap. Both chainId and slug are required, as a slug identifies a
// @Description DEX only within one chain.
// @Tags catalogue
// @Produce json
// @Param chainId query int true "EVM chain ID"
// @Param slug query string true "DEX slug, e.g. uniswap"
// @Success 200 {object} APIResponse{data=uatu.Dex}
// @Failure 400 {object} APIResponse
// @Failure 500 {object} APIResponse
// @Router /blockchains/dex [get]
func (c *chainHandler) GetDex(
	ctx context.Context,
	span trace.Span,
	logger *zap.Logger,
	w http.ResponseWriter,
	r *http.Request,
) (render.Renderer, error) {
	chainID, err := strconv.ParseUint(r.URL.Query().Get("chainId"), 10, 64)
	if err != nil {
		return APIError{
			newAPIResponse(http.StatusBadRequest, "invalid chain id", nil),
		}, err
	}
	dex, err := c.chainRepo.GetDex(ctx, uatu.QueryOptions{
		ChainID: uint(chainID),
	})
	if err != nil {
		logger.Error("Failed to get dex",
			zap.Error(err),
			zap.Uint64("chainID", chainID),
		)
		return APIError{
			newAPIResponse(http.StatusInternalServerError, "an error occurred fetching dex", nil),
		}, err
	}
	return newAPIResponse(http.StatusOK, "Dex fetched successfully", dex), nil
}
