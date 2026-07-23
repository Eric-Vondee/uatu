package uatu

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

//go:embed internal/blockchains/*.json
var files embed.FS

type Token struct {
	bun.BaseModel `bun:"table:tokens,alias:tokens"`

	ID           uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"-"`
	Name         string    `bun:"name,notnull" json:"name"`
	Symbol       string    `bun:"symbol,notnull" json:"symbol"`
	Address      string    `bun:"address,notnull" json:"address"`
	Decimals     uint8     `bun:"decimals,notnull" json:"decimals"`
	Logo         string    `bun:"logo,notnull" json:"logo"`
	Slug         string    `bun:"slug,notnull" json:"slug"`
	BlockchainID uint      `bun:"blockchain_id,notnull" json:"blockchainId"`
	CreatedAt    time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"-"`
	UpdatedAt    time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"-"`
}

type Chain struct {
	bun.BaseModel `bun:"table:blockchains,alias:blockchains"`

	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"-"`
	Name          string    `bun:"name,notnull" json:"name"`
	Symbol        string    `bun:"symbol,notnull" json:"symbol"`
	ChainID       uint      `bun:"chain_id,notnull" json:"chainId"`
	BlockExplorer string    `bun:"block_explorer_url,notnull" json:"blockExplorer"`
	RpcUrl        string    `bun:"-" json:"rpcUrl"`
	NativeToken   string    `bun:"native_token,notnull" json:"nativeToken"`
	Tokens        []Token   `bun:"tokens,type:jsonb,notnull" json:"tokens"`
	Dex           []Dex     `bun:"dex,type:jsonb,notnull" json:"dex"`
	Slug          string    `bun:"slug,notnull" json:"slug"`
	Ecosystem     string    `bun:"ecosystem,notnull" json:"ecosystem"`
	CreatedAt     time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"-"`
	UpdatedAt     time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"-"`
}

type QueryOptions struct {
	TokenIn     string
	TokenOut    string
	PairAddress string
	ChainID     uint
	DexName     string
	Slug        string
}
type ChainRepository interface {
	GetBlockchains(ctx context.Context) ([]Chain, error)
	GetBlockchain(ctx context.Context, opts QueryOptions) (Chain, error)
	GetTokens(ctx context.Context, opts QueryOptions) ([]Token, error)
	GetPools(ctx context.Context, opts QueryOptions) ([]Pool, error)
	GetDex(ctx context.Context, opts QueryOptions) (Dex, error)
}

// blockchainsDir is the embedded directory holding one JSON file per chain.
const blockchainsDir = "internal/blockchains"

// Load parses every embedded chain definition.
func Load() ([]Chain, error) {
	entries, err := files.ReadDir(blockchainsDir)
	if err != nil {
		return nil, err
	}
	chains := make([]Chain, 0, len(entries))
	for _, entry := range entries {
		data, err := files.ReadFile(blockchainsDir + "/" + entry.Name())
		if err != nil {
			return nil, err
		}
		var c Chain
		if err := json.Unmarshal(data, &c); err != nil {
			return nil, fmt.Errorf("parsing %s: %w", entry.Name(), err)
		}
		chains = append(chains, c)
	}
	return chains, nil
}
