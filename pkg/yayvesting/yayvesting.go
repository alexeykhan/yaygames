// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yayvesting

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

// YayvestingMetaData contains all meta data concerning the Yayvesting contract.
var YayvestingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_mercleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_tgeTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"category\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"resultReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"step\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"StepClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TgeClaim\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"alreadyRewarded\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumYayVesting.CategoryNames\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"categories\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalSteps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stepTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"percentBefore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"percentAfter\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_category\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"checkClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_category\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_claimResult\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdrawal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastClaimedStep\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mercleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tgeIsClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tgeTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// YayvestingABI is the input ABI used to generate the binding from.
// Deprecated: Use YayvestingMetaData.ABI instead.
var YayvestingABI = YayvestingMetaData.ABI

// Yayvesting is an auto generated Go binding around an Ethereum contract.
type Yayvesting struct {
	YayvestingCaller     // Read-only binding to the contract
	YayvestingTransactor // Write-only binding to the contract
	YayvestingFilterer   // Log filterer for contract events
}

// YayvestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type YayvestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YayvestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YayvestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YayvestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YayvestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YayvestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YayvestingSession struct {
	Contract     *Yayvesting       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YayvestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YayvestingCallerSession struct {
	Contract *YayvestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// YayvestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YayvestingTransactorSession struct {
	Contract     *YayvestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// YayvestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type YayvestingRaw struct {
	Contract *Yayvesting // Generic contract binding to access the raw methods on
}

// YayvestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YayvestingCallerRaw struct {
	Contract *YayvestingCaller // Generic read-only contract binding to access the raw methods on
}

// YayvestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YayvestingTransactorRaw struct {
	Contract *YayvestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYayvesting creates a new instance of Yayvesting, bound to a specific deployed contract.
func NewYayvesting(address common.Address, backend bind.ContractBackend) (*Yayvesting, error) {
	contract, err := bindYayvesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Yayvesting{YayvestingCaller: YayvestingCaller{contract: contract}, YayvestingTransactor: YayvestingTransactor{contract: contract}, YayvestingFilterer: YayvestingFilterer{contract: contract}}, nil
}

// NewYayvestingCaller creates a new read-only instance of Yayvesting, bound to a specific deployed contract.
func NewYayvestingCaller(address common.Address, caller bind.ContractCaller) (*YayvestingCaller, error) {
	contract, err := bindYayvesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YayvestingCaller{contract: contract}, nil
}

// NewYayvestingTransactor creates a new write-only instance of Yayvesting, bound to a specific deployed contract.
func NewYayvestingTransactor(address common.Address, transactor bind.ContractTransactor) (*YayvestingTransactor, error) {
	contract, err := bindYayvesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YayvestingTransactor{contract: contract}, nil
}

// NewYayvestingFilterer creates a new log filterer instance of Yayvesting, bound to a specific deployed contract.
func NewYayvestingFilterer(address common.Address, filterer bind.ContractFilterer) (*YayvestingFilterer, error) {
	contract, err := bindYayvesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YayvestingFilterer{contract: contract}, nil
}

// bindYayvesting binds a generic wrapper to an already deployed contract.
func bindYayvesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YayvestingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yayvesting *YayvestingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yayvesting.Contract.YayvestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yayvesting *YayvestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yayvesting.Contract.YayvestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yayvesting *YayvestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yayvesting.Contract.YayvestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yayvesting *YayvestingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yayvesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yayvesting *YayvestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yayvesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yayvesting *YayvestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yayvesting.Contract.contract.Transact(opts, method, params...)
}

