// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package compoundLogic

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CompoundLogicABI is the input ABI used to generate the binding from.
const CompoundLogicABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"addToMarkets\",\"type\":\"address[]\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"addToMarkets\",\"type\":\"address[]\"}],\"name\":\"borrowByMultiCollateral\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cEther\",\"outputs\":[{\"internalType\":\"contractCEther\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"contractComptroller\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"incognitoSmartContract\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateBorrow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isredeemUnderlying\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"exitToMarkets\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"repayBorrow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// CompoundLogicBin is the compiled bytecode used for deploying new contracts.
var CompoundLogicBin = "0x608060405234801561001057600080fd5b50611fbf806100206000396000f3fe6080604052600436106100915760003560e01c806372e94bf61161005957806372e94bf6146103a25780637520f7ed146103f9578063abdb5ea8146104c7578063b42a644b1461055c578063b6dbc8ed146105b357610091565b806319b68c001461009357806326d5e251146100ea57806340c10f19146101e15780635fe3b5671461027657806364fd7078146102cd575b005b34801561009f57600080fd5b506100a8610700565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156100f657600080fd5b506101986004803603606081101561010d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561015457600080fd5b82018360208201111561016657600080fd5b8035906020019184602083028401116401000000008311171561018857600080fd5b9091929391929390505050610718565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b61022d600480360360408110156101f757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610b45565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b34801561028257600080fd5b5061028b610d41565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610359600480360360808110156102e357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d59565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b3480156103ae57600080fd5b506103b761102a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561040557600080fd5b5061047e6004803603608081101561041c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803515159060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061102f565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b610513600480360360408110156104dd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061148d565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390f35b34801561056857600080fd5b50610571611675565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156105bf57600080fd5b50610661600480360360608110156105d657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561061d57600080fd5b82018360208201111561062f57600080fd5b8035906020019184602083028401116401000000008311171561065157600080fd5b909192939192939050505061168d565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156106a857808201518184015260208101905061068d565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156106ea5780820151818401526020810190506106cf565b5050505090500194505050505060405180910390f35b73f92fbe0d3c0dcdae407923b2ac17ec223b1084e481565b600080600084849050111561089457731f5d7f3caac149fe41b8bd62a3673fe6ec0ab73b73ffffffffffffffffffffffffffffffffffffffff1663c299823885856040518363ffffffff1660e01b815260040180806020018281038252848482818152602001925060200280828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b1580156107bd57600080fd5b505af11580156107d1573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060208110156107fb57600080fd5b810190808051604051939291908464010000000082111561081b57600080fd5b8382019150602082018581111561083157600080fd5b825186602082028301116401000000008211171561084e57600080fd5b8083526020830192505050908051906020019060200280838360005b8381101561088557808201518184015260208101905061086a565b50505050905001604052505050505b60008673ffffffffffffffffffffffffffffffffffffffff1663c5ebeaec876040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b1580156108e957600080fd5b505af11580156108fd573d6000803e3d6000fd5b505050506040513d602081101561091357600080fd5b81019080805190602001909291905050501461092e57600080fd5b600080905073f92fbe0d3c0dcdae407923b2ac17ec223b1084e473ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16141561099f576109856000611b00565b9050610992600082611bfe565b6000819250925050610b3c565b610a288773ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b1580156109e857600080fd5b505afa1580156109fc573d6000803e3d6000fd5b505050506040513d6020811015610a1257600080fd5b8101908080519060200190929190505050611b00565b9050610ab48773ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b158015610a7357600080fd5b505afa158015610a87573d6000803e3d6000fd5b505050506040513d6020811015610a9d57600080fd5b810190808051906020019092919050505082611bfe565b8673ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b158015610afa57600080fd5b505afa158015610b0e573d6000803e3d6000fd5b505050506040513d6020811015610b2457600080fd5b81019080805190602001909291905050508192509250505b94509492505050565b60008073f92fbe0d3c0dcdae407923b2ac17ec223b1084e473ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610bf6578373ffffffffffffffffffffffffffffffffffffffff16631249c58b346040518263ffffffff1660e01b81526004016000604051808303818588803b158015610bd857600080fd5b505af1158015610bec573d6000803e3d6000fd5b5050505050610d1c565b610c818473ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b158015610c3f57600080fd5b505afa158015610c53573d6000803e3d6000fd5b505050506040513d6020811015610c6957600080fd5b81019080805190602001909291905050508585611d69565b60008473ffffffffffffffffffffffffffffffffffffffff1663a0712d68856040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b158015610cd657600080fd5b505af1158015610cea573d6000803e3d6000fd5b505050506040513d6020811015610d0057600080fd5b810190808051906020019092919050505014610d1b57600080fd5b5b6000610d2785611b00565b9050610d338582611bfe565b848192509250509250929050565b731f5d7f3caac149fe41b8bd62a3673fe6ec0ab73b81565b60008073f92fbe0d3c0dcdae407923b2ac17ec223b1084e473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161415610e75578573ffffffffffffffffffffffffffffffffffffffff1663aae40a2a3487866040518463ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001925050506000604051808303818588803b158015610e5757600080fd5b505af1158015610e6b573d6000803e3d6000fd5b5050505050611003565b610f008673ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b158015610ebe57600080fd5b505afa158015610ed2573d6000803e3d6000fd5b505050506040513d6020811015610ee857600080fd5b81019080805190602001909291905050508786611d69565b60008673ffffffffffffffffffffffffffffffffffffffff1663f5e3c4628787876040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019350505050602060405180830381600087803b158015610fbd57600080fd5b505af1158015610fd1573d6000803e3d6000fd5b505050506040513d6020811015610fe757600080fd5b81019080805190602001909291905050501461100257600080fd5b5b600061100e84611b00565b905061101a8482611bfe565b8381925092505094509492505050565b600081565b600080600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161461113657731f5d7f3caac149fe41b8bd62a3673fe6ec0ab73b73ffffffffffffffffffffffffffffffffffffffff1663ede4edd0846040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1580156110f957600080fd5b505af115801561110d573d6000803e3d6000fd5b505050506040513d602081101561112357600080fd5b8101908080519060200190929190505050505b83156111db5760008673ffffffffffffffffffffffffffffffffffffffff1663852a12e3876040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b15801561119157600080fd5b505af11580156111a5573d6000803e3d6000fd5b505050506040513d60208110156111bb57600080fd5b8101908080519060200190929190505050146111d657600080fd5b611276565b60008673ffffffffffffffffffffffffffffffffffffffff1663db006a75876040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b15801561123057600080fd5b505af1158015611244573d6000803e3d6000fd5b505050506040513d602081101561125a57600080fd5b81019080805190602001909291905050501461127557600080fd5b5b600080905073f92fbe0d3c0dcdae407923b2ac17ec223b1084e473ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1614156112e7576112cd6000611b00565b90506112da600082611bfe565b6000819250925050611484565b6113708773ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b15801561133057600080fd5b505afa158015611344573d6000803e3d6000fd5b505050506040513d602081101561135a57600080fd5b8101908080519060200190929190505050611b00565b90506113fc8773ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b1580156113bb57600080fd5b505afa1580156113cf573d6000803e3d6000fd5b505050506040513d60208110156113e557600080fd5b810190808051906020019092919050505082611bfe565b8673ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b15801561144257600080fd5b505afa158015611456573d6000803e3d6000fd5b505050506040513d602081101561146c57600080fd5b81019080805190602001909291905050508192509250505b94509492505050565b60008073f92fbe0d3c0dcdae407923b2ac17ec223b1084e473ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141561153e578373ffffffffffffffffffffffffffffffffffffffff16634e4d9fea346040518263ffffffff1660e01b81526004016000604051808303818588803b15801561152057600080fd5b505af1158015611534573d6000803e3d6000fd5b5050505050611664565b6115c98473ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b15801561158757600080fd5b505afa15801561159b573d6000803e3d6000fd5b505050506040513d60208110156115b157600080fd5b81019080805190602001909291905050508585611d69565b60008473ffffffffffffffffffffffffffffffffffffffff16630e752702856040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b15801561161e57600080fd5b505af1158015611632573d6000803e3d6000fd5b505050506040513d602081101561164857600080fd5b81019080805190602001909291905050501461166357600080fd5b5b836000809050915091509250929050565b738c4b922e2f7d6d1aba41d79c47f497f6f54e0af881565b606080600084849050111561180957731f5d7f3caac149fe41b8bd62a3673fe6ec0ab73b73ffffffffffffffffffffffffffffffffffffffff1663c299823885856040518363ffffffff1660e01b815260040180806020018281038252848482818152602001925060200280828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b15801561173257600080fd5b505af1158015611746573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561177057600080fd5b810190808051604051939291908464010000000082111561179057600080fd5b838201915060208201858111156117a657600080fd5b82518660208202830111640100000000821117156117c357600080fd5b8083526020830192505050908051906020019060200280838360005b838110156117fa5780820151818401526020810190506117df565b50505050905001604052505050505b60008673ffffffffffffffffffffffffffffffffffffffff1663c5ebeaec876040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b15801561185e57600080fd5b505af1158015611872573d6000803e3d6000fd5b505050506040513d602081101561188857600080fd5b8101908080519060200190929190505050146118a357600080fd5b600080905073f92fbe0d3c0dcdae407923b2ac17ec223b1084e473ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16141561190c576118fa6000611b00565b9050611907600082611bfe565b611a22565b6119958773ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b15801561195557600080fd5b505afa158015611969573d6000803e3d6000fd5b505050506040513d602081101561197f57600080fd5b8101908080519060200190929190505050611b00565b9050611a218773ffffffffffffffffffffffffffffffffffffffff16636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b1580156119e057600080fd5b505afa1580156119f4573d6000803e3d6000fd5b505050506040513d6020811015611a0a57600080fd5b810190808051906020019092919050505082611bfe565b5b60606001604051908082528060200260200182016040528015611a545781602001602082028038833980820191505090505b5090508781600081518110611a6557fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060606001604051908082528060200260200182016040528015611ad15781602001602082028038833980820191505090505b5090508281600081518110611ae257fe5b60200260200101818152505081819450945050505094509492505050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611b3e57479050611bf9565b8173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015611bbb57600080fd5b505afa158015611bcf573d6000803e3d6000fd5b505050506040513d6020811015611be557600080fd5b810190808051906020019092919050505090505b919050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611ca05780471015611c4057600080fd5b738c4b922e2f7d6d1aba41d79c47f497f6f54e0af873ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015611c9a573d6000803e3d6000fd5b50611d65565b8173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb738c4b922e2f7d6d1aba41d79c47f497f6f54e0af8836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b158015611d3b57600080fd5b505af1158015611d4f573d6000803e3d6000fd5b50505050611d5b611f4c565b611d6457600080fd5b5b5050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614611f47578273ffffffffffffffffffffffffffffffffffffffff1663095ea7b38360006040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015611e2557600080fd5b505af1158015611e39573d6000803e3d6000fd5b505050506040513d6020811015611e4f57600080fd5b810190808051906020019092919050505050611e69611f4c565b611e7257600080fd5b8273ffffffffffffffffffffffffffffffffffffffff1663095ea7b383836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015611ef957600080fd5b505af1158015611f0d573d6000803e3d6000fd5b505050506040513d6020811015611f2357600080fd5b810190808051906020019092919050505050611f3d611f4c565b611f4657600080fd5b5b505050565b600080600090503d60008114611f695760208114611f7257611f7e565b60019150611f7e565b60206000803e60005191505b5060008114159150509056fea265627a7a7231582011b7d01464e23962a33a39a62e70158611f6f829337a261d31ef34d7a2afecf164736f6c63430005100032"

