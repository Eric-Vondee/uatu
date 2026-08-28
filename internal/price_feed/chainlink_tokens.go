package feeds

import "github.com/ethereum/go-ethereum/common"

var SupportedTokens = []TokenFeed{
	{
		PriceFeedAddress:  common.HexToAddress("0x3E7d1eAB13ad0104d2750B8863b489D65364e32D"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "ethereum",
		Slugs:             []string{"usdt", "usdt0"},
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x8fFfFfd4AfB6115b954Bd326cbe7B4BA576818f6"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "ethereum",
		Slugs:             []string{"usdc", "usdc.e"},
	},
	{
		PriceFeedAddress:  common.HexToAddress("0xb9fB4e65744E4178894f7C61CF80E8a48A5f224a"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "robinhood",
		Slugs:             []string{"usde"},
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x61B7e5650328764B076A108EFF5fa7282a1B9aD2"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "robinhood",
		Slugs:             []string{"usdg"},
	},
	{
		PriceFeedAddress:  common.HexToAddress("0xAed0c38402a5d19df6E4c03F4E2DceD6e29c1ee9"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "ethereum",
		Slugs:             []string{"dai"},
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "ethereum",
		Slugs:             []string{"eth", "weth"},
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x0567F2323251f0Aab15c8dFb1967E4e8A7D42aeE"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "bsc",
		Slugs:             []string{"bnb", "wbnb"},
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x82BA56a2fADF9C14f17D08bc51bDA0bDB83A8934"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "arbitrum",
		Slugs:             []string{"pol", "wpol"},
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x205aaD468a11fd5D34fA7211bC6Bad5b3deB9b98"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "arbitrum",
		Slugs:             []string{"optimism"},
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x0568fD19986748cEfF3301e55c0eb1E729E0Ab7e"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "celo",
		Slugs:             []string{"celo"},
	},
	{
		PriceFeedAddress:  common.HexToAddress("0xe38A27BE4E7d866327e09736F3C570F256FFd048"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "celo",
		Slugs:             []string{"cusd", "usdm"},
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x0A77230d17318075983913bC2145DB16C7366156"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "avalanche",
		Slugs:             []string{"avax", "wavax"},
	},
	{
		PriceFeedAddress:  common.HexToAddress("0xBcD78f76005B7515837af6b50c7C52BCf73822fb"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "monad",
		Slugs:             []string{"mon", "wmon"},
	},
	{
		PriceFeedAddress:  common.HexToAddress("0xD97F20bEbeD74e8144134C4b148fE93417dd0F96"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "mantle",
		Slugs:             []string{"mnt", "wmnt"},
	},
	{
		PriceFeedAddress:  common.HexToAddress("0xF932477C37715aE6657Ab884414Bd9876FE3f750"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "plasma",
		Slugs:             []string{"xpl", "wxpl"},
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x678df3415fc31947dA4324eC63212874be5a82f8"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "gnosis",
		Slugs:             []string{"xdai", "wxdai"},
	},
	{
		PriceFeedAddress:  common.HexToAddress("0xF4030086522a5bEEa4988F8cA5B36dbC97BeE88c"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "ethereum",
		Slugs:             []string{"wbtc"},
	},
	{
		PriceFeedAddress:  common.HexToAddress("0xdAd6f90429a2C821496B78Fe7482412971E278f1"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "unichain",
		Slugs:             []string{"uni"},
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x22441d81416430A54336aB28765abd31a792Ad37"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "gnosis",
		Slugs:             []string{"gno"},
	},
	{
		PriceFeedAddress:  common.HexToAddress("0x8Bb2943AB030E3eE05a58d9832525B4f60A97FA0"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "worldchain",
		Slugs:             []string{"wld"},
	},
	{
		PriceFeedAddress:  common.HexToAddress("0xb2A824043730FE05F3DA2efaFa1CBbe83fa548D6"),
		PriceFeedProvider: ChainLink,
		ChainSlug:         "arbitrum",
		Slugs:             []string{"arb"},
	},
}
