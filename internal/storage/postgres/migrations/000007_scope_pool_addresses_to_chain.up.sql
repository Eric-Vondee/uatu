ALTER TABLE pools
    DROP CONSTRAINT IF EXISTS pools_pair_address_key;

CREATE UNIQUE INDEX IF NOT EXISTS uq_pools_chain_pair_address
    ON pools (chain_id, LOWER(pair_address));
