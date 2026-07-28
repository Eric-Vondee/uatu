package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/uatu"
	"github.com/uatu/internal/dex"
	"github.com/uptrace/bun"
)

func getChainAssets(chains []uatu.Chain) (
	[]uatu.Token,
	[]uatu.Dex,
	[]uatu.Pool,
) {
	tokens := make([]uatu.Token, 0)
	dexes := make([]uatu.Dex, 0)
	pools := make([]uatu.Pool, 0)
	for _, c := range chains {
		for _, token := range c.Tokens {
			token.BlockchainID = c.ChainID
			tokens = append(tokens, token)
		}
		dexes = append(dexes, chainDexes(c)...)
		if !hasSwapDex(c) {
			continue
		}
		client, err := dex.Provider(c.RpcUrl)
		if err != nil {
			log.Printf("skipping pools for %s: %v", c.Slug, err)
			continue
		}
		pools = append(pools, getPools(client, c)...)
	}
	return tokens, dexes, pools
}

const maxPoolLookups = 8

type poolJob struct {
	tokenIn, tokenOut uatu.Token
	d                 uatu.Dex
}

func getPools(client *dex.Client, chain uatu.Chain) []uatu.Pool {
	tokens := chain.Tokens
	jobs := make([]poolJob, 0)
	for i, tokenIn := range tokens {
		for _, tokenOut := range tokens[i+1:] {
			if uatu.FormatEvmAddress(tokenIn.Address) == (common.Address{}) ||
				uatu.FormatEvmAddress(tokenOut.Address) == (common.Address{}) {
				continue
			}
			for _, d := range chain.Dex {
				if d.Slug == "cowswap" {
					continue
				}
				jobs = append(jobs, poolJob{tokenIn: tokenIn, tokenOut: tokenOut, d: d})
			}
		}
	}

	var (
		mu    sync.Mutex
		wg    sync.WaitGroup
		sem   = make(chan struct{}, maxPoolLookups)
		pools = make([]uatu.Pool, 0)
	)
	for _, job := range jobs {
		wg.Add(1)
		go func(j poolJob) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			found := lookupPools(client, chain, j)
			if len(found) == 0 {
				return
			}
			mu.Lock()
			pools = append(pools, found...)
			mu.Unlock()
		}(job)
	}
	wg.Wait()
	return pools
}

// clPair is a concentrated-liquidity pair normalised across dexes. Uniswap and
// pancakeswap key their pools by fee tier, aerodrome by tick spacing, so either
// field may be unset depending on which factory answered.
type clPair struct {
	address     common.Address
	fee         *big.Int
	tickSpacing *big.Int
}

// lookupPools resolves one token pair against one dex's v2 and v3 factories.
func lookupPools(client *dex.Client, chain uatu.Chain, j poolJob) []uatu.Pool {
	pools := make([]uatu.Pool, 0, 2)
	if j.d.V2FactoryAddress != "" {
		factoryAddress := uatu.FormatEvmAddress(j.d.V2FactoryAddress)
		tokenIn := uatu.FormatEvmAddress(j.tokenIn.Address)
		tokenOut := uatu.FormatEvmAddress(j.tokenOut.Address)
		pair, err := client.GetV2Pair(factoryAddress, tokenIn, tokenOut)
		if err != nil {
			log.Printf("%s %s: v2 pair %s/%s: %v", chain.Slug, j.d.Slug, j.tokenIn.Symbol, j.tokenOut.Symbol, err)
		} else if pair != (common.Address{}) {
			pool := newPool(chain, j.d, j.tokenIn, j.tokenOut, pair.Hex(), "v2", 300, 0)
			pools = append(pools, pool)
		}
	}
	if j.d.V3FactoryAddress != "" {
		factoryAddress := uatu.FormatEvmAddress(j.d.V3FactoryAddress)
		tokenIn := uatu.FormatEvmAddress(j.tokenIn.Address)
		tokenOut := uatu.FormatEvmAddress(j.tokenOut.Address)
		var (
			pairs []clPair
			err   error
		)
		switch j.d.Slug {
		case "uniswap", "pancakeswap", "oku":
			var found []dex.V3Pair
			found, err = client.GetV3Pair(factoryAddress, tokenIn, tokenOut)
			for _, p := range found {
				pairs = append(pairs, clPair{address: p.PairAddress, fee: p.Fee})
			}
		case "aerodrome":
			var found []dex.AerodromePair
			found, err = client.GetAerodromePair(factoryAddress, tokenIn, tokenOut)
			for _, p := range found {
				pairs = append(pairs, clPair{
					address:     p.PairAddress,
					fee:         p.Fee,
					tickSpacing: p.TickSpacing,
				})
			}
		default:
			log.Printf("%s: no v3 factory lookup for dex %s", chain.Slug, j.d.Slug)
			return pools
		}
		if err != nil {
			log.Printf("%s %s: v3 pools %s/%s: %v", chain.Slug, j.d.Slug, j.tokenIn.Symbol, j.tokenOut.Symbol, err)
			return pools
		}
		for _, p := range pairs {
			pool := newPool(
				chain, j.d, j.tokenIn, j.tokenOut, p.address.Hex(), "v3",
				bigToUint(p.fee), bigToUint(p.tickSpacing),
			)
			pools = append(pools, pool)
		}
	}
	return pools
}

