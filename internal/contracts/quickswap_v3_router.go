// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ISwapRouterExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	Deadline         *big.Int
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// ISwapRouterExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactInputSingleParams struct {
	TokenIn          common.Address
	TokenOut         common.Address
	Recipient        common.Address
	Deadline         *big.Int
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
	LimitSqrtPrice   *big.Int
}

// ISwapRouterExactOutputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactOutputParams struct {
	Path            []byte
	Recipient       common.Address
	Deadline        *big.Int
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}

// ISwapRouterExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactOutputSingleParams struct {
	TokenIn         common.Address
	TokenOut        common.Address
	Fee             *big.Int
	Recipient       common.Address
	Deadline        *big.Int
	AmountOut       *big.Int
	AmountInMaximum *big.Int
	LimitSqrtPrice  *big.Int
}

// V3QuickSwapRouterMetaData contains all meta data concerning the V3QuickSwapRouter contract.
var V3QuickSwapRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WNativeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poolDeployer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WNativeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"algebraSwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingleSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactOutputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolDeployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundNativeToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"sweepTokenWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWNativeToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"unwrapWNativeTokenWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// V3QuickSwapRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use V3QuickSwapRouterMetaData.ABI instead.
var V3QuickSwapRouterABI = V3QuickSwapRouterMetaData.ABI

// V3QuickSwapRouter is an auto generated Go binding around an Ethereum contract.
type V3QuickSwapRouter struct {
	V3QuickSwapRouterCaller     // Read-only binding to the contract
	V3QuickSwapRouterTransactor // Write-only binding to the contract
	V3QuickSwapRouterFilterer   // Log filterer for contract events
}

// V3QuickSwapRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type V3QuickSwapRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3QuickSwapRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V3QuickSwapRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3QuickSwapRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V3QuickSwapRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3QuickSwapRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V3QuickSwapRouterSession struct {
	Contract     *V3QuickSwapRouter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// V3QuickSwapRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V3QuickSwapRouterCallerSession struct {
	Contract *V3QuickSwapRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// V3QuickSwapRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V3QuickSwapRouterTransactorSession struct {
	Contract     *V3QuickSwapRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// V3QuickSwapRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type V3QuickSwapRouterRaw struct {
	Contract *V3QuickSwapRouter // Generic contract binding to access the raw methods on
}

// V3QuickSwapRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V3QuickSwapRouterCallerRaw struct {
	Contract *V3QuickSwapRouterCaller // Generic read-only contract binding to access the raw methods on
}

// V3QuickSwapRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V3QuickSwapRouterTransactorRaw struct {
	Contract *V3QuickSwapRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV3QuickSwapRouter creates a new instance of V3QuickSwapRouter, bound to a specific deployed contract.
func NewV3QuickSwapRouter(address common.Address, backend bind.ContractBackend) (*V3QuickSwapRouter, error) {
	contract, err := bindV3QuickSwapRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V3QuickSwapRouter{V3QuickSwapRouterCaller: V3QuickSwapRouterCaller{contract: contract}, V3QuickSwapRouterTransactor: V3QuickSwapRouterTransactor{contract: contract}, V3QuickSwapRouterFilterer: V3QuickSwapRouterFilterer{contract: contract}}, nil
}

// NewV3QuickSwapRouterCaller creates a new read-only instance of V3QuickSwapRouter, bound to a specific deployed contract.
func NewV3QuickSwapRouterCaller(address common.Address, caller bind.ContractCaller) (*V3QuickSwapRouterCaller, error) {
	contract, err := bindV3QuickSwapRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V3QuickSwapRouterCaller{contract: contract}, nil
}

// NewV3QuickSwapRouterTransactor creates a new write-only instance of V3QuickSwapRouter, bound to a specific deployed contract.
func NewV3QuickSwapRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*V3QuickSwapRouterTransactor, error) {
	contract, err := bindV3QuickSwapRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V3QuickSwapRouterTransactor{contract: contract}, nil
}

