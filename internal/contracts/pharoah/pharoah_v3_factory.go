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

// V3PharoahFactoryMetaData contains all meta data concerning the V3PharoahFactory contract.
var V3PharoahFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accessHub\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ADDRESS_ZERO\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"F0\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FEE_TOO_LARGE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IDENTICAL_TOKENS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NOT_ACCESSHUB\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PAIR_EXISTS\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"newFee\",\"type\":\"uint24\"}],\"name\":\"FeeAdjustment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldFeeCollector\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeCollector\",\"type\":\"address\"}],\"name\":\"FeeCollectorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"feeProtocolOld\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"feeProtocolNew\",\"type\":\"uint24\"}],\"name\":\"SetFeeProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"feeProtocolOld\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"feeProtocolNew\",\"type\":\"uint24\"}],\"name\":\"SetPoolFeeProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"TickSpacingEnabled\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_FEE_FLAG\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessHub\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"uint24\",\"name\":\"initialFee\",\"type\":\"uint24\"}],\"name\":\"enableTickSpacing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeProtocol\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"gaugeFeeSplitEnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ramsesV3PoolDeployer\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"isPairV3\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isV3\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parameters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"poolFeeProtocol\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"__poolFeeProtocol\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ramsesV3PoolDeployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_fee\",\"type\":\"uint24\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeCollector\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_feeProtocol\",\"type\":\"uint24\"}],\"name\":\"setFeeProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_feeProtocol\",\"type\":\"uint24\"}],\"name\":\"setPoolFeeProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"setVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"name\":\"tickSpacingInitialFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"initialFee\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// V3PharoahFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use V3PharoahFactoryMetaData.ABI instead.
var V3PharoahFactoryABI = V3PharoahFactoryMetaData.ABI

// V3PharoahFactory is an auto generated Go binding around an Ethereum contract.
type V3PharoahFactory struct {
	V3PharoahFactoryCaller     // Read-only binding to the contract
	V3PharoahFactoryTransactor // Write-only binding to the contract
	V3PharoahFactoryFilterer   // Log filterer for contract events
}

// V3PharoahFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type V3PharoahFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3PharoahFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V3PharoahFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3PharoahFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V3PharoahFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3PharoahFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V3PharoahFactorySession struct {
	Contract     *V3PharoahFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V3PharoahFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V3PharoahFactoryCallerSession struct {
	Contract *V3PharoahFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// V3PharoahFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V3PharoahFactoryTransactorSession struct {
	Contract     *V3PharoahFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// V3PharoahFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type V3PharoahFactoryRaw struct {
	Contract *V3PharoahFactory // Generic contract binding to access the raw methods on
}

// V3PharoahFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V3PharoahFactoryCallerRaw struct {
	Contract *V3PharoahFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// V3PharoahFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V3PharoahFactoryTransactorRaw struct {
	Contract *V3PharoahFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV3PharoahFactory creates a new instance of V3PharoahFactory, bound to a specific deployed contract.
func NewV3PharoahFactory(address common.Address, backend bind.ContractBackend) (*V3PharoahFactory, error) {
	contract, err := bindV3PharoahFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V3PharoahFactory{V3PharoahFactoryCaller: V3PharoahFactoryCaller{contract: contract}, V3PharoahFactoryTransactor: V3PharoahFactoryTransactor{contract: contract}, V3PharoahFactoryFilterer: V3PharoahFactoryFilterer{contract: contract}}, nil
}

// NewV3PharoahFactoryCaller creates a new read-only instance of V3PharoahFactory, bound to a specific deployed contract.
func NewV3PharoahFactoryCaller(address common.Address, caller bind.ContractCaller) (*V3PharoahFactoryCaller, error) {
	contract, err := bindV3PharoahFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V3PharoahFactoryCaller{contract: contract}, nil
}

// NewV3PharoahFactoryTransactor creates a new write-only instance of V3PharoahFactory, bound to a specific deployed contract.
func NewV3PharoahFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*V3PharoahFactoryTransactor, error) {
	contract, err := bindV3PharoahFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V3PharoahFactoryTransactor{contract: contract}, nil
}

// NewV3PharoahFactoryFilterer creates a new log filterer instance of V3PharoahFactory, bound to a specific deployed contract.
func NewV3PharoahFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*V3PharoahFactoryFilterer, error) {
	contract, err := bindV3PharoahFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V3PharoahFactoryFilterer{contract: contract}, nil
}

