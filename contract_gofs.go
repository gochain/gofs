// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gofs

import (
	"math/big"
	"strings"

	gochain "github.com/gochain/gochain/v3"
	"github.com/gochain/gochain/v3/accounts/abi"
	"github.com/gochain/gochain/v3/accounts/abi/bind"
	"github.com/gochain/gochain/v3/common"
	"github.com/gochain/gochain/v3/core/types"
	"github.com/gochain/gochain/v3/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = gochain.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GOFSABI is the input ABI used to generate the binding from.
const GOFSABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"cid\",\"type\":\"bytes\"}],\"name\":\"newWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"setRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cid\",\"type\":\"bytes\"}],\"name\":\"pin\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cid\",\"type\":\"bytes\"}],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cidByHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deployed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"cid\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"bh\",\"type\":\"uint256\"}],\"name\":\"Pinned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"cid\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"CreatedWallet\",\"type\":\"event\"}]"

// GOFS is an auto generated Go binding around an GoChain contract.
type GOFS struct {
	GOFSCaller     // Read-only binding to the contract
	GOFSTransactor // Write-only binding to the contract
	GOFSFilterer   // Log filterer for contract events
}

// GOFSCaller is an auto generated read-only Go binding around an GoChain contract.
type GOFSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GOFSTransactor is an auto generated write-only Go binding around an GoChain contract.
type GOFSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GOFSFilterer is an auto generated log filtering Go binding around an GoChain contract events.
type GOFSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GOFSSession is an auto generated Go binding around an GoChain contract,
// with pre-set call and transact options.
type GOFSSession struct {
	Contract     *GOFS             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GOFSCallerSession is an auto generated read-only Go binding around an GoChain contract,
// with pre-set call options.
type GOFSCallerSession struct {
	Contract *GOFSCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GOFSTransactorSession is an auto generated write-only Go binding around an GoChain contract,
// with pre-set transact options.
type GOFSTransactorSession struct {
	Contract     *GOFSTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GOFSRaw is an auto generated low-level Go binding around an GoChain contract.
type GOFSRaw struct {
	Contract *GOFS // Generic contract binding to access the raw methods on
}

// GOFSCallerRaw is an auto generated low-level read-only Go binding around an GoChain contract.
type GOFSCallerRaw struct {
	Contract *GOFSCaller // Generic read-only contract binding to access the raw methods on
}

// GOFSTransactorRaw is an auto generated low-level write-only Go binding around an GoChain contract.
type GOFSTransactorRaw struct {
	Contract *GOFSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGOFS creates a new instance of GOFS, bound to a specific deployed contract.
func NewGOFS(address common.Address, backend bind.ContractBackend) (*GOFS, error) {
	contract, err := bindGOFS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GOFS{GOFSCaller: GOFSCaller{contract: contract}, GOFSTransactor: GOFSTransactor{contract: contract}, GOFSFilterer: GOFSFilterer{contract: contract}}, nil
}

// NewGOFSCaller creates a new read-only instance of GOFS, bound to a specific deployed contract.
func NewGOFSCaller(address common.Address, caller bind.ContractCaller) (*GOFSCaller, error) {
	contract, err := bindGOFS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GOFSCaller{contract: contract}, nil
}

// NewGOFSTransactor creates a new write-only instance of GOFS, bound to a specific deployed contract.
func NewGOFSTransactor(address common.Address, transactor bind.ContractTransactor) (*GOFSTransactor, error) {
	contract, err := bindGOFS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GOFSTransactor{contract: contract}, nil
}

// NewGOFSFilterer creates a new log filterer instance of GOFS, bound to a specific deployed contract.
func NewGOFSFilterer(address common.Address, filterer bind.ContractFilterer) (*GOFSFilterer, error) {
	contract, err := bindGOFS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GOFSFilterer{contract: contract}, nil
}

// bindGOFS binds a generic wrapper to an already deployed contract.
func bindGOFS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GOFSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GOFS *GOFSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GOFS.Contract.GOFSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GOFS *GOFSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GOFS.Contract.GOFSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GOFS *GOFSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GOFS.Contract.GOFSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GOFS *GOFSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GOFS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GOFS *GOFSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GOFS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GOFS *GOFSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GOFS.Contract.contract.Transact(opts, method, params...)
}

// CidByHash is a free data retrieval call binding the contract method 0xe16cf225.
//
// Solidity: function cidByHash(bytes32 ) view returns(bytes)
func (_GOFS *GOFSCaller) CidByHash(opts *bind.CallOpts, arg0 [32]byte) ([]byte, error) {
	var out []interface{}
	err := _GOFS.contract.Call(opts, &out, "cidByHash", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CidByHash is a free data retrieval call binding the contract method 0xe16cf225.
//
// Solidity: function cidByHash(bytes32 ) view returns(bytes)
func (_GOFS *GOFSSession) CidByHash(arg0 [32]byte) ([]byte, error) {
	return _GOFS.Contract.CidByHash(&_GOFS.CallOpts, arg0)
}

// CidByHash is a free data retrieval call binding the contract method 0xe16cf225.
//
// Solidity: function cidByHash(bytes32 ) view returns(bytes)
func (_GOFS *GOFSCallerSession) CidByHash(arg0 [32]byte) ([]byte, error) {
	return _GOFS.Contract.CidByHash(&_GOFS.CallOpts, arg0)
}

// Deployed is a free data retrieval call binding the contract method 0xf905c15a.
//
// Solidity: function deployed() view returns(uint256)
func (_GOFS *GOFSCaller) Deployed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GOFS.contract.Call(opts, &out, "deployed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deployed is a free data retrieval call binding the contract method 0xf905c15a.
//
// Solidity: function deployed() view returns(uint256)
func (_GOFS *GOFSSession) Deployed() (*big.Int, error) {
	return _GOFS.Contract.Deployed(&_GOFS.CallOpts)
}

// Deployed is a free data retrieval call binding the contract method 0xf905c15a.
//
// Solidity: function deployed() view returns(uint256)
func (_GOFS *GOFSCallerSession) Deployed() (*big.Int, error) {
	return _GOFS.Contract.Deployed(&_GOFS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address addr)
func (_GOFS *GOFSCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GOFS.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address addr)
func (_GOFS *GOFSSession) Owner() (common.Address, error) {
	return _GOFS.Contract.Owner(&_GOFS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address addr)
func (_GOFS *GOFSCallerSession) Owner() (common.Address, error) {
	return _GOFS.Contract.Owner(&_GOFS.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_GOFS *GOFSCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GOFS.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_GOFS *GOFSSession) Rate() (*big.Int, error) {
	return _GOFS.Contract.Rate(&_GOFS.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_GOFS *GOFSCallerSession) Rate() (*big.Int, error) {
	return _GOFS.Contract.Rate(&_GOFS.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0xc5bf2249.
//
// Solidity: function wallet(bytes cid) view returns(address)
func (_GOFS *GOFSCaller) Wallet(opts *bind.CallOpts, cid []byte) (common.Address, error) {
	var out []interface{}
	err := _GOFS.contract.Call(opts, &out, "wallet", cid)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0xc5bf2249.
//
// Solidity: function wallet(bytes cid) view returns(address)
func (_GOFS *GOFSSession) Wallet(cid []byte) (common.Address, error) {
	return _GOFS.Contract.Wallet(&_GOFS.CallOpts, cid)
}

// Wallet is a free data retrieval call binding the contract method 0xc5bf2249.
//
// Solidity: function wallet(bytes cid) view returns(address)
func (_GOFS *GOFSCallerSession) Wallet(cid []byte) (common.Address, error) {
	return _GOFS.Contract.Wallet(&_GOFS.CallOpts, cid)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address addr) returns()
func (_GOFS *GOFSTransactor) ChangeOwner(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _GOFS.contract.Transact(opts, "changeOwner", addr)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address addr) returns()
func (_GOFS *GOFSSession) ChangeOwner(addr common.Address) (*types.Transaction, error) {
	return _GOFS.Contract.ChangeOwner(&_GOFS.TransactOpts, addr)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address addr) returns()
func (_GOFS *GOFSTransactorSession) ChangeOwner(addr common.Address) (*types.Transaction, error) {
	return _GOFS.Contract.ChangeOwner(&_GOFS.TransactOpts, addr)
}

// NewWallet is a paid mutator transaction binding the contract method 0x28c6fa6f.
//
// Solidity: function newWallet(bytes cid) returns()
func (_GOFS *GOFSTransactor) NewWallet(opts *bind.TransactOpts, cid []byte) (*types.Transaction, error) {
	return _GOFS.contract.Transact(opts, "newWallet", cid)
}

// NewWallet is a paid mutator transaction binding the contract method 0x28c6fa6f.
//
// Solidity: function newWallet(bytes cid) returns()
func (_GOFS *GOFSSession) NewWallet(cid []byte) (*types.Transaction, error) {
	return _GOFS.Contract.NewWallet(&_GOFS.TransactOpts, cid)
}

// NewWallet is a paid mutator transaction binding the contract method 0x28c6fa6f.
//
// Solidity: function newWallet(bytes cid) returns()
func (_GOFS *GOFSTransactorSession) NewWallet(cid []byte) (*types.Transaction, error) {
	return _GOFS.Contract.NewWallet(&_GOFS.TransactOpts, cid)
}

// Pin is a paid mutator transaction binding the contract method 0x7d1962f8.
//
// Solidity: function pin(bytes cid) payable returns()
func (_GOFS *GOFSTransactor) Pin(opts *bind.TransactOpts, cid []byte) (*types.Transaction, error) {
	return _GOFS.contract.Transact(opts, "pin", cid)
}

// Pin is a paid mutator transaction binding the contract method 0x7d1962f8.
//
// Solidity: function pin(bytes cid) payable returns()
func (_GOFS *GOFSSession) Pin(cid []byte) (*types.Transaction, error) {
	return _GOFS.Contract.Pin(&_GOFS.TransactOpts, cid)
}

// Pin is a paid mutator transaction binding the contract method 0x7d1962f8.
//
// Solidity: function pin(bytes cid) payable returns()
func (_GOFS *GOFSTransactorSession) Pin(cid []byte) (*types.Transaction, error) {
	return _GOFS.Contract.Pin(&_GOFS.TransactOpts, cid)
}

// SetRate is a paid mutator transaction binding the contract method 0x34fcf437.
//
// Solidity: function setRate(uint256 _rate) returns()
func (_GOFS *GOFSTransactor) SetRate(opts *bind.TransactOpts, _rate *big.Int) (*types.Transaction, error) {
	return _GOFS.contract.Transact(opts, "setRate", _rate)
}

// SetRate is a paid mutator transaction binding the contract method 0x34fcf437.
//
// Solidity: function setRate(uint256 _rate) returns()
func (_GOFS *GOFSSession) SetRate(_rate *big.Int) (*types.Transaction, error) {
	return _GOFS.Contract.SetRate(&_GOFS.TransactOpts, _rate)
}

// SetRate is a paid mutator transaction binding the contract method 0x34fcf437.
//
// Solidity: function setRate(uint256 _rate) returns()
func (_GOFS *GOFSTransactorSession) SetRate(_rate *big.Int) (*types.Transaction, error) {
	return _GOFS.Contract.SetRate(&_GOFS.TransactOpts, _rate)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns()
func (_GOFS *GOFSTransactor) Withdraw(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _GOFS.contract.Transact(opts, "withdraw", to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns()
func (_GOFS *GOFSSession) Withdraw(to common.Address) (*types.Transaction, error) {
	return _GOFS.Contract.Withdraw(&_GOFS.TransactOpts, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns()
func (_GOFS *GOFSTransactorSession) Withdraw(to common.Address) (*types.Transaction, error) {
	return _GOFS.Contract.Withdraw(&_GOFS.TransactOpts, to)
}

// GOFSCreatedWalletIterator is returned from FilterCreatedWallet and is used to iterate over the raw logs and unpacked data for CreatedWallet events raised by the GOFS contract.
type GOFSCreatedWalletIterator struct {
	Event *GOFSCreatedWallet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  gochain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GOFSCreatedWalletIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GOFSCreatedWallet)
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
		it.Event = new(GOFSCreatedWallet)
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
func (it *GOFSCreatedWalletIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GOFSCreatedWalletIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GOFSCreatedWallet represents a CreatedWallet event raised by the GOFS contract.
type GOFSCreatedWallet struct {
	User   common.Address
	Cid    common.Hash
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCreatedWallet is a free log retrieval operation binding the contract event 0x89c3649b91d0d77a0655fa3d84b050b21c775ba31bfbc37e440ec3ee04f28927.
//
// Solidity: event CreatedWallet(address indexed user, bytes indexed cid, address wallet)
func (_GOFS *GOFSFilterer) FilterCreatedWallet(opts *bind.FilterOpts, user []common.Address, cid [][]byte) (*GOFSCreatedWalletIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _GOFS.contract.FilterLogs(opts, "CreatedWallet", userRule, cidRule)
	if err != nil {
		return nil, err
	}
	return &GOFSCreatedWalletIterator{contract: _GOFS.contract, event: "CreatedWallet", logs: logs, sub: sub}, nil
}

// WatchCreatedWallet is a free log subscription operation binding the contract event 0x89c3649b91d0d77a0655fa3d84b050b21c775ba31bfbc37e440ec3ee04f28927.
//
// Solidity: event CreatedWallet(address indexed user, bytes indexed cid, address wallet)
func (_GOFS *GOFSFilterer) WatchCreatedWallet(opts *bind.WatchOpts, sink chan<- *GOFSCreatedWallet, user []common.Address, cid [][]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _GOFS.contract.WatchLogs(opts, "CreatedWallet", userRule, cidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GOFSCreatedWallet)
				if err := _GOFS.contract.UnpackLog(event, "CreatedWallet", log); err != nil {
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

// ParseCreatedWallet is a log parse operation binding the contract event 0x89c3649b91d0d77a0655fa3d84b050b21c775ba31bfbc37e440ec3ee04f28927.
//
// Solidity: event CreatedWallet(address indexed user, bytes indexed cid, address wallet)
func (_GOFS *GOFSFilterer) ParseCreatedWallet(log types.Log) (*GOFSCreatedWallet, error) {
	event := new(GOFSCreatedWallet)
	if err := _GOFS.contract.UnpackLog(event, "CreatedWallet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GOFSPinnedIterator is returned from FilterPinned and is used to iterate over the raw logs and unpacked data for Pinned events raised by the GOFS contract.
type GOFSPinnedIterator struct {
	Event *GOFSPinned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  gochain.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GOFSPinnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GOFSPinned)
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
		it.Event = new(GOFSPinned)
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
func (it *GOFSPinnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GOFSPinnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GOFSPinned represents a Pinned event raised by the GOFS contract.
type GOFSPinned struct {
	User common.Address
	Cid  common.Hash
	Bh   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPinned is a free log retrieval operation binding the contract event 0x7c80eb99758dfe3e8aef5df583c1c9bab5369cf46b930b802f130edcfeac90ca.
//
// Solidity: event Pinned(address indexed user, bytes indexed cid, uint256 bh)
func (_GOFS *GOFSFilterer) FilterPinned(opts *bind.FilterOpts, user []common.Address, cid [][]byte) (*GOFSPinnedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _GOFS.contract.FilterLogs(opts, "Pinned", userRule, cidRule)
	if err != nil {
		return nil, err
	}
	return &GOFSPinnedIterator{contract: _GOFS.contract, event: "Pinned", logs: logs, sub: sub}, nil
}

// WatchPinned is a free log subscription operation binding the contract event 0x7c80eb99758dfe3e8aef5df583c1c9bab5369cf46b930b802f130edcfeac90ca.
//
// Solidity: event Pinned(address indexed user, bytes indexed cid, uint256 bh)
func (_GOFS *GOFSFilterer) WatchPinned(opts *bind.WatchOpts, sink chan<- *GOFSPinned, user []common.Address, cid [][]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _GOFS.contract.WatchLogs(opts, "Pinned", userRule, cidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GOFSPinned)
				if err := _GOFS.contract.UnpackLog(event, "Pinned", log); err != nil {
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

// ParsePinned is a log parse operation binding the contract event 0x7c80eb99758dfe3e8aef5df583c1c9bab5369cf46b930b802f130edcfeac90ca.
//
// Solidity: event Pinned(address indexed user, bytes indexed cid, uint256 bh)
func (_GOFS *GOFSFilterer) ParsePinned(log types.Log) (*GOFSPinned, error) {
	event := new(GOFSPinned)
	if err := _GOFS.contract.UnpackLog(event, "Pinned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

