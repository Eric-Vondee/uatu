package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/uatu/config"

	"github.com/alexlast/bunzap"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bunotel"
	"go.uber.org/zap"
)

func DbConnection(cfg config.Config, logger *zap.Logger) (*bun.DB, error) {
	pgdb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(cfg.Database.Postgres.DSN),
	))
	pgdb.SetMaxOpenConns(25)
	pgdb.SetMaxIdleConns(10)
	pgdb.SetConnMaxLifetime(30 * time.Minute)
	pgdb.SetConnMaxIdleTime(5 * time.Minute)

	db := bun.NewDB(pgdb, pgdialect.New())

	db.WithQueryHook(
		bunzap.NewQueryHook(
			bunzap.QueryHookOptions{
				Logger: logger,
			}))

	db.WithQueryHook(
		bunotel.NewQueryHook(
			bunotel.WithDBName("uatu.database")))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping PostgreSQL: %w", err)
	}
	return db, nil
}
