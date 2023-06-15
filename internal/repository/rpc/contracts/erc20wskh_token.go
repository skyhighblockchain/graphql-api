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
)

// ErcWrappedSkhMetaData contains all meta data concerning the ErcWrappedSkh contract.
var ErcWrappedSkhMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_INVALID_ZERO_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_NO_ERROR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ErcWrappedSkhABI is the input ABI used to generate the binding from.
// Deprecated: Use ErcWrappedSkhMetaData.ABI instead.
var ErcWrappedSkhABI = ErcWrappedSkhMetaData.ABI

// ErcWrappedSkh is an auto generated Go binding around an Ethereum contract.
type ErcWrappedSkh struct {
	ErcWrappedSkhCaller     // Read-only binding to the contract
	ErcWrappedSkhTransactor // Write-only binding to the contract
	ErcWrappedSkhFilterer   // Log filterer for contract events
}

// ErcWrappedSkhCaller is an auto generated read-only Go binding around an Ethereum contract.
type ErcWrappedSkhCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErcWrappedSkhTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ErcWrappedSkhTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErcWrappedSkhFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ErcWrappedSkhFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErcWrappedSkhSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ErcWrappedSkhSession struct {
	Contract     *ErcWrappedSkh    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ErcWrappedSkhCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ErcWrappedSkhCallerSession struct {
	Contract *ErcWrappedSkhCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ErcWrappedSkhTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ErcWrappedSkhTransactorSession struct {
	Contract     *ErcWrappedSkhTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ErcWrappedSkhRaw is an auto generated low-level Go binding around an Ethereum contract.
type ErcWrappedSkhRaw struct {
	Contract *ErcWrappedSkh // Generic contract binding to access the raw methods on
}

// ErcWrappedSkhCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ErcWrappedSkhCallerRaw struct {
	Contract *ErcWrappedSkhCaller // Generic read-only contract binding to access the raw methods on
}

// ErcWrappedSkhTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ErcWrappedSkhTransactorRaw struct {
	Contract *ErcWrappedSkhTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErcWrappedSkh creates a new instance of ErcWrappedSkh, bound to a specific deployed contract.
func NewErcWrappedSkh(address common.Address, backend bind.ContractBackend) (*ErcWrappedSkh, error) {
	contract, err := bindErcWrappedSkh(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedSkh{ErcWrappedSkhCaller: ErcWrappedSkhCaller{contract: contract}, ErcWrappedSkhTransactor: ErcWrappedSkhTransactor{contract: contract}, ErcWrappedSkhFilterer: ErcWrappedSkhFilterer{contract: contract}}, nil
}

// NewErcWrappedSkhCaller creates a new read-only instance of ErcWrappedSkh, bound to a specific deployed contract.
func NewErcWrappedSkhCaller(address common.Address, caller bind.ContractCaller) (*ErcWrappedSkhCaller, error) {
	contract, err := bindErcWrappedSkh(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedSkhCaller{contract: contract}, nil
}

// NewErcWrappedSkhTransactor creates a new write-only instance of ErcWrappedSkh, bound to a specific deployed contract.
func NewErcWrappedSkhTransactor(address common.Address, transactor bind.ContractTransactor) (*ErcWrappedSkhTransactor, error) {
	contract, err := bindErcWrappedSkh(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedSkhTransactor{contract: contract}, nil
}

// NewErcWrappedSkhFilterer creates a new log filterer instance of ErcWrappedSkh, bound to a specific deployed contract.
func NewErcWrappedSkhFilterer(address common.Address, filterer bind.ContractFilterer) (*ErcWrappedSkhFilterer, error) {
	contract, err := bindErcWrappedSkh(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedSkhFilterer{contract: contract}, nil
}

// bindErcWrappedSkh binds a generic wrapper to an already deployed contract.
func bindErcWrappedSkh(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ErcWrappedSkhABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ErcWrappedSkh *ErcWrappedSkhRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ErcWrappedSkh.Contract.ErcWrappedSkhCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ErcWrappedSkh *ErcWrappedSkhRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.ErcWrappedSkhTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ErcWrappedSkh *ErcWrappedSkhRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.ErcWrappedSkhTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ErcWrappedSkh *ErcWrappedSkhCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ErcWrappedSkh.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ErcWrappedSkh *ErcWrappedSkhTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ErcWrappedSkh *ErcWrappedSkhTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.contract.Transact(opts, method, params...)
}

// ERRINVALIDZEROVALUE is a free data retrieval call binding the contract method 0x6d7497b3.
//
// Solidity: function ERR_INVALID_ZERO_VALUE() view returns(uint256)
func (_ErcWrappedSkh *ErcWrappedSkhCaller) ERRINVALIDZEROVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ErcWrappedSkh.contract.Call(opts, &out, "ERR_INVALID_ZERO_VALUE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRINVALIDZEROVALUE is a free data retrieval call binding the contract method 0x6d7497b3.
//
// Solidity: function ERR_INVALID_ZERO_VALUE() view returns(uint256)
func (_ErcWrappedSkh *ErcWrappedSkhSession) ERRINVALIDZEROVALUE() (*big.Int, error) {
	return _ErcWrappedSkh.Contract.ERRINVALIDZEROVALUE(&_ErcWrappedSkh.CallOpts)
}

// ERRINVALIDZEROVALUE is a free data retrieval call binding the contract method 0x6d7497b3.
//
// Solidity: function ERR_INVALID_ZERO_VALUE() view returns(uint256)
func (_ErcWrappedSkh *ErcWrappedSkhCallerSession) ERRINVALIDZEROVALUE() (*big.Int, error) {
	return _ErcWrappedSkh.Contract.ERRINVALIDZEROVALUE(&_ErcWrappedSkh.CallOpts)
}

// ERRNOERROR is a free data retrieval call binding the contract method 0x35052d6e.
//
// Solidity: function ERR_NO_ERROR() view returns(uint256)
func (_ErcWrappedSkh *ErcWrappedSkhCaller) ERRNOERROR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ErcWrappedSkh.contract.Call(opts, &out, "ERR_NO_ERROR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERRNOERROR is a free data retrieval call binding the contract method 0x35052d6e.
//
// Solidity: function ERR_NO_ERROR() view returns(uint256)
func (_ErcWrappedSkh *ErcWrappedSkhSession) ERRNOERROR() (*big.Int, error) {
	return _ErcWrappedSkh.Contract.ERRNOERROR(&_ErcWrappedSkh.CallOpts)
}

// ERRNOERROR is a free data retrieval call binding the contract method 0x35052d6e.
//
// Solidity: function ERR_NO_ERROR() view returns(uint256)
func (_ErcWrappedSkh *ErcWrappedSkhCallerSession) ERRNOERROR() (*big.Int, error) {
	return _ErcWrappedSkh.Contract.ERRNOERROR(&_ErcWrappedSkh.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ErcWrappedSkh *ErcWrappedSkhCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ErcWrappedSkh.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ErcWrappedSkh *ErcWrappedSkhSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ErcWrappedSkh.Contract.Allowance(&_ErcWrappedSkh.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ErcWrappedSkh *ErcWrappedSkhCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ErcWrappedSkh.Contract.Allowance(&_ErcWrappedSkh.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ErcWrappedSkh *ErcWrappedSkhCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ErcWrappedSkh.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ErcWrappedSkh *ErcWrappedSkhSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ErcWrappedSkh.Contract.BalanceOf(&_ErcWrappedSkh.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ErcWrappedSkh *ErcWrappedSkhCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ErcWrappedSkh.Contract.BalanceOf(&_ErcWrappedSkh.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ErcWrappedSkh *ErcWrappedSkhCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ErcWrappedSkh.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ErcWrappedSkh *ErcWrappedSkhSession) Decimals() (uint8, error) {
	return _ErcWrappedSkh.Contract.Decimals(&_ErcWrappedSkh.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ErcWrappedSkh *ErcWrappedSkhCallerSession) Decimals() (uint8, error) {
	return _ErcWrappedSkh.Contract.Decimals(&_ErcWrappedSkh.CallOpts)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_ErcWrappedSkh *ErcWrappedSkhCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _ErcWrappedSkh.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_ErcWrappedSkh *ErcWrappedSkhSession) IsPauser(account common.Address) (bool, error) {
	return _ErcWrappedSkh.Contract.IsPauser(&_ErcWrappedSkh.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_ErcWrappedSkh *ErcWrappedSkhCallerSession) IsPauser(account common.Address) (bool, error) {
	return _ErcWrappedSkh.Contract.IsPauser(&_ErcWrappedSkh.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ErcWrappedSkh *ErcWrappedSkhCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ErcWrappedSkh.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ErcWrappedSkh *ErcWrappedSkhSession) Name() (string, error) {
	return _ErcWrappedSkh.Contract.Name(&_ErcWrappedSkh.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ErcWrappedSkh *ErcWrappedSkhCallerSession) Name() (string, error) {
	return _ErcWrappedSkh.Contract.Name(&_ErcWrappedSkh.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ErcWrappedSkh *ErcWrappedSkhCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ErcWrappedSkh.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ErcWrappedSkh *ErcWrappedSkhSession) Paused() (bool, error) {
	return _ErcWrappedSkh.Contract.Paused(&_ErcWrappedSkh.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ErcWrappedSkh *ErcWrappedSkhCallerSession) Paused() (bool, error) {
	return _ErcWrappedSkh.Contract.Paused(&_ErcWrappedSkh.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ErcWrappedSkh *ErcWrappedSkhCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ErcWrappedSkh.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ErcWrappedSkh *ErcWrappedSkhSession) Symbol() (string, error) {
	return _ErcWrappedSkh.Contract.Symbol(&_ErcWrappedSkh.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ErcWrappedSkh *ErcWrappedSkhCallerSession) Symbol() (string, error) {
	return _ErcWrappedSkh.Contract.Symbol(&_ErcWrappedSkh.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ErcWrappedSkh *ErcWrappedSkhCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ErcWrappedSkh.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ErcWrappedSkh *ErcWrappedSkhSession) TotalSupply() (*big.Int, error) {
	return _ErcWrappedSkh.Contract.TotalSupply(&_ErcWrappedSkh.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ErcWrappedSkh *ErcWrappedSkhCallerSession) TotalSupply() (*big.Int, error) {
	return _ErcWrappedSkh.Contract.TotalSupply(&_ErcWrappedSkh.CallOpts)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ErcWrappedSkh *ErcWrappedSkhTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ErcWrappedSkh.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ErcWrappedSkh *ErcWrappedSkhSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.AddPauser(&_ErcWrappedSkh.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ErcWrappedSkh *ErcWrappedSkhTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.AddPauser(&_ErcWrappedSkh.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ErcWrappedSkh *ErcWrappedSkhTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedSkh.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ErcWrappedSkh *ErcWrappedSkhSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.Approve(&_ErcWrappedSkh.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ErcWrappedSkh *ErcWrappedSkhTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.Approve(&_ErcWrappedSkh.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ErcWrappedSkh *ErcWrappedSkhTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedSkh.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ErcWrappedSkh *ErcWrappedSkhSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.DecreaseAllowance(&_ErcWrappedSkh.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ErcWrappedSkh *ErcWrappedSkhTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.DecreaseAllowance(&_ErcWrappedSkh.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(uint256)
func (_ErcWrappedSkh *ErcWrappedSkhTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedSkh.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(uint256)
func (_ErcWrappedSkh *ErcWrappedSkhSession) Deposit() (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.Deposit(&_ErcWrappedSkh.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(uint256)
func (_ErcWrappedSkh *ErcWrappedSkhTransactorSession) Deposit() (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.Deposit(&_ErcWrappedSkh.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ErcWrappedSkh *ErcWrappedSkhTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedSkh.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ErcWrappedSkh *ErcWrappedSkhSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.IncreaseAllowance(&_ErcWrappedSkh.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ErcWrappedSkh *ErcWrappedSkhTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.IncreaseAllowance(&_ErcWrappedSkh.TransactOpts, spender, addedValue)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ErcWrappedSkh *ErcWrappedSkhTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedSkh.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ErcWrappedSkh *ErcWrappedSkhSession) Pause() (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.Pause(&_ErcWrappedSkh.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ErcWrappedSkh *ErcWrappedSkhTransactorSession) Pause() (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.Pause(&_ErcWrappedSkh.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ErcWrappedSkh *ErcWrappedSkhTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedSkh.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ErcWrappedSkh *ErcWrappedSkhSession) RenouncePauser() (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.RenouncePauser(&_ErcWrappedSkh.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ErcWrappedSkh *ErcWrappedSkhTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.RenouncePauser(&_ErcWrappedSkh.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ErcWrappedSkh *ErcWrappedSkhTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedSkh.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ErcWrappedSkh *ErcWrappedSkhSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.Transfer(&_ErcWrappedSkh.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ErcWrappedSkh *ErcWrappedSkhTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.Transfer(&_ErcWrappedSkh.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ErcWrappedSkh *ErcWrappedSkhTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedSkh.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ErcWrappedSkh *ErcWrappedSkhSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.TransferFrom(&_ErcWrappedSkh.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ErcWrappedSkh *ErcWrappedSkhTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.TransferFrom(&_ErcWrappedSkh.TransactOpts, from, to, value)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ErcWrappedSkh *ErcWrappedSkhTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErcWrappedSkh.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ErcWrappedSkh *ErcWrappedSkhSession) Unpause() (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.Unpause(&_ErcWrappedSkh.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ErcWrappedSkh *ErcWrappedSkhTransactorSession) Unpause() (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.Unpause(&_ErcWrappedSkh.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns(uint256)
func (_ErcWrappedSkh *ErcWrappedSkhTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ErcWrappedSkh.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns(uint256)
func (_ErcWrappedSkh *ErcWrappedSkhSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.Withdraw(&_ErcWrappedSkh.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns(uint256)
func (_ErcWrappedSkh *ErcWrappedSkhTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _ErcWrappedSkh.Contract.Withdraw(&_ErcWrappedSkh.TransactOpts, amount)
}

// ErcWrappedSkhApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ErcWrappedSkh contract.
type ErcWrappedSkhApprovalIterator struct {
	Event *ErcWrappedSkhApproval // Event containing the contract specifics and raw log

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
func (it *ErcWrappedSkhApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedSkhApproval)
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
		it.Event = new(ErcWrappedSkhApproval)
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
func (it *ErcWrappedSkhApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedSkhApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedSkhApproval represents a Approval event raised by the ErcWrappedSkh contract.
type ErcWrappedSkhApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ErcWrappedSkh *ErcWrappedSkhFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ErcWrappedSkhApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ErcWrappedSkh.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedSkhApprovalIterator{contract: _ErcWrappedSkh.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ErcWrappedSkh *ErcWrappedSkhFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ErcWrappedSkhApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ErcWrappedSkh.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedSkhApproval)
				if err := _ErcWrappedSkh.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ErcWrappedSkh *ErcWrappedSkhFilterer) ParseApproval(log types.Log) (*ErcWrappedSkhApproval, error) {
	event := new(ErcWrappedSkhApproval)
	if err := _ErcWrappedSkh.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErcWrappedSkhPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ErcWrappedSkh contract.
type ErcWrappedSkhPausedIterator struct {
	Event *ErcWrappedSkhPaused // Event containing the contract specifics and raw log

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
func (it *ErcWrappedSkhPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedSkhPaused)
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
		it.Event = new(ErcWrappedSkhPaused)
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
func (it *ErcWrappedSkhPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedSkhPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedSkhPaused represents a Paused event raised by the ErcWrappedSkh contract.
type ErcWrappedSkhPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ErcWrappedSkh *ErcWrappedSkhFilterer) FilterPaused(opts *bind.FilterOpts) (*ErcWrappedSkhPausedIterator, error) {

	logs, sub, err := _ErcWrappedSkh.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ErcWrappedSkhPausedIterator{contract: _ErcWrappedSkh.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ErcWrappedSkh *ErcWrappedSkhFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ErcWrappedSkhPaused) (event.Subscription, error) {

	logs, sub, err := _ErcWrappedSkh.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedSkhPaused)
				if err := _ErcWrappedSkh.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ErcWrappedSkh *ErcWrappedSkhFilterer) ParsePaused(log types.Log) (*ErcWrappedSkhPaused, error) {
	event := new(ErcWrappedSkhPaused)
	if err := _ErcWrappedSkh.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErcWrappedSkhPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the ErcWrappedSkh contract.
type ErcWrappedSkhPauserAddedIterator struct {
	Event *ErcWrappedSkhPauserAdded // Event containing the contract specifics and raw log

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
func (it *ErcWrappedSkhPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedSkhPauserAdded)
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
		it.Event = new(ErcWrappedSkhPauserAdded)
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
func (it *ErcWrappedSkhPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedSkhPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedSkhPauserAdded represents a PauserAdded event raised by the ErcWrappedSkh contract.
type ErcWrappedSkhPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_ErcWrappedSkh *ErcWrappedSkhFilterer) FilterPauserAdded(opts *bind.FilterOpts, account []common.Address) (*ErcWrappedSkhPauserAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ErcWrappedSkh.contract.FilterLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedSkhPauserAddedIterator{contract: _ErcWrappedSkh.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_ErcWrappedSkh *ErcWrappedSkhFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *ErcWrappedSkhPauserAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ErcWrappedSkh.contract.WatchLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedSkhPauserAdded)
				if err := _ErcWrappedSkh.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// ParsePauserAdded is a log parse operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_ErcWrappedSkh *ErcWrappedSkhFilterer) ParsePauserAdded(log types.Log) (*ErcWrappedSkhPauserAdded, error) {
	event := new(ErcWrappedSkhPauserAdded)
	if err := _ErcWrappedSkh.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErcWrappedSkhPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the ErcWrappedSkh contract.
type ErcWrappedSkhPauserRemovedIterator struct {
	Event *ErcWrappedSkhPauserRemoved // Event containing the contract specifics and raw log

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
func (it *ErcWrappedSkhPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedSkhPauserRemoved)
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
		it.Event = new(ErcWrappedSkhPauserRemoved)
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
func (it *ErcWrappedSkhPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedSkhPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedSkhPauserRemoved represents a PauserRemoved event raised by the ErcWrappedSkh contract.
type ErcWrappedSkhPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_ErcWrappedSkh *ErcWrappedSkhFilterer) FilterPauserRemoved(opts *bind.FilterOpts, account []common.Address) (*ErcWrappedSkhPauserRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ErcWrappedSkh.contract.FilterLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedSkhPauserRemovedIterator{contract: _ErcWrappedSkh.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_ErcWrappedSkh *ErcWrappedSkhFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *ErcWrappedSkhPauserRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ErcWrappedSkh.contract.WatchLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedSkhPauserRemoved)
				if err := _ErcWrappedSkh.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// ParsePauserRemoved is a log parse operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_ErcWrappedSkh *ErcWrappedSkhFilterer) ParsePauserRemoved(log types.Log) (*ErcWrappedSkhPauserRemoved, error) {
	event := new(ErcWrappedSkhPauserRemoved)
	if err := _ErcWrappedSkh.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErcWrappedSkhTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ErcWrappedSkh contract.
type ErcWrappedSkhTransferIterator struct {
	Event *ErcWrappedSkhTransfer // Event containing the contract specifics and raw log

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
func (it *ErcWrappedSkhTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedSkhTransfer)
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
		it.Event = new(ErcWrappedSkhTransfer)
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
func (it *ErcWrappedSkhTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedSkhTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedSkhTransfer represents a Transfer event raised by the ErcWrappedSkh contract.
type ErcWrappedSkhTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ErcWrappedSkh *ErcWrappedSkhFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ErcWrappedSkhTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ErcWrappedSkh.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ErcWrappedSkhTransferIterator{contract: _ErcWrappedSkh.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ErcWrappedSkh *ErcWrappedSkhFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ErcWrappedSkhTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ErcWrappedSkh.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedSkhTransfer)
				if err := _ErcWrappedSkh.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ErcWrappedSkh *ErcWrappedSkhFilterer) ParseTransfer(log types.Log) (*ErcWrappedSkhTransfer, error) {
	event := new(ErcWrappedSkhTransfer)
	if err := _ErcWrappedSkh.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ErcWrappedSkhUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ErcWrappedSkh contract.
type ErcWrappedSkhUnpausedIterator struct {
	Event *ErcWrappedSkhUnpaused // Event containing the contract specifics and raw log

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
func (it *ErcWrappedSkhUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ErcWrappedSkhUnpaused)
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
		it.Event = new(ErcWrappedSkhUnpaused)
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
func (it *ErcWrappedSkhUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ErcWrappedSkhUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ErcWrappedSkhUnpaused represents a Unpaused event raised by the ErcWrappedSkh contract.
type ErcWrappedSkhUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ErcWrappedSkh *ErcWrappedSkhFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ErcWrappedSkhUnpausedIterator, error) {

	logs, sub, err := _ErcWrappedSkh.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ErcWrappedSkhUnpausedIterator{contract: _ErcWrappedSkh.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ErcWrappedSkh *ErcWrappedSkhFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ErcWrappedSkhUnpaused) (event.Subscription, error) {

	logs, sub, err := _ErcWrappedSkh.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ErcWrappedSkhUnpaused)
				if err := _ErcWrappedSkh.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ErcWrappedSkh *ErcWrappedSkhFilterer) ParseUnpaused(log types.Log) (*ErcWrappedSkhUnpaused, error) {
	event := new(ErcWrappedSkhUnpaused)
	if err := _ErcWrappedSkh.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
