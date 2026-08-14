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
	"github.com/uatu/internal/contracts"
)

type Command string

const (
	V2_SWAP     Command = "0x08"
	V3_SWAP     Command = "0x00"
	WRAP_ETH    Command = "0x0b"
	UNWRAP_WETH Command = "0x0c"
)

// universalRouterRecipient keeps a swap's WETH output in the router so the
// following UNWRAP_WETH command can withdraw it to the user.
var universalRouterRecipient = common.HexToAddress("0x0000000000000000000000000000000000000002")

type V2Pool struct {
	Token0   common.Address
	Token1   common.Address
	Reserve0 *big.Int
	Reserve1 *big.Int
}

type Slot0 struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}

type V3Pool struct {
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	Slot0       Slot0
	Liquidity   *big.Int
	TickSpacing *big.Int
}

type V3Pair struct {
	PairAddress common.Address
	Fee         *big.Int
}

func (c *Client) GetV2Pair(address, token0, token1 common.Address) (common.Address, error) {
	factoryContract, err := contracts.NewV2FactoryCaller(address, c.client)
	if err != nil {
		return common.Address{}, fmt.Errorf("could not bind v2 factory: %w", err)
	}
	pair, err := factoryContract.GetPair(&bind.CallOpts{}, token0, token1)
	if err != nil {
		return common.Address{}, fmt.Errorf("could not get pair: %w", err)
	}
	return pair, nil
}

func (c *Client) GetV2Pool(address common.Address) (*V2Pool, error) {
	poolContract, err := contracts.NewV2PoolCaller(address, c.client)
	if err != nil {
		return nil, fmt.Errorf("could not bind v2 pool: %w", err)
	}
	token0, err := poolContract.Token0(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("could not get token0: %w", err)
	}
	token1, err := poolContract.Token1(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("could not get token1: %w", err)
	}
	reserve, err := poolContract.GetReserves(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("could not get reserves: %w", err)
	}

	return &V2Pool{
		Token0:   token0,
		Token1:   token1,
		Reserve0: reserve.Reserve0,
		Reserve1: reserve.Reserve1,
	}, nil
}

func (c *Client) GetV3Pair(address, token0, token1 common.Address) ([]V3Pair, error) {
	fees := []int64{100, 500, 3000, 10000}
	factoryContract, err := contracts.NewV3FactoryCaller(address, c.client)
	if err != nil {
		return nil, fmt.Errorf("could not bind v3 factory: %w", err)
	}
	pairs := make([]V3Pair, 0, len(fees))
	for _, fee := range fees {
		feeAmount := big.NewInt(fee)
		pair, err := factoryContract.GetPool(&bind.CallOpts{}, token0, token1, feeAmount)
		if err != nil {
			return nil, fmt.Errorf("could not get v3 pool for fee %d: %w", fee, err)
		}
		if pair == (common.Address{}) {
			continue
		}
		pairs = append(pairs, V3Pair{PairAddress: pair, Fee: feeAmount})
	}
	return pairs, nil
}

func (c *Client) GetV3Pool(address common.Address) (*V3Pool, error) {
	poolContract, err := contracts.NewV3PoolCaller(address, c.client)
	if err != nil {
		return nil, fmt.Errorf("could not bind v3 pool: %w", err)
	}
	opts := &bind.CallOpts{}
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

	return &V3Pool{
		Token0:      token0,
		Token1:      token1,
		Fee:         fee,
		Slot0:       Slot0(slot0),
		Liquidity:   liquidity,
		TickSpacing: tickSpacing,
	}, nil
}

func (c *Client) getV2AmountOut(
	ctx context.Context,
	amountIn, reserveIn, reserveOut *big.Int,
	router common.Address,
) (*big.Int, error) {
	routerContract, err := contracts.NewV2RouterCaller(router, c.client)
	if err != nil {
		return nil, fmt.Errorf("could not bind router: %w", err)
	}
	amountOut, err := routerContract.GetAmountOut(
		&bind.CallOpts{Context: ctx},
		amountIn, reserveIn, reserveOut,
	)
	if err != nil {
		return nil, fmt.Errorf("could not get v2 amount out: %w", err)
	}
	return amountOut, nil
}

