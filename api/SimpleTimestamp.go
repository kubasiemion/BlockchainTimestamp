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

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"serial\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"docHash\",\"type\":\"uint256\"}],\"name\":\"IssueCert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Sequence\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allCerts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CertID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"DocHash\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Holder\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"certsByIssue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CertID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"DocHash\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"Holder\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"url\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506002805460ff1916600117905561044d8061002d6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063070122791461005c57806319f707591461007a5780635600f04f1461009f578063652a9506146100b45780637bce5c61146100c7575b600080fd5b6100646100dc565b60405161007191906103a5565b60405180910390f35b61008d6100883660046102de565b6100e5565b60405161007196959493929190610371565b6100a7610127565b604051610071919061031e565b61008d6100c23660046102f6565b610143565b6100da6100d536600461029f565b610185565b005b60025460ff1690565b60006020819052908152604090208054600182015460028301546003840154600485015460059095015493949293919290916001600160a01b03918216911686565b6040518060600160405280603881526020016103e06038913981565b60016020819052600091825260409091208054918101546002820154600383015460048401546005909401549293919290916001600160a01b03908116911686565b6040805160c0810182524281524360208083019182528284018681526060840186815233608086019081526001600160a01b038a811660a0880190815260008b81528087528981209851808a5597516001808b0191825596516002808c0191825596516003808d0191825596516004808e0180549289166001600160a01b0319938416178155965160059e8f018054918a169184169190911781558a5460ff90811688529c8c90529e86209c8d559354998c019990995590548a880155549489019490945590549287018054861693831693909317909255965494909501805490921693909516929092179091558254169190610281836103b3565b91906101000a81548160ff021916908360ff16021790555050505050565b6000806000606084860312156102b3578283fd5b83356001600160a01b03811681146102c9578384fd5b95602085013595506040909401359392505050565b6000602082840312156102ef578081fd5b5035919050565b600060208284031215610307578081fd5b813560ff81168114610317578182fd5b9392505050565b6000602080835283518082850152825b8181101561034a5785810183015185820160400152820161032e565b8181111561035b5783604083870101525b50601f01601f1916929092016040019392505050565b9586526020860194909452604085019290925260608401526001600160a01b0390811660808401521660a082015260c00190565b60ff91909116815260200190565b600060ff821660ff8114156103d657634e487b7160e01b82526011600452602482fd5b6001019291505056fe68747470733a2f2f7265736561726368626f78312e756b736f7574682e636c6f75646170702e617a7572652e636f6d2f6365727474696d65a264697066735822122043fef3c662c9aab5ddec62a8c33c576e1e09152caffd3df20a751967b6edf38f64736f6c63430008000033",
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

// Sequence is a free data retrieval call binding the contract method 0x07012279.
//
// Solidity: function Sequence() view returns(uint8)
func (_Api *ApiCaller) Sequence(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "Sequence")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Sequence is a free data retrieval call binding the contract method 0x07012279.
//
// Solidity: function Sequence() view returns(uint8)
func (_Api *ApiSession) Sequence() (uint8, error) {
	return _Api.Contract.Sequence(&_Api.CallOpts)
}

// Sequence is a free data retrieval call binding the contract method 0x07012279.
//
// Solidity: function Sequence() view returns(uint8)
func (_Api *ApiCallerSession) Sequence() (uint8, error) {
	return _Api.Contract.Sequence(&_Api.CallOpts)
}

