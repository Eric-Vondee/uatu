DROP INDEX IF EXISTS uq_pools_chain_pair_address;

ALTER TABLE pools
    ADD CONSTRAINT pools_pair_address_key UNIQUE (pair_address);
