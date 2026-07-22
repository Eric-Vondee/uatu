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

// V3FactoryMetaData contains all meta data concerning the V3Factory contract.
var V3FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// V3FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use V3FactoryMetaData.ABI instead.
var V3FactoryABI = V3FactoryMetaData.ABI

// V3Factory is an auto generated Go binding around an Ethereum contract.
type V3Factory struct {
	V3FactoryCaller     // Read-only binding to the contract
	V3FactoryTransactor // Write-only binding to the contract
	V3FactoryFilterer   // Log filterer for contract events
}

// V3FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type V3FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V3FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V3FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V3FactorySession struct {
	Contract     *V3Factory        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V3FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V3FactoryCallerSession struct {
	Contract *V3FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// V3FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V3FactoryTransactorSession struct {
	Contract     *V3FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// V3FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type V3FactoryRaw struct {
	Contract *V3Factory // Generic contract binding to access the raw methods on
}

// V3FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V3FactoryCallerRaw struct {
	Contract *V3FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// V3FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V3FactoryTransactorRaw struct {
	Contract *V3FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV3Factory creates a new instance of V3Factory, bound to a specific deployed contract.
func NewV3Factory(address common.Address, backend bind.ContractBackend) (*V3Factory, error) {
	contract, err := bindV3Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V3Factory{V3FactoryCaller: V3FactoryCaller{contract: contract}, V3FactoryTransactor: V3FactoryTransactor{contract: contract}, V3FactoryFilterer: V3FactoryFilterer{contract: contract}}, nil
}

// NewV3FactoryCaller creates a new read-only instance of V3Factory, bound to a specific deployed contract.
func NewV3FactoryCaller(address common.Address, caller bind.ContractCaller) (*V3FactoryCaller, error) {
	contract, err := bindV3Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V3FactoryCaller{contract: contract}, nil
}

// NewV3FactoryTransactor creates a new write-only instance of V3Factory, bound to a specific deployed contract.
func NewV3FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*V3FactoryTransactor, error) {
	contract, err := bindV3Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V3FactoryTransactor{contract: contract}, nil
}

// NewV3FactoryFilterer creates a new log filterer instance of V3Factory, bound to a specific deployed contract.
func NewV3FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*V3FactoryFilterer, error) {
	contract, err := bindV3Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V3FactoryFilterer{contract: contract}, nil
}

// bindV3Factory binds a generic wrapper to an already deployed contract.
func bindV3Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := V3FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3Factory *V3FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3Factory.Contract.V3FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3Factory *V3FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3Factory.Contract.V3FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3Factory *V3FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3Factory.Contract.V3FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3Factory *V3FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3Factory *V3FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3Factory *V3FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3Factory.Contract.contract.Transact(opts, method, params...)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_V3Factory *V3FactoryCaller) GetPool(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _V3Factory.contract.Call(opts, &out, "getPool", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_V3Factory *V3FactorySession) GetPool(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	return _V3Factory.Contract.GetPool(&_V3Factory.CallOpts, arg0, arg1, arg2)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_V3Factory *V3FactoryCallerSession) GetPool(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	return _V3Factory.Contract.GetPool(&_V3Factory.CallOpts, arg0, arg1, arg2)
}
