package jobs

import (
	"context"
	"fmt"

	"github.com/uatu/config"
	feeds "github.com/uatu/internal/price_feed"
	redisstore "github.com/uatu/internal/storage/redis"
)

func SyncChainlinkPrices(
	ctx context.Context,
	cfg config.Config,
	redisClient *redisstore.RedisService,
) error {
	tokenPrices := feeds.FetchTokenPrices(ctx, cfg)
	if len(tokenPrices) == 0 {
		return fmt.Errorf("no prices fetched")
	}

	cacheData := make(map[string]any, len(tokenPrices))
	for _, tokenPrice := range tokenPrices {
		cacheData[feeds.PriceCacheKey(tokenPrice.Slug)] = tokenPrice
	}
	if err := redisClient.InsertAll(ctx, cacheData, tokenPriceCacheTTL); err != nil {
		return fmt.Errorf("bulk write failed: %w", err)
	}
	return nil
}

func syncChainlinkPrices(
	ctx context.Context,
	cfg config.Config,
	redisClient *redisstore.RedisService,
) func() error {
	return func() error {
		return SyncChainlinkPrices(ctx, cfg, redisClient)
	}
}
