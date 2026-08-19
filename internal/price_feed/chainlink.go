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
	Slug              string
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
	prices := make([]TokenFeedResponse, 0, len(SupportedTokens))
	for _, token := range SupportedTokens {
		if strings.ToLower(token.PriceFeedProvider) == "chainlink" {
			price, err := fetchChainlinkPrice(ctx, cfg, token)
			if err != nil {
				continue
			}
			prices = append(prices, price)
		}
	}
	return prices
}

func fetchChainlinkPrice(
	ctx context.Context,
	cfg config.Config,
	token TokenFeed,
) (TokenFeedResponse, error) {
	chainSlug := token.ChainSlug
	rpcURL := cfg.GetRPC(token.ChainSlug)
	if token.PriceFeedAddress == (common.Address{}) {
		return TokenFeedResponse{}, fmt.Errorf("%s: price feed address is empty", token.Slug)
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
		Name:          chainSlug,
		Symbol:        strings.ToUpper(chainSlug),
		Price:         formatAnswer(data.Answer, decimals),
		PriceAnswer:   data.Answer.String(),
		PriceDecimals: decimals,
		UpdatedAt:     data.UpdatedAt.Int64(),
		FetchedAt:     time.Now().Unix(),
		RoundID:       data.RoundId.String(),
		FeedAddress:   token.PriceFeedAddress.Hex(),
		Slug:          token.Slug,
		ChainSlug:     token.ChainSlug,
		Provider:      ChainLink,
	}, nil
}

func formatAnswer(answer *big.Int, decimals uint8) string {
	scale := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	return new(big.Rat).SetFrac(answer, scale).FloatString(18)
}

var SupportedTokens = []TokenFeed{
	{
		PriceFeedAddress:  common.HexToAddress("0x3E7d1eAB13ad0104d2750B8863b489D65364e32D"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "ethereum",
		Slug:              "usdt",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x3E7d1eAB13ad0104d2750B8863b489D65364e32D"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "ethereum",
		Slug:              "usdt0",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x8fFfFfd4AfB6115b954Bd326cbe7B4BA576818f6"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "ethereum",
		Slug:              "usdc",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x8fFfFfd4AfB6115b954Bd326cbe7B4BA576818f6"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "ethereum",
		Slug:              "usdc.e",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0xb9fB4e65744E4178894f7C61CF80E8a48A5f224a"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "robinhood",
		Slug:              "usde",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x61B7e5650328764B076A108EFF5fa7282a1B9aD2"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "robinhood",
		Slug:              "usdg",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0xAed0c38402a5d19df6E4c03F4E2DceD6e29c1ee9"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "ethereum",
		Slug:              "dai",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "ethereum",
		Slug:              "eth",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "ethereum",
		Slug:              "weth",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x0567F2323251f0Aab15c8dFb1967E4e8A7D42aeE"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "bsc",
		Slug:              "bnb",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x0567F2323251f0Aab15c8dFb1967E4e8A7D42aeE"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "bsc",
		Slug:              "wbnb",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x82BA56a2fADF9C14f17D08bc51bDA0bDB83A8934"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "arbitrum",
		Slug:              "pol",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x82BA56a2fADF9C14f17D08bc51bDA0bDB83A8934"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "arbitrum",
		Slug:              "wpol",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x205aaD468a11fd5D34fA7211bC6Bad5b3deB9b98"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "arbitrum",
		Slug:              "optimism",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x0568fD19986748cEfF3301e55c0eb1E729E0Ab7e"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "celo",
		Slug:              "celo",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0xe38A27BE4E7d866327e09736F3C570F256FFd048"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "celo",
		Slug:              "cusd",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x0A77230d17318075983913bC2145DB16C7366156"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "avalanche",
		Slug:              "avax",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x0A77230d17318075983913bC2145DB16C7366156"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "avalanche",
		Slug:              "wavax",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0xBcD78f76005B7515837af6b50c7C52BCf73822fb"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "monad",
		Slug:              "mon",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0xBcD78f76005B7515837af6b50c7C52BCf73822fb"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "monad",
		Slug:              "wmon",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0xD97F20bEbeD74e8144134C4b148fE93417dd0F96"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "mantle",
		Slug:              "mnt",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0xD97F20bEbeD74e8144134C4b148fE93417dd0F96"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "mantle",
		Slug:              "wmnt",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0xF932477C37715aE6657Ab884414Bd9876FE3f750"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "plasma",
		Slug:              "xpl",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0xF932477C37715aE6657Ab884414Bd9876FE3f750"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "plasma",
		Slug:              "wxpl",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x678df3415fc31947dA4324eC63212874be5a82f8"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "gnosis",
		Slug:              "xdai",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x678df3415fc31947da4324ec63212874be5a82f8"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "gnosis",
		Slug:              "wxdai",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0xF4030086522a5bEEa4988F8cA5B36dbC97BeE88c"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "ethereum",
		Slug:              "wbtc",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0xdAd6f90429a2C821496B78Fe7482412971E278f1"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "unichain",
		Slug:              "uni",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x22441d81416430A54336aB28765abd31a792Ad37"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "gnosis",
		Slug:              "gno",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x8Bb2943AB030E3eE05a58d9832525B4f60A97FA0"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "worldchain",
		Slug:              "wld",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0xe38A27BE4E7d866327e09736F3C570F256FFd048"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "celo",
		Slug:              "usdm",
	},
	{
		PriceFeedAddress:  common.HexToAddress("0xb2A824043730FE05F3DA2efaFa1CBbe83fa548D6"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "arbitrum",
		Slug:              "arb",
	},
}
