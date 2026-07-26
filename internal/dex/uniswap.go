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
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/uatu"
	"github.com/uatu/internal/contracts"
)

type Command string

const (
	V2_SWAP Command = "0x08"
	V3_SWAP Command = "0x00"
)

type Client struct {
	client *ethclient.Client
}

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

func Provider(rpcUrl string) (*Client, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, fmt.Errorf("could not create eth client: %w", err)
	}
	return &Client{client: client}, nil
}

func (c *Client) GetV2Pair(address, token0, token1 string) (common.Address, error) {
	factoryAddress := uatu.FormatEvmAddress(address)
	token0Address := uatu.FormatEvmAddress(token0)
	token1Address := uatu.FormatEvmAddress(token1)
	factoryContract, err := contracts.NewV2FactoryCaller(factoryAddress, c.client)
	if err != nil {
		return common.Address{}, fmt.Errorf("could not bind v2 factory: %w", err)
	}
	pair, err := factoryContract.GetPair(&bind.CallOpts{}, token0Address, token1Address)
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

type Permit2Allowance struct {
	Amount     *big.Int
	Expiration *big.Int
	Nonce      *big.Int
}

func (c *Client) GetPermit2TokenAllowance(
	ctx context.Context,
	permit2Address, wallet, token, spender common.Address,
) (*Permit2Allowance, error) {
	contract, err := contracts.NewPermit2Caller(permit2Address, c.client)
	if err != nil {
		return nil, fmt.Errorf("could not bind permit2: %w", err)
	}
	allowance, err := contract.Allowance(
		&bind.CallOpts{Context: ctx},
		wallet,
		token,
		spender,
	)
	if err != nil {
		return nil, fmt.Errorf("could not get permit2 allowance: %w", err)
	}
	return &Permit2Allowance{
		Amount:     allowance.Amount,
		Expiration: allowance.Expiration,
		Nonce:      allowance.Nonce,
	}, nil
}

func (c *Client) GetV2AmountOut(amountIn, reserveIn, reserveOut *big.Int, router string) (*big.Int, error) {
	routerAddress := uatu.FormatEvmAddress(router)
	routerContract, err := contracts.NewV2RouterCaller(routerAddress, c.client)
	if err != nil {
		return nil, fmt.Errorf("could not bind router: %w", err)
	}
	amountOut, err := routerContract.GetAmountOut(&bind.CallOpts{}, amountIn, reserveIn, reserveOut)
	if err != nil {
		return nil, fmt.Errorf("could not get v2 amount out: %w", err)
	}
	return amountOut, nil
}

func (c *Client) GetV3AmountOut(amountIn *big.Int, quoter, tokenIn, tokenOut common.Address, fee *big.Int) (*big.Int, error) {
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

var erc20ABI = func() abi.ABI {
	parsed, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"approve","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"spender","type":"address"}],"name":"allowance","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`))
	if err != nil {
		panic(err)
	}
	return parsed
}()

func EncodeERC20Token(amount *big.Int, spender common.Address) ([]byte, error) {
	encoded, err := erc20ABI.Pack("approve", spender, amount)
	if err != nil {
		return nil, fmt.Errorf("could not encode approve calldata: %w", err)
	}
	return encoded, nil
}

func (c *Client) GetERC20Allowance(
	ctx context.Context,
	token, owner, spender common.Address,
) (*big.Int, error) {
	contract := bind.NewBoundContract(token, erc20ABI, c.client, nil, nil)
	var out []interface{}
	if err := contract.Call(&bind.CallOpts{Context: ctx}, &out, "allowance", owner, spender); err != nil {
		return nil, fmt.Errorf("could not get erc20 allowance: %w", err)
	}
	allowance, ok := out[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("unexpected erc20 allowance type %T", out[0])
	}
	return allowance, nil
}

func EncodePermit2Approval(token, spender common.Address, amount, expiration *big.Int) ([]byte, error) {
	permit2ABI, err := contracts.Permit2MetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("could not parse permit2 abi: %w", err)
	}
	calldata, err := permit2ABI.Pack("approve", token, spender, amount, expiration)
	if err != nil {
		return nil, fmt.Errorf("could not encode permit2 approve calldata: %w", err)
	}
	return calldata, nil
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
func EncodeV2SwapExactIn(
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

func EncodeV3SwapExactSingle(
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

func EncodeUniversalRouterExecute(cmd Command, input []byte, deadline *big.Int) ([]byte, error) {
	calldata, err := universalRouterABI.Pack("execute", common.FromHex(string(cmd)), [][]byte{input}, deadline)
	if err != nil {
		return nil, fmt.Errorf("could not encode execute calldata: %w", err)
	}
	return calldata, nil
}

func (c *Client) Swap(ctx context.Context, d uatu.IDexRequest) (*uatu.IDexResponse, error) {
	switch d.PoolType {
	case "v2":
		return c.BuyV2(ctx, d)
	case "v3":
		return c.BuyV3(ctx, d)
	default:
		return nil, fmt.Errorf("unsupported pool type: %s", d.PoolType)
	}
}

func (c *Client) BuyV2(ctx context.Context, d uatu.IDexRequest) (*uatu.IDexResponse, error) {
	permit2Address := uatu.FormatEvmAddress(d.Dex.Permit2Address)
	universalRouterAddress := uatu.FormatEvmAddress(d.Dex.UniversalRouterAddress)
	tokenIn := d.TokenIn
	permit2Allowance, err := c.GetPermit2TokenAllowance(
		ctx,
		permit2Address,
		d.WalletAddress, d.TokenIn,
		universalRouterAddress,
	)
	if err != nil {
		return nil, err
	}
	erc20Allowance, err := c.GetERC20Allowance(
		ctx,
		tokenIn,
		d.WalletAddress,
		permit2Address,
	)
	if err != nil {
		return nil, err
	}
	now := big.NewInt(time.Now().Unix())
	deadline := big.NewInt(time.Now().Add(swapDeadline).Unix())
	var encodedPermit2Approval, enodedERC20TokenApproval []byte

	if permit2Allowance.Expiration.Cmp(now) <= 0 || permit2Allowance.Amount.Cmp(d.AmountIn) < 0 {
		expiration := big.NewInt(time.Now().Add(permit2ApprovalDuration).Unix())
		encodedPermit2Approval, err = EncodePermit2Approval(tokenIn, universalRouterAddress, maxUint160, expiration)
		if err != nil {
			return nil, err
		}
	}
	if erc20Allowance.Cmp(d.AmountIn) <= 0 {
		enodedERC20TokenApproval, err = EncodeERC20Token(d.AmountIn, permit2Address)
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
	amountOut, err := c.GetV2AmountOut(d.AmountIn, reserveIn, reserveOut, d.Dex.V2RouterAddress)
	if err != nil {
		return nil, err
	}

	input, err := EncodeV2SwapExactIn(
		d.WalletAddress,
		d.AmountIn,
		amountOut,
		[]common.Address{tokenIn, d.TokenOut},
		true,
	)
	if err != nil {
		return nil, err
	}

	swapCalldata, err := EncodeUniversalRouterExecute(V2_SWAP, input, deadline)
	if err != nil {
		return nil, err
	}
	return &uatu.IDexResponse{
		AmountIn:               d.AmountIn,
		AmountOut:              amountOut,
		EncodedData:            swapCalldata,
		EncodedERC20Approval:   enodedERC20TokenApproval,
		EncodedPermit2Approval: encodedPermit2Approval,
		Dex:                    d.Dex,
	}, nil
}

func (c *Client) BuyV3(ctx context.Context, d uatu.IDexRequest) (*uatu.IDexResponse, error) {
	tokenIn := d.TokenIn
	universalRouterAddress := uatu.FormatEvmAddress(d.Dex.UniversalRouterAddress)
	permit2Address := uatu.FormatEvmAddress(d.Dex.Permit2Address)
	quoterAddress := uatu.FormatEvmAddress(d.Dex.V3QuoterAddress)
	permit2Allowance, err := c.GetPermit2TokenAllowance(
		ctx,
		permit2Address,
		d.WalletAddress, d.TokenIn,
		universalRouterAddress,
	)
	if err != nil {
		return nil, err
	}
	erc20Allowance, err := c.GetERC20Allowance(
		ctx,
		tokenIn,
		d.WalletAddress,
		permit2Address,
	)
	if err != nil {
		return nil, err
	}
	now := big.NewInt(time.Now().Unix())
	deadline := big.NewInt(time.Now().Add(swapDeadline).Unix())
	var encodedPermit2Approval, enodedERC20TokenApproval []byte

	if permit2Allowance.Expiration.Cmp(now) <= 0 || permit2Allowance.Amount.Cmp(d.AmountIn) < 0 {
		expiration := big.NewInt(time.Now().Add(permit2ApprovalDuration).Unix())
		encodedPermit2Approval, err = EncodePermit2Approval(tokenIn, universalRouterAddress, maxUint160, expiration)
		if err != nil {
			return nil, err
		}
	}
	if erc20Allowance.Cmp(d.AmountIn) <= 0 {
		enodedERC20TokenApproval, err = EncodeERC20Token(d.AmountIn, permit2Address)
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
	amountOut, err := c.GetV3AmountOut(d.AmountIn, quoterAddress, tokenIn, d.TokenOut, pool.Fee)
	if err != nil {
		return nil, err
	}

	input, err := EncodeV3SwapExactSingle(
		d.WalletAddress,
		d.AmountIn,
		amountOut,
		tokenIn, pool.Fee, d.TokenOut,
		true,
	)
	if err != nil {
		return nil, err
	}
	swapCalldata, err := EncodeUniversalRouterExecute(V3_SWAP, input, deadline)
	if err != nil {
		return nil, err
	}
	return &uatu.IDexResponse{
		AmountIn:               d.AmountIn,
		AmountOut:              amountOut,
		EncodedData:            swapCalldata,
		EncodedERC20Approval:   enodedERC20TokenApproval,
		EncodedPermit2Approval: encodedPermit2Approval,
		Dex:                    d.Dex,
	}, nil
}
