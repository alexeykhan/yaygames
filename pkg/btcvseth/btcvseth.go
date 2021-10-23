// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package btcvseth

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

// BtcvsethMetaData contains all meta data concerning the Btcvseth contract.
var BtcvsethMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"type\":\"address\",\"name\":\"_oracleFirst\",\"internalType\":\"contractAggregatorV3Interface\"},{\"type\":\"address\",\"name\":\"_oracleSecond\",\"internalType\":\"contractAggregatorV3Interface\"},{\"type\":\"address\",\"name\":\"_adminAddress\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_operatorAddress\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_intervalBlocks\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_bufferBlocks\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_minBetAmount\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_oracleUpdateAllowance\",\"internalType\":\"uint256\"}]},{\"type\":\"event\",\"name\":\"BetBear\",\"inputs\":[{\"type\":\"address\",\"name\":\"sender\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"currentEpoch\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BetBull\",\"inputs\":[{\"type\":\"address\",\"name\":\"sender\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"currentEpoch\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Claim\",\"inputs\":[{\"type\":\"address\",\"name\":\"sender\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"currentEpoch\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClaimTreasury\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EndRound\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"epoch\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"blockNumber\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"int256\",\"name\":\"firstPrice\",\"internalType\":\"int256\",\"indexed\":false},{\"type\":\"int256\",\"name\":\"secondPrice\",\"internalType\":\"int256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockRound\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"epoch\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"blockNumber\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"int256\",\"name\":\"firstPrice\",\"internalType\":\"int256\",\"indexed\":false},{\"type\":\"int256\",\"name\":\"secondPrice\",\"internalType\":\"int256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinBetAmountUpdated\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"epoch\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"minBetAmount\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Pause\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"epoch\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RatesUpdated\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"epoch\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"rewardRate\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"treasuryRate\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsCalculated\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"epoch\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"rewardBaseCalAmount\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"rewardAmount\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"treasuryAmount\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"int256\",\"name\":\"firstPercent\",\"internalType\":\"int256\",\"indexed\":false},{\"type\":\"int256\",\"name\":\"secondPercent\",\"internalType\":\"int256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StartRound\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"epoch\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"blockNumber\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpause\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"epoch\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"TOTAL_RATE\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"adminAddress\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"betBear\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"betBull\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"bufferBlocks\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"claim\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"epoch\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"claimTreasury\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"claimable\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"epoch\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"user\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"currentEpoch\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"executeRound\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"genesisLockOnce\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"genesisLockRound\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"genesisStartOnce\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"genesisStartRound\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256[]\",\"name\":\"\",\"internalType\":\"uint256[]\"},{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getUserRounds\",\"inputs\":[{\"type\":\"address\",\"name\":\"user\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"cursor\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"size\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"intervalBlocks\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"position\",\"internalType\":\"enumBtcVsEth.Position\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"},{\"type\":\"bool\",\"name\":\"claimed\",\"internalType\":\"bool\"}],\"name\":\"ledger\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"minBetAmount\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"operatorAddress\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"oracleFirstLatestRoundId\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"oracleSecondLatestRoundId\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"oracleUpdateAllowance\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"owner\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"pause\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"paused\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"refundable\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"epoch\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"user\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"renounceOwnership\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"rewardRate\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"epoch\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"startBlock\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"lockBlock\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"endBlock\",\"internalType\":\"uint256\"},{\"type\":\"int256\",\"name\":\"lockFirstPrice\",\"internalType\":\"int256\"},{\"type\":\"int256\",\"name\":\"closeFirstPrice\",\"internalType\":\"int256\"},{\"type\":\"int256\",\"name\":\"lockSecondPrice\",\"internalType\":\"int256\"},{\"type\":\"int256\",\"name\":\"closeSecondPrice\",\"internalType\":\"int256\"},{\"type\":\"uint256\",\"name\":\"totalAmount\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"bullAmount\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"bearAmount\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"rewardBaseCalAmount\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"rewardAmount\",\"internalType\":\"uint256\"},{\"type\":\"bool\",\"name\":\"oracleCalled\",\"internalType\":\"bool\"}],\"name\":\"rounds\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setAdmin\",\"inputs\":[{\"type\":\"address\",\"name\":\"_adminAddress\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setBufferBlocks\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_bufferBlocks\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setFirstOracle\",\"inputs\":[{\"type\":\"address\",\"name\":\"_oracle\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setIntervalBlocks\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_intervalBlocks\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setMinBetAmount\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_minBetAmount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setOperator\",\"inputs\":[{\"type\":\"address\",\"name\":\"_operatorAddress\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setOracleUpdateAllowance\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_oracleUpdateAllowance\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setRewardRate\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_rewardRate\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setSecondOracle\",\"inputs\":[{\"type\":\"address\",\"name\":\"_oracle\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setTreasuryRate\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_treasuryRate\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"transferOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"treasuryAmount\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"treasuryRate\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"unpause\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"userRounds\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}]}]",
}

// BtcvsethABI is the input ABI used to generate the binding from.
// Deprecated: Use BtcvsethMetaData.ABI instead.
var BtcvsethABI = BtcvsethMetaData.ABI

// Btcvseth is an auto generated Go binding around an Ethereum contract.
type Btcvseth struct {
	BtcvsethCaller     // Read-only binding to the contract
	BtcvsethTransactor // Write-only binding to the contract
	BtcvsethFilterer   // Log filterer for contract events
}

