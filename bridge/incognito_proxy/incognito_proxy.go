// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package incognito_proxy

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IncognitoProxyCommittee is an auto generated low-level Go binding around an user-defined struct.
type IncognitoProxyCommittee struct {
	Pubkeys    []common.Address
	StartBlock *big.Int
	SwapID     *big.Int
}

// IncognitoProxyInstructionProof is an auto generated low-level Go binding around an user-defined struct.
type IncognitoProxyInstructionProof struct {
	SigV    []uint8
	Id      *big.Int
	Path    [][32]byte
	BlkData [32]byte
	SigIdx  []*big.Int
	SigR    [][32]byte
	SigS    [][32]byte
}

// IncognitoProxyMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type IncognitoProxyMerkleProof struct {
	Id   *big.Int
	Path [][32]byte
}

// AdminPausableABI is the input ABI used to generate the binding from.
const AdminPausableABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ndays\",\"type\":\"uint256\"}],\"name\":\"Extend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"extend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_successor\",\"type\":\"address\"}],\"name\":\"retire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"successor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AdminPausableFuncSigs maps the 4-byte function signature to its string representation.
var AdminPausableFuncSigs = map[string]string{
	"f851a440": "admin()",
	"4e71d92d": "claim()",
	"79599f96": "expire()",
	"9714378c": "extend(uint256)",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"9e6371ba": "retire(address)",
	"6ff968c3": "successor()",
	"3f4ba83a": "unpause()",
}

// AdminPausableBin is the compiled bytecode used for deploying new contracts.
var AdminPausableBin = "0x608060405234801561001057600080fd5b506040516106fc3803806106fc8339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b03199092169190911790556001805460ff60a01b191690556301e1338042016002556106808061007c6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c806379599f961161006657806379599f96146100ea5780638456cb59146101045780639714378c1461010c5780639e6371ba14610129578063f851a4401461014f57610093565b80633f4ba83a146100985780634e71d92d146100a25780635c975abb146100aa5780636ff968c3146100c6575b600080fd5b6100a0610157565b005b6100a0610239565b6100b2610320565b604080519115158252519081900360200190f35b6100ce610330565b604080516001600160a01b039092168252519081900360200190f35b6100f261033f565b60408051918252519081900360200190f35b6100a0610345565b6100a06004803603602081101561012257600080fd5b503561046a565b6100a06004803603602081101561013f57600080fd5b50356001600160a01b031661058e565b6100ce61063b565b6000546001600160a01b031633146101a2576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b600154600160a01b900460ff166101f7576040805162461bcd60e51b81526020600482015260146024820152736e6f7420706175736564207269676874206e6f7760601b604482015290519081900360640190fd5b6001805460ff60a01b191690556040805133815290517f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa9181900360200190a1565b6002544210610279576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b6001546001600160a01b031633146102c7576040805162461bcd60e51b815260206004820152600c60248201526b1d5b985d5d1a1bdc9a5e995960a21b604482015290519081900360640190fd5b600154600080546001600160a01b0319166001600160a01b0392831617908190556040805191909216815290517f0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc9181900360200190a1565b600154600160a01b900460ff1681565b6001546001600160a01b031681565b60025481565b6000546001600160a01b03163314610390576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b600154600160a01b900460ff16156103e2576040805162461bcd60e51b815260206004820152601060248201526f706175736564207269676874206e6f7760801b604482015290519081900360640190fd5b6002544210610422576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b6001805460ff60a01b1916600160a01b1790556040805133815290517f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2589181900360200190a1565b6000546001600160a01b031633146104b5576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b60025442106104f5576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b61016e811061054b576040805162461bcd60e51b815260206004820152601a60248201527f63616e6e6f7420657874656e6420666f7220746f6f206c6f6e67000000000000604482015290519081900360640190fd5b600280546201518083020190556040805182815290517f02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e89181900360200190a150565b6000546001600160a01b031633146105d9576040805162461bcd60e51b81526020600482015260096024820152683737ba1030b236b4b760b91b604482015290519081900360640190fd5b6002544210610619576040805162461bcd60e51b8152602060048201526007602482015266195e1c1a5c995960ca1b604482015290519081900360640190fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b03168156fea2646970667358221220f07a79af805ec091ff5b9b9d4b8853e8c61f0e8db1f210d302f28d54ee5b38a364736f6c63430006060033"

// DeployAdminPausable deploys a new Ethereum contract, binding an instance of AdminPausable to it.
func DeployAdminPausable(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address) (common.Address, *types.Transaction, *AdminPausable, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminPausableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AdminPausableBin), backend, _admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AdminPausable{AdminPausableCaller: AdminPausableCaller{contract: contract}, AdminPausableTransactor: AdminPausableTransactor{contract: contract}, AdminPausableFilterer: AdminPausableFilterer{contract: contract}}, nil
}

// AdminPausable is an auto generated Go binding around an Ethereum contract.
type AdminPausable struct {
	AdminPausableCaller     // Read-only binding to the contract
	AdminPausableTransactor // Write-only binding to the contract
	AdminPausableFilterer   // Log filterer for contract events
}

// AdminPausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdminPausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminPausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdminPausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminPausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdminPausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminPausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdminPausableSession struct {
	Contract     *AdminPausable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdminPausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdminPausableCallerSession struct {
	Contract *AdminPausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AdminPausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdminPausableTransactorSession struct {
	Contract     *AdminPausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AdminPausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdminPausableRaw struct {
	Contract *AdminPausable // Generic contract binding to access the raw methods on
}

// AdminPausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdminPausableCallerRaw struct {
	Contract *AdminPausableCaller // Generic read-only contract binding to access the raw methods on
}

// AdminPausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdminPausableTransactorRaw struct {
	Contract *AdminPausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdminPausable creates a new instance of AdminPausable, bound to a specific deployed contract.
func NewAdminPausable(address common.Address, backend bind.ContractBackend) (*AdminPausable, error) {
	contract, err := bindAdminPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdminPausable{AdminPausableCaller: AdminPausableCaller{contract: contract}, AdminPausableTransactor: AdminPausableTransactor{contract: contract}, AdminPausableFilterer: AdminPausableFilterer{contract: contract}}, nil
}

// NewAdminPausableCaller creates a new read-only instance of AdminPausable, bound to a specific deployed contract.
func NewAdminPausableCaller(address common.Address, caller bind.ContractCaller) (*AdminPausableCaller, error) {
	contract, err := bindAdminPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdminPausableCaller{contract: contract}, nil
}

// NewAdminPausableTransactor creates a new write-only instance of AdminPausable, bound to a specific deployed contract.
func NewAdminPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*AdminPausableTransactor, error) {
	contract, err := bindAdminPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdminPausableTransactor{contract: contract}, nil
}

// NewAdminPausableFilterer creates a new log filterer instance of AdminPausable, bound to a specific deployed contract.
func NewAdminPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*AdminPausableFilterer, error) {
	contract, err := bindAdminPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdminPausableFilterer{contract: contract}, nil
}