func (c *Client) getV3AmountOut(
	amountIn *big.Int,
	quoter, tokenIn, tokenOut common.Address,
	fee *big.Int,
) (*big.Int, error) {
	quoterContract, err := contracts.NewV3Quoter(quoter, c.client)
	if err != nil {
		return nil, fmt.Errorf("could not bind v3 quoter: %w", err)
	}
	params := contracts.IQuoterV2QuoteExactInputSingleParams{
		TokenIn:           tokenIn,
		TokenOut:          tokenOut,
		AmountIn:          amountIn,
		Fee:               fee,
		SqrtPriceLimitX96: big.NewInt(0),
	}
	raw := &contracts.V3QuoterRaw{Contract: quoterContract}
	var out []interface{}
	if err := raw.Call(&bind.CallOpts{}, &out, "quoteExactInputSingle", params); err != nil {
		return nil, fmt.Errorf("could not quote v3 amount out: %w", err)
	}
	amountOut, ok := out[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("unexpected quoter amountOut type %T", out[0])
	}
	return amountOut, nil
}

const (
	swapDeadline            = 20 * time.Minute
	permit2ApprovalDuration = 30 * 24 * time.Hour
)

// maxUint160 is permit2's unlimited allowance amount (uint160 max).
var maxUint160 = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 160), big.NewInt(1))

func mustABIType(name string) abi.Type {
	typ, err := abi.NewType(name, "", nil)
	if err != nil {
		panic(err)
	}
	return typ
}

var v2SwapExactInArgs = abi.Arguments{
	{Type: mustABIType("address")},
	{Type: mustABIType("uint256")},
	{Type: mustABIType("uint256")},
	{Type: mustABIType("address[]")},
	{Type: mustABIType("bool")},
}

