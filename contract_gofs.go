// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gofs

import (
	"math/big"
	"strings"

	"github.com/gochain/gochain/v3"
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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GOFSABI is the input ABI used to generate the binding from.
const GOFSABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"cid\",\"type\":\"bytes\"}],\"name\":\"newWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"setRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cid\",\"type\":\"bytes\"}],\"name\":\"pin\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cid\",\"type\":\"bytes\"}],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cidByHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deployed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_rate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"cid\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"bh\",\"type\":\"uint256\"}],\"name\":\"Pinned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"cid\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"CreatedWallet\",\"type\":\"event\"}]"

// GOFSBin is the compiled bytecode used for deploying new contracts.
const GOFSBin = `0x608060405234801561001057600080fd5b506040516020806110098339810180604052602081101561003057600080fd5b505160008054600160a060020a0319163317905560015543600255610faf8061005a6000396000f3fe6080604052600436106100a3576000357c0100000000000000000000000000000000000000000000000000000000900480637d1962f8116100765780637d1962f8146101e1578063a6f9dae114610287578063c5bf2249146102ba578063e16cf22514610389578063f905c15a14610428576100a3565b806328c6fa6f146100a85780632c4e722e1461015d57806334fcf4371461018457806351cff8d9146101ae575b600080fd5b3480156100b457600080fd5b5061015b600480360360208110156100cb57600080fd5b8101906020810181356401000000008111156100e657600080fd5b8201836020820111156100f857600080fd5b8035906020019184600183028401116401000000008311171561011a57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061043d945050505050565b005b34801561016957600080fd5b506101726107a6565b60408051918252519081900360200190f35b34801561019057600080fd5b5061015b600480360360208110156101a757600080fd5b50356107ac565b3480156101ba57600080fd5b5061015b600480360360208110156101d157600080fd5b5035600160a060020a03166107fd565b61015b600480360360208110156101f757600080fd5b81019060208101813564010000000081111561021257600080fd5b82018360208201111561022457600080fd5b8035906020019184600183028401116401000000008311171561024657600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061087f945050505050565b34801561029357600080fd5b5061015b600480360360208110156102aa57600080fd5b5035600160a060020a0316610a47565b3480156102c657600080fd5b5061036d600480360360208110156102dd57600080fd5b8101906020810181356401000000008111156102f857600080fd5b82018360208201111561030a57600080fd5b8035906020019184600183028401116401000000008311171561032c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610ac2945050505050565b60408051600160a060020a039092168252519081900360200190f35b34801561039557600080fd5b506103b3600480360360208110156103ac57600080fd5b5035610b33565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103ed5781810151838201526020016103d5565b50505050905090810190601f16801561041a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561043457600080fd5b50610172610bce565b80600081518110151561044c57fe5b90602001015160f860020a900460f860020a02600160f860020a031916601260f860020a021480156104af575080600181518110151561048857fe5b90602001015160f860020a900460f860020a02600160f860020a031916602060f860020a02145b15610504576040805160e560020a62461bcd02815260206004820152601a60248201527f56657273696f6e203020434944206e6f7420616c6c6f7765642e000000000000604482015290519081900360640190fd5b6000600160a060020a03166003826040518082805190602001908083835b602083106105415780518252601f199092019160209182019101610522565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054600160a060020a03169290921491506105d29050576040805160e560020a62461bcd02815260206004820152601d60248201527f57616c6c657420616c72656164792065786973747320666f7220636964000000604482015290519081900360640190fd5b600030826040516105e290610c2c565b600160a060020a0383168152604060208083018281528451928401929092528351606084019185019080838360005b83811015610629578181015183820152602001610611565b50505050905090810190601f1680156106565780820380516001836020036101000a031916815260200191505b509350505050604051809103906000f080158015610678573d6000803e3d6000fd5b509050806003836040518082805190602001908083835b602083106106ae5780518252601f19909201916020918201910161068f565b51815160209384036101000a60001901801990921691161790529201948552506040519384900381018420805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0396909616959095179094555050835184928291908401908083835b602083106107345780518252601f199092019160209182019101610715565b51815160209384036101000a600019018019909216911617905260408051929094018290038220600160a060020a038816835293519395503394507f89c3649b91d0d77a0655fa3d84b050b21c775ba31bfbc37e440ec3ee04f289279391829003019150a36107a282610bd4565b5050565b60015481565b600054600160a060020a031633146107f85760405160e560020a62461bcd028152600401808060200182810382526022815260200180610f626022913960400191505060405180910390fd5b600155565b600054600160a060020a031633146108495760405160e560020a62461bcd028152600401808060200182810382526022815260200180610f626022913960400191505060405180910390fd5b604051600160a060020a03821690303180156108fc02916000818181858888f193505050501580156107a2573d6000803e3d6000fd5b80600081518110151561088e57fe5b90602001015160f860020a900460f860020a02600160f860020a031916601260f860020a021480156108f157508060018151811015156108ca57fe5b90602001015160f860020a900460f860020a02600160f860020a031916602060f860020a02145b15610946576040805160e560020a62461bcd02815260206004820152601a60248201527f56657273696f6e203020434944206e6f7420616c6c6f7765642e000000000000604482015290519081900360640190fd5b6001543410156109a0576040805160e560020a62461bcd02815260206004820152601a60248201527f43616e6e6f7420707572636861736520302073746f726167652e000000000000604482015290519081900360640190fd5b6000600154348115156109af57fe5b049050816040518082805190602001908083835b602083106109e25780518252601f1990920191602091820191016109c3565b51815160209384036101000a60001901801990921691161790526040805192909401829003822087835293519395503294507f7c80eb99758dfe3e8aef5df583c1c9bab5369cf46b930b802f130edcfeac90ca9391829003019150a36107a282610bd4565b600054600160a060020a03163314610a935760405160e560020a62461bcd028152600401808060200182810382526022815260200180610f626022913960400191505060405180910390fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60006003826040518082805190602001908083835b60208310610af65780518252601f199092019160209182019101610ad7565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054600160a060020a0316949350505050565b60046020908152600091825260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015610bc65780601f10610b9b57610100808354040283529160200191610bc6565b820191906000526020600020905b815481529060010190602001808311610ba957829003601f168201915b505050505081565b60025481565b8051602080830191909120600081815260049092526040909120546002600019610100600184161502019091160415156107a25760008181526004602090815260409091208351610c2792850190610c39565b505050565b61028d80610cd583390190565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610c7a57805160ff1916838001178555610ca7565b82800160010185558215610ca7579182015b82811115610ca7578251825591602001919060010190610c8c565b50610cb3929150610cb7565b5090565b610cd191905b80821115610cb35760008155600101610cbd565b9056fe608060405234801561001057600080fd5b5060405161028d38038061028d8339810180604052604081101561003357600080fd5b81516020830180519193928301929164010000000081111561005457600080fd5b8201602081018481111561006757600080fd5b815164010000000081118282018710171561008157600080fd5b505060008054600160a060020a031916600160a060020a03871617905580519093506100b692506001915060208401906100be565b505050610159565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100ff57805160ff191683800117855561012c565b8280016001018555821561012c579182015b8281111561012c578251825591602001919060010190610111565b5061013892915061013c565b5090565b61015691905b808211156101385760008155600101610142565b90565b610125806101686000396000f3fe608060408190526000547f7d1962f800000000000000000000000000000000000000000000000000000000909152602060849081526001805460026101008284161502600019019091160460a481905273ffffffffffffffffffffffffffffffffffffffff90931692637d1962f892349291819060c4908490801560c15780601f1060975761010080835404028352916020019160c1565b820191906000526020600020905b81548152906001019060200180831160a557829003601f168201915b5050925050506000604051808303818588803b15801560df57600080fd5b505af115801560f2573d6000803e3d6000fd5b505050505000fea165627a7a7230582082709d4c19652becd227633b5b850cdfdb0ccbf1fef97ccc36d62b1733d2f61b00294f6e6c79206f776e65722063616e2063616c6c20746869732066756e6374696f6e2ea165627a7a7230582071491c6e542b583e114033a8c863faed790c9a1b1d4e7621ac93ddbdb24713940029`

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

// CidByHash is a free data retrieval call binding the contract method 0xe16cf225.
//
// Solidity: function cidByHash(bytes32 ) constant returns(bytes)
func (_GOFS *GOFSCaller) CidByHash(opts *bind.CallOpts, arg0 [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _GOFS.contract.Call(opts, out, "cidByHash", arg0)
	return *ret0, err
}

// CidByHash is a free data retrieval call binding the contract method 0xe16cf225.
//
// Solidity: function cidByHash(bytes32 ) constant returns(bytes)
func (_GOFS *GOFSSession) CidByHash(arg0 [32]byte) ([]byte, error) {
	return _GOFS.Contract.CidByHash(&_GOFS.CallOpts, arg0)
}

// CidByHash is a free data retrieval call binding the contract method 0xe16cf225.
//
// Solidity: function cidByHash(bytes32 ) constant returns(bytes)
func (_GOFS *GOFSCallerSession) CidByHash(arg0 [32]byte) ([]byte, error) {
	return _GOFS.Contract.CidByHash(&_GOFS.CallOpts, arg0)
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