// bindAdminPausable binds a generic wrapper to an already deployed contract.
func bindAdminPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminPausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminPausable *AdminPausableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdminPausable.Contract.AdminPausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminPausable *AdminPausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminPausable.Contract.AdminPausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminPausable *AdminPausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminPausable.Contract.AdminPausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminPausable *AdminPausableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdminPausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminPausable *AdminPausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminPausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminPausable *AdminPausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminPausable.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AdminPausable *AdminPausableCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdminPausable.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AdminPausable *AdminPausableSession) Admin() (common.Address, error) {
	return _AdminPausable.Contract.Admin(&_AdminPausable.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_AdminPausable *AdminPausableCallerSession) Admin() (common.Address, error) {
	return _AdminPausable.Contract.Admin(&_AdminPausable.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_AdminPausable *AdminPausableCaller) Expire(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdminPausable.contract.Call(opts, out, "expire")
	return *ret0, err
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_AdminPausable *AdminPausableSession) Expire() (*big.Int, error) {
	return _AdminPausable.Contract.Expire(&_AdminPausable.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_AdminPausable *AdminPausableCallerSession) Expire() (*big.Int, error) {
	return _AdminPausable.Contract.Expire(&_AdminPausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AdminPausable *AdminPausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AdminPausable.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AdminPausable *AdminPausableSession) Paused() (bool, error) {
	return _AdminPausable.Contract.Paused(&_AdminPausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AdminPausable *AdminPausableCallerSession) Paused() (bool, error) {
	return _AdminPausable.Contract.Paused(&_AdminPausable.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_AdminPausable *AdminPausableCaller) Successor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdminPausable.contract.Call(opts, out, "successor")
	return *ret0, err
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_AdminPausable *AdminPausableSession) Successor() (common.Address, error) {
	return _AdminPausable.Contract.Successor(&_AdminPausable.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_AdminPausable *AdminPausableCallerSession) Successor() (common.Address, error) {
	return _AdminPausable.Contract.Successor(&_AdminPausable.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdminPausable *AdminPausableTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminPausable.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdminPausable *AdminPausableSession) Claim() (*types.Transaction, error) {
	return _AdminPausable.Contract.Claim(&_AdminPausable.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_AdminPausable *AdminPausableTransactorSession) Claim() (*types.Transaction, error) {
	return _AdminPausable.Contract.Claim(&_AdminPausable.TransactOpts)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_AdminPausable *AdminPausableTransactor) Extend(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _AdminPausable.contract.Transact(opts, "extend", n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_AdminPausable *AdminPausableSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _AdminPausable.Contract.Extend(&_AdminPausable.TransactOpts, n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_AdminPausable *AdminPausableTransactorSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _AdminPausable.Contract.Extend(&_AdminPausable.TransactOpts, n)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AdminPausable *AdminPausableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminPausable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AdminPausable *AdminPausableSession) Pause() (*types.Transaction, error) {
	return _AdminPausable.Contract.Pause(&_AdminPausable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AdminPausable *AdminPausableTransactorSession) Pause() (*types.Transaction, error) {
	return _AdminPausable.Contract.Pause(&_AdminPausable.TransactOpts)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_AdminPausable *AdminPausableTransactor) Retire(opts *bind.TransactOpts, _successor common.Address) (*types.Transaction, error) {
	return _AdminPausable.contract.Transact(opts, "retire", _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_AdminPausable *AdminPausableSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _AdminPausable.Contract.Retire(&_AdminPausable.TransactOpts, _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_AdminPausable *AdminPausableTransactorSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _AdminPausable.Contract.Retire(&_AdminPausable.TransactOpts, _successor)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AdminPausable *AdminPausableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminPausable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AdminPausable *AdminPausableSession) Unpause() (*types.Transaction, error) {
	return _AdminPausable.Contract.Unpause(&_AdminPausable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AdminPausable *AdminPausableTransactorSession) Unpause() (*types.Transaction, error) {
	return _AdminPausable.Contract.Unpause(&_AdminPausable.TransactOpts)
}

// AdminPausableClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the AdminPausable contract.
type AdminPausableClaimIterator struct {
	Event *AdminPausableClaim // Event containing the contract specifics and raw log

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
func (it *AdminPausableClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminPausableClaim)
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
		it.Event = new(AdminPausableClaim)
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
func (it *AdminPausableClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminPausableClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminPausableClaim represents a Claim event raised by the AdminPausable contract.
type AdminPausableClaim struct {
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_AdminPausable *AdminPausableFilterer) FilterClaim(opts *bind.FilterOpts) (*AdminPausableClaimIterator, error) {

	logs, sub, err := _AdminPausable.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &AdminPausableClaimIterator{contract: _AdminPausable.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_AdminPausable *AdminPausableFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *AdminPausableClaim) (event.Subscription, error) {

	logs, sub, err := _AdminPausable.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminPausableClaim)
				if err := _AdminPausable.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_AdminPausable *AdminPausableFilterer) ParseClaim(log types.Log) (*AdminPausableClaim, error) {
	event := new(AdminPausableClaim)
	if err := _AdminPausable.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdminPausableExtendIterator is returned from FilterExtend and is used to iterate over the raw logs and unpacked data for Extend events raised by the AdminPausable contract.
type AdminPausableExtendIterator struct {
	Event *AdminPausableExtend // Event containing the contract specifics and raw log

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
func (it *AdminPausableExtendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminPausableExtend)
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
		it.Event = new(AdminPausableExtend)
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
func (it *AdminPausableExtendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminPausableExtendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminPausableExtend represents a Extend event raised by the AdminPausable contract.
type AdminPausableExtend struct {
	Ndays *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExtend is a free log retrieval operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_AdminPausable *AdminPausableFilterer) FilterExtend(opts *bind.FilterOpts) (*AdminPausableExtendIterator, error) {

	logs, sub, err := _AdminPausable.contract.FilterLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return &AdminPausableExtendIterator{contract: _AdminPausable.contract, event: "Extend", logs: logs, sub: sub}, nil
}

// WatchExtend is a free log subscription operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_AdminPausable *AdminPausableFilterer) WatchExtend(opts *bind.WatchOpts, sink chan<- *AdminPausableExtend) (event.Subscription, error) {

	logs, sub, err := _AdminPausable.contract.WatchLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminPausableExtend)
				if err := _AdminPausable.contract.UnpackLog(event, "Extend", log); err != nil {
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

// ParseExtend is a log parse operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_AdminPausable *AdminPausableFilterer) ParseExtend(log types.Log) (*AdminPausableExtend, error) {
	event := new(AdminPausableExtend)
	if err := _AdminPausable.contract.UnpackLog(event, "Extend", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdminPausablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AdminPausable contract.
type AdminPausablePausedIterator struct {
	Event *AdminPausablePaused // Event containing the contract specifics and raw log

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
func (it *AdminPausablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminPausablePaused)
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
		it.Event = new(AdminPausablePaused)
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
func (it *AdminPausablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminPausablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminPausablePaused represents a Paused event raised by the AdminPausable contract.
type AdminPausablePaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_AdminPausable *AdminPausableFilterer) FilterPaused(opts *bind.FilterOpts) (*AdminPausablePausedIterator, error) {

	logs, sub, err := _AdminPausable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AdminPausablePausedIterator{contract: _AdminPausable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_AdminPausable *AdminPausableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AdminPausablePaused) (event.Subscription, error) {

	logs, sub, err := _AdminPausable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminPausablePaused)
				if err := _AdminPausable.contract.UnpackLog(event, "Paused", log); err != nil {
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
// Solidity: event Paused(address pauser)
func (_AdminPausable *AdminPausableFilterer) ParsePaused(log types.Log) (*AdminPausablePaused, error) {
	event := new(AdminPausablePaused)
	if err := _AdminPausable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdminPausableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AdminPausable contract.
type AdminPausableUnpausedIterator struct {
	Event *AdminPausableUnpaused // Event containing the contract specifics and raw log

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
func (it *AdminPausableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminPausableUnpaused)
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
		it.Event = new(AdminPausableUnpaused)
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
func (it *AdminPausableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminPausableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminPausableUnpaused represents a Unpaused event raised by the AdminPausable contract.
type AdminPausableUnpaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_AdminPausable *AdminPausableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AdminPausableUnpausedIterator, error) {

	logs, sub, err := _AdminPausable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AdminPausableUnpausedIterator{contract: _AdminPausable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_AdminPausable *AdminPausableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AdminPausableUnpaused) (event.Subscription, error) {

	logs, sub, err := _AdminPausable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminPausableUnpaused)
				if err := _AdminPausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
// Solidity: event Unpaused(address pauser)
func (_AdminPausable *AdminPausableFilterer) ParseUnpaused(log types.Log) (*AdminPausableUnpaused, error) {
	event := new(AdminPausableUnpaused)
	if err := _AdminPausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyABI is the input ABI used to generate the binding from.
const IncognitoProxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"beaconCommittee\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"bridgeCommittee\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startHeight\",\"type\":\"uint256\"}],\"name\":\"BeaconCommitteeSwapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"}],\"name\":\"CandidatePromoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"}],\"name\":\"ChainFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ndays\",\"type\":\"uint256\"}],\"name\":\"Extend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startHeight\",\"type\":\"uint256\"}],\"name\":\"SubmittedBridgeCandidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"beaconCandidates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beaconBlockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"beaconCommittees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapID\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beaconFinality\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"}],\"name\":\"blockIsFinal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bridgeCandidates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beaconBlockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bridgeCommittees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapID\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeFinality\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockHeight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"}],\"name\":\"calcMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"extend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"numVals\",\"type\":\"uint256\"}],\"name\":\"extractCommitteeFromInstruction\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"extractDataFromBlockMerkleInstruction\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"extractMetaFromInstruction\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"}],\"name\":\"filterSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getBeaconCommittee\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"pubkeys\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapID\",\"type\":\"uint256\"}],\"internalType\":\"structIncognitoProxy.Committee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getBridgeCommittee\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"pubkeys\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapID\",\"type\":\"uint256\"}],\"internalType\":\"structIncognitoProxy.Committee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"}],\"name\":\"getLatestSwapID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"instHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIncognitoProxy.InstructionProof\",\"name\":\"instProof\",\"type\":\"tuple\"}],\"name\":\"instructionApprovedBySigners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"swapID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"}],\"name\":\"loadCandidates\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"swapID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIncognitoProxy.MerkleProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"promoteCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_successor\",\"type\":\"address\"}],\"name\":\"retire\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIncognitoProxy.InstructionProof\",\"name\":\"instProof\",\"type\":\"tuple\"}],\"name\":\"submitBeaconCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIncognitoProxy.InstructionProof[2]\",\"name\":\"instProofs\",\"type\":\"tuple[2]\"}],\"name\":\"submitBridgeCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[2]\",\"name\":\"insts\",\"type\":\"bytes[2]\"},{\"components\":[{\"internalType\":\"uint8[]\",\"name\":\"sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"path\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdx\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigS\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIncognitoProxy.InstructionProof[2]\",\"name\":\"instProofs\",\"type\":\"tuple[2]\"},{\"internalType\":\"uint256\",\"name\":\"swapID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isBeacon\",\"type\":\"bool\"}],\"name\":\"submitFinalityProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"successor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"committee\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"verifySig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// IncognitoProxyFuncSigs maps the 4-byte function signature to its string representation.
var IncognitoProxyFuncSigs = map[string]string{
	"f851a440": "admin()",
	"62532738": "beaconCandidates(uint256)",
	"f203a5ed": "beaconCommittees(uint256)",
	"5d7f86d9": "beaconFinality()",
	"24fb3c62": "blockIsFinal(bool,bytes32,uint256,bytes32[])",
	"e83d31ec": "bridgeCandidates(uint256)",
	"9b30b637": "bridgeCommittees(uint256)",
	"2c1c9cac": "bridgeFinality()",
	"6fd760a8": "calcMerkleRoot(bytes32,uint256,bytes32[])",
	"4e71d92d": "claim()",
	"79599f96": "expire()",
	"9714378c": "extend(uint256)",
	"8eb60066": "extractCommitteeFromInstruction(bytes,uint256)",
	"f179dc78": "extractDataFromBlockMerkleInstruction(bytes)",
	"90500bae": "extractMetaFromInstruction(bytes)",
	"6ae474d8": "filterSigners(uint256[],address[])",
	"faea3167": "getBeaconCommittee(uint256)",
	"8ceb69c3": "getBridgeCommittee(uint256)",
	"722a5d7b": "getLatestSwapID(bool)",
	"fc31d302": "instructionApprovedBySigners(bytes32,address[],(uint8[],uint256,bytes32[],bytes32,uint256[],bytes32[],bytes32[]))",
	"6e845d19": "loadCandidates(uint256,bool)",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"5f551ac3": "promoteCandidate(uint256,bool,(uint256,bytes32[]))",
	"9e6371ba": "retire(address)",
	"14205a4b": "submitBeaconCandidate(bytes,(uint8[],uint256,bytes32[],bytes32,uint256[],bytes32[],bytes32[]))",
	"886cfca3": "submitBridgeCandidate(bytes,(uint8[],uint256,bytes32[],bytes32,uint256[],bytes32[],bytes32[])[2])",
	"a42005f5": "submitFinalityProof(bytes[2],(uint8[],uint256,bytes32[],bytes32,uint256[],bytes32[],bytes32[])[2],uint256,bool)",
	"6ff968c3": "successor()",
	"3f4ba83a": "unpause()",
	"3aacfdad": "verifySig(address[],bytes32,uint8[],bytes32[],bytes32[])",
}

// IncognitoProxyBin is the compiled bytecode used for deploying new contracts.
var IncognitoProxyBin = "0x60806040523480156200001157600080fd5b5060405162002d5638038062002d568339810160408190526200003491620003a3565b600080546001600160a01b0385166001600160a01b03199091161781556001805460ff60a01b191681556301e1338042016002556040805160608101825285815260208082018590529181018490526003805493840181558085528151805192949091027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0192620000cc928492909101906200024f565b5060208281015160018301556040928301516002909201919091558151606081018352858152600081830181905292810184905291805260068152815180517f54cdd369e4e8a8515e52ca72ec816c2101831ad1f18bf44102ed171459c9b4f8926200013d9284929101906200024f565b50602082810151600180840191909155604093840151600290930192909255825160608101845285815260008183018190529381018490526004805493840181559093528251805160039093027f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0192620001bc92849201906200024f565b5060208281015160018301556040928301516002909201919091558151606081018352848152600081830181905292810184905291805260058152815180517f05b8ccbb9d4d8fb16ea74ce3c29a41f1b461fbdaff4714a0d9a8eb05499746bc926200022d9284929101906200024f565b506020820151816001015560408201518160020155905050505050506200041e565b828054828255906000526020600020908101928215620002a7579160200282015b82811115620002a757825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000270565b50620002b5929150620002b9565b5090565b620002e091905b80821115620002b55780546001600160a01b0319168155600101620002c0565b90565b80516001600160a01b0381168114620002fb57600080fd5b92915050565b600082601f83011262000312578081fd5b81516001600160401b038082111562000329578283fd5b60208083026040518282820101818110858211171562000347578687fd5b6040528481529450818501925085820181870183018810156200036957600080fd5b600091505b848210156200039857620003838882620002e3565b8452928201926001919091019082016200036e565b505050505092915050565b600080600060608486031215620003b8578283fd5b620003c48585620002e3565b60208501519093506001600160401b0380821115620003e1578384fd5b620003ef8783880162000301565b9350604086015191508082111562000405578283fd5b50620004148682870162000301565b9150509250925092565b612928806200042e6000396000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c806379599f961161010f5780639e6371ba116100a2578063f203a5ed11610071578063f203a5ed1461040d578063f851a44014610420578063faea316714610428578063fc31d3021461043b576101e5565b80639e6371ba146103b2578063a42005f5146103c5578063e83d31ec146103d8578063f179dc78146103eb576101e5565b80638eb60066116100de5780638eb600661461035557806390500bae146103685780639714378c1461038c5780639b30b6371461039f576101e5565b806379599f96146103125780638456cb591461031a578063886cfca3146103225780638ceb69c314610335576101e5565b80635d7f86d9116101875780636e845d19116101565780636e845d19146102b75780636fd760a8146102ca5780636ff968c3146102ea578063722a5d7b146102ff576101e5565b80635d7f86d9146102695780635f551ac31461027157806362532738146102845780636ae474d814610297576101e5565b80633aacfdad116101c35780633aacfdad1461023e5780633f4ba83a146102515780634e71d92d146102595780635c975abb14610261576101e5565b806314205a4b146101ea57806324fb3c62146101ff5780632c1c9cac14610228575b600080fd5b6101fd6101f8366004612290565b61044e565b005b61021261020d3660046120f4565b61076c565b60405161021f9190612483565b60405180910390f35b61023061079a565b60405161021f92919061244e565b61021261024c366004611f04565b6107a3565b6101fd6108a1565b6101fd61093b565b6102126109d7565b6102306109e7565b6101fd61027f36600461236f565b6109f0565b610230610292366004612328565b610ce9565b6102aa6102a5366004612078565b610d05565b60405161021f9190612470565b6102aa6102c5366004612340565b610e2a565b6102dd6102d83660046121bd565b610f5d565b60405161021f9190612445565b6102f261103e565b60405161021f919061245c565b6102dd61030d3660046120d8565b61104d565b6102dd6110ac565b6101fd6110b2565b6101fd61033036600461223a565b61116a565b610348610343366004612328565b61148a565b60405161021f91906127ef565b6102aa6103633660046122e6565b611535565b61037b610376366004612200565b6115de565b60405161021f959493929190612855565b6101fd61039a366004612328565b611650565b6102306103ad366004612328565b611704565b6101fd6103c0366004611ee9565b611735565b6101fd6103d3366004611fb5565b6117a2565b6102306103e6366004612328565b61199d565b6103fe6103f9366004612200565b6119b9565b60405161021f9392919061283a565b61023061041b366004612328565b6119fc565b6102f2611a09565b610348610436366004612328565b611a18565b610212610449366004612154565b611a2d565b600154600160a01b900460ff16156104815760405162461bcd60e51b815260040161047890612746565b60405180910390fd5b610489611ae4565b610492836115de565b60808601526060850152604084015260ff90811660208401521680825260461480156104c55750806020015160ff166001145b6104ce57600080fd5b825160208401206003805460609160009160001981019081106104ed57fe5b9060005260206000209060030201600201549050808460800151116105245760405162461bcd60e51b815260040161047890612502565b80600101846080015114156105bf576105b8856080015160036001600380549050038154811061055057fe5b600091825260209182902060039091020180546040805182850281018501909152818152928301828280156105ae57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610590575b5050505050610d05565b915061063a565b608080860151908501516000190160009081526006602090815260409182902080548351818402810184019094528084526106379493928301828280156105ae576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610590575050505050610d05565b91505b610645838387611a2d565b61064e57600080fd5b606061065e878660600151611535565b905060006106758588602001518960400151610f5d565b9050600087606001518260405160200161069092919061244e565b604051602081830303815290604052805190602001206040516020016106b69190612445565b60408051601f1981840301815282825280516020918201206060840183528684528a8301518483015283830181905260808b01516000908152600683529290922083518051939550909261070d9284920190611b19565b50602082015160018201556040918201516002909101556003548882015191517fc918d4d247e9af7d4a8ef8ac6f0078262006c87209d062a6d846f5b547dcff6692610759929161244e565b60405180910390a1505050505050505050565b600080851561077e5750600854610783565b50600a545b8061078f868686610f5d565b149695505050505050565b600954600a5482565b600082518451146107b357600080fd5b81518451146107c157600080fd5b60005b8451811015610892578681815181106107d957fe5b60200260200101516001600160a01b03166001878784815181106107f957fe5b602002602001015187858151811061080d57fe5b602002602001015187868151811061082157fe5b602002602001015160405160008152602001604052604051610846949392919061248e565b6020604051602081039080840390855afa158015610868573d6000803e3d6000fd5b505050602060405103516001600160a01b03161461088a576000915050610898565b6001016107c4565b50600190505b95945050505050565b6000546001600160a01b031633146108cb5760405162461bcd60e51b81526004016104789061279f565b600154600160a01b900460ff166108f45760405162461bcd60e51b8152600401610478906124ac565b6001805460ff60a01b191690556040517f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa9061093190339061245c565b60405180910390a1565b600254421061095c5760405162461bcd60e51b81526004016104789061258a565b6001546001600160a01b031633146109865760405162461bcd60e51b815260040161047890612608565b600154600080546001600160a01b0319166001600160a01b0392831617908190556040517f0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc9261093192169061245c565b600154600160a01b900460ff1681565b60075460085482565b600154600160a01b900460ff1615610a1a5760405162461bcd60e51b815260040161047890612746565b60008215610a3a5750600083815260066020526040902060020154610a4e565b506000838152600560205260409020600201545b610a636001828460000151856020015161076c565b610a6c57600080fd5b8215610b9057600380548591906000198101908110610a8757fe5b90600052602060002090600302016002015460010114610ab95760405162461bcd60e51b815260040161047890612547565b60408051600086815260066020908152908390208054608092810284018301909452606083018481526003948493919291840182828015610b2357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610b05575b5050509183525050600087815260066020908152604080832060019081015483860152930189905284549283018555938152839020825180519394600390930290910192610b749284920190611b19565b5060208201518160010155604082015181600201555050610caa565b600480548591906000198101908110610ba557fe5b90600052602060002090600302016002015460010114610bd75760405162461bcd60e51b815260040161047890612547565b60408051600086815260056020908152908390208054608092810284018301909452606083018481526004948493919291840182828015610c4157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610c23575b5050509183525050600087815260056020908152604080832060019081015483860152930189905284549283018555938152839020825180519394600390930290910192610c929284920190611b19565b50602082015181600101556040820151816002015550505b7f40906ca8849ddb84e9284fa5813086f804b1a6e4b236420725c43cbf663b61d58484604051610cdb92919061282a565b60405180910390a150505050565b6006602052600090815260409020600181015460029091015482565b60608082516001600160401b0381118015610d1f57600080fd5b50604051908082528060200260200182016040528015610d49578160200160208202803683370190505b50905060005b8451811015610e2057600081118015610d915750846001820381518110610d7257fe5b6020026020010151858281518110610d8657fe5b602002602001015111155b80610db057508351858281518110610da557fe5b602002602001015110155b15610dcd5760405162461bcd60e51b8152600401610478906124da565b83858281518110610dda57fe5b602002602001015181518110610dec57fe5b6020026020010151828281518110610e0057fe5b6001600160a01b0390921660209283029190910190910152600101610d4f565b5090505b92915050565b60608115610ec757600083815260066020526040902054610e5d5760405162461bcd60e51b815260040161047890612770565b60008381526006602090815260409182902080548351818402810184019094528084529091830182828015610ebb57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610e9d575b50505050509050610e24565b600083815260056020526040902054610ef25760405162461bcd60e51b815260040161047890612717565b60008381526005602090815260409182902080548351818402810184019094528084529091830182828015610f5057602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610f32575b5050505050905092915050565b6000805b8251811015611032576001841615610fb657828181518110610f7f57fe5b602002602001015185604051602001610f9992919061244e565b604051602081830303815290604052805190602001209450611026565b828181518110610fc257fe5b60200260200101516000801b1415610fe7578485604051602001610f9992919061244e565b84838281518110610ff457fe5b602002602001015160405160200161100d92919061244e565b6040516020818303038152906040528051906020012094505b600193841c9301610f61565b508390505b9392505050565b6001546001600160a01b031681565b600081156110805760038054600019810190811061106757fe5b90600052602060002090600302016002015490506110a7565b60048054600019810190811061109257fe5b90600052602060002090600302016002015490505b919050565b60025481565b6000546001600160a01b031633146110dc5760405162461bcd60e51b81526004016104789061279f565b600154600160a01b900460ff16156111065760405162461bcd60e51b815260040161047890612746565b60025442106111275760405162461bcd60e51b81526004016104789061258a565b6001805460ff60a01b1916600160a01b1790556040517f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2589061093190339061245c565b600154600160a01b900460ff16156111945760405162461bcd60e51b815260040161047890612746565b61119c611ae4565b6111a5836115de565b60808601526060850152604084015260ff90811660208401521680825260471480156111d85750806020015160ff166001145b6111e157600080fd5b82516020840120606061120b84600060200201516080015160038054600019810190811061055057fe5b84516020810151604090910151919250600091611229918591610f5d565b905061123e83838760005b6020020151611a2d565b61125a5760405162461bcd60e51b815260040161047890612665565b6004805460009190600019810190811061127057fe5b9060005260206000209060030201600201549050808560800151116112a75760405162461bcd60e51b815260040161047890612502565b80600101856080015114156112de576112d786600160200201516080015160048054600019810190811061055057fe5b925061135a565b60208087015160809081015190870151600019016000908152600583526040908190208054825181860281018601909352808352611357948301828280156105ae576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610590575050505050610d05565b92505b6113678484886001611234565b6113835760405162461bcd60e51b81526004016104789061262e565b611391878660600151611535565b8651606001516040519194506000916113af9190859060200161244e565b604051602081830303815290604052805190602001206040516020016113d59190612445565b60408051601f198184030181528282528051602091820120606084018352878452898301518483015283830181905260808a01516000908152600583529290922083518051939550909261142c9284920190611b19565b50602082015160018201556040918201516002909101556004548782015191517fc918d4d247e9af7d4a8ef8ac6f0078262006c87209d062a6d846f5b547dcff6692611478929161244e565b60405180910390a15050505050505050565b611492611b7e565b6004828154811061149f57fe5b90600052602060002090600302016040518060600160405290816000820180548060200260200160405190810160405280929190818152602001828054801561151157602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116114f3575b50505050508152602001600182015481526020016002820154815250509050919050565b60608160200260620183511461154a57600080fd5b6060826001600160401b038111801561156257600080fd5b5060405190808252806020026020018201604052801561158c578160200160208202803683370190505b5090506000805b848110156115d457602081026082870101519150818382815181106115b457fe5b6001600160a01b0390921660209283029190910190910152600101611593565b5090949350505050565b60008060008060006062865110156115f557600080fd5b60008660008151811061160457fe5b602001015160f81c60f81b60f81c905060008760018151811061162357fe5b0160200151602289015160428a01516062909a0151939a60f89290921c9990985096509194509092505050565b6000546001600160a01b0316331461167a5760405162461bcd60e51b81526004016104789061279f565b600254421061169b5760405162461bcd60e51b81526004016104789061258a565b61016e81106116bc5760405162461bcd60e51b8152600401610478906125d1565b600280546201518083020190556040517f02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8906116f9908390612445565b60405180910390a150565b6004818154811061171157fe5b90600052602060002090600302016000915090508060010154908060020154905082565b6000546001600160a01b0316331461175f5760405162461bcd60e51b81526004016104789061279f565b60025442106117805760405162461bcd60e51b81526004016104789061258a565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b600154600160a01b900460ff16156117cc5760405162461bcd60e51b815260040161047890612746565b6117d58161104d565b8210156117f45760405162461bcd60e51b81526004016104789061269c565b60606118008383610e2a565b905060005b600281101561187d57606061182e86836002811061181f57fe5b60200201516080015184610d05565b905061185887836002811061183f57fe5b6020020151805190602001208288856002811061123457fe5b6118745760405162461bcd60e51b8152600401610478906126eb565b50600101611805565b506000808061189288825b60200201516119b9565b91945092509050600080806118a88b6001611888565b925092509250600a81816118b857fe5b04600a8504600101146118dd5760405162461bcd60e51b8152600401610478906127c2565b8560ff1660491480156118f357508260ff166049145b61190f5760405162461bcd60e51b8152600401610478906125ab565b87156119395760408051808201909152600080825260209091018690526007556008859055611959565b6040805180820190915260008082526020909101869052600955600a8590555b7f045c41374537196136848ce62859f8aa22ae91b490103bd2bed75eac6a4d5180886040516119889190612483565b60405180910390a15050505050505050505050565b6005602052600090815260409020600181015460029091015482565b60008060006041845110156119cd57600080fd5b6000846000815181106119dc57fe5b0160200151602186015160419096015160f89190911c9690945092505050565b6003818154811061171157fe5b6000546001600160a01b031681565b611a20611b7e565b6003828154811061149f57fe5b600080611a438584602001518560400151610f5d565b90506000836060015182604051602001611a5e92919061244e565b60405160208183030381529060405280519060200120604051602001611a849190612445565b6040516020818303038152906040528051906020012090506003855160020281611aaa57fe5b0484600001515111611ac157600092505050611037565b611ada858286600001518760a001518860c001516107a3565b9695505050505050565b6040518060a00160405280600060ff168152602001600060ff1681526020016000815260200160008152602001600081525090565b828054828255906000526020600020908101928215611b6e579160200282015b82811115611b6e57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611b39565b50611b7a929150611b9f565b5090565b60405180606001604052806060815260200160008152602001600081525090565b611bc391905b80821115611b7a5780546001600160a01b0319168155600101611ba5565b90565b80356001600160a01b0381168114610e2457600080fd5b600082601f830112611bed578081fd5b8135611c00611bfb826128a6565b612880565b818152915060208083019084810181840286018201871015611c2157600080fd5b60005b84811015611c4857611c368883611bc6565b84529282019290820190600101611c24565b505050505092915050565b600082601f830112611c63578081fd5b8135611c71611bfb826128a6565b818152915060208083019084810181840286018201871015611c9257600080fd5b60005b84811015611c4857813584529282019290820190600101611c95565b600082601f830112611cc1578081fd5b611ccb6040612880565b9050808260005b6002811015611cfd57611ce88683358701611df7565b83526020928301929190910190600101611cd2565b50505092915050565b600082601f830112611d16578081fd5b8135611d24611bfb826128a6565b818152915060208083019084810181840286018201871015611d4557600080fd5b6000805b85811015611d7357823560ff81168114611d61578283fd5b85529383019391830191600101611d49565b50505050505092915050565b80358015158114610e2457600080fd5b600082601f830112611d9f578081fd5b81356001600160401b03811115611db4578182fd5b611dc7601f8201601f1916602001612880565b9150808252836020828501011115611dde57600080fd5b8060208401602084013760009082016020015292915050565b600060e08284031215611e08578081fd5b611e1260e0612880565b905081356001600160401b0380821115611e2b57600080fd5b611e3785838601611d06565b8352602084013560208401526040840135915080821115611e5757600080fd5b611e6385838601611c53565b6040840152606084013560608401526080840135915080821115611e8657600080fd5b611e9285838601611c53565b608084015260a0840135915080821115611eab57600080fd5b611eb785838601611c53565b60a084015260c0840135915080821115611ed057600080fd5b50611edd84828501611c53565b60c08301525092915050565b600060208284031215611efa578081fd5b6110378383611bc6565b600080600080600060a08688031215611f1b578081fd5b85356001600160401b0380821115611f31578283fd5b611f3d89838a01611bdd565b9650602088013595506040880135915080821115611f59578283fd5b611f6589838a01611d06565b94506060880135915080821115611f7a578283fd5b611f8689838a01611c53565b93506080880135915080821115611f9b578283fd5b50611fa888828901611c53565b9150509295509295909350565b60008060008060808587031215611fca578384fd5b84356001600160401b0380821115611fe0578586fd5b81870188601f820112611ff1578687fd5b60029250612001611bfb846128c5565b8082895b8681101561202f5761201a8d83358701611d8f565b84526020938401939190910190600101612005565b509098505050506020870135915080821115612049578485fd5b5061205687828801611cb1565b9350506040850135915061206d8660608701611d7f565b905092959194509250565b6000806040838503121561208a578182fd5b82356001600160401b03808211156120a0578384fd5b6120ac86838701611c53565b935060208501359150808211156120c1578283fd5b506120ce85828601611bdd565b9150509250929050565b6000602082840312156120e9578081fd5b8135611037816128e1565b60008060008060808587031215612109578182fd5b8435612114816128e1565b9350602085013592506040850135915060608501356001600160401b0381111561213c578182fd5b61214887828801611c53565b91505092959194509250565b600080600060608486031215612168578081fd5b8335925060208401356001600160401b0380821115612185578283fd5b61219187838801611bdd565b935060408601359150808211156121a6578283fd5b506121b386828701611df7565b9150509250925092565b6000806000606084860312156121d1578081fd5b833592506020840135915060408401356001600160401b038111156121f4578182fd5b6121b386828701611c53565b600060208284031215612211578081fd5b81356001600160401b03811115612226578182fd5b61223284828501611d8f565b949350505050565b6000806040838503121561224c578182fd5b82356001600160401b0380821115612262578384fd5b61226e86838701611d8f565b93506020850135915080821115612283578283fd5b506120ce85828601611cb1565b600080604083850312156122a2578182fd5b82356001600160401b03808211156122b8578384fd5b6122c486838701611d8f565b935060208501359150808211156122d9578283fd5b506120ce85828601611df7565b600080604083850312156122f8578182fd5b82356001600160401b0381111561230d578283fd5b61231985828601611d8f565b95602094909401359450505050565b600060208284031215612339578081fd5b5035919050565b60008060408385031215612352578182fd5b823591506020830135612364816128e1565b809150509250929050565b600080600060608486031215612383578081fd5b833592506020840135612395816128e1565b915060408401356001600160401b03808211156123b0578283fd5b818601604081890312156123c2578384fd5b6123cc6040612880565b9250803583526020810135828111156123e3578485fd5b6123ef89828401611c53565b6020850152505050809150509250925092565b6000815180845260208085019450808401835b8381101561243a5781516001600160a01b031687529582019590820190600101612415565b509495945050505050565b90815260200190565b918252602082015260400190565b6001600160a01b0391909116815260200190565b6000602082526110376020830184612402565b901515815260200190565b93845260ff9290921660208401526040830152606082015260800190565b6020808252601490820152736e6f7420706175736564207269676874206e6f7760601b604082015260600190565b6020808252600e908201526d1cda59d2591e081a5b9d985b1a5960921b604082015260600190565b60208082526025908201527f63616e6e6f74207375626d69742063616e64696461746520666f72206f6c6420604082015264737761707360d81b606082015260800190565b60208082526023908201527f6d7573742070726f6d6f74652063616e6469646174652073657175656e7469616040820152626c6c7960e81b606082015260800190565b602080825260079082015266195e1c1a5c995960ca1b604082015260600190565b6020808252600c908201526b696e76616c6964206d65746160a01b604082015260600190565b6020808252601a908201527f63616e6e6f7420657874656e6420666f7220746f6f206c6f6e67000000000000604082015260600190565b6020808252600c908201526b1d5b985d5d1a1bdc9a5e995960a21b604082015260600190565b6020808252601a908201527f696e76616c69642062726964676520696e737472756374696f6e000000000000604082015260600190565b6020808252601a908201527f696e76616c696420626561636f6e20696e737472756374696f6e000000000000604082015260600190565b6020808252602f908201527f70726f6f66206d757374206265207369676e6564206279206e657720636f6d6d60408201526e69747465652f63616e64696461746560881b606082015260800190565b602080825260129082015271696e76616c6964207369676e61747572657360701b604082015260600190565b6020808252601590820152741a5b9d985b1a5908189c9a5919d9481cddd85c1251605a1b604082015260600190565b60208082526010908201526f706175736564207269676874206e6f7760801b604082015260600190565b6020808252601590820152741a5b9d985b1a590818995858dbdb881cddd85c1251605a1b604082015260600190565b6020808252600990820152683737ba1030b236b4b760b91b604082015260600190565b6020808252601390820152721c1c9bdc1bdcd9551a5b59481a5b9d985b1a59606a1b604082015260600190565b60006020825282516060602084015261280b6080840182612402565b6020850151604085015260408501516060850152809250505092915050565b9182521515602082015260400190565b60ff9390931683526020830191909152604082015260600190565b60ff958616815293909416602084015260408301919091526060820152608081019190915260a00190565b6040518181016001600160401b038111828210171561289e57600080fd5b604052919050565b60006001600160401b038211156128bb578081fd5b5060209081020190565b60006001600160401b038211156128da578081fd5b5060200290565b80151581146128ef57600080fd5b5056fea2646970667358221220ffa85bc3e01d86a79ef19a1ca8776f0698988e412283551003c60e392e6f833f64736f6c63430006060033"

// DeployIncognitoProxy deploys a new Ethereum contract, binding an instance of IncognitoProxy to it.
func DeployIncognitoProxy(auth *bind.TransactOpts, backend bind.ContractBackend, admin common.Address, beaconCommittee []common.Address, bridgeCommittee []common.Address) (common.Address, *types.Transaction, *IncognitoProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(IncognitoProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IncognitoProxyBin), backend, admin, beaconCommittee, bridgeCommittee)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IncognitoProxy{IncognitoProxyCaller: IncognitoProxyCaller{contract: contract}, IncognitoProxyTransactor: IncognitoProxyTransactor{contract: contract}, IncognitoProxyFilterer: IncognitoProxyFilterer{contract: contract}}, nil
}

// IncognitoProxy is an auto generated Go binding around an Ethereum contract.
type IncognitoProxy struct {
	IncognitoProxyCaller     // Read-only binding to the contract
	IncognitoProxyTransactor // Write-only binding to the contract
	IncognitoProxyFilterer   // Log filterer for contract events
}

// IncognitoProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IncognitoProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IncognitoProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IncognitoProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IncognitoProxySession struct {
	Contract     *IncognitoProxy   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IncognitoProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IncognitoProxyCallerSession struct {
	Contract *IncognitoProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IncognitoProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IncognitoProxyTransactorSession struct {
	Contract     *IncognitoProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IncognitoProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IncognitoProxyRaw struct {
	Contract *IncognitoProxy // Generic contract binding to access the raw methods on
}

// IncognitoProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IncognitoProxyCallerRaw struct {
	Contract *IncognitoProxyCaller // Generic read-only contract binding to access the raw methods on
}

// IncognitoProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IncognitoProxyTransactorRaw struct {
	Contract *IncognitoProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIncognitoProxy creates a new instance of IncognitoProxy, bound to a specific deployed contract.
func NewIncognitoProxy(address common.Address, backend bind.ContractBackend) (*IncognitoProxy, error) {
	contract, err := bindIncognitoProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IncognitoProxy{IncognitoProxyCaller: IncognitoProxyCaller{contract: contract}, IncognitoProxyTransactor: IncognitoProxyTransactor{contract: contract}, IncognitoProxyFilterer: IncognitoProxyFilterer{contract: contract}}, nil
}

// NewIncognitoProxyCaller creates a new read-only instance of IncognitoProxy, bound to a specific deployed contract.
func NewIncognitoProxyCaller(address common.Address, caller bind.ContractCaller) (*IncognitoProxyCaller, error) {
	contract, err := bindIncognitoProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyCaller{contract: contract}, nil
}

// NewIncognitoProxyTransactor creates a new write-only instance of IncognitoProxy, bound to a specific deployed contract.
func NewIncognitoProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*IncognitoProxyTransactor, error) {
	contract, err := bindIncognitoProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyTransactor{contract: contract}, nil
}

// NewIncognitoProxyFilterer creates a new log filterer instance of IncognitoProxy, bound to a specific deployed contract.
func NewIncognitoProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*IncognitoProxyFilterer, error) {
	contract, err := bindIncognitoProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyFilterer{contract: contract}, nil
}

// bindIncognitoProxy binds a generic wrapper to an already deployed contract.
func bindIncognitoProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IncognitoProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncognitoProxy *IncognitoProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IncognitoProxy.Contract.IncognitoProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncognitoProxy *IncognitoProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.IncognitoProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncognitoProxy *IncognitoProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.IncognitoProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncognitoProxy *IncognitoProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IncognitoProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncognitoProxy *IncognitoProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncognitoProxy *IncognitoProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_IncognitoProxy *IncognitoProxyCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_IncognitoProxy *IncognitoProxySession) Admin() (common.Address, error) {
	return _IncognitoProxy.Contract.Admin(&_IncognitoProxy.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_IncognitoProxy *IncognitoProxyCallerSession) Admin() (common.Address, error) {
	return _IncognitoProxy.Contract.Admin(&_IncognitoProxy.CallOpts)
}

// BeaconCandidates is a free data retrieval call binding the contract method 0x62532738.
//
// Solidity: function beaconCandidates(uint256 ) view returns(uint256 startBlock, bytes32 beaconBlockHash)
func (_IncognitoProxy *IncognitoProxyCaller) BeaconCandidates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartBlock      *big.Int
	BeaconBlockHash [32]byte
}, error) {
	ret := new(struct {
		StartBlock      *big.Int
		BeaconBlockHash [32]byte
	})
	out := ret
	err := _IncognitoProxy.contract.Call(opts, out, "beaconCandidates", arg0)
	return *ret, err
}

// BeaconCandidates is a free data retrieval call binding the contract method 0x62532738.
//
// Solidity: function beaconCandidates(uint256 ) view returns(uint256 startBlock, bytes32 beaconBlockHash)
func (_IncognitoProxy *IncognitoProxySession) BeaconCandidates(arg0 *big.Int) (struct {
	StartBlock      *big.Int
	BeaconBlockHash [32]byte
}, error) {
	return _IncognitoProxy.Contract.BeaconCandidates(&_IncognitoProxy.CallOpts, arg0)
}

// BeaconCandidates is a free data retrieval call binding the contract method 0x62532738.
//
// Solidity: function beaconCandidates(uint256 ) view returns(uint256 startBlock, bytes32 beaconBlockHash)
func (_IncognitoProxy *IncognitoProxyCallerSession) BeaconCandidates(arg0 *big.Int) (struct {
	StartBlock      *big.Int
	BeaconBlockHash [32]byte
}, error) {
	return _IncognitoProxy.Contract.BeaconCandidates(&_IncognitoProxy.CallOpts, arg0)
}

// BeaconCommittees is a free data retrieval call binding the contract method 0xf203a5ed.
//
// Solidity: function beaconCommittees(uint256 ) view returns(uint256 startBlock, uint256 swapID)
func (_IncognitoProxy *IncognitoProxyCaller) BeaconCommittees(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartBlock *big.Int
	SwapID     *big.Int
}, error) {
	ret := new(struct {
		StartBlock *big.Int
		SwapID     *big.Int
	})
	out := ret
	err := _IncognitoProxy.contract.Call(opts, out, "beaconCommittees", arg0)
	return *ret, err
}

// BeaconCommittees is a free data retrieval call binding the contract method 0xf203a5ed.
//
// Solidity: function beaconCommittees(uint256 ) view returns(uint256 startBlock, uint256 swapID)
func (_IncognitoProxy *IncognitoProxySession) BeaconCommittees(arg0 *big.Int) (struct {
	StartBlock *big.Int
	SwapID     *big.Int
}, error) {
	return _IncognitoProxy.Contract.BeaconCommittees(&_IncognitoProxy.CallOpts, arg0)
}

// BeaconCommittees is a free data retrieval call binding the contract method 0xf203a5ed.
//
// Solidity: function beaconCommittees(uint256 ) view returns(uint256 startBlock, uint256 swapID)
func (_IncognitoProxy *IncognitoProxyCallerSession) BeaconCommittees(arg0 *big.Int) (struct {
	StartBlock *big.Int
	SwapID     *big.Int
}, error) {
	return _IncognitoProxy.Contract.BeaconCommittees(&_IncognitoProxy.CallOpts, arg0)
}

// BeaconFinality is a free data retrieval call binding the contract method 0x5d7f86d9.
//
// Solidity: function beaconFinality() view returns(uint256 blockHeight, bytes32 rootHash)
func (_IncognitoProxy *IncognitoProxyCaller) BeaconFinality(opts *bind.CallOpts) (struct {
	BlockHeight *big.Int
	RootHash    [32]byte
}, error) {
	ret := new(struct {
		BlockHeight *big.Int
		RootHash    [32]byte
	})
	out := ret
	err := _IncognitoProxy.contract.Call(opts, out, "beaconFinality")
	return *ret, err
}

// BeaconFinality is a free data retrieval call binding the contract method 0x5d7f86d9.
//
// Solidity: function beaconFinality() view returns(uint256 blockHeight, bytes32 rootHash)
func (_IncognitoProxy *IncognitoProxySession) BeaconFinality() (struct {
	BlockHeight *big.Int
	RootHash    [32]byte
}, error) {
	return _IncognitoProxy.Contract.BeaconFinality(&_IncognitoProxy.CallOpts)
}

// BeaconFinality is a free data retrieval call binding the contract method 0x5d7f86d9.
//
// Solidity: function beaconFinality() view returns(uint256 blockHeight, bytes32 rootHash)
func (_IncognitoProxy *IncognitoProxyCallerSession) BeaconFinality() (struct {
	BlockHeight *big.Int
	RootHash    [32]byte
}, error) {
	return _IncognitoProxy.Contract.BeaconFinality(&_IncognitoProxy.CallOpts)
}

// BlockIsFinal is a free data retrieval call binding the contract method 0x24fb3c62.
//
// Solidity: function blockIsFinal(bool isBeacon, bytes32 blockHash, uint256 blockID, bytes32[] path) view returns(bool)
func (_IncognitoProxy *IncognitoProxyCaller) BlockIsFinal(opts *bind.CallOpts, isBeacon bool, blockHash [32]byte, blockID *big.Int, path [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "blockIsFinal", isBeacon, blockHash, blockID, path)
	return *ret0, err
}

// BlockIsFinal is a free data retrieval call binding the contract method 0x24fb3c62.
//
// Solidity: function blockIsFinal(bool isBeacon, bytes32 blockHash, uint256 blockID, bytes32[] path) view returns(bool)
func (_IncognitoProxy *IncognitoProxySession) BlockIsFinal(isBeacon bool, blockHash [32]byte, blockID *big.Int, path [][32]byte) (bool, error) {
	return _IncognitoProxy.Contract.BlockIsFinal(&_IncognitoProxy.CallOpts, isBeacon, blockHash, blockID, path)
}

// BlockIsFinal is a free data retrieval call binding the contract method 0x24fb3c62.
//
// Solidity: function blockIsFinal(bool isBeacon, bytes32 blockHash, uint256 blockID, bytes32[] path) view returns(bool)
func (_IncognitoProxy *IncognitoProxyCallerSession) BlockIsFinal(isBeacon bool, blockHash [32]byte, blockID *big.Int, path [][32]byte) (bool, error) {
	return _IncognitoProxy.Contract.BlockIsFinal(&_IncognitoProxy.CallOpts, isBeacon, blockHash, blockID, path)
}

// BridgeCandidates is a free data retrieval call binding the contract method 0xe83d31ec.
//
// Solidity: function bridgeCandidates(uint256 ) view returns(uint256 startBlock, bytes32 beaconBlockHash)
func (_IncognitoProxy *IncognitoProxyCaller) BridgeCandidates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartBlock      *big.Int
	BeaconBlockHash [32]byte
}, error) {
	ret := new(struct {
		StartBlock      *big.Int
		BeaconBlockHash [32]byte
	})
	out := ret
	err := _IncognitoProxy.contract.Call(opts, out, "bridgeCandidates", arg0)
	return *ret, err
}

// BridgeCandidates is a free data retrieval call binding the contract method 0xe83d31ec.
//
// Solidity: function bridgeCandidates(uint256 ) view returns(uint256 startBlock, bytes32 beaconBlockHash)
func (_IncognitoProxy *IncognitoProxySession) BridgeCandidates(arg0 *big.Int) (struct {
	StartBlock      *big.Int
	BeaconBlockHash [32]byte
}, error) {
	return _IncognitoProxy.Contract.BridgeCandidates(&_IncognitoProxy.CallOpts, arg0)
}

// BridgeCandidates is a free data retrieval call binding the contract method 0xe83d31ec.
//
// Solidity: function bridgeCandidates(uint256 ) view returns(uint256 startBlock, bytes32 beaconBlockHash)
func (_IncognitoProxy *IncognitoProxyCallerSession) BridgeCandidates(arg0 *big.Int) (struct {
	StartBlock      *big.Int
	BeaconBlockHash [32]byte
}, error) {
	return _IncognitoProxy.Contract.BridgeCandidates(&_IncognitoProxy.CallOpts, arg0)
}

// BridgeCommittees is a free data retrieval call binding the contract method 0x9b30b637.
//
// Solidity: function bridgeCommittees(uint256 ) view returns(uint256 startBlock, uint256 swapID)
func (_IncognitoProxy *IncognitoProxyCaller) BridgeCommittees(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartBlock *big.Int
	SwapID     *big.Int
}, error) {
	ret := new(struct {
		StartBlock *big.Int
		SwapID     *big.Int
	})
	out := ret
	err := _IncognitoProxy.contract.Call(opts, out, "bridgeCommittees", arg0)
	return *ret, err
}

// BridgeCommittees is a free data retrieval call binding the contract method 0x9b30b637.
//
// Solidity: function bridgeCommittees(uint256 ) view returns(uint256 startBlock, uint256 swapID)
func (_IncognitoProxy *IncognitoProxySession) BridgeCommittees(arg0 *big.Int) (struct {
	StartBlock *big.Int
	SwapID     *big.Int
}, error) {
	return _IncognitoProxy.Contract.BridgeCommittees(&_IncognitoProxy.CallOpts, arg0)
}

// BridgeCommittees is a free data retrieval call binding the contract method 0x9b30b637.
//
// Solidity: function bridgeCommittees(uint256 ) view returns(uint256 startBlock, uint256 swapID)
func (_IncognitoProxy *IncognitoProxyCallerSession) BridgeCommittees(arg0 *big.Int) (struct {
	StartBlock *big.Int
	SwapID     *big.Int
}, error) {
	return _IncognitoProxy.Contract.BridgeCommittees(&_IncognitoProxy.CallOpts, arg0)
}

// BridgeFinality is a free data retrieval call binding the contract method 0x2c1c9cac.
//
// Solidity: function bridgeFinality() view returns(uint256 blockHeight, bytes32 rootHash)
func (_IncognitoProxy *IncognitoProxyCaller) BridgeFinality(opts *bind.CallOpts) (struct {
	BlockHeight *big.Int
	RootHash    [32]byte
}, error) {
	ret := new(struct {
		BlockHeight *big.Int
		RootHash    [32]byte
	})
	out := ret
	err := _IncognitoProxy.contract.Call(opts, out, "bridgeFinality")
	return *ret, err
}

// BridgeFinality is a free data retrieval call binding the contract method 0x2c1c9cac.
//
// Solidity: function bridgeFinality() view returns(uint256 blockHeight, bytes32 rootHash)
func (_IncognitoProxy *IncognitoProxySession) BridgeFinality() (struct {
	BlockHeight *big.Int
	RootHash    [32]byte
}, error) {
	return _IncognitoProxy.Contract.BridgeFinality(&_IncognitoProxy.CallOpts)
}

// BridgeFinality is a free data retrieval call binding the contract method 0x2c1c9cac.
//
// Solidity: function bridgeFinality() view returns(uint256 blockHeight, bytes32 rootHash)
func (_IncognitoProxy *IncognitoProxyCallerSession) BridgeFinality() (struct {
	BlockHeight *big.Int
	RootHash    [32]byte
}, error) {
	return _IncognitoProxy.Contract.BridgeFinality(&_IncognitoProxy.CallOpts)
}

// CalcMerkleRoot is a free data retrieval call binding the contract method 0x6fd760a8.
//
// Solidity: function calcMerkleRoot(bytes32 leaf, uint256 id, bytes32[] path) pure returns(bytes32)
func (_IncognitoProxy *IncognitoProxyCaller) CalcMerkleRoot(opts *bind.CallOpts, leaf [32]byte, id *big.Int, path [][32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "calcMerkleRoot", leaf, id, path)
	return *ret0, err
}

// CalcMerkleRoot is a free data retrieval call binding the contract method 0x6fd760a8.
//
// Solidity: function calcMerkleRoot(bytes32 leaf, uint256 id, bytes32[] path) pure returns(bytes32)
func (_IncognitoProxy *IncognitoProxySession) CalcMerkleRoot(leaf [32]byte, id *big.Int, path [][32]byte) ([32]byte, error) {
	return _IncognitoProxy.Contract.CalcMerkleRoot(&_IncognitoProxy.CallOpts, leaf, id, path)
}

// CalcMerkleRoot is a free data retrieval call binding the contract method 0x6fd760a8.
//
// Solidity: function calcMerkleRoot(bytes32 leaf, uint256 id, bytes32[] path) pure returns(bytes32)
func (_IncognitoProxy *IncognitoProxyCallerSession) CalcMerkleRoot(leaf [32]byte, id *big.Int, path [][32]byte) ([32]byte, error) {
	return _IncognitoProxy.Contract.CalcMerkleRoot(&_IncognitoProxy.CallOpts, leaf, id, path)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_IncognitoProxy *IncognitoProxyCaller) Expire(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "expire")
	return *ret0, err
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_IncognitoProxy *IncognitoProxySession) Expire() (*big.Int, error) {
	return _IncognitoProxy.Contract.Expire(&_IncognitoProxy.CallOpts)
}

// Expire is a free data retrieval call binding the contract method 0x79599f96.
//
// Solidity: function expire() view returns(uint256)
func (_IncognitoProxy *IncognitoProxyCallerSession) Expire() (*big.Int, error) {
	return _IncognitoProxy.Contract.Expire(&_IncognitoProxy.CallOpts)
}

// ExtractCommitteeFromInstruction is a free data retrieval call binding the contract method 0x8eb60066.
//
// Solidity: function extractCommitteeFromInstruction(bytes inst, uint256 numVals) pure returns(address[])
func (_IncognitoProxy *IncognitoProxyCaller) ExtractCommitteeFromInstruction(opts *bind.CallOpts, inst []byte, numVals *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "extractCommitteeFromInstruction", inst, numVals)
	return *ret0, err
}

// ExtractCommitteeFromInstruction is a free data retrieval call binding the contract method 0x8eb60066.
//
// Solidity: function extractCommitteeFromInstruction(bytes inst, uint256 numVals) pure returns(address[])
func (_IncognitoProxy *IncognitoProxySession) ExtractCommitteeFromInstruction(inst []byte, numVals *big.Int) ([]common.Address, error) {
	return _IncognitoProxy.Contract.ExtractCommitteeFromInstruction(&_IncognitoProxy.CallOpts, inst, numVals)
}

// ExtractCommitteeFromInstruction is a free data retrieval call binding the contract method 0x8eb60066.
//
// Solidity: function extractCommitteeFromInstruction(bytes inst, uint256 numVals) pure returns(address[])
func (_IncognitoProxy *IncognitoProxyCallerSession) ExtractCommitteeFromInstruction(inst []byte, numVals *big.Int) ([]common.Address, error) {
	return _IncognitoProxy.Contract.ExtractCommitteeFromInstruction(&_IncognitoProxy.CallOpts, inst, numVals)
}

// ExtractDataFromBlockMerkleInstruction is a free data retrieval call binding the contract method 0xf179dc78.
//
// Solidity: function extractDataFromBlockMerkleInstruction(bytes inst) pure returns(uint8, bytes32, uint256)
func (_IncognitoProxy *IncognitoProxyCaller) ExtractDataFromBlockMerkleInstruction(opts *bind.CallOpts, inst []byte) (uint8, [32]byte, *big.Int, error) {
	var (
		ret0 = new(uint8)
		ret1 = new([32]byte)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _IncognitoProxy.contract.Call(opts, out, "extractDataFromBlockMerkleInstruction", inst)
	return *ret0, *ret1, *ret2, err
}

// ExtractDataFromBlockMerkleInstruction is a free data retrieval call binding the contract method 0xf179dc78.
//
// Solidity: function extractDataFromBlockMerkleInstruction(bytes inst) pure returns(uint8, bytes32, uint256)
func (_IncognitoProxy *IncognitoProxySession) ExtractDataFromBlockMerkleInstruction(inst []byte) (uint8, [32]byte, *big.Int, error) {
	return _IncognitoProxy.Contract.ExtractDataFromBlockMerkleInstruction(&_IncognitoProxy.CallOpts, inst)
}

// ExtractDataFromBlockMerkleInstruction is a free data retrieval call binding the contract method 0xf179dc78.
//
// Solidity: function extractDataFromBlockMerkleInstruction(bytes inst) pure returns(uint8, bytes32, uint256)
func (_IncognitoProxy *IncognitoProxyCallerSession) ExtractDataFromBlockMerkleInstruction(inst []byte) (uint8, [32]byte, *big.Int, error) {
	return _IncognitoProxy.Contract.ExtractDataFromBlockMerkleInstruction(&_IncognitoProxy.CallOpts, inst)
}

// ExtractMetaFromInstruction is a free data retrieval call binding the contract method 0x90500bae.
//
// Solidity: function extractMetaFromInstruction(bytes inst) pure returns(uint8, uint8, uint256, uint256, uint256)
func (_IncognitoProxy *IncognitoProxyCaller) ExtractMetaFromInstruction(opts *bind.CallOpts, inst []byte) (uint8, uint8, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(uint8)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _IncognitoProxy.contract.Call(opts, out, "extractMetaFromInstruction", inst)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// ExtractMetaFromInstruction is a free data retrieval call binding the contract method 0x90500bae.
//
// Solidity: function extractMetaFromInstruction(bytes inst) pure returns(uint8, uint8, uint256, uint256, uint256)
func (_IncognitoProxy *IncognitoProxySession) ExtractMetaFromInstruction(inst []byte) (uint8, uint8, *big.Int, *big.Int, *big.Int, error) {
	return _IncognitoProxy.Contract.ExtractMetaFromInstruction(&_IncognitoProxy.CallOpts, inst)
}

// ExtractMetaFromInstruction is a free data retrieval call binding the contract method 0x90500bae.
//
// Solidity: function extractMetaFromInstruction(bytes inst) pure returns(uint8, uint8, uint256, uint256, uint256)
func (_IncognitoProxy *IncognitoProxyCallerSession) ExtractMetaFromInstruction(inst []byte) (uint8, uint8, *big.Int, *big.Int, *big.Int, error) {
	return _IncognitoProxy.Contract.ExtractMetaFromInstruction(&_IncognitoProxy.CallOpts, inst)
}

// FilterSigners is a free data retrieval call binding the contract method 0x6ae474d8.
//
// Solidity: function filterSigners(uint256[] sigIdx, address[] signers) pure returns(address[])
func (_IncognitoProxy *IncognitoProxyCaller) FilterSigners(opts *bind.CallOpts, sigIdx []*big.Int, signers []common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "filterSigners", sigIdx, signers)
	return *ret0, err
}

// FilterSigners is a free data retrieval call binding the contract method 0x6ae474d8.
//
// Solidity: function filterSigners(uint256[] sigIdx, address[] signers) pure returns(address[])
func (_IncognitoProxy *IncognitoProxySession) FilterSigners(sigIdx []*big.Int, signers []common.Address) ([]common.Address, error) {
	return _IncognitoProxy.Contract.FilterSigners(&_IncognitoProxy.CallOpts, sigIdx, signers)
}

// FilterSigners is a free data retrieval call binding the contract method 0x6ae474d8.
//
// Solidity: function filterSigners(uint256[] sigIdx, address[] signers) pure returns(address[])
func (_IncognitoProxy *IncognitoProxyCallerSession) FilterSigners(sigIdx []*big.Int, signers []common.Address) ([]common.Address, error) {
	return _IncognitoProxy.Contract.FilterSigners(&_IncognitoProxy.CallOpts, sigIdx, signers)
}

// GetBeaconCommittee is a free data retrieval call binding the contract method 0xfaea3167.
//
// Solidity: function getBeaconCommittee(uint256 i) view returns((address[],uint256,uint256))
func (_IncognitoProxy *IncognitoProxyCaller) GetBeaconCommittee(opts *bind.CallOpts, i *big.Int) (IncognitoProxyCommittee, error) {
	var (
		ret0 = new(IncognitoProxyCommittee)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "getBeaconCommittee", i)
	return *ret0, err
}

// GetBeaconCommittee is a free data retrieval call binding the contract method 0xfaea3167.
//
// Solidity: function getBeaconCommittee(uint256 i) view returns((address[],uint256,uint256))
func (_IncognitoProxy *IncognitoProxySession) GetBeaconCommittee(i *big.Int) (IncognitoProxyCommittee, error) {
	return _IncognitoProxy.Contract.GetBeaconCommittee(&_IncognitoProxy.CallOpts, i)
}

// GetBeaconCommittee is a free data retrieval call binding the contract method 0xfaea3167.
//
// Solidity: function getBeaconCommittee(uint256 i) view returns((address[],uint256,uint256))
func (_IncognitoProxy *IncognitoProxyCallerSession) GetBeaconCommittee(i *big.Int) (IncognitoProxyCommittee, error) {
	return _IncognitoProxy.Contract.GetBeaconCommittee(&_IncognitoProxy.CallOpts, i)
}

// GetBridgeCommittee is a free data retrieval call binding the contract method 0x8ceb69c3.
//
// Solidity: function getBridgeCommittee(uint256 i) view returns((address[],uint256,uint256))
func (_IncognitoProxy *IncognitoProxyCaller) GetBridgeCommittee(opts *bind.CallOpts, i *big.Int) (IncognitoProxyCommittee, error) {
	var (
		ret0 = new(IncognitoProxyCommittee)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "getBridgeCommittee", i)
	return *ret0, err
}

// GetBridgeCommittee is a free data retrieval call binding the contract method 0x8ceb69c3.
//
// Solidity: function getBridgeCommittee(uint256 i) view returns((address[],uint256,uint256))
func (_IncognitoProxy *IncognitoProxySession) GetBridgeCommittee(i *big.Int) (IncognitoProxyCommittee, error) {
	return _IncognitoProxy.Contract.GetBridgeCommittee(&_IncognitoProxy.CallOpts, i)
}

// GetBridgeCommittee is a free data retrieval call binding the contract method 0x8ceb69c3.
//
// Solidity: function getBridgeCommittee(uint256 i) view returns((address[],uint256,uint256))
func (_IncognitoProxy *IncognitoProxyCallerSession) GetBridgeCommittee(i *big.Int) (IncognitoProxyCommittee, error) {
	return _IncognitoProxy.Contract.GetBridgeCommittee(&_IncognitoProxy.CallOpts, i)
}

// InstructionApprovedBySigners is a free data retrieval call binding the contract method 0xfc31d302.
//
// Solidity: function instructionApprovedBySigners(bytes32 instHash, address[] signers, (uint8[],uint256,bytes32[],bytes32,uint256[],bytes32[],bytes32[]) instProof) view returns(bool)
func (_IncognitoProxy *IncognitoProxyCaller) InstructionApprovedBySigners(opts *bind.CallOpts, instHash [32]byte, signers []common.Address, instProof IncognitoProxyInstructionProof) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "instructionApprovedBySigners", instHash, signers, instProof)
	return *ret0, err
}

// InstructionApprovedBySigners is a free data retrieval call binding the contract method 0xfc31d302.
//
// Solidity: function instructionApprovedBySigners(bytes32 instHash, address[] signers, (uint8[],uint256,bytes32[],bytes32,uint256[],bytes32[],bytes32[]) instProof) view returns(bool)
func (_IncognitoProxy *IncognitoProxySession) InstructionApprovedBySigners(instHash [32]byte, signers []common.Address, instProof IncognitoProxyInstructionProof) (bool, error) {
	return _IncognitoProxy.Contract.InstructionApprovedBySigners(&_IncognitoProxy.CallOpts, instHash, signers, instProof)
}

// InstructionApprovedBySigners is a free data retrieval call binding the contract method 0xfc31d302.
//
// Solidity: function instructionApprovedBySigners(bytes32 instHash, address[] signers, (uint8[],uint256,bytes32[],bytes32,uint256[],bytes32[],bytes32[]) instProof) view returns(bool)
func (_IncognitoProxy *IncognitoProxyCallerSession) InstructionApprovedBySigners(instHash [32]byte, signers []common.Address, instProof IncognitoProxyInstructionProof) (bool, error) {
	return _IncognitoProxy.Contract.InstructionApprovedBySigners(&_IncognitoProxy.CallOpts, instHash, signers, instProof)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IncognitoProxy *IncognitoProxyCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IncognitoProxy *IncognitoProxySession) Paused() (bool, error) {
	return _IncognitoProxy.Contract.Paused(&_IncognitoProxy.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IncognitoProxy *IncognitoProxyCallerSession) Paused() (bool, error) {
	return _IncognitoProxy.Contract.Paused(&_IncognitoProxy.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_IncognitoProxy *IncognitoProxyCaller) Successor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "successor")
	return *ret0, err
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_IncognitoProxy *IncognitoProxySession) Successor() (common.Address, error) {
	return _IncognitoProxy.Contract.Successor(&_IncognitoProxy.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() view returns(address)
func (_IncognitoProxy *IncognitoProxyCallerSession) Successor() (common.Address, error) {
	return _IncognitoProxy.Contract.Successor(&_IncognitoProxy.CallOpts)
}

// VerifySig is a free data retrieval call binding the contract method 0x3aacfdad.
//
// Solidity: function verifySig(address[] committee, bytes32 msgHash, uint8[] v, bytes32[] r, bytes32[] s) pure returns(bool)
func (_IncognitoProxy *IncognitoProxyCaller) VerifySig(opts *bind.CallOpts, committee []common.Address, msgHash [32]byte, v []uint8, r [][32]byte, s [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IncognitoProxy.contract.Call(opts, out, "verifySig", committee, msgHash, v, r, s)
	return *ret0, err
}

// VerifySig is a free data retrieval call binding the contract method 0x3aacfdad.
//
// Solidity: function verifySig(address[] committee, bytes32 msgHash, uint8[] v, bytes32[] r, bytes32[] s) pure returns(bool)
func (_IncognitoProxy *IncognitoProxySession) VerifySig(committee []common.Address, msgHash [32]byte, v []uint8, r [][32]byte, s [][32]byte) (bool, error) {
	return _IncognitoProxy.Contract.VerifySig(&_IncognitoProxy.CallOpts, committee, msgHash, v, r, s)
}

// VerifySig is a free data retrieval call binding the contract method 0x3aacfdad.
//
// Solidity: function verifySig(address[] committee, bytes32 msgHash, uint8[] v, bytes32[] r, bytes32[] s) pure returns(bool)
func (_IncognitoProxy *IncognitoProxyCallerSession) VerifySig(committee []common.Address, msgHash [32]byte, v []uint8, r [][32]byte, s [][32]byte) (bool, error) {
	return _IncognitoProxy.Contract.VerifySig(&_IncognitoProxy.CallOpts, committee, msgHash, v, r, s)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_IncognitoProxy *IncognitoProxyTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_IncognitoProxy *IncognitoProxySession) Claim() (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Claim(&_IncognitoProxy.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) Claim() (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Claim(&_IncognitoProxy.TransactOpts)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_IncognitoProxy *IncognitoProxyTransactor) Extend(opts *bind.TransactOpts, n *big.Int) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "extend", n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_IncognitoProxy *IncognitoProxySession) Extend(n *big.Int) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Extend(&_IncognitoProxy.TransactOpts, n)
}

// Extend is a paid mutator transaction binding the contract method 0x9714378c.
//
// Solidity: function extend(uint256 n) returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) Extend(n *big.Int) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Extend(&_IncognitoProxy.TransactOpts, n)
}

// GetLatestSwapID is a paid mutator transaction binding the contract method 0x722a5d7b.
//
// Solidity: function getLatestSwapID(bool isBeacon) returns(uint256)
func (_IncognitoProxy *IncognitoProxyTransactor) GetLatestSwapID(opts *bind.TransactOpts, isBeacon bool) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "getLatestSwapID", isBeacon)
}

// GetLatestSwapID is a paid mutator transaction binding the contract method 0x722a5d7b.
//
// Solidity: function getLatestSwapID(bool isBeacon) returns(uint256)
func (_IncognitoProxy *IncognitoProxySession) GetLatestSwapID(isBeacon bool) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.GetLatestSwapID(&_IncognitoProxy.TransactOpts, isBeacon)
}

// GetLatestSwapID is a paid mutator transaction binding the contract method 0x722a5d7b.
//
// Solidity: function getLatestSwapID(bool isBeacon) returns(uint256)
func (_IncognitoProxy *IncognitoProxyTransactorSession) GetLatestSwapID(isBeacon bool) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.GetLatestSwapID(&_IncognitoProxy.TransactOpts, isBeacon)
}

// LoadCandidates is a paid mutator transaction binding the contract method 0x6e845d19.
//
// Solidity: function loadCandidates(uint256 swapID, bool isBeacon) returns(address[])
func (_IncognitoProxy *IncognitoProxyTransactor) LoadCandidates(opts *bind.TransactOpts, swapID *big.Int, isBeacon bool) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "loadCandidates", swapID, isBeacon)
}

