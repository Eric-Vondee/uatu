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

// V3QuickSwapFactoryMetaData contains all meta data concerning the V3QuickSwapFactory contract.
var V3QuickSwapFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolDeployer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vaultAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFarmingAddress\",\"type\":\"address\"}],\"name\":\"FarmingAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"alpha1\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"alpha2\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"beta1\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"beta2\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"gamma1\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"gamma2\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"volumeBeta\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"volumeGamma\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"baseFee\",\"type\":\"uint16\"}],\"name\":\"FeeConfiguration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"Owner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"Pool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVaultAddress\",\"type\":\"address\"}],\"name\":\"VaultAddress\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"baseFeeConfiguration\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"alpha1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"alpha2\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"beta1\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"beta2\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"gamma1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"gamma2\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"volumeBeta\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"volumeGamma\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"baseFee\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"farmingAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolByPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolDeployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"alpha1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"alpha2\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"beta1\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"beta2\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"gamma1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"gamma2\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"volumeBeta\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"volumeGamma\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"baseFee\",\"type\":\"uint16\"}],\"name\":\"setBaseFeeConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_farmingAddress\",\"type\":\"address\"}],\"name\":\"setFarmingAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultAddress\",\"type\":\"address\"}],\"name\":\"setVaultAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// V3QuickSwapFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use V3QuickSwapFactoryMetaData.ABI instead.
var V3QuickSwapFactoryABI = V3QuickSwapFactoryMetaData.ABI

// V3QuickSwapFactory is an auto generated Go binding around an Ethereum contract.
type V3QuickSwapFactory struct {
	V3QuickSwapFactoryCaller     // Read-only binding to the contract
	V3QuickSwapFactoryTransactor // Write-only binding to the contract
	V3QuickSwapFactoryFilterer   // Log filterer for contract events
}

// V3QuickSwapFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type V3QuickSwapFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3QuickSwapFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V3QuickSwapFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3QuickSwapFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V3QuickSwapFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3QuickSwapFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V3QuickSwapFactorySession struct {
	Contract     *V3QuickSwapFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// V3QuickSwapFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V3QuickSwapFactoryCallerSession struct {
	Contract *V3QuickSwapFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// V3QuickSwapFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V3QuickSwapFactoryTransactorSession struct {
	Contract     *V3QuickSwapFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// V3QuickSwapFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type V3QuickSwapFactoryRaw struct {
	Contract *V3QuickSwapFactory // Generic contract binding to access the raw methods on
}

// V3QuickSwapFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V3QuickSwapFactoryCallerRaw struct {
	Contract *V3QuickSwapFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// V3QuickSwapFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V3QuickSwapFactoryTransactorRaw struct {
	Contract *V3QuickSwapFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV3QuickSwapFactory creates a new instance of V3QuickSwapFactory, bound to a specific deployed contract.
func NewV3QuickSwapFactory(address common.Address, backend bind.ContractBackend) (*V3QuickSwapFactory, error) {
	contract, err := bindV3QuickSwapFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V3QuickSwapFactory{V3QuickSwapFactoryCaller: V3QuickSwapFactoryCaller{contract: contract}, V3QuickSwapFactoryTransactor: V3QuickSwapFactoryTransactor{contract: contract}, V3QuickSwapFactoryFilterer: V3QuickSwapFactoryFilterer{contract: contract}}, nil
}

// NewV3QuickSwapFactoryCaller creates a new read-only instance of V3QuickSwapFactory, bound to a specific deployed contract.
func NewV3QuickSwapFactoryCaller(address common.Address, caller bind.ContractCaller) (*V3QuickSwapFactoryCaller, error) {
	contract, err := bindV3QuickSwapFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V3QuickSwapFactoryCaller{contract: contract}, nil
}

// NewV3QuickSwapFactoryTransactor creates a new write-only instance of V3QuickSwapFactory, bound to a specific deployed contract.
func NewV3QuickSwapFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*V3QuickSwapFactoryTransactor, error) {
	contract, err := bindV3QuickSwapFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V3QuickSwapFactoryTransactor{contract: contract}, nil
}

