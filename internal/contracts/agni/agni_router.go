// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package agni

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
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
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
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountOut         *big.Int
	AmountInMaximum   *big.Int
	SqrtPriceLimitX96 *big.Int
}

// V3AgniRouterMetaData contains all meta data concerning the V3AgniRouter contract.
var V3AgniRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_deployer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WMNT\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WMNT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"agniSwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactOutputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundMNT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"sweepTokenWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWMNT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"unwrapWMNTWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// V3AgniRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use V3AgniRouterMetaData.ABI instead.
var V3AgniRouterABI = V3AgniRouterMetaData.ABI

// V3AgniRouter is an auto generated Go binding around an Ethereum contract.
type V3AgniRouter struct {
	V3AgniRouterCaller     // Read-only binding to the contract
	V3AgniRouterTransactor // Write-only binding to the contract
	V3AgniRouterFilterer   // Log filterer for contract events
}

// V3AgniRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type V3AgniRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3AgniRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V3AgniRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3AgniRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V3AgniRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3AgniRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V3AgniRouterSession struct {
	Contract     *V3AgniRouter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V3AgniRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V3AgniRouterCallerSession struct {
	Contract *V3AgniRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// V3AgniRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V3AgniRouterTransactorSession struct {
	Contract     *V3AgniRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// V3AgniRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type V3AgniRouterRaw struct {
	Contract *V3AgniRouter // Generic contract binding to access the raw methods on
}

// V3AgniRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V3AgniRouterCallerRaw struct {
	Contract *V3AgniRouterCaller // Generic read-only contract binding to access the raw methods on
}

// V3AgniRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V3AgniRouterTransactorRaw struct {
	Contract *V3AgniRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV3AgniRouter creates a new instance of V3AgniRouter, bound to a specific deployed contract.
func NewV3AgniRouter(address common.Address, backend bind.ContractBackend) (*V3AgniRouter, error) {
	contract, err := bindV3AgniRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V3AgniRouter{V3AgniRouterCaller: V3AgniRouterCaller{contract: contract}, V3AgniRouterTransactor: V3AgniRouterTransactor{contract: contract}, V3AgniRouterFilterer: V3AgniRouterFilterer{contract: contract}}, nil
}

// NewV3AgniRouterCaller creates a new read-only instance of V3AgniRouter, bound to a specific deployed contract.
func NewV3AgniRouterCaller(address common.Address, caller bind.ContractCaller) (*V3AgniRouterCaller, error) {
	contract, err := bindV3AgniRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V3AgniRouterCaller{contract: contract}, nil
}

// NewV3AgniRouterTransactor creates a new write-only instance of V3AgniRouter, bound to a specific deployed contract.
func NewV3AgniRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*V3AgniRouterTransactor, error) {
	contract, err := bindV3AgniRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V3AgniRouterTransactor{contract: contract}, nil
}

// NewV3AgniRouterFilterer creates a new log filterer instance of V3AgniRouter, bound to a specific deployed contract.
func NewV3AgniRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*V3AgniRouterFilterer, error) {
	contract, err := bindV3AgniRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V3AgniRouterFilterer{contract: contract}, nil
}

// bindV3AgniRouter binds a generic wrapper to an already deployed contract.
func bindV3AgniRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := V3AgniRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3AgniRouter *V3AgniRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3AgniRouter.Contract.V3AgniRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3AgniRouter *V3AgniRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.V3AgniRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3AgniRouter *V3AgniRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.V3AgniRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3AgniRouter *V3AgniRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3AgniRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3AgniRouter *V3AgniRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3AgniRouter *V3AgniRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.contract.Transact(opts, method, params...)
}

