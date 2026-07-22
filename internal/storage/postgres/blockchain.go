package postgres

import (
	"context"
	"fmt"

	"github.com/uatu"
	"github.com/uptrace/bun"
)

type chainRepo struct {
	db *bun.DB
}

func NewChainRepository(db *bun.DB) uatu.ChainRepository {
	return &chainRepo{
		db: db,
	}
}

func (c *chainRepo) GetBlockchains(ctx context.Context) ([]uatu.Chain, error) {
	var chains []uatu.Chain
	err := c.db.NewSelect().
		Model(&chains).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get blockchains: %w", err)
	}
	return chains, nil
}

func (c *chainRepo) GetBlockchain(ctx context.Context, opts uatu.QueryOptions) (uatu.Chain, error) {
	var chain uatu.Chain
	err := c.db.NewSelect().
		Model(&chain).
		Where("chain_id=?", opts.ChainID).
		Scan(ctx)
	if err != nil {
		return uatu.Chain{}, fmt.Errorf("could not get blockchains: %w", err)
	}
	return chain, nil
}

func (c *chainRepo) GetTokens(ctx context.Context) ([]uatu.Token, error) {
	var tokens []uatu.Token
	err := c.db.NewSelect().
		Model(&tokens).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get tokens: %w", err)
	}
	return tokens, nil
}

func (c *chainRepo) GetDex(ctx context.Context, opts uatu.QueryOptions) (uatu.Dex, error) {
	var dex uatu.Dex
	err := c.db.NewSelect().
		Model(&dex).
		Where("blockchain_id = ?", opts.ChainID).
		Where("slug = ?", opts.Slug).
		Scan(ctx)
	if err != nil {
		return uatu.Dex{}, fmt.Errorf("could not get dex: %w", err)
	}
	return dex, nil
}

func (c *chainRepo) GetPools(ctx context.Context, opts uatu.QueryOptions) ([]uatu.Pool, error) {
	var pools []uatu.Pool
	if err := c.db.NewSelect().
		Model(&pools).
		Where("chain_id = ?", opts.ChainID).
		Where(
			"(base_token_address = ? AND quote_token_address = ?)"+
				" OR (base_token_address = ? AND quote_token_address = ?)",
			opts.TokenIn, opts.TokenOut, opts.TokenOut, opts.TokenIn,
		).
		Scan(ctx); err != nil {
		return nil, fmt.Errorf("could not get pools: %w", err)
	}
	return pools, nil
}
