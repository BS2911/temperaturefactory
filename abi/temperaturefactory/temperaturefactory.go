// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package temperaturefactory

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
)

// TemperaturefactoryMetaData contains all meta data concerning the Temperaturefactory contract.
var TemperaturefactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TEMPERFAC_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TEMPERFAC_WORKER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTemperature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setTemperature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_worker\",\"type\":\"address\"}],\"name\":\"setWorker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"temperature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TemperaturefactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use TemperaturefactoryMetaData.ABI instead.
var TemperaturefactoryABI = TemperaturefactoryMetaData.ABI

// Temperaturefactory is an auto generated Go binding around an Ethereum contract.
type Temperaturefactory struct {
	TemperaturefactoryCaller     // Read-only binding to the contract
	TemperaturefactoryTransactor // Write-only binding to the contract
	TemperaturefactoryFilterer   // Log filterer for contract events
}

// TemperaturefactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TemperaturefactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TemperaturefactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TemperaturefactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TemperaturefactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TemperaturefactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TemperaturefactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TemperaturefactorySession struct {
	Contract     *Temperaturefactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TemperaturefactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TemperaturefactoryCallerSession struct {
	Contract *TemperaturefactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TemperaturefactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TemperaturefactoryTransactorSession struct {
	Contract     *TemperaturefactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TemperaturefactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TemperaturefactoryRaw struct {
	Contract *Temperaturefactory // Generic contract binding to access the raw methods on
}

// TemperaturefactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TemperaturefactoryCallerRaw struct {
	Contract *TemperaturefactoryCaller // Generic read-only contract binding to access the raw methods on
}

// TemperaturefactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TemperaturefactoryTransactorRaw struct {
	Contract *TemperaturefactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTemperaturefactory creates a new instance of Temperaturefactory, bound to a specific deployed contract.
func NewTemperaturefactory(address common.Address, backend bind.ContractBackend) (*Temperaturefactory, error) {
	contract, err := bindTemperaturefactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Temperaturefactory{TemperaturefactoryCaller: TemperaturefactoryCaller{contract: contract}, TemperaturefactoryTransactor: TemperaturefactoryTransactor{contract: contract}, TemperaturefactoryFilterer: TemperaturefactoryFilterer{contract: contract}}, nil
}

// NewTemperaturefactoryCaller creates a new read-only instance of Temperaturefactory, bound to a specific deployed contract.
func NewTemperaturefactoryCaller(address common.Address, caller bind.ContractCaller) (*TemperaturefactoryCaller, error) {
	contract, err := bindTemperaturefactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TemperaturefactoryCaller{contract: contract}, nil
}

// NewTemperaturefactoryTransactor creates a new write-only instance of Temperaturefactory, bound to a specific deployed contract.
func NewTemperaturefactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*TemperaturefactoryTransactor, error) {
	contract, err := bindTemperaturefactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TemperaturefactoryTransactor{contract: contract}, nil
}

// NewTemperaturefactoryFilterer creates a new log filterer instance of Temperaturefactory, bound to a specific deployed contract.
func NewTemperaturefactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*TemperaturefactoryFilterer, error) {
	contract, err := bindTemperaturefactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TemperaturefactoryFilterer{contract: contract}, nil
}

// bindTemperaturefactory binds a generic wrapper to an already deployed contract.
func bindTemperaturefactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TemperaturefactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Temperaturefactory *TemperaturefactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Temperaturefactory.Contract.TemperaturefactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Temperaturefactory *TemperaturefactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Temperaturefactory.Contract.TemperaturefactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Temperaturefactory *TemperaturefactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Temperaturefactory.Contract.TemperaturefactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Temperaturefactory *TemperaturefactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Temperaturefactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Temperaturefactory *TemperaturefactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Temperaturefactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Temperaturefactory *TemperaturefactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Temperaturefactory.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Temperaturefactory *TemperaturefactoryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Temperaturefactory.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Temperaturefactory *TemperaturefactorySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Temperaturefactory.Contract.DEFAULTADMINROLE(&_Temperaturefactory.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Temperaturefactory *TemperaturefactoryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Temperaturefactory.Contract.DEFAULTADMINROLE(&_Temperaturefactory.CallOpts)
}

