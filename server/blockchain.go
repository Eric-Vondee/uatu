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