// LoadCandidates is a paid mutator transaction binding the contract method 0x6e845d19.
//
// Solidity: function loadCandidates(uint256 swapID, bool isBeacon) returns(address[])
func (_IncognitoProxy *IncognitoProxySession) LoadCandidates(swapID *big.Int, isBeacon bool) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.LoadCandidates(&_IncognitoProxy.TransactOpts, swapID, isBeacon)
}

// LoadCandidates is a paid mutator transaction binding the contract method 0x6e845d19.
//
// Solidity: function loadCandidates(uint256 swapID, bool isBeacon) returns(address[])
func (_IncognitoProxy *IncognitoProxyTransactorSession) LoadCandidates(swapID *big.Int, isBeacon bool) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.LoadCandidates(&_IncognitoProxy.TransactOpts, swapID, isBeacon)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IncognitoProxy *IncognitoProxyTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IncognitoProxy *IncognitoProxySession) Pause() (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Pause(&_IncognitoProxy.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) Pause() (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Pause(&_IncognitoProxy.TransactOpts)
}

// PromoteCandidate is a paid mutator transaction binding the contract method 0x5f551ac3.
//
// Solidity: function promoteCandidate(uint256 swapID, bool isBeacon, (uint256,bytes32[]) proof) returns()
func (_IncognitoProxy *IncognitoProxyTransactor) PromoteCandidate(opts *bind.TransactOpts, swapID *big.Int, isBeacon bool, proof IncognitoProxyMerkleProof) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "promoteCandidate", swapID, isBeacon, proof)
}

