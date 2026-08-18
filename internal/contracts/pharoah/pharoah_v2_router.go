// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pharoah

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IRouterroute is an auto generated low-level Go binding around an user-defined struct.
type IRouterroute struct {
	From   common.Address
	To     common.Address
	Stable bool
}

// V2PharoahRouterMetaData contains all meta data concerning the V2PharoahRouter contract.
var V2PharoahRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ETH_TRANSFER_FAILED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EXCESSIVE_INPUT_AMOUNT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EXPIRED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IDENTICAL\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INSUFFICIENT_AMOUNT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INSUFFICIENT_A_AMOUNT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INSUFFICIENT_B_AMOUNT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INSUFFICIENT_LIQUIDITY\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INSUFFICIENT_OUTPUT_AMOUNT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_PATH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_RESERVES\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"NO_GAUGE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityAndStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityETHAndStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structIRouter.route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structIRouter.route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"pairFor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"}],\"name\":\"quoteAddLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"name\":\"quoteRemoveLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETHSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"sortTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structIRouter.route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapETHForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structIRouter.route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structIRouter.route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structIRouter.route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structIRouter.route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETHSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structIRouter.route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structIRouter.route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structIRouter.route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"internalType\":\"structIRouter.route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// V2PharoahRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use V2PharoahRouterMetaData.ABI instead.
var V2PharoahRouterABI = V2PharoahRouterMetaData.ABI

// V2PharoahRouter is an auto generated Go binding around an Ethereum contract.
type V2PharoahRouter struct {
	V2PharoahRouterCaller     // Read-only binding to the contract
	V2PharoahRouterTransactor // Write-only binding to the contract
	V2PharoahRouterFilterer   // Log filterer for contract events
}

// V2PharoahRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type V2PharoahRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V2PharoahRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V2PharoahRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V2PharoahRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V2PharoahRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V2PharoahRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V2PharoahRouterSession struct {
	Contract     *V2PharoahRouter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V2PharoahRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V2PharoahRouterCallerSession struct {
	Contract *V2PharoahRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// V2PharoahRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V2PharoahRouterTransactorSession struct {
	Contract     *V2PharoahRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// V2PharoahRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type V2PharoahRouterRaw struct {
	Contract *V2PharoahRouter // Generic contract binding to access the raw methods on
}

// V2PharoahRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V2PharoahRouterCallerRaw struct {
	Contract *V2PharoahRouterCaller // Generic read-only contract binding to access the raw methods on
}

// V2PharoahRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V2PharoahRouterTransactorRaw struct {
	Contract *V2PharoahRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV2PharoahRouter creates a new instance of V2PharoahRouter, bound to a specific deployed contract.
func NewV2PharoahRouter(address common.Address, backend bind.ContractBackend) (*V2PharoahRouter, error) {
	contract, err := bindV2PharoahRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V2PharoahRouter{V2PharoahRouterCaller: V2PharoahRouterCaller{contract: contract}, V2PharoahRouterTransactor: V2PharoahRouterTransactor{contract: contract}, V2PharoahRouterFilterer: V2PharoahRouterFilterer{contract: contract}}, nil
}

// NewV2PharoahRouterCaller creates a new read-only instance of V2PharoahRouter, bound to a specific deployed contract.
func NewV2PharoahRouterCaller(address common.Address, caller bind.ContractCaller) (*V2PharoahRouterCaller, error) {
	contract, err := bindV2PharoahRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V2PharoahRouterCaller{contract: contract}, nil
}

// NewV2PharoahRouterTransactor creates a new write-only instance of V2PharoahRouter, bound to a specific deployed contract.
func NewV2PharoahRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*V2PharoahRouterTransactor, error) {
	contract, err := bindV2PharoahRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V2PharoahRouterTransactor{contract: contract}, nil
}

// NewV2PharoahRouterFilterer creates a new log filterer instance of V2PharoahRouter, bound to a specific deployed contract.
func NewV2PharoahRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*V2PharoahRouterFilterer, error) {
	contract, err := bindV2PharoahRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V2PharoahRouterFilterer{contract: contract}, nil
}