// AllCerts is a free data retrieval call binding the contract method 0x19f70759.
//
// Solidity: function allCerts(uint256 ) view returns(uint256 blockTime, uint256 blockNumber, uint256 CertID, uint256 DocHash, address issuer, address Holder)
func (_Api *ApiCaller) AllCerts(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BlockTime   *big.Int
	BlockNumber *big.Int
	CertID      *big.Int
	DocHash     *big.Int
	Issuer      common.Address
	Holder      common.Address
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "allCerts", arg0)

	outstruct := new(struct {
		BlockTime   *big.Int
		BlockNumber *big.Int
		CertID      *big.Int
		DocHash     *big.Int
		Issuer      common.Address
		Holder      common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CertID = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DocHash = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Issuer = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Holder = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// AllCerts is a free data retrieval call binding the contract method 0x19f70759.
//
// Solidity: function allCerts(uint256 ) view returns(uint256 blockTime, uint256 blockNumber, uint256 CertID, uint256 DocHash, address issuer, address Holder)
func (_Api *ApiSession) AllCerts(arg0 *big.Int) (struct {
	BlockTime   *big.Int
	BlockNumber *big.Int
	CertID      *big.Int
	DocHash     *big.Int
	Issuer      common.Address
	Holder      common.Address
}, error) {
	return _Api.Contract.AllCerts(&_Api.CallOpts, arg0)
}

// AllCerts is a free data retrieval call binding the contract method 0x19f70759.
//
// Solidity: function allCerts(uint256 ) view returns(uint256 blockTime, uint256 blockNumber, uint256 CertID, uint256 DocHash, address issuer, address Holder)
func (_Api *ApiCallerSession) AllCerts(arg0 *big.Int) (struct {
	BlockTime   *big.Int
	BlockNumber *big.Int
	CertID      *big.Int
	DocHash     *big.Int
	Issuer      common.Address
	Holder      common.Address
}, error) {
	return _Api.Contract.AllCerts(&_Api.CallOpts, arg0)
}

// CertsByIssue is a free data retrieval call binding the contract method 0x652a9506.
//
// Solidity: function certsByIssue(uint8 ) view returns(uint256 blockTime, uint256 blockNumber, uint256 CertID, uint256 DocHash, address issuer, address Holder)
func (_Api *ApiCaller) CertsByIssue(opts *bind.CallOpts, arg0 uint8) (struct {
	BlockTime   *big.Int
	BlockNumber *big.Int
	CertID      *big.Int
	DocHash     *big.Int
	Issuer      common.Address
	Holder      common.Address
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "certsByIssue", arg0)

	outstruct := new(struct {
		BlockTime   *big.Int
		BlockNumber *big.Int
		CertID      *big.Int
		DocHash     *big.Int
		Issuer      common.Address
		Holder      common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CertID = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DocHash = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Issuer = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Holder = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// CertsByIssue is a free data retrieval call binding the contract method 0x652a9506.
//
// Solidity: function certsByIssue(uint8 ) view returns(uint256 blockTime, uint256 blockNumber, uint256 CertID, uint256 DocHash, address issuer, address Holder)
func (_Api *ApiSession) CertsByIssue(arg0 uint8) (struct {
	BlockTime   *big.Int
	BlockNumber *big.Int
	CertID      *big.Int
	DocHash     *big.Int
	Issuer      common.Address
	Holder      common.Address
}, error) {
	return _Api.Contract.CertsByIssue(&_Api.CallOpts, arg0)
}

// CertsByIssue is a free data retrieval call binding the contract method 0x652a9506.
//
// Solidity: function certsByIssue(uint8 ) view returns(uint256 blockTime, uint256 blockNumber, uint256 CertID, uint256 DocHash, address issuer, address Holder)
func (_Api *ApiCallerSession) CertsByIssue(arg0 uint8) (struct {
	BlockTime   *big.Int
	BlockNumber *big.Int
	CertID      *big.Int
	DocHash     *big.Int
	Issuer      common.Address
	Holder      common.Address
}, error) {
	return _Api.Contract.CertsByIssue(&_Api.CallOpts, arg0)
}

// Url is a free data retrieval call binding the contract method 0x5600f04f.
//
// Solidity: function url() view returns(string)
func (_Api *ApiCaller) Url(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "url")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Url is a free data retrieval call binding the contract method 0x5600f04f.
//
// Solidity: function url() view returns(string)
func (_Api *ApiSession) Url() (string, error) {
	return _Api.Contract.Url(&_Api.CallOpts)
}

// Url is a free data retrieval call binding the contract method 0x5600f04f.
//
// Solidity: function url() view returns(string)
func (_Api *ApiCallerSession) Url() (string, error) {
	return _Api.Contract.Url(&_Api.CallOpts)
}

// IssueCert is a paid mutator transaction binding the contract method 0x7bce5c61.
//
// Solidity: function IssueCert(address holder, uint256 serial, uint256 docHash) returns()
func (_Api *ApiTransactor) IssueCert(opts *bind.TransactOpts, holder common.Address, serial *big.Int, docHash *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "IssueCert", holder, serial, docHash)
}

// IssueCert is a paid mutator transaction binding the contract method 0x7bce5c61.
//
// Solidity: function IssueCert(address holder, uint256 serial, uint256 docHash) returns()
func (_Api *ApiSession) IssueCert(holder common.Address, serial *big.Int, docHash *big.Int) (*types.Transaction, error) {
	return _Api.Contract.IssueCert(&_Api.TransactOpts, holder, serial, docHash)
}

// IssueCert is a paid mutator transaction binding the contract method 0x7bce5c61.
//
// Solidity: function IssueCert(address holder, uint256 serial, uint256 docHash) returns()
func (_Api *ApiTransactorSession) IssueCert(holder common.Address, serial *big.Int, docHash *big.Int) (*types.Transaction, error) {
	return _Api.Contract.IssueCert(&_Api.TransactOpts, holder, serial, docHash)
}
