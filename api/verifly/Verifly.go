// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// VeriflyData is an auto generated low-level Go binding around an user-defined struct.
type VeriflyData struct {
	FullName   string
	CourseName string
	CreatedAt  string
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"fullName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"courseName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"createdAt\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structVerifly.Data\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"Added\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_fullName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_courseName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_createdAt\",\"type\":\"string\"}],\"name\":\"add\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"ipfs\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"fullName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"courseName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"createdAt\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"viewData\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"fullName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"courseName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"createdAt\",\"type\":\"string\"}],\"internalType\":\"structVerifly.Data\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600080546001600160a01b03191633179055610b4b806100326000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806336d9d69214610051578063878fc1081461007a5780638da5cb5b1461009a578063a30e6db5146100c5575b600080fd5b61006461005f3660046106f5565b6100e7565b604051610071919061077a565b60405180910390f35b61008d6100883660046107db565b6102f1565b6040516100719190610888565b6000546100ad906001600160a01b031681565b6040516001600160a01b039091168152602001610071565b6100d86100d33660046106f5565b61048d565b604051610071939291906108a2565b61010b60405180606001604052806060815260200160608152602001606081525090565b60018260405161011b91906108e5565b908152602001604051809103902060405180606001604052908160008201805461014490610901565b80601f016020809104026020016040519081016040528092919081815260200182805461017090610901565b80156101bd5780601f10610192576101008083540402835291602001916101bd565b820191906000526020600020905b8154815290600101906020018083116101a057829003601f168201915b505050505081526020016001820180546101d690610901565b80601f016020809104026020016040519081016040528092919081815260200182805461020290610901565b801561024f5780601f106102245761010080835404028352916020019161024f565b820191906000526020600020905b81548152906001019060200180831161023257829003601f168201915b5050505050815260200160028201805461026890610901565b80601f016020809104026020016040519081016040528092919081815260200182805461029490610901565b80156102e15780601f106102b6576101008083540402835291602001916102e1565b820191906000526020600020905b8154815290600101906020018083116102c457829003601f168201915b5050505050815250509050919050565b6000546060906001600160a01b031633146103435760405162461bcd60e51b815260206004820152600e60248201526d139bdd08105d5d1a1bdc9a5e995960921b604482015260640160405180910390fd5b600060018660405161035591906108e5565b908152604051908190036020019020805461036f90610901565b9050111561039e575060408051808201909152600a8152694461746120457869747360b01b6020820152610485565b6040518060600160405280858152602001848152602001838152506001866040516103c991906108e5565b908152604051908190036020019020815181906103e6908261098a565b50602082015160018201906103fb908261098a565b5060408201516002820190610410908261098a565b509050507fec036d0cecfe4858346b9712e64b4ca7c2c49534a29b5d22db1b19049a77296260018660405161044591906108e5565b90815260405190819003602001812061045d91610ac7565b60405180910390a1506040805180820190915260078152665375636365737360c81b60208201525b949350505050565b80516020818301810180516001825292820191909301209152805481906104b390610901565b80601f01602080910402602001604051908101604052809291908181526020018280546104df90610901565b801561052c5780601f106105015761010080835404028352916020019161052c565b820191906000526020600020905b81548152906001019060200180831161050f57829003601f168201915b50505050509080600101805461054190610901565b80601f016020809104026020016040519081016040528092919081815260200182805461056d90610901565b80156105ba5780601f1061058f576101008083540402835291602001916105ba565b820191906000526020600020905b81548152906001019060200180831161059d57829003601f168201915b5050505050908060020180546105cf90610901565b80601f01602080910402602001604051908101604052809291908181526020018280546105fb90610901565b80156106485780601f1061061d57610100808354040283529160200191610648565b820191906000526020600020905b81548152906001019060200180831161062b57829003601f168201915b5050505050905083565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261067957600080fd5b813567ffffffffffffffff8082111561069457610694610652565b604051601f8301601f19908116603f011681019082821181831017156106bc576106bc610652565b816040528381528660208588010111156106d557600080fd5b836020870160208301376000602085830101528094505050505092915050565b60006020828403121561070757600080fd5b813567ffffffffffffffff81111561071e57600080fd5b61048584828501610668565b60005b8381101561074557818101518382015260200161072d565b50506000910152565b6000815180845261076681602086016020860161072a565b601f01601f19169290920160200192915050565b602081526000825160606020840152610796608084018261074e565b90506020840151601f19808584030160408601526107b4838361074e565b92506040860151915080858403016060860152506107d2828261074e565b95945050505050565b600080600080608085870312156107f157600080fd5b843567ffffffffffffffff8082111561080957600080fd5b61081588838901610668565b9550602087013591508082111561082b57600080fd5b61083788838901610668565b9450604087013591508082111561084d57600080fd5b61085988838901610668565b9350606087013591508082111561086f57600080fd5b5061087c87828801610668565b91505092959194509250565b60208152600061089b602083018461074e565b9392505050565b6060815260006108b5606083018661074e565b82810360208401526108c7818661074e565b905082810360408401526108db818561074e565b9695505050505050565b600082516108f781846020870161072a565b9190910192915050565b600181811c9082168061091557607f821691505b60208210810361093557634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561098557600081815260208120601f850160051c810160208610156109625750805b601f850160051c820191505b818110156109815782815560010161096e565b5050505b505050565b815167ffffffffffffffff8111156109a4576109a4610652565b6109b8816109b28454610901565b8461093b565b602080601f8311600181146109ed57600084156109d55750858301515b600019600386901b1c1916600185901b178555610981565b600085815260208120601f198616915b82811015610a1c578886015182559484019460019091019084016109fd565b5085821015610a3a5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60008154610a5781610901565b808552602060018381168015610a745760018114610a8e57610abc565b60ff1985168884015283151560051b880183019550610abc565b866000528260002060005b85811015610ab45781548a8201860152908301908401610a99565b890184019650505b505050505092915050565b60208152606060208201526000610ae16080830184610a4a565b601f1980848303016040850152610afb8260018701610a4a565b915080848303016060850152506104858160028601610a4a56fea2646970667358221220c52240f46cd413822d5ab528685fc527bdc0ed964c5c89aeb3b9a9e773940c0c64736f6c63430008130033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// Ipfs is a free data retrieval call binding the contract method 0xa30e6db5.
//
// Solidity: function ipfs(string ) view returns(string fullName, string courseName, string createdAt)
func (_Api *ApiCaller) Ipfs(opts *bind.CallOpts, arg0 string) (struct {
	FullName   string
	CourseName string
	CreatedAt  string
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "ipfs", arg0)

	outstruct := new(struct {
		FullName   string
		CourseName string
		CreatedAt  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FullName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.CourseName = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.CreatedAt = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// Ipfs is a free data retrieval call binding the contract method 0xa30e6db5.
//
// Solidity: function ipfs(string ) view returns(string fullName, string courseName, string createdAt)
func (_Api *ApiSession) Ipfs(arg0 string) (struct {
	FullName   string
	CourseName string
	CreatedAt  string
}, error) {
	return _Api.Contract.Ipfs(&_Api.CallOpts, arg0)
}

// Ipfs is a free data retrieval call binding the contract method 0xa30e6db5.
//
// Solidity: function ipfs(string ) view returns(string fullName, string courseName, string createdAt)
func (_Api *ApiCallerSession) Ipfs(arg0 string) (struct {
	FullName   string
	CourseName string
	CreatedAt  string
}, error) {
	return _Api.Contract.Ipfs(&_Api.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiSession) Owner() (common.Address, error) {
	return _Api.Contract.Owner(&_Api.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Api *ApiCallerSession) Owner() (common.Address, error) {
	return _Api.Contract.Owner(&_Api.CallOpts)
}

// ViewData is a free data retrieval call binding the contract method 0x36d9d692.
//
// Solidity: function viewData(string _key) view returns((string,string,string))
func (_Api *ApiCaller) ViewData(opts *bind.CallOpts, _key string) (VeriflyData, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "viewData", _key)

	if err != nil {
		return *new(VeriflyData), err
	}

	out0 := *abi.ConvertType(out[0], new(VeriflyData)).(*VeriflyData)

	return out0, err

}

// ViewData is a free data retrieval call binding the contract method 0x36d9d692.
//
// Solidity: function viewData(string _key) view returns((string,string,string))
func (_Api *ApiSession) ViewData(_key string) (VeriflyData, error) {
	return _Api.Contract.ViewData(&_Api.CallOpts, _key)
}

// ViewData is a free data retrieval call binding the contract method 0x36d9d692.
//
// Solidity: function viewData(string _key) view returns((string,string,string))
func (_Api *ApiCallerSession) ViewData(_key string) (VeriflyData, error) {
	return _Api.Contract.ViewData(&_Api.CallOpts, _key)
}

// Add is a paid mutator transaction binding the contract method 0x878fc108.
//
// Solidity: function add(string _key, string _fullName, string _courseName, string _createdAt) returns(string)
func (_Api *ApiTransactor) Add(opts *bind.TransactOpts, _key string, _fullName string, _courseName string, _createdAt string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "add", _key, _fullName, _courseName, _createdAt)
}

// Add is a paid mutator transaction binding the contract method 0x878fc108.
//
// Solidity: function add(string _key, string _fullName, string _courseName, string _createdAt) returns(string)
func (_Api *ApiSession) Add(_key string, _fullName string, _courseName string, _createdAt string) (*types.Transaction, error) {
	return _Api.Contract.Add(&_Api.TransactOpts, _key, _fullName, _courseName, _createdAt)
}

// Add is a paid mutator transaction binding the contract method 0x878fc108.
//
// Solidity: function add(string _key, string _fullName, string _courseName, string _createdAt) returns(string)
func (_Api *ApiTransactorSession) Add(_key string, _fullName string, _courseName string, _createdAt string) (*types.Transaction, error) {
	return _Api.Contract.Add(&_Api.TransactOpts, _key, _fullName, _courseName, _createdAt)
}

// ApiAddedIterator is returned from FilterAdded and is used to iterate over the raw logs and unpacked data for Added events raised by the Api contract.
type ApiAddedIterator struct {
	Event *ApiAdded // Event containing the contract specifics and raw log

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
func (it *ApiAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiAdded)
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
		it.Event = new(ApiAdded)
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
func (it *ApiAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiAdded represents a Added event raised by the Api contract.
type ApiAdded struct {
	Arg0 VeriflyData
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAdded is a free log retrieval operation binding the contract event 0xec036d0cecfe4858346b9712e64b4ca7c2c49534a29b5d22db1b19049a772962.
//
// Solidity: event Added((string,string,string) arg0)
func (_Api *ApiFilterer) FilterAdded(opts *bind.FilterOpts) (*ApiAddedIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "Added")
	if err != nil {
		return nil, err
	}
	return &ApiAddedIterator{contract: _Api.contract, event: "Added", logs: logs, sub: sub}, nil
}

// WatchAdded is a free log subscription operation binding the contract event 0xec036d0cecfe4858346b9712e64b4ca7c2c49534a29b5d22db1b19049a772962.
//
// Solidity: event Added((string,string,string) arg0)
func (_Api *ApiFilterer) WatchAdded(opts *bind.WatchOpts, sink chan<- *ApiAdded) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "Added")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiAdded)
				if err := _Api.contract.UnpackLog(event, "Added", log); err != nil {
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

// ParseAdded is a log parse operation binding the contract event 0xec036d0cecfe4858346b9712e64b4ca7c2c49534a29b5d22db1b19049a772962.
//
// Solidity: event Added((string,string,string) arg0)
func (_Api *ApiFilterer) ParseAdded(log types.Log) (*ApiAdded, error) {
	event := new(ApiAdded)
	if err := _Api.contract.UnpackLog(event, "Added", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
