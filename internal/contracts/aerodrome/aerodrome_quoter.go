// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aerodrome

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

// AerodromeQuoterMetaData contains all meta data concerning the AerodromeQuoter contract.
var AerodromeQuoterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"quoteExactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160[]\",\"name\":\"sqrtPriceX96AfterList\",\"type\":\"uint160[]\"},{\"internalType\":\"uint32[]\",\"name\":\"initializedTicksCrossedList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIQuoterV2.QuoteExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96After\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"initializedTicksCrossed\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"quoteExactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160[]\",\"name\":\"sqrtPriceX96AfterList\",\"type\":\"uint160[]\"},{\"internalType\":\"uint32[]\",\"name\":\"initializedTicksCrossedList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIQuoterV2.QuoteExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"quoteExactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96After\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"initializedTicksCrossed\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasEstimate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AerodromeQuoterABI is the input ABI used to generate the binding from.
// Deprecated: Use AerodromeQuoterMetaData.ABI instead.
var AerodromeQuoterABI = AerodromeQuoterMetaData.ABI

// AerodromeQuoter is an auto generated Go binding around an Ethereum contract.
type AerodromeQuoter struct {
	AerodromeQuoterCaller     // Read-only binding to the contract
	AerodromeQuoterTransactor // Write-only binding to the contract
	AerodromeQuoterFilterer   // Log filterer for contract events
}