// bindV3PharoahFactory binds a generic wrapper to an already deployed contract.
func bindV3PharoahFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := V3PharoahFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3PharoahFactory *V3PharoahFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3PharoahFactory.Contract.V3PharoahFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3PharoahFactory *V3PharoahFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3PharoahFactory.Contract.V3PharoahFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3PharoahFactory *V3PharoahFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3PharoahFactory.Contract.V3PharoahFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3PharoahFactory *V3PharoahFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3PharoahFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3PharoahFactory *V3PharoahFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3PharoahFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3PharoahFactory *V3PharoahFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3PharoahFactory.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTFEEFLAG is a free data retrieval call binding the contract method 0xdd6c5a6b.
//
// Solidity: function DEFAULT_FEE_FLAG() view returns(uint24)
func (_V3PharoahFactory *V3PharoahFactoryCaller) DEFAULTFEEFLAG(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _V3PharoahFactory.contract.Call(opts, &out, "DEFAULT_FEE_FLAG")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTFEEFLAG is a free data retrieval call binding the contract method 0xdd6c5a6b.
//
// Solidity: function DEFAULT_FEE_FLAG() view returns(uint24)
func (_V3PharoahFactory *V3PharoahFactorySession) DEFAULTFEEFLAG() (*big.Int, error) {
	return _V3PharoahFactory.Contract.DEFAULTFEEFLAG(&_V3PharoahFactory.CallOpts)
}

// DEFAULTFEEFLAG is a free data retrieval call binding the contract method 0xdd6c5a6b.
//
// Solidity: function DEFAULT_FEE_FLAG() view returns(uint24)
func (_V3PharoahFactory *V3PharoahFactoryCallerSession) DEFAULTFEEFLAG() (*big.Int, error) {
	return _V3PharoahFactory.Contract.DEFAULTFEEFLAG(&_V3PharoahFactory.CallOpts)
}

// AccessHub is a free data retrieval call binding the contract method 0xe7589b39.
//
// Solidity: function accessHub() view returns(address)
func (_V3PharoahFactory *V3PharoahFactoryCaller) AccessHub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3PharoahFactory.contract.Call(opts, &out, "accessHub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccessHub is a free data retrieval call binding the contract method 0xe7589b39.
//
// Solidity: function accessHub() view returns(address)
func (_V3PharoahFactory *V3PharoahFactorySession) AccessHub() (common.Address, error) {
	return _V3PharoahFactory.Contract.AccessHub(&_V3PharoahFactory.CallOpts)
}

// AccessHub is a free data retrieval call binding the contract method 0xe7589b39.
//
// Solidity: function accessHub() view returns(address)
func (_V3PharoahFactory *V3PharoahFactoryCallerSession) AccessHub() (common.Address, error) {
	return _V3PharoahFactory.Contract.AccessHub(&_V3PharoahFactory.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_V3PharoahFactory *V3PharoahFactoryCaller) FeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3PharoahFactory.contract.Call(opts, &out, "feeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_V3PharoahFactory *V3PharoahFactorySession) FeeCollector() (common.Address, error) {
	return _V3PharoahFactory.Contract.FeeCollector(&_V3PharoahFactory.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_V3PharoahFactory *V3PharoahFactoryCallerSession) FeeCollector() (common.Address, error) {
	return _V3PharoahFactory.Contract.FeeCollector(&_V3PharoahFactory.CallOpts)
}

// FeeProtocol is a free data retrieval call binding the contract method 0x527eb4bc.
//
// Solidity: function feeProtocol() view returns(uint24)
func (_V3PharoahFactory *V3PharoahFactoryCaller) FeeProtocol(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _V3PharoahFactory.contract.Call(opts, &out, "feeProtocol")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeProtocol is a free data retrieval call binding the contract method 0x527eb4bc.
//
// Solidity: function feeProtocol() view returns(uint24)
func (_V3PharoahFactory *V3PharoahFactorySession) FeeProtocol() (*big.Int, error) {
	return _V3PharoahFactory.Contract.FeeProtocol(&_V3PharoahFactory.CallOpts)
}

// FeeProtocol is a free data retrieval call binding the contract method 0x527eb4bc.
//
// Solidity: function feeProtocol() view returns(uint24)
func (_V3PharoahFactory *V3PharoahFactoryCallerSession) FeeProtocol() (*big.Int, error) {
	return _V3PharoahFactory.Contract.FeeProtocol(&_V3PharoahFactory.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0x28af8d0b.
//
// Solidity: function getPool(address tokenA, address tokenB, int24 tickSpacing) view returns(address pool)
func (_V3PharoahFactory *V3PharoahFactoryCaller) GetPool(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, tickSpacing *big.Int) (common.Address, error) {
	var out []interface{}
	err := _V3PharoahFactory.contract.Call(opts, &out, "getPool", tokenA, tokenB, tickSpacing)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x28af8d0b.
//
// Solidity: function getPool(address tokenA, address tokenB, int24 tickSpacing) view returns(address pool)
func (_V3PharoahFactory *V3PharoahFactorySession) GetPool(tokenA common.Address, tokenB common.Address, tickSpacing *big.Int) (common.Address, error) {
	return _V3PharoahFactory.Contract.GetPool(&_V3PharoahFactory.CallOpts, tokenA, tokenB, tickSpacing)
}

// GetPool is a free data retrieval call binding the contract method 0x28af8d0b.
//
// Solidity: function getPool(address tokenA, address tokenB, int24 tickSpacing) view returns(address pool)
func (_V3PharoahFactory *V3PharoahFactoryCallerSession) GetPool(tokenA common.Address, tokenB common.Address, tickSpacing *big.Int) (common.Address, error) {
	return _V3PharoahFactory.Contract.GetPool(&_V3PharoahFactory.CallOpts, tokenA, tokenB, tickSpacing)
}

// IsPairV3 is a free data retrieval call binding the contract method 0x42378e95.
//
// Solidity: function isPairV3(address pool) view returns(bool isV3)
func (_V3PharoahFactory *V3PharoahFactoryCaller) IsPairV3(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var out []interface{}
	err := _V3PharoahFactory.contract.Call(opts, &out, "isPairV3", pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPairV3 is a free data retrieval call binding the contract method 0x42378e95.
//
// Solidity: function isPairV3(address pool) view returns(bool isV3)
func (_V3PharoahFactory *V3PharoahFactorySession) IsPairV3(pool common.Address) (bool, error) {
	return _V3PharoahFactory.Contract.IsPairV3(&_V3PharoahFactory.CallOpts, pool)
}

// IsPairV3 is a free data retrieval call binding the contract method 0x42378e95.
//
// Solidity: function isPairV3(address pool) view returns(bool isV3)
func (_V3PharoahFactory *V3PharoahFactoryCallerSession) IsPairV3(pool common.Address) (bool, error) {
	return _V3PharoahFactory.Contract.IsPairV3(&_V3PharoahFactory.CallOpts, pool)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_V3PharoahFactory *V3PharoahFactoryCaller) Parameters(opts *bind.CallOpts) (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	var out []interface{}
	err := _V3PharoahFactory.contract.Call(opts, &out, "parameters")

	outstruct := new(struct {
		Factory     common.Address
		Token0      common.Address
		Token1      common.Address
		Fee         *big.Int
		TickSpacing *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Factory = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token0 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TickSpacing = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_V3PharoahFactory *V3PharoahFactorySession) Parameters() (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	return _V3PharoahFactory.Contract.Parameters(&_V3PharoahFactory.CallOpts)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_V3PharoahFactory *V3PharoahFactoryCallerSession) Parameters() (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	return _V3PharoahFactory.Contract.Parameters(&_V3PharoahFactory.CallOpts)
}

// PoolFeeProtocol is a free data retrieval call binding the contract method 0xebb0d9f7.
//
// Solidity: function poolFeeProtocol(address pool) view returns(uint24 __poolFeeProtocol)
func (_V3PharoahFactory *V3PharoahFactoryCaller) PoolFeeProtocol(opts *bind.CallOpts, pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _V3PharoahFactory.contract.Call(opts, &out, "poolFeeProtocol", pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolFeeProtocol is a free data retrieval call binding the contract method 0xebb0d9f7.
//
// Solidity: function poolFeeProtocol(address pool) view returns(uint24 __poolFeeProtocol)
func (_V3PharoahFactory *V3PharoahFactorySession) PoolFeeProtocol(pool common.Address) (*big.Int, error) {
	return _V3PharoahFactory.Contract.PoolFeeProtocol(&_V3PharoahFactory.CallOpts, pool)
}

// PoolFeeProtocol is a free data retrieval call binding the contract method 0xebb0d9f7.
//
// Solidity: function poolFeeProtocol(address pool) view returns(uint24 __poolFeeProtocol)
func (_V3PharoahFactory *V3PharoahFactoryCallerSession) PoolFeeProtocol(pool common.Address) (*big.Int, error) {
	return _V3PharoahFactory.Contract.PoolFeeProtocol(&_V3PharoahFactory.CallOpts, pool)
}

// RamsesV3PoolDeployer is a free data retrieval call binding the contract method 0xbf49a292.
//
// Solidity: function ramsesV3PoolDeployer() view returns(address)
func (_V3PharoahFactory *V3PharoahFactoryCaller) RamsesV3PoolDeployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3PharoahFactory.contract.Call(opts, &out, "ramsesV3PoolDeployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RamsesV3PoolDeployer is a free data retrieval call binding the contract method 0xbf49a292.
//
// Solidity: function ramsesV3PoolDeployer() view returns(address)
func (_V3PharoahFactory *V3PharoahFactorySession) RamsesV3PoolDeployer() (common.Address, error) {
	return _V3PharoahFactory.Contract.RamsesV3PoolDeployer(&_V3PharoahFactory.CallOpts)
}

// RamsesV3PoolDeployer is a free data retrieval call binding the contract method 0xbf49a292.
//
// Solidity: function ramsesV3PoolDeployer() view returns(address)
func (_V3PharoahFactory *V3PharoahFactoryCallerSession) RamsesV3PoolDeployer() (common.Address, error) {
	return _V3PharoahFactory.Contract.RamsesV3PoolDeployer(&_V3PharoahFactory.CallOpts)
}

// TickSpacingInitialFee is a free data retrieval call binding the contract method 0xcf3a52a6.
//
// Solidity: function tickSpacingInitialFee(int24 tickSpacing) view returns(uint24 initialFee)
func (_V3PharoahFactory *V3PharoahFactoryCaller) TickSpacingInitialFee(opts *bind.CallOpts, tickSpacing *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _V3PharoahFactory.contract.Call(opts, &out, "tickSpacingInitialFee", tickSpacing)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSpacingInitialFee is a free data retrieval call binding the contract method 0xcf3a52a6.
//
// Solidity: function tickSpacingInitialFee(int24 tickSpacing) view returns(uint24 initialFee)
func (_V3PharoahFactory *V3PharoahFactorySession) TickSpacingInitialFee(tickSpacing *big.Int) (*big.Int, error) {
	return _V3PharoahFactory.Contract.TickSpacingInitialFee(&_V3PharoahFactory.CallOpts, tickSpacing)
}

// TickSpacingInitialFee is a free data retrieval call binding the contract method 0xcf3a52a6.
//
// Solidity: function tickSpacingInitialFee(int24 tickSpacing) view returns(uint24 initialFee)
func (_V3PharoahFactory *V3PharoahFactoryCallerSession) TickSpacingInitialFee(tickSpacing *big.Int) (*big.Int, error) {
	return _V3PharoahFactory.Contract.TickSpacingInitialFee(&_V3PharoahFactory.CallOpts, tickSpacing)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_V3PharoahFactory *V3PharoahFactoryCaller) Voter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3PharoahFactory.contract.Call(opts, &out, "voter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_V3PharoahFactory *V3PharoahFactorySession) Voter() (common.Address, error) {
	return _V3PharoahFactory.Contract.Voter(&_V3PharoahFactory.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_V3PharoahFactory *V3PharoahFactoryCallerSession) Voter() (common.Address, error) {
	return _V3PharoahFactory.Contract.Voter(&_V3PharoahFactory.CallOpts)
}

// CreatePool is a paid mutator transaction binding the contract method 0x232aa5ac.
//
// Solidity: function createPool(address tokenA, address tokenB, int24 tickSpacing, uint160 sqrtPriceX96) returns(address pool)
func (_V3PharoahFactory *V3PharoahFactoryTransactor) CreatePool(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, tickSpacing *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _V3PharoahFactory.contract.Transact(opts, "createPool", tokenA, tokenB, tickSpacing, sqrtPriceX96)
}

// CreatePool is a paid mutator transaction binding the contract method 0x232aa5ac.
//
// Solidity: function createPool(address tokenA, address tokenB, int24 tickSpacing, uint160 sqrtPriceX96) returns(address pool)
func (_V3PharoahFactory *V3PharoahFactorySession) CreatePool(tokenA common.Address, tokenB common.Address, tickSpacing *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _V3PharoahFactory.Contract.CreatePool(&_V3PharoahFactory.TransactOpts, tokenA, tokenB, tickSpacing, sqrtPriceX96)
}

// CreatePool is a paid mutator transaction binding the contract method 0x232aa5ac.
//
// Solidity: function createPool(address tokenA, address tokenB, int24 tickSpacing, uint160 sqrtPriceX96) returns(address pool)
func (_V3PharoahFactory *V3PharoahFactoryTransactorSession) CreatePool(tokenA common.Address, tokenB common.Address, tickSpacing *big.Int, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _V3PharoahFactory.Contract.CreatePool(&_V3PharoahFactory.TransactOpts, tokenA, tokenB, tickSpacing, sqrtPriceX96)
}

// EnableTickSpacing is a paid mutator transaction binding the contract method 0xeee0fdb4.
//
// Solidity: function enableTickSpacing(int24 tickSpacing, uint24 initialFee) returns()
func (_V3PharoahFactory *V3PharoahFactoryTransactor) EnableTickSpacing(opts *bind.TransactOpts, tickSpacing *big.Int, initialFee *big.Int) (*types.Transaction, error) {
	return _V3PharoahFactory.contract.Transact(opts, "enableTickSpacing", tickSpacing, initialFee)
}

// EnableTickSpacing is a paid mutator transaction binding the contract method 0xeee0fdb4.
//
// Solidity: function enableTickSpacing(int24 tickSpacing, uint24 initialFee) returns()
func (_V3PharoahFactory *V3PharoahFactorySession) EnableTickSpacing(tickSpacing *big.Int, initialFee *big.Int) (*types.Transaction, error) {
	return _V3PharoahFactory.Contract.EnableTickSpacing(&_V3PharoahFactory.TransactOpts, tickSpacing, initialFee)
}

// EnableTickSpacing is a paid mutator transaction binding the contract method 0xeee0fdb4.
//
// Solidity: function enableTickSpacing(int24 tickSpacing, uint24 initialFee) returns()
func (_V3PharoahFactory *V3PharoahFactoryTransactorSession) EnableTickSpacing(tickSpacing *big.Int, initialFee *big.Int) (*types.Transaction, error) {
	return _V3PharoahFactory.Contract.EnableTickSpacing(&_V3PharoahFactory.TransactOpts, tickSpacing, initialFee)
}

// GaugeFeeSplitEnable is a paid mutator transaction binding the contract method 0x3cb08b53.
//
// Solidity: function gaugeFeeSplitEnable(address pool) returns()
func (_V3PharoahFactory *V3PharoahFactoryTransactor) GaugeFeeSplitEnable(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _V3PharoahFactory.contract.Transact(opts, "gaugeFeeSplitEnable", pool)
}

// GaugeFeeSplitEnable is a paid mutator transaction binding the contract method 0x3cb08b53.
//
// Solidity: function gaugeFeeSplitEnable(address pool) returns()
func (_V3PharoahFactory *V3PharoahFactorySession) GaugeFeeSplitEnable(pool common.Address) (*types.Transaction, error) {
	return _V3PharoahFactory.Contract.GaugeFeeSplitEnable(&_V3PharoahFactory.TransactOpts, pool)
}

// GaugeFeeSplitEnable is a paid mutator transaction binding the contract method 0x3cb08b53.
//
// Solidity: function gaugeFeeSplitEnable(address pool) returns()
func (_V3PharoahFactory *V3PharoahFactoryTransactorSession) GaugeFeeSplitEnable(pool common.Address) (*types.Transaction, error) {
	return _V3PharoahFactory.Contract.GaugeFeeSplitEnable(&_V3PharoahFactory.TransactOpts, pool)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _ramsesV3PoolDeployer) returns()
func (_V3PharoahFactory *V3PharoahFactoryTransactor) Initialize(opts *bind.TransactOpts, _ramsesV3PoolDeployer common.Address) (*types.Transaction, error) {
	return _V3PharoahFactory.contract.Transact(opts, "initialize", _ramsesV3PoolDeployer)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _ramsesV3PoolDeployer) returns()
func (_V3PharoahFactory *V3PharoahFactorySession) Initialize(_ramsesV3PoolDeployer common.Address) (*types.Transaction, error) {
	return _V3PharoahFactory.Contract.Initialize(&_V3PharoahFactory.TransactOpts, _ramsesV3PoolDeployer)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _ramsesV3PoolDeployer) returns()
func (_V3PharoahFactory *V3PharoahFactoryTransactorSession) Initialize(_ramsesV3PoolDeployer common.Address) (*types.Transaction, error) {
	return _V3PharoahFactory.Contract.Initialize(&_V3PharoahFactory.TransactOpts, _ramsesV3PoolDeployer)
}

// SetFee is a paid mutator transaction binding the contract method 0xba364c3d.
//
// Solidity: function setFee(address _pool, uint24 _fee) returns()
func (_V3PharoahFactory *V3PharoahFactoryTransactor) SetFee(opts *bind.TransactOpts, _pool common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _V3PharoahFactory.contract.Transact(opts, "setFee", _pool, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0xba364c3d.
//
// Solidity: function setFee(address _pool, uint24 _fee) returns()
func (_V3PharoahFactory *V3PharoahFactorySession) SetFee(_pool common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _V3PharoahFactory.Contract.SetFee(&_V3PharoahFactory.TransactOpts, _pool, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0xba364c3d.
//
// Solidity: function setFee(address _pool, uint24 _fee) returns()
func (_V3PharoahFactory *V3PharoahFactoryTransactorSession) SetFee(_pool common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _V3PharoahFactory.Contract.SetFee(&_V3PharoahFactory.TransactOpts, _pool, _fee)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_V3PharoahFactory *V3PharoahFactoryTransactor) SetFeeCollector(opts *bind.TransactOpts, _feeCollector common.Address) (*types.Transaction, error) {
	return _V3PharoahFactory.contract.Transact(opts, "setFeeCollector", _feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_V3PharoahFactory *V3PharoahFactorySession) SetFeeCollector(_feeCollector common.Address) (*types.Transaction, error) {
	return _V3PharoahFactory.Contract.SetFeeCollector(&_V3PharoahFactory.TransactOpts, _feeCollector)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address _feeCollector) returns()
func (_V3PharoahFactory *V3PharoahFactoryTransactorSession) SetFeeCollector(_feeCollector common.Address) (*types.Transaction, error) {
	return _V3PharoahFactory.Contract.SetFeeCollector(&_V3PharoahFactory.TransactOpts, _feeCollector)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x7fe35510.
//
// Solidity: function setFeeProtocol(uint24 _feeProtocol) returns()
func (_V3PharoahFactory *V3PharoahFactoryTransactor) SetFeeProtocol(opts *bind.TransactOpts, _feeProtocol *big.Int) (*types.Transaction, error) {
	return _V3PharoahFactory.contract.Transact(opts, "setFeeProtocol", _feeProtocol)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x7fe35510.
//
// Solidity: function setFeeProtocol(uint24 _feeProtocol) returns()
func (_V3PharoahFactory *V3PharoahFactorySession) SetFeeProtocol(_feeProtocol *big.Int) (*types.Transaction, error) {
	return _V3PharoahFactory.Contract.SetFeeProtocol(&_V3PharoahFactory.TransactOpts, _feeProtocol)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x7fe35510.
//
// Solidity: function setFeeProtocol(uint24 _feeProtocol) returns()
func (_V3PharoahFactory *V3PharoahFactoryTransactorSession) SetFeeProtocol(_feeProtocol *big.Int) (*types.Transaction, error) {
	return _V3PharoahFactory.Contract.SetFeeProtocol(&_V3PharoahFactory.TransactOpts, _feeProtocol)
}

// SetPoolFeeProtocol is a paid mutator transaction binding the contract method 0x7ab4974d.
//
// Solidity: function setPoolFeeProtocol(address pool, uint24 _feeProtocol) returns()
func (_V3PharoahFactory *V3PharoahFactoryTransactor) SetPoolFeeProtocol(opts *bind.TransactOpts, pool common.Address, _feeProtocol *big.Int) (*types.Transaction, error) {
	return _V3PharoahFactory.contract.Transact(opts, "setPoolFeeProtocol", pool, _feeProtocol)
}

// SetPoolFeeProtocol is a paid mutator transaction binding the contract method 0x7ab4974d.
//
// Solidity: function setPoolFeeProtocol(address pool, uint24 _feeProtocol) returns()
func (_V3PharoahFactory *V3PharoahFactorySession) SetPoolFeeProtocol(pool common.Address, _feeProtocol *big.Int) (*types.Transaction, error) {
	return _V3PharoahFactory.Contract.SetPoolFeeProtocol(&_V3PharoahFactory.TransactOpts, pool, _feeProtocol)
}

// SetPoolFeeProtocol is a paid mutator transaction binding the contract method 0x7ab4974d.
//
// Solidity: function setPoolFeeProtocol(address pool, uint24 _feeProtocol) returns()
func (_V3PharoahFactory *V3PharoahFactoryTransactorSession) SetPoolFeeProtocol(pool common.Address, _feeProtocol *big.Int) (*types.Transaction, error) {
	return _V3PharoahFactory.Contract.SetPoolFeeProtocol(&_V3PharoahFactory.TransactOpts, pool, _feeProtocol)
}

// SetVoter is a paid mutator transaction binding the contract method 0x4bc2a657.
//
// Solidity: function setVoter(address _voter) returns()
func (_V3PharoahFactory *V3PharoahFactoryTransactor) SetVoter(opts *bind.TransactOpts, _voter common.Address) (*types.Transaction, error) {
	return _V3PharoahFactory.contract.Transact(opts, "setVoter", _voter)
}

// SetVoter is a paid mutator transaction binding the contract method 0x4bc2a657.
//
// Solidity: function setVoter(address _voter) returns()
func (_V3PharoahFactory *V3PharoahFactorySession) SetVoter(_voter common.Address) (*types.Transaction, error) {
	return _V3PharoahFactory.Contract.SetVoter(&_V3PharoahFactory.TransactOpts, _voter)
}

// SetVoter is a paid mutator transaction binding the contract method 0x4bc2a657.
//
// Solidity: function setVoter(address _voter) returns()
func (_V3PharoahFactory *V3PharoahFactoryTransactorSession) SetVoter(_voter common.Address) (*types.Transaction, error) {
	return _V3PharoahFactory.Contract.SetVoter(&_V3PharoahFactory.TransactOpts, _voter)
}

// V3PharoahFactoryFeeAdjustmentIterator is returned from FilterFeeAdjustment and is used to iterate over the raw logs and unpacked data for FeeAdjustment events raised by the V3PharoahFactory contract.
type V3PharoahFactoryFeeAdjustmentIterator struct {
	Event *V3PharoahFactoryFeeAdjustment // Event containing the contract specifics and raw log

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
func (it *V3PharoahFactoryFeeAdjustmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3PharoahFactoryFeeAdjustment)
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
		it.Event = new(V3PharoahFactoryFeeAdjustment)
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
func (it *V3PharoahFactoryFeeAdjustmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3PharoahFactoryFeeAdjustmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3PharoahFactoryFeeAdjustment represents a FeeAdjustment event raised by the V3PharoahFactory contract.
type V3PharoahFactoryFeeAdjustment struct {
	Pool   common.Address
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeAdjustment is a free log retrieval operation binding the contract event 0xe4accbaee82fb833ac207d4c4454c5a04e85f5e1e9a20a9e2c98e54e8706ff2b.
//
// Solidity: event FeeAdjustment(address pool, uint24 newFee)
func (_V3PharoahFactory *V3PharoahFactoryFilterer) FilterFeeAdjustment(opts *bind.FilterOpts) (*V3PharoahFactoryFeeAdjustmentIterator, error) {

	logs, sub, err := _V3PharoahFactory.contract.FilterLogs(opts, "FeeAdjustment")
	if err != nil {
		return nil, err
	}
	return &V3PharoahFactoryFeeAdjustmentIterator{contract: _V3PharoahFactory.contract, event: "FeeAdjustment", logs: logs, sub: sub}, nil
}

// WatchFeeAdjustment is a free log subscription operation binding the contract event 0xe4accbaee82fb833ac207d4c4454c5a04e85f5e1e9a20a9e2c98e54e8706ff2b.
//
// Solidity: event FeeAdjustment(address pool, uint24 newFee)
func (_V3PharoahFactory *V3PharoahFactoryFilterer) WatchFeeAdjustment(opts *bind.WatchOpts, sink chan<- *V3PharoahFactoryFeeAdjustment) (event.Subscription, error) {

	logs, sub, err := _V3PharoahFactory.contract.WatchLogs(opts, "FeeAdjustment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3PharoahFactoryFeeAdjustment)
				if err := _V3PharoahFactory.contract.UnpackLog(event, "FeeAdjustment", log); err != nil {
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

// ParseFeeAdjustment is a log parse operation binding the contract event 0xe4accbaee82fb833ac207d4c4454c5a04e85f5e1e9a20a9e2c98e54e8706ff2b.
//
// Solidity: event FeeAdjustment(address pool, uint24 newFee)
func (_V3PharoahFactory *V3PharoahFactoryFilterer) ParseFeeAdjustment(log types.Log) (*V3PharoahFactoryFeeAdjustment, error) {
	event := new(V3PharoahFactoryFeeAdjustment)
	if err := _V3PharoahFactory.contract.UnpackLog(event, "FeeAdjustment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3PharoahFactoryFeeCollectorChangedIterator is returned from FilterFeeCollectorChanged and is used to iterate over the raw logs and unpacked data for FeeCollectorChanged events raised by the V3PharoahFactory contract.
type V3PharoahFactoryFeeCollectorChangedIterator struct {
	Event *V3PharoahFactoryFeeCollectorChanged // Event containing the contract specifics and raw log

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
func (it *V3PharoahFactoryFeeCollectorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3PharoahFactoryFeeCollectorChanged)
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
		it.Event = new(V3PharoahFactoryFeeCollectorChanged)
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
func (it *V3PharoahFactoryFeeCollectorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3PharoahFactoryFeeCollectorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3PharoahFactoryFeeCollectorChanged represents a FeeCollectorChanged event raised by the V3PharoahFactory contract.
type V3PharoahFactoryFeeCollectorChanged struct {
	OldFeeCollector common.Address
	NewFeeCollector common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFeeCollectorChanged is a free log retrieval operation binding the contract event 0x649c5e3d0ed183894196148e193af316452b0037e77d2ff0fef23b7dc722bed0.
//
// Solidity: event FeeCollectorChanged(address indexed oldFeeCollector, address indexed newFeeCollector)
func (_V3PharoahFactory *V3PharoahFactoryFilterer) FilterFeeCollectorChanged(opts *bind.FilterOpts, oldFeeCollector []common.Address, newFeeCollector []common.Address) (*V3PharoahFactoryFeeCollectorChangedIterator, error) {

	var oldFeeCollectorRule []interface{}
	for _, oldFeeCollectorItem := range oldFeeCollector {
		oldFeeCollectorRule = append(oldFeeCollectorRule, oldFeeCollectorItem)
	}
	var newFeeCollectorRule []interface{}
	for _, newFeeCollectorItem := range newFeeCollector {
		newFeeCollectorRule = append(newFeeCollectorRule, newFeeCollectorItem)
	}

	logs, sub, err := _V3PharoahFactory.contract.FilterLogs(opts, "FeeCollectorChanged", oldFeeCollectorRule, newFeeCollectorRule)
	if err != nil {
		return nil, err
	}
	return &V3PharoahFactoryFeeCollectorChangedIterator{contract: _V3PharoahFactory.contract, event: "FeeCollectorChanged", logs: logs, sub: sub}, nil
}

// WatchFeeCollectorChanged is a free log subscription operation binding the contract event 0x649c5e3d0ed183894196148e193af316452b0037e77d2ff0fef23b7dc722bed0.
//
// Solidity: event FeeCollectorChanged(address indexed oldFeeCollector, address indexed newFeeCollector)
func (_V3PharoahFactory *V3PharoahFactoryFilterer) WatchFeeCollectorChanged(opts *bind.WatchOpts, sink chan<- *V3PharoahFactoryFeeCollectorChanged, oldFeeCollector []common.Address, newFeeCollector []common.Address) (event.Subscription, error) {

	var oldFeeCollectorRule []interface{}
	for _, oldFeeCollectorItem := range oldFeeCollector {
		oldFeeCollectorRule = append(oldFeeCollectorRule, oldFeeCollectorItem)
	}
	var newFeeCollectorRule []interface{}
	for _, newFeeCollectorItem := range newFeeCollector {
		newFeeCollectorRule = append(newFeeCollectorRule, newFeeCollectorItem)
	}

	logs, sub, err := _V3PharoahFactory.contract.WatchLogs(opts, "FeeCollectorChanged", oldFeeCollectorRule, newFeeCollectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3PharoahFactoryFeeCollectorChanged)
				if err := _V3PharoahFactory.contract.UnpackLog(event, "FeeCollectorChanged", log); err != nil {
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

// ParseFeeCollectorChanged is a log parse operation binding the contract event 0x649c5e3d0ed183894196148e193af316452b0037e77d2ff0fef23b7dc722bed0.
//
// Solidity: event FeeCollectorChanged(address indexed oldFeeCollector, address indexed newFeeCollector)
func (_V3PharoahFactory *V3PharoahFactoryFilterer) ParseFeeCollectorChanged(log types.Log) (*V3PharoahFactoryFeeCollectorChanged, error) {
	event := new(V3PharoahFactoryFeeCollectorChanged)
	if err := _V3PharoahFactory.contract.UnpackLog(event, "FeeCollectorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3PharoahFactoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the V3PharoahFactory contract.
type V3PharoahFactoryPoolCreatedIterator struct {
	Event *V3PharoahFactoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *V3PharoahFactoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3PharoahFactoryPoolCreated)
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
		it.Event = new(V3PharoahFactoryPoolCreated)
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
func (it *V3PharoahFactoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3PharoahFactoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3PharoahFactoryPoolCreated represents a PoolCreated event raised by the V3PharoahFactory contract.
type V3PharoahFactoryPoolCreated struct {
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Pool        common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_V3PharoahFactory *V3PharoahFactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address, fee []*big.Int) (*V3PharoahFactoryPoolCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _V3PharoahFactory.contract.FilterLogs(opts, "PoolCreated", token0Rule, token1Rule, feeRule)
	if err != nil {
		return nil, err
	}
	return &V3PharoahFactoryPoolCreatedIterator{contract: _V3PharoahFactory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_V3PharoahFactory *V3PharoahFactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *V3PharoahFactoryPoolCreated, token0 []common.Address, token1 []common.Address, fee []*big.Int) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _V3PharoahFactory.contract.WatchLogs(opts, "PoolCreated", token0Rule, token1Rule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3PharoahFactoryPoolCreated)
				if err := _V3PharoahFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_V3PharoahFactory *V3PharoahFactoryFilterer) ParsePoolCreated(log types.Log) (*V3PharoahFactoryPoolCreated, error) {
	event := new(V3PharoahFactoryPoolCreated)
	if err := _V3PharoahFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3PharoahFactorySetFeeProtocolIterator is returned from FilterSetFeeProtocol and is used to iterate over the raw logs and unpacked data for SetFeeProtocol events raised by the V3PharoahFactory contract.
type V3PharoahFactorySetFeeProtocolIterator struct {
	Event *V3PharoahFactorySetFeeProtocol // Event containing the contract specifics and raw log

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
func (it *V3PharoahFactorySetFeeProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3PharoahFactorySetFeeProtocol)
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
		it.Event = new(V3PharoahFactorySetFeeProtocol)
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
func (it *V3PharoahFactorySetFeeProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3PharoahFactorySetFeeProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3PharoahFactorySetFeeProtocol represents a SetFeeProtocol event raised by the V3PharoahFactory contract.
type V3PharoahFactorySetFeeProtocol struct {
	FeeProtocolOld *big.Int
	FeeProtocolNew *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetFeeProtocol is a free log retrieval operation binding the contract event 0x67a069e4d951485f3e494a1edfa67d7334e991e8514ba748fd1636270acd1c97.
//
// Solidity: event SetFeeProtocol(uint24 feeProtocolOld, uint24 feeProtocolNew)
func (_V3PharoahFactory *V3PharoahFactoryFilterer) FilterSetFeeProtocol(opts *bind.FilterOpts) (*V3PharoahFactorySetFeeProtocolIterator, error) {

	logs, sub, err := _V3PharoahFactory.contract.FilterLogs(opts, "SetFeeProtocol")
	if err != nil {
		return nil, err
	}
	return &V3PharoahFactorySetFeeProtocolIterator{contract: _V3PharoahFactory.contract, event: "SetFeeProtocol", logs: logs, sub: sub}, nil
}

// WatchSetFeeProtocol is a free log subscription operation binding the contract event 0x67a069e4d951485f3e494a1edfa67d7334e991e8514ba748fd1636270acd1c97.
//
// Solidity: event SetFeeProtocol(uint24 feeProtocolOld, uint24 feeProtocolNew)
func (_V3PharoahFactory *V3PharoahFactoryFilterer) WatchSetFeeProtocol(opts *bind.WatchOpts, sink chan<- *V3PharoahFactorySetFeeProtocol) (event.Subscription, error) {

	logs, sub, err := _V3PharoahFactory.contract.WatchLogs(opts, "SetFeeProtocol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3PharoahFactorySetFeeProtocol)
				if err := _V3PharoahFactory.contract.UnpackLog(event, "SetFeeProtocol", log); err != nil {
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

// ParseSetFeeProtocol is a log parse operation binding the contract event 0x67a069e4d951485f3e494a1edfa67d7334e991e8514ba748fd1636270acd1c97.
//
// Solidity: event SetFeeProtocol(uint24 feeProtocolOld, uint24 feeProtocolNew)
func (_V3PharoahFactory *V3PharoahFactoryFilterer) ParseSetFeeProtocol(log types.Log) (*V3PharoahFactorySetFeeProtocol, error) {
	event := new(V3PharoahFactorySetFeeProtocol)
	if err := _V3PharoahFactory.contract.UnpackLog(event, "SetFeeProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3PharoahFactorySetPoolFeeProtocolIterator is returned from FilterSetPoolFeeProtocol and is used to iterate over the raw logs and unpacked data for SetPoolFeeProtocol events raised by the V3PharoahFactory contract.
type V3PharoahFactorySetPoolFeeProtocolIterator struct {
	Event *V3PharoahFactorySetPoolFeeProtocol // Event containing the contract specifics and raw log

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
func (it *V3PharoahFactorySetPoolFeeProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3PharoahFactorySetPoolFeeProtocol)
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
		it.Event = new(V3PharoahFactorySetPoolFeeProtocol)
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
func (it *V3PharoahFactorySetPoolFeeProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3PharoahFactorySetPoolFeeProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3PharoahFactorySetPoolFeeProtocol represents a SetPoolFeeProtocol event raised by the V3PharoahFactory contract.
type V3PharoahFactorySetPoolFeeProtocol struct {
	Pool           common.Address
	FeeProtocolOld *big.Int
	FeeProtocolNew *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetPoolFeeProtocol is a free log retrieval operation binding the contract event 0x1fb49ee35e38c4a757469d4a1c37187e7b3821f994a5556fde452ba2607ee235.
//
// Solidity: event SetPoolFeeProtocol(address pool, uint24 feeProtocolOld, uint24 feeProtocolNew)
func (_V3PharoahFactory *V3PharoahFactoryFilterer) FilterSetPoolFeeProtocol(opts *bind.FilterOpts) (*V3PharoahFactorySetPoolFeeProtocolIterator, error) {

	logs, sub, err := _V3PharoahFactory.contract.FilterLogs(opts, "SetPoolFeeProtocol")
	if err != nil {
		return nil, err
	}
	return &V3PharoahFactorySetPoolFeeProtocolIterator{contract: _V3PharoahFactory.contract, event: "SetPoolFeeProtocol", logs: logs, sub: sub}, nil
}

// WatchSetPoolFeeProtocol is a free log subscription operation binding the contract event 0x1fb49ee35e38c4a757469d4a1c37187e7b3821f994a5556fde452ba2607ee235.
//
// Solidity: event SetPoolFeeProtocol(address pool, uint24 feeProtocolOld, uint24 feeProtocolNew)
func (_V3PharoahFactory *V3PharoahFactoryFilterer) WatchSetPoolFeeProtocol(opts *bind.WatchOpts, sink chan<- *V3PharoahFactorySetPoolFeeProtocol) (event.Subscription, error) {

	logs, sub, err := _V3PharoahFactory.contract.WatchLogs(opts, "SetPoolFeeProtocol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3PharoahFactorySetPoolFeeProtocol)
				if err := _V3PharoahFactory.contract.UnpackLog(event, "SetPoolFeeProtocol", log); err != nil {
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

// ParseSetPoolFeeProtocol is a log parse operation binding the contract event 0x1fb49ee35e38c4a757469d4a1c37187e7b3821f994a5556fde452ba2607ee235.
//
// Solidity: event SetPoolFeeProtocol(address pool, uint24 feeProtocolOld, uint24 feeProtocolNew)
func (_V3PharoahFactory *V3PharoahFactoryFilterer) ParseSetPoolFeeProtocol(log types.Log) (*V3PharoahFactorySetPoolFeeProtocol, error) {
	event := new(V3PharoahFactorySetPoolFeeProtocol)
	if err := _V3PharoahFactory.contract.UnpackLog(event, "SetPoolFeeProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3PharoahFactoryTickSpacingEnabledIterator is returned from FilterTickSpacingEnabled and is used to iterate over the raw logs and unpacked data for TickSpacingEnabled events raised by the V3PharoahFactory contract.
type V3PharoahFactoryTickSpacingEnabledIterator struct {
	Event *V3PharoahFactoryTickSpacingEnabled // Event containing the contract specifics and raw log

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
func (it *V3PharoahFactoryTickSpacingEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3PharoahFactoryTickSpacingEnabled)
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
		it.Event = new(V3PharoahFactoryTickSpacingEnabled)
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
func (it *V3PharoahFactoryTickSpacingEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3PharoahFactoryTickSpacingEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3PharoahFactoryTickSpacingEnabled represents a TickSpacingEnabled event raised by the V3PharoahFactory contract.
type V3PharoahFactoryTickSpacingEnabled struct {
	TickSpacing *big.Int
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTickSpacingEnabled is a free log retrieval operation binding the contract event 0xebafae466a4a780a1d87f5fab2f52fad33be9151a7f69d099e8934c8de85b747.
//
// Solidity: event TickSpacingEnabled(int24 indexed tickSpacing, uint24 indexed fee)
func (_V3PharoahFactory *V3PharoahFactoryFilterer) FilterTickSpacingEnabled(opts *bind.FilterOpts, tickSpacing []*big.Int, fee []*big.Int) (*V3PharoahFactoryTickSpacingEnabledIterator, error) {

	var tickSpacingRule []interface{}
	for _, tickSpacingItem := range tickSpacing {
		tickSpacingRule = append(tickSpacingRule, tickSpacingItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _V3PharoahFactory.contract.FilterLogs(opts, "TickSpacingEnabled", tickSpacingRule, feeRule)
	if err != nil {
		return nil, err
	}
	return &V3PharoahFactoryTickSpacingEnabledIterator{contract: _V3PharoahFactory.contract, event: "TickSpacingEnabled", logs: logs, sub: sub}, nil
}

// WatchTickSpacingEnabled is a free log subscription operation binding the contract event 0xebafae466a4a780a1d87f5fab2f52fad33be9151a7f69d099e8934c8de85b747.
//
// Solidity: event TickSpacingEnabled(int24 indexed tickSpacing, uint24 indexed fee)
func (_V3PharoahFactory *V3PharoahFactoryFilterer) WatchTickSpacingEnabled(opts *bind.WatchOpts, sink chan<- *V3PharoahFactoryTickSpacingEnabled, tickSpacing []*big.Int, fee []*big.Int) (event.Subscription, error) {

	var tickSpacingRule []interface{}
	for _, tickSpacingItem := range tickSpacing {
		tickSpacingRule = append(tickSpacingRule, tickSpacingItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _V3PharoahFactory.contract.WatchLogs(opts, "TickSpacingEnabled", tickSpacingRule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3PharoahFactoryTickSpacingEnabled)
				if err := _V3PharoahFactory.contract.UnpackLog(event, "TickSpacingEnabled", log); err != nil {
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
func (_V3PharoahFactory *V3PharoahFactoryFilterer) ParseTickSpacingEnabled(log types.Log) (*V3PharoahFactoryTickSpacingEnabled, error) {
	event := new(V3PharoahFactoryTickSpacingEnabled)
	if err := _V3PharoahFactory.contract.UnpackLog(event, "TickSpacingEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
