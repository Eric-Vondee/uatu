package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/uatu"
	"github.com/uatu/config"
	"github.com/uatu/internal/storage/postgres"
	"go.uber.org/zap"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed blockchains, tokens, dexes, and pools into the database",
	Long: "Loads the embedded chain definitions, discovers pools via the " +
		"configured RPC endpoints, and upserts blockchains, tokens, dexes, and " +
		"pools into the database.",
	RunE: runSeed,
}

func init() {
	rootCmd.AddCommand(seedCmd)
}

func runSeed(cmd *cobra.Command, _ []string) error {
	cfg, err := config.InitializeConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return fmt.Errorf("create logger: %w", err)
	}
	defer func() { _ = logger.Sync() }()

	db, err := postgres.DbConnection(*cfg, logger)
	if err != nil {
		return fmt.Errorf("connect to database: %w", err)
	}
	defer func() { _ = db.Close() }()

	chains, err := uatu.Load()
	if err != nil {
		return fmt.Errorf("load blockchains: %w", err)
	}

	tokens, dexes, pools := getChainAssets(chains)
	log.Printf("Discovered %d pools", len(pools))

	ctx := cmd.Context()
	if err := createBlockchains(ctx, db, chains); err != nil {
		return fmt.Errorf("seed blockchains: %w", err)
	}
	log.Printf("Seeded %d blockchains", len(chains))

	if err := createTokens(ctx, db, tokens); err != nil {
		return fmt.Errorf("seed tokens: %w", err)
	}
	log.Printf("Seeded %d tokens", len(tokens))

	if err := createDexes(ctx, db, dexes); err != nil {
		return fmt.Errorf("seed dexes: %w", err)
	}
	log.Printf("Seeded %d dexes", len(dexes))

	if err := createPools(ctx, db, pools); err != nil {
		return fmt.Errorf("seed pools: %w", err)
	}
	log.Printf("Seeded %d pools", len(pools))

	return nil
}
