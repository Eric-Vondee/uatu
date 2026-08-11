package uatu

import (
	"context"
	"math/big"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/uptrace/bun"
)

type TransactionStatus string

const (
	Pending   TransactionStatus = "pending"
	Completed TransactionStatus = "completed"
	Failed    TransactionStatus = "failed"
)

// Actions is a single transaction the client must submit to execute a quote.
type Actions struct {
	Message string `json:"message"`
	Amount  string `json:"amount"`
	Data    string `json:"data,omitempty"`
	Spender string `json:"spender,omitempty"`
	From    string `json:"from"`
	To      string `json:"to"`
	ChainID uint   `json:"chainId,omitempty"`
}

type QuoteRequest struct {
	Amount           decimal.Decimal `json:"amount" validate:"required" swaggertype:"string"`
	Chain            string          `json:"chain" validate:"required"`
	ChainID          uint            `json:"chainId" validate:"required"`
	RecipientAddress string          `json:"recipientAddress" validate:"required"`
	TokenIn          string          `json:"tokenIn" validate:"required"`
	TokenOut         string          `json:"tokenOut" validate:"required"`
	GenericRequest
}

type QuoteResponse struct {
	Quote
}

// RouteQuote is a non-persisted quote option for a direct DEX route.
type RouteQuote struct {
	AmountIn  string `json:"amountIn"`
	AmountOut string `json:"amountOut"`
	TokenIn   Token  `json:"tokenIn"`
	TokenOut  Token  `json:"tokenOut"`
	Route     Route  `json:"route"`
}

type Quote struct {
	ID                 uuid.UUID         `bun:"type:uuid,default:uuid_v4(),pk" json:"-"`
	QuoteID            string            `bun:"quote_id,notnull,unique" json:"quoteId"`
	AmountIn           string            `bun:"amount_in,notnull,type:numeric(78,0)" json:"amountIn"`
	AmountOut          string            `bun:"amount_out,notnull,type:numeric(78,0)" json:"amountOut"`
	AmountInFloat      decimal.Decimal   `bun:"amount_in_float" json:"amountInFloat" swaggertype:"string"`
	AmountOutFloat     decimal.Decimal   `bun:"amount_out_float" json:"amountOutFloat" swaggertype:"string"`
	OriginChain        string            `bun:"origin_chain,notnull" json:"originChain"`
	OriginChainId      uint              `bun:"origin_chain_id,notnull" json:"originChainId"`
	DestinationChain   string            `bun:"destination_chain,notnull" json:"destinationChain"`
	DestinationChainId uint              `bun:"destination_chain_id,notnull" json:"destinationChainId"`
	WalletAddress      string            `bun:"wallet_address,notnull" json:"walletAddress"`
	RecipientAddress   string            `bun:"recipient_address,notnull" json:"recipientAddress"`
	TokenIn            Token             `bun:"type:jsonb,notnull" json:"tokenIn"`
	TokenOut           Token             `bun:"type:jsonb,notnull" json:"tokenOut"`
	PairAddress        string            `bun:"type:pair_address" json:"pairAddress"`
	Hash               string            `bun:"hash" json:"hash,omitempty"`
	ExplorerUrl        string            `bun:"explorer_url" json:"explorerUrl,omitempty"`
	Steps              []Actions         `bun:"steps,notnull" json:"steps"`
	Status             TransactionStatus `bun:"type:varchar(100),default:'pending',notnull" json:"status" swaggertype:"string" enums:"pending,completed,failed"`
	Route              Route             `bun:"type:jsonb,notnull" json:"route"`
	Deadline           *big.Int          `bun:"deadline,notnull" json:"deadline" swaggertype:"integer"`
	CreatedAt          time.Time         `bun:"created_at,notnull" json:"createdAt"`
	UpdatedAt          time.Time         `bun:"updated_at,notnull" json:"updatedAt"`

	bun.BaseModel `bun:"table:quotes,alias:quotes" json:"-"`
}

type QuoteRepository interface {
	Create(ctx context.Context, quote *Quote) error
	Get(ctx context.Context, id string) (*Quote, error)
}
