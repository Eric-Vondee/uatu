---
name: Bug report
about: Report incorrect behaviour in the API, quotes, seeding or CLI
title: "[BUG] "
labels: bug
assignees: ""
---

## Describe the bug

<!-- What went wrong? One or two sentences. -->

## Expected behaviour

<!-- What should have happened instead? -->

## Actual behaviour

<!-- What actually happened — wrong price, HTTP 500, empty pool list, reverted swap, job crash. -->

## Where does it happen?

- [ ] `POST /quotes` — single quote
- [ ] `POST /quotes/routes` — route search
- [ ] `GET /blockchains` (or `/tokens`, `/pools`, `/dex`)
- [ ] Seeding / discovery CLI (`uatu-cli seed`)
- [ ] Migrations (`uatu-cli migrate`)
- [ ] Server startup / config
- [ ] Other:

## Reproduction

<!--
  A reproducible request beats a description. Redact API keys and RPC URLs.
-->

**Request**

```http
POST /quotes HTTP/1.1
Content-Type: application/json

{ }
```

**Response**

```json

```

**Steps**

1.
2.
3.

## On-chain context

<!-- Delete this section for bugs that are not chain-specific (config, startup, migrations). -->

- **Chain / chainId**:
- **DEX / protocol**: <!-- Uniswap, SushiSwap, PancakeSwap, QuickSwap, Aerodrome, Pharoah, Agni Finance, OkuTrade, CowSwap -->
- **Token pair (addresses, not symbols)**:
- **Pool address** (if known):
- **Amount in / out**:
- **Block number** (if the result is block-sensitive):
- **Expected price vs. returned price**: <!-- and where you sourced the expected value -->

## Environment

- **uatu version / commit SHA**:
- **Go version**: <!-- `go version`, repo targets 1.26+ -->
- **How it was run**: <!-- local `make run`, Docker, deployed -->
- **RPC provider** for the affected chain: <!-- e.g. Alchemy, Infura, self-hosted — no URLs/keys -->
- **Postgres version**:

## Impact

- [ ] Returns wrong or unsafe quote data (could cause a bad swap)
- [ ] Endpoint errors or hangs
- [ ] Degraded but usable
- [ ] Cosmetic / docs

## Additional context

<!-- Regression? When did it last work? Related issues or PRs? -->