// NewV3QuickSwapFactoryFilterer creates a new log filterer instance of V3QuickSwapFactory, bound to a specific deployed contract.
func NewV3QuickSwapFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*V3QuickSwapFactoryFilterer, error) {
	contract, err := bindV3QuickSwapFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V3QuickSwapFactoryFilterer{contract: contract}, nil
}

// bindV3QuickSwapFactory binds a generic wrapper to an already deployed contract.
func bindV3QuickSwapFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := V3QuickSwapFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3QuickSwapFactory *V3QuickSwapFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3QuickSwapFactory.Contract.V3QuickSwapFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3QuickSwapFactory *V3QuickSwapFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3QuickSwapFactory.Contract.V3QuickSwapFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3QuickSwapFactory *V3QuickSwapFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3QuickSwapFactory.Contract.V3QuickSwapFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3QuickSwapFactory *V3QuickSwapFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3QuickSwapFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3QuickSwapFactory *V3QuickSwapFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3QuickSwapFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3QuickSwapFactory *V3QuickSwapFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3QuickSwapFactory.Contract.contract.Transact(opts, method, params...)
}

// BaseFeeConfiguration is a free data retrieval call binding the contract method 0x9832853a.
//
// Solidity: function baseFeeConfiguration() view returns(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_V3QuickSwapFactory *V3QuickSwapFactoryCaller) BaseFeeConfiguration(opts *bind.CallOpts) (struct {
	Alpha1      uint16
	Alpha2      uint16
	Beta1       uint32
	Beta2       uint32
	Gamma1      uint16
	Gamma2      uint16
	VolumeBeta  uint32
	VolumeGamma uint16
	BaseFee     uint16
}, error) {
	var out []interface{}
	err := _V3QuickSwapFactory.contract.Call(opts, &out, "baseFeeConfiguration")

	outstruct := new(struct {
		Alpha1      uint16
		Alpha2      uint16
		Beta1       uint32
		Beta2       uint32
		Gamma1      uint16
		Gamma2      uint16
		VolumeBeta  uint32
		VolumeGamma uint16
		BaseFee     uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Alpha1 = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.Alpha2 = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.Beta1 = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Beta2 = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.Gamma1 = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.Gamma2 = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.VolumeBeta = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.VolumeGamma = *abi.ConvertType(out[7], new(uint16)).(*uint16)
	outstruct.BaseFee = *abi.ConvertType(out[8], new(uint16)).(*uint16)

	return *outstruct, err

}

// BaseFeeConfiguration is a free data retrieval call binding the contract method 0x9832853a.
//
// Solidity: function baseFeeConfiguration() view returns(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_V3QuickSwapFactory *V3QuickSwapFactorySession) BaseFeeConfiguration() (struct {
	Alpha1      uint16
	Alpha2      uint16
	Beta1       uint32
	Beta2       uint32
	Gamma1      uint16
	Gamma2      uint16
	VolumeBeta  uint32
	VolumeGamma uint16
	BaseFee     uint16
}, error) {
	return _V3QuickSwapFactory.Contract.BaseFeeConfiguration(&_V3QuickSwapFactory.CallOpts)
}

// BaseFeeConfiguration is a free data retrieval call binding the contract method 0x9832853a.
//
// Solidity: function baseFeeConfiguration() view returns(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_V3QuickSwapFactory *V3QuickSwapFactoryCallerSession) BaseFeeConfiguration() (struct {
	Alpha1      uint16
	Alpha2      uint16
	Beta1       uint32
	Beta2       uint32
	Gamma1      uint16
	Gamma2      uint16
	VolumeBeta  uint32
	VolumeGamma uint16
	BaseFee     uint16
}, error) {
	return _V3QuickSwapFactory.Contract.BaseFeeConfiguration(&_V3QuickSwapFactory.CallOpts)
}