// AerodromeQuoterCaller is an auto generated read-only Go binding around an Ethereum contract.
type AerodromeQuoterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AerodromeQuoterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AerodromeQuoterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AerodromeQuoterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AerodromeQuoterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AerodromeQuoterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AerodromeQuoterSession struct {
	Contract     *AerodromeQuoter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AerodromeQuoterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AerodromeQuoterCallerSession struct {
	Contract *AerodromeQuoterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AerodromeQuoterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AerodromeQuoterTransactorSession struct {
	Contract     *AerodromeQuoterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AerodromeQuoterRaw is an auto generated low-level Go binding around an Ethereum contract.
type AerodromeQuoterRaw struct {
	Contract *AerodromeQuoter // Generic contract binding to access the raw methods on
}

// AerodromeQuoterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AerodromeQuoterCallerRaw struct {
	Contract *AerodromeQuoterCaller // Generic read-only contract binding to access the raw methods on
}

// AerodromeQuoterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AerodromeQuoterTransactorRaw struct {
	Contract *AerodromeQuoterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAerodromeQuoter creates a new instance of AerodromeQuoter, bound to a specific deployed contract.
func NewAerodromeQuoter(address common.Address, backend bind.ContractBackend) (*AerodromeQuoter, error) {
	contract, err := bindAerodromeQuoter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AerodromeQuoter{AerodromeQuoterCaller: AerodromeQuoterCaller{contract: contract}, AerodromeQuoterTransactor: AerodromeQuoterTransactor{contract: contract}, AerodromeQuoterFilterer: AerodromeQuoterFilterer{contract: contract}}, nil
}

// NewAerodromeQuoterCaller creates a new read-only instance of AerodromeQuoter, bound to a specific deployed contract.
func NewAerodromeQuoterCaller(address common.Address, caller bind.ContractCaller) (*AerodromeQuoterCaller, error) {
	contract, err := bindAerodromeQuoter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AerodromeQuoterCaller{contract: contract}, nil
}

// NewAerodromeQuoterTransactor creates a new write-only instance of AerodromeQuoter, bound to a specific deployed contract.
func NewAerodromeQuoterTransactor(address common.Address, transactor bind.ContractTransactor) (*AerodromeQuoterTransactor, error) {
	contract, err := bindAerodromeQuoter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AerodromeQuoterTransactor{contract: contract}, nil
}

// NewAerodromeQuoterFilterer creates a new log filterer instance of AerodromeQuoter, bound to a specific deployed contract.
func NewAerodromeQuoterFilterer(address common.Address, filterer bind.ContractFilterer) (*AerodromeQuoterFilterer, error) {
	contract, err := bindAerodromeQuoter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AerodromeQuoterFilterer{contract: contract}, nil
}

// bindAerodromeQuoter binds a generic wrapper to an already deployed contract.
func bindAerodromeQuoter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AerodromeQuoterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AerodromeQuoter *AerodromeQuoterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AerodromeQuoter.Contract.AerodromeQuoterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AerodromeQuoter *AerodromeQuoterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AerodromeQuoter.Contract.AerodromeQuoterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AerodromeQuoter *AerodromeQuoterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AerodromeQuoter.Contract.AerodromeQuoterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AerodromeQuoter *AerodromeQuoterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AerodromeQuoter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AerodromeQuoter *AerodromeQuoterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AerodromeQuoter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AerodromeQuoter *AerodromeQuoterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AerodromeQuoter.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_AerodromeQuoter *AerodromeQuoterCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AerodromeQuoter.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_AerodromeQuoter *AerodromeQuoterSession) WETH9() (common.Address, error) {
	return _AerodromeQuoter.Contract.WETH9(&_AerodromeQuoter.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_AerodromeQuoter *AerodromeQuoterCallerSession) WETH9() (common.Address, error) {
	return _AerodromeQuoter.Contract.WETH9(&_AerodromeQuoter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_AerodromeQuoter *AerodromeQuoterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AerodromeQuoter.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_AerodromeQuoter *AerodromeQuoterSession) Factory() (common.Address, error) {
	return _AerodromeQuoter.Contract.Factory(&_AerodromeQuoter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_AerodromeQuoter *AerodromeQuoterCallerSession) Factory() (common.Address, error) {
	return _AerodromeQuoter.Contract.Factory(&_AerodromeQuoter.CallOpts)
}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_AerodromeQuoter *AerodromeQuoterCaller) UniswapV3SwapCallback(opts *bind.CallOpts, amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	var out []interface{}
	err := _AerodromeQuoter.contract.Call(opts, &out, "uniswapV3SwapCallback", amount0Delta, amount1Delta, path)

	if err != nil {
		return err
	}

	return err

}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_AerodromeQuoter *AerodromeQuoterSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _AerodromeQuoter.Contract.UniswapV3SwapCallback(&_AerodromeQuoter.CallOpts, amount0Delta, amount1Delta, path)
}

// UniswapV3SwapCallback is a free data retrieval call binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes path) view returns()
func (_AerodromeQuoter *AerodromeQuoterCallerSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, path []byte) error {
	return _AerodromeQuoter.Contract.UniswapV3SwapCallback(&_AerodromeQuoter.CallOpts, amount0Delta, amount1Delta, path)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_AerodromeQuoter *AerodromeQuoterTransactor) QuoteExactInput(opts *bind.TransactOpts, path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _AerodromeQuoter.contract.Transact(opts, "quoteExactInput", path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_AerodromeQuoter *AerodromeQuoterSession) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _AerodromeQuoter.Contract.QuoteExactInput(&_AerodromeQuoter.TransactOpts, path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_AerodromeQuoter *AerodromeQuoterTransactorSession) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _AerodromeQuoter.Contract.QuoteExactInput(&_AerodromeQuoter.TransactOpts, path, amountIn)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0x9e7defe6.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,int24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_AerodromeQuoter *AerodromeQuoterTransactor) QuoteExactInputSingle(opts *bind.TransactOpts, params IQuoterV2QuoteExactInputSingleParams) (*types.Transaction, error) {
	return _AerodromeQuoter.contract.Transact(opts, "quoteExactInputSingle", params)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0x9e7defe6.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,int24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_AerodromeQuoter *AerodromeQuoterSession) QuoteExactInputSingle(params IQuoterV2QuoteExactInputSingleParams) (*types.Transaction, error) {
	return _AerodromeQuoter.Contract.QuoteExactInputSingle(&_AerodromeQuoter.TransactOpts, params)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0x9e7defe6.
//
// Solidity: function quoteExactInputSingle((address,address,uint256,int24,uint160) params) returns(uint256 amountOut, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_AerodromeQuoter *AerodromeQuoterTransactorSession) QuoteExactInputSingle(params IQuoterV2QuoteExactInputSingleParams) (*types.Transaction, error) {
	return _AerodromeQuoter.Contract.QuoteExactInputSingle(&_AerodromeQuoter.TransactOpts, params)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_AerodromeQuoter *AerodromeQuoterTransactor) QuoteExactOutput(opts *bind.TransactOpts, path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _AerodromeQuoter.contract.Transact(opts, "quoteExactOutput", path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_AerodromeQuoter *AerodromeQuoterSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _AerodromeQuoter.Contract.QuoteExactOutput(&_AerodromeQuoter.TransactOpts, path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn, uint160[] sqrtPriceX96AfterList, uint32[] initializedTicksCrossedList, uint256 gasEstimate)
func (_AerodromeQuoter *AerodromeQuoterTransactorSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _AerodromeQuoter.Contract.QuoteExactOutput(&_AerodromeQuoter.TransactOpts, path, amountOut)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xfa6af908.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,int24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_AerodromeQuoter *AerodromeQuoterTransactor) QuoteExactOutputSingle(opts *bind.TransactOpts, params IQuoterV2QuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _AerodromeQuoter.contract.Transact(opts, "quoteExactOutputSingle", params)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xfa6af908.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,int24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_AerodromeQuoter *AerodromeQuoterSession) QuoteExactOutputSingle(params IQuoterV2QuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _AerodromeQuoter.Contract.QuoteExactOutputSingle(&_AerodromeQuoter.TransactOpts, params)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0xfa6af908.
//
// Solidity: function quoteExactOutputSingle((address,address,uint256,int24,uint160) params) returns(uint256 amountIn, uint160 sqrtPriceX96After, uint32 initializedTicksCrossed, uint256 gasEstimate)
func (_AerodromeQuoter *AerodromeQuoterTransactorSession) QuoteExactOutputSingle(params IQuoterV2QuoteExactOutputSingleParams) (*types.Transaction, error) {
	return _AerodromeQuoter.Contract.QuoteExactOutputSingle(&_AerodromeQuoter.TransactOpts, params)
}
