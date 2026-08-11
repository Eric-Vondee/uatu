package dex

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/uatu"
	"github.com/uatu/internal/contracts/aerodrome"
)

type AerodromePool struct {
	Token0      common.Address
	Token1      common.Address
	TickSpacing *big.Int
	Fee         *big.Int
	Stable      bool
	Slot0       CLSlot0
	Liquidity   *big.Int
}

type CLSlot0 struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	Unlocked                   bool
}

type AerodromePair struct {
	PairAddress common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}

func (c *Client) GetAerodromePair(address, token0, token1 common.Address) ([]AerodromePair, error) {
	tickSpacings := []int64{1, 50, 100, 200, 2000}
	factoryContract, err := aerodrome.NewAerodromeFactoryCaller(address, c.client)
	if err != nil {
		return nil, fmt.Errorf("could not bind v2 factory: %w", err)
	}
	pairs := make([]AerodromePair, 0, len(tickSpacings))
	for _, tick := range tickSpacings {
		tickSpacing := big.NewInt(tick)
		pair, err := factoryContract.GetPool(&bind.CallOpts{}, token0, token1, tickSpacing)
		if err != nil {
			return nil, fmt.Errorf("could not get v3 pool for fee %d: %w", tick, err)
		}
		if pair == (common.Address{}) {
			continue
		}
		pairs = append(pairs, AerodromePair{PairAddress: pair, TickSpacing: tickSpacing})
	}
	return pairs, nil
}

func (c *Client) GetAerodromePool(ctx context.Context, address common.Address) (*AerodromePool, error) {
	poolContract, err := aerodrome.NewAerodromePoolCaller(address, c.client)
	if err != nil {
		return nil, fmt.Errorf("could not bind aerodrome pool: %w", err)
	}
	opts := &bind.CallOpts{Context: ctx}
	token0, err := poolContract.Token0(opts)
	if err != nil {
		return nil, fmt.Errorf("could not get token0: %w", err)
	}
	token1, err := poolContract.Token1(opts)
	if err != nil {
		return nil, fmt.Errorf("could not get token1: %w", err)
	}
	fee, err := poolContract.Fee(opts)
	if err != nil {
		return nil, fmt.Errorf("could not get fee: %w", err)
	}
	slot0, err := poolContract.Slot0(opts)
	if err != nil {
		return nil, fmt.Errorf("could not get slot0: %w", err)
	}
	liquidity, err := poolContract.Liquidity(opts)
	if err != nil {
		return nil, fmt.Errorf("could not get liquidity: %w", err)
	}
	tickSpacing, err := poolContract.TickSpacing(opts)
	if err != nil {
		return nil, fmt.Errorf("could not get tick spacing: %w", err)
	}

	return &AerodromePool{
		Token0:      token0,
		Token1:      token1,
		Fee:         fee,
		TickSpacing: tickSpacing,
		Slot0:       CLSlot0(slot0),
		Liquidity:   liquidity,
	}, nil
}

func (c *Client) getAerodromeAmountsOut(
	ctx context.Context,
	amountIn *big.Int,
	quoter, tokenIn, tokenOut common.Address,
	tickSpacing *big.Int,
) (*big.Int, error) {
	quoterContract, err := aerodrome.NewAerodromeQuoter(quoter, c.client)
	if err != nil {
		return nil, fmt.Errorf("could not bind aerodrome quoter: %w", err)
	}
	params := aerodrome.IQuoterV2QuoteExactInputSingleParams{
		TokenIn:           tokenIn,
		TokenOut:          tokenOut,
		AmountIn:          amountIn,
		TickSpacing:       tickSpacing,
		SqrtPriceLimitX96: big.NewInt(0),
	}
	raw := &aerodrome.AerodromeQuoterRaw{Contract: quoterContract}
	var out []interface{}
	if err := raw.Call(&bind.CallOpts{Context: ctx}, &out, "quoteExactInputSingle", params); err != nil {
		return nil, fmt.Errorf("could not quote v3 amount out: %w", err)
	}
	amountOut, ok := out[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("unexpected quoter amountOut type %T", out[0])
	}
	return amountOut, nil
}

func encodeExactInputSingle(
	recipient common.Address,
	amountIn, amountOutMin *big.Int,
	tokenIn, tokenOut common.Address,
	tickSpacing, deadline *big.Int,
) ([]byte, error) {
	routerABI, err := aerodrome.AerodromeRouterMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("could not parse cl swap router abi: %w", err)
	}
	params := aerodrome.ISwapRouterExactInputSingleParams{
		TokenIn:           tokenIn,
		TokenOut:          tokenOut,
		TickSpacing:       tickSpacing,
		Recipient:         recipient,
		Deadline:          deadline,
		AmountIn:          amountIn,
		AmountOutMinimum:  amountOutMin,
		SqrtPriceLimitX96: big.NewInt(0),
	}
	calldata, err := routerABI.Pack("exactInputSingle", params)
	if err != nil {
		return nil, fmt.Errorf("could not encode cl exactInputSingle calldata: %w", err)
	}
	return calldata, nil
}

func (c *Client) Aerodrome(ctx context.Context, d uatu.IDexRequest) (*uatu.IDexResponse, error) {
	tokenIn := d.TokenIn
	quoterAddress := uatu.FormatEvmAddress(d.Dex.V3QuoterAddress)
	routerAddress := uatu.FormatEvmAddress(d.Dex.V3RouterAddress)
	erc20Allowance, err := c.getERC20Allowance(
		ctx,
		tokenIn,
		d.WalletAddress,
		routerAddress,
	)
	if err != nil {
		return nil, err
	}
	deadline := big.NewInt(time.Now().Add(swapDeadline).Unix())

	var enodedERC20TokenApproval []byte
	if erc20Allowance.Cmp(d.AmountIn) <= 0 {
		enodedERC20TokenApproval, err = encodeERC20Token(d.AmountIn, routerAddress)
		if err != nil {
			return nil, err
		}
	}
	pool, err := c.GetAerodromePool(ctx, d.PairAddress)
	if err != nil {
		return nil, err
	}
	if tokenIn != pool.Token0 && tokenIn != pool.Token1 {
		return nil, fmt.Errorf("token %s is not in pool %s", d.TokenIn, d.PairAddress)
	}
	amountOut, err := c.getAerodromeAmountsOut(
		ctx,
		d.AmountIn, quoterAddress,
		tokenIn, d.TokenOut,
		pool.TickSpacing,
	)
	if err != nil {
		return nil, err
	}

	swapCallData, err := encodeExactInputSingle(
		d.WalletAddress,
		d.AmountIn, amountOut,
		tokenIn, d.TokenOut,
		pool.TickSpacing, deadline,
	)
	if err != nil {
		return nil, err
	}
	return &uatu.IDexResponse{
		AmountIn:             d.AmountIn,
		AmountOut:            amountOut,
		EncodedData:          swapCallData,
		EncodedERC20Approval: enodedERC20TokenApproval,
		Dex:                  d.Dex,
		RouterAddress:        routerAddress,
		PairAddress:          d.PairAddress,
		Route:                DexRoutes["aerodrome"],
	}, nil
}
