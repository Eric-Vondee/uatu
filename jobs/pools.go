package jobs

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/uatu"
	"github.com/uatu/config"
	"github.com/uatu/internal/dex"
	feeds "github.com/uatu/internal/price_feed"
	redisstore "github.com/uatu/internal/storage/redis"
)

const (
	maxPoolSyncs     = 8
	minPoolLiquidity = 100
)

func SyncPools(
	ctx context.Context,
	cfg config.Config,
	chainRepo uatu.ChainRepository,
	priceCache *redisstore.RedisService,
) error {
	chains, err := chainRepo.GetBlockchains(ctx)
	if err != nil {
		return fmt.Errorf("get blockchains: %w", err)
	}
	pools, err := chainRepo.GetPools(ctx, uatu.QueryOptions{IncludeInactive: true})
	if err != nil {
		return fmt.Errorf("get pools: %w", err)
	}

	chainsByID := make(map[uint]uatu.Chain, len(chains))
	for _, chain := range chains {
		chainsByID[chain.ChainID] = chain
	}
	clients := make(map[uint]*dex.Client, len(chains))
	for _, pool := range pools {
		if _, exists := clients[pool.ChainID]; exists {
			continue
		}
		chain, ok := chainsByID[pool.ChainID]
		if !ok {
			return fmt.Errorf("pool %s references unknown chain %d", pool.PairAddress, pool.ChainID)
		}
		rpcURL := cfg.GetRPC(chain.Slug)
		if rpcURL == "" {
			return fmt.Errorf("rpc URL is not configured for %s", chain.Slug)
		}
		client, err := dex.Provider(rpcURL)
		if err != nil {
			return fmt.Errorf("connect to %s RPC: %w", chain.Slug, err)
		}
		clients[pool.ChainID] = client
	}

	var (
		wg   sync.WaitGroup
		mu   sync.Mutex
		errs []error
		sem  = make(chan struct{}, maxPoolSyncs)
	)
	for i := range pools {
		pool := pools[i]
		client := clients[pool.ChainID]
		wg.Go(func() {
			sem <- struct{}{}
			defer func() { <-sem }()

			if err := syncPool(ctx, client, chainRepo, priceCache, &pool); err != nil {
				mu.Lock()
				errs = append(errs, fmt.Errorf("%s: %w", pool.PairAddress, err))
				mu.Unlock()
			}
		})
	}
	wg.Wait()
	return errors.Join(errs...)
}

func syncPools(
	ctx context.Context,
	cfg config.Config,
	chainRepo uatu.ChainRepository,
	priceCache *redisstore.RedisService,
) func() error {
	return func() error {
		return SyncPools(ctx, cfg, chainRepo, priceCache)
	}
}

func syncPool(
	ctx context.Context,
	client *dex.Client,
	chainRepo uatu.ChainRepository,
	priceCache *redisstore.RedisService,
	pool *uatu.Pool,
) error {
	pairAddress := uatu.FormatEvmAddress(pool.PairAddress)
	baseTokenAddress := uatu.FormatEvmAddress(pool.BaseTokenAddress)
	quoteTokenAddress := uatu.FormatEvmAddress(pool.QuoteTokenAddress)
	if pairAddress == (common.Address{}) || baseTokenAddress == (common.Address{}) || quoteTokenAddress == (common.Address{}) {
		return fmt.Errorf("pool or token address is invalid")
	}

	baseBalance, err := client.GetTokenBalance(ctx, baseTokenAddress, pairAddress)
	if err != nil {
		return fmt.Errorf("get base token balance: %w", err)
	}
	quoteBalance, err := client.GetTokenBalance(ctx, quoteTokenAddress, pairAddress)
	if err != nil {
		return fmt.Errorf("get quote token balance: %w", err)
	}

	baseValue, err := tokenBalanceUSD(ctx, priceCache, pool.BaseToken, baseBalance)
	if err != nil {
		return fmt.Errorf("value base token: %w", err)
	}
	quoteValue, err := tokenBalanceUSD(ctx, priceCache, pool.QuoteToken, quoteBalance)
	if err != nil {
		return fmt.Errorf("value quote token: %w", err)
	}
	totalValue := new(big.Rat).Add(baseValue, quoteValue)

	pool.BaseTokenBalance = baseBalance.String()
	pool.QuoteTokenBalance = quoteBalance.String()
	pool.Liquidity = uatu.PoolLiquidity{
		BaseTokenBalance:  baseBalance.String(),
		QuoteTokenBalance: quoteBalance.String(),
	}
	pool.LiquidityInUSD = uatu.PoolLiquidityUSD{
		BaseTokenBalance:  baseValue.FloatString(18),
		QuoteTokenBalance: quoteValue.FloatString(18),
		TotalPoolBalance:  totalValue.FloatString(18),
	}
	pool.IsActivePool = totalValue.Cmp(big.NewRat(minPoolLiquidity, 1)) >= 0

	if err := chainRepo.UpdatePoolLiquidity(ctx, pool); err != nil {
		return err
	}
	return nil
}

func tokenBalanceUSD(
	ctx context.Context,
	priceCache *redisstore.RedisService,
	token uatu.Token,
	balance *big.Int,
) (*big.Rat, error) {
	if balance == nil || balance.Sign() < 0 {
		return nil, fmt.Errorf("token balance is invalid")
	}

	price, err := cachedTokenPrice(ctx, priceCache, token)
	if err != nil {
		return nil, err
	}
	tokenScale := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(token.Decimals)), nil)
	value := new(big.Rat).Mul(new(big.Rat).SetInt(balance), price)
	return value.Quo(value, new(big.Rat).SetInt(tokenScale)), nil
}

func cachedTokenPrice(
	ctx context.Context,
	priceCache *redisstore.RedisService,
	token uatu.Token,
) (*big.Rat, error) {
	if priceCache == nil {
		return nil, fmt.Errorf("price cache is not configured")
	}

	var response feeds.TokenFeedResponse
	if err := priceCache.Get(ctx, feeds.PriceCacheKey(token.Slug), &response); err != nil {
		if errors.Is(err, redisstore.ErrCacheMiss) && token.IsStableCoin {
			return big.NewRat(1, 1), nil
		}
		return nil, fmt.Errorf("get cached price: %w", err)
	}
	answer, ok := new(big.Int).SetString(response.PriceAnswer, 10)
	if !ok || answer.Sign() <= 0 {
		return nil, fmt.Errorf("cached price is invalid")
	}
	priceScale := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(response.PriceDecimals)), nil)
	return new(big.Rat).SetFrac(answer, priceScale), nil
}