// FarmingAddress is a free data retrieval call binding the contract method 0x8a2ade58.
//
// Solidity: function farmingAddress() view returns(address)
func (_V3QuickSwapFactory *V3QuickSwapFactoryCaller) FarmingAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3QuickSwapFactory.contract.Call(opts, &out, "farmingAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FarmingAddress is a free data retrieval call binding the contract method 0x8a2ade58.
//
// Solidity: function farmingAddress() view returns(address)
func (_V3QuickSwapFactory *V3QuickSwapFactorySession) FarmingAddress() (common.Address, error) {
	return _V3QuickSwapFactory.Contract.FarmingAddress(&_V3QuickSwapFactory.CallOpts)
}

// FarmingAddress is a free data retrieval call binding the contract method 0x8a2ade58.
//
// Solidity: function farmingAddress() view returns(address)
func (_V3QuickSwapFactory *V3QuickSwapFactoryCallerSession) FarmingAddress() (common.Address, error) {
	return _V3QuickSwapFactory.Contract.FarmingAddress(&_V3QuickSwapFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_V3QuickSwapFactory *V3QuickSwapFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3QuickSwapFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_V3QuickSwapFactory *V3QuickSwapFactorySession) Owner() (common.Address, error) {
	return _V3QuickSwapFactory.Contract.Owner(&_V3QuickSwapFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_V3QuickSwapFactory *V3QuickSwapFactoryCallerSession) Owner() (common.Address, error) {
	return _V3QuickSwapFactory.Contract.Owner(&_V3QuickSwapFactory.CallOpts)
}

// PoolByPair is a free data retrieval call binding the contract method 0xd9a641e1.
//
// Solidity: function poolByPair(address , address ) view returns(address)
func (_V3QuickSwapFactory *V3QuickSwapFactoryCaller) PoolByPair(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _V3QuickSwapFactory.contract.Call(opts, &out, "poolByPair", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolByPair is a free data retrieval call binding the contract method 0xd9a641e1.
//
// Solidity: function poolByPair(address , address ) view returns(address)
func (_V3QuickSwapFactory *V3QuickSwapFactorySession) PoolByPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _V3QuickSwapFactory.Contract.PoolByPair(&_V3QuickSwapFactory.CallOpts, arg0, arg1)
}

// PoolByPair is a free data retrieval call binding the contract method 0xd9a641e1.
//
// Solidity: function poolByPair(address , address ) view returns(address)
func (_V3QuickSwapFactory *V3QuickSwapFactoryCallerSession) PoolByPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _V3QuickSwapFactory.Contract.PoolByPair(&_V3QuickSwapFactory.CallOpts, arg0, arg1)
}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_V3QuickSwapFactory *V3QuickSwapFactoryCaller) PoolDeployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3QuickSwapFactory.contract.Call(opts, &out, "poolDeployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_V3QuickSwapFactory *V3QuickSwapFactorySession) PoolDeployer() (common.Address, error) {
	return _V3QuickSwapFactory.Contract.PoolDeployer(&_V3QuickSwapFactory.CallOpts)
}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_V3QuickSwapFactory *V3QuickSwapFactoryCallerSession) PoolDeployer() (common.Address, error) {
	return _V3QuickSwapFactory.Contract.PoolDeployer(&_V3QuickSwapFactory.CallOpts)
}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_V3QuickSwapFactory *V3QuickSwapFactoryCaller) VaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V3QuickSwapFactory.contract.Call(opts, &out, "vaultAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_V3QuickSwapFactory *V3QuickSwapFactorySession) VaultAddress() (common.Address, error) {
	return _V3QuickSwapFactory.Contract.VaultAddress(&_V3QuickSwapFactory.CallOpts)
}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_V3QuickSwapFactory *V3QuickSwapFactoryCallerSession) VaultAddress() (common.Address, error) {
	return _V3QuickSwapFactory.Contract.VaultAddress(&_V3QuickSwapFactory.CallOpts)
}

// CreatePool is a paid mutator transaction binding the contract method 0xe3433615.
//
// Solidity: function createPool(address tokenA, address tokenB) returns(address pool)
func (_V3QuickSwapFactory *V3QuickSwapFactoryTransactor) CreatePool(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _V3QuickSwapFactory.contract.Transact(opts, "createPool", tokenA, tokenB)
}

// CreatePool is a paid mutator transaction binding the contract method 0xe3433615.
//
// Solidity: function createPool(address tokenA, address tokenB) returns(address pool)
func (_V3QuickSwapFactory *V3QuickSwapFactorySession) CreatePool(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _V3QuickSwapFactory.Contract.CreatePool(&_V3QuickSwapFactory.TransactOpts, tokenA, tokenB)
}

