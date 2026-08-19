<!--
  Title this PR like a Conventional Commit: `type(scope): short imperative summary`
  e.g. feat(quote): support exact-output quotes on Uniswap v3
  Allowed types: feat, fix, refactor, perf, chore, docs, test, build, ci, style
  See AGENTS.md for the full convention.
-->

## What & why

<!-- What does this change do, and why is it needed? Explain the *why* — the diff already shows the how. -->

Closes #

## Type of change

- [ ] `feat` — new feature or capability
- [ ] `fix` — bug fix
- [ ] `refactor` / `perf` — no behaviour change (or faster only)
- [ ] `chore` / `build` / `ci` — tooling, deps, config
- [ ] `docs` / `test` / `style`
- [ ] **Breaking change** — see "Breaking changes" below

## Areas touched

<!-- Tick what this PR reaches into. This is a routing hint for reviewers, not a quality gate. -->

- [ ] Quote / routing
- [ ] Chain or token configs
- [ ] HTTP API — handlers, routes, rate limiting
- [ ] Storage / migrations
- [ ] Contract bindings
- [ ] Jobs / seeding
- [ ] Docs / OpenAPI spec

## How it was verified

<!--
  Be specific: which chain, which pair, which endpoint. "Tested locally" tells a
  reviewer nothing. Paste the request/response or CLI output where it helps.
-->

- [ ] `go test ./...` passes locally
- [ ] `make vet` is clean
- [ ] `golangci-lint run` is clean (CI runs this on new issues only)
- [ ] Exercised against a live chain / RPC — chain(s):
- [ ] Not runtime-verifiable (docs, config-only) — explain why below

## Repo-specific checks

<!-- Delete any line that does not apply to this PR. -->

- [ ] **Swagger** — changed handler annotations, so I ran `make docs` and committed `swagger/` (`make docs-check` passes)
- [ ] **Migrations** — new migration is additive/reversible, and `make migrate` runs clean on a fresh database
- [ ] **Chain configs** — addresses checksummed and verified against a block explorer; decimals and `chainId` confirmed on-chain
- [ ] **Contract bindings** — regenerated with `go generate`, not hand-edited
- [ ] **Config** — new keys documented in the README and defaulted safely when unset
- [ ] **Dependencies** — `make tidy` run, and `go.mod` / `go.sum` changes are intentional

## Breaking changes

<!--
  Anything that changes an API response shape, an endpoint path, a config key,
  or on-chain behaviour is breaking. Describe the break and the migration path,
  and mark the commit with `!` or a `BREAKING CHANGE:` footer.
  Write "None" if there are none.
-->

None

## Risk & rollback

<!--
  What's the blast radius if this is wrong — bad quotes, failed swaps, a stuck
  seeding job? How do we roll it back (revert, config flip, re-run migration)?
-->

## Screenshots / sample payloads

<!-- API request+response pairs, CLI output, dashboards — anything that makes review faster. Optional. -->

## Additional context

<!-- Related PRs, design docs, third-party protocol docs, known follow-ups. -->

---

**Author checklist**

- [ ] PR title follows Conventional Commits and commits are split by logical change
- [ ] I self-reviewed the diff and it contains no unrelated changes, debug prints, or commented-out code
- [ ] Non-obvious logic is commented, and exported behaviour is documented
- [ ] Tests cover the new behaviour (or I explained above why they don't)
- [ ] No secrets, private keys, or RPC credentials are in the diff
