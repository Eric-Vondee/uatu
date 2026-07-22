package postgres

import (
	"context"
	"fmt"

	"github.com/uatu"
	"github.com/uptrace/bun"
)

type quoteRepo struct {
	db *bun.DB
}

func NewQuoteRepository(db *bun.DB) uatu.QuoteRepository {
	return &quoteRepo{
		db: db,
	}
}

func (q *quoteRepo) Create(ctx context.Context, quote *uatu.Quote) error {
	_, err := q.db.NewInsert().
		Model(quote).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to create quote")
	}
	return err
}

func (q *quoteRepo) Get(ctx context.Context, id string) (*uatu.Quote, error) {
	var quote uatu.Quote
	err := q.db.NewSelect().
		Model(&quote).
		Where("quote_id = ?", id).
		Scan(ctx, &quote)
	if err != nil {
		return nil, err
	}
	return &quote, nil
}
