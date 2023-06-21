// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// ZKMerkleTrieMetaData contains all meta data concerning the ZKMerkleTrie contract.
var ZKMerkleTrieMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poseidon2\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"POSEIDON2\",\"outputs\":[{\"internalType\":\"contractIPoseidon2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"_proofs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_proofs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"verifyInclusionProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051611ceb380380611ceb83398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b608051611c456100a66000396000818160940152818161091a01528181610b6401528181610c600152610dce0152611c456000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806312e64a7214610046578063c423b1e81461006e578063dc8b50381461008f575b600080fd5b610059610054366004611879565b6100db565b60405190151581526020015b60405180910390f35b61008161007c3660046118ee565b6101a6565b6040516100659291906119b8565b6100b67f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610065565b6040517fc423b1e800000000000000000000000000000000000000000000000000000000815260009081908190309063c423b1e890610122908a90899089906004016119d3565b600060405180830381865afa15801561013f573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526101859190810190611a63565b9150915081801561019b575061019b868261089b565b979650505050505050565b600060606002845110156102275760405162461bcd60e51b815260206004820152602960248201527f5a4b4d65726b6c65547269653a2070726f76696465642070726f6f662069732060448201527f746f6f2073686f7274000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b61027e84600186516102399190611b24565b8151811061024957610249611b3b565b602002602001015180516020909101207f950654da67865a81bc70e45f3230f5179f08e29c66184bf746f71050f117b3b81490565b6102f05760405162461bcd60e51b815260206004820152602d60248201527f5a4b4d65726b6c65547269653a20746865206c617374206974656d206973206e60448201527f6f74206d61676963206861736800000000000000000000000000000000000000606482015260840161021e565b60006102fb866108b7565b9050600061030886610996565b90506103526040805161010081019091528060008152600060208201819052604082018190526060808301829052608083015260a0820181905260c0820181905260e09091015290565b60408051602081019091526000808252835190918291829190829061037990600290611b24565b90505b86818151811061038e5761038e611b3b565b6020026020010151602001519550600060038111156103af576103af611b6a565b865160038111156103c2576103c2611b6a565b036105015760006103d38983610a89565b90508015610433578660200151861461042e5760405162461bcd60e51b815260206004820152601b60248201527f5a4b4d65726b6c65547269653a20696e76616c6964206b6579204c0000000000604482015260640161021e565b610486565b866040015186146104865760405162461bcd60e51b815260206004820152601b60248201527f5a4b4d65726b6c65547269653a20696e76616c6964206b657920520000000000604482015260640161021e565b6040805160028082526060820183526000926020830190803683370190505090508760200151816000815181106104bf576104bf611b3b565b6020026020010181815250508760400151816001815181106104e3576104e3611b3b565b6020026020010181815250506104f881610ae6565b96505050610805565b60018651600381111561051657610516611b6a565b0361077257831561058f5760405162461bcd60e51b815260206004820152602260248201527f5a4b4d65726b6c65547269653a206475706c696361746564206c656166206e6f60448201527f6465000000000000000000000000000000000000000000000000000000000000606482015260840161021e565b604080516003808252608082019092526001955060009160208201606080368337019050509050600160001b816000815181106105ce576105ce611b3b565b6020026020010181815250508660600151816001815181106105f2576105f2611b3b565b6020026020010181815250506106108760a001518860800151610f2f565b8160028151811061062357610623611b3b565b60200260200101818152505061063881610ae6565b9550866080015151602061064c9190611b99565b67ffffffffffffffff811115610664576106646116c0565b6040519080825280601f01601f19166020018201604052801561068e576020820181803683370190505b509250602060005b8860800151518110156106d9576000896080015182815181106106bb576106bb611b3b565b60209081029190910181015187850152929092019150600101610696565b5060e08801511561076b578e8860e0015114806106f95750898860e00151145b61076b5760405162461bcd60e51b815260206004820152602260248201527f5a4b4d65726b6c65547269653a20696e76616c6964206b657920707265696d6160448201527f6765000000000000000000000000000000000000000000000000000000000000606482015260840161021e565b5050610805565b60028651600381111561078757610787611b6a565b036108055782156108005760405162461bcd60e51b815260206004820152602360248201527f5a4b4d65726b6c65547269653a206475706c69636174656420656d707479206e60448201527f6f64650000000000000000000000000000000000000000000000000000000000606482015260840161021e565b600192505b80600003610861578a851461085c5760405162461bcd60e51b815260206004820152601960248201527f5a4b4d65726b65547269653a20696e76616c696420726f6f7400000000000000604482015260640161021e565b610888565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161037c565b50919b919a509098505050505050505050565b6000818051906020012083805190602001201490505b92915050565b60008060006108c5846110cb565b6040805180820182528381526020810183905290517f299e566000000000000000000000000000000000000000000000000000000000815292945090925073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169163299e56609161094d91600401611bd6565b602060405180830381865afa15801561096a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061098e9190611c07565b949350505050565b805160609060008167ffffffffffffffff8111156109b6576109b66116c0565b6040519080825280602002602001820160405280156109ef57816020015b6109dc611661565b8152602001906001900390816109d45790505b50905060005b610a00600184611b24565b811015610a81576000610a2b868381518110610a1e57610a1e611b3b565b60200260200101516110f3565b90506040518060400160405280878481518110610a4a57610a4a611b3b565b6020026020010151815260200182815250838381518110610a6d57610a6d611b3b565b6020908102919091010152506001016109f5565b509392505050565b60006101008210610adc5760405162461bcd60e51b815260206004820152601c60248201527f5a4b4d65726b6c65547269653a20746f6f206c6f6e6720646570746800000000604482015260640161021e565b506001901b161590565b6000600282511015610b605760405162461bcd60e51b815260206004820152602b60248201527f5a4b547269654861736865723a20746f6f206665772076616c75657320666f7260448201527f205f68617368456c656d73000000000000000000000000000000000000000000606482015260840161021e565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663299e5660604051806040016040528086600081518110610bbc57610bbc611b3b565b6020026020010151815260200186600181518110610bdc57610bdc611b3b565b60200260200101518152506040518263ffffffff1660e01b8152600401610c039190611bd6565b602060405180830381865afa158015610c20573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c449190611c07565b90508251600203610c555792915050565b8251600303610d2d577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663299e5660604051806040016040528084815260200186600281518110610cbe57610cbe611b3b565b60200260200101518152506040518263ffffffff1660e01b8152600401610ce59190611bd6565b602060405180830381865afa158015610d02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d269190611c07565b9392505050565b6000600184516001610d3f9190611c20565b901c67ffffffffffffffff811115610d5957610d596116c0565b604051908082528060200260200182016040528015610d82578160200160208202803683370190505b5090508181600081518110610d9957610d99611b3b565b6020908102919091010152600060015b8251811015610f1c578551600182811b9350610dc6908490611c20565b1015610edb577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663299e56606040518060400160405280898681518110610e2557610e25611b3b565b6020026020010151815260200189866001610e409190611c20565b81518110610e5057610e50611b3b565b60200260200101518152506040518263ffffffff1660e01b8152600401610e779190611bd6565b602060405180830381865afa158015610e94573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eb89190611c07565b838281518110610eca57610eca611b3b565b602002602001018181525050610f14565b858281518110610eed57610eed611b3b565b6020026020010151838281518110610f0757610f07611b3b565b6020026020010181815250505b600101610da9565b50610f2682610ae6565b95945050505050565b6000600182511015610fa95760405162461bcd60e51b815260206004820152602b60248201527f5a4b547269654861736865723a20746f6f206665772076616c75657320666f7260448201527f205f76616c756548617368000000000000000000000000000000000000000000606482015260840161021e565b6000825167ffffffffffffffff811115610fc557610fc56116c0565b604051908082528060200260200182016040528015610fee578160200160208202803683370190505b50905060005b8351811015611094576001811b851663ffffffff16156110535761103084828151811061102357611023611b3b565b60200260200101516108b7565b82828151811061104257611042611b3b565b60200260200101818152505061108c565b83818151811061106557611065611b3b565b602002602001015182828151811061107f5761107f611b3b565b6020026020010181815250505b600101610ff4565b506002835110156110c257806000815181106110b2576110b2611b3b565b60200260200101519150506108b1565b61098e81610ae6565b60008082608081901b6110de8260801c90565b6110e88260801c90565b935093505050915091565b60408051610100810182526000808252602082018190529181018290526060808201839052608082015260a0810182905260c0810182905260e081019190915260408051610100810182526000808252602082018190529181018290526060808201839052608082015260a0810182905260c0810182905260e081019190915260006111a7846040805180820182526060815260006020918201528151808301909252828101825291519181019190915290565b905060006111b4826113d3565b60ff169050806111df576111c78261145a565b60208401526111d58261145a565b604084015261138f565b600181036112c9576111f08261145a565b6060840152600080611201846114f8565b63ffffffff821660a088015290925060ff1690508067ffffffffffffffff81111561122e5761122e6116c0565b604051908082528060200260200182016040528015611257578160200160208202803683370190505b50608086015260005b8181101561129a576112718561145a565b8660800151828151811061128757611287611b3b565b6020908102919091010152600101611260565b5060006112a6856113d3565b60ff16905080156112c1576112bb8582611592565b60e08701525b50505061138f565b6002811461138f57600381036113475760405162461bcd60e51b815260206004820152602560248201527f4e6f64655265616465723a20756e657870656374656420726f6f74206e6f646560448201527f2074797065000000000000000000000000000000000000000000000000000000606482015260840161021e565b60405162461bcd60e51b815260206004820152601d60248201527f4e6f64655265616465723a20696e76616c6964206e6f64652074797065000000604482015260640161021e565b8060038111156113a1576113a1611b6a565b839060038111156113b4576113b4611b6a565b908160038111156113c7576113c7611b6a565b90525091949350505050565b600060018260200151101561142a5760405162461bcd60e51b815260206004820152601f60248201527f4e6f64655265616465723a20746f6f2073686f727420666f722075696e743800604482015260640161021e565b81518051600180830180865260208601805191949360f81c929161144f908390611b24565b905250949350505050565b60006020826020015110156114d75760405162461bcd60e51b815260206004820152602160248201527f4e6f64655265616465723a20746f6f2073686f727420666f722062797465733360448201527f3200000000000000000000000000000000000000000000000000000000000000606482015260840161021e565b8151805160208083018086528186018051919493929161144f908390611b24565b6000806004836020015110156115505760405162461bcd60e51b815260206004820181905260248201527f4e6f64655265616465723a20746f6f2073686f727420666f722075696e743332604482015260640161021e565b8251805160048083018087526020870180519194939260f084901c9260f885901c929061157e908390611b24565b90525060ff90911697909650945050505050565b6000818360200151101561160e5760405162461bcd60e51b815260206004820152602160248201527f4e6f64655265616465723a20746f6f2073686f727420666f72206e206279746560448201527f7300000000000000000000000000000000000000000000000000000000000000606482015260840161021e565b8251606090600080611621866008611b99565b61162d90610100611b24565b8351848801808a5260208a01805191975091831c94509192508791611653908390611b24565b905250909695505050505050565b6040518060400160405280606081526020016116bb6040805161010081019091528060008152600060208201819052604082018190526060808301829052608083015260a0820181905260c0820181905260e09091015290565b905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611736576117366116c0565b604052919050565b600067ffffffffffffffff821115611758576117586116c0565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f83011261179557600080fd5b81356117a86117a38261173e565b6116ef565b8181528460208386010111156117bd57600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f8301126117eb57600080fd5b8135602067ffffffffffffffff80831115611808576118086116c0565b8260051b6118178382016116ef565b938452858101830193838101908886111561183157600080fd5b84880192505b8583101561186d5782358481111561184f5760008081fd5b61185d8a87838c0101611784565b8352509184019190840190611837565b98975050505050505050565b6000806000806080858703121561188f57600080fd5b84359350602085013567ffffffffffffffff808211156118ae57600080fd5b6118ba88838901611784565b945060408701359150808211156118d057600080fd5b506118dd878288016117da565b949793965093946060013593505050565b60008060006060848603121561190357600080fd5b83359250602084013567ffffffffffffffff81111561192157600080fd5b61192d868287016117da565b925050604084013590509250925092565b60005b83811015611959578181015183820152602001611941565b83811115611968576000848401525b50505050565b6000815180845261198681602086016020860161193e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b821515815260406020820152600061098e604083018461196e565b600060608201858352602060608185015281865180845260808601915060808160051b870101935082880160005b82811015611a4d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80888703018452611a3b86835161196e565b95509284019290840190600101611a01565b5050505050604092909201929092529392505050565b60008060408385031215611a7657600080fd5b82518015158114611a8657600080fd5b602084015190925067ffffffffffffffff811115611aa357600080fd5b8301601f81018513611ab457600080fd5b8051611ac26117a38261173e565b818152866020838501011115611ad757600080fd5b611ae882602083016020860161193e565b8093505050509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015611b3657611b36611af5565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611bd157611bd1611af5565b500290565b60408101818360005b6002811015611bfe578151835260209283019290910190600101611bdf565b50505092915050565b600060208284031215611c1957600080fd5b5051919050565b60008219821115611c3357611c33611af5565b50019056fea164736f6c634300080f000a",
}