// CreatePool is a paid mutator transaction binding the contract method 0xe3433615.
//
// Solidity: function createPool(address tokenA, address tokenB) returns(address pool)
func (_V3QuickSwapFactory *V3QuickSwapFactoryTransactorSession) CreatePool(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _V3QuickSwapFactory.Contract.CreatePool(&_V3QuickSwapFactory.TransactOpts, tokenA, tokenB)
}

// SetBaseFeeConfiguration is a paid mutator transaction binding the contract method 0x5d6d7e93.
//
// Solidity: function setBaseFeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee) returns()
func (_V3QuickSwapFactory *V3QuickSwapFactoryTransactor) SetBaseFeeConfiguration(opts *bind.TransactOpts, alpha1 uint16, alpha2 uint16, beta1 uint32, beta2 uint32, gamma1 uint16, gamma2 uint16, volumeBeta uint32, volumeGamma uint16, baseFee uint16) (*types.Transaction, error) {
	return _V3QuickSwapFactory.contract.Transact(opts, "setBaseFeeConfiguration", alpha1, alpha2, beta1, beta2, gamma1, gamma2, volumeBeta, volumeGamma, baseFee)
}

// SetBaseFeeConfiguration is a paid mutator transaction binding the contract method 0x5d6d7e93.
//
// Solidity: function setBaseFeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee) returns()
func (_V3QuickSwapFactory *V3QuickSwapFactorySession) SetBaseFeeConfiguration(alpha1 uint16, alpha2 uint16, beta1 uint32, beta2 uint32, gamma1 uint16, gamma2 uint16, volumeBeta uint32, volumeGamma uint16, baseFee uint16) (*types.Transaction, error) {
	return _V3QuickSwapFactory.Contract.SetBaseFeeConfiguration(&_V3QuickSwapFactory.TransactOpts, alpha1, alpha2, beta1, beta2, gamma1, gamma2, volumeBeta, volumeGamma, baseFee)
}

// SetBaseFeeConfiguration is a paid mutator transaction binding the contract method 0x5d6d7e93.
//
// Solidity: function setBaseFeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee) returns()
func (_V3QuickSwapFactory *V3QuickSwapFactoryTransactorSession) SetBaseFeeConfiguration(alpha1 uint16, alpha2 uint16, beta1 uint32, beta2 uint32, gamma1 uint16, gamma2 uint16, volumeBeta uint32, volumeGamma uint16, baseFee uint16) (*types.Transaction, error) {
	return _V3QuickSwapFactory.Contract.SetBaseFeeConfiguration(&_V3QuickSwapFactory.TransactOpts, alpha1, alpha2, beta1, beta2, gamma1, gamma2, volumeBeta, volumeGamma, baseFee)
}

// SetFarmingAddress is a paid mutator transaction binding the contract method 0xb001f618.
//
// Solidity: function setFarmingAddress(address _farmingAddress) returns()
func (_V3QuickSwapFactory *V3QuickSwapFactoryTransactor) SetFarmingAddress(opts *bind.TransactOpts, _farmingAddress common.Address) (*types.Transaction, error) {
	return _V3QuickSwapFactory.contract.Transact(opts, "setFarmingAddress", _farmingAddress)
}

// SetFarmingAddress is a paid mutator transaction binding the contract method 0xb001f618.
//
// Solidity: function setFarmingAddress(address _farmingAddress) returns()
func (_V3QuickSwapFactory *V3QuickSwapFactorySession) SetFarmingAddress(_farmingAddress common.Address) (*types.Transaction, error) {
	return _V3QuickSwapFactory.Contract.SetFarmingAddress(&_V3QuickSwapFactory.TransactOpts, _farmingAddress)
}

