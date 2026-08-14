package dex

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/uatu"
	"github.com/uatu/internal/assets"
	feeds "github.com/uatu/internal/price_feed"
	redisstore "github.com/uatu/internal/storage/redis"
	"go.uber.org/zap"
)

const (
	basisPointsDenominator = 10000
	poolFeeDenominator     = 1_000_000
	minOracleOutputBps     = 9800
	maxOracleOutputBps     = 10200
	maxOraclePriceAge      = 90 * time.Second
)

var DexRoutes = map[string]uatu.Route{
	"agni":        {Name: "Agni Finance", Logo: assets.Agni, Dex: "agni"},
	"aerodrome":   {Name: "Aerodrome", Logo: assets.Aerodrome, Dex: "aerodrome"},
	"cowswap":     {Name: "CowSwap", Logo: assets.CowSwap, Dex: "cowswap"},
	"oku":         {Name: "OkuTrade", Logo: assets.Oku, Dex: "oku"},
	"pancakeswap": {Name: "PancakeSwap", Logo: assets.PancakSwap, Dex: "pancakeswap"},
	"quickswap":   {Name: "QuickSwap", Logo: assets.QuickSwap, Dex: "quickswap"},
	"sushiswap":   {Name: "SushiSwap", Logo: assets.SushiSwap, Dex: "sushiswap"},
	"uniswap":     {Name: "Uniswap", Logo: assets.Uniswap, Dex: "uniswap"},
}

type quoteRoute struct {
	pool *uatu.Pool
	dex  uatu.Dex
}

type quoteResult struct {
	response *uatu.IDexResponse
	dex      string
	err      error
}

type oraclePrice struct {
	answer    *big.Int
	decimals  uint8
	updatedAt time.Time
	fetchedAt time.Time
}

func routeWithPoolFee(route uatu.Route, amountIn *big.Int, poolFee uint) uatu.Route {
	return routeWithFee(
		route,
		poolFeeAmount(amountIn, new(big.Int).SetUint64(uint64(poolFee))),
	)
}

func routeWithFee(route uatu.Route, fee *big.Int) uatu.Route {
	route.Fees = new(big.Int).Set(fee).String()
	return route
}

func poolFeeAmount(amountIn, poolFee *big.Int) *big.Int {
	return new(big.Int).Div(
		new(big.Int).Mul(amountIn, poolFee),
		big.NewInt(poolFeeDenominator),
	)
}

type BestQuoteParams struct {
	AmountIn           *big.Int
	Chain              uatu.Chain
	TokenIn            uatu.Token
	TokenOut           uatu.Token
	ExecutionTokenIn   uatu.Token
	ExecutionTokenOut  uatu.Token
	WrapNativeInput    bool
	UnwrapNativeOutput bool
	WalletAddress      common.Address
	Pools              []uatu.Pool
	RPCURL             string
	PriceCache         *redisstore.RedisService
}

func WrappedNativeToken(chain uatu.Chain, token uatu.Token) (uatu.Token, bool, error) {
	if uatu.FormatEvmAddress(token.Address) != (common.Address{}) {
		return token, false, nil
	}

	wrappedSymbol := "W" + token.Symbol
	for _, candidate := range chain.Tokens {
		if strings.EqualFold(candidate.Symbol, wrappedSymbol) {
			return candidate, true, nil
		}
	}

	return uatu.Token{}, false, fmt.Errorf("%s is not configured for %s", wrappedSymbol, chain.Name)
}