// DeployCompoundLogic deploys a new Ethereum contract, binding an instance of CompoundLogic to it.
func DeployCompoundLogic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CompoundLogic, error) {
	parsed, err := abi.JSON(strings.NewReader(CompoundLogicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CompoundLogicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CompoundLogic{CompoundLogicCaller: CompoundLogicCaller{contract: contract}, CompoundLogicTransactor: CompoundLogicTransactor{contract: contract}, CompoundLogicFilterer: CompoundLogicFilterer{contract: contract}}, nil
}

// CompoundLogic is an auto generated Go binding around an Ethereum contract.
type CompoundLogic struct {
	CompoundLogicCaller     // Read-only binding to the contract
	CompoundLogicTransactor // Write-only binding to the contract
	CompoundLogicFilterer   // Log filterer for contract events
}

// CompoundLogicCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompoundLogicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundLogicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompoundLogicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundLogicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompoundLogicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundLogicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompoundLogicSession struct {
	Contract     *CompoundLogic    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CompoundLogicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompoundLogicCallerSession struct {
	Contract *CompoundLogicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CompoundLogicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompoundLogicTransactorSession struct {
	Contract     *CompoundLogicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CompoundLogicRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompoundLogicRaw struct {
	Contract *CompoundLogic // Generic contract binding to access the raw methods on
}

// CompoundLogicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompoundLogicCallerRaw struct {
	Contract *CompoundLogicCaller // Generic read-only contract binding to access the raw methods on
}

// CompoundLogicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompoundLogicTransactorRaw struct {
	Contract *CompoundLogicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompoundLogic creates a new instance of CompoundLogic, bound to a specific deployed contract.
func NewCompoundLogic(address common.Address, backend bind.ContractBackend) (*CompoundLogic, error) {
	contract, err := bindCompoundLogic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompoundLogic{CompoundLogicCaller: CompoundLogicCaller{contract: contract}, CompoundLogicTransactor: CompoundLogicTransactor{contract: contract}, CompoundLogicFilterer: CompoundLogicFilterer{contract: contract}}, nil
}

// NewCompoundLogicCaller creates a new read-only instance of CompoundLogic, bound to a specific deployed contract.
func NewCompoundLogicCaller(address common.Address, caller bind.ContractCaller) (*CompoundLogicCaller, error) {
	contract, err := bindCompoundLogic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundLogicCaller{contract: contract}, nil
}

// NewCompoundLogicTransactor creates a new write-only instance of CompoundLogic, bound to a specific deployed contract.
func NewCompoundLogicTransactor(address common.Address, transactor bind.ContractTransactor) (*CompoundLogicTransactor, error) {
	contract, err := bindCompoundLogic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundLogicTransactor{contract: contract}, nil
}

// NewCompoundLogicFilterer creates a new log filterer instance of CompoundLogic, bound to a specific deployed contract.
func NewCompoundLogicFilterer(address common.Address, filterer bind.ContractFilterer) (*CompoundLogicFilterer, error) {
	contract, err := bindCompoundLogic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompoundLogicFilterer{contract: contract}, nil
}

// bindCompoundLogic binds a generic wrapper to an already deployed contract.
func bindCompoundLogic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CompoundLogicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompoundLogic *CompoundLogicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CompoundLogic.Contract.CompoundLogicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompoundLogic *CompoundLogicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompoundLogic.Contract.CompoundLogicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompoundLogic *CompoundLogicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompoundLogic.Contract.CompoundLogicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompoundLogic *CompoundLogicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CompoundLogic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompoundLogic *CompoundLogicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompoundLogic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompoundLogic *CompoundLogicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompoundLogic.Contract.contract.Transact(opts, method, params...)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_CompoundLogic *CompoundLogicCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CompoundLogic.contract.Call(opts, out, "ETH_CONTRACT_ADDRESS")
	return *ret0, err
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_CompoundLogic *CompoundLogicSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _CompoundLogic.Contract.ETHCONTRACTADDRESS(&_CompoundLogic.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() constant returns(address)
func (_CompoundLogic *CompoundLogicCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _CompoundLogic.Contract.ETHCONTRACTADDRESS(&_CompoundLogic.CallOpts)
}

// CEther is a free data retrieval call binding the contract method 0x19b68c00.
//
// Solidity: function cEther() constant returns(address)
func (_CompoundLogic *CompoundLogicCaller) CEther(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CompoundLogic.contract.Call(opts, out, "cEther")
	return *ret0, err
}

// CEther is a free data retrieval call binding the contract method 0x19b68c00.
//
// Solidity: function cEther() constant returns(address)
func (_CompoundLogic *CompoundLogicSession) CEther() (common.Address, error) {
	return _CompoundLogic.Contract.CEther(&_CompoundLogic.CallOpts)
}

// CEther is a free data retrieval call binding the contract method 0x19b68c00.
//
// Solidity: function cEther() constant returns(address)
func (_CompoundLogic *CompoundLogicCallerSession) CEther() (common.Address, error) {
	return _CompoundLogic.Contract.CEther(&_CompoundLogic.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() constant returns(address)
func (_CompoundLogic *CompoundLogicCaller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CompoundLogic.contract.Call(opts, out, "comptroller")
	return *ret0, err
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() constant returns(address)
func (_CompoundLogic *CompoundLogicSession) Comptroller() (common.Address, error) {
	return _CompoundLogic.Contract.Comptroller(&_CompoundLogic.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() constant returns(address)
func (_CompoundLogic *CompoundLogicCallerSession) Comptroller() (common.Address, error) {
	return _CompoundLogic.Contract.Comptroller(&_CompoundLogic.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_CompoundLogic *CompoundLogicCaller) IncognitoSmartContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CompoundLogic.contract.Call(opts, out, "incognitoSmartContract")
	return *ret0, err
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_CompoundLogic *CompoundLogicSession) IncognitoSmartContract() (common.Address, error) {
	return _CompoundLogic.Contract.IncognitoSmartContract(&_CompoundLogic.CallOpts)
}

// IncognitoSmartContract is a free data retrieval call binding the contract method 0xb42a644b.
//
// Solidity: function incognitoSmartContract() constant returns(address)
func (_CompoundLogic *CompoundLogicCallerSession) IncognitoSmartContract() (common.Address, error) {
	return _CompoundLogic.Contract.IncognitoSmartContract(&_CompoundLogic.CallOpts)
}

// Borrow is a paid mutator transaction binding the contract method 0x26d5e251.
//
// Solidity: function borrow(address cToken, uint256 amount, address[] addToMarkets) returns(address, uint256)
func (_CompoundLogic *CompoundLogicTransactor) Borrow(opts *bind.TransactOpts, cToken common.Address, amount *big.Int, addToMarkets []common.Address) (*types.Transaction, error) {
	return _CompoundLogic.contract.Transact(opts, "borrow", cToken, amount, addToMarkets)
}

// Borrow is a paid mutator transaction binding the contract method 0x26d5e251.
//
// Solidity: function borrow(address cToken, uint256 amount, address[] addToMarkets) returns(address, uint256)
func (_CompoundLogic *CompoundLogicSession) Borrow(cToken common.Address, amount *big.Int, addToMarkets []common.Address) (*types.Transaction, error) {
	return _CompoundLogic.Contract.Borrow(&_CompoundLogic.TransactOpts, cToken, amount, addToMarkets)
}

// Borrow is a paid mutator transaction binding the contract method 0x26d5e251.
//
// Solidity: function borrow(address cToken, uint256 amount, address[] addToMarkets) returns(address, uint256)
func (_CompoundLogic *CompoundLogicTransactorSession) Borrow(cToken common.Address, amount *big.Int, addToMarkets []common.Address) (*types.Transaction, error) {
	return _CompoundLogic.Contract.Borrow(&_CompoundLogic.TransactOpts, cToken, amount, addToMarkets)
}

// BorrowByMultiCollateral is a paid mutator transaction binding the contract method 0xb6dbc8ed.
//
// Solidity: function borrowByMultiCollateral(address cToken, uint256 amount, address[] addToMarkets) returns(address[], uint256[])
func (_CompoundLogic *CompoundLogicTransactor) BorrowByMultiCollateral(opts *bind.TransactOpts, cToken common.Address, amount *big.Int, addToMarkets []common.Address) (*types.Transaction, error) {
	return _CompoundLogic.contract.Transact(opts, "borrowByMultiCollateral", cToken, amount, addToMarkets)
}

// BorrowByMultiCollateral is a paid mutator transaction binding the contract method 0xb6dbc8ed.
//
// Solidity: function borrowByMultiCollateral(address cToken, uint256 amount, address[] addToMarkets) returns(address[], uint256[])
func (_CompoundLogic *CompoundLogicSession) BorrowByMultiCollateral(cToken common.Address, amount *big.Int, addToMarkets []common.Address) (*types.Transaction, error) {
	return _CompoundLogic.Contract.BorrowByMultiCollateral(&_CompoundLogic.TransactOpts, cToken, amount, addToMarkets)
}

// BorrowByMultiCollateral is a paid mutator transaction binding the contract method 0xb6dbc8ed.
//
// Solidity: function borrowByMultiCollateral(address cToken, uint256 amount, address[] addToMarkets) returns(address[], uint256[])
func (_CompoundLogic *CompoundLogicTransactorSession) BorrowByMultiCollateral(cToken common.Address, amount *big.Int, addToMarkets []common.Address) (*types.Transaction, error) {
	return _CompoundLogic.Contract.BorrowByMultiCollateral(&_CompoundLogic.TransactOpts, cToken, amount, addToMarkets)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0x64fd7078.
//
// Solidity: function liquidateBorrow(address cToken, address borrower, uint256 repayAmount, address cTokenCollateral) returns(address, uint256)
func (_CompoundLogic *CompoundLogicTransactor) LiquidateBorrow(opts *bind.TransactOpts, cToken common.Address, borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _CompoundLogic.contract.Transact(opts, "liquidateBorrow", cToken, borrower, repayAmount, cTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0x64fd7078.
//
// Solidity: function liquidateBorrow(address cToken, address borrower, uint256 repayAmount, address cTokenCollateral) returns(address, uint256)
func (_CompoundLogic *CompoundLogicSession) LiquidateBorrow(cToken common.Address, borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _CompoundLogic.Contract.LiquidateBorrow(&_CompoundLogic.TransactOpts, cToken, borrower, repayAmount, cTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0x64fd7078.
//
// Solidity: function liquidateBorrow(address cToken, address borrower, uint256 repayAmount, address cTokenCollateral) returns(address, uint256)
func (_CompoundLogic *CompoundLogicTransactorSession) LiquidateBorrow(cToken common.Address, borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _CompoundLogic.Contract.LiquidateBorrow(&_CompoundLogic.TransactOpts, cToken, borrower, repayAmount, cTokenCollateral)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address cToken, uint256 amount) returns(address, uint256)
func (_CompoundLogic *CompoundLogicTransactor) Mint(opts *bind.TransactOpts, cToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CompoundLogic.contract.Transact(opts, "mint", cToken, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address cToken, uint256 amount) returns(address, uint256)
func (_CompoundLogic *CompoundLogicSession) Mint(cToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CompoundLogic.Contract.Mint(&_CompoundLogic.TransactOpts, cToken, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address cToken, uint256 amount) returns(address, uint256)
func (_CompoundLogic *CompoundLogicTransactorSession) Mint(cToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CompoundLogic.Contract.Mint(&_CompoundLogic.TransactOpts, cToken, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0x7520f7ed.
//
// Solidity: function redeem(address cToken, uint256 amount, bool isredeemUnderlying, address exitToMarkets) returns(address, uint256)
func (_CompoundLogic *CompoundLogicTransactor) Redeem(opts *bind.TransactOpts, cToken common.Address, amount *big.Int, isredeemUnderlying bool, exitToMarkets common.Address) (*types.Transaction, error) {
	return _CompoundLogic.contract.Transact(opts, "redeem", cToken, amount, isredeemUnderlying, exitToMarkets)
}

// Redeem is a paid mutator transaction binding the contract method 0x7520f7ed.
//
// Solidity: function redeem(address cToken, uint256 amount, bool isredeemUnderlying, address exitToMarkets) returns(address, uint256)
func (_CompoundLogic *CompoundLogicSession) Redeem(cToken common.Address, amount *big.Int, isredeemUnderlying bool, exitToMarkets common.Address) (*types.Transaction, error) {
	return _CompoundLogic.Contract.Redeem(&_CompoundLogic.TransactOpts, cToken, amount, isredeemUnderlying, exitToMarkets)
}

// Redeem is a paid mutator transaction binding the contract method 0x7520f7ed.
//
// Solidity: function redeem(address cToken, uint256 amount, bool isredeemUnderlying, address exitToMarkets) returns(address, uint256)
func (_CompoundLogic *CompoundLogicTransactorSession) Redeem(cToken common.Address, amount *big.Int, isredeemUnderlying bool, exitToMarkets common.Address) (*types.Transaction, error) {
	return _CompoundLogic.Contract.Redeem(&_CompoundLogic.TransactOpts, cToken, amount, isredeemUnderlying, exitToMarkets)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0xabdb5ea8.
//
// Solidity: function repayBorrow(address cToken, uint256 amount) returns(address, uint256)
func (_CompoundLogic *CompoundLogicTransactor) RepayBorrow(opts *bind.TransactOpts, cToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CompoundLogic.contract.Transact(opts, "repayBorrow", cToken, amount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0xabdb5ea8.
//
// Solidity: function repayBorrow(address cToken, uint256 amount) returns(address, uint256)
func (_CompoundLogic *CompoundLogicSession) RepayBorrow(cToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CompoundLogic.Contract.RepayBorrow(&_CompoundLogic.TransactOpts, cToken, amount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0xabdb5ea8.
//
// Solidity: function repayBorrow(address cToken, uint256 amount) returns(address, uint256)
func (_CompoundLogic *CompoundLogicTransactorSession) RepayBorrow(cToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CompoundLogic.Contract.RepayBorrow(&_CompoundLogic.TransactOpts, cToken, amount)
}
