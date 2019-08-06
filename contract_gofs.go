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

// GOFSABI is the input ABI used to generate the binding from.
const GOFSABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"cid\",\"type\":\"bytes\"}],\"name\":\"newWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"setRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cid\",\"type\":\"bytes\"}],\"name\":\"pin\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cid\",\"type\":\"bytes\"}],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deployed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_rate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"cid\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"bh\",\"type\":\"uint256\"}],\"name\":\"Pinned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"cid\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"CreatedWallet\",\"type\":\"event\"}]"

// GOFSBin is the compiled bytecode used for deploying new contracts.
const GOFSBin = `0x608060405234801561001057600080fd5b50604051602080610dc78339810180604052602081101561003057600080fd5b505160008054600160a060020a0319163317905560015543600255610d6d8061005a6000396000f3fe608060405260043610610098576000357c0100000000000000000000000000000000000000000000000000000000900480637d1962f81161006b5780637d1962f8146101d6578063a6f9dae11461027c578063c5bf2249146102af578063f905c15a1461037e57610098565b806328c6fa6f1461009d5780632c4e722e1461015257806334fcf4371461017957806351cff8d9146101a3575b600080fd5b3480156100a957600080fd5b50610150600480360360208110156100c057600080fd5b8101906020810181356401000000008111156100db57600080fd5b8201836020820111156100ed57600080fd5b8035906020019184600183028401116401000000008311171561010f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610393945050505050565b005b34801561015e57600080fd5b506101676106f3565b60408051918252519081900360200190f35b34801561018557600080fd5b506101506004803603602081101561019c57600080fd5b50356106f9565b3480156101af57600080fd5b50610150600480360360208110156101c657600080fd5b5035600160a060020a031661074a565b610150600480360360208110156101ec57600080fd5b81019060208101813564010000000081111561020757600080fd5b82018360208201111561021957600080fd5b8035906020019184600183028401116401000000008311171561023b57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506107d0945050505050565b34801561028857600080fd5b506101506004803603602081101561029f57600080fd5b5035600160a060020a0316610993565b3480156102bb57600080fd5b50610362600480360360208110156102d257600080fd5b8101906020810181356401000000008111156102ed57600080fd5b8201836020820111156102ff57600080fd5b8035906020019184600183028401116401000000008311171561032157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a0e945050505050565b60408051600160a060020a039092168252519081900360200190f35b34801561038a57600080fd5b50610167610a7f565b8060008151811015156103a257fe5b90602001015160f860020a900460f860020a02600160f860020a031916601260f860020a0214801561040557508060018151811015156103de57fe5b90602001015160f860020a900460f860020a02600160f860020a031916602060f860020a02145b1561045a576040805160e560020a62461bcd02815260206004820152601a60248201527f56657273696f6e203020434944206e6f7420616c6c6f7765642e000000000000604482015290519081900360640190fd5b6000600160a060020a03166003826040518082805190602001908083835b602083106104975780518252601f199092019160209182019101610478565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054600160a060020a03169290921491506105289050576040805160e560020a62461bcd02815260206004820152601d60248201527f57616c6c657420616c72656164792065786973747320666f7220636964000000604482015290519081900360640190fd5b6000308260405161053890610a85565b600160a060020a0383168152604060208083018281528451928401929092528351606084019185019080838360005b8381101561057f578181015183820152602001610567565b50505050905090810190601f1680156105ac5780820380516001836020036101000a031916815260200191505b509350505050604051809103906000f0801580156105ce573d6000803e3d6000fd5b509050806003836040518082805190602001908083835b602083106106045780518252601f1990920191602091820191016105e5565b51815160209384036101000a60001901801990921691161790529201948552506040519384900381018420805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0396909616959095179094555050835184928291908401908083835b6020831061068a5780518252601f19909201916020918201910161066b565b51815160209384036101000a600019018019909216911617905260408051929094018290038220600160a060020a038816835293519395503394507f89c3649b91d0d77a0655fa3d84b050b21c775ba31bfbc37e440ec3ee04f289279391829003019150a35050565b60015481565b600054600160a060020a031633146107455760405160e560020a62461bcd028152600401808060200182810382526022815260200180610d206022913960400191505060405180910390fd5b600155565b600054600160a060020a031633146107965760405160e560020a62461bcd028152600401808060200182810382526022815260200180610d206022913960400191505060405180910390fd5b604051600160a060020a03821690303180156108fc02916000818181858888f193505050501580156107cc573d6000803e3d6000fd5b5050565b8060008151811015156107df57fe5b90602001015160f860020a900460f860020a02600160f860020a031916601260f860020a02148015610842575080600181518110151561081b57fe5b90602001015160f860020a900460f860020a02600160f860020a031916602060f860020a02145b15610897576040805160e560020a62461bcd02815260206004820152601a60248201527f56657273696f6e203020434944206e6f7420616c6c6f7765642e000000000000604482015290519081900360640190fd5b6001543410156108f1576040805160e560020a62461bcd02815260206004820152601a60248201527f43616e6e6f7420707572636861736520302073746f726167652e000000000000604482015290519081900360640190fd5b60006001543481151561090057fe5b049050816040518082805190602001908083835b602083106109335780518252601f199092019160209182019101610914565b51815160209384036101000a60001901801990921691161790526040805192909401829003822087835293519395503294507f7c80eb99758dfe3e8aef5df583c1c9bab5369cf46b930b802f130edcfeac90ca9391829003019150a35050565b600054600160a060020a031633146109df5760405160e560020a62461bcd028152600401808060200182810382526022815260200180610d206022913960400191505060405180910390fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60006003826040518082805190602001908083835b60208310610a425780518252601f199092019160209182019101610a23565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054600160a060020a0316949350505050565b60025481565b61028d80610a938339019056fe608060405234801561001057600080fd5b5060405161028d38038061028d8339810180604052604081101561003357600080fd5b81516020830180519193928301929164010000000081111561005457600080fd5b8201602081018481111561006757600080fd5b815164010000000081118282018710171561008157600080fd5b505060008054600160a060020a031916600160a060020a03871617905580519093506100b692506001915060208401906100be565b505050610159565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100ff57805160ff191683800117855561012c565b8280016001018555821561012c579182015b8281111561012c578251825591602001919060010190610111565b5061013892915061013c565b5090565b61015691905b808211156101385760008155600101610142565b90565b610125806101686000396000f3fe608060408190526000547f7d1962f800000000000000000000000000000000000000000000000000000000909152602060849081526001805460026101008284161502600019019091160460a481905273ffffffffffffffffffffffffffffffffffffffff90931692637d1962f892349291819060c4908490801560c15780601f1060975761010080835404028352916020019160c1565b820191906000526020600020905b81548152906001019060200180831160a557829003601f168201915b5050925050506000604051808303818588803b15801560df57600080fd5b505af115801560f2573d6000803e3d6000fd5b505050505000fea165627a7a72305820f7cf0bfec932874e3688b55bee7638938879a192eda477d0c2420b67ad74fa8d00294f6e6c79206f776e65722063616e2063616c6c20746869732066756e6374696f6e2ea165627a7a72305820ec51aad0e58fdaf76881e18913e72b852a843decbeeaa0fad3bdd37f0f5c8ff40029`

