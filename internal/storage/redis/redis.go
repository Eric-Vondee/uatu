package redisstore

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/uatu/config"
)

var ErrCacheMiss = errors.New("cache miss")

type RedisService struct {
	redis *redis.Client
}

func InitializeRedis(ctx context.Context, cfg config.RedisConfig) (*RedisService, error) {
	opts, err := redis.ParseURL(cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("parse redis dsn: %w", err)
	}
	opts.PoolSize = 10
	opts.MinIdleConns = 2
	opts.DialTimeout = 5 * time.Second
	opts.ReadTimeout = 3 * time.Second
	opts.WriteTimeout = 3 * time.Second

	redisClient := redis.NewClient(opts)
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := redisClient.Ping(pingCtx).Err(); err != nil {
		_ = redisClient.Close()
		return nil, fmt.Errorf("ping Redis: %w", err)
	}

	return &RedisService{redis: redisClient}, nil
}

func (r *RedisService) Close() error {
	if r == nil || r.redis == nil {
		return nil
	}
	return r.redis.Close()
}

func (r *RedisService) Exists(ctx context.Context, key string) (bool, error) {
	value, err := r.redis.Exists(ctx, strings.ToLower(key)).Result()
	if err != nil {
		return false, err
	}
	return value > 0, nil
}

func (r *RedisService) Get(ctx context.Context, key string, result any) error {
	value, err := r.redis.Get(ctx, strings.ToLower(key)).Result()
	if err == redis.Nil {
		return ErrCacheMiss
	}
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(value), result); err != nil {
		return fmt.Errorf("decode redis value: %w", err)
	}
	return nil
}

func (r *RedisService) Set(ctx context.Context, key string, payload any, ttl time.Duration) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return r.redis.Set(ctx, strings.ToLower(key), data, ttl).Err()
}

func (r *RedisService) Delete(ctx context.Context, key string) (bool, error) {
	value, err := r.redis.Del(ctx, strings.ToLower(key)).Result()
	if err != nil {
		return false, err
	}
	return value > 0, nil
}

func (r *RedisService) InsertAll(ctx context.Context, data map[string]any, ttl time.Duration) error {
	if len(data) == 0 {
		return nil
	}
	pipeline := r.redis.TxPipeline()
	for key, payload := range data {
		encoded, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("encode %s: %w", key, err)
		}
		pipeline.Set(ctx, strings.ToLower(key), encoded, ttl)
	}
	_, err := pipeline.Exec(ctx)
	return err
}
