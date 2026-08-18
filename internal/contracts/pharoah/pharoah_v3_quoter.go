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

// IQuoterV2QuoteExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IQuoterV2QuoteExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	AmountIn          *big.Int
	TickSpacing       *big.Int
	SqrtPriceLimitX96 *big.Int
}

// IQuoterV2QuoteExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IQuoterV2QuoteExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Amount            *big.Int
	TickSpacing       *big.Int
	SqrtPriceLimitX96 *big.Int
}

// V3PharoahQuoterMetaData contains all meta data concerning the V3PharoahQuoter contract.
var V3PharoahQuoterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_deployer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"OutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Overflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"quoteExactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160[]\",\"name\":\"sqrtPriceX96AfterList\",\"type\":\"uint160[]\"},{\"internalType\":\"uint32[]\",\"name\":\"initializedTicksCrossedList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIQuoterV2.QuoteExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96After\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"initializedTicksCrossed\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"quoteExactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160[]\",\"name\":\"sqrtPriceX96AfterList\",\"type\":\"uint160[]\"},{\"internalType\":\"uint32[]\",\"name\":\"initializedTicksCrossedList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIQuoterV2.QuoteExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96After\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"initializedTicksCrossed\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// V3PharoahQuoterABI is the input ABI used to generate the binding from.
// Deprecated: Use V3PharoahQuoterMetaData.ABI instead.
var V3PharoahQuoterABI = V3PharoahQuoterMetaData.ABI

// V3PharoahQuoter is an auto generated Go binding around an Ethereum contract.
type V3PharoahQuoter struct {
	V3PharoahQuoterCaller     // Read-only binding to the contract
	V3PharoahQuoterTransactor // Write-only binding to the contract
	V3PharoahQuoterFilterer   // Log filterer for contract events
}