var universalRouterABI = func() abi.ABI {
	parsed, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"bytes","name":"commands","type":"bytes"},{"internalType":"bytes[]","name":"inputs","type":"bytes[]"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"execute","outputs":[],"stateMutability":"payable","type":"function"}]`))
	if err != nil {
		panic(err)
	}
	return parsed
}()

// EncodeV2SwapExactIn encodes the V2_SWAP_EXACT_IN command input:
// (recipient, amountIn, amountOutMin, path, payerIsUser).
func encodeV2SwapExactIn(
	recipient common.Address,
	amountIn, amountOutMin *big.Int,
	path []common.Address,
	payerIsUser bool,
) ([]byte, error) {
	input, err := v2SwapExactInArgs.Pack(recipient, amountIn, amountOutMin, path, payerIsUser)
	if err != nil {
		return nil, fmt.Errorf("could not encode v2 swap input: %w", err)
	}
	return input, nil
}

var v3SwapExactInArgs = abi.Arguments{
	{Type: mustABIType("address")},
	{Type: mustABIType("uint256")},
	{Type: mustABIType("uint256")},
	{Type: mustABIType("bytes")},
	{Type: mustABIType("bool")},
}

var unwrapWETHArgs = abi.Arguments{
	{Type: mustABIType("address")},
	{Type: mustABIType("uint256")},
}

func encodeUnwrapWETH(recipient common.Address, amountMinimum *big.Int) ([]byte, error) {
	input, err := unwrapWETHArgs.Pack(recipient, amountMinimum)
	if err != nil {
		return nil, fmt.Errorf("could not encode WETH unwrap input: %w", err)
	}
	return input, nil
}

func encodeWrapETH(recipient common.Address, amount *big.Int) ([]byte, error) {
	input, err := unwrapWETHArgs.Pack(recipient, amount)
	if err != nil {
		return nil, fmt.Errorf("could not encode native token wrap input: %w", err)
	}
	return input, nil
}

func encodeV3InputParams(
	recipient common.Address,
	amountIn, amountOutMin *big.Int,
	tokenIn common.Address, fee *big.Int, tokenOut common.Address,
	payerIsUser bool,
) ([]byte, error) {
	if fee.BitLen() > 24 {
		return nil, fmt.Errorf("fee %s does not fit uint24", fee)
	}
	path := make([]byte, 0, 43)
	path = append(path, tokenIn.Bytes()...)
	path = append(path, common.LeftPadBytes(fee.Bytes(), 3)...)
	path = append(path, tokenOut.Bytes()...)

	input, err := v3SwapExactInArgs.Pack(recipient, amountIn, amountOutMin, path, payerIsUser)
	if err != nil {
		return nil, fmt.Errorf("could not encode v3 swap input: %w", err)
	}
	return input, nil
}

// func encodeUniswapExactInputSingle(
// 	recipient common.Address,
// 	amountIn, amountOutMin, fee *big.Int,
// 	tokenIn, tokenOut common.Address,
// ) ([]byte, error) {
// 	routerABI, err := contracts.V3UniswapRouterMetaData.GetAbi()
// 	if err != nil {
// 		return nil, fmt.Errorf("could not parse uniswap v3swap router abi: %w", err)
// 	}
// 	params := contracts.IV3SwapRouterExactInputSingleParams{
// 		TokenIn:           tokenIn,
// 		TokenOut:          tokenOut,
// 		Fee:               fee,
// 		Recipient:         recipient,
// 		AmountIn:          amountIn,
// 		AmountOutMinimum:  amountOutMin,
// 		SqrtPriceLimitX96: big.NewInt(0),
// 	}
// 	calldata, err := routerABI.Pack("exactInputSingle", params)
// 	if err != nil {
// 		return nil, fmt.Errorf("could not encode uniswap exactInputSingle calldata: %w", err)
// 	}
// 	return calldata, nil
// }

func encodeUniversalRouterExecute(commands []Command, inputs [][]byte, deadline *big.Int) ([]byte, error) {
	if len(commands) == 0 || len(commands) != len(inputs) {
		return nil, fmt.Errorf("commands and inputs must be non-empty and have equal length")
	}

	commandBytes := make([]byte, len(commands))
	for i, command := range commands {
		encoded := common.FromHex(string(command))
		if len(encoded) != 1 {
			return nil, fmt.Errorf("invalid universal router command %q", command)
		}
		commandBytes[i] = encoded[0]
	}

	calldata, err := universalRouterABI.Pack("execute",
		commandBytes,
		inputs,
		deadline,
	)
	if err != nil {
		return nil, fmt.Errorf("could not encode execute calldata: %w", err)
	}
	return calldata, nil
}

func (c *Client) swapV2UniversalRouter(ctx context.Context, d uatu.IDexRequest) (*uatu.IDexResponse, error) {
	permit2Address := uatu.FormatEvmAddress(d.Dex.Permit2Address)
	universalRouterAddress := uatu.FormatEvmAddress(d.Dex.UniversalRouterAddress)
	tokenIn := d.TokenIn
	now := big.NewInt(time.Now().Unix())
	deadline := big.NewInt(time.Now().Add(swapDeadline).Unix())
	var encodedPermit2Approval, enodedERC20TokenApproval []byte

	if !d.WrapNativeInput {
		permit2Allowance, err := c.getPermit2TokenAllowance(ctx, permit2Address, d.WalletAddress, d.TokenIn, universalRouterAddress)
		if err != nil {
			return nil, err
		}
		erc20Allowance, err := c.getERC20Allowance(ctx, tokenIn, d.WalletAddress, permit2Address)
		if err != nil {
			return nil, err
		}
		if permit2Allowance.Expiration.Cmp(now) <= 0 || permit2Allowance.Amount.Cmp(d.AmountIn) < 0 {
			expiration := big.NewInt(time.Now().Add(permit2ApprovalDuration).Unix())
			encodedPermit2Approval, err = encodePermit2Approval(tokenIn, universalRouterAddress, maxUint160, expiration)
			if err != nil {
				return nil, err
			}
		}
		if erc20Allowance.Cmp(d.AmountIn) <= 0 {
			enodedERC20TokenApproval, err = encodeERC20Token(d.AmountIn, permit2Address)
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
	if tokenIn == pool.Token1 {
		reserveIn, reserveOut = pool.Reserve1, pool.Reserve0
	} else if tokenIn != pool.Token0 {
		return nil, fmt.Errorf("token %s is not in pool %s", d.TokenIn, d.PairAddress)
	}
	amountOut, err := c.getV2AmountOut(
		ctx, d.AmountIn,
		reserveIn, reserveOut,
		uatu.FormatEvmAddress(d.Dex.V2RouterAddress),
	)
	if err != nil {
		return nil, err
	}

	recipient := d.WalletAddress
	if d.UnwrapNativeOutput {
		recipient = universalRouterRecipient
	}
	input, err := encodeV2SwapExactIn(
		recipient,
		d.AmountIn,
		amountOut,
		[]common.Address{tokenIn, d.TokenOut},
		!d.WrapNativeInput,
	)
	if err != nil {
		return nil, err
	}

	commands := []Command{V2_SWAP}
	inputs := [][]byte{input}
	if d.WrapNativeInput {
		wrapInput, err := encodeWrapETH(universalRouterRecipient, d.AmountIn)
		if err != nil {
			return nil, err
		}
		commands = append([]Command{WRAP_ETH}, commands...)
		inputs = append([][]byte{wrapInput}, inputs...)
	}
	if d.UnwrapNativeOutput {
		unwrapInput, err := encodeUnwrapWETH(d.WalletAddress, amountOut)
		if err != nil {
			return nil, err
		}
		commands = append(commands, UNWRAP_WETH)
		inputs = append(inputs, unwrapInput)
	}
	swapCalldata, err := encodeUniversalRouterExecute(commands, inputs, deadline)
	if err != nil {
		return nil, err
	}
	return &uatu.IDexResponse{
		AmountIn:               d.AmountIn,
		AmountOut:              amountOut,
		EncodedData:            swapCalldata,
		EncodedERC20Approval:   enodedERC20TokenApproval,
		EncodedPermit2Approval: encodedPermit2Approval,
		NativeValue:            nativeValue(d.WrapNativeInput, d.AmountIn),
		Dex:                    d.Dex,
		RouterAddress:          universalRouterAddress,
		PairAddress:            d.PairAddress,
		Route:                  routeWithPoolFee(DexRoutes[d.Dex.Slug], d.AmountIn, d.PoolFee),
	}, nil
}

// func (c *Client) swapV3ExactInputSingle(ctx context.Context, d uatu.IDexRequest) (*uatu.IDexResponse, error) {
// 	tokenIn := d.TokenIn
// 	quoterAddress := uatu.FormatEvmAddress(d.Dex.V3QuoterAddress)
// 	routerAddress := uatu.FormatEvmAddress(d.Dex.V3RouterAddress)
// 	erc20Allowance, err := c.getERC20Allowance(
// 		ctx,
// 		tokenIn,
// 		d.WalletAddress,
// 		routerAddress,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var enodedERC20TokenApproval []byte
// 	if erc20Allowance.Cmp(d.AmountIn) <= 0 {
// 		enodedERC20TokenApproval, err = encodeERC20Token(d.AmountIn, routerAddress)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	pool, err := c.GetV3Pool(d.PairAddress)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if tokenIn != pool.Token0 && tokenIn != pool.Token1 {
// 		return nil, fmt.Errorf("token %s is not in pool %s", d.TokenIn, d.PairAddress)
// 	}
// 	amountOut, err := c.getV3AmountOut(d.AmountIn, quoterAddress, tokenIn, d.TokenOut, pool.Fee)
// 	if err != nil {
// 		return nil, err
// 	}
// 	swapCallData, err := encodeUniswapExactInputSingle(
// 		d.WalletAddress,
// 		d.AmountIn, amountOut, pool.Fee,
// 		tokenIn, d.TokenOut,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &uatu.IDexResponse{
// 		AmountIn:             d.AmountIn,
// 		AmountOut:            amountOut,
// 		EncodedData:          swapCallData,
// 		EncodedERC20Approval: enodedERC20TokenApproval,
// 		Dex:                  d.Dex,
// 		RouterAddress:        routerAddress,
// 		PairAddress:          d.PairAddress,
// 		Route:                routeWithFee(DexRoutes[d.Dex.Slug], poolFeeAmount(d.AmountIn, pool.Fee)),
// 	}, nil
// }

func (c *Client) swapV3UniversalRouter(ctx context.Context, d uatu.IDexRequest) (*uatu.IDexResponse, error) {
	tokenIn := d.TokenIn
	universalRouterAddress := uatu.FormatEvmAddress(d.Dex.UniversalRouterAddress)
	permit2Address := uatu.FormatEvmAddress(d.Dex.Permit2Address)
	quoterAddress := uatu.FormatEvmAddress(d.Dex.V3QuoterAddress)
	now := big.NewInt(time.Now().Unix())
	deadline := big.NewInt(time.Now().Add(swapDeadline).Unix())
	var encodedPermit2Approval, enodedERC20TokenApproval []byte

	if !d.WrapNativeInput {
		permit2Allowance, err := c.getPermit2TokenAllowance(ctx, permit2Address, d.WalletAddress, d.TokenIn, universalRouterAddress)
		if err != nil {
			return nil, err
		}
		erc20Allowance, err := c.getERC20Allowance(ctx, tokenIn, d.WalletAddress, permit2Address)
		if err != nil {
			return nil, err
		}
		if permit2Allowance.Expiration.Cmp(now) <= 0 || permit2Allowance.Amount.Cmp(d.AmountIn) < 0 {
			expiration := big.NewInt(time.Now().Add(permit2ApprovalDuration).Unix())
			encodedPermit2Approval, err = encodePermit2Approval(tokenIn, universalRouterAddress, maxUint160, expiration)
			if err != nil {
				return nil, err
			}
		}
		if erc20Allowance.Cmp(d.AmountIn) <= 0 {
			enodedERC20TokenApproval, err = encodeERC20Token(d.AmountIn, permit2Address)
			if err != nil {
				return nil, err
			}
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

	recipient := d.WalletAddress
	if d.UnwrapNativeOutput {
		recipient = universalRouterRecipient
	}
	input, err := encodeV3InputParams(
		recipient,
		d.AmountIn,
		amountOut,
		tokenIn, pool.Fee, d.TokenOut,
		!d.WrapNativeInput,
	)
	if err != nil {
		return nil, err
	}
	commands := []Command{V3_SWAP}
	inputs := [][]byte{input}
	if d.WrapNativeInput {
		wrapInput, err := encodeWrapETH(universalRouterRecipient, d.AmountIn)
		if err != nil {
			return nil, err
		}
		commands = append([]Command{WRAP_ETH}, commands...)
		inputs = append([][]byte{wrapInput}, inputs...)
	}
	if d.UnwrapNativeOutput {
		unwrapInput, err := encodeUnwrapWETH(d.WalletAddress, amountOut)
		if err != nil {
			return nil, err
		}
		commands = append(commands, UNWRAP_WETH)
		inputs = append(inputs, unwrapInput)
	}
	swapCalldata, err := encodeUniversalRouterExecute(commands, inputs, deadline)
	if err != nil {
		return nil, err
	}
	return &uatu.IDexResponse{
		AmountIn:               d.AmountIn,
		AmountOut:              amountOut,
		EncodedData:            swapCalldata,
		EncodedERC20Approval:   enodedERC20TokenApproval,
		EncodedPermit2Approval: encodedPermit2Approval,
		NativeValue:            nativeValue(d.WrapNativeInput, d.AmountIn),
		Dex:                    d.Dex,
		RouterAddress:          universalRouterAddress,
		PairAddress:            d.PairAddress,
		Route:                  routeWithFee(DexRoutes[d.Dex.Slug], poolFeeAmount(d.AmountIn, pool.Fee)),
	}, nil
}

func nativeValue(wrapNativeInput bool, amountIn *big.Int) *big.Int {
	if !wrapNativeInput {
		return nil
	}
	return new(big.Int).Set(amountIn)
}

func (c *Client) Uniswap(ctx context.Context, d uatu.IDexRequest) (*uatu.IDexResponse, error) {
	switch d.PoolType {
	case "v2":
		return c.swapV2UniversalRouter(ctx, d)
	case "v3":
		return c.swapV3UniversalRouter(ctx, d)
	default:
		return nil, fmt.Errorf("unsupported pool type: %s", d.PoolType)
	}
}
