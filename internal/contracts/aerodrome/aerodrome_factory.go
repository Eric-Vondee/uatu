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

// AerodromeFactoryMetaData contains all meta data concerning the AerodromeFactory contract.
var AerodromeFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poolImplementation\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"oldUnstakedFee\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"newUnstakedFee\",\"type\":\"uint24\"}],\"name\":\"DefaultUnstakedFeeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldFeeManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeManager\",\"type\":\"address\"}],\"name\":\"SwapFeeManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldFeeModule\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeModule\",\"type\":\"address\"}],\"name\":\"SwapFeeModuleChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"TickSpacingEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldFeeManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeManager\",\"type\":\"address\"}],\"name\":\"UnstakedFeeManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldFeeModule\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeModule\",\"type\":\"address\"}],\"name\":\"UnstakedFeeModuleChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPoolsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultUnstakedFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"enableTickSpacing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryRegistry\",\"outputs\":[{\"internalType\":\"contractIFactoryRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getSwapFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getUnstakedFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_defaultUnstakedFee\",\"type\":\"uint24\"}],\"name\":\"setDefaultUnstakedFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_swapFeeManager\",\"type\":\"address\"}],\"name\":\"setSwapFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_swapFeeModule\",\"type\":\"address\"}],\"name\":\"setSwapFeeModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_unstakedFeeManager\",\"type\":\"address\"}],\"name\":\"setUnstakedFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_unstakedFeeModule\",\"type\":\"address\"}],\"name\":\"setUnstakedFeeModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapFeeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapFeeModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"name\":\"tickSpacingToFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSpacings\",\"outputs\":[{\"internalType\":\"int24[]\",\"name\":\"\",\"type\":\"int24[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstakedFeeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstakedFeeModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voter\",\"outputs\":[{\"internalType\":\"contractIVoter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AerodromeFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use AerodromeFactoryMetaData.ABI instead.
var AerodromeFactoryABI = AerodromeFactoryMetaData.ABI

// AerodromeFactory is an auto generated Go binding around an Ethereum contract.
type AerodromeFactory struct {
	AerodromeFactoryCaller     // Read-only binding to the contract
	AerodromeFactoryTransactor // Write-only binding to the contract
	AerodromeFactoryFilterer   // Log filterer for contract events
}

// AerodromeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AerodromeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AerodromeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AerodromeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AerodromeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AerodromeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AerodromeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AerodromeFactorySession struct {
	Contract     *AerodromeFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AerodromeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AerodromeFactoryCallerSession struct {
	Contract *AerodromeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// AerodromeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AerodromeFactoryTransactorSession struct {
	Contract     *AerodromeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AerodromeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AerodromeFactoryRaw struct {
	Contract *AerodromeFactory // Generic contract binding to access the raw methods on
}

// AerodromeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AerodromeFactoryCallerRaw struct {
	Contract *AerodromeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// AerodromeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AerodromeFactoryTransactorRaw struct {
	Contract *AerodromeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAerodromeFactory creates a new instance of AerodromeFactory, bound to a specific deployed contract.
func NewAerodromeFactory(address common.Address, backend bind.ContractBackend) (*AerodromeFactory, error) {
	contract, err := bindAerodromeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AerodromeFactory{AerodromeFactoryCaller: AerodromeFactoryCaller{contract: contract}, AerodromeFactoryTransactor: AerodromeFactoryTransactor{contract: contract}, AerodromeFactoryFilterer: AerodromeFactoryFilterer{contract: contract}}, nil
}

// NewAerodromeFactoryCaller creates a new read-only instance of AerodromeFactory, bound to a specific deployed contract.
func NewAerodromeFactoryCaller(address common.Address, caller bind.ContractCaller) (*AerodromeFactoryCaller, error) {
	contract, err := bindAerodromeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AerodromeFactoryCaller{contract: contract}, nil
}

// NewAerodromeFactoryTransactor creates a new write-only instance of AerodromeFactory, bound to a specific deployed contract.
func NewAerodromeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AerodromeFactoryTransactor, error) {
	contract, err := bindAerodromeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AerodromeFactoryTransactor{contract: contract}, nil
}

// NewAerodromeFactoryFilterer creates a new log filterer instance of AerodromeFactory, bound to a specific deployed contract.
func NewAerodromeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AerodromeFactoryFilterer, error) {
	contract, err := bindAerodromeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AerodromeFactoryFilterer{contract: contract}, nil
}

