# uatu

A DEX aggregation and swap-quote API for EVM chains. It maintains a catalogue of
blockchains, tokens, DEXes and liquidity pools across 18 networks, and serves
executable swap quotes built from on-chain AMM data.

## What it does

- **Quotes** — given a token pair, amount and chain, it resolves the relevant
  pool, prices the swap against the DEX's on-chain contracts, and returns a
  quote with the encoded calldata and approval steps needed to execute it.
- **Catalogue** — exposes the blockchains, tokens and pools it knows about, so
  clients can discover what's swappable where.
- **Discovery** — a seeding CLI walks each chain's token list, queries the v2 and
  v3 factories over RPC to find real pools, and persists what it finds.

Supported protocols include Uniswap v2/v3 (+ Universal Router), PancakeSwap,
CoW Swap, Velodrome, Aerodrome, Merchant Moe and Lithos. Which protocols are
available on a given chain is defined per chain in `internal/blockchains/`.

## Layout

```
cmd/               HTTP server entrypoint
cmd/cli/           Cobra CLI (database seeding)
config/            Viper config loading
swagger/           generated API spec (swag) — committed, do not hand-edit
server/            chi routes, handlers, OpenTelemetry setup
internal/blockchains/  chain definitions (JSON, embedded at build time)
internal/contracts/    generated Uniswap v2/v3 + Permit2 bindings
internal/dex/          on-chain pool and quote lookups
internal/storage/postgres/  bun repositories and SQL migrations
```

Domain models (`Chain`, `Token`, `Dex`, `Pool`, `Quote`) live in the root
`uatu` package.

## Requirements

- Go 1.26+
- PostgreSQL
- An RPC endpoint per chain you intend to support

## Configuration

Configuration comes from two places: a `.env` file and the process environment.
Both are read, and **environment variables win** where a key is defined in both.

The `.env` file is a convenience for local development and is entirely optional.
Anywhere the environment already supplies these values — a container, or a
systemd unit using `EnvironmentFile=` — no file needs to exist on disk.

| Variable | Purpose |
| --- | --- |
| `PORT` | HTTP listen port |
| `POSTGRES_DSN` | Postgres connection string |
| `<CHAIN>_RPC_URL` | RPC endpoint per chain (e.g. `ETHEREUM_RPC_URL`) |
| `OTEL_ENABLED` | Toggle OpenTelemetry export |
| `OTEL_ENDPOINT`, `OTEL_USE_TLS`, `OTEL_HEADERS` | OTLP exporter settings |

All fields are validated at startup; the process exits if any required value is
missing.

## Running

```sh
# Create the schema (safe to re-run; a no-op when already up to date)
go run ./cmd/cli migrate

# Seed blockchains, tokens, dexes and pools
go run ./cmd/cli seed

# Start the API server
go run ./cmd
```

To build binaries instead:

```sh
go build -o bin/uatu     ./cmd
go build -o bin/uatu-cli ./cmd/cli
```

Migrations live in `internal/storage/postgres/migrations` and are embedded into
the binary, so `migrate` needs only `POSTGRES_DSN` — no source tree and no
separate migration tool.

The seed step makes live RPC calls across every configured chain, so it takes a
while and logs non-fatal warnings for endpoints that don't respond. Run it
against a clean database — only the `blockchains` insert upserts, so re-running
can produce duplicate-key errors on the other tables.

## API

Swagger UI is served at `/swagger/` and the raw spec at `/swagger/doc.json`.

Regenerate the spec after changing any handler annotation:

```sh
go generate ./...
```

## Adding a chain

Drop a JSON definition into `internal/blockchains/`, add its `*_RPC_URL` to the
config struct and env, then re-run `uatu-cli seed`. Definitions are embedded
into the binary at build time, so a rebuild is required.
