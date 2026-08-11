CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS blockchains (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    symbol VARCHAR(100) NOT NULL,
    chain_id BIGINT UNIQUE NOT NULL,
    block_explorer_url VARCHAR(255) NOT NULL,
    blockchain_logo VARCHAR(255),
    native_token VARCHAR(255) NOT NULL,
    tokens JSONB NOT NULL DEFAULT '[]'::jsonb,
    dex JSONB NOT NULL DEFAULT '[]'::jsonb,
    slug VARCHAR(255) NOT NULL,
    ecosystem VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