// NewV3QuickSwapRouterFilterer creates a new log filterer instance of V3QuickSwapRouter, bound to a specific deployed contract.
func NewV3QuickSwapRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*V3QuickSwapRouterFilterer, error) {
	contract, err := bindV3QuickSwapRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V3QuickSwapRouterFilterer{contract: contract}, nil
}

// bindV3QuickSwapRouter binds a generic wrapper to an already deployed contract.
func bindV3QuickSwapRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := V3QuickSwapRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3QuickSwapRouter *V3QuickSwapRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3QuickSwapRouter.Contract.V3QuickSwapRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3QuickSwapRouter *V3QuickSwapRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.V3QuickSwapRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3QuickSwapRouter *V3QuickSwapRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.V3QuickSwapRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3QuickSwapRouter *V3QuickSwapRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3QuickSwapRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.contract.Transact(opts, method, params...)
}

// WNativeToken is a free data retrieval call binding the contract method 0x8af3ac85.
//
// Solidity: function WNativeToken() view returns(address)
func (_V3QuickSwapRouter *V3QuickSwapRouterCaller) WNativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3QuickSwapRouter.contract.Call(opts, &out, "WNativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WNativeToken is a free data retrieval call binding the contract method 0x8af3ac85.
//
// Solidity: function WNativeToken() view returns(address)
func (_V3QuickSwapRouter *V3QuickSwapRouterSession) WNativeToken() (common.Address, error) {
	return _V3QuickSwapRouter.Contract.WNativeToken(&_V3QuickSwapRouter.CallOpts)
}

// WNativeToken is a free data retrieval call binding the contract method 0x8af3ac85.
//
// Solidity: function WNativeToken() view returns(address)
func (_V3QuickSwapRouter *V3QuickSwapRouterCallerSession) WNativeToken() (common.Address, error) {
	return _V3QuickSwapRouter.Contract.WNativeToken(&_V3QuickSwapRouter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_V3QuickSwapRouter *V3QuickSwapRouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3QuickSwapRouter.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_V3QuickSwapRouter *V3QuickSwapRouterSession) Factory() (common.Address, error) {
	return _V3QuickSwapRouter.Contract.Factory(&_V3QuickSwapRouter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_V3QuickSwapRouter *V3QuickSwapRouterCallerSession) Factory() (common.Address, error) {
	return _V3QuickSwapRouter.Contract.Factory(&_V3QuickSwapRouter.CallOpts)
}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_V3QuickSwapRouter *V3QuickSwapRouterCaller) PoolDeployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3QuickSwapRouter.contract.Call(opts, &out, "poolDeployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_V3QuickSwapRouter *V3QuickSwapRouterSession) PoolDeployer() (common.Address, error) {
	return _V3QuickSwapRouter.Contract.PoolDeployer(&_V3QuickSwapRouter.CallOpts)
}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_V3QuickSwapRouter *V3QuickSwapRouterCallerSession) PoolDeployer() (common.Address, error) {
	return _V3QuickSwapRouter.Contract.PoolDeployer(&_V3QuickSwapRouter.CallOpts)
}

// AlgebraSwapCallback is a paid mutator transaction binding the contract method 0x2c8958f6.
//
// Solidity: function algebraSwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactor) AlgebraSwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _V3QuickSwapRouter.contract.Transact(opts, "algebraSwapCallback", amount0Delta, amount1Delta, _data)
}

// AlgebraSwapCallback is a paid mutator transaction binding the contract method 0x2c8958f6.
//
// Solidity: function algebraSwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterSession) AlgebraSwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.AlgebraSwapCallback(&_V3QuickSwapRouter.TransactOpts, amount0Delta, amount1Delta, _data)
}