// PromoteCandidate is a paid mutator transaction binding the contract method 0x5f551ac3.
//
// Solidity: function promoteCandidate(uint256 swapID, bool isBeacon, (uint256,bytes32[]) proof) returns()
func (_IncognitoProxy *IncognitoProxySession) PromoteCandidate(swapID *big.Int, isBeacon bool, proof IncognitoProxyMerkleProof) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.PromoteCandidate(&_IncognitoProxy.TransactOpts, swapID, isBeacon, proof)
}

// PromoteCandidate is a paid mutator transaction binding the contract method 0x5f551ac3.
//
// Solidity: function promoteCandidate(uint256 swapID, bool isBeacon, (uint256,bytes32[]) proof) returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) PromoteCandidate(swapID *big.Int, isBeacon bool, proof IncognitoProxyMerkleProof) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.PromoteCandidate(&_IncognitoProxy.TransactOpts, swapID, isBeacon, proof)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_IncognitoProxy *IncognitoProxyTransactor) Retire(opts *bind.TransactOpts, _successor common.Address) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "retire", _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_IncognitoProxy *IncognitoProxySession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Retire(&_IncognitoProxy.TransactOpts, _successor)
}

// Retire is a paid mutator transaction binding the contract method 0x9e6371ba.
//
// Solidity: function retire(address _successor) returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) Retire(_successor common.Address) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Retire(&_IncognitoProxy.TransactOpts, _successor)
}

