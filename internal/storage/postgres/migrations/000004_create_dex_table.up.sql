CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS dexes(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(100) NOT NULL,
    version VARCHAR(20) NOT NULL,
    blockchain_id BIGINT NOT NULL REFERENCES blockchains(chain_id),
    v2_factory_address VARCHAR(255),
    v2_router_address VARCHAR(255),
    v3_factory_address VARCHAR(255),
    v3_router_address VARCHAR(255),
    v3_quoter_address VARCHAR(255),
    permit2_address VARCHAR(255),
    settlement_address VARCHAR(255),
    universal_router_address VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE INDEX IF NOT EXISTS idx_dexes_chain ON dexes(blockchain_id);