// WMNT is a free data retrieval call binding the contract method 0xe8c47196.
//
// Solidity: function WMNT() view returns(address)
func (_V3AgniRouter *V3AgniRouterCaller) WMNT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3AgniRouter.contract.Call(opts, &out, "WMNT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WMNT is a free data retrieval call binding the contract method 0xe8c47196.
//
// Solidity: function WMNT() view returns(address)
func (_V3AgniRouter *V3AgniRouterSession) WMNT() (common.Address, error) {
	return _V3AgniRouter.Contract.WMNT(&_V3AgniRouter.CallOpts)
}

// WMNT is a free data retrieval call binding the contract method 0xe8c47196.
//
// Solidity: function WMNT() view returns(address)
func (_V3AgniRouter *V3AgniRouterCallerSession) WMNT() (common.Address, error) {
	return _V3AgniRouter.Contract.WMNT(&_V3AgniRouter.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_V3AgniRouter *V3AgniRouterCaller) Deployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3AgniRouter.contract.Call(opts, &out, "deployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_V3AgniRouter *V3AgniRouterSession) Deployer() (common.Address, error) {
	return _V3AgniRouter.Contract.Deployer(&_V3AgniRouter.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_V3AgniRouter *V3AgniRouterCallerSession) Deployer() (common.Address, error) {
	return _V3AgniRouter.Contract.Deployer(&_V3AgniRouter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_V3AgniRouter *V3AgniRouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3AgniRouter.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_V3AgniRouter *V3AgniRouterSession) Factory() (common.Address, error) {
	return _V3AgniRouter.Contract.Factory(&_V3AgniRouter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_V3AgniRouter *V3AgniRouterCallerSession) Factory() (common.Address, error) {
	return _V3AgniRouter.Contract.Factory(&_V3AgniRouter.CallOpts)
}

// AgniSwapCallback is a paid mutator transaction binding the contract method 0x5bee97a3.
//
// Solidity: function agniSwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_V3AgniRouter *V3AgniRouterTransactor) AgniSwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _V3AgniRouter.contract.Transact(opts, "agniSwapCallback", amount0Delta, amount1Delta, _data)
}

// AgniSwapCallback is a paid mutator transaction binding the contract method 0x5bee97a3.
//
// Solidity: function agniSwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_V3AgniRouter *V3AgniRouterSession) AgniSwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.AgniSwapCallback(&_V3AgniRouter.TransactOpts, amount0Delta, amount1Delta, _data)
}

// AgniSwapCallback is a paid mutator transaction binding the contract method 0x5bee97a3.
//
// Solidity: function agniSwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_V3AgniRouter *V3AgniRouterTransactorSession) AgniSwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.AgniSwapCallback(&_V3AgniRouter.TransactOpts, amount0Delta, amount1Delta, _data)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_V3AgniRouter *V3AgniRouterTransactor) ExactInput(opts *bind.TransactOpts, params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _V3AgniRouter.contract.Transact(opts, "exactInput", params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_V3AgniRouter *V3AgniRouterSession) ExactInput(params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.ExactInput(&_V3AgniRouter.TransactOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_V3AgniRouter *V3AgniRouterTransactorSession) ExactInput(params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.ExactInput(&_V3AgniRouter.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_V3AgniRouter *V3AgniRouterTransactor) ExactInputSingle(opts *bind.TransactOpts, params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _V3AgniRouter.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_V3AgniRouter *V3AgniRouterSession) ExactInputSingle(params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.ExactInputSingle(&_V3AgniRouter.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_V3AgniRouter *V3AgniRouterTransactorSession) ExactInputSingle(params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.ExactInputSingle(&_V3AgniRouter.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_V3AgniRouter *V3AgniRouterTransactor) ExactOutput(opts *bind.TransactOpts, params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _V3AgniRouter.contract.Transact(opts, "exactOutput", params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_V3AgniRouter *V3AgniRouterSession) ExactOutput(params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.ExactOutput(&_V3AgniRouter.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_V3AgniRouter *V3AgniRouterTransactorSession) ExactOutput(params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.ExactOutput(&_V3AgniRouter.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_V3AgniRouter *V3AgniRouterTransactor) ExactOutputSingle(opts *bind.TransactOpts, params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _V3AgniRouter.contract.Transact(opts, "exactOutputSingle", params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_V3AgniRouter *V3AgniRouterSession) ExactOutputSingle(params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.ExactOutputSingle(&_V3AgniRouter.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_V3AgniRouter *V3AgniRouterTransactorSession) ExactOutputSingle(params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.ExactOutputSingle(&_V3AgniRouter.TransactOpts, params)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_V3AgniRouter *V3AgniRouterTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _V3AgniRouter.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_V3AgniRouter *V3AgniRouterSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.Multicall(&_V3AgniRouter.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_V3AgniRouter *V3AgniRouterTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.Multicall(&_V3AgniRouter.TransactOpts, data)
}

// RefundMNT is a paid mutator transaction binding the contract method 0x63f5daa5.
//
// Solidity: function refundMNT() payable returns()
func (_V3AgniRouter *V3AgniRouterTransactor) RefundMNT(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3AgniRouter.contract.Transact(opts, "refundMNT")
}

// RefundMNT is a paid mutator transaction binding the contract method 0x63f5daa5.
//
// Solidity: function refundMNT() payable returns()
func (_V3AgniRouter *V3AgniRouterSession) RefundMNT() (*types.Transaction, error) {
	return _V3AgniRouter.Contract.RefundMNT(&_V3AgniRouter.TransactOpts)
}

// RefundMNT is a paid mutator transaction binding the contract method 0x63f5daa5.
//
// Solidity: function refundMNT() payable returns()
func (_V3AgniRouter *V3AgniRouterTransactorSession) RefundMNT() (*types.Transaction, error) {
	return _V3AgniRouter.Contract.RefundMNT(&_V3AgniRouter.TransactOpts)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3AgniRouter *V3AgniRouterTransactor) SelfPermit(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3AgniRouter.contract.Transact(opts, "selfPermit", token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3AgniRouter *V3AgniRouterSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.SelfPermit(&_V3AgniRouter.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3AgniRouter *V3AgniRouterTransactorSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.SelfPermit(&_V3AgniRouter.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3AgniRouter *V3AgniRouterTransactor) SelfPermitAllowed(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3AgniRouter.contract.Transact(opts, "selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3AgniRouter *V3AgniRouterSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.SelfPermitAllowed(&_V3AgniRouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3AgniRouter *V3AgniRouterTransactorSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.SelfPermitAllowed(&_V3AgniRouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3AgniRouter *V3AgniRouterTransactor) SelfPermitAllowedIfNecessary(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3AgniRouter.contract.Transact(opts, "selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3AgniRouter *V3AgniRouterSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.SelfPermitAllowedIfNecessary(&_V3AgniRouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3AgniRouter *V3AgniRouterTransactorSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.SelfPermitAllowedIfNecessary(&_V3AgniRouter.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3AgniRouter *V3AgniRouterTransactor) SelfPermitIfNecessary(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3AgniRouter.contract.Transact(opts, "selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3AgniRouter *V3AgniRouterSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.SelfPermitIfNecessary(&_V3AgniRouter.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_V3AgniRouter *V3AgniRouterTransactorSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.SelfPermitIfNecessary(&_V3AgniRouter.TransactOpts, token, value, deadline, v, r, s)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_V3AgniRouter *V3AgniRouterTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3AgniRouter.contract.Transact(opts, "sweepToken", token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_V3AgniRouter *V3AgniRouterSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.SweepToken(&_V3AgniRouter.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_V3AgniRouter *V3AgniRouterTransactorSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.SweepToken(&_V3AgniRouter.TransactOpts, token, amountMinimum, recipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_V3AgniRouter *V3AgniRouterTransactor) SweepTokenWithFee(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _V3AgniRouter.contract.Transact(opts, "sweepTokenWithFee", token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_V3AgniRouter *V3AgniRouterSession) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.SweepTokenWithFee(&_V3AgniRouter.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_V3AgniRouter *V3AgniRouterTransactorSession) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.SweepTokenWithFee(&_V3AgniRouter.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWMNT is a paid mutator transaction binding the contract method 0xcf3f26d6.
//
// Solidity: function unwrapWMNT(uint256 amountMinimum, address recipient) payable returns()
func (_V3AgniRouter *V3AgniRouterTransactor) UnwrapWMNT(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3AgniRouter.contract.Transact(opts, "unwrapWMNT", amountMinimum, recipient)
}

// UnwrapWMNT is a paid mutator transaction binding the contract method 0xcf3f26d6.
//
// Solidity: function unwrapWMNT(uint256 amountMinimum, address recipient) payable returns()
func (_V3AgniRouter *V3AgniRouterSession) UnwrapWMNT(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.UnwrapWMNT(&_V3AgniRouter.TransactOpts, amountMinimum, recipient)
}

// UnwrapWMNT is a paid mutator transaction binding the contract method 0xcf3f26d6.
//
// Solidity: function unwrapWMNT(uint256 amountMinimum, address recipient) payable returns()
func (_V3AgniRouter *V3AgniRouterTransactorSession) UnwrapWMNT(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.UnwrapWMNT(&_V3AgniRouter.TransactOpts, amountMinimum, recipient)
}

// UnwrapWMNTWithFee is a paid mutator transaction binding the contract method 0x30e71e8e.
//
// Solidity: function unwrapWMNTWithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_V3AgniRouter *V3AgniRouterTransactor) UnwrapWMNTWithFee(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _V3AgniRouter.contract.Transact(opts, "unwrapWMNTWithFee", amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWMNTWithFee is a paid mutator transaction binding the contract method 0x30e71e8e.
//
// Solidity: function unwrapWMNTWithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_V3AgniRouter *V3AgniRouterSession) UnwrapWMNTWithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.UnwrapWMNTWithFee(&_V3AgniRouter.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWMNTWithFee is a paid mutator transaction binding the contract method 0x30e71e8e.
//
// Solidity: function unwrapWMNTWithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_V3AgniRouter *V3AgniRouterTransactorSession) UnwrapWMNTWithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _V3AgniRouter.Contract.UnwrapWMNTWithFee(&_V3AgniRouter.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_V3AgniRouter *V3AgniRouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3AgniRouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_V3AgniRouter *V3AgniRouterSession) Receive() (*types.Transaction, error) {
	return _V3AgniRouter.Contract.Receive(&_V3AgniRouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_V3AgniRouter *V3AgniRouterTransactorSession) Receive() (*types.Transaction, error) {
	return _V3AgniRouter.Contract.Receive(&_V3AgniRouter.TransactOpts)
}
