package main

import (
	"log"

	"github.com/uatu/config"
	"github.com/uatu/internal/storage/postgres"
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

	db, err := postgres.DbConnection(*cfg, logger)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() { _ = db.Close() }()

	quoteRepo := postgres.NewQuoteRepository(db)
	chainRepo := postgres.NewChainRepository(db)

	srv := server.New(*cfg, logger, quoteRepo, chainRepo)
	if err := srv.Run(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