// bindAerodromeFactory binds a generic wrapper to an already deployed contract.
func bindAerodromeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AerodromeFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AerodromeFactory *AerodromeFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AerodromeFactory.Contract.AerodromeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AerodromeFactory *AerodromeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.AerodromeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AerodromeFactory *AerodromeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.AerodromeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AerodromeFactory *AerodromeFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AerodromeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AerodromeFactory *AerodromeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AerodromeFactory *AerodromeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.contract.Transact(opts, method, params...)
}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (_AerodromeFactory *AerodromeFactoryCaller) AllPools(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AerodromeFactory.contract.Call(opts, &out, "allPools", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (_AerodromeFactory *AerodromeFactorySession) AllPools(arg0 *big.Int) (common.Address, error) {
	return _AerodromeFactory.Contract.AllPools(&_AerodromeFactory.CallOpts, arg0)
}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (_AerodromeFactory *AerodromeFactoryCallerSession) AllPools(arg0 *big.Int) (common.Address, error) {
	return _AerodromeFactory.Contract.AllPools(&_AerodromeFactory.CallOpts, arg0)
}

// AllPoolsLength is a free data retrieval call binding the contract method 0xefde4e64.
//
// Solidity: function allPoolsLength() view returns(uint256)
func (_AerodromeFactory *AerodromeFactoryCaller) AllPoolsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AerodromeFactory.contract.Call(opts, &out, "allPoolsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPoolsLength is a free data retrieval call binding the contract method 0xefde4e64.
//
// Solidity: function allPoolsLength() view returns(uint256)
func (_AerodromeFactory *AerodromeFactorySession) AllPoolsLength() (*big.Int, error) {
	return _AerodromeFactory.Contract.AllPoolsLength(&_AerodromeFactory.CallOpts)
}

// AllPoolsLength is a free data retrieval call binding the contract method 0xefde4e64.
//
// Solidity: function allPoolsLength() view returns(uint256)
func (_AerodromeFactory *AerodromeFactoryCallerSession) AllPoolsLength() (*big.Int, error) {
	return _AerodromeFactory.Contract.AllPoolsLength(&_AerodromeFactory.CallOpts)
}

// DefaultUnstakedFee is a free data retrieval call binding the contract method 0xe2824832.
//
// Solidity: function defaultUnstakedFee() view returns(uint24)
func (_AerodromeFactory *AerodromeFactoryCaller) DefaultUnstakedFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AerodromeFactory.contract.Call(opts, &out, "defaultUnstakedFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultUnstakedFee is a free data retrieval call binding the contract method 0xe2824832.
//
// Solidity: function defaultUnstakedFee() view returns(uint24)
func (_AerodromeFactory *AerodromeFactorySession) DefaultUnstakedFee() (*big.Int, error) {
	return _AerodromeFactory.Contract.DefaultUnstakedFee(&_AerodromeFactory.CallOpts)
}

// DefaultUnstakedFee is a free data retrieval call binding the contract method 0xe2824832.
//
// Solidity: function defaultUnstakedFee() view returns(uint24)
func (_AerodromeFactory *AerodromeFactoryCallerSession) DefaultUnstakedFee() (*big.Int, error) {
	return _AerodromeFactory.Contract.DefaultUnstakedFee(&_AerodromeFactory.CallOpts)
}

// FactoryRegistry is a free data retrieval call binding the contract method 0x3bf0c9fb.
//
// Solidity: function factoryRegistry() view returns(address)
func (_AerodromeFactory *AerodromeFactoryCaller) FactoryRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AerodromeFactory.contract.Call(opts, &out, "factoryRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryRegistry is a free data retrieval call binding the contract method 0x3bf0c9fb.
//
// Solidity: function factoryRegistry() view returns(address)
func (_AerodromeFactory *AerodromeFactorySession) FactoryRegistry() (common.Address, error) {
	return _AerodromeFactory.Contract.FactoryRegistry(&_AerodromeFactory.CallOpts)
}

// FactoryRegistry is a free data retrieval call binding the contract method 0x3bf0c9fb.
//
// Solidity: function factoryRegistry() view returns(address)
func (_AerodromeFactory *AerodromeFactoryCallerSession) FactoryRegistry() (common.Address, error) {
	return _AerodromeFactory.Contract.FactoryRegistry(&_AerodromeFactory.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0x28af8d0b.
//
// Solidity: function getPool(address , address , int24 ) view returns(address)
func (_AerodromeFactory *AerodromeFactoryCaller) GetPool(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AerodromeFactory.contract.Call(opts, &out, "getPool", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x28af8d0b.
//
// Solidity: function getPool(address , address , int24 ) view returns(address)
func (_AerodromeFactory *AerodromeFactorySession) GetPool(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	return _AerodromeFactory.Contract.GetPool(&_AerodromeFactory.CallOpts, arg0, arg1, arg2)
}

// GetPool is a free data retrieval call binding the contract method 0x28af8d0b.
//
// Solidity: function getPool(address , address , int24 ) view returns(address)
func (_AerodromeFactory *AerodromeFactoryCallerSession) GetPool(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	return _AerodromeFactory.Contract.GetPool(&_AerodromeFactory.CallOpts, arg0, arg1, arg2)
}

// GetSwapFee is a free data retrieval call binding the contract method 0x35458dcc.
//
// Solidity: function getSwapFee(address pool) view returns(uint24)
func (_AerodromeFactory *AerodromeFactoryCaller) GetSwapFee(opts *bind.CallOpts, pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AerodromeFactory.contract.Call(opts, &out, "getSwapFee", pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFee is a free data retrieval call binding the contract method 0x35458dcc.
//
// Solidity: function getSwapFee(address pool) view returns(uint24)
func (_AerodromeFactory *AerodromeFactorySession) GetSwapFee(pool common.Address) (*big.Int, error) {
	return _AerodromeFactory.Contract.GetSwapFee(&_AerodromeFactory.CallOpts, pool)
}

// GetSwapFee is a free data retrieval call binding the contract method 0x35458dcc.
//
// Solidity: function getSwapFee(address pool) view returns(uint24)
func (_AerodromeFactory *AerodromeFactoryCallerSession) GetSwapFee(pool common.Address) (*big.Int, error) {
	return _AerodromeFactory.Contract.GetSwapFee(&_AerodromeFactory.CallOpts, pool)
}

// GetUnstakedFee is a free data retrieval call binding the contract method 0x48cf7a43.
//
// Solidity: function getUnstakedFee(address pool) view returns(uint24)
func (_AerodromeFactory *AerodromeFactoryCaller) GetUnstakedFee(opts *bind.CallOpts, pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AerodromeFactory.contract.Call(opts, &out, "getUnstakedFee", pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnstakedFee is a free data retrieval call binding the contract method 0x48cf7a43.
//
// Solidity: function getUnstakedFee(address pool) view returns(uint24)
func (_AerodromeFactory *AerodromeFactorySession) GetUnstakedFee(pool common.Address) (*big.Int, error) {
	return _AerodromeFactory.Contract.GetUnstakedFee(&_AerodromeFactory.CallOpts, pool)
}

// GetUnstakedFee is a free data retrieval call binding the contract method 0x48cf7a43.
//
// Solidity: function getUnstakedFee(address pool) view returns(uint24)
func (_AerodromeFactory *AerodromeFactoryCallerSession) GetUnstakedFee(pool common.Address) (*big.Int, error) {
	return _AerodromeFactory.Contract.GetUnstakedFee(&_AerodromeFactory.CallOpts, pool)
}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address pool) view returns(bool)
func (_AerodromeFactory *AerodromeFactoryCaller) IsPool(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var out []interface{}
	err := _AerodromeFactory.contract.Call(opts, &out, "isPool", pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address pool) view returns(bool)
func (_AerodromeFactory *AerodromeFactorySession) IsPool(pool common.Address) (bool, error) {
	return _AerodromeFactory.Contract.IsPool(&_AerodromeFactory.CallOpts, pool)
}

// IsPool is a free data retrieval call binding the contract method 0x5b16ebb7.
//
// Solidity: function isPool(address pool) view returns(bool)
func (_AerodromeFactory *AerodromeFactoryCallerSession) IsPool(pool common.Address) (bool, error) {
	return _AerodromeFactory.Contract.IsPool(&_AerodromeFactory.CallOpts, pool)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AerodromeFactory *AerodromeFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AerodromeFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AerodromeFactory *AerodromeFactorySession) Owner() (common.Address, error) {
	return _AerodromeFactory.Contract.Owner(&_AerodromeFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AerodromeFactory *AerodromeFactoryCallerSession) Owner() (common.Address, error) {
	return _AerodromeFactory.Contract.Owner(&_AerodromeFactory.CallOpts)
}

// PoolImplementation is a free data retrieval call binding the contract method 0xcefa7799.
//
// Solidity: function poolImplementation() view returns(address)
func (_AerodromeFactory *AerodromeFactoryCaller) PoolImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AerodromeFactory.contract.Call(opts, &out, "poolImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolImplementation is a free data retrieval call binding the contract method 0xcefa7799.
//
// Solidity: function poolImplementation() view returns(address)
func (_AerodromeFactory *AerodromeFactorySession) PoolImplementation() (common.Address, error) {
	return _AerodromeFactory.Contract.PoolImplementation(&_AerodromeFactory.CallOpts)
}

// PoolImplementation is a free data retrieval call binding the contract method 0xcefa7799.
//
// Solidity: function poolImplementation() view returns(address)
func (_AerodromeFactory *AerodromeFactoryCallerSession) PoolImplementation() (common.Address, error) {
	return _AerodromeFactory.Contract.PoolImplementation(&_AerodromeFactory.CallOpts)
}

// SwapFeeManager is a free data retrieval call binding the contract method 0xd574afa9.
//
// Solidity: function swapFeeManager() view returns(address)
func (_AerodromeFactory *AerodromeFactoryCaller) SwapFeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AerodromeFactory.contract.Call(opts, &out, "swapFeeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapFeeManager is a free data retrieval call binding the contract method 0xd574afa9.
//
// Solidity: function swapFeeManager() view returns(address)
func (_AerodromeFactory *AerodromeFactorySession) SwapFeeManager() (common.Address, error) {
	return _AerodromeFactory.Contract.SwapFeeManager(&_AerodromeFactory.CallOpts)
}

// SwapFeeManager is a free data retrieval call binding the contract method 0xd574afa9.
//
// Solidity: function swapFeeManager() view returns(address)
func (_AerodromeFactory *AerodromeFactoryCallerSession) SwapFeeManager() (common.Address, error) {
	return _AerodromeFactory.Contract.SwapFeeManager(&_AerodromeFactory.CallOpts)
}

// SwapFeeModule is a free data retrieval call binding the contract method 0x23c43a51.
//
// Solidity: function swapFeeModule() view returns(address)
func (_AerodromeFactory *AerodromeFactoryCaller) SwapFeeModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AerodromeFactory.contract.Call(opts, &out, "swapFeeModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapFeeModule is a free data retrieval call binding the contract method 0x23c43a51.
//
// Solidity: function swapFeeModule() view returns(address)
func (_AerodromeFactory *AerodromeFactorySession) SwapFeeModule() (common.Address, error) {
	return _AerodromeFactory.Contract.SwapFeeModule(&_AerodromeFactory.CallOpts)
}

// SwapFeeModule is a free data retrieval call binding the contract method 0x23c43a51.
//
// Solidity: function swapFeeModule() view returns(address)
func (_AerodromeFactory *AerodromeFactoryCallerSession) SwapFeeModule() (common.Address, error) {
	return _AerodromeFactory.Contract.SwapFeeModule(&_AerodromeFactory.CallOpts)
}

// TickSpacingToFee is a free data retrieval call binding the contract method 0x380dc1c2.
//
// Solidity: function tickSpacingToFee(int24 ) view returns(uint24)
func (_AerodromeFactory *AerodromeFactoryCaller) TickSpacingToFee(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AerodromeFactory.contract.Call(opts, &out, "tickSpacingToFee", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSpacingToFee is a free data retrieval call binding the contract method 0x380dc1c2.
//
// Solidity: function tickSpacingToFee(int24 ) view returns(uint24)
func (_AerodromeFactory *AerodromeFactorySession) TickSpacingToFee(arg0 *big.Int) (*big.Int, error) {
	return _AerodromeFactory.Contract.TickSpacingToFee(&_AerodromeFactory.CallOpts, arg0)
}

// TickSpacingToFee is a free data retrieval call binding the contract method 0x380dc1c2.
//
// Solidity: function tickSpacingToFee(int24 ) view returns(uint24)
func (_AerodromeFactory *AerodromeFactoryCallerSession) TickSpacingToFee(arg0 *big.Int) (*big.Int, error) {
	return _AerodromeFactory.Contract.TickSpacingToFee(&_AerodromeFactory.CallOpts, arg0)
}

// TickSpacings is a free data retrieval call binding the contract method 0x9cbbbe86.
//
// Solidity: function tickSpacings() view returns(int24[])
func (_AerodromeFactory *AerodromeFactoryCaller) TickSpacings(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _AerodromeFactory.contract.Call(opts, &out, "tickSpacings")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TickSpacings is a free data retrieval call binding the contract method 0x9cbbbe86.
//
// Solidity: function tickSpacings() view returns(int24[])
func (_AerodromeFactory *AerodromeFactorySession) TickSpacings() ([]*big.Int, error) {
	return _AerodromeFactory.Contract.TickSpacings(&_AerodromeFactory.CallOpts)
}

// TickSpacings is a free data retrieval call binding the contract method 0x9cbbbe86.
//
// Solidity: function tickSpacings() view returns(int24[])
func (_AerodromeFactory *AerodromeFactoryCallerSession) TickSpacings() ([]*big.Int, error) {
	return _AerodromeFactory.Contract.TickSpacings(&_AerodromeFactory.CallOpts)
}

// UnstakedFeeManager is a free data retrieval call binding the contract method 0x82e189e0.
//
// Solidity: function unstakedFeeManager() view returns(address)
func (_AerodromeFactory *AerodromeFactoryCaller) UnstakedFeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AerodromeFactory.contract.Call(opts, &out, "unstakedFeeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnstakedFeeManager is a free data retrieval call binding the contract method 0x82e189e0.
//
// Solidity: function unstakedFeeManager() view returns(address)
func (_AerodromeFactory *AerodromeFactorySession) UnstakedFeeManager() (common.Address, error) {
	return _AerodromeFactory.Contract.UnstakedFeeManager(&_AerodromeFactory.CallOpts)
}

// UnstakedFeeManager is a free data retrieval call binding the contract method 0x82e189e0.
//
// Solidity: function unstakedFeeManager() view returns(address)
func (_AerodromeFactory *AerodromeFactoryCallerSession) UnstakedFeeManager() (common.Address, error) {
	return _AerodromeFactory.Contract.UnstakedFeeManager(&_AerodromeFactory.CallOpts)
}

// UnstakedFeeModule is a free data retrieval call binding the contract method 0x7693bc11.
//
// Solidity: function unstakedFeeModule() view returns(address)
func (_AerodromeFactory *AerodromeFactoryCaller) UnstakedFeeModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AerodromeFactory.contract.Call(opts, &out, "unstakedFeeModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnstakedFeeModule is a free data retrieval call binding the contract method 0x7693bc11.
//
// Solidity: function unstakedFeeModule() view returns(address)
func (_AerodromeFactory *AerodromeFactorySession) UnstakedFeeModule() (common.Address, error) {
	return _AerodromeFactory.Contract.UnstakedFeeModule(&_AerodromeFactory.CallOpts)
}

// UnstakedFeeModule is a free data retrieval call binding the contract method 0x7693bc11.
//
// Solidity: function unstakedFeeModule() view returns(address)
func (_AerodromeFactory *AerodromeFactoryCallerSession) UnstakedFeeModule() (common.Address, error) {
	return _AerodromeFactory.Contract.UnstakedFeeModule(&_AerodromeFactory.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_AerodromeFactory *AerodromeFactoryCaller) Voter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AerodromeFactory.contract.Call(opts, &out, "voter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_AerodromeFactory *AerodromeFactorySession) Voter() (common.Address, error) {
	return _AerodromeFactory.Contract.Voter(&_AerodromeFactory.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_AerodromeFactory *AerodromeFactoryCallerSession) Voter() (common.Address, error) {
	return _AerodromeFactory.Contract.Voter(&_AerodromeFactory.CallOpts)
}

// CreatePool is a paid mutator transaction binding the contract method 0x232aa5ac.
//
// Solidity: function createPool(address tokenA, address tokenB, int24 tickSpacing, uint160 sqrtPriceX96) returns(address pool)
func (_AerodromeFactory *AerodromeFactoryTransactor) CreatePool(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, tickSpacing *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _AerodromeFactory.contract.Transact(opts, "createPool", tokenA, tokenB, tickSpacing, sqrtPriceX96)
}

// CreatePool is a paid mutator transaction binding the contract method 0x232aa5ac.
//
// Solidity: function createPool(address tokenA, address tokenB, int24 tickSpacing, uint160 sqrtPriceX96) returns(address pool)
func (_AerodromeFactory *AerodromeFactorySession) CreatePool(tokenA common.Address, tokenB common.Address, tickSpacing *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.CreatePool(&_AerodromeFactory.TransactOpts, tokenA, tokenB, tickSpacing, sqrtPriceX96)
}

// CreatePool is a paid mutator transaction binding the contract method 0x232aa5ac.
//
// Solidity: function createPool(address tokenA, address tokenB, int24 tickSpacing, uint160 sqrtPriceX96) returns(address pool)
func (_AerodromeFactory *AerodromeFactoryTransactorSession) CreatePool(tokenA common.Address, tokenB common.Address, tickSpacing *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.CreatePool(&_AerodromeFactory.TransactOpts, tokenA, tokenB, tickSpacing, sqrtPriceX96)
}

// EnableTickSpacing is a paid mutator transaction binding the contract method 0xeee0fdb4.
//
// Solidity: function enableTickSpacing(int24 tickSpacing, uint24 fee) returns()
func (_AerodromeFactory *AerodromeFactoryTransactor) EnableTickSpacing(opts *bind.TransactOpts, tickSpacing *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _AerodromeFactory.contract.Transact(opts, "enableTickSpacing", tickSpacing, fee)
}

// EnableTickSpacing is a paid mutator transaction binding the contract method 0xeee0fdb4.
//
// Solidity: function enableTickSpacing(int24 tickSpacing, uint24 fee) returns()
func (_AerodromeFactory *AerodromeFactorySession) EnableTickSpacing(tickSpacing *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.EnableTickSpacing(&_AerodromeFactory.TransactOpts, tickSpacing, fee)
}

// EnableTickSpacing is a paid mutator transaction binding the contract method 0xeee0fdb4.
//
// Solidity: function enableTickSpacing(int24 tickSpacing, uint24 fee) returns()
func (_AerodromeFactory *AerodromeFactoryTransactorSession) EnableTickSpacing(tickSpacing *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.EnableTickSpacing(&_AerodromeFactory.TransactOpts, tickSpacing, fee)
}

// SetDefaultUnstakedFee is a paid mutator transaction binding the contract method 0xa2f97f42.
//
// Solidity: function setDefaultUnstakedFee(uint24 _defaultUnstakedFee) returns()
func (_AerodromeFactory *AerodromeFactoryTransactor) SetDefaultUnstakedFee(opts *bind.TransactOpts, _defaultUnstakedFee *big.Int) (*types.Transaction, error) {
	return _AerodromeFactory.contract.Transact(opts, "setDefaultUnstakedFee", _defaultUnstakedFee)
}

// SetDefaultUnstakedFee is a paid mutator transaction binding the contract method 0xa2f97f42.
//
// Solidity: function setDefaultUnstakedFee(uint24 _defaultUnstakedFee) returns()
func (_AerodromeFactory *AerodromeFactorySession) SetDefaultUnstakedFee(_defaultUnstakedFee *big.Int) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.SetDefaultUnstakedFee(&_AerodromeFactory.TransactOpts, _defaultUnstakedFee)
}

// SetDefaultUnstakedFee is a paid mutator transaction binding the contract method 0xa2f97f42.
//
// Solidity: function setDefaultUnstakedFee(uint24 _defaultUnstakedFee) returns()
func (_AerodromeFactory *AerodromeFactoryTransactorSession) SetDefaultUnstakedFee(_defaultUnstakedFee *big.Int) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.SetDefaultUnstakedFee(&_AerodromeFactory.TransactOpts, _defaultUnstakedFee)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_AerodromeFactory *AerodromeFactoryTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _AerodromeFactory.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_AerodromeFactory *AerodromeFactorySession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.SetOwner(&_AerodromeFactory.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_AerodromeFactory *AerodromeFactoryTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.SetOwner(&_AerodromeFactory.TransactOpts, _owner)
}

// SetSwapFeeManager is a paid mutator transaction binding the contract method 0xffb4d9d1.
//
// Solidity: function setSwapFeeManager(address _swapFeeManager) returns()
func (_AerodromeFactory *AerodromeFactoryTransactor) SetSwapFeeManager(opts *bind.TransactOpts, _swapFeeManager common.Address) (*types.Transaction, error) {
	return _AerodromeFactory.contract.Transact(opts, "setSwapFeeManager", _swapFeeManager)
}

// SetSwapFeeManager is a paid mutator transaction binding the contract method 0xffb4d9d1.
//
// Solidity: function setSwapFeeManager(address _swapFeeManager) returns()
func (_AerodromeFactory *AerodromeFactorySession) SetSwapFeeManager(_swapFeeManager common.Address) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.SetSwapFeeManager(&_AerodromeFactory.TransactOpts, _swapFeeManager)
}

// SetSwapFeeManager is a paid mutator transaction binding the contract method 0xffb4d9d1.
//
// Solidity: function setSwapFeeManager(address _swapFeeManager) returns()
func (_AerodromeFactory *AerodromeFactoryTransactorSession) SetSwapFeeManager(_swapFeeManager common.Address) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.SetSwapFeeManager(&_AerodromeFactory.TransactOpts, _swapFeeManager)
}

// SetSwapFeeModule is a paid mutator transaction binding the contract method 0x61b9c3ec.
//
// Solidity: function setSwapFeeModule(address _swapFeeModule) returns()
func (_AerodromeFactory *AerodromeFactoryTransactor) SetSwapFeeModule(opts *bind.TransactOpts, _swapFeeModule common.Address) (*types.Transaction, error) {
	return _AerodromeFactory.contract.Transact(opts, "setSwapFeeModule", _swapFeeModule)
}

// SetSwapFeeModule is a paid mutator transaction binding the contract method 0x61b9c3ec.
//
// Solidity: function setSwapFeeModule(address _swapFeeModule) returns()
func (_AerodromeFactory *AerodromeFactorySession) SetSwapFeeModule(_swapFeeModule common.Address) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.SetSwapFeeModule(&_AerodromeFactory.TransactOpts, _swapFeeModule)
}

// SetSwapFeeModule is a paid mutator transaction binding the contract method 0x61b9c3ec.
//
// Solidity: function setSwapFeeModule(address _swapFeeModule) returns()
func (_AerodromeFactory *AerodromeFactoryTransactorSession) SetSwapFeeModule(_swapFeeModule common.Address) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.SetSwapFeeModule(&_AerodromeFactory.TransactOpts, _swapFeeModule)
}

// SetUnstakedFeeManager is a paid mutator transaction binding the contract method 0x93ce8627.
//
// Solidity: function setUnstakedFeeManager(address _unstakedFeeManager) returns()
func (_AerodromeFactory *AerodromeFactoryTransactor) SetUnstakedFeeManager(opts *bind.TransactOpts, _unstakedFeeManager common.Address) (*types.Transaction, error) {
	return _AerodromeFactory.contract.Transact(opts, "setUnstakedFeeManager", _unstakedFeeManager)
}

// SetUnstakedFeeManager is a paid mutator transaction binding the contract method 0x93ce8627.
//
// Solidity: function setUnstakedFeeManager(address _unstakedFeeManager) returns()
func (_AerodromeFactory *AerodromeFactorySession) SetUnstakedFeeManager(_unstakedFeeManager common.Address) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.SetUnstakedFeeManager(&_AerodromeFactory.TransactOpts, _unstakedFeeManager)
}

// SetUnstakedFeeManager is a paid mutator transaction binding the contract method 0x93ce8627.
//
// Solidity: function setUnstakedFeeManager(address _unstakedFeeManager) returns()
func (_AerodromeFactory *AerodromeFactoryTransactorSession) SetUnstakedFeeManager(_unstakedFeeManager common.Address) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.SetUnstakedFeeManager(&_AerodromeFactory.TransactOpts, _unstakedFeeManager)
}

// SetUnstakedFeeModule is a paid mutator transaction binding the contract method 0x1b31d878.
//
// Solidity: function setUnstakedFeeModule(address _unstakedFeeModule) returns()
func (_AerodromeFactory *AerodromeFactoryTransactor) SetUnstakedFeeModule(opts *bind.TransactOpts, _unstakedFeeModule common.Address) (*types.Transaction, error) {
	return _AerodromeFactory.contract.Transact(opts, "setUnstakedFeeModule", _unstakedFeeModule)
}

// SetUnstakedFeeModule is a paid mutator transaction binding the contract method 0x1b31d878.
//
// Solidity: function setUnstakedFeeModule(address _unstakedFeeModule) returns()
func (_AerodromeFactory *AerodromeFactorySession) SetUnstakedFeeModule(_unstakedFeeModule common.Address) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.SetUnstakedFeeModule(&_AerodromeFactory.TransactOpts, _unstakedFeeModule)
}

// SetUnstakedFeeModule is a paid mutator transaction binding the contract method 0x1b31d878.
//
// Solidity: function setUnstakedFeeModule(address _unstakedFeeModule) returns()
func (_AerodromeFactory *AerodromeFactoryTransactorSession) SetUnstakedFeeModule(_unstakedFeeModule common.Address) (*types.Transaction, error) {
	return _AerodromeFactory.Contract.SetUnstakedFeeModule(&_AerodromeFactory.TransactOpts, _unstakedFeeModule)
}

// AerodromeFactoryDefaultUnstakedFeeChangedIterator is returned from FilterDefaultUnstakedFeeChanged and is used to iterate over the raw logs and unpacked data for DefaultUnstakedFeeChanged events raised by the AerodromeFactory contract.
type AerodromeFactoryDefaultUnstakedFeeChangedIterator struct {
	Event *AerodromeFactoryDefaultUnstakedFeeChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AerodromeFactoryDefaultUnstakedFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AerodromeFactoryDefaultUnstakedFeeChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AerodromeFactoryDefaultUnstakedFeeChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AerodromeFactoryDefaultUnstakedFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AerodromeFactoryDefaultUnstakedFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AerodromeFactoryDefaultUnstakedFeeChanged represents a DefaultUnstakedFeeChanged event raised by the AerodromeFactory contract.
type AerodromeFactoryDefaultUnstakedFeeChanged struct {
	OldUnstakedFee *big.Int
	NewUnstakedFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefaultUnstakedFeeChanged is a free log retrieval operation binding the contract event 0xcbca61144322b913ada4febfb591864cad7617559d7ee0d3e29b48eb93fcc78e.
//
// Solidity: event DefaultUnstakedFeeChanged(uint24 indexed oldUnstakedFee, uint24 indexed newUnstakedFee)
func (_AerodromeFactory *AerodromeFactoryFilterer) FilterDefaultUnstakedFeeChanged(opts *bind.FilterOpts, oldUnstakedFee []*big.Int, newUnstakedFee []*big.Int) (*AerodromeFactoryDefaultUnstakedFeeChangedIterator, error) {

	var oldUnstakedFeeRule []interface{}
	for _, oldUnstakedFeeItem := range oldUnstakedFee {
		oldUnstakedFeeRule = append(oldUnstakedFeeRule, oldUnstakedFeeItem)
	}
	var newUnstakedFeeRule []interface{}
	for _, newUnstakedFeeItem := range newUnstakedFee {
		newUnstakedFeeRule = append(newUnstakedFeeRule, newUnstakedFeeItem)
	}

	logs, sub, err := _AerodromeFactory.contract.FilterLogs(opts, "DefaultUnstakedFeeChanged", oldUnstakedFeeRule, newUnstakedFeeRule)
	if err != nil {
		return nil, err
	}
	return &AerodromeFactoryDefaultUnstakedFeeChangedIterator{contract: _AerodromeFactory.contract, event: "DefaultUnstakedFeeChanged", logs: logs, sub: sub}, nil
}

// WatchDefaultUnstakedFeeChanged is a free log subscription operation binding the contract event 0xcbca61144322b913ada4febfb591864cad7617559d7ee0d3e29b48eb93fcc78e.
//
// Solidity: event DefaultUnstakedFeeChanged(uint24 indexed oldUnstakedFee, uint24 indexed newUnstakedFee)
func (_AerodromeFactory *AerodromeFactoryFilterer) WatchDefaultUnstakedFeeChanged(opts *bind.WatchOpts, sink chan<- *AerodromeFactoryDefaultUnstakedFeeChanged, oldUnstakedFee []*big.Int, newUnstakedFee []*big.Int) (event.Subscription, error) {

	var oldUnstakedFeeRule []interface{}
	for _, oldUnstakedFeeItem := range oldUnstakedFee {
		oldUnstakedFeeRule = append(oldUnstakedFeeRule, oldUnstakedFeeItem)
	}
	var newUnstakedFeeRule []interface{}
	for _, newUnstakedFeeItem := range newUnstakedFee {
		newUnstakedFeeRule = append(newUnstakedFeeRule, newUnstakedFeeItem)
	}

	logs, sub, err := _AerodromeFactory.contract.WatchLogs(opts, "DefaultUnstakedFeeChanged", oldUnstakedFeeRule, newUnstakedFeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AerodromeFactoryDefaultUnstakedFeeChanged)
				if err := _AerodromeFactory.contract.UnpackLog(event, "DefaultUnstakedFeeChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultUnstakedFeeChanged is a log parse operation binding the contract event 0xcbca61144322b913ada4febfb591864cad7617559d7ee0d3e29b48eb93fcc78e.
//
// Solidity: event DefaultUnstakedFeeChanged(uint24 indexed oldUnstakedFee, uint24 indexed newUnstakedFee)
func (_AerodromeFactory *AerodromeFactoryFilterer) ParseDefaultUnstakedFeeChanged(log types.Log) (*AerodromeFactoryDefaultUnstakedFeeChanged, error) {
	event := new(AerodromeFactoryDefaultUnstakedFeeChanged)
	if err := _AerodromeFactory.contract.UnpackLog(event, "DefaultUnstakedFeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AerodromeFactoryOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the AerodromeFactory contract.
type AerodromeFactoryOwnerChangedIterator struct {
	Event *AerodromeFactoryOwnerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AerodromeFactoryOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AerodromeFactoryOwnerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AerodromeFactoryOwnerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AerodromeFactoryOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AerodromeFactoryOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AerodromeFactoryOwnerChanged represents a OwnerChanged event raised by the AerodromeFactory contract.
type AerodromeFactoryOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_AerodromeFactory *AerodromeFactoryFilterer) FilterOwnerChanged(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*AerodromeFactoryOwnerChangedIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AerodromeFactory.contract.FilterLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AerodromeFactoryOwnerChangedIterator{contract: _AerodromeFactory.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_AerodromeFactory *AerodromeFactoryFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *AerodromeFactoryOwnerChanged, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AerodromeFactory.contract.WatchLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AerodromeFactoryOwnerChanged)
				if err := _AerodromeFactory.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_AerodromeFactory *AerodromeFactoryFilterer) ParseOwnerChanged(log types.Log) (*AerodromeFactoryOwnerChanged, error) {
	event := new(AerodromeFactoryOwnerChanged)
	if err := _AerodromeFactory.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AerodromeFactoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the AerodromeFactory contract.
type AerodromeFactoryPoolCreatedIterator struct {
	Event *AerodromeFactoryPoolCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AerodromeFactoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AerodromeFactoryPoolCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AerodromeFactoryPoolCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AerodromeFactoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AerodromeFactoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AerodromeFactoryPoolCreated represents a PoolCreated event raised by the AerodromeFactory contract.
type AerodromeFactoryPoolCreated struct {
	Token0      common.Address
	Token1      common.Address
	TickSpacing *big.Int
	Pool        common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0xab0d57f0df537bb25e80245ef7748fa62353808c54d6e528a9dd20887aed9ac2.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, int24 indexed tickSpacing, address pool)
func (_AerodromeFactory *AerodromeFactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address, tickSpacing []*big.Int) (*AerodromeFactoryPoolCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var tickSpacingRule []interface{}
	for _, tickSpacingItem := range tickSpacing {
		tickSpacingRule = append(tickSpacingRule, tickSpacingItem)
	}

	logs, sub, err := _AerodromeFactory.contract.FilterLogs(opts, "PoolCreated", token0Rule, token1Rule, tickSpacingRule)
	if err != nil {
		return nil, err
	}
	return &AerodromeFactoryPoolCreatedIterator{contract: _AerodromeFactory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0xab0d57f0df537bb25e80245ef7748fa62353808c54d6e528a9dd20887aed9ac2.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, int24 indexed tickSpacing, address pool)
func (_AerodromeFactory *AerodromeFactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *AerodromeFactoryPoolCreated, token0 []common.Address, token1 []common.Address, tickSpacing []*big.Int) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var tickSpacingRule []interface{}
	for _, tickSpacingItem := range tickSpacing {
		tickSpacingRule = append(tickSpacingRule, tickSpacingItem)
	}

	logs, sub, err := _AerodromeFactory.contract.WatchLogs(opts, "PoolCreated", token0Rule, token1Rule, tickSpacingRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AerodromeFactoryPoolCreated)
				if err := _AerodromeFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolCreated is a log parse operation binding the contract event 0xab0d57f0df537bb25e80245ef7748fa62353808c54d6e528a9dd20887aed9ac2.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, int24 indexed tickSpacing, address pool)
func (_AerodromeFactory *AerodromeFactoryFilterer) ParsePoolCreated(log types.Log) (*AerodromeFactoryPoolCreated, error) {
	event := new(AerodromeFactoryPoolCreated)
	if err := _AerodromeFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AerodromeFactorySwapFeeManagerChangedIterator is returned from FilterSwapFeeManagerChanged and is used to iterate over the raw logs and unpacked data for SwapFeeManagerChanged events raised by the AerodromeFactory contract.
type AerodromeFactorySwapFeeManagerChangedIterator struct {
	Event *AerodromeFactorySwapFeeManagerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AerodromeFactorySwapFeeManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AerodromeFactorySwapFeeManagerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AerodromeFactorySwapFeeManagerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AerodromeFactorySwapFeeManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AerodromeFactorySwapFeeManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AerodromeFactorySwapFeeManagerChanged represents a SwapFeeManagerChanged event raised by the AerodromeFactory contract.
type AerodromeFactorySwapFeeManagerChanged struct {
	OldFeeManager common.Address
	NewFeeManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSwapFeeManagerChanged is a free log retrieval operation binding the contract event 0x7ae0007229b3333719d97e8ef5829c888f560776012974f87409c158e5b7eb91.
//
// Solidity: event SwapFeeManagerChanged(address indexed oldFeeManager, address indexed newFeeManager)
func (_AerodromeFactory *AerodromeFactoryFilterer) FilterSwapFeeManagerChanged(opts *bind.FilterOpts, oldFeeManager []common.Address, newFeeManager []common.Address) (*AerodromeFactorySwapFeeManagerChangedIterator, error) {

	var oldFeeManagerRule []interface{}
	for _, oldFeeManagerItem := range oldFeeManager {
		oldFeeManagerRule = append(oldFeeManagerRule, oldFeeManagerItem)
	}
	var newFeeManagerRule []interface{}
	for _, newFeeManagerItem := range newFeeManager {
		newFeeManagerRule = append(newFeeManagerRule, newFeeManagerItem)
	}

	logs, sub, err := _AerodromeFactory.contract.FilterLogs(opts, "SwapFeeManagerChanged", oldFeeManagerRule, newFeeManagerRule)
	if err != nil {
		return nil, err
	}
	return &AerodromeFactorySwapFeeManagerChangedIterator{contract: _AerodromeFactory.contract, event: "SwapFeeManagerChanged", logs: logs, sub: sub}, nil
}

// WatchSwapFeeManagerChanged is a free log subscription operation binding the contract event 0x7ae0007229b3333719d97e8ef5829c888f560776012974f87409c158e5b7eb91.
//
// Solidity: event SwapFeeManagerChanged(address indexed oldFeeManager, address indexed newFeeManager)
func (_AerodromeFactory *AerodromeFactoryFilterer) WatchSwapFeeManagerChanged(opts *bind.WatchOpts, sink chan<- *AerodromeFactorySwapFeeManagerChanged, oldFeeManager []common.Address, newFeeManager []common.Address) (event.Subscription, error) {

	var oldFeeManagerRule []interface{}
	for _, oldFeeManagerItem := range oldFeeManager {
		oldFeeManagerRule = append(oldFeeManagerRule, oldFeeManagerItem)
	}
	var newFeeManagerRule []interface{}
	for _, newFeeManagerItem := range newFeeManager {
		newFeeManagerRule = append(newFeeManagerRule, newFeeManagerItem)
	}

	logs, sub, err := _AerodromeFactory.contract.WatchLogs(opts, "SwapFeeManagerChanged", oldFeeManagerRule, newFeeManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AerodromeFactorySwapFeeManagerChanged)
				if err := _AerodromeFactory.contract.UnpackLog(event, "SwapFeeManagerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwapFeeManagerChanged is a log parse operation binding the contract event 0x7ae0007229b3333719d97e8ef5829c888f560776012974f87409c158e5b7eb91.
//
// Solidity: event SwapFeeManagerChanged(address indexed oldFeeManager, address indexed newFeeManager)
func (_AerodromeFactory *AerodromeFactoryFilterer) ParseSwapFeeManagerChanged(log types.Log) (*AerodromeFactorySwapFeeManagerChanged, error) {
	event := new(AerodromeFactorySwapFeeManagerChanged)
	if err := _AerodromeFactory.contract.UnpackLog(event, "SwapFeeManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AerodromeFactorySwapFeeModuleChangedIterator is returned from FilterSwapFeeModuleChanged and is used to iterate over the raw logs and unpacked data for SwapFeeModuleChanged events raised by the AerodromeFactory contract.
type AerodromeFactorySwapFeeModuleChangedIterator struct {
	Event *AerodromeFactorySwapFeeModuleChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AerodromeFactorySwapFeeModuleChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AerodromeFactorySwapFeeModuleChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AerodromeFactorySwapFeeModuleChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AerodromeFactorySwapFeeModuleChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AerodromeFactorySwapFeeModuleChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AerodromeFactorySwapFeeModuleChanged represents a SwapFeeModuleChanged event raised by the AerodromeFactory contract.
type AerodromeFactorySwapFeeModuleChanged struct {
	OldFeeModule common.Address
	NewFeeModule common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwapFeeModuleChanged is a free log retrieval operation binding the contract event 0xdf24ed64a7bcd761cf1132e79f94ea269a1d570e7a6ca0ab99a8f5ccd6f5022f.
//
// Solidity: event SwapFeeModuleChanged(address indexed oldFeeModule, address indexed newFeeModule)
func (_AerodromeFactory *AerodromeFactoryFilterer) FilterSwapFeeModuleChanged(opts *bind.FilterOpts, oldFeeModule []common.Address, newFeeModule []common.Address) (*AerodromeFactorySwapFeeModuleChangedIterator, error) {

	var oldFeeModuleRule []interface{}
	for _, oldFeeModuleItem := range oldFeeModule {
		oldFeeModuleRule = append(oldFeeModuleRule, oldFeeModuleItem)
	}
	var newFeeModuleRule []interface{}
	for _, newFeeModuleItem := range newFeeModule {
		newFeeModuleRule = append(newFeeModuleRule, newFeeModuleItem)
	}

	logs, sub, err := _AerodromeFactory.contract.FilterLogs(opts, "SwapFeeModuleChanged", oldFeeModuleRule, newFeeModuleRule)
	if err != nil {
		return nil, err
	}
	return &AerodromeFactorySwapFeeModuleChangedIterator{contract: _AerodromeFactory.contract, event: "SwapFeeModuleChanged", logs: logs, sub: sub}, nil
}

// WatchSwapFeeModuleChanged is a free log subscription operation binding the contract event 0xdf24ed64a7bcd761cf1132e79f94ea269a1d570e7a6ca0ab99a8f5ccd6f5022f.
//
// Solidity: event SwapFeeModuleChanged(address indexed oldFeeModule, address indexed newFeeModule)
func (_AerodromeFactory *AerodromeFactoryFilterer) WatchSwapFeeModuleChanged(opts *bind.WatchOpts, sink chan<- *AerodromeFactorySwapFeeModuleChanged, oldFeeModule []common.Address, newFeeModule []common.Address) (event.Subscription, error) {

	var oldFeeModuleRule []interface{}
	for _, oldFeeModuleItem := range oldFeeModule {
		oldFeeModuleRule = append(oldFeeModuleRule, oldFeeModuleItem)
	}
	var newFeeModuleRule []interface{}
	for _, newFeeModuleItem := range newFeeModule {
		newFeeModuleRule = append(newFeeModuleRule, newFeeModuleItem)
	}

	logs, sub, err := _AerodromeFactory.contract.WatchLogs(opts, "SwapFeeModuleChanged", oldFeeModuleRule, newFeeModuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AerodromeFactorySwapFeeModuleChanged)
				if err := _AerodromeFactory.contract.UnpackLog(event, "SwapFeeModuleChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwapFeeModuleChanged is a log parse operation binding the contract event 0xdf24ed64a7bcd761cf1132e79f94ea269a1d570e7a6ca0ab99a8f5ccd6f5022f.
//
// Solidity: event SwapFeeModuleChanged(address indexed oldFeeModule, address indexed newFeeModule)
func (_AerodromeFactory *AerodromeFactoryFilterer) ParseSwapFeeModuleChanged(log types.Log) (*AerodromeFactorySwapFeeModuleChanged, error) {
	event := new(AerodromeFactorySwapFeeModuleChanged)
	if err := _AerodromeFactory.contract.UnpackLog(event, "SwapFeeModuleChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AerodromeFactoryTickSpacingEnabledIterator is returned from FilterTickSpacingEnabled and is used to iterate over the raw logs and unpacked data for TickSpacingEnabled events raised by the AerodromeFactory contract.
type AerodromeFactoryTickSpacingEnabledIterator struct {
	Event *AerodromeFactoryTickSpacingEnabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AerodromeFactoryTickSpacingEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AerodromeFactoryTickSpacingEnabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AerodromeFactoryTickSpacingEnabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AerodromeFactoryTickSpacingEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AerodromeFactoryTickSpacingEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AerodromeFactoryTickSpacingEnabled represents a TickSpacingEnabled event raised by the AerodromeFactory contract.
type AerodromeFactoryTickSpacingEnabled struct {
	TickSpacing *big.Int
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTickSpacingEnabled is a free log retrieval operation binding the contract event 0xebafae466a4a780a1d87f5fab2f52fad33be9151a7f69d099e8934c8de85b747.
//
// Solidity: event TickSpacingEnabled(int24 indexed tickSpacing, uint24 indexed fee)
func (_AerodromeFactory *AerodromeFactoryFilterer) FilterTickSpacingEnabled(opts *bind.FilterOpts, tickSpacing []*big.Int, fee []*big.Int) (*AerodromeFactoryTickSpacingEnabledIterator, error) {

	var tickSpacingRule []interface{}
	for _, tickSpacingItem := range tickSpacing {
		tickSpacingRule = append(tickSpacingRule, tickSpacingItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _AerodromeFactory.contract.FilterLogs(opts, "TickSpacingEnabled", tickSpacingRule, feeRule)
	if err != nil {
		return nil, err
	}
	return &AerodromeFactoryTickSpacingEnabledIterator{contract: _AerodromeFactory.contract, event: "TickSpacingEnabled", logs: logs, sub: sub}, nil
}

// WatchTickSpacingEnabled is a free log subscription operation binding the contract event 0xebafae466a4a780a1d87f5fab2f52fad33be9151a7f69d099e8934c8de85b747.
//
// Solidity: event TickSpacingEnabled(int24 indexed tickSpacing, uint24 indexed fee)
func (_AerodromeFactory *AerodromeFactoryFilterer) WatchTickSpacingEnabled(opts *bind.WatchOpts, sink chan<- *AerodromeFactoryTickSpacingEnabled, tickSpacing []*big.Int, fee []*big.Int) (event.Subscription, error) {

	var tickSpacingRule []interface{}
	for _, tickSpacingItem := range tickSpacing {
		tickSpacingRule = append(tickSpacingRule, tickSpacingItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _AerodromeFactory.contract.WatchLogs(opts, "TickSpacingEnabled", tickSpacingRule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AerodromeFactoryTickSpacingEnabled)
				if err := _AerodromeFactory.contract.UnpackLog(event, "TickSpacingEnabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTickSpacingEnabled is a log parse operation binding the contract event 0xebafae466a4a780a1d87f5fab2f52fad33be9151a7f69d099e8934c8de85b747.
//
// Solidity: event TickSpacingEnabled(int24 indexed tickSpacing, uint24 indexed fee)
func (_AerodromeFactory *AerodromeFactoryFilterer) ParseTickSpacingEnabled(log types.Log) (*AerodromeFactoryTickSpacingEnabled, error) {
	event := new(AerodromeFactoryTickSpacingEnabled)
	if err := _AerodromeFactory.contract.UnpackLog(event, "TickSpacingEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AerodromeFactoryUnstakedFeeManagerChangedIterator is returned from FilterUnstakedFeeManagerChanged and is used to iterate over the raw logs and unpacked data for UnstakedFeeManagerChanged events raised by the AerodromeFactory contract.
type AerodromeFactoryUnstakedFeeManagerChangedIterator struct {
	Event *AerodromeFactoryUnstakedFeeManagerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AerodromeFactoryUnstakedFeeManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AerodromeFactoryUnstakedFeeManagerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AerodromeFactoryUnstakedFeeManagerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AerodromeFactoryUnstakedFeeManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AerodromeFactoryUnstakedFeeManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AerodromeFactoryUnstakedFeeManagerChanged represents a UnstakedFeeManagerChanged event raised by the AerodromeFactory contract.
type AerodromeFactoryUnstakedFeeManagerChanged struct {
	OldFeeManager common.Address
	NewFeeManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnstakedFeeManagerChanged is a free log retrieval operation binding the contract event 0x3d7ebe96182c99643ca0c997a416a2a3409baab225f85f50c29fcf0591c820c1.
//
// Solidity: event UnstakedFeeManagerChanged(address indexed oldFeeManager, address indexed newFeeManager)
func (_AerodromeFactory *AerodromeFactoryFilterer) FilterUnstakedFeeManagerChanged(opts *bind.FilterOpts, oldFeeManager []common.Address, newFeeManager []common.Address) (*AerodromeFactoryUnstakedFeeManagerChangedIterator, error) {

	var oldFeeManagerRule []interface{}
	for _, oldFeeManagerItem := range oldFeeManager {
		oldFeeManagerRule = append(oldFeeManagerRule, oldFeeManagerItem)
	}
	var newFeeManagerRule []interface{}
	for _, newFeeManagerItem := range newFeeManager {
		newFeeManagerRule = append(newFeeManagerRule, newFeeManagerItem)
	}

	logs, sub, err := _AerodromeFactory.contract.FilterLogs(opts, "UnstakedFeeManagerChanged", oldFeeManagerRule, newFeeManagerRule)
	if err != nil {
		return nil, err
	}
	return &AerodromeFactoryUnstakedFeeManagerChangedIterator{contract: _AerodromeFactory.contract, event: "UnstakedFeeManagerChanged", logs: logs, sub: sub}, nil
}

// WatchUnstakedFeeManagerChanged is a free log subscription operation binding the contract event 0x3d7ebe96182c99643ca0c997a416a2a3409baab225f85f50c29fcf0591c820c1.
//
// Solidity: event UnstakedFeeManagerChanged(address indexed oldFeeManager, address indexed newFeeManager)
func (_AerodromeFactory *AerodromeFactoryFilterer) WatchUnstakedFeeManagerChanged(opts *bind.WatchOpts, sink chan<- *AerodromeFactoryUnstakedFeeManagerChanged, oldFeeManager []common.Address, newFeeManager []common.Address) (event.Subscription, error) {

	var oldFeeManagerRule []interface{}
	for _, oldFeeManagerItem := range oldFeeManager {
		oldFeeManagerRule = append(oldFeeManagerRule, oldFeeManagerItem)
	}
	var newFeeManagerRule []interface{}
	for _, newFeeManagerItem := range newFeeManager {
		newFeeManagerRule = append(newFeeManagerRule, newFeeManagerItem)
	}

	logs, sub, err := _AerodromeFactory.contract.WatchLogs(opts, "UnstakedFeeManagerChanged", oldFeeManagerRule, newFeeManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AerodromeFactoryUnstakedFeeManagerChanged)
				if err := _AerodromeFactory.contract.UnpackLog(event, "UnstakedFeeManagerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnstakedFeeManagerChanged is a log parse operation binding the contract event 0x3d7ebe96182c99643ca0c997a416a2a3409baab225f85f50c29fcf0591c820c1.
//
// Solidity: event UnstakedFeeManagerChanged(address indexed oldFeeManager, address indexed newFeeManager)
func (_AerodromeFactory *AerodromeFactoryFilterer) ParseUnstakedFeeManagerChanged(log types.Log) (*AerodromeFactoryUnstakedFeeManagerChanged, error) {
	event := new(AerodromeFactoryUnstakedFeeManagerChanged)
	if err := _AerodromeFactory.contract.UnpackLog(event, "UnstakedFeeManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AerodromeFactoryUnstakedFeeModuleChangedIterator is returned from FilterUnstakedFeeModuleChanged and is used to iterate over the raw logs and unpacked data for UnstakedFeeModuleChanged events raised by the AerodromeFactory contract.
type AerodromeFactoryUnstakedFeeModuleChangedIterator struct {
	Event *AerodromeFactoryUnstakedFeeModuleChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AerodromeFactoryUnstakedFeeModuleChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AerodromeFactoryUnstakedFeeModuleChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AerodromeFactoryUnstakedFeeModuleChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AerodromeFactoryUnstakedFeeModuleChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AerodromeFactoryUnstakedFeeModuleChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AerodromeFactoryUnstakedFeeModuleChanged represents a UnstakedFeeModuleChanged event raised by the AerodromeFactory contract.
type AerodromeFactoryUnstakedFeeModuleChanged struct {
	OldFeeModule common.Address
	NewFeeModule common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnstakedFeeModuleChanged is a free log retrieval operation binding the contract event 0x6520f404f3831947cee8673060459cdfb181b7332aa7580bcce9bf90ef1f0e20.
//
// Solidity: event UnstakedFeeModuleChanged(address indexed oldFeeModule, address indexed newFeeModule)
func (_AerodromeFactory *AerodromeFactoryFilterer) FilterUnstakedFeeModuleChanged(opts *bind.FilterOpts, oldFeeModule []common.Address, newFeeModule []common.Address) (*AerodromeFactoryUnstakedFeeModuleChangedIterator, error) {

	var oldFeeModuleRule []interface{}
	for _, oldFeeModuleItem := range oldFeeModule {
		oldFeeModuleRule = append(oldFeeModuleRule, oldFeeModuleItem)
	}
	var newFeeModuleRule []interface{}
	for _, newFeeModuleItem := range newFeeModule {
		newFeeModuleRule = append(newFeeModuleRule, newFeeModuleItem)
	}

	logs, sub, err := _AerodromeFactory.contract.FilterLogs(opts, "UnstakedFeeModuleChanged", oldFeeModuleRule, newFeeModuleRule)
	if err != nil {
		return nil, err
	}
	return &AerodromeFactoryUnstakedFeeModuleChangedIterator{contract: _AerodromeFactory.contract, event: "UnstakedFeeModuleChanged", logs: logs, sub: sub}, nil
}

// WatchUnstakedFeeModuleChanged is a free log subscription operation binding the contract event 0x6520f404f3831947cee8673060459cdfb181b7332aa7580bcce9bf90ef1f0e20.
//
// Solidity: event UnstakedFeeModuleChanged(address indexed oldFeeModule, address indexed newFeeModule)
func (_AerodromeFactory *AerodromeFactoryFilterer) WatchUnstakedFeeModuleChanged(opts *bind.WatchOpts, sink chan<- *AerodromeFactoryUnstakedFeeModuleChanged, oldFeeModule []common.Address, newFeeModule []common.Address) (event.Subscription, error) {

	var oldFeeModuleRule []interface{}
	for _, oldFeeModuleItem := range oldFeeModule {
		oldFeeModuleRule = append(oldFeeModuleRule, oldFeeModuleItem)
	}
	var newFeeModuleRule []interface{}
	for _, newFeeModuleItem := range newFeeModule {
		newFeeModuleRule = append(newFeeModuleRule, newFeeModuleItem)
	}

	logs, sub, err := _AerodromeFactory.contract.WatchLogs(opts, "UnstakedFeeModuleChanged", oldFeeModuleRule, newFeeModuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AerodromeFactoryUnstakedFeeModuleChanged)
				if err := _AerodromeFactory.contract.UnpackLog(event, "UnstakedFeeModuleChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnstakedFeeModuleChanged is a log parse operation binding the contract event 0x6520f404f3831947cee8673060459cdfb181b7332aa7580bcce9bf90ef1f0e20.
//
// Solidity: event UnstakedFeeModuleChanged(address indexed oldFeeModule, address indexed newFeeModule)
func (_AerodromeFactory *AerodromeFactoryFilterer) ParseUnstakedFeeModuleChanged(log types.Log) (*AerodromeFactoryUnstakedFeeModuleChanged, error) {
	event := new(AerodromeFactoryUnstakedFeeModuleChanged)
	if err := _AerodromeFactory.contract.UnpackLog(event, "UnstakedFeeModuleChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
