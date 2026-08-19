package dex

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/uatu"
	"github.com/uatu/internal/contracts/pharoah"
)

var pharoahV2FactoryABI = func() abi.ABI {
	parsed, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"address","name":"","type":"address"},{"internalType":"address","name":"","type":"address"},{"internalType":"bool","name":"","type":"bool"}],"name":"getPair","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`))
	if err != nil {
		panic(err)
	}
	return parsed
}()

type PharoahPair struct {
	PairAddress common.Address
	Stable      bool
}

func (c *Client) GetPharoahV2Pair(address, token0, token1 common.Address) ([]PharoahPair, error) {
	factoryContract := bind.NewBoundContract(address, pharoahV2FactoryABI, c.client, nil, nil)
	pairs := make([]PharoahPair, 0, 2)
	for _, stable := range []bool{false, true} {
		var out []interface{}
		if err := factoryContract.Call(&bind.CallOpts{}, &out, "getPair", token0, token1, stable); err != nil {
			return nil, fmt.Errorf("could not get v2 pair for stable %t: %w", stable, err)
		}
		pair, ok := out[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("unexpected v2 pair type %T", out[0])
		}
		if pair == (common.Address{}) {
			continue
		}
		pairs = append(pairs, PharoahPair{PairAddress: pair, Stable: stable})
	}
	return pairs, nil
}

type PharoahPool struct {
	Token0      common.Address
	Token1      common.Address
	TickSpacing *big.Int
	Fee         *big.Int
	Slot0       PharoahSlot0
	Liquidity   *big.Int
}

type PharoahSlot0 struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                *big.Int
	Unlocked                   bool
}

type PharoahV3Pair struct {
	PairAddress common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}

func (c *Client) GetPharoahV3Pair(address, token0, token1 common.Address) ([]PharoahV3Pair, error) {
	tickSpacings := []int64{1, 5, 10, 50, 100, 200, 500, 3000, 10000}
	factoryContract, err := pharoah.NewV3PharoahFactoryCaller(address, c.client)
	if err != nil {
		return nil, fmt.Errorf("could not bind v3 factory: %w", err)
	}
	pairs := make([]PharoahV3Pair, 0, len(tickSpacings))
	for _, tick := range tickSpacings {
		tickSpacing := big.NewInt(tick)
		pair, err := factoryContract.GetPool(&bind.CallOpts{}, token0, token1, tickSpacing)
		if err != nil {
			return nil, fmt.Errorf("could not get v3 pool for tick spacing %d: %w", tick, err)
		}
		if pair == (common.Address{}) {
			continue
		}
		pairs = append(pairs, PharoahV3Pair{PairAddress: pair, TickSpacing: tickSpacing})
	}
	return pairs, nil
}

func (c *Client) GetPharoahPool(ctx context.Context, address common.Address) (*PharoahPool, error) {
	poolContract, err := pharoah.NewV3PharoahPoolCaller(address, c.client)
	if err != nil {
		return nil, fmt.Errorf("could not bind pharoah pool: %w", err)
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

	return &PharoahPool{
		Token0:      token0,
		Token1:      token1,
		Fee:         fee,
		TickSpacing: tickSpacing,
		Slot0:       PharoahSlot0(slot0),
		Liquidity:   liquidity,
	}, nil
}

func (c *Client) getPharoahAmountsOut(
	ctx context.Context,
	amountIn *big.Int,
	quoter, tokenIn, tokenOut common.Address,
	tickSpacing *big.Int,
) (*big.Int, error) {
	quoterContract, err := pharoah.NewV3PharoahQuoter(quoter, c.client)
	if err != nil {
		return nil, fmt.Errorf("could not bind pharoah quoter: %w", err)
	}
	params := pharoah.IQuoterV2QuoteExactInputSingleParams{
		TokenIn:           tokenIn,
		TokenOut:          tokenOut,
		AmountIn:          amountIn,
		TickSpacing:       tickSpacing,
		SqrtPriceLimitX96: big.NewInt(0),
	}
	raw := &pharoah.V3PharoahQuoterRaw{Contract: quoterContract}
	var out []interface{}
	if err := raw.Call(&bind.CallOpts{Context: ctx}, &out, "quoteExactInputSingle", params); err != nil {
		return nil, fmt.Errorf("could not quote pharoah amount out: %w", err)
	}
	amountOut, ok := out[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("unexpected quoter amountOut type %T", out[0])
	}
	return amountOut, nil
}

func encodePharoahExactInputSingle(
	recipient common.Address,
	amountIn, amountOutMin *big.Int,
	tokenIn, tokenOut common.Address,
	tickSpacing, deadline *big.Int,
) ([]byte, error) {
	routerABI, err := pharoah.V3PharoahRouterMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("could not parse pharoah swap router abi: %w", err)
	}
	params := pharoah.ISwapRouterExactInputSingleParams{
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
		return nil, fmt.Errorf("could not encode pharoah exactInputSingle calldata: %w", err)
	}
	return calldata, nil
}

func (c *Client) Pharoah(ctx context.Context, d uatu.IDexRequest) (*uatu.IDexResponse, error) {
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
	pool, err := c.GetPharoahPool(ctx, d.PairAddress)
	if err != nil {
		return nil, err
	}
	if tokenIn != pool.Token0 && tokenIn != pool.Token1 {
		return nil, fmt.Errorf("token %s is not in pool %s", d.TokenIn, d.PairAddress)
	}
	amountOut, err := c.getPharoahAmountsOut(
		ctx,
		d.AmountIn, quoterAddress,
		tokenIn, d.TokenOut,
		pool.TickSpacing,
	)
	if err != nil {
		return nil, err
	}
	amountOutMin := applySlippage(amountOut, d.SlippageBps)
	recipient := d.WalletAddress
	if d.UnwrapNativeOutput {
		recipient = routerAddress
	}
	swapCallData, err := encodePharoahExactInputSingle(
		recipient,
		d.AmountIn, amountOutMin,
		tokenIn, d.TokenOut,
		pool.TickSpacing, deadline,
	)
	if err != nil {
		return nil, err
	}
	if d.UnwrapNativeOutput {
		routerABI, err := pharoah.V3PharoahRouterMetaData.GetAbi()
		if err != nil {
			return nil, fmt.Errorf("could not parse pharoah swap router abi: %w", err)
		}
		unwrapCallData, err := routerABI.Pack("unwrapWETH9", amountOutMin, d.WalletAddress)
		if err != nil {
			return nil, fmt.Errorf("could not encode WETH unwrap calldata: %w", err)
		}
		swapCallData, err = routerABI.Pack("multicall", [][]byte{swapCallData, unwrapCallData})
		if err != nil {
			return nil, fmt.Errorf("could not encode native output multicall: %w", err)
		}
	}
	return &uatu.IDexResponse{
		AmountIn:             d.AmountIn,
		AmountOut:            amountOut,
		AmountOutMinimum:     amountOutMin,
		EncodedData:          swapCallData,
		EncodedERC20Approval: enodedERC20TokenApproval,
		Dex:                  d.Dex,
		RouterAddress:        routerAddress,
		PairAddress:          d.PairAddress,
		Route:                routeWithFee(DexRoutes["pharoah"], poolFeeAmount(d.AmountIn, pool.Fee)),
	}, nil
}
