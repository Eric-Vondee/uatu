package uatu

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Dex struct {
	ID                     uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"-"`
	Name                   string    `bun:"name,notnull" json:"name"`
	Slug                   string    `bun:"slug,notnull" json:"slug"`
	Version                string    `bun:"version,notnull" json:"version,omitempty"`
	BlockchainID           uint      `bun:"blockchain_id,notnull" json:"-"`
	V2RouterAddress        string    `bun:"v2_router_address" json:"v2RouterAddress,omitempty"`
	V2FactoryAddress       string    `bun:"v2_factory_address" json:"v2FactoryAddress,omitempty"`
	V3FactoryAddress       string    `bun:"v3_factory_address,nullzero" json:"v3FactoryAddress,omitempty"`
	V3RouterAddress        string    `bun:"v3_router_address,nullzero" json:"v3RouterAddress,omitempty"`
	V3QuoterAddress        string    `bun:"v3_quoter_address,nullzero" json:"v3QuoterAddress,omitempty"`
	Permit2Address         string    `bun:"permit2_address,nullzero" json:"permit2Address,omitempty"`
	SettlementAddress      string    `bun:"settlement_address,nullzero" json:"settlementAddress,omitempty"`
	UniversalRouterAddress string    `bun:"universal_router_address,nullzero" json:"universalRouterAddress,omitempty"`
	CreatedAt              time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"-"`
	UpdatedAt              time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"-"`

	bun.BaseModel `bun:"table:dexes,alias:dexes" json:"-"`
}

type IDexRequest struct {
	AmountIn                                      *big.Int
	PoolFee                                       uint
	ChainId                                       uint
	Dex                                           Dex
	TokenIn, TokenOut, WalletAddress, PairAddress common.Address
	PoolType                                      string
}

type IDexResponse struct {
	AmountIn, AmountOut                                       *big.Int
	EncodedData, EncodedERC20Approval, EncodedPermit2Approval []byte

	Dex           Dex
	RouterAddress common.Address
	PairAddress   common.Address
	RegisterOrder func(context.Context) error
}
