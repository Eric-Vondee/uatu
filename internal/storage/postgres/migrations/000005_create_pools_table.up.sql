CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS pools (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    dex_name VARCHAR(100) NOT NULL,
    symbol VARCHAR(100) NOT NULL,
    pair_address VARCHAR(255) UNIQUE NOT NULL,
    base_token_address VARCHAR(255) NOT NULL,
    quote_token_address VARCHAR(255) NOT NULL,
    base_token JSONB NOT NULL DEFAULT '{}'::jsonb,
    quote_token JSONB NOT NULL DEFAULT '{}'::jsonb,
    chain_id BIGINT NOT NULL,
    pool_fee BIGINT,
    pool_type VARCHAR(100) NOT NULL,
    tick_spacing BIGINT,
    market_cap NUMERIC(78,0),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_pools_pair_address ON pools(pair_address);
CREATE INDEX IF NOT EXISTS idx_pools_chain_id ON pools(chain_id);
CREATE INDEX IF NOT EXISTS idx_pools_token_pair ON pools(base_token_address, quote_token_address, chain_id)
