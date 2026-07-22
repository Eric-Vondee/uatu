package postgres

import (
	"database/sql"
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

	db := bun.NewDB(pgdb, pgdialect.New())

	db.WithQueryHook(
		bunzap.NewQueryHook(
			bunzap.QueryHookOptions{
				Logger: logger,
			}))

	db.WithQueryHook(
		bunotel.NewQueryHook(
			bunotel.WithDBName("uatu.database")))

	return db, db.Ping()
}
