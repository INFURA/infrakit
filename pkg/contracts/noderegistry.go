// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = abi.ConvertType
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fixedDeposit\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"Approved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"Disapproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Exit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"}],\"name\":\"approveProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvedList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_provider\",\"type\":\"address\"}],\"name\":\"disapproveProvider\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fixedDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registered\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"rpc\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_rpc\",\"type\":\"string\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_rpc\",\"type\":\"string\"}],\"name\":\"updateRPC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001b2238038062001b22833981810160405281019062000037919062000171565b620000576200004b6200006560201b60201c565b6200006d60201b60201c565b8060038190555050620001a3565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b6000819050919050565b6200014b8162000136565b81146200015757600080fd5b50565b6000815190506200016b8162000140565b92915050565b6000602082840312156200018a576200018962000131565b5b60006200019a848285016200015a565b91505092915050565b61196f80620001b36000396000f3fe60806040526004361061009c5760003560e01c80637a28399b116100645780637a28399b146101445780638da5cb5b1461016d578063b2dd5c0714610198578063f2fde38b146101d7578063f67ae71e14610200578063f76e947b1461021c5761009c565b806310ae97b0146100a15780632def6620146100ca5780633d68ad43146100d457806346f45b8d14610111578063715018a61461012d575b600080fd5b3480156100ad57600080fd5b506100c860048036038101906100c39190610f1a565b610247565b005b6100d2610329565b005b3480156100e057600080fd5b506100fb60048036038101906100f69190610fc1565b610552565b6040516101089190611009565b60405180910390f35b61012b60048036038101906101269190610f1a565b610572565b005b34801561013957600080fd5b506101426107f4565b005b34801561015057600080fd5b5061016b60048036038101906101669190610fc1565b610808565b005b34801561017957600080fd5b506101826108a2565b60405161018f9190611033565b60405180910390f35b3480156101a457600080fd5b506101bf60048036038101906101ba9190610fc1565b6108cb565b6040516101ce939291906110e6565b60405180910390f35b3480156101e357600080fd5b506101fe60048036038101906101f99190610fc1565b61098a565b005b61021a60048036038101906102159190611162565b610a0d565b005b34801561022857600080fd5b50610231610c13565b60405161023e919061118f565b60405180910390f35b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160009054906101000a900460ff166102d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102cd906111f6565b60405180910390fd5b80600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000190816103259190611422565b5050565b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154116103ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103a590611566565b60405180910390fd5b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101549050600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600080820160006104459190610d63565b60018201600090556002820160006101000a81549060ff02191690555050600033905060008173ffffffffffffffffffffffffffffffffffffffff168360405161048e906115b7565b60006040518083038185875af1925050503d80600081146104cb576040519150601f19603f3d011682016040523d82523d6000602084013e6104d0565b606091505b5050905080610514576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050b90611618565b60405180910390fd5b7f22d324652c93739755cf4581508b60875ebdd78c20c0cff5cf8e23452b2996313384604051610545929190611638565b60405180910390a1505050565b60026020528060005260406000206000915054906101000a900460ff1681565b60035434146105b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ad906116ad565b60405180910390fd5b600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610642576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106399061173f565b60405180910390fd5b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154146106c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106be906117ab565b60405180910390fd5b60018060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160006101000a81548160ff02191690831515021790555080600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000190816107709190611422565b5034600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055507f6f3bf3fa84e4763a43b3d23f9d79be242d6d5c834941ff4c1111b67469e1150c33346040516107e9929190611638565b60405180910390a150565b6107fc610c19565b6108066000610c97565b565b610810610c19565b6001600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f5d91bd0cecc45fef102af61de92c5462fadc884a5ce9d21c15e8a85198f2349e816040516108979190611033565b60405180910390a150565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60016020528060005260406000206000915090508060000180546108ee90611245565b80601f016020809104026020016040519081016040528092919081815260200182805461091a90611245565b80156109675780601f1061093c57610100808354040283529160200191610967565b820191906000526020600020905b81548152906001019060200180831161094a57829003601f168201915b5050505050908060010154908060020160009054906101000a900460ff16905083565b610992610c19565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610a01576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109f89061183d565b60405180910390fd5b610a0a81610c97565b50565b610a15610c19565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806060016040529081600082018054610ac990611245565b80601f0160208091040260200160405190810160405280929190818152602001828054610af590611245565b8015610b425780601f10610b1757610100808354040283529160200191610b42565b820191906000526020600020905b815481529060010190602001808311610b2557829003601f168201915b50505050508152602001600182015481526020016002820160009054906101000a900460ff1615151515815250509050806040015115610bd8576000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160006101000a81548160ff0219169083151502179055505b7f397e549cd5c892fc5a94e2bdf95875a84a2298a36b47123ad898f3ca3af496e882604051610c0791906118b2565b60405180910390a15050565b60035481565b610c21610d5b565b73ffffffffffffffffffffffffffffffffffffffff16610c3f6108a2565b73ffffffffffffffffffffffffffffffffffffffff1614610c95576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c8c90611919565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b508054610d6f90611245565b6000825580601f10610d815750610da0565b601f016020900490600052602060002090810190610d9f9190610da3565b5b50565b5b80821115610dbc576000816000905550600101610da4565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610e2782610dde565b810181811067ffffffffffffffff82111715610e4657610e45610def565b5b80604052505050565b6000610e59610dc0565b9050610e658282610e1e565b919050565b600067ffffffffffffffff821115610e8557610e84610def565b5b610e8e82610dde565b9050602081019050919050565b82818337600083830152505050565b6000610ebd610eb884610e6a565b610e4f565b905082815260208101848484011115610ed957610ed8610dd9565b5b610ee4848285610e9b565b509392505050565b600082601f830112610f0157610f00610dd4565b5b8135610f11848260208601610eaa565b91505092915050565b600060208284031215610f3057610f2f610dca565b5b600082013567ffffffffffffffff811115610f4e57610f4d610dcf565b5b610f5a84828501610eec565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610f8e82610f63565b9050919050565b610f9e81610f83565b8114610fa957600080fd5b50565b600081359050610fbb81610f95565b92915050565b600060208284031215610fd757610fd6610dca565b5b6000610fe584828501610fac565b91505092915050565b60008115159050919050565b61100381610fee565b82525050565b600060208201905061101e6000830184610ffa565b92915050565b61102d81610f83565b82525050565b60006020820190506110486000830184611024565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561108857808201518184015260208101905061106d565b60008484015250505050565b600061109f8261104e565b6110a98185611059565b93506110b981856020860161106a565b6110c281610dde565b840191505092915050565b6000819050919050565b6110e0816110cd565b82525050565b600060608201905081810360008301526111008186611094565b905061110f60208301856110d7565b61111c6040830184610ffa565b949350505050565b600061112f82610f63565b9050919050565b61113f81611124565b811461114a57600080fd5b50565b60008135905061115c81611136565b92915050565b60006020828403121561117857611177610dca565b5b60006111868482850161114d565b91505092915050565b60006020820190506111a460008301846110d7565b92915050565b7f4f6e6c79206163746976652070726f7669646572732e00000000000000000000600082015250565b60006111e0601683611059565b91506111eb826111aa565b602082019050919050565b6000602082019050818103600083015261120f816111d3565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061125d57607f821691505b6020821081036112705761126f611216565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026112d87fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261129b565b6112e2868361129b565b95508019841693508086168417925050509392505050565b6000819050919050565b600061131f61131a611315846110cd565b6112fa565b6110cd565b9050919050565b6000819050919050565b61133983611304565b61134d61134582611326565b8484546112a8565b825550505050565b600090565b611362611355565b61136d818484611330565b505050565b5b818110156113915761138660008261135a565b600181019050611373565b5050565b601f8211156113d6576113a781611276565b6113b08461128b565b810160208510156113bf578190505b6113d36113cb8561128b565b830182611372565b50505b505050565b600082821c905092915050565b60006113f9600019846008026113db565b1980831691505092915050565b600061141283836113e8565b9150826002028217905092915050565b61142b8261104e565b67ffffffffffffffff81111561144457611443610def565b5b61144e8254611245565b611459828285611395565b600060209050601f83116001811461148c576000841561147a578287015190505b6114848582611406565b8655506114ec565b601f19841661149a86611276565b60005b828110156114c25784890151825560018201915060208501945060208101905061149d565b868310156114df57848901516114db601f8916826113e8565b8355505b6001600288020188555050505b505050505050565b7f43616c6c6572206d757374206861766520616e206163746976652062616c616e60008201527f63652e0000000000000000000000000000000000000000000000000000000000602082015250565b6000611550602383611059565b915061155b826114f4565b604082019050919050565b6000602082019050818103600083015261157f81611543565b9050919050565b600081905092915050565b50565b60006115a1600083611586565b91506115ac82611591565b600082019050919050565b60006115c282611594565b9150819050919050565b7f4661696c656420746f2073656e64204574686572000000000000000000000000600082015250565b6000611602601483611059565b915061160d826115cc565b602082019050919050565b60006020820190508181036000830152611631816115f5565b9050919050565b600060408201905061164d6000830185611024565b61165a60208301846110d7565b9392505050565b7f4669786564206465706f7369742069732072657175697265642e000000000000600082015250565b6000611697601a83611059565b91506116a282611661565b602082019050919050565b600060208201905081810360008301526116c68161168a565b9050919050565b7f50726f7669646572206d757374206265207072652d617070726f76656420627960008201527f206f776e65722e00000000000000000000000000000000000000000000000000602082015250565b6000611729602783611059565b9150611734826116cd565b604082019050919050565b600060208201905081810360",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, _fixedDeposit *big.Int) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend, _fixedDeposit)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// ApprovedList is a free data retrieval call binding the contract method 0x3d68ad43.
//
// Solidity: function approvedList(address ) view returns(bool)
func (_Contract *ContractCaller) ApprovedList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "approvedList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedList is a free data retrieval call binding the contract method 0x3d68ad43.
//
// Solidity: function approvedList(address ) view returns(bool)
func (_Contract *ContractSession) ApprovedList(arg0 common.Address) (bool, error) {
	return _Contract.Contract.ApprovedList(&_Contract.CallOpts, arg0)
}