func bigToUint(v *big.Int) uint {
	if v == nil {
		return 0
	}
	return uint(v.Uint64())
}

func newPool(
	chain uatu.Chain,
	d uatu.Dex,
	base, quote uatu.Token,
	address, poolType string,
	fee, tickSpacing uint,
) uatu.Pool {
	return uatu.Pool{
		Name:              fmt.Sprintf("%s %s/%s", d.Name, base.Symbol, quote.Symbol),
		DexName:           strings.ToLower(d.Slug),
		Symbol:            fmt.Sprintf("%s/%s", base.Symbol, quote.Symbol),
		PairAddress:       address,
		BaseTokenAddress:  uatu.FormatEvmAddress(base.Address).String(),
		QuoteTokenAddress: uatu.FormatEvmAddress(quote.Address).String(),
		BaseToken:         base,
		QuoteToken:        quote,
		ChainID:           chain.ChainID,
		PoolFee:           fee,
		PoolType:          poolType,
		TickSpacing:       tickSpacing,
	}
}

func hasSwapDex(c uatu.Chain) bool {
	for _, d := range c.Dex {
		if d.Slug == "uniswap" || d.Slug == "pancakeswap" {
			return true
		}
	}
	return false
}

func chainDexes(c uatu.Chain) []uatu.Dex {
	dexes := make([]uatu.Dex, 0, len(c.Dex))
	for _, d := range c.Dex {
		d.BlockchainID = c.ChainID
		hasV2 := d.V2RouterAddress != "" || d.V2FactoryAddress != ""
		hasV3 := d.V3FactoryAddress != "" || d.V3RouterAddress != ""
		if hasV2 {
			v2 := d
			v2.Version = "v2"
			dexes = append(dexes, v2)
		}
		if hasV3 {
			v3 := d
			v3.Version = "v3"
			dexes = append(dexes, v3)
		}
		// dexes without AMM contracts (e.g. cowswap) get a single unversioned row
		if !hasV2 && !hasV3 {
			dexes = append(dexes, d)
		}
	}
	return dexes
}

func createBlockchains(ctx context.Context, db *bun.DB, chains []uatu.Chain) error {
	if len(chains) == 0 {
		return nil
	}
	_, err := db.NewInsert().
		Model(&chains).
		On("CONFLICT (chain_id) DO UPDATE").
		Set("name = EXCLUDED.name").
		Set("symbol = EXCLUDED.symbol").
		Set("block_explorer_url = EXCLUDED.block_explorer_url").
		Set("native_token = EXCLUDED.native_token").
		Set("tokens = EXCLUDED.tokens").
		Set("dex = EXCLUDED.dex").
		Set("slug = EXCLUDED.slug").
		Set("ecosystem = EXCLUDED.ecosystem").
		Set("updated_at = CURRENT_TIMESTAMP").
		Exec(ctx)
	return err
}

func createTokens(ctx context.Context, db *bun.DB, tokens []uatu.Token) error {
	if len(tokens) == 0 {
		return nil
	}
	_, err := db.NewInsert().
		Model(&tokens).
		Exec(ctx)
	return err
}

func createDexes(ctx context.Context, db *bun.DB, dexes []uatu.Dex) error {
	if len(dexes) == 0 {
		return nil
	}
	_, err := db.NewInsert().
		Model(&dexes).
		Exec(ctx)
	return err
}

func createPools(ctx context.Context, db *bun.DB, pools []uatu.Pool) error {
	if len(pools) == 0 {
		return nil
	}
	_, err := db.NewInsert().
		Model(&pools).
		Exec(ctx)
	return err
}
