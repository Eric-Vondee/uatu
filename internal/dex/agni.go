package dex

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/uatu"
	"github.com/uatu/internal/contracts/agni"
)

type AgniSlot0 struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint32
	Unlocked                   bool
}

type AgniV3Pool struct {
	Token0 common.Address
	Token1 common.Address
	Fee    *big.Int
	Slot0  AgniSlot0
}

func (c *Client) GetAgniV3Pool(ctx context.Context, address common.Address) (*AgniV3Pool, error) {
	poolContract, err := agni.NewV3AgniPoolCaller(address, c.client)
	if err != nil {
		return nil, fmt.Errorf("could not bind Agni v3 pool: %w", err)
	}
	opts := &bind.CallOpts{Context: ctx}
	token0, err := poolContract.Token0(opts)
	if err != nil {
		return nil, fmt.Errorf("could not get Agni pool token0: %w", err)
	}
	token1, err := poolContract.Token1(opts)
	if err != nil {
		return nil, fmt.Errorf("could not get Agni pool token1: %w", err)
	}
	fee, err := poolContract.Fee(opts)
	if err != nil {
		return nil, fmt.Errorf("could not get Agni pool fee: %w", err)
	}
	slot0, err := poolContract.Slot0(opts)
	if err != nil {
		return nil, fmt.Errorf("could not get Agni pool slot0: %w", err)
	}

	return &AgniV3Pool{
		Token0: token0,
		Token1: token1,
		Fee:    fee,
		Slot0: AgniSlot0{
			SqrtPriceX96:               slot0.SqrtPriceX96,
			Tick:                       slot0.Tick,
			ObservationIndex:           slot0.ObservationIndex,
			ObservationCardinality:     slot0.ObservationCardinality,
			ObservationCardinalityNext: slot0.ObservationCardinalityNext,
			FeeProtocol:                slot0.FeeProtocol,
			Unlocked:                   slot0.Unlocked,
		},
	}, nil
}

func encodeAgniExactInputSingle(
	recipient common.Address,
	amountIn, amountOutMinimum, fee *big.Int,
	tokenIn, tokenOut common.Address,
	deadline *big.Int,
) ([]byte, error) {
	routerABI, err := agni.V3AgniRouterMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("could not parse Agni router ABI: %w", err)
	}
	return routerABI.Pack("exactInputSingle", agni.ISwapRouterExactInputSingleParams{
		TokenIn:           tokenIn,
		TokenOut:          tokenOut,
		Fee:               fee,
		Recipient:         recipient,
		Deadline:          deadline,
		AmountIn:          amountIn,
		AmountOutMinimum:  amountOutMinimum,
		SqrtPriceLimitX96: big.NewInt(0),
	})
}

func (c *Client) Agni(ctx context.Context, d uatu.IDexRequest) (*uatu.IDexResponse, error) {
	tokenIn := d.TokenIn
	routerAddress := uatu.FormatEvmAddress(d.Dex.V3RouterAddress)
	erc20Allowance, err := c.getERC20Allowance(ctx, tokenIn, d.WalletAddress, routerAddress)
	if err != nil {
		return nil, err
	}

	var encodedERC20TokenApproval []byte
	if erc20Allowance.Cmp(d.AmountIn) <= 0 {
		encodedERC20TokenApproval, err = encodeERC20Token(d.AmountIn, routerAddress)
		if err != nil {
			return nil, err
		}
	}

	pool, err := c.GetAgniV3Pool(ctx, d.PairAddress)
	if err != nil {
		return nil, err
	}
	if tokenIn != pool.Token0 && tokenIn != pool.Token1 {
		return nil, fmt.Errorf("token %s is not in pool %s", d.TokenIn, d.PairAddress)
	}

	quoterAddress := uatu.FormatEvmAddress(d.Dex.V3QuoterAddress)
	amountOut, err := c.getV3AmountOut(d.AmountIn, quoterAddress, tokenIn, d.TokenOut, pool.Fee)
	if err != nil {
		return nil, err
	}
	amountOutMin := applySlippage(amountOut, d.SlippageBps)

	swapCallData, err := encodeAgniExactInputSingle(
		d.WalletAddress,
		d.AmountIn,
		amountOutMin,
		pool.Fee,
		tokenIn,
		d.TokenOut,
		big.NewInt(time.Now().Add(SwapDeadline).Unix()),
	)
	if err != nil {
		return nil, fmt.Errorf("could not encode Agni exactInputSingle calldata: %w", err)
	}

	return &uatu.IDexResponse{
		AmountIn:             d.AmountIn,
		AmountOut:            amountOut,
		AmountOutMinimum:     amountOutMin,
		EncodedData:          swapCallData,
		EncodedERC20Approval: encodedERC20TokenApproval,
		Dex:                  d.Dex,
		RouterAddress:        routerAddress,
		PairAddress:          d.PairAddress,
		Route:                routeWithFee(DexRoutes["agni"], poolFeeAmount(d.AmountIn, pool.Fee)),
	}, nil
}