// AlgebraSwapCallback is a paid mutator transaction binding the contract method 0x2c8958f6.
//
// Solidity: function algebraSwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactorSession) AlgebraSwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.AlgebraSwapCallback(&_V3QuickSwapRouter.TransactOpts, amount0Delta, amount1Delta, _data)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactor) ExactInput(opts *bind.TransactOpts, params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _V3QuickSwapRouter.contract.Transact(opts, "exactInput", params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_V3QuickSwapRouter *V3QuickSwapRouterSession) ExactInput(params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.ExactInput(&_V3QuickSwapRouter.TransactOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactorSession) ExactInput(params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.ExactInput(&_V3QuickSwapRouter.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0xbc651188.
//
// Solidity: function exactInputSingle((address,address,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactor) ExactInputSingle(opts *bind.TransactOpts, params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _V3QuickSwapRouter.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0xbc651188.
//
// Solidity: function exactInputSingle((address,address,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_V3QuickSwapRouter *V3QuickSwapRouterSession) ExactInputSingle(params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.ExactInputSingle(&_V3QuickSwapRouter.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0xbc651188.
//
// Solidity: function exactInputSingle((address,address,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactorSession) ExactInputSingle(params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.ExactInputSingle(&_V3QuickSwapRouter.TransactOpts, params)
}

// ExactInputSingleSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb87d2524.
//
// Solidity: function exactInputSingleSupportingFeeOnTransferTokens((address,address,address,uint256,uint256,uint256,uint160) params) returns(uint256 amountOut)
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactor) ExactInputSingleSupportingFeeOnTransferTokens(opts *bind.TransactOpts, params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _V3QuickSwapRouter.contract.Transact(opts, "exactInputSingleSupportingFeeOnTransferTokens", params)
}

// ExactInputSingleSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb87d2524.
//
// Solidity: function exactInputSingleSupportingFeeOnTransferTokens((address,address,address,uint256,uint256,uint256,uint160) params) returns(uint256 amountOut)
func (_V3QuickSwapRouter *V3QuickSwapRouterSession) ExactInputSingleSupportingFeeOnTransferTokens(params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.ExactInputSingleSupportingFeeOnTransferTokens(&_V3QuickSwapRouter.TransactOpts, params)
}

// ExactInputSingleSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb87d2524.
//
// Solidity: function exactInputSingleSupportingFeeOnTransferTokens((address,address,address,uint256,uint256,uint256,uint160) params) returns(uint256 amountOut)
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactorSession) ExactInputSingleSupportingFeeOnTransferTokens(params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.ExactInputSingleSupportingFeeOnTransferTokens(&_V3QuickSwapRouter.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactor) ExactOutput(opts *bind.TransactOpts, params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _V3QuickSwapRouter.contract.Transact(opts, "exactOutput", params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_V3QuickSwapRouter *V3QuickSwapRouterSession) ExactOutput(params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.ExactOutput(&_V3QuickSwapRouter.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactorSession) ExactOutput(params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.ExactOutput(&_V3QuickSwapRouter.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactor) ExactOutputSingle(opts *bind.TransactOpts, params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _V3QuickSwapRouter.contract.Transact(opts, "exactOutputSingle", params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_V3QuickSwapRouter *V3QuickSwapRouterSession) ExactOutputSingle(params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.ExactOutputSingle(&_V3QuickSwapRouter.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactorSession) ExactOutputSingle(params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.ExactOutputSingle(&_V3QuickSwapRouter.TransactOpts, params)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _V3QuickSwapRouter.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_V3QuickSwapRouter *V3QuickSwapRouterSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.Multicall(&_V3QuickSwapRouter.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.Multicall(&_V3QuickSwapRouter.TransactOpts, data)
}

// RefundNativeToken is a paid mutator transaction binding the contract method 0x41865270.
//
// Solidity: function refundNativeToken() payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactor) RefundNativeToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3QuickSwapRouter.contract.Transact(opts, "refundNativeToken")
}

// RefundNativeToken is a paid mutator transaction binding the contract method 0x41865270.
//
// Solidity: function refundNativeToken() payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterSession) RefundNativeToken() (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.RefundNativeToken(&_V3QuickSwapRouter.TransactOpts)
}

// RefundNativeToken is a paid mutator transaction binding the contract method 0x41865270.
//
// Solidity: function refundNativeToken() payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactorSession) RefundNativeToken() (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.RefundNativeToken(&_V3QuickSwapRouter.TransactOpts)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactor) SelfPermit(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3QuickSwapRouter.contract.Transact(opts, "selfPermit", token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.SelfPermit(&_V3QuickSwapRouter.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactorSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.SelfPermit(&_V3QuickSwapRouter.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactor) SelfPermitAllowed(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3QuickSwapRouter.contract.Transact(opts, "selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.SelfPermitAllowed(&_V3QuickSwapRouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactorSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.SelfPermitAllowed(&_V3QuickSwapRouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactor) SelfPermitAllowedIfNecessary(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3QuickSwapRouter.contract.Transact(opts, "selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.SelfPermitAllowedIfNecessary(&_V3QuickSwapRouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactorSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.SelfPermitAllowedIfNecessary(&_V3QuickSwapRouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactor) SelfPermitIfNecessary(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3QuickSwapRouter.contract.Transact(opts, "selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.SelfPermitIfNecessary(&_V3QuickSwapRouter.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactorSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.SelfPermitIfNecessary(&_V3QuickSwapRouter.TransactOpts, token, value, deadline, v, r, s)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3QuickSwapRouter.contract.Transact(opts, "sweepToken", token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.SweepToken(&_V3QuickSwapRouter.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactorSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.SweepToken(&_V3QuickSwapRouter.TransactOpts, token, amountMinimum, recipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactor) SweepTokenWithFee(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _V3QuickSwapRouter.contract.Transact(opts, "sweepTokenWithFee", token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterSession) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.SweepTokenWithFee(&_V3QuickSwapRouter.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactorSession) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.SweepTokenWithFee(&_V3QuickSwapRouter.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWNativeToken is a paid mutator transaction binding the contract method 0x69bc35b2.
//
// Solidity: function unwrapWNativeToken(uint256 amountMinimum, address recipient) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactor) UnwrapWNativeToken(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3QuickSwapRouter.contract.Transact(opts, "unwrapWNativeToken", amountMinimum, recipient)
}

// UnwrapWNativeToken is a paid mutator transaction binding the contract method 0x69bc35b2.
//
// Solidity: function unwrapWNativeToken(uint256 amountMinimum, address recipient) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterSession) UnwrapWNativeToken(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.UnwrapWNativeToken(&_V3QuickSwapRouter.TransactOpts, amountMinimum, recipient)
}

// UnwrapWNativeToken is a paid mutator transaction binding the contract method 0x69bc35b2.
//
// Solidity: function unwrapWNativeToken(uint256 amountMinimum, address recipient) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactorSession) UnwrapWNativeToken(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.UnwrapWNativeToken(&_V3QuickSwapRouter.TransactOpts, amountMinimum, recipient)
}

// UnwrapWNativeTokenWithFee is a paid mutator transaction binding the contract method 0xc60696ec.
//
// Solidity: function unwrapWNativeTokenWithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactor) UnwrapWNativeTokenWithFee(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _V3QuickSwapRouter.contract.Transact(opts, "unwrapWNativeTokenWithFee", amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWNativeTokenWithFee is a paid mutator transaction binding the contract method 0xc60696ec.
//
// Solidity: function unwrapWNativeTokenWithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterSession) UnwrapWNativeTokenWithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.UnwrapWNativeTokenWithFee(&_V3QuickSwapRouter.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWNativeTokenWithFee is a paid mutator transaction binding the contract method 0xc60696ec.
//
// Solidity: function unwrapWNativeTokenWithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactorSession) UnwrapWNativeTokenWithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.UnwrapWNativeTokenWithFee(&_V3QuickSwapRouter.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3QuickSwapRouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterSession) Receive() (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.Receive(&_V3QuickSwapRouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_V3QuickSwapRouter *V3QuickSwapRouterTransactorSession) Receive() (*types.Transaction, error) {
	return _V3QuickSwapRouter.Contract.Receive(&_V3QuickSwapRouter.TransactOpts)
}
