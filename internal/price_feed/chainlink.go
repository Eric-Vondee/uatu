package feeds

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/uatu/config"
	"github.com/uatu/internal/contracts"
)

const (
	ChainLink           = "chainlink"
	PriceCacheKeyPrefix = "chainlink:price:"
	priceFetchLimit     = 10
)

func PriceCacheKey(tokenSlug string) string {
	return PriceCacheKeyPrefix + ":" + strings.ToLower(tokenSlug)
}

type TokenFeed struct {
	PriceFeedAddress  common.Address
	PriceFeedProvider string
	ChainSlug         string
	Slugs             []string
}

type TokenFeedResponse struct {
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	Logo          string `json:"logo"`
	Slug          string `json:"slug"`
	ChainSlug     string `json:"chainSlug"`
	Price         string `json:"price"`
	PriceAnswer   string `json:"priceAnswer"`
	PriceDecimals uint8  `json:"priceDecimals"`
	UpdatedAt     int64  `json:"updatedAt"`
	FetchedAt     int64  `json:"fetchedAt"`
	RoundID       string `json:"roundId,omitempty"`
	FeedAddress   string `json:"feedAddress,omitempty"`
	Provider      string `json:"provider"`
}

func FetchTokenPrices(
	ctx context.Context,
	cfg config.Config,
) []TokenFeedResponse {
	prices := make([]TokenFeedResponse, 0, countSlugs(SupportedTokens))
	for _, token := range SupportedTokens {
		if !strings.EqualFold(token.PriceFeedProvider, ChainLink) {
			continue
		}
		currentToken, err := fetchChainlinkPrice(ctx, cfg, token)
		if err != nil {
			continue
		}
		for _, slug := range token.Slugs {
			price := currentToken
			price.Slug = slug
			price.Name = slug
			price.Symbol = strings.ToUpper(slug)
			prices = append(prices, price)
		}
	}
	return prices
}

func countSlugs(tokens []TokenFeed) int {
	total := 0
	for _, token := range tokens {
		total += len(token.Slugs)
	}
	return total
}

func fetchChainlinkPrice(
	ctx context.Context,
	cfg config.Config,
	token TokenFeed,
) (TokenFeedResponse, error) {
	chainSlug := token.ChainSlug
	rpcURL := cfg.GetRPC(token.ChainSlug)
	if len(token.Slugs) == 0 {
		return TokenFeedResponse{}, fmt.Errorf("%s: price feed has no token slugs", chainSlug)
	}
	if token.PriceFeedAddress == (common.Address{}) {
		return TokenFeedResponse{}, fmt.Errorf("%s: price feed address is empty", strings.Join(token.Slugs, "/"))
	}
	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return TokenFeedResponse{}, fmt.Errorf("client connection failed for %s: %w", chainSlug, err)
	}
	defer client.Close()

	priceFeedContract, err := contracts.NewContracts(token.PriceFeedAddress, client)
	if err != nil {
		return TokenFeedResponse{}, fmt.Errorf("contract connection failed for %s: %w", chainSlug, err)
	}
	callOpts := &bind.CallOpts{Context: ctx}
	decimals, err := priceFeedContract.Decimals(callOpts)
	if err != nil {
		return TokenFeedResponse{}, fmt.Errorf("decimals fetch failed for %s: %w", chainSlug, err)
	}

	data, err := priceFeedContract.LatestRoundData(callOpts)
	if err != nil {
		return TokenFeedResponse{}, fmt.Errorf("latest round data failed for %s: %w", chainSlug, err)
	}
	if data.Answer == nil || data.Answer.Sign() <= 0 {
		return TokenFeedResponse{}, fmt.Errorf("chainlink returned a non-positive answer")
	}
	if data.UpdatedAt == nil || data.UpdatedAt.Sign() <= 0 {
		return TokenFeedResponse{}, fmt.Errorf("chainlink returned an invalid update time")
	}
	if data.RoundId == nil || data.RoundId.Sign() <= 0 {
		return TokenFeedResponse{}, fmt.Errorf("chainlink returned an invalid round id")
	}

	return TokenFeedResponse{
		Price:         formatAnswer(data.Answer, decimals),
		PriceAnswer:   data.Answer.String(),
		PriceDecimals: decimals,
		UpdatedAt:     data.UpdatedAt.Int64(),
		FetchedAt:     time.Now().Unix(),
		RoundID:       data.RoundId.String(),
		FeedAddress:   token.PriceFeedAddress.Hex(),
		ChainSlug:     token.ChainSlug,
		Provider:      ChainLink,
	}, nil
}

func formatAnswer(answer *big.Int, decimals uint8) string {
	scale := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	return new(big.Rat).SetFrac(answer, scale).FloatString(18)
}