// AlreadyRewarded is a free data retrieval call binding the contract method 0x3ae4fb1b.
//
// Solidity: function alreadyRewarded(address ) view returns(uint256)
func (_Yayvesting *YayvestingCaller) AlreadyRewarded(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yayvesting.contract.Call(opts, &out, "alreadyRewarded", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AlreadyRewarded is a free data retrieval call binding the contract method 0x3ae4fb1b.
//
// Solidity: function alreadyRewarded(address ) view returns(uint256)
func (_Yayvesting *YayvestingSession) AlreadyRewarded(arg0 common.Address) (*big.Int, error) {
	return _Yayvesting.Contract.AlreadyRewarded(&_Yayvesting.CallOpts, arg0)
}

// AlreadyRewarded is a free data retrieval call binding the contract method 0x3ae4fb1b.
//
// Solidity: function alreadyRewarded(address ) view returns(uint256)
func (_Yayvesting *YayvestingCallerSession) AlreadyRewarded(arg0 common.Address) (*big.Int, error) {
	return _Yayvesting.Contract.AlreadyRewarded(&_Yayvesting.CallOpts, arg0)
}

// Categories is a free data retrieval call binding the contract method 0x7be6277b.
//
// Solidity: function categories(uint8 ) view returns(uint256 totalSteps, uint256 stepTime, uint256 percentBefore, uint256 percentAfter)
func (_Yayvesting *YayvestingCaller) Categories(opts *bind.CallOpts, arg0 uint8) (struct {
	TotalSteps    *big.Int
	StepTime      *big.Int
	PercentBefore *big.Int
	PercentAfter  *big.Int
}, error) {
	var out []interface{}
	err := _Yayvesting.contract.Call(opts, &out, "categories", arg0)

	outstruct := new(struct {
		TotalSteps    *big.Int
		StepTime      *big.Int
		PercentBefore *big.Int
		PercentAfter  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalSteps = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StepTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PercentBefore = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PercentAfter = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Categories is a free data retrieval call binding the contract method 0x7be6277b.
//
// Solidity: function categories(uint8 ) view returns(uint256 totalSteps, uint256 stepTime, uint256 percentBefore, uint256 percentAfter)
func (_Yayvesting *YayvestingSession) Categories(arg0 uint8) (struct {
	TotalSteps    *big.Int
	StepTime      *big.Int
	PercentBefore *big.Int
	PercentAfter  *big.Int
}, error) {
	return _Yayvesting.Contract.Categories(&_Yayvesting.CallOpts, arg0)
}

// Categories is a free data retrieval call binding the contract method 0x7be6277b.
//
// Solidity: function categories(uint8 ) view returns(uint256 totalSteps, uint256 stepTime, uint256 percentBefore, uint256 percentAfter)
func (_Yayvesting *YayvestingCallerSession) Categories(arg0 uint8) (struct {
	TotalSteps    *big.Int
	StepTime      *big.Int
	PercentBefore *big.Int
	PercentAfter  *big.Int
}, error) {
	return _Yayvesting.Contract.Categories(&_Yayvesting.CallOpts, arg0)
}

// CheckClaim is a free data retrieval call binding the contract method 0x359ebd9f.
//
// Solidity: function checkClaim(address _target, uint256 _category, uint256 _amount, bytes32[] _merkleProof) view returns(bool)
func (_Yayvesting *YayvestingCaller) CheckClaim(opts *bind.CallOpts, _target common.Address, _category *big.Int, _amount *big.Int, _merkleProof [][32]byte) (bool, error) {
	var out []interface{}
	err := _Yayvesting.contract.Call(opts, &out, "checkClaim", _target, _category, _amount, _merkleProof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckClaim is a free data retrieval call binding the contract method 0x359ebd9f.
//
// Solidity: function checkClaim(address _target, uint256 _category, uint256 _amount, bytes32[] _merkleProof) view returns(bool)
func (_Yayvesting *YayvestingSession) CheckClaim(_target common.Address, _category *big.Int, _amount *big.Int, _merkleProof [][32]byte) (bool, error) {
	return _Yayvesting.Contract.CheckClaim(&_Yayvesting.CallOpts, _target, _category, _amount, _merkleProof)
}

// CheckClaim is a free data retrieval call binding the contract method 0x359ebd9f.
//
// Solidity: function checkClaim(address _target, uint256 _category, uint256 _amount, bytes32[] _merkleProof) view returns(bool)
func (_Yayvesting *YayvestingCallerSession) CheckClaim(_target common.Address, _category *big.Int, _amount *big.Int, _merkleProof [][32]byte) (bool, error) {
	return _Yayvesting.Contract.CheckClaim(&_Yayvesting.CallOpts, _target, _category, _amount, _merkleProof)
}

// LastClaimedStep is a free data retrieval call binding the contract method 0x8d41ebe0.
//
// Solidity: function lastClaimedStep(address ) view returns(uint256)
func (_Yayvesting *YayvestingCaller) LastClaimedStep(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yayvesting.contract.Call(opts, &out, "lastClaimedStep", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastClaimedStep is a free data retrieval call binding the contract method 0x8d41ebe0.
//
// Solidity: function lastClaimedStep(address ) view returns(uint256)
func (_Yayvesting *YayvestingSession) LastClaimedStep(arg0 common.Address) (*big.Int, error) {
	return _Yayvesting.Contract.LastClaimedStep(&_Yayvesting.CallOpts, arg0)
}

// LastClaimedStep is a free data retrieval call binding the contract method 0x8d41ebe0.
//
// Solidity: function lastClaimedStep(address ) view returns(uint256)
func (_Yayvesting *YayvestingCallerSession) LastClaimedStep(arg0 common.Address) (*big.Int, error) {
	return _Yayvesting.Contract.LastClaimedStep(&_Yayvesting.CallOpts, arg0)
}

// MercleRoot is a free data retrieval call binding the contract method 0xd060d6fe.
//
// Solidity: function mercleRoot() view returns(bytes32)
func (_Yayvesting *YayvestingCaller) MercleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Yayvesting.contract.Call(opts, &out, "mercleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MercleRoot is a free data retrieval call binding the contract method 0xd060d6fe.
//
// Solidity: function mercleRoot() view returns(bytes32)
func (_Yayvesting *YayvestingSession) MercleRoot() ([32]byte, error) {
	return _Yayvesting.Contract.MercleRoot(&_Yayvesting.CallOpts)
}

// MercleRoot is a free data retrieval call binding the contract method 0xd060d6fe.
//
// Solidity: function mercleRoot() view returns(bytes32)
func (_Yayvesting *YayvestingCallerSession) MercleRoot() ([32]byte, error) {
	return _Yayvesting.Contract.MercleRoot(&_Yayvesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Yayvesting *YayvestingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yayvesting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Yayvesting *YayvestingSession) Owner() (common.Address, error) {
	return _Yayvesting.Contract.Owner(&_Yayvesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Yayvesting *YayvestingCallerSession) Owner() (common.Address, error) {
	return _Yayvesting.Contract.Owner(&_Yayvesting.CallOpts)
}

// TgeIsClaimed is a free data retrieval call binding the contract method 0xe989a9b5.
//
// Solidity: function tgeIsClaimed(address ) view returns(bool)
func (_Yayvesting *YayvestingCaller) TgeIsClaimed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Yayvesting.contract.Call(opts, &out, "tgeIsClaimed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TgeIsClaimed is a free data retrieval call binding the contract method 0xe989a9b5.
//
// Solidity: function tgeIsClaimed(address ) view returns(bool)
func (_Yayvesting *YayvestingSession) TgeIsClaimed(arg0 common.Address) (bool, error) {
	return _Yayvesting.Contract.TgeIsClaimed(&_Yayvesting.CallOpts, arg0)
}

// TgeIsClaimed is a free data retrieval call binding the contract method 0xe989a9b5.
//
// Solidity: function tgeIsClaimed(address ) view returns(bool)
func (_Yayvesting *YayvestingCallerSession) TgeIsClaimed(arg0 common.Address) (bool, error) {
	return _Yayvesting.Contract.TgeIsClaimed(&_Yayvesting.CallOpts, arg0)
}

// TgeTimestamp is a free data retrieval call binding the contract method 0xa4317ef4.
//
// Solidity: function tgeTimestamp() view returns(uint256)
func (_Yayvesting *YayvestingCaller) TgeTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yayvesting.contract.Call(opts, &out, "tgeTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TgeTimestamp is a free data retrieval call binding the contract method 0xa4317ef4.
//
// Solidity: function tgeTimestamp() view returns(uint256)
func (_Yayvesting *YayvestingSession) TgeTimestamp() (*big.Int, error) {
	return _Yayvesting.Contract.TgeTimestamp(&_Yayvesting.CallOpts)
}

// TgeTimestamp is a free data retrieval call binding the contract method 0xa4317ef4.
//
// Solidity: function tgeTimestamp() view returns(uint256)
func (_Yayvesting *YayvestingCallerSession) TgeTimestamp() (*big.Int, error) {
	return _Yayvesting.Contract.TgeTimestamp(&_Yayvesting.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Yayvesting *YayvestingCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yayvesting.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Yayvesting *YayvestingSession) Token() (common.Address, error) {
	return _Yayvesting.Contract.Token(&_Yayvesting.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Yayvesting *YayvestingCallerSession) Token() (common.Address, error) {
	return _Yayvesting.Contract.Token(&_Yayvesting.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0xae0b51df.
//
// Solidity: function claim(uint256 _category, uint256 _amount, bytes32[] _merkleProof) returns(uint256 _claimResult)
func (_Yayvesting *YayvestingTransactor) Claim(opts *bind.TransactOpts, _category *big.Int, _amount *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _Yayvesting.contract.Transact(opts, "claim", _category, _amount, _merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0xae0b51df.
//
// Solidity: function claim(uint256 _category, uint256 _amount, bytes32[] _merkleProof) returns(uint256 _claimResult)
func (_Yayvesting *YayvestingSession) Claim(_category *big.Int, _amount *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _Yayvesting.Contract.Claim(&_Yayvesting.TransactOpts, _category, _amount, _merkleProof)
}

// Claim is a paid mutator transaction binding the contract method 0xae0b51df.
//
// Solidity: function claim(uint256 _category, uint256 _amount, bytes32[] _merkleProof) returns(uint256 _claimResult)
func (_Yayvesting *YayvestingTransactorSession) Claim(_category *big.Int, _amount *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _Yayvesting.Contract.Claim(&_Yayvesting.TransactOpts, _category, _amount, _merkleProof)
}

// EmergencyWithdrawal is a paid mutator transaction binding the contract method 0x0039522c.
//
// Solidity: function emergencyWithdrawal(uint256 amount) returns(bool)
func (_Yayvesting *YayvestingTransactor) EmergencyWithdrawal(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Yayvesting.contract.Transact(opts, "emergencyWithdrawal", amount)
}

// EmergencyWithdrawal is a paid mutator transaction binding the contract method 0x0039522c.
//
// Solidity: function emergencyWithdrawal(uint256 amount) returns(bool)
func (_Yayvesting *YayvestingSession) EmergencyWithdrawal(amount *big.Int) (*types.Transaction, error) {
	return _Yayvesting.Contract.EmergencyWithdrawal(&_Yayvesting.TransactOpts, amount)
}

// EmergencyWithdrawal is a paid mutator transaction binding the contract method 0x0039522c.
//
// Solidity: function emergencyWithdrawal(uint256 amount) returns(bool)
func (_Yayvesting *YayvestingTransactorSession) EmergencyWithdrawal(amount *big.Int) (*types.Transaction, error) {
	return _Yayvesting.Contract.EmergencyWithdrawal(&_Yayvesting.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Yayvesting *YayvestingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yayvesting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Yayvesting *YayvestingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Yayvesting.Contract.RenounceOwnership(&_Yayvesting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Yayvesting *YayvestingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Yayvesting.Contract.RenounceOwnership(&_Yayvesting.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Yayvesting *YayvestingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Yayvesting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Yayvesting *YayvestingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Yayvesting.Contract.TransferOwnership(&_Yayvesting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Yayvesting *YayvestingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Yayvesting.Contract.TransferOwnership(&_Yayvesting.TransactOpts, newOwner)
}

// YayvestingClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Yayvesting contract.
type YayvestingClaimIterator struct {
	Event *YayvestingClaim // Event containing the contract specifics and raw log

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
func (it *YayvestingClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YayvestingClaim)
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
		it.Event = new(YayvestingClaim)
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
func (it *YayvestingClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YayvestingClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YayvestingClaim represents a Claim event raised by the Yayvesting contract.
type YayvestingClaim struct {
	Target       common.Address
	Category     *big.Int
	Amount       *big.Int
	MerkleProof  [][32]byte
	ResultReward *big.Int
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0xed397d6389ffb93b4028067de164a2b9ad3612bf5d4937c1bbf1a2e0fe7c2e91.
//
// Solidity: event Claim(address indexed target, uint256 indexed category, uint256 amount, bytes32[] merkleProof, uint256 resultReward, uint256 timestamp)
func (_Yayvesting *YayvestingFilterer) FilterClaim(opts *bind.FilterOpts, target []common.Address, category []*big.Int) (*YayvestingClaimIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var categoryRule []interface{}
	for _, categoryItem := range category {
		categoryRule = append(categoryRule, categoryItem)
	}

	logs, sub, err := _Yayvesting.contract.FilterLogs(opts, "Claim", targetRule, categoryRule)
	if err != nil {
		return nil, err
	}
	return &YayvestingClaimIterator{contract: _Yayvesting.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0xed397d6389ffb93b4028067de164a2b9ad3612bf5d4937c1bbf1a2e0fe7c2e91.
//
// Solidity: event Claim(address indexed target, uint256 indexed category, uint256 amount, bytes32[] merkleProof, uint256 resultReward, uint256 timestamp)
func (_Yayvesting *YayvestingFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *YayvestingClaim, target []common.Address, category []*big.Int) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var categoryRule []interface{}
	for _, categoryItem := range category {
		categoryRule = append(categoryRule, categoryItem)
	}

	logs, sub, err := _Yayvesting.contract.WatchLogs(opts, "Claim", targetRule, categoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YayvestingClaim)
				if err := _Yayvesting.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0xed397d6389ffb93b4028067de164a2b9ad3612bf5d4937c1bbf1a2e0fe7c2e91.
//
// Solidity: event Claim(address indexed target, uint256 indexed category, uint256 amount, bytes32[] merkleProof, uint256 resultReward, uint256 timestamp)
func (_Yayvesting *YayvestingFilterer) ParseClaim(log types.Log) (*YayvestingClaim, error) {
	event := new(YayvestingClaim)
	if err := _Yayvesting.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YayvestingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Yayvesting contract.
type YayvestingOwnershipTransferredIterator struct {
	Event *YayvestingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *YayvestingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YayvestingOwnershipTransferred)
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
		it.Event = new(YayvestingOwnershipTransferred)
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
func (it *YayvestingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YayvestingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YayvestingOwnershipTransferred represents a OwnershipTransferred event raised by the Yayvesting contract.
type YayvestingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Yayvesting *YayvestingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*YayvestingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Yayvesting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &YayvestingOwnershipTransferredIterator{contract: _Yayvesting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Yayvesting *YayvestingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *YayvestingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Yayvesting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YayvestingOwnershipTransferred)
				if err := _Yayvesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Yayvesting *YayvestingFilterer) ParseOwnershipTransferred(log types.Log) (*YayvestingOwnershipTransferred, error) {
	event := new(YayvestingOwnershipTransferred)
	if err := _Yayvesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YayvestingStepClaimIterator is returned from FilterStepClaim and is used to iterate over the raw logs and unpacked data for StepClaim events raised by the Yayvesting contract.
type YayvestingStepClaimIterator struct {
	Event *YayvestingStepClaim // Event containing the contract specifics and raw log

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
func (it *YayvestingStepClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YayvestingStepClaim)
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
		it.Event = new(YayvestingStepClaim)
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
func (it *YayvestingStepClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YayvestingStepClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YayvestingStepClaim represents a StepClaim event raised by the Yayvesting contract.
type YayvestingStepClaim struct {
	Target    common.Address
	Step      *big.Int
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStepClaim is a free log retrieval operation binding the contract event 0xb5fedb6d86259eba2fd9023651406f62391c1dbe56a7b4f6922924b299c9d2a5.
//
// Solidity: event StepClaim(address indexed target, uint256 indexed step, uint256 value, uint256 timestamp)
func (_Yayvesting *YayvestingFilterer) FilterStepClaim(opts *bind.FilterOpts, target []common.Address, step []*big.Int) (*YayvestingStepClaimIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var stepRule []interface{}
	for _, stepItem := range step {
		stepRule = append(stepRule, stepItem)
	}

	logs, sub, err := _Yayvesting.contract.FilterLogs(opts, "StepClaim", targetRule, stepRule)
	if err != nil {
		return nil, err
	}
	return &YayvestingStepClaimIterator{contract: _Yayvesting.contract, event: "StepClaim", logs: logs, sub: sub}, nil
}

// WatchStepClaim is a free log subscription operation binding the contract event 0xb5fedb6d86259eba2fd9023651406f62391c1dbe56a7b4f6922924b299c9d2a5.
//
// Solidity: event StepClaim(address indexed target, uint256 indexed step, uint256 value, uint256 timestamp)
func (_Yayvesting *YayvestingFilterer) WatchStepClaim(opts *bind.WatchOpts, sink chan<- *YayvestingStepClaim, target []common.Address, step []*big.Int) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var stepRule []interface{}
	for _, stepItem := range step {
		stepRule = append(stepRule, stepItem)
	}

	logs, sub, err := _Yayvesting.contract.WatchLogs(opts, "StepClaim", targetRule, stepRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YayvestingStepClaim)
				if err := _Yayvesting.contract.UnpackLog(event, "StepClaim", log); err != nil {
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

// ParseStepClaim is a log parse operation binding the contract event 0xb5fedb6d86259eba2fd9023651406f62391c1dbe56a7b4f6922924b299c9d2a5.
//
// Solidity: event StepClaim(address indexed target, uint256 indexed step, uint256 value, uint256 timestamp)
func (_Yayvesting *YayvestingFilterer) ParseStepClaim(log types.Log) (*YayvestingStepClaim, error) {
	event := new(YayvestingStepClaim)
	if err := _Yayvesting.contract.UnpackLog(event, "StepClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YayvestingTgeClaimIterator is returned from FilterTgeClaim and is used to iterate over the raw logs and unpacked data for TgeClaim events raised by the Yayvesting contract.
type YayvestingTgeClaimIterator struct {
	Event *YayvestingTgeClaim // Event containing the contract specifics and raw log

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
func (it *YayvestingTgeClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YayvestingTgeClaim)
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
		it.Event = new(YayvestingTgeClaim)
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
func (it *YayvestingTgeClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YayvestingTgeClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YayvestingTgeClaim represents a TgeClaim event raised by the Yayvesting contract.
type YayvestingTgeClaim struct {
	Target    common.Address
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTgeClaim is a free log retrieval operation binding the contract event 0x8166115a145a1fe0553bf0f9581055c766cccb00c1f73f661870f4a0f41b2fe9.
//
// Solidity: event TgeClaim(address indexed target, uint256 value, uint256 timestamp)
func (_Yayvesting *YayvestingFilterer) FilterTgeClaim(opts *bind.FilterOpts, target []common.Address) (*YayvestingTgeClaimIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Yayvesting.contract.FilterLogs(opts, "TgeClaim", targetRule)
	if err != nil {
		return nil, err
	}
	return &YayvestingTgeClaimIterator{contract: _Yayvesting.contract, event: "TgeClaim", logs: logs, sub: sub}, nil
}

// WatchTgeClaim is a free log subscription operation binding the contract event 0x8166115a145a1fe0553bf0f9581055c766cccb00c1f73f661870f4a0f41b2fe9.
//
// Solidity: event TgeClaim(address indexed target, uint256 value, uint256 timestamp)
func (_Yayvesting *YayvestingFilterer) WatchTgeClaim(opts *bind.WatchOpts, sink chan<- *YayvestingTgeClaim, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Yayvesting.contract.WatchLogs(opts, "TgeClaim", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YayvestingTgeClaim)
				if err := _Yayvesting.contract.UnpackLog(event, "TgeClaim", log); err != nil {
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

// ParseTgeClaim is a log parse operation binding the contract event 0x8166115a145a1fe0553bf0f9581055c766cccb00c1f73f661870f4a0f41b2fe9.
//
// Solidity: event TgeClaim(address indexed target, uint256 value, uint256 timestamp)
func (_Yayvesting *YayvestingFilterer) ParseTgeClaim(log types.Log) (*YayvestingTgeClaim, error) {
	event := new(YayvestingTgeClaim)
	if err := _Yayvesting.contract.UnpackLog(event, "TgeClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