// bindV2PharoahRouter binds a generic wrapper to an already deployed contract.
func bindV2PharoahRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := V2PharoahRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V2PharoahRouter *V2PharoahRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V2PharoahRouter.Contract.V2PharoahRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V2PharoahRouter *V2PharoahRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.V2PharoahRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V2PharoahRouter *V2PharoahRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.V2PharoahRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V2PharoahRouter *V2PharoahRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V2PharoahRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V2PharoahRouter *V2PharoahRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V2PharoahRouter *V2PharoahRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_V2PharoahRouter *V2PharoahRouterCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V2PharoahRouter.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_V2PharoahRouter *V2PharoahRouterSession) WETH() (common.Address, error) {
	return _V2PharoahRouter.Contract.WETH(&_V2PharoahRouter.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_V2PharoahRouter *V2PharoahRouterCallerSession) WETH() (common.Address, error) {
	return _V2PharoahRouter.Contract.WETH(&_V2PharoahRouter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_V2PharoahRouter *V2PharoahRouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V2PharoahRouter.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_V2PharoahRouter *V2PharoahRouterSession) Factory() (common.Address, error) {
	return _V2PharoahRouter.Contract.Factory(&_V2PharoahRouter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_V2PharoahRouter *V2PharoahRouterCallerSession) Factory() (common.Address, error) {
	return _V2PharoahRouter.Contract.Factory(&_V2PharoahRouter.CallOpts)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x5e1e6325.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn, address tokenOut) view returns(uint256 amount, bool stable)
func (_V2PharoahRouter *V2PharoahRouterCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, tokenIn common.Address, tokenOut common.Address) (struct {
	Amount *big.Int
	Stable bool
}, error) {
	var out []interface{}
	err := _V2PharoahRouter.contract.Call(opts, &out, "getAmountOut", amountIn, tokenIn, tokenOut)

	outstruct := new(struct {
		Amount *big.Int
		Stable bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Stable = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x5e1e6325.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn, address tokenOut) view returns(uint256 amount, bool stable)
func (_V2PharoahRouter *V2PharoahRouterSession) GetAmountOut(amountIn *big.Int, tokenIn common.Address, tokenOut common.Address) (struct {
	Amount *big.Int
	Stable bool
}, error) {
	return _V2PharoahRouter.Contract.GetAmountOut(&_V2PharoahRouter.CallOpts, amountIn, tokenIn, tokenOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x5e1e6325.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn, address tokenOut) view returns(uint256 amount, bool stable)
func (_V2PharoahRouter *V2PharoahRouterCallerSession) GetAmountOut(amountIn *big.Int, tokenIn common.Address, tokenOut common.Address) (struct {
	Amount *big.Int
	Stable bool
}, error) {
	return _V2PharoahRouter.Contract.GetAmountOut(&_V2PharoahRouter.CallOpts, amountIn, tokenIn, tokenOut)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0xd5c54e33.
//
// Solidity: function getAmountsIn(uint256 amountOut, (address,address,bool)[] routes) view returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterCaller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, routes []IRouterroute) ([]*big.Int, error) {
	var out []interface{}
	err := _V2PharoahRouter.contract.Call(opts, &out, "getAmountsIn", amountOut, routes)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0xd5c54e33.
//
// Solidity: function getAmountsIn(uint256 amountOut, (address,address,bool)[] routes) view returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterSession) GetAmountsIn(amountOut *big.Int, routes []IRouterroute) ([]*big.Int, error) {
	return _V2PharoahRouter.Contract.GetAmountsIn(&_V2PharoahRouter.CallOpts, amountOut, routes)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0xd5c54e33.
//
// Solidity: function getAmountsIn(uint256 amountOut, (address,address,bool)[] routes) view returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterCallerSession) GetAmountsIn(amountOut *big.Int, routes []IRouterroute) ([]*big.Int, error) {
	return _V2PharoahRouter.Contract.GetAmountsIn(&_V2PharoahRouter.CallOpts, amountOut, routes)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x9881fcb4.
//
// Solidity: function getAmountsOut(uint256 amountIn, (address,address,bool)[] routes) view returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterCaller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, routes []IRouterroute) ([]*big.Int, error) {
	var out []interface{}
	err := _V2PharoahRouter.contract.Call(opts, &out, "getAmountsOut", amountIn, routes)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0x9881fcb4.
//
// Solidity: function getAmountsOut(uint256 amountIn, (address,address,bool)[] routes) view returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterSession) GetAmountsOut(amountIn *big.Int, routes []IRouterroute) ([]*big.Int, error) {
	return _V2PharoahRouter.Contract.GetAmountsOut(&_V2PharoahRouter.CallOpts, amountIn, routes)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x9881fcb4.
//
// Solidity: function getAmountsOut(uint256 amountIn, (address,address,bool)[] routes) view returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterCallerSession) GetAmountsOut(amountIn *big.Int, routes []IRouterroute) ([]*big.Int, error) {
	return _V2PharoahRouter.Contract.GetAmountsOut(&_V2PharoahRouter.CallOpts, amountIn, routes)
}

// GetReserves is a free data retrieval call binding the contract method 0x5e60dab5.
//
// Solidity: function getReserves(address tokenA, address tokenB, bool stable) view returns(uint256 reserveA, uint256 reserveB)
func (_V2PharoahRouter *V2PharoahRouterCaller) GetReserves(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	var out []interface{}
	err := _V2PharoahRouter.contract.Call(opts, &out, "getReserves", tokenA, tokenB, stable)

	outstruct := new(struct {
		ReserveA *big.Int
		ReserveB *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReserveA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReserveB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x5e60dab5.
//
// Solidity: function getReserves(address tokenA, address tokenB, bool stable) view returns(uint256 reserveA, uint256 reserveB)
func (_V2PharoahRouter *V2PharoahRouterSession) GetReserves(tokenA common.Address, tokenB common.Address, stable bool) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _V2PharoahRouter.Contract.GetReserves(&_V2PharoahRouter.CallOpts, tokenA, tokenB, stable)
}

// GetReserves is a free data retrieval call binding the contract method 0x5e60dab5.
//
// Solidity: function getReserves(address tokenA, address tokenB, bool stable) view returns(uint256 reserveA, uint256 reserveB)
func (_V2PharoahRouter *V2PharoahRouterCallerSession) GetReserves(tokenA common.Address, tokenB common.Address, stable bool) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _V2PharoahRouter.Contract.GetReserves(&_V2PharoahRouter.CallOpts, tokenA, tokenB, stable)
}

// PairFor is a free data retrieval call binding the contract method 0x4c1ee03e.
//
// Solidity: function pairFor(address tokenA, address tokenB, bool stable) view returns(address pair)
func (_V2PharoahRouter *V2PharoahRouterCaller) PairFor(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	var out []interface{}
	err := _V2PharoahRouter.contract.Call(opts, &out, "pairFor", tokenA, tokenB, stable)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairFor is a free data retrieval call binding the contract method 0x4c1ee03e.
//
// Solidity: function pairFor(address tokenA, address tokenB, bool stable) view returns(address pair)
func (_V2PharoahRouter *V2PharoahRouterSession) PairFor(tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	return _V2PharoahRouter.Contract.PairFor(&_V2PharoahRouter.CallOpts, tokenA, tokenB, stable)
}

// PairFor is a free data retrieval call binding the contract method 0x4c1ee03e.
//
// Solidity: function pairFor(address tokenA, address tokenB, bool stable) view returns(address pair)
func (_V2PharoahRouter *V2PharoahRouterCallerSession) PairFor(tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	return _V2PharoahRouter.Contract.PairFor(&_V2PharoahRouter.CallOpts, tokenA, tokenB, stable)
}

// QuoteAddLiquidity is a free data retrieval call binding the contract method 0x98a0fb3c.
//
// Solidity: function quoteAddLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired) view returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_V2PharoahRouter *V2PharoahRouterCaller) QuoteAddLiquidity(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int) (struct {
	AmountA   *big.Int
	AmountB   *big.Int
	Liquidity *big.Int
}, error) {
	var out []interface{}
	err := _V2PharoahRouter.contract.Call(opts, &out, "quoteAddLiquidity", tokenA, tokenB, stable, amountADesired, amountBDesired)

	outstruct := new(struct {
		AmountA   *big.Int
		AmountB   *big.Int
		Liquidity *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Liquidity = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuoteAddLiquidity is a free data retrieval call binding the contract method 0x98a0fb3c.
//
// Solidity: function quoteAddLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired) view returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_V2PharoahRouter *V2PharoahRouterSession) QuoteAddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int) (struct {
	AmountA   *big.Int
	AmountB   *big.Int
	Liquidity *big.Int
}, error) {
	return _V2PharoahRouter.Contract.QuoteAddLiquidity(&_V2PharoahRouter.CallOpts, tokenA, tokenB, stable, amountADesired, amountBDesired)
}

// QuoteAddLiquidity is a free data retrieval call binding the contract method 0x98a0fb3c.
//
// Solidity: function quoteAddLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired) view returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_V2PharoahRouter *V2PharoahRouterCallerSession) QuoteAddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int) (struct {
	AmountA   *big.Int
	AmountB   *big.Int
	Liquidity *big.Int
}, error) {
	return _V2PharoahRouter.Contract.QuoteAddLiquidity(&_V2PharoahRouter.CallOpts, tokenA, tokenB, stable, amountADesired, amountBDesired)
}

// QuoteRemoveLiquidity is a free data retrieval call binding the contract method 0x4386e63c.
//
// Solidity: function quoteRemoveLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity) view returns(uint256 amountA, uint256 amountB)
func (_V2PharoahRouter *V2PharoahRouterCaller) QuoteRemoveLiquidity(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	var out []interface{}
	err := _V2PharoahRouter.contract.Call(opts, &out, "quoteRemoveLiquidity", tokenA, tokenB, stable, liquidity)

	outstruct := new(struct {
		AmountA *big.Int
		AmountB *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QuoteRemoveLiquidity is a free data retrieval call binding the contract method 0x4386e63c.
//
// Solidity: function quoteRemoveLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity) view returns(uint256 amountA, uint256 amountB)
func (_V2PharoahRouter *V2PharoahRouterSession) QuoteRemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	return _V2PharoahRouter.Contract.QuoteRemoveLiquidity(&_V2PharoahRouter.CallOpts, tokenA, tokenB, stable, liquidity)
}

// QuoteRemoveLiquidity is a free data retrieval call binding the contract method 0x4386e63c.
//
// Solidity: function quoteRemoveLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity) view returns(uint256 amountA, uint256 amountB)
func (_V2PharoahRouter *V2PharoahRouterCallerSession) QuoteRemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int) (struct {
	AmountA *big.Int
	AmountB *big.Int
}, error) {
	return _V2PharoahRouter.Contract.QuoteRemoveLiquidity(&_V2PharoahRouter.CallOpts, tokenA, tokenB, stable, liquidity)
}

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (_V2PharoahRouter *V2PharoahRouterCaller) SortTokens(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
}, error) {
	var out []interface{}
	err := _V2PharoahRouter.contract.Call(opts, &out, "sortTokens", tokenA, tokenB)

	outstruct := new(struct {
		Token0 common.Address
		Token1 common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (_V2PharoahRouter *V2PharoahRouterSession) SortTokens(tokenA common.Address, tokenB common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
}, error) {
	return _V2PharoahRouter.Contract.SortTokens(&_V2PharoahRouter.CallOpts, tokenA, tokenB)
}

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (_V2PharoahRouter *V2PharoahRouterCallerSession) SortTokens(tokenA common.Address, tokenB common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
}, error) {
	return _V2PharoahRouter.Contract.SortTokens(&_V2PharoahRouter.CallOpts, tokenA, tokenB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x5a47ddc3.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_V2PharoahRouter *V2PharoahRouterTransactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.contract.Transact(opts, "addLiquidity", tokenA, tokenB, stable, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x5a47ddc3.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_V2PharoahRouter *V2PharoahRouterSession) AddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.AddLiquidity(&_V2PharoahRouter.TransactOpts, tokenA, tokenB, stable, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x5a47ddc3.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_V2PharoahRouter *V2PharoahRouterTransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.AddLiquidity(&_V2PharoahRouter.TransactOpts, tokenA, tokenB, stable, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityAndStake is a paid mutator transaction binding the contract method 0xaa729753.
//
// Solidity: function addLiquidityAndStake(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_V2PharoahRouter *V2PharoahRouterTransactor) AddLiquidityAndStake(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.contract.Transact(opts, "addLiquidityAndStake", tokenA, tokenB, stable, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityAndStake is a paid mutator transaction binding the contract method 0xaa729753.
//
// Solidity: function addLiquidityAndStake(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_V2PharoahRouter *V2PharoahRouterSession) AddLiquidityAndStake(tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.AddLiquidityAndStake(&_V2PharoahRouter.TransactOpts, tokenA, tokenB, stable, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityAndStake is a paid mutator transaction binding the contract method 0xaa729753.
//
// Solidity: function addLiquidityAndStake(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_V2PharoahRouter *V2PharoahRouterTransactorSession) AddLiquidityAndStake(tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.AddLiquidityAndStake(&_V2PharoahRouter.TransactOpts, tokenA, tokenB, stable, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xb7e0d4c0.
//
// Solidity: function addLiquidityETH(address token, bool stable, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_V2PharoahRouter *V2PharoahRouterTransactor) AddLiquidityETH(opts *bind.TransactOpts, token common.Address, stable bool, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.contract.Transact(opts, "addLiquidityETH", token, stable, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xb7e0d4c0.
//
// Solidity: function addLiquidityETH(address token, bool stable, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_V2PharoahRouter *V2PharoahRouterSession) AddLiquidityETH(token common.Address, stable bool, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.AddLiquidityETH(&_V2PharoahRouter.TransactOpts, token, stable, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0xb7e0d4c0.
//
// Solidity: function addLiquidityETH(address token, bool stable, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_V2PharoahRouter *V2PharoahRouterTransactorSession) AddLiquidityETH(token common.Address, stable bool, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.AddLiquidityETH(&_V2PharoahRouter.TransactOpts, token, stable, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETHAndStake is a paid mutator transaction binding the contract method 0x34fa2bb2.
//
// Solidity: function addLiquidityETHAndStake(address token, bool stable, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_V2PharoahRouter *V2PharoahRouterTransactor) AddLiquidityETHAndStake(opts *bind.TransactOpts, token common.Address, stable bool, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.contract.Transact(opts, "addLiquidityETHAndStake", token, stable, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETHAndStake is a paid mutator transaction binding the contract method 0x34fa2bb2.
//
// Solidity: function addLiquidityETHAndStake(address token, bool stable, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_V2PharoahRouter *V2PharoahRouterSession) AddLiquidityETHAndStake(token common.Address, stable bool, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.AddLiquidityETHAndStake(&_V2PharoahRouter.TransactOpts, token, stable, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETHAndStake is a paid mutator transaction binding the contract method 0x34fa2bb2.
//
// Solidity: function addLiquidityETHAndStake(address token, bool stable, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_V2PharoahRouter *V2PharoahRouterTransactorSession) AddLiquidityETHAndStake(token common.Address, stable bool, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.AddLiquidityETHAndStake(&_V2PharoahRouter.TransactOpts, token, stable, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0dede6c4.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_V2PharoahRouter *V2PharoahRouterTransactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, stable, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0dede6c4.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_V2PharoahRouter *V2PharoahRouterSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.RemoveLiquidity(&_V2PharoahRouter.TransactOpts, tokenA, tokenB, stable, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0dede6c4.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_V2PharoahRouter *V2PharoahRouterTransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.RemoveLiquidity(&_V2PharoahRouter.TransactOpts, tokenA, tokenB, stable, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xd7b0e0a5.
//
// Solidity: function removeLiquidityETH(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_V2PharoahRouter *V2PharoahRouterTransactor) RemoveLiquidityETH(opts *bind.TransactOpts, token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.contract.Transact(opts, "removeLiquidityETH", token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xd7b0e0a5.
//
// Solidity: function removeLiquidityETH(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_V2PharoahRouter *V2PharoahRouterSession) RemoveLiquidityETH(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.RemoveLiquidityETH(&_V2PharoahRouter.TransactOpts, token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xd7b0e0a5.
//
// Solidity: function removeLiquidityETH(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_V2PharoahRouter *V2PharoahRouterTransactorSession) RemoveLiquidityETH(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.RemoveLiquidityETH(&_V2PharoahRouter.TransactOpts, token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xfe411f14.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_V2PharoahRouter *V2PharoahRouterTransactor) RemoveLiquidityETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.contract.Transact(opts, "removeLiquidityETHSupportingFeeOnTransferTokens", token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xfe411f14.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_V2PharoahRouter *V2PharoahRouterSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_V2PharoahRouter.TransactOpts, token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xfe411f14.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (_V2PharoahRouter *V2PharoahRouterTransactorSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_V2PharoahRouter.TransactOpts, token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0x633afc92.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, (address,address,bool)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterTransactor) SwapETHForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.contract.Transact(opts, "swapETHForExactTokens", amountOut, routes, to, deadline)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0x633afc92.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, (address,address,bool)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterSession) SwapETHForExactTokens(amountOut *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.SwapETHForExactTokens(&_V2PharoahRouter.TransactOpts, amountOut, routes, to, deadline)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0x633afc92.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, (address,address,bool)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterTransactorSession) SwapETHForExactTokens(amountOut *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.SwapETHForExactTokens(&_V2PharoahRouter.TransactOpts, amountOut, routes, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x67ffb66a.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterTransactor) SwapExactETHForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.contract.Transact(opts, "swapExactETHForTokens", amountOutMin, routes, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x67ffb66a.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterSession) SwapExactETHForTokens(amountOutMin *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.SwapExactETHForTokens(&_V2PharoahRouter.TransactOpts, amountOutMin, routes, to, deadline)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x67ffb66a.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterTransactorSession) SwapExactETHForTokens(amountOutMin *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.SwapExactETHForTokens(&_V2PharoahRouter.TransactOpts, amountOutMin, routes, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x76c72751.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) payable returns()
func (_V2PharoahRouter *V2PharoahRouterTransactor) SwapExactETHForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.contract.Transact(opts, "swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, routes, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x76c72751.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) payable returns()
func (_V2PharoahRouter *V2PharoahRouterSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_V2PharoahRouter.TransactOpts, amountOutMin, routes, to, deadline)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x76c72751.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) payable returns()
func (_V2PharoahRouter *V2PharoahRouterTransactorSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_V2PharoahRouter.TransactOpts, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18a13086.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterTransactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.contract.Transact(opts, "swapExactTokensForETH", amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18a13086.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.SwapExactTokensForETH(&_V2PharoahRouter.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18a13086.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterTransactorSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.SwapExactTokensForETH(&_V2PharoahRouter.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x7af728c8.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns()
func (_V2PharoahRouter *V2PharoahRouterTransactor) SwapExactTokensForETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.contract.Transact(opts, "swapExactTokensForETHSupportingFeeOnTransferTokens", amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x7af728c8.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns()
func (_V2PharoahRouter *V2PharoahRouterSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_V2PharoahRouter.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x7af728c8.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns()
func (_V2PharoahRouter *V2PharoahRouterTransactorSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_V2PharoahRouter.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xf41766d8.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xf41766d8.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.SwapExactTokensForTokens(&_V2PharoahRouter.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xf41766d8.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.SwapExactTokensForTokens(&_V2PharoahRouter.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x6cc1ae13.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns()
func (_V2PharoahRouter *V2PharoahRouterTransactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x6cc1ae13.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns()
func (_V2PharoahRouter *V2PharoahRouterSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_V2PharoahRouter.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x6cc1ae13.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool)[] routes, address to, uint256 deadline) returns()
func (_V2PharoahRouter *V2PharoahRouterTransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_V2PharoahRouter.TransactOpts, amountIn, amountOutMin, routes, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0xd69f344c.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterTransactor) SwapTokensForExactETH(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.contract.Transact(opts, "swapTokensForExactETH", amountOut, amountInMax, routes, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0xd69f344c.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterSession) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.SwapTokensForExactETH(&_V2PharoahRouter.TransactOpts, amountOut, amountInMax, routes, to, deadline)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0xd69f344c.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterTransactorSession) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.SwapTokensForExactETH(&_V2PharoahRouter.TransactOpts, amountOut, amountInMax, routes, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0xed1fbca2.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterTransactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, routes, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0xed1fbca2.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.SwapTokensForExactTokens(&_V2PharoahRouter.TransactOpts, amountOut, amountInMax, routes, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0xed1fbca2.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, (address,address,bool)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (_V2PharoahRouter *V2PharoahRouterTransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, routes []IRouterroute, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.SwapTokensForExactTokens(&_V2PharoahRouter.TransactOpts, amountOut, amountInMax, routes, to, deadline)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_V2PharoahRouter *V2PharoahRouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2PharoahRouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_V2PharoahRouter *V2PharoahRouterSession) Receive() (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.Receive(&_V2PharoahRouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_V2PharoahRouter *V2PharoahRouterTransactorSession) Receive() (*types.Transaction, error) {
	return _V2PharoahRouter.Contract.Receive(&_V2PharoahRouter.TransactOpts)
}