// DeployGOFS deploys a new GoChain contract, binding an instance of GOFS to it.
func DeployGOFS(auth *bind.TransactOpts, backend bind.ContractBackend, _rate *big.Int) (common.Address, *types.Transaction, *GOFS, error) {
	parsed, err := abi.JSON(strings.NewReader(GOFSABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GOFSBin), backend, _rate)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GOFS{GOFSCaller: GOFSCaller{contract: contract}, GOFSTransactor: GOFSTransactor{contract: contract}, GOFSFilterer: GOFSFilterer{contract: contract}}, nil
}

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
func (_GOFS *GOFSRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_GOFS *GOFSCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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

// Deployed is a free data retrieval call binding the contract method 0xf905c15a.
//
// Solidity: function deployed() constant returns(uint256)
func (_GOFS *GOFSCaller) Deployed(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GOFS.contract.Call(opts, out, "deployed")
	return *ret0, err
}

// Deployed is a free data retrieval call binding the contract method 0xf905c15a.
//
// Solidity: function deployed() constant returns(uint256)
func (_GOFS *GOFSSession) Deployed() (*big.Int, error) {
	return _GOFS.Contract.Deployed(&_GOFS.CallOpts)
}

// Deployed is a free data retrieval call binding the contract method 0xf905c15a.
//
// Solidity: function deployed() constant returns(uint256)
func (_GOFS *GOFSCallerSession) Deployed() (*big.Int, error) {
	return _GOFS.Contract.Deployed(&_GOFS.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_GOFS *GOFSCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GOFS.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_GOFS *GOFSSession) Rate() (*big.Int, error) {
	return _GOFS.Contract.Rate(&_GOFS.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_GOFS *GOFSCallerSession) Rate() (*big.Int, error) {
	return _GOFS.Contract.Rate(&_GOFS.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0xc5bf2249.
//
// Solidity: function wallet(bytes cid) constant returns(address)
func (_GOFS *GOFSCaller) Wallet(opts *bind.CallOpts, cid []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GOFS.contract.Call(opts, out, "wallet", cid)
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0xc5bf2249.
//
// Solidity: function wallet(bytes cid) constant returns(address)
func (_GOFS *GOFSSession) Wallet(cid []byte) (common.Address, error) {
	return _GOFS.Contract.Wallet(&_GOFS.CallOpts, cid)
}

// Wallet is a free data retrieval call binding the contract method 0xc5bf2249.
//
// Solidity: function wallet(bytes cid) constant returns(address)
func (_GOFS *GOFSCallerSession) Wallet(cid []byte) (common.Address, error) {
	return _GOFS.Contract.Wallet(&_GOFS.CallOpts, cid)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_GOFS *GOFSTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GOFS.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_GOFS *GOFSSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _GOFS.Contract.ChangeOwner(&_GOFS.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_GOFS *GOFSTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _GOFS.Contract.ChangeOwner(&_GOFS.TransactOpts, newOwner)
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
// Solidity: function pin(bytes cid) returns()
func (_GOFS *GOFSTransactor) Pin(opts *bind.TransactOpts, cid []byte) (*types.Transaction, error) {
	return _GOFS.contract.Transact(opts, "pin", cid)
}

// Pin is a paid mutator transaction binding the contract method 0x7d1962f8.
//
// Solidity: function pin(bytes cid) returns()
func (_GOFS *GOFSSession) Pin(cid []byte) (*types.Transaction, error) {
	return _GOFS.Contract.Pin(&_GOFS.TransactOpts, cid)
}

// Pin is a paid mutator transaction binding the contract method 0x7d1962f8.
//
// Solidity: function pin(bytes cid) returns()
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
