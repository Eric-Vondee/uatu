// @title uatu API
// @version 0.1.0
// @description DEX aggregation and swap-quote API for EVM chains.
// @description Maintains a catalogue of blockchains, tokens, DEXes and liquidity
// @description pools, and serves executable swap quotes built from on-chain AMM data.

// @BasePath /
// @schemes http https

// host is intentionally unset so the Swagger UI issues requests against
// whatever origin serves the page — the listen port is configuration, not a
// build-time constant.
package uatu
