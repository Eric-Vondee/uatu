package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/uatu/config"
	"github.com/uatu/internal/storage/postgres"
	redisstore "github.com/uatu/internal/storage/redis"
	"github.com/uatu/jobs"
	"github.com/uatu/server"
	"go.uber.org/zap"
)

func main() {
	if err := run(); err != nil {
		log.Printf("uatu server stopped: %v", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.InitializeConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	if err = cfg.Validate(); err != nil {
		return fmt.Errorf("validate config: %w", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return fmt.Errorf("create logger: %w", err)
	}
	defer func() { _ = logger.Sync() }()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	otelShutdown, err := server.InitOTELCapabilities(ctx, *cfg)
	if err != nil {
		return fmt.Errorf("initialize OpenTelemetry: %w", err)
	}
	defer func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := otelShutdown(shutdownCtx); err != nil {
			logger.Error("Failed to shut down OpenTelemetry", zap.Error(err))
		}
	}()

	db, err := postgres.DbConnection(*cfg, logger)
	if err != nil {
		return fmt.Errorf("connect to database: %w", err)
	}
	defer func() { _ = db.Close() }()

	redisClient, err := redisstore.InitializeRedis(ctx, cfg.Redis)
	if err != nil {
		return fmt.Errorf("connect to Redis: %w", err)
	}
	defer func() { _ = redisClient.Close() }()

	quoteRepo := postgres.NewQuoteRepository(db)
	chainRepo := postgres.NewChainRepository(db)

	stopJobs, err := jobs.Startup(ctx, *cfg, redisClient, chainRepo)
	if err != nil {
		return fmt.Errorf("start background jobs: %w", err)
	}
	defer stopJobs()

	srv, err := server.New(*cfg, logger, quoteRepo, chainRepo, redisClient)
	if err != nil {
		return fmt.Errorf("create server: %w", err)
	}
	return srv.Run(ctx)
}
