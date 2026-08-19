package dex

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/uatu"
	"github.com/uatu/internal/contracts/uniswap"
)

func encodeSushiSwapV2ExactInput(
	recipient common.Address,
	amountIn, amountOutMin, deadline *big.Int,
	tokenIn, tokenOut common.Address,
	wrapNativeInput, unwrapNativeOutput bool,
) ([]byte, error) {
	routerABI, err := uniswap.V2RouterMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("could not parse SushiSwap V2 router ABI: %w", err)
	}
	path := []common.Address{tokenIn, tokenOut}
	if wrapNativeInput {
		return routerABI.Pack("swapExactETHForTokens", amountOutMin, path, recipient, deadline)
	}
	method := "swapExactTokensForTokens"
	if unwrapNativeOutput {
		method = "swapExactTokensForETH"
	}
	return routerABI.Pack(method, amountIn, amountOutMin, path, recipient, deadline)
}

func (c *Client) SushiSwap(ctx context.Context, d uatu.IDexRequest) (*uatu.IDexResponse, error) {
	switch d.PoolType {
	case "v2":
	default:
		return nil, fmt.Errorf("unsupported SushiSwap pool type: %s", d.PoolType)
	}
	routerAddress := uatu.FormatEvmAddress(d.Dex.V2RouterAddress)
	if routerAddress == (common.Address{}) {
		return nil, fmt.Errorf("SushiSwap V2 router is not configured")
	}

	var approval []byte
	if !d.WrapNativeInput {
		allowance, err := c.getERC20Allowance(ctx, d.TokenIn, d.WalletAddress, routerAddress)
		if err != nil {
			return nil, err
		}
		if allowance.Cmp(d.AmountIn) <= 0 {
			approval, err = encodeERC20Token(d.AmountIn, routerAddress)
			if err != nil {
				return nil, err
			}
		}
	}

	pool, err := c.GetV2Pool(d.PairAddress)
	if err != nil {
		return nil, err
	}
	reserveIn, reserveOut := pool.Reserve0, pool.Reserve1
	if d.TokenIn == pool.Token1 {
		reserveIn, reserveOut = pool.Reserve1, pool.Reserve0
	} else if d.TokenIn != pool.Token0 {
		return nil, fmt.Errorf("token %s is not in pool %s", d.TokenIn, d.PairAddress)
	}
	amountOut, err := c.getV2AmountOut(ctx, d.AmountIn, reserveIn, reserveOut, routerAddress)
	if err != nil {
		return nil, err
	}
	amountOutMin := applySlippage(amountOut, d.SlippageBps)
	calldata, err := encodeSushiSwapV2ExactInput(
		d.WalletAddress, d.AmountIn, amountOutMin,
		big.NewInt(time.Now().Add(SwapDeadline).Unix()),
		d.TokenIn, d.TokenOut, d.WrapNativeInput, d.UnwrapNativeOutput,
	)
	if err != nil {
		return nil, fmt.Errorf("could not encode SushiSwap V2 swap: %w", err)
	}
	return &uatu.IDexResponse{
		AmountIn: d.AmountIn, AmountOut: amountOut, AmountOutMinimum: amountOutMin,
		EncodedData:          calldata,
		EncodedERC20Approval: approval, NativeValue: nativeValue(d.WrapNativeInput, d.AmountIn),
		Dex: d.Dex, RouterAddress: routerAddress, PairAddress: d.PairAddress,
		Route: routeWithPoolFee(DexRoutes["sushiswap"], d.AmountIn, d.PoolFee),
	}, nil
}