// BtcvsethCaller is an auto generated read-only Go binding around an Ethereum contract.
type BtcvsethCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BtcvsethTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BtcvsethTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BtcvsethFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BtcvsethFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BtcvsethSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BtcvsethSession struct {
	Contract     *Btcvseth         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BtcvsethCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BtcvsethCallerSession struct {
	Contract *BtcvsethCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BtcvsethTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BtcvsethTransactorSession struct {
	Contract     *BtcvsethTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BtcvsethRaw is an auto generated low-level Go binding around an Ethereum contract.
type BtcvsethRaw struct {
	Contract *Btcvseth // Generic contract binding to access the raw methods on
}

// BtcvsethCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BtcvsethCallerRaw struct {
	Contract *BtcvsethCaller // Generic read-only contract binding to access the raw methods on
}

// BtcvsethTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BtcvsethTransactorRaw struct {
	Contract *BtcvsethTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBtcvseth creates a new instance of Btcvseth, bound to a specific deployed contract.
func NewBtcvseth(address common.Address, backend bind.ContractBackend) (*Btcvseth, error) {
	contract, err := bindBtcvseth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Btcvseth{BtcvsethCaller: BtcvsethCaller{contract: contract}, BtcvsethTransactor: BtcvsethTransactor{contract: contract}, BtcvsethFilterer: BtcvsethFilterer{contract: contract}}, nil
}

// NewBtcvsethCaller creates a new read-only instance of Btcvseth, bound to a specific deployed contract.
func NewBtcvsethCaller(address common.Address, caller bind.ContractCaller) (*BtcvsethCaller, error) {
	contract, err := bindBtcvseth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BtcvsethCaller{contract: contract}, nil
}

// NewBtcvsethTransactor creates a new write-only instance of Btcvseth, bound to a specific deployed contract.
func NewBtcvsethTransactor(address common.Address, transactor bind.ContractTransactor) (*BtcvsethTransactor, error) {
	contract, err := bindBtcvseth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BtcvsethTransactor{contract: contract}, nil
}

// NewBtcvsethFilterer creates a new log filterer instance of Btcvseth, bound to a specific deployed contract.
func NewBtcvsethFilterer(address common.Address, filterer bind.ContractFilterer) (*BtcvsethFilterer, error) {
	contract, err := bindBtcvseth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BtcvsethFilterer{contract: contract}, nil
}

// bindBtcvseth binds a generic wrapper to an already deployed contract.
func bindBtcvseth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BtcvsethABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Btcvseth *BtcvsethRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Btcvseth.Contract.BtcvsethCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Btcvseth *BtcvsethRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Btcvseth.Contract.BtcvsethTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Btcvseth *BtcvsethRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Btcvseth.Contract.BtcvsethTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Btcvseth *BtcvsethCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Btcvseth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Btcvseth *BtcvsethTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Btcvseth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Btcvseth *BtcvsethTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Btcvseth.Contract.contract.Transact(opts, method, params...)
}

// TOTALRATE is a free data retrieval call binding the contract method 0xb29c299b.
//
// Solidity: function TOTAL_RATE() view returns(uint256)
func (_Btcvseth *BtcvsethCaller) TOTALRATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "TOTAL_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALRATE is a free data retrieval call binding the contract method 0xb29c299b.
//
// Solidity: function TOTAL_RATE() view returns(uint256)
func (_Btcvseth *BtcvsethSession) TOTALRATE() (*big.Int, error) {
	return _Btcvseth.Contract.TOTALRATE(&_Btcvseth.CallOpts)
}

// TOTALRATE is a free data retrieval call binding the contract method 0xb29c299b.
//
// Solidity: function TOTAL_RATE() view returns(uint256)
func (_Btcvseth *BtcvsethCallerSession) TOTALRATE() (*big.Int, error) {
	return _Btcvseth.Contract.TOTALRATE(&_Btcvseth.CallOpts)
}

// AdminAddress is a free data retrieval call binding the contract method 0xfc6f9468.
//
// Solidity: function adminAddress() view returns(address)
func (_Btcvseth *BtcvsethCaller) AdminAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "adminAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminAddress is a free data retrieval call binding the contract method 0xfc6f9468.
//
// Solidity: function adminAddress() view returns(address)
func (_Btcvseth *BtcvsethSession) AdminAddress() (common.Address, error) {
	return _Btcvseth.Contract.AdminAddress(&_Btcvseth.CallOpts)
}

// AdminAddress is a free data retrieval call binding the contract method 0xfc6f9468.
//
// Solidity: function adminAddress() view returns(address)
func (_Btcvseth *BtcvsethCallerSession) AdminAddress() (common.Address, error) {
	return _Btcvseth.Contract.AdminAddress(&_Btcvseth.CallOpts)
}

// BufferBlocks is a free data retrieval call binding the contract method 0x9780a752.
//
// Solidity: function bufferBlocks() view returns(uint256)
func (_Btcvseth *BtcvsethCaller) BufferBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "bufferBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BufferBlocks is a free data retrieval call binding the contract method 0x9780a752.
//
// Solidity: function bufferBlocks() view returns(uint256)
func (_Btcvseth *BtcvsethSession) BufferBlocks() (*big.Int, error) {
	return _Btcvseth.Contract.BufferBlocks(&_Btcvseth.CallOpts)
}

// BufferBlocks is a free data retrieval call binding the contract method 0x9780a752.
//
// Solidity: function bufferBlocks() view returns(uint256)
func (_Btcvseth *BtcvsethCallerSession) BufferBlocks() (*big.Int, error) {
	return _Btcvseth.Contract.BufferBlocks(&_Btcvseth.CallOpts)
}

// Claimable is a free data retrieval call binding the contract method 0xa0c7f71c.
//
// Solidity: function claimable(uint256 epoch, address user) view returns(bool)
func (_Btcvseth *BtcvsethCaller) Claimable(opts *bind.CallOpts, epoch *big.Int, user common.Address) (bool, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "claimable", epoch, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0xa0c7f71c.
//
// Solidity: function claimable(uint256 epoch, address user) view returns(bool)
func (_Btcvseth *BtcvsethSession) Claimable(epoch *big.Int, user common.Address) (bool, error) {
	return _Btcvseth.Contract.Claimable(&_Btcvseth.CallOpts, epoch, user)
}

// Claimable is a free data retrieval call binding the contract method 0xa0c7f71c.
//
// Solidity: function claimable(uint256 epoch, address user) view returns(bool)
func (_Btcvseth *BtcvsethCallerSession) Claimable(epoch *big.Int, user common.Address) (bool, error) {
	return _Btcvseth.Contract.Claimable(&_Btcvseth.CallOpts, epoch, user)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Btcvseth *BtcvsethCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Btcvseth *BtcvsethSession) CurrentEpoch() (*big.Int, error) {
	return _Btcvseth.Contract.CurrentEpoch(&_Btcvseth.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Btcvseth *BtcvsethCallerSession) CurrentEpoch() (*big.Int, error) {
	return _Btcvseth.Contract.CurrentEpoch(&_Btcvseth.CallOpts)
}

// GenesisLockOnce is a free data retrieval call binding the contract method 0x0f74174f.
//
// Solidity: function genesisLockOnce() view returns(bool)
func (_Btcvseth *BtcvsethCaller) GenesisLockOnce(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "genesisLockOnce")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GenesisLockOnce is a free data retrieval call binding the contract method 0x0f74174f.
//
// Solidity: function genesisLockOnce() view returns(bool)
func (_Btcvseth *BtcvsethSession) GenesisLockOnce() (bool, error) {
	return _Btcvseth.Contract.GenesisLockOnce(&_Btcvseth.CallOpts)
}

// GenesisLockOnce is a free data retrieval call binding the contract method 0x0f74174f.
//
// Solidity: function genesisLockOnce() view returns(bool)
func (_Btcvseth *BtcvsethCallerSession) GenesisLockOnce() (bool, error) {
	return _Btcvseth.Contract.GenesisLockOnce(&_Btcvseth.CallOpts)
}

// GenesisStartOnce is a free data retrieval call binding the contract method 0xf7fdec28.
//
// Solidity: function genesisStartOnce() view returns(bool)
func (_Btcvseth *BtcvsethCaller) GenesisStartOnce(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "genesisStartOnce")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GenesisStartOnce is a free data retrieval call binding the contract method 0xf7fdec28.
//
// Solidity: function genesisStartOnce() view returns(bool)
func (_Btcvseth *BtcvsethSession) GenesisStartOnce() (bool, error) {
	return _Btcvseth.Contract.GenesisStartOnce(&_Btcvseth.CallOpts)
}

// GenesisStartOnce is a free data retrieval call binding the contract method 0xf7fdec28.
//
// Solidity: function genesisStartOnce() view returns(bool)
func (_Btcvseth *BtcvsethCallerSession) GenesisStartOnce() (bool, error) {
	return _Btcvseth.Contract.GenesisStartOnce(&_Btcvseth.CallOpts)
}

// GetUserRounds is a free data retrieval call binding the contract method 0x951fd600.
//
// Solidity: function getUserRounds(address user, uint256 cursor, uint256 size) view returns(uint256[], uint256)
func (_Btcvseth *BtcvsethCaller) GetUserRounds(opts *bind.CallOpts, user common.Address, cursor *big.Int, size *big.Int) ([]*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "getUserRounds", user, cursor, size)

	if err != nil {
		return *new([]*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetUserRounds is a free data retrieval call binding the contract method 0x951fd600.
//
// Solidity: function getUserRounds(address user, uint256 cursor, uint256 size) view returns(uint256[], uint256)
func (_Btcvseth *BtcvsethSession) GetUserRounds(user common.Address, cursor *big.Int, size *big.Int) ([]*big.Int, *big.Int, error) {
	return _Btcvseth.Contract.GetUserRounds(&_Btcvseth.CallOpts, user, cursor, size)
}

// GetUserRounds is a free data retrieval call binding the contract method 0x951fd600.
//
// Solidity: function getUserRounds(address user, uint256 cursor, uint256 size) view returns(uint256[], uint256)
func (_Btcvseth *BtcvsethCallerSession) GetUserRounds(user common.Address, cursor *big.Int, size *big.Int) ([]*big.Int, *big.Int, error) {
	return _Btcvseth.Contract.GetUserRounds(&_Btcvseth.CallOpts, user, cursor, size)
}

// IntervalBlocks is a free data retrieval call binding the contract method 0x1ec9f34b.
//
// Solidity: function intervalBlocks() view returns(uint256)
func (_Btcvseth *BtcvsethCaller) IntervalBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "intervalBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntervalBlocks is a free data retrieval call binding the contract method 0x1ec9f34b.
//
// Solidity: function intervalBlocks() view returns(uint256)
func (_Btcvseth *BtcvsethSession) IntervalBlocks() (*big.Int, error) {
	return _Btcvseth.Contract.IntervalBlocks(&_Btcvseth.CallOpts)
}

// IntervalBlocks is a free data retrieval call binding the contract method 0x1ec9f34b.
//
// Solidity: function intervalBlocks() view returns(uint256)
func (_Btcvseth *BtcvsethCallerSession) IntervalBlocks() (*big.Int, error) {
	return _Btcvseth.Contract.IntervalBlocks(&_Btcvseth.CallOpts)
}

// Ledger is a free data retrieval call binding the contract method 0x7285c58b.
//
// Solidity: function ledger(uint256 , address ) view returns(uint8 position, uint256 amount, bool claimed)
func (_Btcvseth *BtcvsethCaller) Ledger(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Position uint8
	Amount   *big.Int
	Claimed  bool
}, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "ledger", arg0, arg1)

	outstruct := new(struct {
		Position uint8
		Amount   *big.Int
		Claimed  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Position = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Claimed = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Ledger is a free data retrieval call binding the contract method 0x7285c58b.
//
// Solidity: function ledger(uint256 , address ) view returns(uint8 position, uint256 amount, bool claimed)
func (_Btcvseth *BtcvsethSession) Ledger(arg0 *big.Int, arg1 common.Address) (struct {
	Position uint8
	Amount   *big.Int
	Claimed  bool
}, error) {
	return _Btcvseth.Contract.Ledger(&_Btcvseth.CallOpts, arg0, arg1)
}

// Ledger is a free data retrieval call binding the contract method 0x7285c58b.
//
// Solidity: function ledger(uint256 , address ) view returns(uint8 position, uint256 amount, bool claimed)
func (_Btcvseth *BtcvsethCallerSession) Ledger(arg0 *big.Int, arg1 common.Address) (struct {
	Position uint8
	Amount   *big.Int
	Claimed  bool
}, error) {
	return _Btcvseth.Contract.Ledger(&_Btcvseth.CallOpts, arg0, arg1)
}

// MinBetAmount is a free data retrieval call binding the contract method 0xfa968eea.
//
// Solidity: function minBetAmount() view returns(uint256)
func (_Btcvseth *BtcvsethCaller) MinBetAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "minBetAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBetAmount is a free data retrieval call binding the contract method 0xfa968eea.
//
// Solidity: function minBetAmount() view returns(uint256)
func (_Btcvseth *BtcvsethSession) MinBetAmount() (*big.Int, error) {
	return _Btcvseth.Contract.MinBetAmount(&_Btcvseth.CallOpts)
}

// MinBetAmount is a free data retrieval call binding the contract method 0xfa968eea.
//
// Solidity: function minBetAmount() view returns(uint256)
func (_Btcvseth *BtcvsethCallerSession) MinBetAmount() (*big.Int, error) {
	return _Btcvseth.Contract.MinBetAmount(&_Btcvseth.CallOpts)
}

// OperatorAddress is a free data retrieval call binding the contract method 0x127effb2.
//
// Solidity: function operatorAddress() view returns(address)
func (_Btcvseth *BtcvsethCaller) OperatorAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "operatorAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorAddress is a free data retrieval call binding the contract method 0x127effb2.
//
// Solidity: function operatorAddress() view returns(address)
func (_Btcvseth *BtcvsethSession) OperatorAddress() (common.Address, error) {
	return _Btcvseth.Contract.OperatorAddress(&_Btcvseth.CallOpts)
}

// OperatorAddress is a free data retrieval call binding the contract method 0x127effb2.
//
// Solidity: function operatorAddress() view returns(address)
func (_Btcvseth *BtcvsethCallerSession) OperatorAddress() (common.Address, error) {
	return _Btcvseth.Contract.OperatorAddress(&_Btcvseth.CallOpts)
}

// OracleFirstLatestRoundId is a free data retrieval call binding the contract method 0x91e5ae4f.
//
// Solidity: function oracleFirstLatestRoundId() view returns(uint256)
func (_Btcvseth *BtcvsethCaller) OracleFirstLatestRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "oracleFirstLatestRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OracleFirstLatestRoundId is a free data retrieval call binding the contract method 0x91e5ae4f.
//
// Solidity: function oracleFirstLatestRoundId() view returns(uint256)
func (_Btcvseth *BtcvsethSession) OracleFirstLatestRoundId() (*big.Int, error) {
	return _Btcvseth.Contract.OracleFirstLatestRoundId(&_Btcvseth.CallOpts)
}

// OracleFirstLatestRoundId is a free data retrieval call binding the contract method 0x91e5ae4f.
//
// Solidity: function oracleFirstLatestRoundId() view returns(uint256)
func (_Btcvseth *BtcvsethCallerSession) OracleFirstLatestRoundId() (*big.Int, error) {
	return _Btcvseth.Contract.OracleFirstLatestRoundId(&_Btcvseth.CallOpts)
}

// OracleSecondLatestRoundId is a free data retrieval call binding the contract method 0xb1802631.
//
// Solidity: function oracleSecondLatestRoundId() view returns(uint256)
func (_Btcvseth *BtcvsethCaller) OracleSecondLatestRoundId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "oracleSecondLatestRoundId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OracleSecondLatestRoundId is a free data retrieval call binding the contract method 0xb1802631.
//
// Solidity: function oracleSecondLatestRoundId() view returns(uint256)
func (_Btcvseth *BtcvsethSession) OracleSecondLatestRoundId() (*big.Int, error) {
	return _Btcvseth.Contract.OracleSecondLatestRoundId(&_Btcvseth.CallOpts)
}

// OracleSecondLatestRoundId is a free data retrieval call binding the contract method 0xb1802631.
//
// Solidity: function oracleSecondLatestRoundId() view returns(uint256)
func (_Btcvseth *BtcvsethCallerSession) OracleSecondLatestRoundId() (*big.Int, error) {
	return _Btcvseth.Contract.OracleSecondLatestRoundId(&_Btcvseth.CallOpts)
}

// OracleUpdateAllowance is a free data retrieval call binding the contract method 0x60554011.
//
// Solidity: function oracleUpdateAllowance() view returns(uint256)
func (_Btcvseth *BtcvsethCaller) OracleUpdateAllowance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "oracleUpdateAllowance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OracleUpdateAllowance is a free data retrieval call binding the contract method 0x60554011.
//
// Solidity: function oracleUpdateAllowance() view returns(uint256)
func (_Btcvseth *BtcvsethSession) OracleUpdateAllowance() (*big.Int, error) {
	return _Btcvseth.Contract.OracleUpdateAllowance(&_Btcvseth.CallOpts)
}

// OracleUpdateAllowance is a free data retrieval call binding the contract method 0x60554011.
//
// Solidity: function oracleUpdateAllowance() view returns(uint256)
func (_Btcvseth *BtcvsethCallerSession) OracleUpdateAllowance() (*big.Int, error) {
	return _Btcvseth.Contract.OracleUpdateAllowance(&_Btcvseth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Btcvseth *BtcvsethCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Btcvseth *BtcvsethSession) Owner() (common.Address, error) {
	return _Btcvseth.Contract.Owner(&_Btcvseth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Btcvseth *BtcvsethCallerSession) Owner() (common.Address, error) {
	return _Btcvseth.Contract.Owner(&_Btcvseth.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Btcvseth *BtcvsethCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Btcvseth *BtcvsethSession) Paused() (bool, error) {
	return _Btcvseth.Contract.Paused(&_Btcvseth.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Btcvseth *BtcvsethCallerSession) Paused() (bool, error) {
	return _Btcvseth.Contract.Paused(&_Btcvseth.CallOpts)
}

// Refundable is a free data retrieval call binding the contract method 0x7bf41254.
//
// Solidity: function refundable(uint256 epoch, address user) view returns(bool)
func (_Btcvseth *BtcvsethCaller) Refundable(opts *bind.CallOpts, epoch *big.Int, user common.Address) (bool, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "refundable", epoch, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Refundable is a free data retrieval call binding the contract method 0x7bf41254.
//
// Solidity: function refundable(uint256 epoch, address user) view returns(bool)
func (_Btcvseth *BtcvsethSession) Refundable(epoch *big.Int, user common.Address) (bool, error) {
	return _Btcvseth.Contract.Refundable(&_Btcvseth.CallOpts, epoch, user)
}

// Refundable is a free data retrieval call binding the contract method 0x7bf41254.
//
// Solidity: function refundable(uint256 epoch, address user) view returns(bool)
func (_Btcvseth *BtcvsethCallerSession) Refundable(epoch *big.Int, user common.Address) (bool, error) {
	return _Btcvseth.Contract.Refundable(&_Btcvseth.CallOpts, epoch, user)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Btcvseth *BtcvsethCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Btcvseth *BtcvsethSession) RewardRate() (*big.Int, error) {
	return _Btcvseth.Contract.RewardRate(&_Btcvseth.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Btcvseth *BtcvsethCallerSession) RewardRate() (*big.Int, error) {
	return _Btcvseth.Contract.RewardRate(&_Btcvseth.CallOpts)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(uint256 epoch, uint256 startBlock, uint256 lockBlock, uint256 endBlock, int256 lockFirstPrice, int256 closeFirstPrice, int256 lockSecondPrice, int256 closeSecondPrice, uint256 totalAmount, uint256 bullAmount, uint256 bearAmount, uint256 rewardBaseCalAmount, uint256 rewardAmount, bool oracleCalled)
func (_Btcvseth *BtcvsethCaller) Rounds(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Epoch               *big.Int
	StartBlock          *big.Int
	LockBlock           *big.Int
	EndBlock            *big.Int
	LockFirstPrice      *big.Int
	CloseFirstPrice     *big.Int
	LockSecondPrice     *big.Int
	CloseSecondPrice    *big.Int
	TotalAmount         *big.Int
	BullAmount          *big.Int
	BearAmount          *big.Int
	RewardBaseCalAmount *big.Int
	RewardAmount        *big.Int
	OracleCalled        bool
}, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "rounds", arg0)

	outstruct := new(struct {
		Epoch               *big.Int
		StartBlock          *big.Int
		LockBlock           *big.Int
		EndBlock            *big.Int
		LockFirstPrice      *big.Int
		CloseFirstPrice     *big.Int
		LockSecondPrice     *big.Int
		CloseSecondPrice    *big.Int
		TotalAmount         *big.Int
		BullAmount          *big.Int
		BearAmount          *big.Int
		RewardBaseCalAmount *big.Int
		RewardAmount        *big.Int
		OracleCalled        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Epoch = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LockBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LockFirstPrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CloseFirstPrice = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LockSecondPrice = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.CloseSecondPrice = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.TotalAmount = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.BullAmount = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.BearAmount = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.RewardBaseCalAmount = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.RewardAmount = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)
	outstruct.OracleCalled = *abi.ConvertType(out[13], new(bool)).(*bool)

	return *outstruct, err

}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(uint256 epoch, uint256 startBlock, uint256 lockBlock, uint256 endBlock, int256 lockFirstPrice, int256 closeFirstPrice, int256 lockSecondPrice, int256 closeSecondPrice, uint256 totalAmount, uint256 bullAmount, uint256 bearAmount, uint256 rewardBaseCalAmount, uint256 rewardAmount, bool oracleCalled)
func (_Btcvseth *BtcvsethSession) Rounds(arg0 *big.Int) (struct {
	Epoch               *big.Int
	StartBlock          *big.Int
	LockBlock           *big.Int
	EndBlock            *big.Int
	LockFirstPrice      *big.Int
	CloseFirstPrice     *big.Int
	LockSecondPrice     *big.Int
	CloseSecondPrice    *big.Int
	TotalAmount         *big.Int
	BullAmount          *big.Int
	BearAmount          *big.Int
	RewardBaseCalAmount *big.Int
	RewardAmount        *big.Int
	OracleCalled        bool
}, error) {
	return _Btcvseth.Contract.Rounds(&_Btcvseth.CallOpts, arg0)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(uint256 epoch, uint256 startBlock, uint256 lockBlock, uint256 endBlock, int256 lockFirstPrice, int256 closeFirstPrice, int256 lockSecondPrice, int256 closeSecondPrice, uint256 totalAmount, uint256 bullAmount, uint256 bearAmount, uint256 rewardBaseCalAmount, uint256 rewardAmount, bool oracleCalled)
func (_Btcvseth *BtcvsethCallerSession) Rounds(arg0 *big.Int) (struct {
	Epoch               *big.Int
	StartBlock          *big.Int
	LockBlock           *big.Int
	EndBlock            *big.Int
	LockFirstPrice      *big.Int
	CloseFirstPrice     *big.Int
	LockSecondPrice     *big.Int
	CloseSecondPrice    *big.Int
	TotalAmount         *big.Int
	BullAmount          *big.Int
	BearAmount          *big.Int
	RewardBaseCalAmount *big.Int
	RewardAmount        *big.Int
	OracleCalled        bool
}, error) {
	return _Btcvseth.Contract.Rounds(&_Btcvseth.CallOpts, arg0)
}

// TreasuryAmount is a free data retrieval call binding the contract method 0x368acb09.
//
// Solidity: function treasuryAmount() view returns(uint256)
func (_Btcvseth *BtcvsethCaller) TreasuryAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "treasuryAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryAmount is a free data retrieval call binding the contract method 0x368acb09.
//
// Solidity: function treasuryAmount() view returns(uint256)
func (_Btcvseth *BtcvsethSession) TreasuryAmount() (*big.Int, error) {
	return _Btcvseth.Contract.TreasuryAmount(&_Btcvseth.CallOpts)
}

// TreasuryAmount is a free data retrieval call binding the contract method 0x368acb09.
//
// Solidity: function treasuryAmount() view returns(uint256)
func (_Btcvseth *BtcvsethCallerSession) TreasuryAmount() (*big.Int, error) {
	return _Btcvseth.Contract.TreasuryAmount(&_Btcvseth.CallOpts)
}

// TreasuryRate is a free data retrieval call binding the contract method 0xe4b72516.
//
// Solidity: function treasuryRate() view returns(uint256)
func (_Btcvseth *BtcvsethCaller) TreasuryRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "treasuryRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryRate is a free data retrieval call binding the contract method 0xe4b72516.
//
// Solidity: function treasuryRate() view returns(uint256)
func (_Btcvseth *BtcvsethSession) TreasuryRate() (*big.Int, error) {
	return _Btcvseth.Contract.TreasuryRate(&_Btcvseth.CallOpts)
}

// TreasuryRate is a free data retrieval call binding the contract method 0xe4b72516.
//
// Solidity: function treasuryRate() view returns(uint256)
func (_Btcvseth *BtcvsethCallerSession) TreasuryRate() (*big.Int, error) {
	return _Btcvseth.Contract.TreasuryRate(&_Btcvseth.CallOpts)
}

// UserRounds is a free data retrieval call binding the contract method 0xdd1f7596.
//
// Solidity: function userRounds(address , uint256 ) view returns(uint256)
func (_Btcvseth *BtcvsethCaller) UserRounds(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Btcvseth.contract.Call(opts, &out, "userRounds", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRounds is a free data retrieval call binding the contract method 0xdd1f7596.
//
// Solidity: function userRounds(address , uint256 ) view returns(uint256)
func (_Btcvseth *BtcvsethSession) UserRounds(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Btcvseth.Contract.UserRounds(&_Btcvseth.CallOpts, arg0, arg1)
}

// UserRounds is a free data retrieval call binding the contract method 0xdd1f7596.
//
// Solidity: function userRounds(address , uint256 ) view returns(uint256)
func (_Btcvseth *BtcvsethCallerSession) UserRounds(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Btcvseth.Contract.UserRounds(&_Btcvseth.CallOpts, arg0, arg1)
}

// BetBear is a paid mutator transaction binding the contract method 0x0088160f.
//
// Solidity: function betBear() payable returns()
func (_Btcvseth *BtcvsethTransactor) BetBear(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Btcvseth.contract.Transact(opts, "betBear")
}

// BetBear is a paid mutator transaction binding the contract method 0x0088160f.
//
// Solidity: function betBear() payable returns()
func (_Btcvseth *BtcvsethSession) BetBear() (*types.Transaction, error) {
	return _Btcvseth.Contract.BetBear(&_Btcvseth.TransactOpts)
}

// BetBear is a paid mutator transaction binding the contract method 0x0088160f.
//
// Solidity: function betBear() payable returns()
func (_Btcvseth *BtcvsethTransactorSession) BetBear() (*types.Transaction, error) {
	return _Btcvseth.Contract.BetBear(&_Btcvseth.TransactOpts)
}

// BetBull is a paid mutator transaction binding the contract method 0x821daba1.
//
// Solidity: function betBull() payable returns()
func (_Btcvseth *BtcvsethTransactor) BetBull(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Btcvseth.contract.Transact(opts, "betBull")
}

// BetBull is a paid mutator transaction binding the contract method 0x821daba1.
//
// Solidity: function betBull() payable returns()
func (_Btcvseth *BtcvsethSession) BetBull() (*types.Transaction, error) {
	return _Btcvseth.Contract.BetBull(&_Btcvseth.TransactOpts)
}

// BetBull is a paid mutator transaction binding the contract method 0x821daba1.
//
// Solidity: function betBull() payable returns()
func (_Btcvseth *BtcvsethTransactorSession) BetBull() (*types.Transaction, error) {
	return _Btcvseth.Contract.BetBull(&_Btcvseth.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 epoch) returns()
func (_Btcvseth *BtcvsethTransactor) Claim(opts *bind.TransactOpts, epoch *big.Int) (*types.Transaction, error) {
	return _Btcvseth.contract.Transact(opts, "claim", epoch)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 epoch) returns()
func (_Btcvseth *BtcvsethSession) Claim(epoch *big.Int) (*types.Transaction, error) {
	return _Btcvseth.Contract.Claim(&_Btcvseth.TransactOpts, epoch)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 epoch) returns()
func (_Btcvseth *BtcvsethTransactorSession) Claim(epoch *big.Int) (*types.Transaction, error) {
	return _Btcvseth.Contract.Claim(&_Btcvseth.TransactOpts, epoch)
}

// ClaimTreasury is a paid mutator transaction binding the contract method 0x003bdc74.
//
// Solidity: function claimTreasury() returns()
func (_Btcvseth *BtcvsethTransactor) ClaimTreasury(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Btcvseth.contract.Transact(opts, "claimTreasury")
}

// ClaimTreasury is a paid mutator transaction binding the contract method 0x003bdc74.
//
// Solidity: function claimTreasury() returns()
func (_Btcvseth *BtcvsethSession) ClaimTreasury() (*types.Transaction, error) {
	return _Btcvseth.Contract.ClaimTreasury(&_Btcvseth.TransactOpts)
}

// ClaimTreasury is a paid mutator transaction binding the contract method 0x003bdc74.
//
// Solidity: function claimTreasury() returns()
func (_Btcvseth *BtcvsethTransactorSession) ClaimTreasury() (*types.Transaction, error) {
	return _Btcvseth.Contract.ClaimTreasury(&_Btcvseth.TransactOpts)
}

// ExecuteRound is a paid mutator transaction binding the contract method 0x7b3205f5.
//
// Solidity: function executeRound() returns()
func (_Btcvseth *BtcvsethTransactor) ExecuteRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Btcvseth.contract.Transact(opts, "executeRound")
}

// ExecuteRound is a paid mutator transaction binding the contract method 0x7b3205f5.
//
// Solidity: function executeRound() returns()
func (_Btcvseth *BtcvsethSession) ExecuteRound() (*types.Transaction, error) {
	return _Btcvseth.Contract.ExecuteRound(&_Btcvseth.TransactOpts)
}

// ExecuteRound is a paid mutator transaction binding the contract method 0x7b3205f5.
//
// Solidity: function executeRound() returns()
func (_Btcvseth *BtcvsethTransactorSession) ExecuteRound() (*types.Transaction, error) {
	return _Btcvseth.Contract.ExecuteRound(&_Btcvseth.TransactOpts)
}

// GenesisLockRound is a paid mutator transaction binding the contract method 0xd9d55eac.
//
// Solidity: function genesisLockRound() returns()
func (_Btcvseth *BtcvsethTransactor) GenesisLockRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Btcvseth.contract.Transact(opts, "genesisLockRound")
}

// GenesisLockRound is a paid mutator transaction binding the contract method 0xd9d55eac.
//
// Solidity: function genesisLockRound() returns()
func (_Btcvseth *BtcvsethSession) GenesisLockRound() (*types.Transaction, error) {
	return _Btcvseth.Contract.GenesisLockRound(&_Btcvseth.TransactOpts)
}

// GenesisLockRound is a paid mutator transaction binding the contract method 0xd9d55eac.
//
// Solidity: function genesisLockRound() returns()
func (_Btcvseth *BtcvsethTransactorSession) GenesisLockRound() (*types.Transaction, error) {
	return _Btcvseth.Contract.GenesisLockRound(&_Btcvseth.TransactOpts)
}

// GenesisStartRound is a paid mutator transaction binding the contract method 0x452fd75a.
//
// Solidity: function genesisStartRound() returns()
func (_Btcvseth *BtcvsethTransactor) GenesisStartRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Btcvseth.contract.Transact(opts, "genesisStartRound")
}

// GenesisStartRound is a paid mutator transaction binding the contract method 0x452fd75a.
//
// Solidity: function genesisStartRound() returns()
func (_Btcvseth *BtcvsethSession) GenesisStartRound() (*types.Transaction, error) {
	return _Btcvseth.Contract.GenesisStartRound(&_Btcvseth.TransactOpts)
}

// GenesisStartRound is a paid mutator transaction binding the contract method 0x452fd75a.
//
// Solidity: function genesisStartRound() returns()
func (_Btcvseth *BtcvsethTransactorSession) GenesisStartRound() (*types.Transaction, error) {
	return _Btcvseth.Contract.GenesisStartRound(&_Btcvseth.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Btcvseth *BtcvsethTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Btcvseth.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Btcvseth *BtcvsethSession) Pause() (*types.Transaction, error) {
	return _Btcvseth.Contract.Pause(&_Btcvseth.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Btcvseth *BtcvsethTransactorSession) Pause() (*types.Transaction, error) {
	return _Btcvseth.Contract.Pause(&_Btcvseth.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Btcvseth *BtcvsethTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Btcvseth.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Btcvseth *BtcvsethSession) RenounceOwnership() (*types.Transaction, error) {
	return _Btcvseth.Contract.RenounceOwnership(&_Btcvseth.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Btcvseth *BtcvsethTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Btcvseth.Contract.RenounceOwnership(&_Btcvseth.TransactOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _adminAddress) returns()
func (_Btcvseth *BtcvsethTransactor) SetAdmin(opts *bind.TransactOpts, _adminAddress common.Address) (*types.Transaction, error) {
	return _Btcvseth.contract.Transact(opts, "setAdmin", _adminAddress)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _adminAddress) returns()
func (_Btcvseth *BtcvsethSession) SetAdmin(_adminAddress common.Address) (*types.Transaction, error) {
	return _Btcvseth.Contract.SetAdmin(&_Btcvseth.TransactOpts, _adminAddress)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _adminAddress) returns()
func (_Btcvseth *BtcvsethTransactorSession) SetAdmin(_adminAddress common.Address) (*types.Transaction, error) {
	return _Btcvseth.Contract.SetAdmin(&_Btcvseth.TransactOpts, _adminAddress)
}

// SetBufferBlocks is a paid mutator transaction binding the contract method 0x10c0ee64.
//
// Solidity: function setBufferBlocks(uint256 _bufferBlocks) returns()
func (_Btcvseth *BtcvsethTransactor) SetBufferBlocks(opts *bind.TransactOpts, _bufferBlocks *big.Int) (*types.Transaction, error) {
	return _Btcvseth.contract.Transact(opts, "setBufferBlocks", _bufferBlocks)
}

// SetBufferBlocks is a paid mutator transaction binding the contract method 0x10c0ee64.
//
// Solidity: function setBufferBlocks(uint256 _bufferBlocks) returns()
func (_Btcvseth *BtcvsethSession) SetBufferBlocks(_bufferBlocks *big.Int) (*types.Transaction, error) {
	return _Btcvseth.Contract.SetBufferBlocks(&_Btcvseth.TransactOpts, _bufferBlocks)
}

// SetBufferBlocks is a paid mutator transaction binding the contract method 0x10c0ee64.
//
// Solidity: function setBufferBlocks(uint256 _bufferBlocks) returns()
func (_Btcvseth *BtcvsethTransactorSession) SetBufferBlocks(_bufferBlocks *big.Int) (*types.Transaction, error) {
	return _Btcvseth.Contract.SetBufferBlocks(&_Btcvseth.TransactOpts, _bufferBlocks)
}

// SetFirstOracle is a paid mutator transaction binding the contract method 0xb4a26343.
//
// Solidity: function setFirstOracle(address _oracle) returns()
func (_Btcvseth *BtcvsethTransactor) SetFirstOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _Btcvseth.contract.Transact(opts, "setFirstOracle", _oracle)
}

// SetFirstOracle is a paid mutator transaction binding the contract method 0xb4a26343.
//
// Solidity: function setFirstOracle(address _oracle) returns()
func (_Btcvseth *BtcvsethSession) SetFirstOracle(_oracle common.Address) (*types.Transaction, error) {
	return _Btcvseth.Contract.SetFirstOracle(&_Btcvseth.TransactOpts, _oracle)
}

// SetFirstOracle is a paid mutator transaction binding the contract method 0xb4a26343.
//
// Solidity: function setFirstOracle(address _oracle) returns()
func (_Btcvseth *BtcvsethTransactorSession) SetFirstOracle(_oracle common.Address) (*types.Transaction, error) {
	return _Btcvseth.Contract.SetFirstOracle(&_Btcvseth.TransactOpts, _oracle)
}

// SetIntervalBlocks is a paid mutator transaction binding the contract method 0xa067455b.
//
// Solidity: function setIntervalBlocks(uint256 _intervalBlocks) returns()
func (_Btcvseth *BtcvsethTransactor) SetIntervalBlocks(opts *bind.TransactOpts, _intervalBlocks *big.Int) (*types.Transaction, error) {
	return _Btcvseth.contract.Transact(opts, "setIntervalBlocks", _intervalBlocks)
}

// SetIntervalBlocks is a paid mutator transaction binding the contract method 0xa067455b.
//
// Solidity: function setIntervalBlocks(uint256 _intervalBlocks) returns()
func (_Btcvseth *BtcvsethSession) SetIntervalBlocks(_intervalBlocks *big.Int) (*types.Transaction, error) {
	return _Btcvseth.Contract.SetIntervalBlocks(&_Btcvseth.TransactOpts, _intervalBlocks)
}

// SetIntervalBlocks is a paid mutator transaction binding the contract method 0xa067455b.
//
// Solidity: function setIntervalBlocks(uint256 _intervalBlocks) returns()
func (_Btcvseth *BtcvsethTransactorSession) SetIntervalBlocks(_intervalBlocks *big.Int) (*types.Transaction, error) {
	return _Btcvseth.Contract.SetIntervalBlocks(&_Btcvseth.TransactOpts, _intervalBlocks)
}

// SetMinBetAmount is a paid mutator transaction binding the contract method 0x6c188593.
//
// Solidity: function setMinBetAmount(uint256 _minBetAmount) returns()
func (_Btcvseth *BtcvsethTransactor) SetMinBetAmount(opts *bind.TransactOpts, _minBetAmount *big.Int) (*types.Transaction, error) {
	return _Btcvseth.contract.Transact(opts, "setMinBetAmount", _minBetAmount)
}

// SetMinBetAmount is a paid mutator transaction binding the contract method 0x6c188593.
//
// Solidity: function setMinBetAmount(uint256 _minBetAmount) returns()
func (_Btcvseth *BtcvsethSession) SetMinBetAmount(_minBetAmount *big.Int) (*types.Transaction, error) {
	return _Btcvseth.Contract.SetMinBetAmount(&_Btcvseth.TransactOpts, _minBetAmount)
}

// SetMinBetAmount is a paid mutator transaction binding the contract method 0x6c188593.
//
// Solidity: function setMinBetAmount(uint256 _minBetAmount) returns()
func (_Btcvseth *BtcvsethTransactorSession) SetMinBetAmount(_minBetAmount *big.Int) (*types.Transaction, error) {
	return _Btcvseth.Contract.SetMinBetAmount(&_Btcvseth.TransactOpts, _minBetAmount)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operatorAddress) returns()
func (_Btcvseth *BtcvsethTransactor) SetOperator(opts *bind.TransactOpts, _operatorAddress common.Address) (*types.Transaction, error) {
	return _Btcvseth.contract.Transact(opts, "setOperator", _operatorAddress)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operatorAddress) returns()
func (_Btcvseth *BtcvsethSession) SetOperator(_operatorAddress common.Address) (*types.Transaction, error) {
	return _Btcvseth.Contract.SetOperator(&_Btcvseth.TransactOpts, _operatorAddress)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operatorAddress) returns()
func (_Btcvseth *BtcvsethTransactorSession) SetOperator(_operatorAddress common.Address) (*types.Transaction, error) {
	return _Btcvseth.Contract.SetOperator(&_Btcvseth.TransactOpts, _operatorAddress)
}

// SetOracleUpdateAllowance is a paid mutator transaction binding the contract method 0xcf2f5039.
//
// Solidity: function setOracleUpdateAllowance(uint256 _oracleUpdateAllowance) returns()
func (_Btcvseth *BtcvsethTransactor) SetOracleUpdateAllowance(opts *bind.TransactOpts, _oracleUpdateAllowance *big.Int) (*types.Transaction, error) {
	return _Btcvseth.contract.Transact(opts, "setOracleUpdateAllowance", _oracleUpdateAllowance)
}

// SetOracleUpdateAllowance is a paid mutator transaction binding the contract method 0xcf2f5039.
//
// Solidity: function setOracleUpdateAllowance(uint256 _oracleUpdateAllowance) returns()
func (_Btcvseth *BtcvsethSession) SetOracleUpdateAllowance(_oracleUpdateAllowance *big.Int) (*types.Transaction, error) {
	return _Btcvseth.Contract.SetOracleUpdateAllowance(&_Btcvseth.TransactOpts, _oracleUpdateAllowance)
}

// SetOracleUpdateAllowance is a paid mutator transaction binding the contract method 0xcf2f5039.
//
// Solidity: function setOracleUpdateAllowance(uint256 _oracleUpdateAllowance) returns()
func (_Btcvseth *BtcvsethTransactorSession) SetOracleUpdateAllowance(_oracleUpdateAllowance *big.Int) (*types.Transaction, error) {
	return _Btcvseth.Contract.SetOracleUpdateAllowance(&_Btcvseth.TransactOpts, _oracleUpdateAllowance)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x9e447fc6.
//
// Solidity: function setRewardRate(uint256 _rewardRate) returns()
func (_Btcvseth *BtcvsethTransactor) SetRewardRate(opts *bind.TransactOpts, _rewardRate *big.Int) (*types.Transaction, error) {
	return _Btcvseth.contract.Transact(opts, "setRewardRate", _rewardRate)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x9e447fc6.
//
// Solidity: function setRewardRate(uint256 _rewardRate) returns()
func (_Btcvseth *BtcvsethSession) SetRewardRate(_rewardRate *big.Int) (*types.Transaction, error) {
	return _Btcvseth.Contract.SetRewardRate(&_Btcvseth.TransactOpts, _rewardRate)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x9e447fc6.
//
// Solidity: function setRewardRate(uint256 _rewardRate) returns()
func (_Btcvseth *BtcvsethTransactorSession) SetRewardRate(_rewardRate *big.Int) (*types.Transaction, error) {
	return _Btcvseth.Contract.SetRewardRate(&_Btcvseth.TransactOpts, _rewardRate)
}

// SetSecondOracle is a paid mutator transaction binding the contract method 0xedb5e796.
//
// Solidity: function setSecondOracle(address _oracle) returns()
func (_Btcvseth *BtcvsethTransactor) SetSecondOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _Btcvseth.contract.Transact(opts, "setSecondOracle", _oracle)
}

// SetSecondOracle is a paid mutator transaction binding the contract method 0xedb5e796.
//
// Solidity: function setSecondOracle(address _oracle) returns()
func (_Btcvseth *BtcvsethSession) SetSecondOracle(_oracle common.Address) (*types.Transaction, error) {
	return _Btcvseth.Contract.SetSecondOracle(&_Btcvseth.TransactOpts, _oracle)
}

// SetSecondOracle is a paid mutator transaction binding the contract method 0xedb5e796.
//
// Solidity: function setSecondOracle(address _oracle) returns()
func (_Btcvseth *BtcvsethTransactorSession) SetSecondOracle(_oracle common.Address) (*types.Transaction, error) {
	return _Btcvseth.Contract.SetSecondOracle(&_Btcvseth.TransactOpts, _oracle)
}

// SetTreasuryRate is a paid mutator transaction binding the contract method 0xd0bf9c54.
//
// Solidity: function setTreasuryRate(uint256 _treasuryRate) returns()
func (_Btcvseth *BtcvsethTransactor) SetTreasuryRate(opts *bind.TransactOpts, _treasuryRate *big.Int) (*types.Transaction, error) {
	return _Btcvseth.contract.Transact(opts, "setTreasuryRate", _treasuryRate)
}

// SetTreasuryRate is a paid mutator transaction binding the contract method 0xd0bf9c54.
//
// Solidity: function setTreasuryRate(uint256 _treasuryRate) returns()
func (_Btcvseth *BtcvsethSession) SetTreasuryRate(_treasuryRate *big.Int) (*types.Transaction, error) {
	return _Btcvseth.Contract.SetTreasuryRate(&_Btcvseth.TransactOpts, _treasuryRate)
}

// SetTreasuryRate is a paid mutator transaction binding the contract method 0xd0bf9c54.
//
// Solidity: function setTreasuryRate(uint256 _treasuryRate) returns()
func (_Btcvseth *BtcvsethTransactorSession) SetTreasuryRate(_treasuryRate *big.Int) (*types.Transaction, error) {
	return _Btcvseth.Contract.SetTreasuryRate(&_Btcvseth.TransactOpts, _treasuryRate)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Btcvseth *BtcvsethTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Btcvseth.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Btcvseth *BtcvsethSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Btcvseth.Contract.TransferOwnership(&_Btcvseth.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Btcvseth *BtcvsethTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Btcvseth.Contract.TransferOwnership(&_Btcvseth.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Btcvseth *BtcvsethTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Btcvseth.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Btcvseth *BtcvsethSession) Unpause() (*types.Transaction, error) {
	return _Btcvseth.Contract.Unpause(&_Btcvseth.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Btcvseth *BtcvsethTransactorSession) Unpause() (*types.Transaction, error) {
	return _Btcvseth.Contract.Unpause(&_Btcvseth.TransactOpts)
}

// BtcvsethBetBearIterator is returned from FilterBetBear and is used to iterate over the raw logs and unpacked data for BetBear events raised by the Btcvseth contract.
type BtcvsethBetBearIterator struct {
	Event *BtcvsethBetBear // Event containing the contract specifics and raw log

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
func (it *BtcvsethBetBearIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcvsethBetBear)
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
		it.Event = new(BtcvsethBetBear)
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
func (it *BtcvsethBetBearIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcvsethBetBearIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcvsethBetBear represents a BetBear event raised by the Btcvseth contract.
type BtcvsethBetBear struct {
	Sender       common.Address
	CurrentEpoch *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBetBear is a free log retrieval operation binding the contract event 0x0d8c1fe3e67ab767116a81f122b83c2557a8c2564019cb7c4f83de1aeb1f1f0d.
//
// Solidity: event BetBear(address indexed sender, uint256 indexed currentEpoch, uint256 amount)
func (_Btcvseth *BtcvsethFilterer) FilterBetBear(opts *bind.FilterOpts, sender []common.Address, currentEpoch []*big.Int) (*BtcvsethBetBearIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var currentEpochRule []interface{}
	for _, currentEpochItem := range currentEpoch {
		currentEpochRule = append(currentEpochRule, currentEpochItem)
	}

	logs, sub, err := _Btcvseth.contract.FilterLogs(opts, "BetBear", senderRule, currentEpochRule)
	if err != nil {
		return nil, err
	}
	return &BtcvsethBetBearIterator{contract: _Btcvseth.contract, event: "BetBear", logs: logs, sub: sub}, nil
}

// WatchBetBear is a free log subscription operation binding the contract event 0x0d8c1fe3e67ab767116a81f122b83c2557a8c2564019cb7c4f83de1aeb1f1f0d.
//
// Solidity: event BetBear(address indexed sender, uint256 indexed currentEpoch, uint256 amount)
func (_Btcvseth *BtcvsethFilterer) WatchBetBear(opts *bind.WatchOpts, sink chan<- *BtcvsethBetBear, sender []common.Address, currentEpoch []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var currentEpochRule []interface{}
	for _, currentEpochItem := range currentEpoch {
		currentEpochRule = append(currentEpochRule, currentEpochItem)
	}

	logs, sub, err := _Btcvseth.contract.WatchLogs(opts, "BetBear", senderRule, currentEpochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcvsethBetBear)
				if err := _Btcvseth.contract.UnpackLog(event, "BetBear", log); err != nil {
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

// ParseBetBear is a log parse operation binding the contract event 0x0d8c1fe3e67ab767116a81f122b83c2557a8c2564019cb7c4f83de1aeb1f1f0d.
//
// Solidity: event BetBear(address indexed sender, uint256 indexed currentEpoch, uint256 amount)
func (_Btcvseth *BtcvsethFilterer) ParseBetBear(log types.Log) (*BtcvsethBetBear, error) {
	event := new(BtcvsethBetBear)
	if err := _Btcvseth.contract.UnpackLog(event, "BetBear", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtcvsethBetBullIterator is returned from FilterBetBull and is used to iterate over the raw logs and unpacked data for BetBull events raised by the Btcvseth contract.
type BtcvsethBetBullIterator struct {
	Event *BtcvsethBetBull // Event containing the contract specifics and raw log

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
func (it *BtcvsethBetBullIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcvsethBetBull)
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
		it.Event = new(BtcvsethBetBull)
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
func (it *BtcvsethBetBullIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcvsethBetBullIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcvsethBetBull represents a BetBull event raised by the Btcvseth contract.
type BtcvsethBetBull struct {
	Sender       common.Address
	CurrentEpoch *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBetBull is a free log retrieval operation binding the contract event 0x438122d8cff518d18388099a5181f0d17a12b4f1b55faedf6e4a6acee0060c12.
//
// Solidity: event BetBull(address indexed sender, uint256 indexed currentEpoch, uint256 amount)
func (_Btcvseth *BtcvsethFilterer) FilterBetBull(opts *bind.FilterOpts, sender []common.Address, currentEpoch []*big.Int) (*BtcvsethBetBullIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var currentEpochRule []interface{}
	for _, currentEpochItem := range currentEpoch {
		currentEpochRule = append(currentEpochRule, currentEpochItem)
	}

	logs, sub, err := _Btcvseth.contract.FilterLogs(opts, "BetBull", senderRule, currentEpochRule)
	if err != nil {
		return nil, err
	}
	return &BtcvsethBetBullIterator{contract: _Btcvseth.contract, event: "BetBull", logs: logs, sub: sub}, nil
}

// WatchBetBull is a free log subscription operation binding the contract event 0x438122d8cff518d18388099a5181f0d17a12b4f1b55faedf6e4a6acee0060c12.
//
// Solidity: event BetBull(address indexed sender, uint256 indexed currentEpoch, uint256 amount)
func (_Btcvseth *BtcvsethFilterer) WatchBetBull(opts *bind.WatchOpts, sink chan<- *BtcvsethBetBull, sender []common.Address, currentEpoch []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var currentEpochRule []interface{}
	for _, currentEpochItem := range currentEpoch {
		currentEpochRule = append(currentEpochRule, currentEpochItem)
	}

	logs, sub, err := _Btcvseth.contract.WatchLogs(opts, "BetBull", senderRule, currentEpochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcvsethBetBull)
				if err := _Btcvseth.contract.UnpackLog(event, "BetBull", log); err != nil {
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

// ParseBetBull is a log parse operation binding the contract event 0x438122d8cff518d18388099a5181f0d17a12b4f1b55faedf6e4a6acee0060c12.
//
// Solidity: event BetBull(address indexed sender, uint256 indexed currentEpoch, uint256 amount)
func (_Btcvseth *BtcvsethFilterer) ParseBetBull(log types.Log) (*BtcvsethBetBull, error) {
	event := new(BtcvsethBetBull)
	if err := _Btcvseth.contract.UnpackLog(event, "BetBull", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtcvsethClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Btcvseth contract.
type BtcvsethClaimIterator struct {
	Event *BtcvsethClaim // Event containing the contract specifics and raw log

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
func (it *BtcvsethClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcvsethClaim)
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
		it.Event = new(BtcvsethClaim)
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
func (it *BtcvsethClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcvsethClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcvsethClaim represents a Claim event raised by the Btcvseth contract.
type BtcvsethClaim struct {
	Sender       common.Address
	CurrentEpoch *big.Int
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x34fcbac0073d7c3d388e51312faf357774904998eeb8fca628b9e6f65ee1cbf7.
//
// Solidity: event Claim(address indexed sender, uint256 indexed currentEpoch, uint256 amount)
func (_Btcvseth *BtcvsethFilterer) FilterClaim(opts *bind.FilterOpts, sender []common.Address, currentEpoch []*big.Int) (*BtcvsethClaimIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var currentEpochRule []interface{}
	for _, currentEpochItem := range currentEpoch {
		currentEpochRule = append(currentEpochRule, currentEpochItem)
	}

	logs, sub, err := _Btcvseth.contract.FilterLogs(opts, "Claim", senderRule, currentEpochRule)
	if err != nil {
		return nil, err
	}
	return &BtcvsethClaimIterator{contract: _Btcvseth.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x34fcbac0073d7c3d388e51312faf357774904998eeb8fca628b9e6f65ee1cbf7.
//
// Solidity: event Claim(address indexed sender, uint256 indexed currentEpoch, uint256 amount)
func (_Btcvseth *BtcvsethFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *BtcvsethClaim, sender []common.Address, currentEpoch []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var currentEpochRule []interface{}
	for _, currentEpochItem := range currentEpoch {
		currentEpochRule = append(currentEpochRule, currentEpochItem)
	}

	logs, sub, err := _Btcvseth.contract.WatchLogs(opts, "Claim", senderRule, currentEpochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcvsethClaim)
				if err := _Btcvseth.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x34fcbac0073d7c3d388e51312faf357774904998eeb8fca628b9e6f65ee1cbf7.
//
// Solidity: event Claim(address indexed sender, uint256 indexed currentEpoch, uint256 amount)
func (_Btcvseth *BtcvsethFilterer) ParseClaim(log types.Log) (*BtcvsethClaim, error) {
	event := new(BtcvsethClaim)
	if err := _Btcvseth.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtcvsethClaimTreasuryIterator is returned from FilterClaimTreasury and is used to iterate over the raw logs and unpacked data for ClaimTreasury events raised by the Btcvseth contract.
type BtcvsethClaimTreasuryIterator struct {
	Event *BtcvsethClaimTreasury // Event containing the contract specifics and raw log

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
func (it *BtcvsethClaimTreasuryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcvsethClaimTreasury)
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
		it.Event = new(BtcvsethClaimTreasury)
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
func (it *BtcvsethClaimTreasuryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcvsethClaimTreasuryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcvsethClaimTreasury represents a ClaimTreasury event raised by the Btcvseth contract.
type BtcvsethClaimTreasury struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimTreasury is a free log retrieval operation binding the contract event 0x609175abb7f12481e4f200d1ef4fc834e6caac3d9eadba42d664352f6d0932ca.
//
// Solidity: event ClaimTreasury(uint256 amount)
func (_Btcvseth *BtcvsethFilterer) FilterClaimTreasury(opts *bind.FilterOpts) (*BtcvsethClaimTreasuryIterator, error) {

	logs, sub, err := _Btcvseth.contract.FilterLogs(opts, "ClaimTreasury")
	if err != nil {
		return nil, err
	}
	return &BtcvsethClaimTreasuryIterator{contract: _Btcvseth.contract, event: "ClaimTreasury", logs: logs, sub: sub}, nil
}

// WatchClaimTreasury is a free log subscription operation binding the contract event 0x609175abb7f12481e4f200d1ef4fc834e6caac3d9eadba42d664352f6d0932ca.
//
// Solidity: event ClaimTreasury(uint256 amount)
func (_Btcvseth *BtcvsethFilterer) WatchClaimTreasury(opts *bind.WatchOpts, sink chan<- *BtcvsethClaimTreasury) (event.Subscription, error) {

	logs, sub, err := _Btcvseth.contract.WatchLogs(opts, "ClaimTreasury")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcvsethClaimTreasury)
				if err := _Btcvseth.contract.UnpackLog(event, "ClaimTreasury", log); err != nil {
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

// ParseClaimTreasury is a log parse operation binding the contract event 0x609175abb7f12481e4f200d1ef4fc834e6caac3d9eadba42d664352f6d0932ca.
//
// Solidity: event ClaimTreasury(uint256 amount)
func (_Btcvseth *BtcvsethFilterer) ParseClaimTreasury(log types.Log) (*BtcvsethClaimTreasury, error) {
	event := new(BtcvsethClaimTreasury)
	if err := _Btcvseth.contract.UnpackLog(event, "ClaimTreasury", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtcvsethEndRoundIterator is returned from FilterEndRound and is used to iterate over the raw logs and unpacked data for EndRound events raised by the Btcvseth contract.
type BtcvsethEndRoundIterator struct {
	Event *BtcvsethEndRound // Event containing the contract specifics and raw log

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
func (it *BtcvsethEndRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcvsethEndRound)
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
		it.Event = new(BtcvsethEndRound)
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
func (it *BtcvsethEndRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcvsethEndRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcvsethEndRound represents a EndRound event raised by the Btcvseth contract.
type BtcvsethEndRound struct {
	Epoch       *big.Int
	BlockNumber *big.Int
	FirstPrice  *big.Int
	SecondPrice *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEndRound is a free log retrieval operation binding the contract event 0xe9a54ac39aff99983e21985115ac91548a416592079c5051ac3ea7e6716ded89.
//
// Solidity: event EndRound(uint256 indexed epoch, uint256 blockNumber, int256 firstPrice, int256 secondPrice)
func (_Btcvseth *BtcvsethFilterer) FilterEndRound(opts *bind.FilterOpts, epoch []*big.Int) (*BtcvsethEndRoundIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Btcvseth.contract.FilterLogs(opts, "EndRound", epochRule)
	if err != nil {
		return nil, err
	}
	return &BtcvsethEndRoundIterator{contract: _Btcvseth.contract, event: "EndRound", logs: logs, sub: sub}, nil
}

// WatchEndRound is a free log subscription operation binding the contract event 0xe9a54ac39aff99983e21985115ac91548a416592079c5051ac3ea7e6716ded89.
//
// Solidity: event EndRound(uint256 indexed epoch, uint256 blockNumber, int256 firstPrice, int256 secondPrice)
func (_Btcvseth *BtcvsethFilterer) WatchEndRound(opts *bind.WatchOpts, sink chan<- *BtcvsethEndRound, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Btcvseth.contract.WatchLogs(opts, "EndRound", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcvsethEndRound)
				if err := _Btcvseth.contract.UnpackLog(event, "EndRound", log); err != nil {
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

// ParseEndRound is a log parse operation binding the contract event 0xe9a54ac39aff99983e21985115ac91548a416592079c5051ac3ea7e6716ded89.
//
// Solidity: event EndRound(uint256 indexed epoch, uint256 blockNumber, int256 firstPrice, int256 secondPrice)
func (_Btcvseth *BtcvsethFilterer) ParseEndRound(log types.Log) (*BtcvsethEndRound, error) {
	event := new(BtcvsethEndRound)
	if err := _Btcvseth.contract.UnpackLog(event, "EndRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtcvsethLockRoundIterator is returned from FilterLockRound and is used to iterate over the raw logs and unpacked data for LockRound events raised by the Btcvseth contract.
type BtcvsethLockRoundIterator struct {
	Event *BtcvsethLockRound // Event containing the contract specifics and raw log

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
func (it *BtcvsethLockRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcvsethLockRound)
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
		it.Event = new(BtcvsethLockRound)
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
func (it *BtcvsethLockRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcvsethLockRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcvsethLockRound represents a LockRound event raised by the Btcvseth contract.
type BtcvsethLockRound struct {
	Epoch       *big.Int
	BlockNumber *big.Int
	FirstPrice  *big.Int
	SecondPrice *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLockRound is a free log retrieval operation binding the contract event 0x21ce5b865a74204726d5062101961300ef664b01ba3f6ddd9ee1bf71846f8e15.
//
// Solidity: event LockRound(uint256 indexed epoch, uint256 blockNumber, int256 firstPrice, int256 secondPrice)
func (_Btcvseth *BtcvsethFilterer) FilterLockRound(opts *bind.FilterOpts, epoch []*big.Int) (*BtcvsethLockRoundIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Btcvseth.contract.FilterLogs(opts, "LockRound", epochRule)
	if err != nil {
		return nil, err
	}
	return &BtcvsethLockRoundIterator{contract: _Btcvseth.contract, event: "LockRound", logs: logs, sub: sub}, nil
}

// WatchLockRound is a free log subscription operation binding the contract event 0x21ce5b865a74204726d5062101961300ef664b01ba3f6ddd9ee1bf71846f8e15.
//
// Solidity: event LockRound(uint256 indexed epoch, uint256 blockNumber, int256 firstPrice, int256 secondPrice)
func (_Btcvseth *BtcvsethFilterer) WatchLockRound(opts *bind.WatchOpts, sink chan<- *BtcvsethLockRound, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Btcvseth.contract.WatchLogs(opts, "LockRound", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcvsethLockRound)
				if err := _Btcvseth.contract.UnpackLog(event, "LockRound", log); err != nil {
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

// ParseLockRound is a log parse operation binding the contract event 0x21ce5b865a74204726d5062101961300ef664b01ba3f6ddd9ee1bf71846f8e15.
//
// Solidity: event LockRound(uint256 indexed epoch, uint256 blockNumber, int256 firstPrice, int256 secondPrice)
func (_Btcvseth *BtcvsethFilterer) ParseLockRound(log types.Log) (*BtcvsethLockRound, error) {
	event := new(BtcvsethLockRound)
	if err := _Btcvseth.contract.UnpackLog(event, "LockRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtcvsethMinBetAmountUpdatedIterator is returned from FilterMinBetAmountUpdated and is used to iterate over the raw logs and unpacked data for MinBetAmountUpdated events raised by the Btcvseth contract.
type BtcvsethMinBetAmountUpdatedIterator struct {
	Event *BtcvsethMinBetAmountUpdated // Event containing the contract specifics and raw log

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
func (it *BtcvsethMinBetAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcvsethMinBetAmountUpdated)
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
		it.Event = new(BtcvsethMinBetAmountUpdated)
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
func (it *BtcvsethMinBetAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcvsethMinBetAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcvsethMinBetAmountUpdated represents a MinBetAmountUpdated event raised by the Btcvseth contract.
type BtcvsethMinBetAmountUpdated struct {
	Epoch        *big.Int
	MinBetAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMinBetAmountUpdated is a free log retrieval operation binding the contract event 0x82480e97991520549ac4b5d0e1d97edb3bfac2666202923b6903d2a953dc7608.
//
// Solidity: event MinBetAmountUpdated(uint256 indexed epoch, uint256 minBetAmount)
func (_Btcvseth *BtcvsethFilterer) FilterMinBetAmountUpdated(opts *bind.FilterOpts, epoch []*big.Int) (*BtcvsethMinBetAmountUpdatedIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Btcvseth.contract.FilterLogs(opts, "MinBetAmountUpdated", epochRule)
	if err != nil {
		return nil, err
	}
	return &BtcvsethMinBetAmountUpdatedIterator{contract: _Btcvseth.contract, event: "MinBetAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchMinBetAmountUpdated is a free log subscription operation binding the contract event 0x82480e97991520549ac4b5d0e1d97edb3bfac2666202923b6903d2a953dc7608.
//
// Solidity: event MinBetAmountUpdated(uint256 indexed epoch, uint256 minBetAmount)
func (_Btcvseth *BtcvsethFilterer) WatchMinBetAmountUpdated(opts *bind.WatchOpts, sink chan<- *BtcvsethMinBetAmountUpdated, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Btcvseth.contract.WatchLogs(opts, "MinBetAmountUpdated", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcvsethMinBetAmountUpdated)
				if err := _Btcvseth.contract.UnpackLog(event, "MinBetAmountUpdated", log); err != nil {
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

// ParseMinBetAmountUpdated is a log parse operation binding the contract event 0x82480e97991520549ac4b5d0e1d97edb3bfac2666202923b6903d2a953dc7608.
//
// Solidity: event MinBetAmountUpdated(uint256 indexed epoch, uint256 minBetAmount)
func (_Btcvseth *BtcvsethFilterer) ParseMinBetAmountUpdated(log types.Log) (*BtcvsethMinBetAmountUpdated, error) {
	event := new(BtcvsethMinBetAmountUpdated)
	if err := _Btcvseth.contract.UnpackLog(event, "MinBetAmountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtcvsethOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Btcvseth contract.
type BtcvsethOwnershipTransferredIterator struct {
	Event *BtcvsethOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BtcvsethOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcvsethOwnershipTransferred)
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
		it.Event = new(BtcvsethOwnershipTransferred)
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
func (it *BtcvsethOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcvsethOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcvsethOwnershipTransferred represents a OwnershipTransferred event raised by the Btcvseth contract.
type BtcvsethOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Btcvseth *BtcvsethFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BtcvsethOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Btcvseth.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BtcvsethOwnershipTransferredIterator{contract: _Btcvseth.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Btcvseth *BtcvsethFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BtcvsethOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Btcvseth.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcvsethOwnershipTransferred)
				if err := _Btcvseth.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Btcvseth *BtcvsethFilterer) ParseOwnershipTransferred(log types.Log) (*BtcvsethOwnershipTransferred, error) {
	event := new(BtcvsethOwnershipTransferred)
	if err := _Btcvseth.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtcvsethPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Btcvseth contract.
type BtcvsethPauseIterator struct {
	Event *BtcvsethPause // Event containing the contract specifics and raw log

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
func (it *BtcvsethPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcvsethPause)
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
		it.Event = new(BtcvsethPause)
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
func (it *BtcvsethPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcvsethPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcvsethPause represents a Pause event raised by the Btcvseth contract.
type BtcvsethPause struct {
	Epoch *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x68b095021b1f40fe513109f513c66692f0b3219aee674a69f4efc57badb8201d.
//
// Solidity: event Pause(uint256 epoch)
func (_Btcvseth *BtcvsethFilterer) FilterPause(opts *bind.FilterOpts) (*BtcvsethPauseIterator, error) {

	logs, sub, err := _Btcvseth.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &BtcvsethPauseIterator{contract: _Btcvseth.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x68b095021b1f40fe513109f513c66692f0b3219aee674a69f4efc57badb8201d.
//
// Solidity: event Pause(uint256 epoch)
func (_Btcvseth *BtcvsethFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *BtcvsethPause) (event.Subscription, error) {

	logs, sub, err := _Btcvseth.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcvsethPause)
				if err := _Btcvseth.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x68b095021b1f40fe513109f513c66692f0b3219aee674a69f4efc57badb8201d.
//
// Solidity: event Pause(uint256 epoch)
func (_Btcvseth *BtcvsethFilterer) ParsePause(log types.Log) (*BtcvsethPause, error) {
	event := new(BtcvsethPause)
	if err := _Btcvseth.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtcvsethPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Btcvseth contract.
type BtcvsethPausedIterator struct {
	Event *BtcvsethPaused // Event containing the contract specifics and raw log

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
func (it *BtcvsethPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcvsethPaused)
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
		it.Event = new(BtcvsethPaused)
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
func (it *BtcvsethPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcvsethPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcvsethPaused represents a Paused event raised by the Btcvseth contract.
type BtcvsethPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Btcvseth *BtcvsethFilterer) FilterPaused(opts *bind.FilterOpts) (*BtcvsethPausedIterator, error) {

	logs, sub, err := _Btcvseth.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BtcvsethPausedIterator{contract: _Btcvseth.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Btcvseth *BtcvsethFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BtcvsethPaused) (event.Subscription, error) {

	logs, sub, err := _Btcvseth.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcvsethPaused)
				if err := _Btcvseth.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Btcvseth *BtcvsethFilterer) ParsePaused(log types.Log) (*BtcvsethPaused, error) {
	event := new(BtcvsethPaused)
	if err := _Btcvseth.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtcvsethRatesUpdatedIterator is returned from FilterRatesUpdated and is used to iterate over the raw logs and unpacked data for RatesUpdated events raised by the Btcvseth contract.
type BtcvsethRatesUpdatedIterator struct {
	Event *BtcvsethRatesUpdated // Event containing the contract specifics and raw log

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
func (it *BtcvsethRatesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcvsethRatesUpdated)
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
		it.Event = new(BtcvsethRatesUpdated)
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
func (it *BtcvsethRatesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcvsethRatesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcvsethRatesUpdated represents a RatesUpdated event raised by the Btcvseth contract.
type BtcvsethRatesUpdated struct {
	Epoch        *big.Int
	RewardRate   *big.Int
	TreasuryRate *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRatesUpdated is a free log retrieval operation binding the contract event 0x023010bc68e7f4c0be9887f513c570c7a0f5f511b9716abccd42bf3b8943532b.
//
// Solidity: event RatesUpdated(uint256 indexed epoch, uint256 rewardRate, uint256 treasuryRate)
func (_Btcvseth *BtcvsethFilterer) FilterRatesUpdated(opts *bind.FilterOpts, epoch []*big.Int) (*BtcvsethRatesUpdatedIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Btcvseth.contract.FilterLogs(opts, "RatesUpdated", epochRule)
	if err != nil {
		return nil, err
	}
	return &BtcvsethRatesUpdatedIterator{contract: _Btcvseth.contract, event: "RatesUpdated", logs: logs, sub: sub}, nil
}

// WatchRatesUpdated is a free log subscription operation binding the contract event 0x023010bc68e7f4c0be9887f513c570c7a0f5f511b9716abccd42bf3b8943532b.
//
// Solidity: event RatesUpdated(uint256 indexed epoch, uint256 rewardRate, uint256 treasuryRate)
func (_Btcvseth *BtcvsethFilterer) WatchRatesUpdated(opts *bind.WatchOpts, sink chan<- *BtcvsethRatesUpdated, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Btcvseth.contract.WatchLogs(opts, "RatesUpdated", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcvsethRatesUpdated)
				if err := _Btcvseth.contract.UnpackLog(event, "RatesUpdated", log); err != nil {
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

// ParseRatesUpdated is a log parse operation binding the contract event 0x023010bc68e7f4c0be9887f513c570c7a0f5f511b9716abccd42bf3b8943532b.
//
// Solidity: event RatesUpdated(uint256 indexed epoch, uint256 rewardRate, uint256 treasuryRate)
func (_Btcvseth *BtcvsethFilterer) ParseRatesUpdated(log types.Log) (*BtcvsethRatesUpdated, error) {
	event := new(BtcvsethRatesUpdated)
	if err := _Btcvseth.contract.UnpackLog(event, "RatesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtcvsethRewardsCalculatedIterator is returned from FilterRewardsCalculated and is used to iterate over the raw logs and unpacked data for RewardsCalculated events raised by the Btcvseth contract.
type BtcvsethRewardsCalculatedIterator struct {
	Event *BtcvsethRewardsCalculated // Event containing the contract specifics and raw log

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
func (it *BtcvsethRewardsCalculatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcvsethRewardsCalculated)
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
		it.Event = new(BtcvsethRewardsCalculated)
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
func (it *BtcvsethRewardsCalculatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcvsethRewardsCalculatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcvsethRewardsCalculated represents a RewardsCalculated event raised by the Btcvseth contract.
type BtcvsethRewardsCalculated struct {
	Epoch               *big.Int
	RewardBaseCalAmount *big.Int
	RewardAmount        *big.Int
	TreasuryAmount      *big.Int
	FirstPercent        *big.Int
	SecondPercent       *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRewardsCalculated is a free log retrieval operation binding the contract event 0x1a772e2472fb70d2e556c23d2829be73279aa2fa7997b9bd4660dbe1ff6e31f8.
//
// Solidity: event RewardsCalculated(uint256 indexed epoch, uint256 rewardBaseCalAmount, uint256 rewardAmount, uint256 treasuryAmount, int256 firstPercent, int256 secondPercent)
func (_Btcvseth *BtcvsethFilterer) FilterRewardsCalculated(opts *bind.FilterOpts, epoch []*big.Int) (*BtcvsethRewardsCalculatedIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Btcvseth.contract.FilterLogs(opts, "RewardsCalculated", epochRule)
	if err != nil {
		return nil, err
	}
	return &BtcvsethRewardsCalculatedIterator{contract: _Btcvseth.contract, event: "RewardsCalculated", logs: logs, sub: sub}, nil
}

// WatchRewardsCalculated is a free log subscription operation binding the contract event 0x1a772e2472fb70d2e556c23d2829be73279aa2fa7997b9bd4660dbe1ff6e31f8.
//
// Solidity: event RewardsCalculated(uint256 indexed epoch, uint256 rewardBaseCalAmount, uint256 rewardAmount, uint256 treasuryAmount, int256 firstPercent, int256 secondPercent)
func (_Btcvseth *BtcvsethFilterer) WatchRewardsCalculated(opts *bind.WatchOpts, sink chan<- *BtcvsethRewardsCalculated, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Btcvseth.contract.WatchLogs(opts, "RewardsCalculated", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcvsethRewardsCalculated)
				if err := _Btcvseth.contract.UnpackLog(event, "RewardsCalculated", log); err != nil {
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

// ParseRewardsCalculated is a log parse operation binding the contract event 0x1a772e2472fb70d2e556c23d2829be73279aa2fa7997b9bd4660dbe1ff6e31f8.
//
// Solidity: event RewardsCalculated(uint256 indexed epoch, uint256 rewardBaseCalAmount, uint256 rewardAmount, uint256 treasuryAmount, int256 firstPercent, int256 secondPercent)
func (_Btcvseth *BtcvsethFilterer) ParseRewardsCalculated(log types.Log) (*BtcvsethRewardsCalculated, error) {
	event := new(BtcvsethRewardsCalculated)
	if err := _Btcvseth.contract.UnpackLog(event, "RewardsCalculated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtcvsethStartRoundIterator is returned from FilterStartRound and is used to iterate over the raw logs and unpacked data for StartRound events raised by the Btcvseth contract.
type BtcvsethStartRoundIterator struct {
	Event *BtcvsethStartRound // Event containing the contract specifics and raw log

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
func (it *BtcvsethStartRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcvsethStartRound)
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
		it.Event = new(BtcvsethStartRound)
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
func (it *BtcvsethStartRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcvsethStartRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcvsethStartRound represents a StartRound event raised by the Btcvseth contract.
type BtcvsethStartRound struct {
	Epoch       *big.Int
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStartRound is a free log retrieval operation binding the contract event 0x0e5543feb86a4cd302f2b88b26c42be2d1673013a34e1f98bd6d524dd3b4ab41.
//
// Solidity: event StartRound(uint256 indexed epoch, uint256 blockNumber)
func (_Btcvseth *BtcvsethFilterer) FilterStartRound(opts *bind.FilterOpts, epoch []*big.Int) (*BtcvsethStartRoundIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Btcvseth.contract.FilterLogs(opts, "StartRound", epochRule)
	if err != nil {
		return nil, err
	}
	return &BtcvsethStartRoundIterator{contract: _Btcvseth.contract, event: "StartRound", logs: logs, sub: sub}, nil
}

// WatchStartRound is a free log subscription operation binding the contract event 0x0e5543feb86a4cd302f2b88b26c42be2d1673013a34e1f98bd6d524dd3b4ab41.
//
// Solidity: event StartRound(uint256 indexed epoch, uint256 blockNumber)
func (_Btcvseth *BtcvsethFilterer) WatchStartRound(opts *bind.WatchOpts, sink chan<- *BtcvsethStartRound, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Btcvseth.contract.WatchLogs(opts, "StartRound", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcvsethStartRound)
				if err := _Btcvseth.contract.UnpackLog(event, "StartRound", log); err != nil {
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

// ParseStartRound is a log parse operation binding the contract event 0x0e5543feb86a4cd302f2b88b26c42be2d1673013a34e1f98bd6d524dd3b4ab41.
//
// Solidity: event StartRound(uint256 indexed epoch, uint256 blockNumber)
func (_Btcvseth *BtcvsethFilterer) ParseStartRound(log types.Log) (*BtcvsethStartRound, error) {
	event := new(BtcvsethStartRound)
	if err := _Btcvseth.contract.UnpackLog(event, "StartRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtcvsethUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Btcvseth contract.
type BtcvsethUnpauseIterator struct {
	Event *BtcvsethUnpause // Event containing the contract specifics and raw log

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
func (it *BtcvsethUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcvsethUnpause)
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
		it.Event = new(BtcvsethUnpause)
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
func (it *BtcvsethUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcvsethUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcvsethUnpause represents a Unpause event raised by the Btcvseth contract.
type BtcvsethUnpause struct {
	Epoch *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0xaaa520fdd7d2c83061d632fa017b0432407e798818af63ea908589fceda39ab7.
//
// Solidity: event Unpause(uint256 epoch)
func (_Btcvseth *BtcvsethFilterer) FilterUnpause(opts *bind.FilterOpts) (*BtcvsethUnpauseIterator, error) {

	logs, sub, err := _Btcvseth.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &BtcvsethUnpauseIterator{contract: _Btcvseth.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0xaaa520fdd7d2c83061d632fa017b0432407e798818af63ea908589fceda39ab7.
//
// Solidity: event Unpause(uint256 epoch)
func (_Btcvseth *BtcvsethFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *BtcvsethUnpause) (event.Subscription, error) {

	logs, sub, err := _Btcvseth.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcvsethUnpause)
				if err := _Btcvseth.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0xaaa520fdd7d2c83061d632fa017b0432407e798818af63ea908589fceda39ab7.
//
// Solidity: event Unpause(uint256 epoch)
func (_Btcvseth *BtcvsethFilterer) ParseUnpause(log types.Log) (*BtcvsethUnpause, error) {
	event := new(BtcvsethUnpause)
	if err := _Btcvseth.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtcvsethUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Btcvseth contract.
type BtcvsethUnpausedIterator struct {
	Event *BtcvsethUnpaused // Event containing the contract specifics and raw log

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
func (it *BtcvsethUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcvsethUnpaused)
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
		it.Event = new(BtcvsethUnpaused)
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
func (it *BtcvsethUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcvsethUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcvsethUnpaused represents a Unpaused event raised by the Btcvseth contract.
type BtcvsethUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Btcvseth *BtcvsethFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BtcvsethUnpausedIterator, error) {

	logs, sub, err := _Btcvseth.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BtcvsethUnpausedIterator{contract: _Btcvseth.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Btcvseth *BtcvsethFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BtcvsethUnpaused) (event.Subscription, error) {

	logs, sub, err := _Btcvseth.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcvsethUnpaused)
				if err := _Btcvseth.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Btcvseth *BtcvsethFilterer) ParseUnpaused(log types.Log) (*BtcvsethUnpaused, error) {
	event := new(BtcvsethUnpaused)
	if err := _Btcvseth.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
