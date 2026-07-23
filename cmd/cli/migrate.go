package main

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	// registers the "postgres"/"postgresql" database driver with migrate
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/spf13/cobra"
	"github.com/uatu"
	"github.com/uatu/config"
	"go.uber.org/zap"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Apply pending database migrations",
	Long: "Applies any migrations that have not yet run against the configured " +
		"database. Migrations are embedded in the binary, so no source tree is " +
		"required. Safe to re-run: it is a no-op when already up to date.",
	RunE: runMigrateCmd,
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func runMigrateCmd(_ *cobra.Command, _ []string) error {
	cfg, err := config.InitializeConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return fmt.Errorf("create logger: %w", err)
	}
	defer func() { _ = logger.Sync() }()

	return runMigrations(logger, cfg)
}

func runMigrations(logger *zap.Logger, cfg *config.Config) error {
	if cfg.Database.Postgres.DSN == "" {
		return errors.New("POSTGRES_DSN is not set")
	}

	d, err := iofs.New(uatu.Migrations, uatu.MigrationsDir)
	if err != nil {
		logger.Error("could not set up embedded migrations", zap.Error(err))
		return err
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, cfg.Database.Postgres.DSN)
	if err != nil {
		logger.Error("could not set up migrations", zap.Error(err))
		return err
	}
	defer func() {
		if srcErr, dbErr := m.Close(); srcErr != nil || dbErr != nil {
			logger.Error("could not close migrations", zap.Error(errors.Join(srcErr, dbErr)))
		}
	}()

	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		logger.Info("no new migration to run")
		return nil
	}

	if err != nil {
		logger.Error("could not run migrations", zap.Error(err))
		return err
	}

	logger.Info("migrations successful")
	return nil
}
