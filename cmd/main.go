package main

import (
	"context"
	"log"

	"github.com/uatu/config"
	"github.com/uatu/internal/storage/postgres"
	redisstore "github.com/uatu/internal/storage/redis"
	"github.com/uatu/jobs"
	"github.com/uatu/server"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.InitializeConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err = cfg.Validate(); err != nil {
		log.Fatalf("Invalid config: %v", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}
	defer func() { _ = logger.Sync() }()

	otelShutdown, err := server.InitOTELCapabilities(context.Background(), *cfg)
	if err != nil {
		log.Fatalf("Failed to initialize OpenTelemetry: %v", err)
	}
	defer func() {
		if err := otelShutdown(context.Background()); err != nil {
			logger.Error("Failed to shut down OpenTelemetry", zap.Error(err))
		}
	}()

	db, err := postgres.DbConnection(*cfg, logger)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() { _ = db.Close() }()

	redisClient, err := redisstore.InitializeRedis(context.Background(), cfg.Redis)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer func() { _ = redisClient.Close() }()

	stopJobs, err := jobs.Startup(context.Background(), *cfg, redisClient)
	if err != nil {
		log.Fatalf("Failed to start background jobs: %v", err)
	}
	defer stopJobs()

	quoteRepo := postgres.NewQuoteRepository(db)
	chainRepo := postgres.NewChainRepository(db)

	srv := server.New(*cfg, logger, quoteRepo, chainRepo, redisClient)
	if err := srv.Run(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