// TEMPERFACMANAGERROLE is a free data retrieval call binding the contract method 0x5cf55085.
//
// Solidity: function TEMPERFAC_MANAGER_ROLE() view returns(bytes32)
func (_Temperaturefactory *TemperaturefactoryCaller) TEMPERFACMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Temperaturefactory.contract.Call(opts, &out, "TEMPERFAC_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TEMPERFACMANAGERROLE is a free data retrieval call binding the contract method 0x5cf55085.
//
// Solidity: function TEMPERFAC_MANAGER_ROLE() view returns(bytes32)
func (_Temperaturefactory *TemperaturefactorySession) TEMPERFACMANAGERROLE() ([32]byte, error) {
	return _Temperaturefactory.Contract.TEMPERFACMANAGERROLE(&_Temperaturefactory.CallOpts)
}

// TEMPERFACMANAGERROLE is a free data retrieval call binding the contract method 0x5cf55085.
//
// Solidity: function TEMPERFAC_MANAGER_ROLE() view returns(bytes32)
func (_Temperaturefactory *TemperaturefactoryCallerSession) TEMPERFACMANAGERROLE() ([32]byte, error) {
	return _Temperaturefactory.Contract.TEMPERFACMANAGERROLE(&_Temperaturefactory.CallOpts)
}

// TEMPERFACWORKERROLE is a free data retrieval call binding the contract method 0xe68b69aa.
//
// Solidity: function TEMPERFAC_WORKER_ROLE() view returns(bytes32)
func (_Temperaturefactory *TemperaturefactoryCaller) TEMPERFACWORKERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Temperaturefactory.contract.Call(opts, &out, "TEMPERFAC_WORKER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TEMPERFACWORKERROLE is a free data retrieval call binding the contract method 0xe68b69aa.
//
// Solidity: function TEMPERFAC_WORKER_ROLE() view returns(bytes32)
func (_Temperaturefactory *TemperaturefactorySession) TEMPERFACWORKERROLE() ([32]byte, error) {
	return _Temperaturefactory.Contract.TEMPERFACWORKERROLE(&_Temperaturefactory.CallOpts)
}

// TEMPERFACWORKERROLE is a free data retrieval call binding the contract method 0xe68b69aa.
//
// Solidity: function TEMPERFAC_WORKER_ROLE() view returns(bytes32)
func (_Temperaturefactory *TemperaturefactoryCallerSession) TEMPERFACWORKERROLE() ([32]byte, error) {
	return _Temperaturefactory.Contract.TEMPERFACWORKERROLE(&_Temperaturefactory.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Temperaturefactory *TemperaturefactoryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Temperaturefactory.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Temperaturefactory *TemperaturefactorySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Temperaturefactory.Contract.GetRoleAdmin(&_Temperaturefactory.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Temperaturefactory *TemperaturefactoryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Temperaturefactory.Contract.GetRoleAdmin(&_Temperaturefactory.CallOpts, role)
}

// GetTemperature is a free data retrieval call binding the contract method 0x6421d04b.
//
// Solidity: function getTemperature() view returns(uint256)
func (_Temperaturefactory *TemperaturefactoryCaller) GetTemperature(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Temperaturefactory.contract.Call(opts, &out, "getTemperature")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTemperature is a free data retrieval call binding the contract method 0x6421d04b.
//
// Solidity: function getTemperature() view returns(uint256)
func (_Temperaturefactory *TemperaturefactorySession) GetTemperature() (*big.Int, error) {
	return _Temperaturefactory.Contract.GetTemperature(&_Temperaturefactory.CallOpts)
}

// GetTemperature is a free data retrieval call binding the contract method 0x6421d04b.
//
// Solidity: function getTemperature() view returns(uint256)
func (_Temperaturefactory *TemperaturefactoryCallerSession) GetTemperature() (*big.Int, error) {
	return _Temperaturefactory.Contract.GetTemperature(&_Temperaturefactory.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Temperaturefactory *TemperaturefactoryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Temperaturefactory.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Temperaturefactory *TemperaturefactorySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Temperaturefactory.Contract.HasRole(&_Temperaturefactory.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Temperaturefactory *TemperaturefactoryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Temperaturefactory.Contract.HasRole(&_Temperaturefactory.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Temperaturefactory *TemperaturefactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Temperaturefactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Temperaturefactory *TemperaturefactorySession) Owner() (common.Address, error) {
	return _Temperaturefactory.Contract.Owner(&_Temperaturefactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Temperaturefactory *TemperaturefactoryCallerSession) Owner() (common.Address, error) {
	return _Temperaturefactory.Contract.Owner(&_Temperaturefactory.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Temperaturefactory *TemperaturefactoryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Temperaturefactory.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Temperaturefactory *TemperaturefactorySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Temperaturefactory.Contract.SupportsInterface(&_Temperaturefactory.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Temperaturefactory *TemperaturefactoryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Temperaturefactory.Contract.SupportsInterface(&_Temperaturefactory.CallOpts, interfaceId)
}

// Temperature is a free data retrieval call binding the contract method 0xadccea12.
//
// Solidity: function temperature() view returns(uint256)
func (_Temperaturefactory *TemperaturefactoryCaller) Temperature(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Temperaturefactory.contract.Call(opts, &out, "temperature")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Temperature is a free data retrieval call binding the contract method 0xadccea12.
//
// Solidity: function temperature() view returns(uint256)
func (_Temperaturefactory *TemperaturefactorySession) Temperature() (*big.Int, error) {
	return _Temperaturefactory.Contract.Temperature(&_Temperaturefactory.CallOpts)
}

// Temperature is a free data retrieval call binding the contract method 0xadccea12.
//
// Solidity: function temperature() view returns(uint256)
func (_Temperaturefactory *TemperaturefactoryCallerSession) Temperature() (*big.Int, error) {
	return _Temperaturefactory.Contract.Temperature(&_Temperaturefactory.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Temperaturefactory *TemperaturefactoryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Temperaturefactory.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Temperaturefactory *TemperaturefactorySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Temperaturefactory.Contract.GrantRole(&_Temperaturefactory.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Temperaturefactory *TemperaturefactoryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Temperaturefactory.Contract.GrantRole(&_Temperaturefactory.TransactOpts, role, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Temperaturefactory *TemperaturefactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Temperaturefactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Temperaturefactory *TemperaturefactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _Temperaturefactory.Contract.RenounceOwnership(&_Temperaturefactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Temperaturefactory *TemperaturefactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Temperaturefactory.Contract.RenounceOwnership(&_Temperaturefactory.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Temperaturefactory *TemperaturefactoryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Temperaturefactory.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Temperaturefactory *TemperaturefactorySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Temperaturefactory.Contract.RenounceRole(&_Temperaturefactory.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Temperaturefactory *TemperaturefactoryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Temperaturefactory.Contract.RenounceRole(&_Temperaturefactory.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Temperaturefactory *TemperaturefactoryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Temperaturefactory.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Temperaturefactory *TemperaturefactorySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Temperaturefactory.Contract.RevokeRole(&_Temperaturefactory.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Temperaturefactory *TemperaturefactoryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Temperaturefactory.Contract.RevokeRole(&_Temperaturefactory.TransactOpts, role, account)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address _manager) returns()
func (_Temperaturefactory *TemperaturefactoryTransactor) SetManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _Temperaturefactory.contract.Transact(opts, "setManager", _manager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address _manager) returns()
func (_Temperaturefactory *TemperaturefactorySession) SetManager(_manager common.Address) (*types.Transaction, error) {
	return _Temperaturefactory.Contract.SetManager(&_Temperaturefactory.TransactOpts, _manager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address _manager) returns()
func (_Temperaturefactory *TemperaturefactoryTransactorSession) SetManager(_manager common.Address) (*types.Transaction, error) {
	return _Temperaturefactory.Contract.SetManager(&_Temperaturefactory.TransactOpts, _manager)
}

// SetTemperature is a paid mutator transaction binding the contract method 0x48f14470.
//
// Solidity: function setTemperature(uint256 _value) returns()
func (_Temperaturefactory *TemperaturefactoryTransactor) SetTemperature(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Temperaturefactory.contract.Transact(opts, "setTemperature", _value)
}

// SetTemperature is a paid mutator transaction binding the contract method 0x48f14470.
//
// Solidity: function setTemperature(uint256 _value) returns()
func (_Temperaturefactory *TemperaturefactorySession) SetTemperature(_value *big.Int) (*types.Transaction, error) {
	return _Temperaturefactory.Contract.SetTemperature(&_Temperaturefactory.TransactOpts, _value)
}

// SetTemperature is a paid mutator transaction binding the contract method 0x48f14470.
//
// Solidity: function setTemperature(uint256 _value) returns()
func (_Temperaturefactory *TemperaturefactoryTransactorSession) SetTemperature(_value *big.Int) (*types.Transaction, error) {
	return _Temperaturefactory.Contract.SetTemperature(&_Temperaturefactory.TransactOpts, _value)
}

// SetWorker is a paid mutator transaction binding the contract method 0xc26f6d44.
//
// Solidity: function setWorker(address _worker) returns()
func (_Temperaturefactory *TemperaturefactoryTransactor) SetWorker(opts *bind.TransactOpts, _worker common.Address) (*types.Transaction, error) {
	return _Temperaturefactory.contract.Transact(opts, "setWorker", _worker)
}

// SetWorker is a paid mutator transaction binding the contract method 0xc26f6d44.
//
// Solidity: function setWorker(address _worker) returns()
func (_Temperaturefactory *TemperaturefactorySession) SetWorker(_worker common.Address) (*types.Transaction, error) {
	return _Temperaturefactory.Contract.SetWorker(&_Temperaturefactory.TransactOpts, _worker)
}

// SetWorker is a paid mutator transaction binding the contract method 0xc26f6d44.
//
// Solidity: function setWorker(address _worker) returns()
func (_Temperaturefactory *TemperaturefactoryTransactorSession) SetWorker(_worker common.Address) (*types.Transaction, error) {
	return _Temperaturefactory.Contract.SetWorker(&_Temperaturefactory.TransactOpts, _worker)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Temperaturefactory *TemperaturefactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Temperaturefactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Temperaturefactory *TemperaturefactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Temperaturefactory.Contract.TransferOwnership(&_Temperaturefactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Temperaturefactory *TemperaturefactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Temperaturefactory.Contract.TransferOwnership(&_Temperaturefactory.TransactOpts, newOwner)
}

// TemperaturefactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Temperaturefactory contract.
type TemperaturefactoryOwnershipTransferredIterator struct {
	Event *TemperaturefactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TemperaturefactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TemperaturefactoryOwnershipTransferred)
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
		it.Event = new(TemperaturefactoryOwnershipTransferred)
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
func (it *TemperaturefactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TemperaturefactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TemperaturefactoryOwnershipTransferred represents a OwnershipTransferred event raised by the Temperaturefactory contract.
type TemperaturefactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Temperaturefactory *TemperaturefactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TemperaturefactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Temperaturefactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TemperaturefactoryOwnershipTransferredIterator{contract: _Temperaturefactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Temperaturefactory *TemperaturefactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TemperaturefactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Temperaturefactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TemperaturefactoryOwnershipTransferred)
				if err := _Temperaturefactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Temperaturefactory *TemperaturefactoryFilterer) ParseOwnershipTransferred(log types.Log) (*TemperaturefactoryOwnershipTransferred, error) {
	event := new(TemperaturefactoryOwnershipTransferred)
	if err := _Temperaturefactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TemperaturefactoryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Temperaturefactory contract.
type TemperaturefactoryRoleAdminChangedIterator struct {
	Event *TemperaturefactoryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TemperaturefactoryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TemperaturefactoryRoleAdminChanged)
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
		it.Event = new(TemperaturefactoryRoleAdminChanged)
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
func (it *TemperaturefactoryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TemperaturefactoryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TemperaturefactoryRoleAdminChanged represents a RoleAdminChanged event raised by the Temperaturefactory contract.
type TemperaturefactoryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Temperaturefactory *TemperaturefactoryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TemperaturefactoryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Temperaturefactory.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TemperaturefactoryRoleAdminChangedIterator{contract: _Temperaturefactory.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Temperaturefactory *TemperaturefactoryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TemperaturefactoryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Temperaturefactory.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TemperaturefactoryRoleAdminChanged)
				if err := _Temperaturefactory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Temperaturefactory *TemperaturefactoryFilterer) ParseRoleAdminChanged(log types.Log) (*TemperaturefactoryRoleAdminChanged, error) {
	event := new(TemperaturefactoryRoleAdminChanged)
	if err := _Temperaturefactory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TemperaturefactoryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Temperaturefactory contract.
type TemperaturefactoryRoleGrantedIterator struct {
	Event *TemperaturefactoryRoleGranted // Event containing the contract specifics and raw log

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
func (it *TemperaturefactoryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TemperaturefactoryRoleGranted)
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
		it.Event = new(TemperaturefactoryRoleGranted)
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
func (it *TemperaturefactoryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TemperaturefactoryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TemperaturefactoryRoleGranted represents a RoleGranted event raised by the Temperaturefactory contract.
type TemperaturefactoryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Temperaturefactory *TemperaturefactoryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TemperaturefactoryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Temperaturefactory.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TemperaturefactoryRoleGrantedIterator{contract: _Temperaturefactory.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Temperaturefactory *TemperaturefactoryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TemperaturefactoryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Temperaturefactory.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TemperaturefactoryRoleGranted)
				if err := _Temperaturefactory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Temperaturefactory *TemperaturefactoryFilterer) ParseRoleGranted(log types.Log) (*TemperaturefactoryRoleGranted, error) {
	event := new(TemperaturefactoryRoleGranted)
	if err := _Temperaturefactory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TemperaturefactoryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Temperaturefactory contract.
type TemperaturefactoryRoleRevokedIterator struct {
	Event *TemperaturefactoryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TemperaturefactoryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TemperaturefactoryRoleRevoked)
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
		it.Event = new(TemperaturefactoryRoleRevoked)
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
func (it *TemperaturefactoryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TemperaturefactoryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TemperaturefactoryRoleRevoked represents a RoleRevoked event raised by the Temperaturefactory contract.
type TemperaturefactoryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Temperaturefactory *TemperaturefactoryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TemperaturefactoryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Temperaturefactory.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TemperaturefactoryRoleRevokedIterator{contract: _Temperaturefactory.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Temperaturefactory *TemperaturefactoryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TemperaturefactoryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Temperaturefactory.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TemperaturefactoryRoleRevoked)
				if err := _Temperaturefactory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Temperaturefactory *TemperaturefactoryFilterer) ParseRoleRevoked(log types.Log) (*TemperaturefactoryRoleRevoked, error) {
	event := new(TemperaturefactoryRoleRevoked)
	if err := _Temperaturefactory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