func cachedOraclePrice(
	ctx context.Context,
	token uatu.Token,
	priceCache *redisstore.RedisService,
) (oraclePrice, error) {
	if priceCache == nil {
		return oraclePrice{}, fmt.Errorf("price cache is not configured")
	}

	var response feeds.TokenFeedResponse

	if err := priceCache.Get(ctx, feeds.PriceCacheKey(token.Slug), &response); err != nil {
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

func GetDexQuotes(
	ctx context.Context,
	params *BestQuoteParams,
) ([]*uatu.IDexResponse, error) {
	if params == nil {
		return nil, fmt.Errorf("best quote parameters are required")
	}
	amountIn := params.AmountIn
	chain := params.Chain
	tokenIn := params.TokenIn
	tokenOut := params.TokenOut
	walletAddress := params.WalletAddress
	pools := params.Pools
	rpcURL := params.RPCURL
	priceCache := params.PriceCache

	client, err := Provider(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("could not connect to %s rpc: %w", chain.Slug, err)
	}

	routes := make([]quoteRoute, 0, len(pools)+1)
	for i := range pools {
		pool := pools[i]
		d, ok := dexForSlug(chain.Dex, pool.DexName)
		if !ok {
			return nil, fmt.Errorf("could not find dex %s on %s", pool.DexName, chain.Slug)
		}
		routes = append(routes, quoteRoute{pool: &pool, dex: d})
	}

	for _, d := range chain.Dex {
		if d.Slug != CowSlug {
			continue
		}
		if !params.WrapNativeInput && !params.UnwrapNativeOutput {
			routes = append(routes, quoteRoute{dex: d})
		}
		break
	}
	if len(routes) == 0 {
		return nil, fmt.Errorf("no eligible %s/%s pools found on %s", params.ExecutionTokenIn.Symbol, params.ExecutionTokenOut.Symbol, chain.Name)
	}

	priceIn, err := cachedOraclePrice(ctx, tokenIn, priceCache)
	if err != nil {
		return nil, fmt.Errorf("could not load oracle price for %s: %w", tokenIn.Symbol, err)
	}
	priceOut, err := cachedOraclePrice(ctx, tokenOut, priceCache)
	if err != nil {
		return nil, fmt.Errorf("could not load oracle price for %s: %w", tokenOut.Symbol, err)
	}

	tokenInAddress := uatu.FormatEvmAddress(params.ExecutionTokenIn.Address)
	tokenOutAddress := uatu.FormatEvmAddress(params.ExecutionTokenOut.Address)

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
				TokenIn:            tokenInAddress,
				TokenOut:           tokenOutAddress,
				AmountIn:           amountIn,
				PairAddress:        pairAddress,
				PoolFee:            poolFee,
				WalletAddress:      walletAddress,
				ChainId:            chain.ChainID,
				Dex:                route.dex,
				WrapNativeInput:    params.WrapNativeInput,
				UnwrapNativeOutput: params.UnwrapNativeOutput,
				PoolType:           poolType,
			}
			var (
				output *uatu.IDexResponse
				err    error
			)
			switch route.dex.Slug {
			case "uniswap", "pancakeswap", "oku":
				output, err = client.Uniswap(ctx, dexRequest)
			case "agni":
				output, err = client.Agni(ctx, dexRequest)
			case "quickswap":
				output, err = client.QuickSwap(ctx, dexRequest)
			case "aerodrome":
				output, err = client.Aerodrome(ctx, dexRequest)
			case CowSlug:
				output, err = client.CowSwap(ctx, dexRequest)
			default:
				err = fmt.Errorf("unsupported dex %q", route.dex.Slug)
			}
			results <- quoteResult{response: output, dex: route.dex.Name, err: err}
		}(route)
	}

	var (
		quotesByDex = make(map[string]*uatu.IDexResponse, len(routes))
		errs        = make([]error, 0, len(routes))
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
			errs = append(errs, fmt.Errorf("%s: %w", r.dex, r.err))
			continue
		case r.response == nil || r.response.AmountOut == nil:
			errs = append(errs, fmt.Errorf("%s: did not return a quote", r.dex))
			continue
		}
		if err := guardSwapOutput(amountIn, r.response.AmountOut, tokenIn, tokenOut, priceIn, priceOut); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", r.dex, err))
			continue
		}

		best, exists := quotesByDex[r.response.Dex.Slug]
		if !exists || r.response.AmountOut.Cmp(best.AmountOut) > 0 {
			quotesByDex[r.response.Dex.Slug] = r.response
		}
	}
	quotes := make([]*uatu.IDexResponse, 0, len(quotesByDex))
	for _, quote := range quotesByDex {
		quotes = append(quotes, quote)
	}
	if len(quotes) == 0 {
		if len(errs) > 0 {
			return nil, fmt.Errorf("no acceptable route found: %w", errors.Join(errs...))
		}
		return nil, fmt.Errorf("no acceptable route found")
	}
	sort.Slice(quotes, func(i, j int) bool {
		return quotes[i].AmountOut.Cmp(quotes[j].AmountOut) > 0
	})
	return quotes, nil
}

func GetBestDexQuote(
	ctx context.Context,
	params *BestQuoteParams,
) (*uatu.IDexResponse, error) {
	quotes, err := GetDexQuotes(ctx, params)
	if err != nil {
		return nil, err
	}
	best := quotes[0]
	if best.RegisterOrder != nil {
		if err := best.RegisterOrder(ctx); err != nil {
			return nil, fmt.Errorf("could not register selected DEX order: %w", err)
		}
	}
	return best, nil
}

func dexForSlug(dexes []uatu.Dex, slug string) (uatu.Dex, bool) {
	for _, d := range dexes {
		if d.Slug == slug {
			return d, true
		}
	}
	return uatu.Dex{}, false
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
			zap.String("amountOut", amountOut.String()),
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