// SubmitBeaconCandidate is a paid mutator transaction binding the contract method 0x14205a4b.
//
// Solidity: function submitBeaconCandidate(bytes inst, (uint8[],uint256,bytes32[],bytes32,uint256[],bytes32[],bytes32[]) instProof) returns()
func (_IncognitoProxy *IncognitoProxyTransactor) SubmitBeaconCandidate(opts *bind.TransactOpts, inst []byte, instProof IncognitoProxyInstructionProof) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "submitBeaconCandidate", inst, instProof)
}

// SubmitBeaconCandidate is a paid mutator transaction binding the contract method 0x14205a4b.
//
// Solidity: function submitBeaconCandidate(bytes inst, (uint8[],uint256,bytes32[],bytes32,uint256[],bytes32[],bytes32[]) instProof) returns()
func (_IncognitoProxy *IncognitoProxySession) SubmitBeaconCandidate(inst []byte, instProof IncognitoProxyInstructionProof) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.SubmitBeaconCandidate(&_IncognitoProxy.TransactOpts, inst, instProof)
}

// SubmitBeaconCandidate is a paid mutator transaction binding the contract method 0x14205a4b.
//
// Solidity: function submitBeaconCandidate(bytes inst, (uint8[],uint256,bytes32[],bytes32,uint256[],bytes32[],bytes32[]) instProof) returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) SubmitBeaconCandidate(inst []byte, instProof IncognitoProxyInstructionProof) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.SubmitBeaconCandidate(&_IncognitoProxy.TransactOpts, inst, instProof)
}

