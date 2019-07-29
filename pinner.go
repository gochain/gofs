// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gofs

import (
	"math/big"
	"strings"

	"github.com/gochain-io/gochain/v3"
	"github.com/gochain-io/gochain/v3/accounts/abi"
	"github.com/gochain-io/gochain/v3/accounts/abi/bind"
	"github.com/gochain-io/gochain/v3/common"
	"github.com/gochain-io/gochain/v3/core/types"
	"github.com/gochain-io/gochain/v3/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = gochain.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PinnerABI is the input ABI used to generate the binding from.
const PinnerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cid\",\"type\":\"bytes\"}],\"name\":\"pin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deployed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"cid\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"bh\",\"type\":\"uint256\"}],\"name\":\"Pinned\",\"type\":\"event\"}]"

// PinnerBin is the compiled bytecode used for deploying new contracts.
const PinnerBin = `0x`

// DeployPinner deploys a new GoChain contract, binding an instance of Pinner to it.
func DeployPinner(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pinner, error) {
	parsed, err := abi.JSON(strings.NewReader(PinnerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PinnerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pinner{PinnerCaller: PinnerCaller{contract: contract}, PinnerTransactor: PinnerTransactor{contract: contract}, PinnerFilterer: PinnerFilterer{contract: contract}}, nil
}

// Pinner is an auto generated Go binding around an GoChain contract.
type Pinner struct {
	PinnerCaller     // Read-only binding to the contract
	PinnerTransactor // Write-only binding to the contract
	PinnerFilterer   // Log filterer for contract events
}

// PinnerCaller is an auto generated read-only Go binding around an GoChain contract.
type PinnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PinnerTransactor is an auto generated write-only Go binding around an GoChain contract.
type PinnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PinnerFilterer is an auto generated log filtering Go binding around an GoChain contract events.
type PinnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PinnerSession is an auto generated Go binding around an GoChain contract,
// with pre-set call and transact options.
type PinnerSession struct {
	Contract     *Pinner           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PinnerCallerSession is an auto generated read-only Go binding around an GoChain contract,
// with pre-set call options.
type PinnerCallerSession struct {
	Contract *PinnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PinnerTransactorSession is an auto generated write-only Go binding around an GoChain contract,
// with pre-set transact options.
type PinnerTransactorSession struct {
	Contract     *PinnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PinnerRaw is an auto generated low-level Go binding around an GoChain contract.
type PinnerRaw struct {
	Contract *Pinner // Generic contract binding to access the raw methods on
}

// PinnerCallerRaw is an auto generated low-level read-only Go binding around an GoChain contract.
type PinnerCallerRaw struct {
	Contract *PinnerCaller // Generic read-only contract binding to access the raw methods on
}

// PinnerTransactorRaw is an auto generated low-level write-only Go binding around an GoChain contract.
type PinnerTransactorRaw struct {
	Contract *PinnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPinner creates a new instance of Pinner, bound to a specific deployed contract.
func NewPinner(address common.Address, backend bind.ContractBackend) (*Pinner, error) {
	contract, err := bindPinner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pinner{PinnerCaller: PinnerCaller{contract: contract}, PinnerTransactor: PinnerTransactor{contract: contract}, PinnerFilterer: PinnerFilterer{contract: contract}}, nil
}

// NewPinnerCaller creates a new read-only instance of Pinner, bound to a specific deployed contract.
func NewPinnerCaller(address common.Address, caller bind.ContractCaller) (*PinnerCaller, error) {
	contract, err := bindPinner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PinnerCaller{contract: contract}, nil
}

// NewPinnerTransactor creates a new write-only instance of Pinner, bound to a specific deployed contract.
func NewPinnerTransactor(address common.Address, transactor bind.ContractTransactor) (*PinnerTransactor, error) {
	contract, err := bindPinner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PinnerTransactor{contract: contract}, nil
}

// NewPinnerFilterer creates a new log filterer instance of Pinner, bound to a specific deployed contract.
func NewPinnerFilterer(address common.Address, filterer bind.ContractFilterer) (*PinnerFilterer, error) {
	contract, err := bindPinner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PinnerFilterer{contract: contract}, nil
}

// bindPinner binds a generic wrapper to an already deployed contract.
func bindPinner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PinnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pinner *PinnerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pinner.Contract.PinnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pinner *PinnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pinner.Contract.PinnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pinner *PinnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pinner.Contract.PinnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pinner *PinnerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pinner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pinner *PinnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pinner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pinner *PinnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pinner.Contract.contract.Transact(opts, method, params...)
}

// Deployed is a free data retrieval call binding the contract method 0xf905c15a.
//
// Solidity: function deployed() constant returns(uint256)
func (_Pinner *PinnerCaller) Deployed(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pinner.contract.Call(opts, out, "deployed")
	return *ret0, err
}

// Deployed is a free data retrieval call binding the contract method 0xf905c15a.
//
// Solidity: function deployed() constant returns(uint256)
func (_Pinner *PinnerSession) Deployed() (*big.Int, error) {
	return _Pinner.Contract.Deployed(&_Pinner.CallOpts)
}

// Deployed is a free data retrieval call binding the contract method 0xf905c15a.
//
// Solidity: function deployed() constant returns(uint256)
func (_Pinner *PinnerCallerSession) Deployed() (*big.Int, error) {
	return _Pinner.Contract.Deployed(&_Pinner.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_Pinner *PinnerCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pinner.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_Pinner *PinnerSession) Rate() (*big.Int, error) {
	return _Pinner.Contract.Rate(&_Pinner.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_Pinner *PinnerCallerSession) Rate() (*big.Int, error) {
	return _Pinner.Contract.Rate(&_Pinner.CallOpts)
}

// Pin is a paid mutator transaction binding the contract method 0x7d1962f8.
//
// Solidity: function pin(bytes cid) returns(bool)
func (_Pinner *PinnerTransactor) Pin(opts *bind.TransactOpts, cid []byte) (*types.Transaction, error) {
	return _Pinner.contract.Transact(opts, "pin", cid)
}

// Pin is a paid mutator transaction binding the contract method 0x7d1962f8.
//
// Solidity: function pin(bytes cid) returns(bool)
func (_Pinner *PinnerSession) Pin(cid []byte) (*types.Transaction, error) {
	return _Pinner.Contract.Pin(&_Pinner.TransactOpts, cid)
}

// Pin is a paid mutator transaction binding the contract method 0x7d1962f8.
//
// Solidity: function pin(bytes cid) returns(bool)
func (_Pinner *PinnerTransactorSession) Pin(cid []byte) (*types.Transaction, error) {
	return _Pinner.Contract.Pin(&_Pinner.TransactOpts, cid)
}

// PinnerPinnedIterator is returned from FilterPinned and is used to iterate over the raw logs and unpacked data for Pinned events raised by the Pinner contract.
type PinnerPinnedIterator struct {
	Event *PinnerPinned // Event containing the contract specifics and raw log

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
func (it *PinnerPinnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PinnerPinned)
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
		it.Event = new(PinnerPinned)
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
func (it *PinnerPinnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PinnerPinnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PinnerPinned represents a Pinned event raised by the Pinner contract.
type PinnerPinned struct {
	User common.Address
	Cid  common.Hash
	Bh   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPinned is a free log retrieval operation binding the contract event 0x7c80eb99758dfe3e8aef5df583c1c9bab5369cf46b930b802f130edcfeac90ca.
//
// Solidity: event Pinned(address indexed user, bytes indexed cid, uint256 bh)
func (_Pinner *PinnerFilterer) FilterPinned(opts *bind.FilterOpts, user []common.Address, cid [][]byte) (*PinnerPinnedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _Pinner.contract.FilterLogs(opts, "Pinned", userRule, cidRule)
	if err != nil {
		return nil, err
	}
	return &PinnerPinnedIterator{contract: _Pinner.contract, event: "Pinned", logs: logs, sub: sub}, nil
}

// WatchPinned is a free log subscription operation binding the contract event 0x7c80eb99758dfe3e8aef5df583c1c9bab5369cf46b930b802f130edcfeac90ca.
//
// Solidity: event Pinned(address indexed user, bytes indexed cid, uint256 bh)
func (_Pinner *PinnerFilterer) WatchPinned(opts *bind.WatchOpts, sink chan<- *PinnerPinned, user []common.Address, cid [][]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _Pinner.contract.WatchLogs(opts, "Pinned", userRule, cidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PinnerPinned)
				if err := _Pinner.contract.UnpackLog(event, "Pinned", log); err != nil {
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