// ZKMerkleTrieABI is the input ABI used to generate the binding from.
// Deprecated: Use ZKMerkleTrieMetaData.ABI instead.
var ZKMerkleTrieABI = ZKMerkleTrieMetaData.ABI

// ZKMerkleTrieBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZKMerkleTrieMetaData.Bin instead.
var ZKMerkleTrieBin = ZKMerkleTrieMetaData.Bin

// DeployZKMerkleTrie deploys a new Ethereum contract, binding an instance of ZKMerkleTrie to it.
func DeployZKMerkleTrie(auth *bind.TransactOpts, backend bind.ContractBackend, _poseidon2 common.Address) (common.Address, *types.Transaction, *ZKMerkleTrie, error) {
	parsed, err := ZKMerkleTrieMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZKMerkleTrieBin), backend, _poseidon2)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZKMerkleTrie{ZKMerkleTrieCaller: ZKMerkleTrieCaller{contract: contract}, ZKMerkleTrieTransactor: ZKMerkleTrieTransactor{contract: contract}, ZKMerkleTrieFilterer: ZKMerkleTrieFilterer{contract: contract}}, nil
}

// ZKMerkleTrie is an auto generated Go binding around an Ethereum contract.
type ZKMerkleTrie struct {
	ZKMerkleTrieCaller     // Read-only binding to the contract
	ZKMerkleTrieTransactor // Write-only binding to the contract
	ZKMerkleTrieFilterer   // Log filterer for contract events
}

// ZKMerkleTrieCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZKMerkleTrieCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKMerkleTrieTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZKMerkleTrieTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKMerkleTrieFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZKMerkleTrieFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKMerkleTrieSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZKMerkleTrieSession struct {
	Contract     *ZKMerkleTrie     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZKMerkleTrieCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZKMerkleTrieCallerSession struct {
	Contract *ZKMerkleTrieCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ZKMerkleTrieTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZKMerkleTrieTransactorSession struct {
	Contract     *ZKMerkleTrieTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ZKMerkleTrieRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZKMerkleTrieRaw struct {
	Contract *ZKMerkleTrie // Generic contract binding to access the raw methods on
}

// ZKMerkleTrieCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZKMerkleTrieCallerRaw struct {
	Contract *ZKMerkleTrieCaller // Generic read-only contract binding to access the raw methods on
}

// ZKMerkleTrieTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZKMerkleTrieTransactorRaw struct {
	Contract *ZKMerkleTrieTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZKMerkleTrie creates a new instance of ZKMerkleTrie, bound to a specific deployed contract.
func NewZKMerkleTrie(address common.Address, backend bind.ContractBackend) (*ZKMerkleTrie, error) {
	contract, err := bindZKMerkleTrie(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZKMerkleTrie{ZKMerkleTrieCaller: ZKMerkleTrieCaller{contract: contract}, ZKMerkleTrieTransactor: ZKMerkleTrieTransactor{contract: contract}, ZKMerkleTrieFilterer: ZKMerkleTrieFilterer{contract: contract}}, nil
}

// NewZKMerkleTrieCaller creates a new read-only instance of ZKMerkleTrie, bound to a specific deployed contract.
func NewZKMerkleTrieCaller(address common.Address, caller bind.ContractCaller) (*ZKMerkleTrieCaller, error) {
	contract, err := bindZKMerkleTrie(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZKMerkleTrieCaller{contract: contract}, nil
}

// NewZKMerkleTrieTransactor creates a new write-only instance of ZKMerkleTrie, bound to a specific deployed contract.
func NewZKMerkleTrieTransactor(address common.Address, transactor bind.ContractTransactor) (*ZKMerkleTrieTransactor, error) {
	contract, err := bindZKMerkleTrie(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZKMerkleTrieTransactor{contract: contract}, nil
}

// NewZKMerkleTrieFilterer creates a new log filterer instance of ZKMerkleTrie, bound to a specific deployed contract.
func NewZKMerkleTrieFilterer(address common.Address, filterer bind.ContractFilterer) (*ZKMerkleTrieFilterer, error) {
	contract, err := bindZKMerkleTrie(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZKMerkleTrieFilterer{contract: contract}, nil
}

// bindZKMerkleTrie binds a generic wrapper to an already deployed contract.
func bindZKMerkleTrie(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZKMerkleTrieMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKMerkleTrie *ZKMerkleTrieRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKMerkleTrie.Contract.ZKMerkleTrieCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKMerkleTrie *ZKMerkleTrieRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKMerkleTrie.Contract.ZKMerkleTrieTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKMerkleTrie *ZKMerkleTrieRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKMerkleTrie.Contract.ZKMerkleTrieTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKMerkleTrie *ZKMerkleTrieCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKMerkleTrie.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKMerkleTrie *ZKMerkleTrieTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKMerkleTrie.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKMerkleTrie *ZKMerkleTrieTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKMerkleTrie.Contract.contract.Transact(opts, method, params...)
}

// POSEIDON2 is a free data retrieval call binding the contract method 0xdc8b5038.
//
// Solidity: function POSEIDON2() view returns(address)
func (_ZKMerkleTrie *ZKMerkleTrieCaller) POSEIDON2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKMerkleTrie.contract.Call(opts, &out, "POSEIDON2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// POSEIDON2 is a free data retrieval call binding the contract method 0xdc8b5038.
//
// Solidity: function POSEIDON2() view returns(address)
func (_ZKMerkleTrie *ZKMerkleTrieSession) POSEIDON2() (common.Address, error) {
	return _ZKMerkleTrie.Contract.POSEIDON2(&_ZKMerkleTrie.CallOpts)
}

// POSEIDON2 is a free data retrieval call binding the contract method 0xdc8b5038.
//
// Solidity: function POSEIDON2() view returns(address)
func (_ZKMerkleTrie *ZKMerkleTrieCallerSession) POSEIDON2() (common.Address, error) {
	return _ZKMerkleTrie.Contract.POSEIDON2(&_ZKMerkleTrie.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0xc423b1e8.
//
// Solidity: function get(bytes32 _key, bytes[] _proofs, bytes32 _root) view returns(bool, bytes)
func (_ZKMerkleTrie *ZKMerkleTrieCaller) Get(opts *bind.CallOpts, _key [32]byte, _proofs [][]byte, _root [32]byte) (bool, []byte, error) {
	var out []interface{}
	err := _ZKMerkleTrie.contract.Call(opts, &out, "get", _key, _proofs, _root)

	if err != nil {
		return *new(bool), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// Get is a free data retrieval call binding the contract method 0xc423b1e8.
//
// Solidity: function get(bytes32 _key, bytes[] _proofs, bytes32 _root) view returns(bool, bytes)
func (_ZKMerkleTrie *ZKMerkleTrieSession) Get(_key [32]byte, _proofs [][]byte, _root [32]byte) (bool, []byte, error) {
	return _ZKMerkleTrie.Contract.Get(&_ZKMerkleTrie.CallOpts, _key, _proofs, _root)
}

// Get is a free data retrieval call binding the contract method 0xc423b1e8.
//
// Solidity: function get(bytes32 _key, bytes[] _proofs, bytes32 _root) view returns(bool, bytes)
func (_ZKMerkleTrie *ZKMerkleTrieCallerSession) Get(_key [32]byte, _proofs [][]byte, _root [32]byte) (bool, []byte, error) {
	return _ZKMerkleTrie.Contract.Get(&_ZKMerkleTrie.CallOpts, _key, _proofs, _root)
}

// VerifyInclusionProof is a free data retrieval call binding the contract method 0x12e64a72.
//
// Solidity: function verifyInclusionProof(bytes32 _key, bytes _value, bytes[] _proofs, bytes32 _root) view returns(bool)
func (_ZKMerkleTrie *ZKMerkleTrieCaller) VerifyInclusionProof(opts *bind.CallOpts, _key [32]byte, _value []byte, _proofs [][]byte, _root [32]byte) (bool, error) {
	var out []interface{}
	err := _ZKMerkleTrie.contract.Call(opts, &out, "verifyInclusionProof", _key, _value, _proofs, _root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyInclusionProof is a free data retrieval call binding the contract method 0x12e64a72.
//
// Solidity: function verifyInclusionProof(bytes32 _key, bytes _value, bytes[] _proofs, bytes32 _root) view returns(bool)
func (_ZKMerkleTrie *ZKMerkleTrieSession) VerifyInclusionProof(_key [32]byte, _value []byte, _proofs [][]byte, _root [32]byte) (bool, error) {
	return _ZKMerkleTrie.Contract.VerifyInclusionProof(&_ZKMerkleTrie.CallOpts, _key, _value, _proofs, _root)
}

// VerifyInclusionProof is a free data retrieval call binding the contract method 0x12e64a72.
//
// Solidity: function verifyInclusionProof(bytes32 _key, bytes _value, bytes[] _proofs, bytes32 _root) view returns(bool)
func (_ZKMerkleTrie *ZKMerkleTrieCallerSession) VerifyInclusionProof(_key [32]byte, _value []byte, _proofs [][]byte, _root [32]byte) (bool, error) {
	return _ZKMerkleTrie.Contract.VerifyInclusionProof(&_ZKMerkleTrie.CallOpts, _key, _value, _proofs, _root)
}