// SubmitBridgeCandidate is a paid mutator transaction binding the contract method 0x886cfca3.
//
// Solidity: function submitBridgeCandidate(bytes inst, (uint8[],uint256,bytes32[],bytes32,uint256[],bytes32[],bytes32[])[2] instProofs) returns()
func (_IncognitoProxy *IncognitoProxyTransactor) SubmitBridgeCandidate(opts *bind.TransactOpts, inst []byte, instProofs [2]IncognitoProxyInstructionProof) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "submitBridgeCandidate", inst, instProofs)
}

// SubmitBridgeCandidate is a paid mutator transaction binding the contract method 0x886cfca3.
//
// Solidity: function submitBridgeCandidate(bytes inst, (uint8[],uint256,bytes32[],bytes32,uint256[],bytes32[],bytes32[])[2] instProofs) returns()
func (_IncognitoProxy *IncognitoProxySession) SubmitBridgeCandidate(inst []byte, instProofs [2]IncognitoProxyInstructionProof) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.SubmitBridgeCandidate(&_IncognitoProxy.TransactOpts, inst, instProofs)
}

// SubmitBridgeCandidate is a paid mutator transaction binding the contract method 0x886cfca3.
//
// Solidity: function submitBridgeCandidate(bytes inst, (uint8[],uint256,bytes32[],bytes32,uint256[],bytes32[],bytes32[])[2] instProofs) returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) SubmitBridgeCandidate(inst []byte, instProofs [2]IncognitoProxyInstructionProof) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.SubmitBridgeCandidate(&_IncognitoProxy.TransactOpts, inst, instProofs)
}

