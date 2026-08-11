package jobs

import (
	"context"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/uatu"
	"github.com/uatu/config"
	redisstore "github.com/uatu/internal/storage/redis"
)

const (
	tokenPriceRefreshInterval = 1 * time.Second
	tokenPriceCacheTTL        = 2 * time.Minute
	poolRefreshInterval       = 3 * time.Minute
)

func Startup(
	ctx context.Context,
	cfg config.Config,
	redisClient *redisstore.RedisService,
	chainRepo uatu.ChainRepository,
) (func(), error) {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	jobCtx, cancel := context.WithCancel(ctx)

	_, err = scheduler.NewJob(
		gocron.DurationJob(tokenPriceRefreshInterval),
		gocron.NewTask(syncChainlinkPrices(jobCtx, cfg, redisClient)),
		gocron.WithContext(jobCtx),
		gocron.WithSingletonMode(gocron.LimitModeReschedule),
		gocron.WithStartAt(gocron.WithStartImmediately()),
	)

	if err != nil {
		cancel()
		return nil, err
	}
	_, err = scheduler.NewJob(
		gocron.DurationJob(poolRefreshInterval),
		gocron.NewTask(syncPools(jobCtx, cfg, chainRepo, redisClient)),
		gocron.WithContext(jobCtx),
		gocron.WithSingletonMode(gocron.LimitModeReschedule),
	)
	if err != nil {
		cancel()
		return nil, err
	}

	scheduler.Start()

	return func() {
		cancel()
		_ = scheduler.Shutdown()
	}, nil
}
