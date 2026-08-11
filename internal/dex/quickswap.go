package dex

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func encodeQuickSwapV2Path(
	recipient common.Address,
	amountIn, amountOutMin, deadline *big.Int,
	tokenIn, tokenOut common.Address,
) ([]byte, error) {
	routerABI, err := contracts.V2RouterMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("could not parse QuickSwap V2 router ABI: %w", err)
	}
	path := []common.Address{tokenIn, tokenOut}
	calldata, err := routerABI.Pack(
		"swapExactTokensForTokens",
		amountIn, amountOutMin, path, recipient, deadline,
	)
	if err != nil {
		return nil, fmt.Errorf("could not encode QuickSwap swapExactTokensForTokens calldata: %w", err)
	}
	return calldata, nil
}

func (c *Client) GetQuickSwapV3Pair(address, token0, token1 common.Address) ([]V3Pair, error) {
	factoryContract, err := contracts.NewV3QuickSwapFactoryCaller(address, c.client)
	if err != nil {
		return nil, fmt.Errorf("could not bind QuickSwap v3 factory: %w", err)
	}
	pair, err := factoryContract.PoolByPair(&bind.CallOpts{}, token0, token1)
	if err != nil {
		return nil, fmt.Errorf("could not get QuickSwap v3 pool: %w", err)
	}
	if pair == (common.Address{}) {
		return nil, nil
	}
	return []V3Pair{{PairAddress: pair}}, nil
}

func (c *Client) quickSwapV2(ctx context.Context, d uatu.IDexRequest) (*uatu.IDexResponse, error) {
	tokenIn := d.TokenIn
	routerAddress := uatu.FormatEvmAddress(d.Dex.V2RouterAddress)
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

	pool, err := c.GetV2Pool(d.PairAddress)
	if err != nil {
		return nil, err
	}

	reserveIn, reserveOut := pool.Reserve0, pool.Reserve1
	if tokenIn == pool.Token1 {
		reserveIn, reserveOut = pool.Reserve1, pool.Reserve0
	} else if tokenIn != pool.Token0 {
		return nil, fmt.Errorf("token %s is not in pool %s", d.TokenIn, d.PairAddress)
	}
	amountOut, err := c.getV2AmountOut(
		ctx, d.AmountIn,
		reserveIn, reserveOut,
		routerAddress,
	)
	if err != nil {
		return nil, err
	}

	swapCalldata, err := encodeQuickSwapV2Path(
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
		EncodedData:          swapCalldata,
		EncodedERC20Approval: enodedERC20TokenApproval,
		Dex:                  d.Dex,
		RouterAddress:        routerAddress,
		PairAddress:          d.PairAddress,
		Route:                DexRoutes["quickswap"],
	}, nil
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
		Route:                DexRoutes["quickswap"],
	}, nil
}

func (c *Client) QuickSwap(ctx context.Context, d uatu.IDexRequest) (*uatu.IDexResponse, error) {
	switch d.PoolType {
	case "v2":
		return c.quickSwapV2(ctx, d)
	case "v3":
		return c.quickSwapV3(ctx, d)
	default:
		return nil, fmt.Errorf("unsupported pool type: %s", d.PoolType)
	}
}