// SubmitFinalityProof is a paid mutator transaction binding the contract method 0xa42005f5.
//
// Solidity: function submitFinalityProof(bytes[2] insts, (uint8[],uint256,bytes32[],bytes32,uint256[],bytes32[],bytes32[])[2] instProofs, uint256 swapID, bool isBeacon) returns()
func (_IncognitoProxy *IncognitoProxyTransactor) SubmitFinalityProof(opts *bind.TransactOpts, insts [2][]byte, instProofs [2]IncognitoProxyInstructionProof, swapID *big.Int, isBeacon bool) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "submitFinalityProof", insts, instProofs, swapID, isBeacon)
}

// SubmitFinalityProof is a paid mutator transaction binding the contract method 0xa42005f5.
//
// Solidity: function submitFinalityProof(bytes[2] insts, (uint8[],uint256,bytes32[],bytes32,uint256[],bytes32[],bytes32[])[2] instProofs, uint256 swapID, bool isBeacon) returns()
func (_IncognitoProxy *IncognitoProxySession) SubmitFinalityProof(insts [2][]byte, instProofs [2]IncognitoProxyInstructionProof, swapID *big.Int, isBeacon bool) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.SubmitFinalityProof(&_IncognitoProxy.TransactOpts, insts, instProofs, swapID, isBeacon)
}

