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

// IGOFSABI is the input ABI used to generate the binding from.
const IGOFSABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"cid\",\"type\":\"bytes\"}],\"name\":\"newWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cid\",\"type\":\"bytes\"}],\"name\":\"pin\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cid\",\"type\":\"bytes\"}],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deployed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"cid\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"bh\",\"type\":\"uint256\"}],\"name\":\"Pinned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"cid\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"CreatedWallet\",\"type\":\"event\"}]"

// IGOFSBin is the compiled bytecode used for deploying new contracts.
const IGOFSBin = `0x`

// DeployIGOFS deploys a new GoChain contract, binding an instance of IGOFS to it.
func DeployIGOFS(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IGOFS, error) {
	parsed, err := abi.JSON(strings.NewReader(IGOFSABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IGOFSBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IGOFS{IGOFSCaller: IGOFSCaller{contract: contract}, IGOFSTransactor: IGOFSTransactor{contract: contract}, IGOFSFilterer: IGOFSFilterer{contract: contract}}, nil
}

// IGOFS is an auto generated Go binding around an GoChain contract.
type IGOFS struct {
	IGOFSCaller     // Read-only binding to the contract
	IGOFSTransactor // Write-only binding to the contract
	IGOFSFilterer   // Log filterer for contract events
}

// IGOFSCaller is an auto generated read-only Go binding around an GoChain contract.
type IGOFSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGOFSTransactor is an auto generated write-only Go binding around an GoChain contract.
type IGOFSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGOFSFilterer is an auto generated log filtering Go binding around an GoChain contract events.
type IGOFSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGOFSSession is an auto generated Go binding around an GoChain contract,
// with pre-set call and transact options.
type IGOFSSession struct {
	Contract     *IGOFS            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGOFSCallerSession is an auto generated read-only Go binding around an GoChain contract,
// with pre-set call options.
type IGOFSCallerSession struct {
	Contract *IGOFSCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IGOFSTransactorSession is an auto generated write-only Go binding around an GoChain contract,
// with pre-set transact options.
type IGOFSTransactorSession struct {
	Contract     *IGOFSTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGOFSRaw is an auto generated low-level Go binding around an GoChain contract.
type IGOFSRaw struct {
	Contract *IGOFS // Generic contract binding to access the raw methods on
}

// IGOFSCallerRaw is an auto generated low-level read-only Go binding around an GoChain contract.
type IGOFSCallerRaw struct {
	Contract *IGOFSCaller // Generic read-only contract binding to access the raw methods on
}

// IGOFSTransactorRaw is an auto generated low-level write-only Go binding around an GoChain contract.
type IGOFSTransactorRaw struct {
	Contract *IGOFSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGOFS creates a new instance of IGOFS, bound to a specific deployed contract.
func NewIGOFS(address common.Address, backend bind.ContractBackend) (*IGOFS, error) {
	contract, err := bindIGOFS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGOFS{IGOFSCaller: IGOFSCaller{contract: contract}, IGOFSTransactor: IGOFSTransactor{contract: contract}, IGOFSFilterer: IGOFSFilterer{contract: contract}}, nil
}

// NewIGOFSCaller creates a new read-only instance of IGOFS, bound to a specific deployed contract.
func NewIGOFSCaller(address common.Address, caller bind.ContractCaller) (*IGOFSCaller, error) {
	contract, err := bindIGOFS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGOFSCaller{contract: contract}, nil
}

// NewIGOFSTransactor creates a new write-only instance of IGOFS, bound to a specific deployed contract.
func NewIGOFSTransactor(address common.Address, transactor bind.ContractTransactor) (*IGOFSTransactor, error) {
	contract, err := bindIGOFS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGOFSTransactor{contract: contract}, nil
}

// NewIGOFSFilterer creates a new log filterer instance of IGOFS, bound to a specific deployed contract.
func NewIGOFSFilterer(address common.Address, filterer bind.ContractFilterer) (*IGOFSFilterer, error) {
	contract, err := bindIGOFS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGOFSFilterer{contract: contract}, nil
}

// bindIGOFS binds a generic wrapper to an already deployed contract.
func bindIGOFS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IGOFSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGOFS *IGOFSRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IGOFS.Contract.IGOFSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGOFS *IGOFSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGOFS.Contract.IGOFSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGOFS *IGOFSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGOFS.Contract.IGOFSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGOFS *IGOFSCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IGOFS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGOFS *IGOFSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGOFS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGOFS *IGOFSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGOFS.Contract.contract.Transact(opts, method, params...)
}

// Deployed is a free data retrieval call binding the contract method 0xf905c15a.
//
// Solidity: function deployed() constant returns(uint256)
func (_IGOFS *IGOFSCaller) Deployed(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IGOFS.contract.Call(opts, out, "deployed")
	return *ret0, err
}

// Deployed is a free data retrieval call binding the contract method 0xf905c15a.
//
// Solidity: function deployed() constant returns(uint256)
func (_IGOFS *IGOFSSession) Deployed() (*big.Int, error) {
	return _IGOFS.Contract.Deployed(&_IGOFS.CallOpts)
}

// Deployed is a free data retrieval call binding the contract method 0xf905c15a.
//
// Solidity: function deployed() constant returns(uint256)
func (_IGOFS *IGOFSCallerSession) Deployed() (*big.Int, error) {
	return _IGOFS.Contract.Deployed(&_IGOFS.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_IGOFS *IGOFSCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IGOFS.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_IGOFS *IGOFSSession) Rate() (*big.Int, error) {
	return _IGOFS.Contract.Rate(&_IGOFS.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_IGOFS *IGOFSCallerSession) Rate() (*big.Int, error) {
	return _IGOFS.Contract.Rate(&_IGOFS.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0xc5bf2249.
//
// Solidity: function wallet(bytes cid) constant returns(address)
func (_IGOFS *IGOFSCaller) Wallet(opts *bind.CallOpts, cid []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IGOFS.contract.Call(opts, out, "wallet", cid)
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0xc5bf2249.
//
// Solidity: function wallet(bytes cid) constant returns(address)
func (_IGOFS *IGOFSSession) Wallet(cid []byte) (common.Address, error) {
	return _IGOFS.Contract.Wallet(&_IGOFS.CallOpts, cid)
}

// Wallet is a free data retrieval call binding the contract method 0xc5bf2249.
//
// Solidity: function wallet(bytes cid) constant returns(address)
func (_IGOFS *IGOFSCallerSession) Wallet(cid []byte) (common.Address, error) {
	return _IGOFS.Contract.Wallet(&_IGOFS.CallOpts, cid)
}

// NewWallet is a paid mutator transaction binding the contract method 0x28c6fa6f.
//
// Solidity: function newWallet(bytes cid) returns()
func (_IGOFS *IGOFSTransactor) NewWallet(opts *bind.TransactOpts, cid []byte) (*types.Transaction, error) {
	return _IGOFS.contract.Transact(opts, "newWallet", cid)
}

// NewWallet is a paid mutator transaction binding the contract method 0x28c6fa6f.
//
// Solidity: function newWallet(bytes cid) returns()
func (_IGOFS *IGOFSSession) NewWallet(cid []byte) (*types.Transaction, error) {
	return _IGOFS.Contract.NewWallet(&_IGOFS.TransactOpts, cid)
}

// NewWallet is a paid mutator transaction binding the contract method 0x28c6fa6f.
//
// Solidity: function newWallet(bytes cid) returns()
func (_IGOFS *IGOFSTransactorSession) NewWallet(cid []byte) (*types.Transaction, error) {
	return _IGOFS.Contract.NewWallet(&_IGOFS.TransactOpts, cid)
}

// Pin is a paid mutator transaction binding the contract method 0x7d1962f8.
//
// Solidity: function pin(bytes cid) returns()
func (_IGOFS *IGOFSTransactor) Pin(opts *bind.TransactOpts, cid []byte) (*types.Transaction, error) {
	return _IGOFS.contract.Transact(opts, "pin", cid)
}

// Pin is a paid mutator transaction binding the contract method 0x7d1962f8.
//
// Solidity: function pin(bytes cid) returns()
func (_IGOFS *IGOFSSession) Pin(cid []byte) (*types.Transaction, error) {
	return _IGOFS.Contract.Pin(&_IGOFS.TransactOpts, cid)
}

// Pin is a paid mutator transaction binding the contract method 0x7d1962f8.
//
// Solidity: function pin(bytes cid) returns()
func (_IGOFS *IGOFSTransactorSession) Pin(cid []byte) (*types.Transaction, error) {
	return _IGOFS.Contract.Pin(&_IGOFS.TransactOpts, cid)
}

// IGOFSCreatedWalletIterator is returned from FilterCreatedWallet and is used to iterate over the raw logs and unpacked data for CreatedWallet events raised by the IGOFS contract.
type IGOFSCreatedWalletIterator struct {
	Event *IGOFSCreatedWallet // Event containing the contract specifics and raw log

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
func (it *IGOFSCreatedWalletIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGOFSCreatedWallet)
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
		it.Event = new(IGOFSCreatedWallet)
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
func (it *IGOFSCreatedWalletIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGOFSCreatedWalletIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGOFSCreatedWallet represents a CreatedWallet event raised by the IGOFS contract.
type IGOFSCreatedWallet struct {
	User   common.Address
	Cid    common.Hash
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCreatedWallet is a free log retrieval operation binding the contract event 0x89c3649b91d0d77a0655fa3d84b050b21c775ba31bfbc37e440ec3ee04f28927.
//
// Solidity: event CreatedWallet(address indexed user, bytes indexed cid, address wallet)
func (_IGOFS *IGOFSFilterer) FilterCreatedWallet(opts *bind.FilterOpts, user []common.Address, cid [][]byte) (*IGOFSCreatedWalletIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _IGOFS.contract.FilterLogs(opts, "CreatedWallet", userRule, cidRule)
	if err != nil {
		return nil, err
	}
	return &IGOFSCreatedWalletIterator{contract: _IGOFS.contract, event: "CreatedWallet", logs: logs, sub: sub}, nil
}

// WatchCreatedWallet is a free log subscription operation binding the contract event 0x89c3649b91d0d77a0655fa3d84b050b21c775ba31bfbc37e440ec3ee04f28927.
//
// Solidity: event CreatedWallet(address indexed user, bytes indexed cid, address wallet)
func (_IGOFS *IGOFSFilterer) WatchCreatedWallet(opts *bind.WatchOpts, sink chan<- *IGOFSCreatedWallet, user []common.Address, cid [][]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _IGOFS.contract.WatchLogs(opts, "CreatedWallet", userRule, cidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGOFSCreatedWallet)
				if err := _IGOFS.contract.UnpackLog(event, "CreatedWallet", log); err != nil {
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

// IGOFSPinnedIterator is returned from FilterPinned and is used to iterate over the raw logs and unpacked data for Pinned events raised by the IGOFS contract.
type IGOFSPinnedIterator struct {
	Event *IGOFSPinned // Event containing the contract specifics and raw log

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
func (it *IGOFSPinnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGOFSPinned)
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
		it.Event = new(IGOFSPinned)
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
func (it *IGOFSPinnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGOFSPinnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGOFSPinned represents a Pinned event raised by the IGOFS contract.
type IGOFSPinned struct {
	User common.Address
	Cid  common.Hash
	Bh   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPinned is a free log retrieval operation binding the contract event 0x7c80eb99758dfe3e8aef5df583c1c9bab5369cf46b930b802f130edcfeac90ca.
//
// Solidity: event Pinned(address indexed user, bytes indexed cid, uint256 bh)
func (_IGOFS *IGOFSFilterer) FilterPinned(opts *bind.FilterOpts, user []common.Address, cid [][]byte) (*IGOFSPinnedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _IGOFS.contract.FilterLogs(opts, "Pinned", userRule, cidRule)
	if err != nil {
		return nil, err
	}
	return &IGOFSPinnedIterator{contract: _IGOFS.contract, event: "Pinned", logs: logs, sub: sub}, nil
}

// WatchPinned is a free log subscription operation binding the contract event 0x7c80eb99758dfe3e8aef5df583c1c9bab5369cf46b930b802f130edcfeac90ca.
//
// Solidity: event Pinned(address indexed user, bytes indexed cid, uint256 bh)
func (_IGOFS *IGOFSFilterer) WatchPinned(opts *bind.WatchOpts, sink chan<- *IGOFSPinned, user []common.Address, cid [][]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _IGOFS.contract.WatchLogs(opts, "Pinned", userRule, cidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGOFSPinned)
				if err := _IGOFS.contract.UnpackLog(event, "Pinned", log); err != nil {
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