// SetFarmingAddress is a paid mutator transaction binding the contract method 0xb001f618.
//
// Solidity: function setFarmingAddress(address _farmingAddress) returns()
func (_V3QuickSwapFactory *V3QuickSwapFactoryTransactorSession) SetFarmingAddress(_farmingAddress common.Address) (*types.Transaction, error) {
	return _V3QuickSwapFactory.Contract.SetFarmingAddress(&_V3QuickSwapFactory.TransactOpts, _farmingAddress)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_V3QuickSwapFactory *V3QuickSwapFactoryTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _V3QuickSwapFactory.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_V3QuickSwapFactory *V3QuickSwapFactorySession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _V3QuickSwapFactory.Contract.SetOwner(&_V3QuickSwapFactory.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_V3QuickSwapFactory *V3QuickSwapFactoryTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _V3QuickSwapFactory.Contract.SetOwner(&_V3QuickSwapFactory.TransactOpts, _owner)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x85535cc5.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (_V3QuickSwapFactory *V3QuickSwapFactoryTransactor) SetVaultAddress(opts *bind.TransactOpts, _vaultAddress common.Address) (*types.Transaction, error) {
	return _V3QuickSwapFactory.contract.Transact(opts, "setVaultAddress", _vaultAddress)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x85535cc5.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (_V3QuickSwapFactory *V3QuickSwapFactorySession) SetVaultAddress(_vaultAddress common.Address) (*types.Transaction, error) {
	return _V3QuickSwapFactory.Contract.SetVaultAddress(&_V3QuickSwapFactory.TransactOpts, _vaultAddress)
}

// SetVaultAddress is a paid mutator transaction binding the contract method 0x85535cc5.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (_V3QuickSwapFactory *V3QuickSwapFactoryTransactorSession) SetVaultAddress(_vaultAddress common.Address) (*types.Transaction, error) {
	return _V3QuickSwapFactory.Contract.SetVaultAddress(&_V3QuickSwapFactory.TransactOpts, _vaultAddress)
}

// V3QuickSwapFactoryFarmingAddressIterator is returned from FilterFarmingAddress and is used to iterate over the raw logs and unpacked data for FarmingAddress events raised by the V3QuickSwapFactory contract.
type V3QuickSwapFactoryFarmingAddressIterator struct {
	Event *V3QuickSwapFactoryFarmingAddress // Event containing the contract specifics and raw log

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
func (it *V3QuickSwapFactoryFarmingAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3QuickSwapFactoryFarmingAddress)
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
		it.Event = new(V3QuickSwapFactoryFarmingAddress)
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
func (it *V3QuickSwapFactoryFarmingAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3QuickSwapFactoryFarmingAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3QuickSwapFactoryFarmingAddress represents a FarmingAddress event raised by the V3QuickSwapFactory contract.
type V3QuickSwapFactoryFarmingAddress struct {
	NewFarmingAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFarmingAddress is a free log retrieval operation binding the contract event 0x56b9e8342f530796ceed0d5529abdcdeae6e4f2ac1dc456ceb73bbda898e0cd3.
//
// Solidity: event FarmingAddress(address indexed newFarmingAddress)
func (_V3QuickSwapFactory *V3QuickSwapFactoryFilterer) FilterFarmingAddress(opts *bind.FilterOpts, newFarmingAddress []common.Address) (*V3QuickSwapFactoryFarmingAddressIterator, error) {

	var newFarmingAddressRule []interface{}
	for _, newFarmingAddressItem := range newFarmingAddress {
		newFarmingAddressRule = append(newFarmingAddressRule, newFarmingAddressItem)
	}

	logs, sub, err := _V3QuickSwapFactory.contract.FilterLogs(opts, "FarmingAddress", newFarmingAddressRule)
	if err != nil {
		return nil, err
	}
	return &V3QuickSwapFactoryFarmingAddressIterator{contract: _V3QuickSwapFactory.contract, event: "FarmingAddress", logs: logs, sub: sub}, nil
}

// WatchFarmingAddress is a free log subscription operation binding the contract event 0x56b9e8342f530796ceed0d5529abdcdeae6e4f2ac1dc456ceb73bbda898e0cd3.
//
// Solidity: event FarmingAddress(address indexed newFarmingAddress)
func (_V3QuickSwapFactory *V3QuickSwapFactoryFilterer) WatchFarmingAddress(opts *bind.WatchOpts, sink chan<- *V3QuickSwapFactoryFarmingAddress, newFarmingAddress []common.Address) (event.Subscription, error) {

	var newFarmingAddressRule []interface{}
	for _, newFarmingAddressItem := range newFarmingAddress {
		newFarmingAddressRule = append(newFarmingAddressRule, newFarmingAddressItem)
	}

	logs, sub, err := _V3QuickSwapFactory.contract.WatchLogs(opts, "FarmingAddress", newFarmingAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3QuickSwapFactoryFarmingAddress)
				if err := _V3QuickSwapFactory.contract.UnpackLog(event, "FarmingAddress", log); err != nil {
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

// ParseFarmingAddress is a log parse operation binding the contract event 0x56b9e8342f530796ceed0d5529abdcdeae6e4f2ac1dc456ceb73bbda898e0cd3.
//
// Solidity: event FarmingAddress(address indexed newFarmingAddress)
func (_V3QuickSwapFactory *V3QuickSwapFactoryFilterer) ParseFarmingAddress(log types.Log) (*V3QuickSwapFactoryFarmingAddress, error) {
	event := new(V3QuickSwapFactoryFarmingAddress)
	if err := _V3QuickSwapFactory.contract.UnpackLog(event, "FarmingAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3QuickSwapFactoryFeeConfigurationIterator is returned from FilterFeeConfiguration and is used to iterate over the raw logs and unpacked data for FeeConfiguration events raised by the V3QuickSwapFactory contract.
type V3QuickSwapFactoryFeeConfigurationIterator struct {
	Event *V3QuickSwapFactoryFeeConfiguration // Event containing the contract specifics and raw log

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
func (it *V3QuickSwapFactoryFeeConfigurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3QuickSwapFactoryFeeConfiguration)
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
		it.Event = new(V3QuickSwapFactoryFeeConfiguration)
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
func (it *V3QuickSwapFactoryFeeConfigurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3QuickSwapFactoryFeeConfigurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3QuickSwapFactoryFeeConfiguration represents a FeeConfiguration event raised by the V3QuickSwapFactory contract.
type V3QuickSwapFactoryFeeConfiguration struct {
	Alpha1      uint16
	Alpha2      uint16
	Beta1       uint32
	Beta2       uint32
	Gamma1      uint16
	Gamma2      uint16
	VolumeBeta  uint32
	VolumeGamma uint16
	BaseFee     uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeConfiguration is a free log retrieval operation binding the contract event 0x4035ab409f15e202f9f114632e1fb14a0552325955722be18503403e7f98730c.
//
// Solidity: event FeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_V3QuickSwapFactory *V3QuickSwapFactoryFilterer) FilterFeeConfiguration(opts *bind.FilterOpts) (*V3QuickSwapFactoryFeeConfigurationIterator, error) {

	logs, sub, err := _V3QuickSwapFactory.contract.FilterLogs(opts, "FeeConfiguration")
	if err != nil {
		return nil, err
	}
	return &V3QuickSwapFactoryFeeConfigurationIterator{contract: _V3QuickSwapFactory.contract, event: "FeeConfiguration", logs: logs, sub: sub}, nil
}

// WatchFeeConfiguration is a free log subscription operation binding the contract event 0x4035ab409f15e202f9f114632e1fb14a0552325955722be18503403e7f98730c.
//
// Solidity: event FeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_V3QuickSwapFactory *V3QuickSwapFactoryFilterer) WatchFeeConfiguration(opts *bind.WatchOpts, sink chan<- *V3QuickSwapFactoryFeeConfiguration) (event.Subscription, error) {

	logs, sub, err := _V3QuickSwapFactory.contract.WatchLogs(opts, "FeeConfiguration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3QuickSwapFactoryFeeConfiguration)
				if err := _V3QuickSwapFactory.contract.UnpackLog(event, "FeeConfiguration", log); err != nil {
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

// ParseFeeConfiguration is a log parse operation binding the contract event 0x4035ab409f15e202f9f114632e1fb14a0552325955722be18503403e7f98730c.
//
// Solidity: event FeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (_V3QuickSwapFactory *V3QuickSwapFactoryFilterer) ParseFeeConfiguration(log types.Log) (*V3QuickSwapFactoryFeeConfiguration, error) {
	event := new(V3QuickSwapFactoryFeeConfiguration)
	if err := _V3QuickSwapFactory.contract.UnpackLog(event, "FeeConfiguration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3QuickSwapFactoryOwnerIterator is returned from FilterOwner and is used to iterate over the raw logs and unpacked data for Owner events raised by the V3QuickSwapFactory contract.
type V3QuickSwapFactoryOwnerIterator struct {
	Event *V3QuickSwapFactoryOwner // Event containing the contract specifics and raw log

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
func (it *V3QuickSwapFactoryOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3QuickSwapFactoryOwner)
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
		it.Event = new(V3QuickSwapFactoryOwner)
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
func (it *V3QuickSwapFactoryOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3QuickSwapFactoryOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3QuickSwapFactoryOwner represents a Owner event raised by the V3QuickSwapFactory contract.
type V3QuickSwapFactoryOwner struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwner is a free log retrieval operation binding the contract event 0xa5e220c2c27d986cc8efeafa8f34ba6ea6bf96a34e146b29b6bdd8587771b130.
//
// Solidity: event Owner(address indexed newOwner)
func (_V3QuickSwapFactory *V3QuickSwapFactoryFilterer) FilterOwner(opts *bind.FilterOpts, newOwner []common.Address) (*V3QuickSwapFactoryOwnerIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _V3QuickSwapFactory.contract.FilterLogs(opts, "Owner", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &V3QuickSwapFactoryOwnerIterator{contract: _V3QuickSwapFactory.contract, event: "Owner", logs: logs, sub: sub}, nil
}

// WatchOwner is a free log subscription operation binding the contract event 0xa5e220c2c27d986cc8efeafa8f34ba6ea6bf96a34e146b29b6bdd8587771b130.
//
// Solidity: event Owner(address indexed newOwner)
func (_V3QuickSwapFactory *V3QuickSwapFactoryFilterer) WatchOwner(opts *bind.WatchOpts, sink chan<- *V3QuickSwapFactoryOwner, newOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _V3QuickSwapFactory.contract.WatchLogs(opts, "Owner", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3QuickSwapFactoryOwner)
				if err := _V3QuickSwapFactory.contract.UnpackLog(event, "Owner", log); err != nil {
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

// ParseOwner is a log parse operation binding the contract event 0xa5e220c2c27d986cc8efeafa8f34ba6ea6bf96a34e146b29b6bdd8587771b130.
//
// Solidity: event Owner(address indexed newOwner)
func (_V3QuickSwapFactory *V3QuickSwapFactoryFilterer) ParseOwner(log types.Log) (*V3QuickSwapFactoryOwner, error) {
	event := new(V3QuickSwapFactoryOwner)
	if err := _V3QuickSwapFactory.contract.UnpackLog(event, "Owner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3QuickSwapFactoryPoolIterator is returned from FilterPool and is used to iterate over the raw logs and unpacked data for Pool events raised by the V3QuickSwapFactory contract.
type V3QuickSwapFactoryPoolIterator struct {
	Event *V3QuickSwapFactoryPool // Event containing the contract specifics and raw log

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
func (it *V3QuickSwapFactoryPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3QuickSwapFactoryPool)
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
		it.Event = new(V3QuickSwapFactoryPool)
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
func (it *V3QuickSwapFactoryPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3QuickSwapFactoryPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3QuickSwapFactoryPool represents a Pool event raised by the V3QuickSwapFactory contract.
type V3QuickSwapFactoryPool struct {
	Token0 common.Address
	Token1 common.Address
	Pool   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPool is a free log retrieval operation binding the contract event 0x91ccaa7a278130b65168c3a0c8d3bcae84cf5e43704342bd3ec0b59e59c036db.
//
// Solidity: event Pool(address indexed token0, address indexed token1, address pool)
func (_V3QuickSwapFactory *V3QuickSwapFactoryFilterer) FilterPool(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*V3QuickSwapFactoryPoolIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _V3QuickSwapFactory.contract.FilterLogs(opts, "Pool", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &V3QuickSwapFactoryPoolIterator{contract: _V3QuickSwapFactory.contract, event: "Pool", logs: logs, sub: sub}, nil
}

// WatchPool is a free log subscription operation binding the contract event 0x91ccaa7a278130b65168c3a0c8d3bcae84cf5e43704342bd3ec0b59e59c036db.
//
// Solidity: event Pool(address indexed token0, address indexed token1, address pool)
func (_V3QuickSwapFactory *V3QuickSwapFactoryFilterer) WatchPool(opts *bind.WatchOpts, sink chan<- *V3QuickSwapFactoryPool, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _V3QuickSwapFactory.contract.WatchLogs(opts, "Pool", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3QuickSwapFactoryPool)
				if err := _V3QuickSwapFactory.contract.UnpackLog(event, "Pool", log); err != nil {
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

// ParsePool is a log parse operation binding the contract event 0x91ccaa7a278130b65168c3a0c8d3bcae84cf5e43704342bd3ec0b59e59c036db.
//
// Solidity: event Pool(address indexed token0, address indexed token1, address pool)
func (_V3QuickSwapFactory *V3QuickSwapFactoryFilterer) ParsePool(log types.Log) (*V3QuickSwapFactoryPool, error) {
	event := new(V3QuickSwapFactoryPool)
	if err := _V3QuickSwapFactory.contract.UnpackLog(event, "Pool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V3QuickSwapFactoryVaultAddressIterator is returned from FilterVaultAddress and is used to iterate over the raw logs and unpacked data for VaultAddress events raised by the V3QuickSwapFactory contract.
type V3QuickSwapFactoryVaultAddressIterator struct {
	Event *V3QuickSwapFactoryVaultAddress // Event containing the contract specifics and raw log

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
func (it *V3QuickSwapFactoryVaultAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V3QuickSwapFactoryVaultAddress)
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
		it.Event = new(V3QuickSwapFactoryVaultAddress)
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
func (it *V3QuickSwapFactoryVaultAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V3QuickSwapFactoryVaultAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V3QuickSwapFactoryVaultAddress represents a VaultAddress event raised by the V3QuickSwapFactory contract.
type V3QuickSwapFactoryVaultAddress struct {
	NewVaultAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVaultAddress is a free log retrieval operation binding the contract event 0xb9c265ae4414f501736ec5d4961edc3309e4385eb2ff3feeecb30fb36621dd83.
//
// Solidity: event VaultAddress(address indexed newVaultAddress)
func (_V3QuickSwapFactory *V3QuickSwapFactoryFilterer) FilterVaultAddress(opts *bind.FilterOpts, newVaultAddress []common.Address) (*V3QuickSwapFactoryVaultAddressIterator, error) {

	var newVaultAddressRule []interface{}
	for _, newVaultAddressItem := range newVaultAddress {
		newVaultAddressRule = append(newVaultAddressRule, newVaultAddressItem)
	}

	logs, sub, err := _V3QuickSwapFactory.contract.FilterLogs(opts, "VaultAddress", newVaultAddressRule)
	if err != nil {
		return nil, err
	}
	return &V3QuickSwapFactoryVaultAddressIterator{contract: _V3QuickSwapFactory.contract, event: "VaultAddress", logs: logs, sub: sub}, nil
}

// WatchVaultAddress is a free log subscription operation binding the contract event 0xb9c265ae4414f501736ec5d4961edc3309e4385eb2ff3feeecb30fb36621dd83.
//
// Solidity: event VaultAddress(address indexed newVaultAddress)
func (_V3QuickSwapFactory *V3QuickSwapFactoryFilterer) WatchVaultAddress(opts *bind.WatchOpts, sink chan<- *V3QuickSwapFactoryVaultAddress, newVaultAddress []common.Address) (event.Subscription, error) {

	var newVaultAddressRule []interface{}
	for _, newVaultAddressItem := range newVaultAddress {
		newVaultAddressRule = append(newVaultAddressRule, newVaultAddressItem)
	}

	logs, sub, err := _V3QuickSwapFactory.contract.WatchLogs(opts, "VaultAddress", newVaultAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V3QuickSwapFactoryVaultAddress)
				if err := _V3QuickSwapFactory.contract.UnpackLog(event, "VaultAddress", log); err != nil {
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

// ParseVaultAddress is a log parse operation binding the contract event 0xb9c265ae4414f501736ec5d4961edc3309e4385eb2ff3feeecb30fb36621dd83.
//
// Solidity: event VaultAddress(address indexed newVaultAddress)
func (_V3QuickSwapFactory *V3QuickSwapFactoryFilterer) ParseVaultAddress(log types.Log) (*V3QuickSwapFactoryVaultAddress, error) {
	event := new(V3QuickSwapFactoryVaultAddress)
	if err := _V3QuickSwapFactory.contract.UnpackLog(event, "VaultAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