// SubmitFinalityProof is a paid mutator transaction binding the contract method 0xa42005f5.
//
// Solidity: function submitFinalityProof(bytes[2] insts, (uint8[],uint256,bytes32[],bytes32,uint256[],bytes32[],bytes32[])[2] instProofs, uint256 swapID, bool isBeacon) returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) SubmitFinalityProof(insts [2][]byte, instProofs [2]IncognitoProxyInstructionProof, swapID *big.Int, isBeacon bool) (*types.Transaction, error) {
	return _IncognitoProxy.Contract.SubmitFinalityProof(&_IncognitoProxy.TransactOpts, insts, instProofs, swapID, isBeacon)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IncognitoProxy *IncognitoProxyTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncognitoProxy.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IncognitoProxy *IncognitoProxySession) Unpause() (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Unpause(&_IncognitoProxy.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IncognitoProxy *IncognitoProxyTransactorSession) Unpause() (*types.Transaction, error) {
	return _IncognitoProxy.Contract.Unpause(&_IncognitoProxy.TransactOpts)
}

// IncognitoProxyBeaconCommitteeSwappedIterator is returned from FilterBeaconCommitteeSwapped and is used to iterate over the raw logs and unpacked data for BeaconCommitteeSwapped events raised by the IncognitoProxy contract.
type IncognitoProxyBeaconCommitteeSwappedIterator struct {
	Event *IncognitoProxyBeaconCommitteeSwapped // Event containing the contract specifics and raw log

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
func (it *IncognitoProxyBeaconCommitteeSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyBeaconCommitteeSwapped)
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
		it.Event = new(IncognitoProxyBeaconCommitteeSwapped)
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
func (it *IncognitoProxyBeaconCommitteeSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyBeaconCommitteeSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyBeaconCommitteeSwapped represents a BeaconCommitteeSwapped event raised by the IncognitoProxy contract.
type IncognitoProxyBeaconCommitteeSwapped struct {
	Id          *big.Int
	StartHeight *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBeaconCommitteeSwapped is a free log retrieval operation binding the contract event 0xe15e1a9dec6ad906dd5985b062bfa5ee8bc5d5738e46e4deb8a2df2fbbbb59d1.
//
// Solidity: event BeaconCommitteeSwapped(uint256 id, uint256 startHeight)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterBeaconCommitteeSwapped(opts *bind.FilterOpts) (*IncognitoProxyBeaconCommitteeSwappedIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "BeaconCommitteeSwapped")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyBeaconCommitteeSwappedIterator{contract: _IncognitoProxy.contract, event: "BeaconCommitteeSwapped", logs: logs, sub: sub}, nil
}

// WatchBeaconCommitteeSwapped is a free log subscription operation binding the contract event 0xe15e1a9dec6ad906dd5985b062bfa5ee8bc5d5738e46e4deb8a2df2fbbbb59d1.
//
// Solidity: event BeaconCommitteeSwapped(uint256 id, uint256 startHeight)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchBeaconCommitteeSwapped(opts *bind.WatchOpts, sink chan<- *IncognitoProxyBeaconCommitteeSwapped) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "BeaconCommitteeSwapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyBeaconCommitteeSwapped)
				if err := _IncognitoProxy.contract.UnpackLog(event, "BeaconCommitteeSwapped", log); err != nil {
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

// ParseBeaconCommitteeSwapped is a log parse operation binding the contract event 0xe15e1a9dec6ad906dd5985b062bfa5ee8bc5d5738e46e4deb8a2df2fbbbb59d1.
//
// Solidity: event BeaconCommitteeSwapped(uint256 id, uint256 startHeight)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseBeaconCommitteeSwapped(log types.Log) (*IncognitoProxyBeaconCommitteeSwapped, error) {
	event := new(IncognitoProxyBeaconCommitteeSwapped)
	if err := _IncognitoProxy.contract.UnpackLog(event, "BeaconCommitteeSwapped", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyCandidatePromotedIterator is returned from FilterCandidatePromoted and is used to iterate over the raw logs and unpacked data for CandidatePromoted events raised by the IncognitoProxy contract.
type IncognitoProxyCandidatePromotedIterator struct {
	Event *IncognitoProxyCandidatePromoted // Event containing the contract specifics and raw log

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
func (it *IncognitoProxyCandidatePromotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyCandidatePromoted)
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
		it.Event = new(IncognitoProxyCandidatePromoted)
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
func (it *IncognitoProxyCandidatePromotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyCandidatePromotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyCandidatePromoted represents a CandidatePromoted event raised by the IncognitoProxy contract.
type IncognitoProxyCandidatePromoted struct {
	SwapID   *big.Int
	IsBeacon bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCandidatePromoted is a free log retrieval operation binding the contract event 0x40906ca8849ddb84e9284fa5813086f804b1a6e4b236420725c43cbf663b61d5.
//
// Solidity: event CandidatePromoted(uint256 swapID, bool isBeacon)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterCandidatePromoted(opts *bind.FilterOpts) (*IncognitoProxyCandidatePromotedIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "CandidatePromoted")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyCandidatePromotedIterator{contract: _IncognitoProxy.contract, event: "CandidatePromoted", logs: logs, sub: sub}, nil
}

// WatchCandidatePromoted is a free log subscription operation binding the contract event 0x40906ca8849ddb84e9284fa5813086f804b1a6e4b236420725c43cbf663b61d5.
//
// Solidity: event CandidatePromoted(uint256 swapID, bool isBeacon)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchCandidatePromoted(opts *bind.WatchOpts, sink chan<- *IncognitoProxyCandidatePromoted) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "CandidatePromoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyCandidatePromoted)
				if err := _IncognitoProxy.contract.UnpackLog(event, "CandidatePromoted", log); err != nil {
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

// ParseCandidatePromoted is a log parse operation binding the contract event 0x40906ca8849ddb84e9284fa5813086f804b1a6e4b236420725c43cbf663b61d5.
//
// Solidity: event CandidatePromoted(uint256 swapID, bool isBeacon)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseCandidatePromoted(log types.Log) (*IncognitoProxyCandidatePromoted, error) {
	event := new(IncognitoProxyCandidatePromoted)
	if err := _IncognitoProxy.contract.UnpackLog(event, "CandidatePromoted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyChainFinalizedIterator is returned from FilterChainFinalized and is used to iterate over the raw logs and unpacked data for ChainFinalized events raised by the IncognitoProxy contract.
type IncognitoProxyChainFinalizedIterator struct {
	Event *IncognitoProxyChainFinalized // Event containing the contract specifics and raw log

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
func (it *IncognitoProxyChainFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyChainFinalized)
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
		it.Event = new(IncognitoProxyChainFinalized)
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
func (it *IncognitoProxyChainFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyChainFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyChainFinalized represents a ChainFinalized event raised by the IncognitoProxy contract.
type IncognitoProxyChainFinalized struct {
	IsBeacon bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChainFinalized is a free log retrieval operation binding the contract event 0x045c41374537196136848ce62859f8aa22ae91b490103bd2bed75eac6a4d5180.
//
// Solidity: event ChainFinalized(bool isBeacon)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterChainFinalized(opts *bind.FilterOpts) (*IncognitoProxyChainFinalizedIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "ChainFinalized")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyChainFinalizedIterator{contract: _IncognitoProxy.contract, event: "ChainFinalized", logs: logs, sub: sub}, nil
}

// WatchChainFinalized is a free log subscription operation binding the contract event 0x045c41374537196136848ce62859f8aa22ae91b490103bd2bed75eac6a4d5180.
//
// Solidity: event ChainFinalized(bool isBeacon)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchChainFinalized(opts *bind.WatchOpts, sink chan<- *IncognitoProxyChainFinalized) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "ChainFinalized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyChainFinalized)
				if err := _IncognitoProxy.contract.UnpackLog(event, "ChainFinalized", log); err != nil {
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

// ParseChainFinalized is a log parse operation binding the contract event 0x045c41374537196136848ce62859f8aa22ae91b490103bd2bed75eac6a4d5180.
//
// Solidity: event ChainFinalized(bool isBeacon)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseChainFinalized(log types.Log) (*IncognitoProxyChainFinalized, error) {
	event := new(IncognitoProxyChainFinalized)
	if err := _IncognitoProxy.contract.UnpackLog(event, "ChainFinalized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the IncognitoProxy contract.
type IncognitoProxyClaimIterator struct {
	Event *IncognitoProxyClaim // Event containing the contract specifics and raw log

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
func (it *IncognitoProxyClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyClaim)
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
		it.Event = new(IncognitoProxyClaim)
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
func (it *IncognitoProxyClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyClaim represents a Claim event raised by the IncognitoProxy contract.
type IncognitoProxyClaim struct {
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterClaim(opts *bind.FilterOpts) (*IncognitoProxyClaimIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyClaimIterator{contract: _IncognitoProxy.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *IncognitoProxyClaim) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyClaim)
				if err := _IncognitoProxy.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address claimer)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseClaim(log types.Log) (*IncognitoProxyClaim, error) {
	event := new(IncognitoProxyClaim)
	if err := _IncognitoProxy.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyExtendIterator is returned from FilterExtend and is used to iterate over the raw logs and unpacked data for Extend events raised by the IncognitoProxy contract.
type IncognitoProxyExtendIterator struct {
	Event *IncognitoProxyExtend // Event containing the contract specifics and raw log

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
func (it *IncognitoProxyExtendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyExtend)
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
		it.Event = new(IncognitoProxyExtend)
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
func (it *IncognitoProxyExtendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyExtendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyExtend represents a Extend event raised by the IncognitoProxy contract.
type IncognitoProxyExtend struct {
	Ndays *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExtend is a free log retrieval operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterExtend(opts *bind.FilterOpts) (*IncognitoProxyExtendIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyExtendIterator{contract: _IncognitoProxy.contract, event: "Extend", logs: logs, sub: sub}, nil
}

// WatchExtend is a free log subscription operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchExtend(opts *bind.WatchOpts, sink chan<- *IncognitoProxyExtend) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "Extend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyExtend)
				if err := _IncognitoProxy.contract.UnpackLog(event, "Extend", log); err != nil {
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

// ParseExtend is a log parse operation binding the contract event 0x02ef6561d311451dadc920679eb21192a61d96ee8ead94241b8ff073029ca6e8.
//
// Solidity: event Extend(uint256 ndays)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseExtend(log types.Log) (*IncognitoProxyExtend, error) {
	event := new(IncognitoProxyExtend)
	if err := _IncognitoProxy.contract.UnpackLog(event, "Extend", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the IncognitoProxy contract.
type IncognitoProxyPausedIterator struct {
	Event *IncognitoProxyPaused // Event containing the contract specifics and raw log

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
func (it *IncognitoProxyPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyPaused)
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
		it.Event = new(IncognitoProxyPaused)
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
func (it *IncognitoProxyPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyPaused represents a Paused event raised by the IncognitoProxy contract.
type IncognitoProxyPaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterPaused(opts *bind.FilterOpts) (*IncognitoProxyPausedIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyPausedIterator{contract: _IncognitoProxy.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address pauser)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *IncognitoProxyPaused) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyPaused)
				if err := _IncognitoProxy.contract.UnpackLog(event, "Paused", log); err != nil {
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
// Solidity: event Paused(address pauser)
func (_IncognitoProxy *IncognitoProxyFilterer) ParsePaused(log types.Log) (*IncognitoProxyPaused, error) {
	event := new(IncognitoProxyPaused)
	if err := _IncognitoProxy.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxySubmittedBridgeCandidateIterator is returned from FilterSubmittedBridgeCandidate and is used to iterate over the raw logs and unpacked data for SubmittedBridgeCandidate events raised by the IncognitoProxy contract.
type IncognitoProxySubmittedBridgeCandidateIterator struct {
	Event *IncognitoProxySubmittedBridgeCandidate // Event containing the contract specifics and raw log

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
func (it *IncognitoProxySubmittedBridgeCandidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxySubmittedBridgeCandidate)
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
		it.Event = new(IncognitoProxySubmittedBridgeCandidate)
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
func (it *IncognitoProxySubmittedBridgeCandidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxySubmittedBridgeCandidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxySubmittedBridgeCandidate represents a SubmittedBridgeCandidate event raised by the IncognitoProxy contract.
type IncognitoProxySubmittedBridgeCandidate struct {
	Id          *big.Int
	StartHeight *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSubmittedBridgeCandidate is a free log retrieval operation binding the contract event 0xc918d4d247e9af7d4a8ef8ac6f0078262006c87209d062a6d846f5b547dcff66.
//
// Solidity: event SubmittedBridgeCandidate(uint256 id, uint256 startHeight)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterSubmittedBridgeCandidate(opts *bind.FilterOpts) (*IncognitoProxySubmittedBridgeCandidateIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "SubmittedBridgeCandidate")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxySubmittedBridgeCandidateIterator{contract: _IncognitoProxy.contract, event: "SubmittedBridgeCandidate", logs: logs, sub: sub}, nil
}

// WatchSubmittedBridgeCandidate is a free log subscription operation binding the contract event 0xc918d4d247e9af7d4a8ef8ac6f0078262006c87209d062a6d846f5b547dcff66.
//
// Solidity: event SubmittedBridgeCandidate(uint256 id, uint256 startHeight)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchSubmittedBridgeCandidate(opts *bind.WatchOpts, sink chan<- *IncognitoProxySubmittedBridgeCandidate) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "SubmittedBridgeCandidate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxySubmittedBridgeCandidate)
				if err := _IncognitoProxy.contract.UnpackLog(event, "SubmittedBridgeCandidate", log); err != nil {
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

// ParseSubmittedBridgeCandidate is a log parse operation binding the contract event 0xc918d4d247e9af7d4a8ef8ac6f0078262006c87209d062a6d846f5b547dcff66.
//
// Solidity: event SubmittedBridgeCandidate(uint256 id, uint256 startHeight)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseSubmittedBridgeCandidate(log types.Log) (*IncognitoProxySubmittedBridgeCandidate, error) {
	event := new(IncognitoProxySubmittedBridgeCandidate)
	if err := _IncognitoProxy.contract.UnpackLog(event, "SubmittedBridgeCandidate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IncognitoProxyUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the IncognitoProxy contract.
type IncognitoProxyUnpausedIterator struct {
	Event *IncognitoProxyUnpaused // Event containing the contract specifics and raw log

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
func (it *IncognitoProxyUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncognitoProxyUnpaused)
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
		it.Event = new(IncognitoProxyUnpaused)
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
func (it *IncognitoProxyUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncognitoProxyUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncognitoProxyUnpaused represents a Unpaused event raised by the IncognitoProxy contract.
type IncognitoProxyUnpaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_IncognitoProxy *IncognitoProxyFilterer) FilterUnpaused(opts *bind.FilterOpts) (*IncognitoProxyUnpausedIterator, error) {

	logs, sub, err := _IncognitoProxy.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &IncognitoProxyUnpausedIterator{contract: _IncognitoProxy.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address pauser)
func (_IncognitoProxy *IncognitoProxyFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *IncognitoProxyUnpaused) (event.Subscription, error) {

	logs, sub, err := _IncognitoProxy.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncognitoProxyUnpaused)
				if err := _IncognitoProxy.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
// Solidity: event Unpaused(address pauser)
func (_IncognitoProxy *IncognitoProxyFilterer) ParseUnpaused(log types.Log) (*IncognitoProxyUnpaused, error) {
	event := new(IncognitoProxyUnpaused)
	if err := _IncognitoProxy.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}
