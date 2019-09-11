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
const GOFSABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"cid\",\"type\":\"bytes\"}],\"name\":\"newWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"setRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cid\",\"type\":\"bytes\"}],\"name\":\"pin\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cid\",\"type\":\"bytes\"}],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cidByHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deployed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"cid\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"bh\",\"type\":\"uint256\"}],\"name\":\"Pinned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"cid\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"CreatedWallet\",\"type\":\"event\"}]"

// GOFSBin is the compiled bytecode used for deploying new contracts.
const GOFSBin = `0x608060405234801561001057600080fd5b5061101f806100206000396000f3fe6080604052600436106100ae576000357c0100000000000000000000000000000000000000000000000000000000900480638da5cb5b116100765780638da5cb5b14610292578063a6f9dae1146102c3578063c5bf2249146102f6578063e16cf225146103a9578063f905c15a14610448576100ae565b806328c6fa6f146100b35780632c4e722e1461016857806334fcf4371461018f57806351cff8d9146101b95780637d1962f8146101ec575b600080fd5b3480156100bf57600080fd5b50610166600480360360208110156100d657600080fd5b8101906020810181356401000000008111156100f157600080fd5b82018360208201111561010357600080fd5b8035906020019184600183028401116401000000008311171561012557600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061045d945050505050565b005b34801561017457600080fd5b5061017d61075e565b60408051918252519081900360200190f35b34801561019b57600080fd5b50610166600480360360208110156101b257600080fd5b5035610764565b3480156101c557600080fd5b50610166600480360360208110156101dc57600080fd5b5035600160a060020a03166107c8565b6101666004803603602081101561020257600080fd5b81019060208101813564010000000081111561021d57600080fd5b82018360208201111561022f57600080fd5b8035906020019184600183028401116401000000008311171561025157600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061084f945050505050565b34801561029e57600080fd5b506102a76109ae565b60408051600160a060020a039092168252519081900360200190f35b3480156102cf57600080fd5b50610166600480360360208110156102e657600080fd5b5035600160a060020a03166109e4565b34801561030257600080fd5b506102a76004803603602081101561031957600080fd5b81019060208101813564010000000081111561033457600080fd5b82018360208201111561034657600080fd5b8035906020019184600183028401116401000000008311171561036857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a6a945050505050565b3480156103b557600080fd5b506103d3600480360360208110156103cc57600080fd5b5035610add565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561040d5781810151838201526020016103f5565b50505050905090810190601f16801561043a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561045457600080fd5b5061017d610b78565b61046681610b7e565b15156104bc576040805160e560020a62461bcd02815260206004820152601a60248201527f56657273696f6e203020434944206e6f7420616c6c6f7765642e000000000000604482015290519081900360640190fd5b6000600160a060020a03166002826040518082805190602001908083835b602083106104f95780518252601f1990920191602091820191016104da565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054600160a060020a031692909214915061058a9050576040805160e560020a62461bcd02815260206004820152601d60248201527f57616c6c657420616c72656164792065786973747320666f7220636964000000604482015290519081900360640190fd5b6000308260405161059a90610c9c565b600160a060020a0383168152604060208083018281528451928401929092528351606084019185019080838360005b838110156105e15781810151838201526020016105c9565b50505050905090810190601f16801561060e5780820380516001836020036101000a031916815260200191505b509350505050604051809103906000f080158015610630573d6000803e3d6000fd5b509050806002836040518082805190602001908083835b602083106106665780518252601f199092019160209182019101610647565b51815160209384036101000a60001901801990921691161790529201948552506040519384900381018420805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0396909616959095179094555050835184928291908401908083835b602083106106ec5780518252601f1990920191602091820191016106cd565b51815160209384036101000a600019018019909216911617905260408051929094018290038220600160a060020a038816835293519395503394507f89c3649b91d0d77a0655fa3d84b050b21c775ba31bfbc37e440ec3ee04f289279391829003019150a361075a82610c44565b5050565b60005481565b61076c6109ae565b600160a060020a031633146107b55760405160e560020a62461bcd028152600401808060200182810382526022815260200180610fd26022913960400191505060405180910390fd5b60015415156107c357436001555b600055565b6107d06109ae565b600160a060020a031633146108195760405160e560020a62461bcd028152600401808060200182810382526022815260200180610fd26022913960400191505060405180910390fd5b604051600160a060020a03821690303180156108fc02916000818181858888f1935050505015801561075a573d6000803e3d6000fd5b61085881610b7e565b15156108ae576040805160e560020a62461bcd02815260206004820152601a60248201527f56657273696f6e203020434944206e6f7420616c6c6f7765642e000000000000604482015290519081900360640190fd5b600054341015610908576040805160e560020a62461bcd02815260206004820152601a60248201527f43616e6e6f7420707572636861736520302073746f726167652e000000000000604482015290519081900360640190fd5b600080543481151561091657fe5b049050816040518082805190602001908083835b602083106109495780518252601f19909201916020918201910161092a565b51815160209384036101000a60001901801990921691161790526040805192909401829003822087835293519395503294507f7c80eb99758dfe3e8aef5df583c1c9bab5369cf46b930b802f130edcfeac90ca9391829003019150a361075a82610c44565b604080517f676f636861696e2e70726f78792e6f776e657200000000000000000000000000815290519081900360130190205490565b6109ec6109ae565b600160a060020a03163314610a355760405160e560020a62461bcd028152600401808060200182810382526022815260200180610fd26022913960400191505060405180910390fd5b604080517f676f636861696e2e70726f78792e6f776e6572000000000000000000000000008152905190819003601301902055565b60006002826040518082805190602001908083835b60208310610a9e5780518252601f199092019160209182019101610a7f565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922054600160a060020a0316925050505b919050565b60036020908152600091825260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015610b705780601f10610b4557610100808354040283529160200191610b70565b820191906000526020600020905b815481529060010190602001808311610b5357829003601f168201915b505050505081565b60015481565b600081516022148015610bda5750816000815181101515610b9b57fe5b90602001015160f860020a900460f860020a027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916601260f860020a02145b8015610c2f5750816001815181101515610bf057fe5b90602001015160f860020a900460f860020a027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916602060f860020a02145b15610c3c57506000610ad8565b506001919050565b80516020808301919091206000818152600390925260409091205460026000196101006001841615020190911604151561075a5760008181526003602090815260409091208351610c9792850190610ca9565b505050565b61028d80610d4583390190565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610cea57805160ff1916838001178555610d17565b82800160010185558215610d17579182015b82811115610d17578251825591602001919060010190610cfc565b50610d23929150610d27565b5090565b610d4191905b80821115610d235760008155600101610d2d565b9056fe608060405234801561001057600080fd5b5060405161028d38038061028d8339810180604052604081101561003357600080fd5b81516020830180519193928301929164010000000081111561005457600080fd5b8201602081018481111561006757600080fd5b815164010000000081118282018710171561008157600080fd5b505060008054600160a060020a031916600160a060020a03871617905580519093506100b692506001915060208401906100be565b505050610159565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100ff57805160ff191683800117855561012c565b8280016001018555821561012c579182015b8281111561012c578251825591602001919060010190610111565b5061013892915061013c565b5090565b61015691905b808211156101385760008155600101610142565b90565b610125806101686000396000f3fe608060408190526000547f7d1962f800000000000000000000000000000000000000000000000000000000909152602060849081526001805460026101008284161502600019019091160460a481905273ffffffffffffffffffffffffffffffffffffffff90931692637d1962f892349291819060c4908490801560c15780601f1060975761010080835404028352916020019160c1565b820191906000526020600020905b81548152906001019060200180831160a557829003601f168201915b5050925050506000604051808303818588803b15801560df57600080fd5b505af115801560f2573d6000803e3d6000fd5b505050505000fea165627a7a72305820ba6a11b392088b6ea800bdd92da14bf7345769669f458dc2a9c0aefc430a1a2e00294f6e6c79206f776e65722063616e2063616c6c20746869732066756e6374696f6e2ea165627a7a72305820dfa34cb8a695316331588acb1584a3bfa5c59b1af89c77c9af761acbbd3486250029`

// DeployGOFS deploys a new GoChain contract, binding an instance of GOFS to it.
func DeployGOFS(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GOFS, error) {
	parsed, err := abi.JSON(strings.NewReader(GOFSABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GOFSBin), backend)
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address addr)
func (_GOFS *GOFSCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GOFS.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address addr)
func (_GOFS *GOFSSession) Owner() (common.Address, error) {
	return _GOFS.Contract.Owner(&_GOFS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address addr)
func (_GOFS *GOFSCallerSession) Owner() (common.Address, error) {
	return _GOFS.Contract.Owner(&_GOFS.CallOpts)
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