// ApprovedList is a free data retrieval call binding the contract method 0x3d68ad43.
//
// Solidity: function approvedList(address ) view returns(bool)
func (_Contract *ContractCallerSession) ApprovedList(arg0 common.Address) (bool, error) {
	return _Contract.Contract.ApprovedList(&_Contract.CallOpts, arg0)
}

// FixedDeposit is a free data retrieval call binding the contract method 0xf76e947b.
//
// Solidity: function fixedDeposit() view returns(uint256)
func (_Contract *ContractCaller) FixedDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "fixedDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FixedDeposit is a free data retrieval call binding the contract method 0xf76e947b.
//
// Solidity: function fixedDeposit() view returns(uint256)
func (_Contract *ContractSession) FixedDeposit() (*big.Int, error) {
	return _Contract.Contract.FixedDeposit(&_Contract.CallOpts)
}

// FixedDeposit is a free data retrieval call binding the contract method 0xf76e947b.
//
// Solidity: function fixedDeposit() view returns(uint256)
func (_Contract *ContractCallerSession) FixedDeposit() (*big.Int, error) {
	return _Contract.Contract.FixedDeposit(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Registered is a free data retrieval call binding the contract method 0xb2dd5c07.
//
// Solidity: function registered(address ) view returns(string rpc, uint256 balance, bool active)
func (_Contract *ContractCaller) Registered(opts *bind.CallOpts, arg0 common.Address) (struct {
	Rpc     string
	Balance *big.Int
	Active  bool
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "registered", arg0)

	outstruct := new(struct {
		Rpc     string
		Balance *big.Int
		Active  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Rpc = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Balance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Active = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Registered is a free data retrieval call binding the contract method 0xb2dd5c07.
//
// Solidity: function registered(address ) view returns(string rpc, uint256 balance, bool active)
func (_Contract *ContractSession) Registered(arg0 common.Address) (struct {
	Rpc     string
	Balance *big.Int
	Active  bool
}, error) {
	return _Contract.Contract.Registered(&_Contract.CallOpts, arg0)
}

// Registered is a free data retrieval call binding the contract method 0xb2dd5c07.
//
// Solidity: function registered(address ) view returns(string rpc, uint256 balance, bool active)
func (_Contract *ContractCallerSession) Registered(arg0 common.Address) (struct {
	Rpc     string
	Balance *big.Int
	Active  bool
}, error) {
	return _Contract.Contract.Registered(&_Contract.CallOpts, arg0)
}

// ApproveProvider is a paid mutator transaction binding the contract method 0x7a28399b.
//
// Solidity: function approveProvider(address _provider) returns()
func (_Contract *ContractTransactor) ApproveProvider(opts *bind.TransactOpts, _provider common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "approveProvider", _provider)
}

// ApproveProvider is a paid mutator transaction binding the contract method 0x7a28399b.
//
// Solidity: function approveProvider(address _provider) returns()
func (_Contract *ContractSession) ApproveProvider(_provider common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ApproveProvider(&_Contract.TransactOpts, _provider)
}

// ApproveProvider is a paid mutator transaction binding the contract method 0x7a28399b.
//
// Solidity: function approveProvider(address _provider) returns()
func (_Contract *ContractTransactorSession) ApproveProvider(_provider common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ApproveProvider(&_Contract.TransactOpts, _provider)
}

// DisapproveProvider is a paid mutator transaction binding the contract method 0xf67ae71e.
//
// Solidity: function disapproveProvider(address _provider) payable returns()
func (_Contract *ContractTransactor) DisapproveProvider(opts *bind.TransactOpts, _provider common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "disapproveProvider", _provider)
}

// DisapproveProvider is a paid mutator transaction binding the contract method 0xf67ae71e.
//
// Solidity: function disapproveProvider(address _provider) payable returns()
func (_Contract *ContractSession) DisapproveProvider(_provider common.Address) (*types.Transaction, error) {
	return _Contract.Contract.DisapproveProvider(&_Contract.TransactOpts, _provider)
}

// DisapproveProvider is a paid mutator transaction binding the contract method 0xf67ae71e.
//
// Solidity: function disapproveProvider(address _provider) payable returns()
func (_Contract *ContractTransactorSession) DisapproveProvider(_provider common.Address) (*types.Transaction, error) {
	return _Contract.Contract.DisapproveProvider(&_Contract.TransactOpts, _provider)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x46f45b8d.
//
// Solidity: function stake(string _rpc) payable returns()
func (_Contract *ContractTransactor) Stake(opts *bind.TransactOpts, _rpc string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "stake", _rpc)
}

// Stake is a paid mutator transaction binding the contract method 0x46f45b8d.
//
// Solidity: function stake(string _rpc) payable returns()
func (_Contract *ContractSession) Stake(_rpc string) (*types.Transaction, error) {
	return _Contract.Contract.Stake(&_Contract.TransactOpts, _rpc)
}

// Stake is a paid mutator transaction binding the contract method 0x46f45b8d.
//
// Solidity: function stake(string _rpc) payable returns()
func (_Contract *ContractTransactorSession) Stake(_rpc string) (*types.Transaction, error) {
	return _Contract.Contract.Stake(&_Contract.TransactOpts, _rpc)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() payable returns()
func (_Contract *ContractTransactor) Unstake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unstake")
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() payable returns()
func (_Contract *ContractSession) Unstake() (*types.Transaction, error) {
	return _Contract.Contract.Unstake(&_Contract.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() payable returns()
func (_Contract *ContractTransactorSession) Unstake() (*types.Transaction, error) {
	return _Contract.Contract.Unstake(&_Contract.TransactOpts)
}

// UpdateRPC is a paid mutator transaction binding the contract method 0x10ae97b0.
//
// Solidity: function updateRPC(string _rpc) returns()
func (_Contract *ContractTransactor) UpdateRPC(opts *bind.TransactOpts, _rpc string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateRPC", _rpc)
}

// UpdateRPC is a paid mutator transaction binding the contract method 0x10ae97b0.
//
// Solidity: function updateRPC(string _rpc) returns()
func (_Contract *ContractSession) UpdateRPC(_rpc string) (*types.Transaction, error) {
	return _Contract.Contract.UpdateRPC(&_Contract.TransactOpts, _rpc)
}

// UpdateRPC is a paid mutator transaction binding the contract method 0x10ae97b0.
//
// Solidity: function updateRPC(string _rpc) returns()
func (_Contract *ContractTransactorSession) UpdateRPC(_rpc string) (*types.Transaction, error) {
	return _Contract.Contract.UpdateRPC(&_Contract.TransactOpts, _rpc)
}

// ContractApprovedIterator is returned from FilterApproved and is used to iterate over the raw logs and unpacked data for Approved events raised by the Contract contract.
type ContractApprovedIterator struct {
	Event *ContractApproved // Event containing the contract specifics and raw log

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
func (it *ContractApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApproved)
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
		it.Event = new(ContractApproved)
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
func (it *ContractApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractApproved represents a Approved event raised by the Contract contract.
type ContractApproved struct {
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproved is a free log retrieval operation binding the contract event 0x5d91bd0cecc45fef102af61de92c5462fadc884a5ce9d21c15e8a85198f2349e.
//
// Solidity: event Approved(address provider)
func (_Contract *ContractFilterer) FilterApproved(opts *bind.FilterOpts) (*ContractApprovedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Approved")
	if err != nil {
		return nil, err
	}
	return &ContractApprovedIterator{contract: _Contract.contract, event: "Approved", logs: logs, sub: sub}, nil
}

// WatchApproved is a free log subscription operation binding the contract event 0x5d91bd0cecc45fef102af61de92c5462fadc884a5ce9d21c15e8a85198f2349e.
//
// Solidity: event Approved(address provider)
func (_Contract *ContractFilterer) WatchApproved(opts *bind.WatchOpts, sink chan<- *ContractApproved) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Approved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractApproved)
				if err := _Contract.contract.UnpackLog(event, "Approved", log); err != nil {
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

// ParseApproved is a log parse operation binding the contract event 0x5d91bd0cecc45fef102af61de92c5462fadc884a5ce9d21c15e8a85198f2349e.
//
// Solidity: event Approved(address provider)
func (_Contract *ContractFilterer) ParseApproved(log types.Log) (*ContractApproved, error) {
	event := new(ContractApproved)
	if err := _Contract.contract.UnpackLog(event, "Approved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDisapprovedIterator is returned from FilterDisapproved and is used to iterate over the raw logs and unpacked data for Disapproved events raised by the Contract contract.
type ContractDisapprovedIterator struct {
	Event *ContractDisapproved // Event containing the contract specifics and raw log

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
func (it *ContractDisapprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDisapproved)
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
		it.Event = new(ContractDisapproved)
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
func (it *ContractDisapprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDisapprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDisapproved represents a Disapproved event raised by the Contract contract.
type ContractDisapproved struct {
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDisapproved is a free log retrieval operation binding the contract event 0x397e549cd5c892fc5a94e2bdf95875a84a2298a36b47123ad898f3ca3af496e8.
//
// Solidity: event Disapproved(address provider)
func (_Contract *ContractFilterer) FilterDisapproved(opts *bind.FilterOpts) (*ContractDisapprovedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Disapproved")
	if err != nil {
		return nil, err
	}
	return &ContractDisapprovedIterator{contract: _Contract.contract, event: "Disapproved", logs: logs, sub: sub}, nil
}

// WatchDisapproved is a free log subscription operation binding the contract event 0x397e549cd5c892fc5a94e2bdf95875a84a2298a36b47123ad898f3ca3af496e8.
//
// Solidity: event Disapproved(address provider)
func (_Contract *ContractFilterer) WatchDisapproved(opts *bind.WatchOpts, sink chan<- *ContractDisapproved) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Disapproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDisapproved)
				if err := _Contract.contract.UnpackLog(event, "Disapproved", log); err != nil {
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

// ParseDisapproved is a log parse operation binding the contract event 0x397e549cd5c892fc5a94e2bdf95875a84a2298a36b47123ad898f3ca3af496e8.
//
// Solidity: event Disapproved(address provider)
func (_Contract *ContractFilterer) ParseDisapproved(log types.Log) (*ContractDisapproved, error) {
	event := new(ContractDisapproved)
	if err := _Contract.contract.UnpackLog(event, "Disapproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractExitIterator is returned from FilterExit and is used to iterate over the raw logs and unpacked data for Exit events raised by the Contract contract.
type ContractExitIterator struct {
	Event *ContractExit // Event containing the contract specifics and raw log

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
func (it *ContractExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractExit)
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
		it.Event = new(ContractExit)
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
func (it *ContractExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractExit represents a Exit event raised by the Contract contract.
type ContractExit struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExit is a free log retrieval operation binding the contract event 0x22d324652c93739755cf4581508b60875ebdd78c20c0cff5cf8e23452b299631.
//
// Solidity: event Exit(address provider, uint256 amount)
func (_Contract *ContractFilterer) FilterExit(opts *bind.FilterOpts) (*ContractExitIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Exit")
	if err != nil {
		return nil, err
	}
	return &ContractExitIterator{contract: _Contract.contract, event: "Exit", logs: logs, sub: sub}, nil
}

// WatchExit is a free log subscription operation binding the contract event 0x22d324652c93739755cf4581508b60875ebdd78c20c0cff5cf8e23452b299631.
//
// Solidity: event Exit(address provider, uint256 amount)
func (_Contract *ContractFilterer) WatchExit(opts *bind.WatchOpts, sink chan<- *ContractExit) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Exit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractExit)
				if err := _Contract.contract.UnpackLog(event, "Exit", log); err != nil {
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

// ParseExit is a log parse operation binding the contract event 0x22d324652c93739755cf4581508b60875ebdd78c20c0cff5cf8e23452b299631.
//
// Solidity: event Exit(address provider, uint256 amount)
func (_Contract *ContractFilterer) ParseExit(log types.Log) (*ContractExit, error) {
	event := new(ContractExit)
	if err := _Contract.contract.UnpackLog(event, "Exit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the Contract contract.
type ContractRegisteredIterator struct {
	Event *ContractRegistered // Event containing the contract specifics and raw log

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
func (it *ContractRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistered)
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
		it.Event = new(ContractRegistered)
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
func (it *ContractRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistered represents a Registered event raised by the Contract contract.
type ContractRegistered struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x6f3bf3fa84e4763a43b3d23f9d79be242d6d5c834941ff4c1111b67469e1150c.
//
// Solidity: event Registered(address provider, uint256 amount)
func (_Contract *ContractFilterer) FilterRegistered(opts *bind.FilterOpts) (*ContractRegisteredIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return &ContractRegisteredIterator{contract: _Contract.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x6f3bf3fa84e4763a43b3d23f9d79be242d6d5c834941ff4c1111b67469e1150c.
//
// Solidity: event Registered(address provider, uint256 amount)
func (_Contract *ContractFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *ContractRegistered) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistered)
				if err := _Contract.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0x6f3bf3fa84e4763a43b3d23f9d79be242d6d5c834941ff4c1111b67469e1150c.
//
// Solidity: event Registered(address provider, uint256 amount)
func (_Contract *ContractFilterer) ParseRegistered(log types.Log) (*ContractRegistered, error) {
	event := new(ContractRegistered)
	if err := _Contract.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
