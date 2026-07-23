package uatu

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Pool struct {
	ID                uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"-"`
	Name              string    `bun:"name,notnull" json:"name"`
	DexName           string    `bun:"dex_name,notnull" json:"dexName"`
	Symbol            string    `bun:"symbol,notnull" json:"symbol"`
	PairAddress       string    `bun:"pair_address,notnull" json:"pairAddress"`
	BaseTokenAddress  string    `bun:"base_token_address,notnull" json:"baseTokenAddress"`
	QuoteTokenAddress string    `bun:"quote_token_address,notnull" json:"quoteTokenAddress"`
	BaseToken         Token     `bun:"base_token,type:jsonb,notnull" json:"baseToken"`
	QuoteToken        Token     `bun:"quote_token,type:jsonb,notnull" json:"quoteToken"`
	ChainID           uint      `bun:"chain_id,notnull" json:"-"`
	PoolFee           uint      `bun:"pool_fee,nullzero" json:"poolFee,omitempty"`
	PoolType          string    `bun:"pool_type,notnull" json:"poolType,omitempty"`
	TickSpacing       uint      `bun:"tick_spacing,nullzero" json:"tickSpacing,omitempty"`
	MarketCap         string    `bun:"market_cap,nullzero" json:"marketCap,omitempty"`
	CreatedAt         time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"-"`
	UpdatedAt         time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"-"`

	bun.BaseModel `bun:"table:pools,alias:pools" json:"-"`
}
