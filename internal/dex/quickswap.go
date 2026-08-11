package dex

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/uatu"
	"github.com/uatu/internal/contracts"
)

func encodeQuickSwapExactInputSingle(
	recipient common.Address,
	amountIn, amountOutMin, deadline *big.Int,
	tokenIn, tokenOut common.Address,
) ([]byte, error) {
	routerABI, err := contracts.V3QuickSwapRouterMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("could not parse QuickSwap V3 router ABI: %w", err)
	}
	params := contracts.ISwapRouterExactInputSingleParams{
		TokenIn:          tokenIn,
		TokenOut:         tokenOut,
		Recipient:        recipient,
		Deadline:         deadline,
		AmountIn:         amountIn,
		AmountOutMinimum: amountOutMin,
		LimitSqrtPrice:   big.NewInt(0),
	}
	calldata, err := routerABI.Pack("exactInputSingle", params)
	if err != nil {
		return nil, fmt.Errorf("could not encode QuickSwap exactInputSingle calldata: %w", err)
	}
	return calldata, nil
}

func (c *Client) quickSwapV3(ctx context.Context, d uatu.IDexRequest) (*uatu.IDexResponse, error) {
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

	var encodedERC20TokenApproval []byte
	if erc20Allowance.Cmp(d.AmountIn) <= 0 {
		encodedERC20TokenApproval, err = encodeERC20Token(d.AmountIn, routerAddress)
		if err != nil {
			return nil, err
		}
	}
	pool, err := c.GetV3Pool(d.PairAddress)
	if err != nil {
		return nil, err
	}
	if tokenIn != pool.Token0 && tokenIn != pool.Token1 {
		return nil, fmt.Errorf("token %s is not in pool %s", d.TokenIn, d.PairAddress)
	}

	amountOut, err := c.getV3AmountOut(d.AmountIn, quoterAddress, tokenIn, d.TokenOut, pool.Fee)
	if err != nil {
		return nil, err
	}

	swapCallData, err := encodeQuickSwapExactInputSingle(
		d.WalletAddress,
		d.AmountIn, amountOut, deadline,
		tokenIn, d.TokenOut,
	)
	if err != nil {
		return nil, err
	}
	return &uatu.IDexResponse{
		AmountIn:             d.AmountIn,
		AmountOut:            amountOut,
		EncodedData:          swapCallData,
		EncodedERC20Approval: encodedERC20TokenApproval,
		Dex:                  d.Dex,
		RouterAddress:        routerAddress,
		PairAddress:          d.PairAddress,
	}, nil
}

func (c *Client) QuickSwap(ctx context.Context, d uatu.IDexRequest) (*uatu.IDexResponse, error) {
	switch d.PoolType {
	case "v2":
		return c.swapV2(ctx, d)
	case "v3":
		return c.quickSwapV3(ctx, d)
	default:
		return nil, fmt.Errorf("unsupported pool type: %s", d.PoolType)
	}
}
