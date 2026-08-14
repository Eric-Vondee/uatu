package redisstore

import (
	"errors"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/sethvargo/go-limiter"
	limiterredis "github.com/sethvargo/go-redisstore"
	"github.com/uatu/config"
)

func NewRateLimitStore(cfg config.RedisConfig, tokens uint64, interval time.Duration) (limiter.Store, error) {
	if cfg.DSN == "" {
		return nil, errors.New("redis DSN is required for the rate limiter")
	}

	store, err := limiterredis.New(&limiterredis.Config{
		Tokens:   tokens,
		Interval: interval,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(cfg.DSN)
		},
	})
	if err != nil {
		return nil, fmt.Errorf("create Redis rate limit store: %w", err)
	}
	return store, nil
}