// V3PharoahQuoterCaller is an auto generated read-only Go binding around an Ethereum contract.
type V3PharoahQuoterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3PharoahQuoterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V3PharoahQuoterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3PharoahQuoterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V3PharoahQuoterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3PharoahQuoterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V3PharoahQuoterSession struct {
	Contract     *V3PharoahQuoter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V3PharoahQuoterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V3PharoahQuoterCallerSession struct {
	Contract *V3PharoahQuoterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// V3PharoahQuoterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V3PharoahQuoterTransactorSession struct {
	Contract     *V3PharoahQuoterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// V3PharoahQuoterRaw is an auto generated low-level Go binding around an Ethereum contract.
type V3PharoahQuoterRaw struct {
	Contract *V3PharoahQuoter // Generic contract binding to access the raw methods on
}

// V3PharoahQuoterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V3PharoahQuoterCallerRaw struct {
	Contract *V3PharoahQuoterCaller // Generic read-only contract binding to access the raw methods on
}

// V3PharoahQuoterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V3PharoahQuoterTransactorRaw struct {
	Contract *V3PharoahQuoterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV3PharoahQuoter creates a new instance of V3PharoahQuoter, bound to a specific deployed contract.
func NewV3PharoahQuoter(address common.Address, backend bind.ContractBackend) (*V3PharoahQuoter, error) {
	contract, err := bindV3PharoahQuoter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V3PharoahQuoter{V3PharoahQuoterCaller: V3PharoahQuoterCaller{contract: contract}, V3PharoahQuoterTransactor: V3PharoahQuoterTransactor{contract: contract}, V3PharoahQuoterFilterer: V3PharoahQuoterFilterer{contract: contract}}, nil
}

// NewV3PharoahQuoterCaller creates a new read-only instance of V3PharoahQuoter, bound to a specific deployed contract.
func NewV3PharoahQuoterCaller(address common.Address, caller bind.ContractCaller) (*V3PharoahQuoterCaller, error) {
	contract, err := bindV3PharoahQuoter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V3PharoahQuoterCaller{contract: contract}, nil
}

// NewV3PharoahQuoterTransactor creates a new write-only instance of V3PharoahQuoter, bound to a specific deployed contract.
func NewV3PharoahQuoterTransactor(address common.Address, transactor bind.ContractTransactor) (*V3PharoahQuoterTransactor, error) {
	contract, err := bindV3PharoahQuoter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V3PharoahQuoterTransactor{contract: contract}, nil
}

// NewV3PharoahQuoterFilterer creates a new log filterer instance of V3PharoahQuoter, bound to a specific deployed contract.
func NewV3PharoahQuoterFilterer(address common.Address, filterer bind.ContractFilterer) (*V3PharoahQuoterFilterer, error) {
	contract, err := bindV3PharoahQuoter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V3PharoahQuoterFilterer{contract: contract}, nil
}

// bindV3PharoahQuoter binds a generic wrapper to an already deployed contract.
func bindV3PharoahQuoter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := V3PharoahQuoterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3PharoahQuoter *V3PharoahQuoterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3PharoahQuoter.Contract.V3PharoahQuoterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3PharoahQuoter *V3PharoahQuoterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3PharoahQuoter.Contract.V3PharoahQuoterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3PharoahQuoter *V3PharoahQuoterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3PharoahQuoter.Contract.V3PharoahQuoterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3PharoahQuoter *V3PharoahQuoterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3PharoahQuoter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3PharoahQuoter *V3PharoahQuoterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3PharoahQuoter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3PharoahQuoter *V3PharoahQuoterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3PharoahQuoter.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_V3PharoahQuoter *V3PharoahQuoterCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3PharoahQuoter.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_V3PharoahQuoter *V3PharoahQuoterSession) WETH9() (common.Address, error) {
	return _V3PharoahQuoter.Contract.WETH9(&_V3PharoahQuoter.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_V3PharoahQuoter *V3PharoahQuoterCallerSession) WETH9() (common.Address, error) {
	return _V3PharoahQuoter.Contract.WETH9(&_V3PharoahQuoter.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_V3PharoahQuoter *V3PharoahQuoterCaller) Deployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3PharoahQuoter.contract.Call(opts, &out, "deployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_V3PharoahQuoter *V3PharoahQuoterSession) Deployer() (common.Address, error) {
	return _V3PharoahQuoter.Contract.Deployer(&_V3PharoahQuoter.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_V3PharoahQuoter *V3PharoahQuoterCallerSession) Deployer() (common.Address, error) {
	return _V3PharoahQuoter.Contract.Deployer(&_V3PharoahQuoter.CallOpts)
}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_V3PharoahQuoter *V3PharoahQuoterCaller) UniswapV3SwapCallback(opts *bind.CallOpts, amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	var out []interface{}
	err := _V3PharoahQuoter.contract.Call(opts, &out, "uniswapV3SwapCallback", amount0Delta, amount1Delta, path)

	if err != nil {
		return err
	}

	return err

}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_V3PharoahQuoter *V3PharoahQuoterSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _V3PharoahQuoter.Contract.UniswapV3SwapCallback(&_V3PharoahQuoter.CallOpts, amount0Delta, amount1Delta, path)
}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_V3PharoahQuoter *V3PharoahQuoterCallerSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _V3PharoahQuoter.Contract.UniswapV3SwapCallback(&_V3PharoahQuoter.CallOpts, amount0Delta, amount1Delta, path)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_V3PharoahQuoter *V3PharoahQuoterTransactor) QuoteExactInput(opts *bind.TransactOpts, path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _V3PharoahQuoter.contract.Transact(opts, "quoteExactInput", path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_V3PharoahQuoter *V3PharoahQuoterSession) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _V3PharoahQuoter.Contract.QuoteExactInput(&_V3PharoahQuoter.TransactOpts, path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_V3PharoahQuoter *V3PharoahQuoterTransactorSession) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _V3PharoahQuoter.Contract.QuoteExactInput(&_V3PharoahQuoter.TransactOpts, path, amountIn)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0x9e7defe6.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,int24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_V3PharoahQuoter *V3PharoahQuoterTransactor) QuoteExactInputSingle(opts *bind.TransactOpts, params IQuoterV2QuoteExactInputSingleParams) (*types.Transaction, error) {
	return _V3PharoahQuoter.contract.Transact(opts, "quoteExactInputSingle", params)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0x9e7defe6.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,int24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_V3PharoahQuoter *V3PharoahQuoterSession) QuoteExactInputSingle(params IQuoterV2QuoteExactInputSingleParams) (*types.Transaction, error) {
	return _V3PharoahQuoter.Contract.QuoteExactInputSingle(&_V3PharoahQuoter.TransactOpts, params)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0x9e7defe6.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,int24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_V3PharoahQuoter *V3PharoahQuoterTransactorSession) QuoteExactInputSingle(params IQuoterV2QuoteExactInputSingleParams) (*types.Transaction, error) {
	return _V3PharoahQuoter.Contract.QuoteExactInputSingle(&_V3PharoahQuoter.TransactOpts, params)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_V3PharoahQuoter *V3PharoahQuoterTransactor) QuoteExactOutput(opts *bind.TransactOpts, path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _V3PharoahQuoter.contract.Transact(opts, "quoteExactOutput", path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_V3PharoahQuoter *V3PharoahQuoterSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _V3PharoahQuoter.Contract.QuoteExactOutput(&_V3PharoahQuoter.TransactOpts, path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_V3PharoahQuoter *V3PharoahQuoterTransactorSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _V3PharoahQuoter.Contract.QuoteExactOutput(&_V3PharoahQuoter.TransactOpts, path, amountOut)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xfa6af908.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,int24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_V3PharoahQuoter *V3PharoahQuoterTransactor) QuoteExactOutputSingle(opts *bind.TransactOpts, params IQuoterV2QuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _V3PharoahQuoter.contract.Transact(opts, "quoteExactOutputSingle", params)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xfa6af908.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,int24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_V3PharoahQuoter *V3PharoahQuoterSession) QuoteExactOutputSingle(params IQuoterV2QuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _V3PharoahQuoter.Contract.QuoteExactOutputSingle(&_V3PharoahQuoter.TransactOpts, params)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xfa6af908.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,int24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_V3PharoahQuoter *V3PharoahQuoterTransactorSession) QuoteExactOutputSingle(params IQuoterV2QuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _V3PharoahQuoter.Contract.QuoteExactOutputSingle(&_V3PharoahQuoter.TransactOpts, params)
}
