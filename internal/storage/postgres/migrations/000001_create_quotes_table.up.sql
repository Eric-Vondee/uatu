CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE If NOT EXISTS quotes(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    quote_id  VARCHAR(255)  NOT NULL UNIQUE,
    amount_in NUMERIC(78,0) NOT NULL CHECK (amount_in >= 0),
    amount_out NUMERIC(78,0) NOT NULL CHECK (amount_out >= 0),
    amount_out_minimum NUMERIC(78,0) NOT NULL DEFAULT 0 CHECK (amount_out_minimum >= 0),
    slippage_bps INTEGER NOT NULL DEFAULT 0,
    amount_in_float  DOUBLE PRECISION NOT NULL,
    amount_out_float DOUBLE PRECISION NOT NULL,
    origin_chain_id VARCHAR(125) NOT NULL,
    origin_chain VARCHAR(125) NOT NULL,
    destination_chain_id VARCHAR(125) NOT NULL,
    destination_chain VARCHAR(125) NOT NULL,
    wallet_address  VARCHAR(255)  NOT NULL,
    recipient_address VARCHAR(255)  NOT NULL,
    token_in jsonb NOT NULL DEFAULT '{}'::jsonb,
    token_out jsonb NOT NULL DEFAULT '{}'::jsonb,
    pair_address VARCHAR(255),
    hash VARCHAR(255),
    explorer_url VARCHAR(512),
    steps jsonb NOT NULL DEFAULT '[]'::jsonb,
    status VARCHAR(100) NOT NULL DEFAULT 'pending',
    route jsonb NOT NULL DEFAULT '{}'::jsonb,
    deadline NUMERIC(78,0) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_quotes_quote_id ON quotes(quote_id);
