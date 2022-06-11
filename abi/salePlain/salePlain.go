// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package salePlain

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

// SalePlainABI is the input ABI used to generate the binding from.
const SalePlainABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"master\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"CancelSale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"solder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"ConfirmSale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_due\",\"type\":\"uint256\"}],\"name\":\"CreateSale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"NewBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_due\",\"type\":\"uint256\"}],\"name\":\"UpdateSale\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"bidWithToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleId\",\"type\":\"uint256\"}],\"name\":\"bidWithValue\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleId\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleId\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleId\",\"type\":\"uint256\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_due\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"record\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getSaleId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleId\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getSaleInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleId\",\"type\":\"uint256\"}],\"name\":\"getSaleInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isOpen\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_due\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"masterAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_due\",\"type\":\"uint256\"}],\"name\":\"updateDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SalePlainBin is the compiled bytecode used for deploying new contracts.
const SalePlainBin = `608060405234801561001057600080fd5b50604051611a97380380611a9783398101604081905261002f91610054565b600180546001600160a01b0319166001600160a01b0392909216919091179055610082565b600060208284031215610065578081fd5b81516001600160a01b038116811461007b578182fd5b9392505050565b611a06806100916000396000f3fe6080604052600436106100a75760003560e01c806340e58ee51161006457806340e58ee51461019657806364b3b844146101b65780638ff05a2a146101d6578063d1d58b25146101f6578063d365a08e14610223578063faa21b1114610245576100a7565b8063087db5b8146100ac5780630f0a1b9e146100e75780632ad1f1db14610109578063370b008114610129578063379607f5146101565780633969e1ae14610176575b600080fd5b3480156100b857600080fd5b506100cc6100c7366004611534565b610258565b6040516100de969594939291906116f4565b60405180910390f35b3480156100f357600080fd5b50610107610102366004611640565b610283565b005b34801561011557600080fd5b50610107610124366004611677565b61042d565b34801561013557600080fd5b5061014961014436600461158c565b610544565b6040516100de919061184c565b34801561016257600080fd5b50610107610171366004611610565b610796565b34801561018257600080fd5b50610149610191366004611534565b610799565b3480156101a257600080fd5b506101076101b1366004611610565b6107b6565b3480156101c257600080fd5b506100cc6101d1366004611610565b61094c565b3480156101e257600080fd5b506101076101f1366004611677565b61099c565b34801561020257600080fd5b50610216610211366004611610565b610ffa565b6040516100de91906116e9565b34801561022f57600080fd5b50610238611000565b6040516100de9190611698565b610107610253366004611610565b61100f565b60008060008060008061026e6101d18989610799565b949d939c50919a509850965090945092505050565b60008061028f856114dc565b6000878152602081905260409020549193506001600160601b031691506001600160a01b031633146102dc5760405162461bcd60e51b81526004016102d3906117a8565b60405180910390fd5b600154604051630e70619560e41b81526001600160a01b039091169063e70619509061030c908790600401611698565b60206040518083038186803b15801561032457600080fd5b505afa158015610338573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061035c91906115f0565b6103785760405162461bcd60e51b81526004016102d390611800565b600085815260208190526040908190206001810180546001600160a01b0319166001600160a01b038881169190911791829055600290920180546001600160801b0319166001600160801b0388811691909117918290559351868416947ffe8937da95814c69018f7a05dcec87068d26502ab6dabf6516ea314dc1dca32b9461041e94889433949190921692821691600160801b900467ffffffffffffffff169061189a565b60405180910390a25050505050565b600080610439846114dc565b6000868152602081905260409020549193506001600160601b031691506001600160a01b0316331461047d5760405162461bcd60e51b81526004016102d3906117a8565b42831161049c5760405162461bcd60e51b81526004016102d3906117d8565b6000848152602081905260409081902060028101805467ffffffffffffffff60801b1916600160801b67ffffffffffffffff8881168202929092179283905560019093015493516001600160a01b03878116957ffe8937da95814c69018f7a05dcec87068d26502ab6dabf6516ea314dc1dca32b95610536958995339594909316936001600160801b03821693929091049091169061189a565b60405180910390a250505050565b6040516331a9108f60e11b815260009030906001600160a01b03891690636352211e90610575908a9060040161184c565b60206040518083038186803b15801561058d57600080fd5b505afa1580156105a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105c59190611518565b6001600160a01b0316146105d857600080fd5b4282116105f75760405162461bcd60e51b81526004016102d3906117d8565b826001600160801b0316831461061f5760405162461bcd60e51b81526004016102d390611752565b61062987876114e4565b604080516080810182526001600160a01b03808916825287811660208084019182526001600160801b03808a1685870190815267ffffffffffffffff808b166060880190815260008a81529485905293889020965187549087166001600160a01b031991821617885594516001808901805492891692909716919091179095559051600290960180549351909116600160801b0267ffffffffffffffff60801b19969092166001600160801b0319909316929092179490941693909317909255905491516308d3fed360e01b815292935016906308d3fed3906107129084908990600401611855565b600060405180830381600087803b15801561072c57600080fd5b505af1158015610740573d6000803e3d6000fd5b50505050866001600160a01b03167e8f298a385d495e86787887da324d1603491ea0ea157acffd6c5851d6ddfa2f87878787876040516107849594939291906118db565b60405180910390a29695505050505050565b50565b60006107af838367ffffffffffffffff166114e4565b9392505050565b6000806107c2836114dc565b6000858152602081905260409020549193506001600160601b031691506001600160a01b031633146108065760405162461bcd60e51b81526004016102d3906117a8565b6000838152602081905260409081902080546001600160a01b03199081168255600182018054909116905560020180546001600160c01b0319169055516323b872dd60e01b81526001600160a01b038316906323b872dd90610870903090339086906004016116ac565b600060405180830381600087803b15801561088a57600080fd5b505af115801561089e573d6000803e3d6000fd5b5050600154604051632084212160e11b81526001600160a01b039091169250634108424291506108d49086903390600401611855565b600060405180830381600087803b1580156108ee57600080fd5b505af1158015610902573d6000803e3d6000fd5b50505050816001600160a01b03167f8b0a052bd72fec557d5735593db719bcb565d7afec8d76ee165e317e8cc984098260405161093f919061184c565b60405180910390a2505050565b600090815260208190526040812060018101546002820154915442600160801b840467ffffffffffffffff16908111956001600160a01b03938416956001600160801b0390951694919390921691565b6000806109a8846114dc565b6000868152602081905260409020600201549193506001600160601b0316915042600160801b90910467ffffffffffffffff16116109f85760405162461bcd60e51b81526004016102d39061172c565b6000848152602081905260409020600101546001600160a01b0316610a2f5760405162461bcd60e51b81526004016102d390611825565b6000848152602081905260409020600201546001600160801b0316831015610a695760405162461bcd60e51b81526004016102d39061177a565b600084815260208190526040908190206001015490516323b872dd60e01b81526001600160a01b03909116906323b872dd90610aad903390309088906004016116ac565b602060405180830381600087803b158015610ac757600080fd5b505af1158015610adb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aff91906115f0565b50600080836001600160a01b0316632a55205a84876040518363ffffffff1660e01b8152600401610b31929190611929565b604080518083038186803b158015610b4857600080fd5b505afa158015610b5c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b80919061155f565b91509150600061271086600160009054906101000a90046001600160a01b03166001600160a01b0316630e716d976040518163ffffffff1660e01b815260040160206040518083038186803b158015610bd857600080fd5b505afa158015610bec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c109190611628565b610c1a919061196f565b610c24919061194f565b60008881526020818152604091829020600190810154905483516331056e5760e21b815293519495506001600160a01b039182169463a9059cbb94919092169263c415b95c92600480840193829003018186803b158015610c8457600080fd5b505afa158015610c98573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cbc9190611518565b836040518363ffffffff1660e01b8152600401610cda9291906116d0565b602060405180830381600087803b158015610cf457600080fd5b505af1158015610d08573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d2c91906115f0565b506000878152602081905260409081902060010154905163a9059cbb60e01b81526001600160a01b039091169063a9059cbb90610d6f90869086906004016116d0565b602060405180830381600087803b158015610d8957600080fd5b505af1158015610d9d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dc191906115f0565b506000878152602081905260409020600181015490546001600160a01b039182169163a9059cbb911684610df5858b61198e565b610dff919061198e565b6040518363ffffffff1660e01b8152600401610e1c9291906116d0565b602060405180830381600087803b158015610e3657600080fd5b505af1158015610e4a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e6e91906115f0565b506040516323b872dd60e01b81526001600160a01b038616906323b872dd90610e9f903090339089906004016116ac565b600060405180830381600087803b158015610eb957600080fd5b505af1158015610ecd573d6000803e3d6000fd5b50506001805460008b81526020819052604090819020805493015490516346eb67ef60e01b81526001600160a01b0392831695506346eb67ef9450610f21938d939081169233929116908d9060040161186c565b600060405180830381600087803b158015610f3b57600080fd5b505af1158015610f4f573d6000803e3d6000fd5b50505060008881526020819052604090819020805460019091015491513393506001600160a01b0391821692898316927f8484265a68fd7cea22f85d6a338d7534f34ae4ff7d204bfa9faaec9e9e696ba392610fb0928b9216908d9061190a565b60405180910390a4505050600093845250505060208190526040902080546001600160a01b03199081168255600182018054909116905560020180546001600160c01b0319169055565b50600090565b6001546001600160a01b031681565b60008061101b836114dc565b6000858152602081905260409020600201549193506001600160601b0316915042600160801b90910467ffffffffffffffff161161106b5760405162461bcd60e51b81526004016102d39061172c565b6000838152602081905260409020600101546001600160a01b0316156110a35760405162461bcd60e51b81526004016102d390611825565b6000838152602081905260409020600201546001600160801b03163410156110dd5760405162461bcd60e51b81526004016102d39061177a565b600080836001600160a01b0316632a55205a84346040518363ffffffff1660e01b815260040161110e929190611929565b604080518083038186803b15801561112557600080fd5b505afa158015611139573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061115d919061155f565b91509150600061271034600160009054906101000a90046001600160a01b03166001600160a01b0316630e716d976040518163ffffffff1660e01b815260040160206040518083038186803b1580156111b557600080fd5b505afa1580156111c9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111ed9190611628565b6111f7919061196f565b611201919061194f565b9050600160009054906101000a90046001600160a01b03166001600160a01b031663c415b95c6040518163ffffffff1660e01b815260040160206040518083038186803b15801561125157600080fd5b505afa158015611265573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112899190611518565b6001600160a01b03166108fc829081150290604051600060405180830381858888f193505050501580156112c1573d6000803e3d6000fd5b506040516001600160a01b0384169083156108fc029084906000818181858888f193505050501580156112f8573d6000803e3d6000fd5b506000868152602081905260409020546001600160a01b03166108fc8361131f843461198e565b611329919061198e565b6040518115909202916000818181858888f19350505050158015611351573d6000803e3d6000fd5b506040516323b872dd60e01b81526001600160a01b038616906323b872dd90611382903090339089906004016116ac565b600060405180830381600087803b15801561139c57600080fd5b505af11580156113b0573d6000803e3d6000fd5b50506001805460008a81526020819052604090819020805493015490516346eb67ef60e01b81526001600160a01b0392831695506346eb67ef9450611404938c93908116923392911690349060040161186c565b600060405180830381600087803b15801561141e57600080fd5b505af1158015611432573d6000803e3d6000fd5b50505060008781526020819052604090819020805460019091015491513393506001600160a01b0391821692898316927f8484265a68fd7cea22f85d6a338d7534f34ae4ff7d204bfa9faaec9e9e696ba392611493928b921690349061190a565b60405180910390a45050506000928352505060208190526040902080546001600160a01b03199081168255600182018054909116905560020180546001600160c01b0319169055565b606081901c91565b60006115036bffffffffffffffffffffffff19606085901b1682611937565b90506107af6001600160601b03831682611937565b600060208284031215611529578081fd5b81516107af816119bb565b60008060408385031215611546578081fd5b8235611551816119bb565b946020939093013593505050565b60008060408385031215611571578182fd5b825161157c816119bb565b6020939093015192949293505050565b60008060008060008060c087890312156115a4578182fd5b86356115af816119bb565b95506020870135945060408701356115c6816119bb565b935060608701356115d6816119bb565b9598949750929560808101359460a0909101359350915050565b600060208284031215611601578081fd5b815180151581146107af578182fd5b600060208284031215611621578081fd5b5035919050565b600060208284031215611639578081fd5b5051919050565b600080600060608486031215611654578283fd5b833592506020840135611666816119bb565b929592945050506040919091013590565b60008060408385031215611689578182fd5b50508035926020909101359150565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03929092168252602082015260400190565b901515815260200190565b95151586526001600160a01b03948516602087015260408601939093526060850191909152821660808401521660a082015260c00190565b6020808252600c908201526b496e76616c69642073616c6560a01b604082015260600190565b6020808252600e908201526d5072696365206f766572666c6f7760901b604082015260600190565b602080825260149082015273125b9cdd59999a58da595b9d081c185e5b595b9d60621b604082015260600190565b602080825260169082015275139bdd081bdddb995c881bdc881a5d195b481cdbdb1960521b604082015260600190565b6020808252600e908201526d54696d6520756e646572666c6f7760901b604082015260600190565b6020808252600b908201526a2bb937b733902a37b5b2b760a91b604082015260600190565b6020808252600d908201526c15dc9bdb99c81c185e5b595b9d609a1b604082015260600190565b90815260200190565b9182526001600160a01b0316602082015260400190565b9485526001600160a01b03938416602086015291831660408501529091166060830152608082015260a00190565b9485526001600160a01b0393841660208601529190921660408401526001600160801b03909116606083015267ffffffffffffffff16608082015260a00190565b9485526001600160a01b0393841660208601529190921660408401526060830191909152608082015260a00190565b9283526001600160a01b03919091166020830152604082015260600190565b918252602082015260400190565b6000821982111561194a5761194a6119a5565b500190565b60008261196a57634e487b7160e01b81526012600452602481fd5b500490565b6000816000190483118215151615611989576119896119a5565b500290565b6000828210156119a0576119a06119a5565b500390565b634e487b7160e01b600052601160045260246000fd5b6001600160a01b038116811461079657600080fdfea26469706673582212209e514d7e996794be1549af1d6a27f9ef9a9db496dab71e801ecda656fe5cb78e64736f6c63430008010033`

// DeploySalePlain deploys a new Ethereum contract, binding an instance of SalePlain to it.
func DeploySalePlain(auth *bind.TransactOpts, backend bind.ContractBackend, master common.Address) (common.Address, *types.Transaction, *SalePlain, error) {
	parsed, err := abi.JSON(strings.NewReader(SalePlainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SalePlainBin), backend, master)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SalePlain{SalePlainCaller: SalePlainCaller{contract: contract}, SalePlainTransactor: SalePlainTransactor{contract: contract}, SalePlainFilterer: SalePlainFilterer{contract: contract}}, nil
}

// SalePlain is an auto generated Go binding around an Ethereum contract.
type SalePlain struct {
	SalePlainCaller     // Read-only binding to the contract
	SalePlainTransactor // Write-only binding to the contract
	SalePlainFilterer   // Log filterer for contract events
}

// SalePlainCaller is an auto generated read-only Go binding around an Ethereum contract.
type SalePlainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SalePlainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SalePlainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SalePlainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SalePlainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SalePlainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SalePlainSession struct {
	Contract     *SalePlain        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SalePlainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SalePlainCallerSession struct {
	Contract *SalePlainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SalePlainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SalePlainTransactorSession struct {
	Contract     *SalePlainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SalePlainRaw is an auto generated low-level Go binding around an Ethereum contract.
type SalePlainRaw struct {
	Contract *SalePlain // Generic contract binding to access the raw methods on
}

// SalePlainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SalePlainCallerRaw struct {
	Contract *SalePlainCaller // Generic read-only contract binding to access the raw methods on
}

// SalePlainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SalePlainTransactorRaw struct {
	Contract *SalePlainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSalePlain creates a new instance of SalePlain, bound to a specific deployed contract.
func NewSalePlain(address common.Address, backend bind.ContractBackend) (*SalePlain, error) {
	contract, err := bindSalePlain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SalePlain{SalePlainCaller: SalePlainCaller{contract: contract}, SalePlainTransactor: SalePlainTransactor{contract: contract}, SalePlainFilterer: SalePlainFilterer{contract: contract}}, nil
}

// NewSalePlainCaller creates a new read-only instance of SalePlain, bound to a specific deployed contract.
func NewSalePlainCaller(address common.Address, caller bind.ContractCaller) (*SalePlainCaller, error) {
	contract, err := bindSalePlain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SalePlainCaller{contract: contract}, nil
}

// NewSalePlainTransactor creates a new write-only instance of SalePlain, bound to a specific deployed contract.
func NewSalePlainTransactor(address common.Address, transactor bind.ContractTransactor) (*SalePlainTransactor, error) {
	contract, err := bindSalePlain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SalePlainTransactor{contract: contract}, nil
}

// NewSalePlainFilterer creates a new log filterer instance of SalePlain, bound to a specific deployed contract.
func NewSalePlainFilterer(address common.Address, filterer bind.ContractFilterer) (*SalePlainFilterer, error) {
	contract, err := bindSalePlain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SalePlainFilterer{contract: contract}, nil
}

// bindSalePlain binds a generic wrapper to an already deployed contract.
func bindSalePlain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SalePlainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SalePlain *SalePlainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SalePlain.Contract.SalePlainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SalePlain *SalePlainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SalePlain.Contract.SalePlainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SalePlain *SalePlainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SalePlain.Contract.SalePlainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SalePlain *SalePlainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SalePlain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SalePlain *SalePlainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SalePlain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SalePlain *SalePlainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SalePlain.Contract.contract.Transact(opts, method, params...)
}

// BidWithToken is a paid mutator transaction binding the contract method 0x8ff05a2a.
//
// Solidity: function bidWithToken(uint256 _saleId, uint256 _amount) returns()
func (_SalePlain *SalePlainTransactor) BidWithToken(opts *bind.TransactOpts, _saleId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _SalePlain.contract.Transact(opts, "bidWithToken", _saleId, _amount)
}

// BidWithToken is a paid mutator transaction binding the contract method 0x8ff05a2a.
//
// Solidity: function bidWithToken(uint256 _saleId, uint256 _amount) returns()
func (_SalePlain *SalePlainSession) BidWithToken(_saleId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _SalePlain.Contract.BidWithToken(&_SalePlain.TransactOpts, _saleId, _amount)
}

// BidWithToken is a paid mutator transaction binding the contract method 0x8ff05a2a.
//
// Solidity: function bidWithToken(uint256 _saleId, uint256 _amount) returns()
func (_SalePlain *SalePlainTransactorSession) BidWithToken(_saleId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _SalePlain.Contract.BidWithToken(&_SalePlain.TransactOpts, _saleId, _amount)
}

// BidWithValue is a paid mutator transaction binding the contract method 0xfaa21b11.
//
// Solidity: function bidWithValue(uint256 _saleId) returns()
func (_SalePlain *SalePlainTransactor) BidWithValue(opts *bind.TransactOpts, _saleId *big.Int) (*types.Transaction, error) {
	return _SalePlain.contract.Transact(opts, "bidWithValue", _saleId)
}

// BidWithValue is a paid mutator transaction binding the contract method 0xfaa21b11.
//
// Solidity: function bidWithValue(uint256 _saleId) returns()
func (_SalePlain *SalePlainSession) BidWithValue(_saleId *big.Int) (*types.Transaction, error) {
	return _SalePlain.Contract.BidWithValue(&_SalePlain.TransactOpts, _saleId)
}

// BidWithValue is a paid mutator transaction binding the contract method 0xfaa21b11.
//
// Solidity: function bidWithValue(uint256 _saleId) returns()
func (_SalePlain *SalePlainTransactorSession) BidWithValue(_saleId *big.Int) (*types.Transaction, error) {
	return _SalePlain.Contract.BidWithValue(&_SalePlain.TransactOpts, _saleId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 _saleId) returns()
func (_SalePlain *SalePlainTransactor) Cancel(opts *bind.TransactOpts, _saleId *big.Int) (*types.Transaction, error) {
	return _SalePlain.contract.Transact(opts, "cancel", _saleId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 _saleId) returns()
func (_SalePlain *SalePlainSession) Cancel(_saleId *big.Int) (*types.Transaction, error) {
	return _SalePlain.Contract.Cancel(&_SalePlain.TransactOpts, _saleId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 _saleId) returns()
func (_SalePlain *SalePlainTransactorSession) Cancel(_saleId *big.Int) (*types.Transaction, error) {
	return _SalePlain.Contract.Cancel(&_SalePlain.TransactOpts, _saleId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _saleId) returns()
func (_SalePlain *SalePlainTransactor) Claim(opts *bind.TransactOpts, _saleId *big.Int) (*types.Transaction, error) {
	return _SalePlain.contract.Transact(opts, "claim", _saleId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _saleId) returns()
func (_SalePlain *SalePlainSession) Claim(_saleId *big.Int) (*types.Transaction, error) {
	return _SalePlain.Contract.Claim(&_SalePlain.TransactOpts, _saleId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _saleId) returns()
func (_SalePlain *SalePlainTransactorSession) Claim(_saleId *big.Int) (*types.Transaction, error) {
	return _SalePlain.Contract.Claim(&_SalePlain.TransactOpts, _saleId)
}

// Claimable is a paid mutator transaction binding the contract method 0xd1d58b25.
//
// Solidity: function claimable(uint256 _saleId) returns(bool)
func (_SalePlain *SalePlainTransactor) Claimable(opts *bind.TransactOpts, _saleId *big.Int) (*types.Transaction, error) {
	return _SalePlain.contract.Transact(opts, "claimable", _saleId)
}

// Claimable is a paid mutator transaction binding the contract method 0xd1d58b25.
//
// Solidity: function claimable(uint256 _saleId) returns(bool)
func (_SalePlain *SalePlainSession) Claimable(_saleId *big.Int) (*types.Transaction, error) {
	return _SalePlain.Contract.Claimable(&_SalePlain.TransactOpts, _saleId)
}

// Claimable is a paid mutator transaction binding the contract method 0xd1d58b25.
//
// Solidity: function claimable(uint256 _saleId) returns(bool)
func (_SalePlain *SalePlainTransactorSession) Claimable(_saleId *big.Int) (*types.Transaction, error) {
	return _SalePlain.Contract.Claimable(&_SalePlain.TransactOpts, _saleId)
}

// Create is a paid mutator transaction binding the contract method 0x370b0081.
//
// Solidity: function create(address _contract, uint256 _tokenId, address seller, address _token, uint256 _price, uint256 _due) returns(uint256 record)
func (_SalePlain *SalePlainTransactor) Create(opts *bind.TransactOpts, _contract common.Address, _tokenId *big.Int, seller common.Address, _token common.Address, _price *big.Int, _due *big.Int) (*types.Transaction, error) {
	return _SalePlain.contract.Transact(opts, "create", _contract, _tokenId, seller, _token, _price, _due)
}

// Create is a paid mutator transaction binding the contract method 0x370b0081.
//
// Solidity: function create(address _contract, uint256 _tokenId, address seller, address _token, uint256 _price, uint256 _due) returns(uint256 record)
func (_SalePlain *SalePlainSession) Create(_contract common.Address, _tokenId *big.Int, seller common.Address, _token common.Address, _price *big.Int, _due *big.Int) (*types.Transaction, error) {
	return _SalePlain.Contract.Create(&_SalePlain.TransactOpts, _contract, _tokenId, seller, _token, _price, _due)
}

// Create is a paid mutator transaction binding the contract method 0x370b0081.
//
// Solidity: function create(address _contract, uint256 _tokenId, address seller, address _token, uint256 _price, uint256 _due) returns(uint256 record)
func (_SalePlain *SalePlainTransactorSession) Create(_contract common.Address, _tokenId *big.Int, seller common.Address, _token common.Address, _price *big.Int, _due *big.Int) (*types.Transaction, error) {
	return _SalePlain.Contract.Create(&_SalePlain.TransactOpts, _contract, _tokenId, seller, _token, _price, _due)
}

// GetSaleId is a paid mutator transaction binding the contract method 0x3969e1ae.
//
// Solidity: function getSaleId(address _contract, uint256 _tokenId) returns(uint256 _saleId)
func (_SalePlain *SalePlainTransactor) GetSaleId(opts *bind.TransactOpts, _contract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _SalePlain.contract.Transact(opts, "getSaleId", _contract, _tokenId)
}

// GetSaleId is a paid mutator transaction binding the contract method 0x3969e1ae.
//
// Solidity: function getSaleId(address _contract, uint256 _tokenId) returns(uint256 _saleId)
func (_SalePlain *SalePlainSession) GetSaleId(_contract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _SalePlain.Contract.GetSaleId(&_SalePlain.TransactOpts, _contract, _tokenId)
}

// GetSaleId is a paid mutator transaction binding the contract method 0x3969e1ae.
//
// Solidity: function getSaleId(address _contract, uint256 _tokenId) returns(uint256 _saleId)
func (_SalePlain *SalePlainTransactorSession) GetSaleId(_contract common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _SalePlain.Contract.GetSaleId(&_SalePlain.TransactOpts, _contract, _tokenId)
}

// GetSaleInfo is a paid mutator transaction binding the contract method 0x64b3b844.
//
// Solidity: function getSaleInfo(uint256 _saleId) returns(bool _isOpen, address _token, uint256 _price, uint256 _due, address _seller, address _bidder)
func (_SalePlain *SalePlainTransactor) GetSaleInfo(opts *bind.TransactOpts, _saleId *big.Int) (*types.Transaction, error) {
	return _SalePlain.contract.Transact(opts, "getSaleInfo", _saleId)
}

// GetSaleInfo is a paid mutator transaction binding the contract method 0x64b3b844.
//
// Solidity: function getSaleInfo(uint256 _saleId) returns(bool _isOpen, address _token, uint256 _price, uint256 _due, address _seller, address _bidder)
func (_SalePlain *SalePlainSession) GetSaleInfo(_saleId *big.Int) (*types.Transaction, error) {
	return _SalePlain.Contract.GetSaleInfo(&_SalePlain.TransactOpts, _saleId)
}

// GetSaleInfo is a paid mutator transaction binding the contract method 0x64b3b844.
//
// Solidity: function getSaleInfo(uint256 _saleId) returns(bool _isOpen, address _token, uint256 _price, uint256 _due, address _seller, address _bidder)
func (_SalePlain *SalePlainTransactorSession) GetSaleInfo(_saleId *big.Int) (*types.Transaction, error) {
	return _SalePlain.Contract.GetSaleInfo(&_SalePlain.TransactOpts, _saleId)
}

// MasterAddress is a paid mutator transaction binding the contract method 0xd365a08e.
//
// Solidity: function masterAddress() returns(address)
func (_SalePlain *SalePlainTransactor) MasterAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SalePlain.contract.Transact(opts, "masterAddress")
}

// MasterAddress is a paid mutator transaction binding the contract method 0xd365a08e.
//
// Solidity: function masterAddress() returns(address)
func (_SalePlain *SalePlainSession) MasterAddress() (*types.Transaction, error) {
	return _SalePlain.Contract.MasterAddress(&_SalePlain.TransactOpts)
}

// MasterAddress is a paid mutator transaction binding the contract method 0xd365a08e.
//
// Solidity: function masterAddress() returns(address)
func (_SalePlain *SalePlainTransactorSession) MasterAddress() (*types.Transaction, error) {
	return _SalePlain.Contract.MasterAddress(&_SalePlain.TransactOpts)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0x2ad1f1db.
//
// Solidity: function updateDuration(uint256 _saleId, uint256 _due) returns()
func (_SalePlain *SalePlainTransactor) UpdateDuration(opts *bind.TransactOpts, _saleId *big.Int, _due *big.Int) (*types.Transaction, error) {
	return _SalePlain.contract.Transact(opts, "updateDuration", _saleId, _due)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0x2ad1f1db.
//
// Solidity: function updateDuration(uint256 _saleId, uint256 _due) returns()
func (_SalePlain *SalePlainSession) UpdateDuration(_saleId *big.Int, _due *big.Int) (*types.Transaction, error) {
	return _SalePlain.Contract.UpdateDuration(&_SalePlain.TransactOpts, _saleId, _due)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0x2ad1f1db.
//
// Solidity: function updateDuration(uint256 _saleId, uint256 _due) returns()
func (_SalePlain *SalePlainTransactorSession) UpdateDuration(_saleId *big.Int, _due *big.Int) (*types.Transaction, error) {
	return _SalePlain.Contract.UpdateDuration(&_SalePlain.TransactOpts, _saleId, _due)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x0f0a1b9e.
//
// Solidity: function updatePrice(uint256 _saleId, address _token, uint256 _price) returns()
func (_SalePlain *SalePlainTransactor) UpdatePrice(opts *bind.TransactOpts, _saleId *big.Int, _token common.Address, _price *big.Int) (*types.Transaction, error) {
	return _SalePlain.contract.Transact(opts, "updatePrice", _saleId, _token, _price)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x0f0a1b9e.
//
// Solidity: function updatePrice(uint256 _saleId, address _token, uint256 _price) returns()
func (_SalePlain *SalePlainSession) UpdatePrice(_saleId *big.Int, _token common.Address, _price *big.Int) (*types.Transaction, error) {
	return _SalePlain.Contract.UpdatePrice(&_SalePlain.TransactOpts, _saleId, _token, _price)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x0f0a1b9e.
//
// Solidity: function updatePrice(uint256 _saleId, address _token, uint256 _price) returns()
func (_SalePlain *SalePlainTransactorSession) UpdatePrice(_saleId *big.Int, _token common.Address, _price *big.Int) (*types.Transaction, error) {
	return _SalePlain.Contract.UpdatePrice(&_SalePlain.TransactOpts, _saleId, _token, _price)
}

// SalePlainCancelSaleIterator is returned from FilterCancelSale and is used to iterate over the raw logs and unpacked data for CancelSale events raised by the SalePlain contract.
type SalePlainCancelSaleIterator struct {
	Event *SalePlainCancelSale // Event containing the contract specifics and raw log

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
func (it *SalePlainCancelSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SalePlainCancelSale)
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
		it.Event = new(SalePlainCancelSale)
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
func (it *SalePlainCancelSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SalePlainCancelSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SalePlainCancelSale represents a CancelSale event raised by the SalePlain contract.
type SalePlainCancelSale struct {
	Contract common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCancelSale is a free log retrieval operation binding the contract event 0x8b0a052bd72fec557d5735593db719bcb565d7afec8d76ee165e317e8cc98409.
//
// Solidity: event CancelSale(address indexed _contract, uint256 _tokenId)
func (_SalePlain *SalePlainFilterer) FilterCancelSale(opts *bind.FilterOpts, _contract []common.Address) (*SalePlainCancelSaleIterator, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	logs, sub, err := _SalePlain.contract.FilterLogs(opts, "CancelSale", _contractRule)
	if err != nil {
		return nil, err
	}
	return &SalePlainCancelSaleIterator{contract: _SalePlain.contract, event: "CancelSale", logs: logs, sub: sub}, nil
}

// WatchCancelSale is a free log subscription operation binding the contract event 0x8b0a052bd72fec557d5735593db719bcb565d7afec8d76ee165e317e8cc98409.
//
// Solidity: event CancelSale(address indexed _contract, uint256 _tokenId)
func (_SalePlain *SalePlainFilterer) WatchCancelSale(opts *bind.WatchOpts, sink chan<- *SalePlainCancelSale, _contract []common.Address) (event.Subscription, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	logs, sub, err := _SalePlain.contract.WatchLogs(opts, "CancelSale", _contractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SalePlainCancelSale)
				if err := _SalePlain.contract.UnpackLog(event, "CancelSale", log); err != nil {
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

// SalePlainConfirmSaleIterator is returned from FilterConfirmSale and is used to iterate over the raw logs and unpacked data for ConfirmSale events raised by the SalePlain contract.
type SalePlainConfirmSaleIterator struct {
	Event *SalePlainConfirmSale // Event containing the contract specifics and raw log

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
func (it *SalePlainConfirmSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SalePlainConfirmSale)
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
		it.Event = new(SalePlainConfirmSale)
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
func (it *SalePlainConfirmSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SalePlainConfirmSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SalePlainConfirmSale represents a ConfirmSale event raised by the SalePlain contract.
type SalePlainConfirmSale struct {
	Contract common.Address
	TokenId  *big.Int
	Solder   common.Address
	Buyer    common.Address
	Token    common.Address
	Price    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConfirmSale is a free log retrieval operation binding the contract event 0x8484265a68fd7cea22f85d6a338d7534f34ae4ff7d204bfa9faaec9e9e696ba3.
//
// Solidity: event ConfirmSale(address indexed _contract, uint256 _tokenId, address indexed solder, address indexed buyer, address token, uint256 _price)
func (_SalePlain *SalePlainFilterer) FilterConfirmSale(opts *bind.FilterOpts, _contract []common.Address, solder []common.Address, buyer []common.Address) (*SalePlainConfirmSaleIterator, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	var solderRule []interface{}
	for _, solderItem := range solder {
		solderRule = append(solderRule, solderItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _SalePlain.contract.FilterLogs(opts, "ConfirmSale", _contractRule, solderRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &SalePlainConfirmSaleIterator{contract: _SalePlain.contract, event: "ConfirmSale", logs: logs, sub: sub}, nil
}

// WatchConfirmSale is a free log subscription operation binding the contract event 0x8484265a68fd7cea22f85d6a338d7534f34ae4ff7d204bfa9faaec9e9e696ba3.
//
// Solidity: event ConfirmSale(address indexed _contract, uint256 _tokenId, address indexed solder, address indexed buyer, address token, uint256 _price)
func (_SalePlain *SalePlainFilterer) WatchConfirmSale(opts *bind.WatchOpts, sink chan<- *SalePlainConfirmSale, _contract []common.Address, solder []common.Address, buyer []common.Address) (event.Subscription, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	var solderRule []interface{}
	for _, solderItem := range solder {
		solderRule = append(solderRule, solderItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _SalePlain.contract.WatchLogs(opts, "ConfirmSale", _contractRule, solderRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SalePlainConfirmSale)
				if err := _SalePlain.contract.UnpackLog(event, "ConfirmSale", log); err != nil {
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

// SalePlainCreateSaleIterator is returned from FilterCreateSale and is used to iterate over the raw logs and unpacked data for CreateSale events raised by the SalePlain contract.
type SalePlainCreateSaleIterator struct {
	Event *SalePlainCreateSale // Event containing the contract specifics and raw log

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
func (it *SalePlainCreateSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SalePlainCreateSale)
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
		it.Event = new(SalePlainCreateSale)
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
func (it *SalePlainCreateSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SalePlainCreateSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SalePlainCreateSale represents a CreateSale event raised by the SalePlain contract.
type SalePlainCreateSale struct {
	Contract common.Address
	TokenId  *big.Int
	Seller   common.Address
	Token    common.Address
	Price    *big.Int
	Due      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCreateSale is a free log retrieval operation binding the contract event 0x008f298a385d495e86787887da324d1603491ea0ea157acffd6c5851d6ddfa2f.
//
// Solidity: event CreateSale(address indexed _contract, uint256 _tokenId, address seller, address token, uint256 _price, uint256 _due)
func (_SalePlain *SalePlainFilterer) FilterCreateSale(opts *bind.FilterOpts, _contract []common.Address) (*SalePlainCreateSaleIterator, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	logs, sub, err := _SalePlain.contract.FilterLogs(opts, "CreateSale", _contractRule)
	if err != nil {
		return nil, err
	}
	return &SalePlainCreateSaleIterator{contract: _SalePlain.contract, event: "CreateSale", logs: logs, sub: sub}, nil
}

// WatchCreateSale is a free log subscription operation binding the contract event 0x008f298a385d495e86787887da324d1603491ea0ea157acffd6c5851d6ddfa2f.
//
// Solidity: event CreateSale(address indexed _contract, uint256 _tokenId, address seller, address token, uint256 _price, uint256 _due)
func (_SalePlain *SalePlainFilterer) WatchCreateSale(opts *bind.WatchOpts, sink chan<- *SalePlainCreateSale, _contract []common.Address) (event.Subscription, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	logs, sub, err := _SalePlain.contract.WatchLogs(opts, "CreateSale", _contractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SalePlainCreateSale)
				if err := _SalePlain.contract.UnpackLog(event, "CreateSale", log); err != nil {
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

// SalePlainNewBidIterator is returned from FilterNewBid and is used to iterate over the raw logs and unpacked data for NewBid events raised by the SalePlain contract.
type SalePlainNewBidIterator struct {
	Event *SalePlainNewBid // Event containing the contract specifics and raw log

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
func (it *SalePlainNewBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SalePlainNewBid)
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
		it.Event = new(SalePlainNewBid)
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
func (it *SalePlainNewBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SalePlainNewBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SalePlainNewBid represents a NewBid event raised by the SalePlain contract.
type SalePlainNewBid struct {
	Contract common.Address
	TokenId  *big.Int
	Bidder   common.Address
	Token    common.Address
	Price    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewBid is a free log retrieval operation binding the contract event 0x3a62da8721790771856162caef0b31ff0c28626c1a7bcd82ed8a6a2a6b0009a2.
//
// Solidity: event NewBid(address indexed _contract, uint256 _tokenId, address indexed bidder, address token, uint256 _price)
func (_SalePlain *SalePlainFilterer) FilterNewBid(opts *bind.FilterOpts, _contract []common.Address, bidder []common.Address) (*SalePlainNewBidIterator, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _SalePlain.contract.FilterLogs(opts, "NewBid", _contractRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &SalePlainNewBidIterator{contract: _SalePlain.contract, event: "NewBid", logs: logs, sub: sub}, nil
}

// WatchNewBid is a free log subscription operation binding the contract event 0x3a62da8721790771856162caef0b31ff0c28626c1a7bcd82ed8a6a2a6b0009a2.
//
// Solidity: event NewBid(address indexed _contract, uint256 _tokenId, address indexed bidder, address token, uint256 _price)
func (_SalePlain *SalePlainFilterer) WatchNewBid(opts *bind.WatchOpts, sink chan<- *SalePlainNewBid, _contract []common.Address, bidder []common.Address) (event.Subscription, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _SalePlain.contract.WatchLogs(opts, "NewBid", _contractRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SalePlainNewBid)
				if err := _SalePlain.contract.UnpackLog(event, "NewBid", log); err != nil {
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

// SalePlainUpdateSaleIterator is returned from FilterUpdateSale and is used to iterate over the raw logs and unpacked data for UpdateSale events raised by the SalePlain contract.
type SalePlainUpdateSaleIterator struct {
	Event *SalePlainUpdateSale // Event containing the contract specifics and raw log

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
func (it *SalePlainUpdateSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SalePlainUpdateSale)
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
		it.Event = new(SalePlainUpdateSale)
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
func (it *SalePlainUpdateSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SalePlainUpdateSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SalePlainUpdateSale represents a UpdateSale event raised by the SalePlain contract.
type SalePlainUpdateSale struct {
	Contract common.Address
	TokenId  *big.Int
	Seller   common.Address
	Token    common.Address
	Price    *big.Int
	Due      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateSale is a free log retrieval operation binding the contract event 0xfe8937da95814c69018f7a05dcec87068d26502ab6dabf6516ea314dc1dca32b.
//
// Solidity: event UpdateSale(address indexed _contract, uint256 _tokenId, address seller, address token, uint256 _price, uint256 _due)
func (_SalePlain *SalePlainFilterer) FilterUpdateSale(opts *bind.FilterOpts, _contract []common.Address) (*SalePlainUpdateSaleIterator, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	logs, sub, err := _SalePlain.contract.FilterLogs(opts, "UpdateSale", _contractRule)
	if err != nil {
		return nil, err
	}
	return &SalePlainUpdateSaleIterator{contract: _SalePlain.contract, event: "UpdateSale", logs: logs, sub: sub}, nil
}

// WatchUpdateSale is a free log subscription operation binding the contract event 0xfe8937da95814c69018f7a05dcec87068d26502ab6dabf6516ea314dc1dca32b.
//
// Solidity: event UpdateSale(address indexed _contract, uint256 _tokenId, address seller, address token, uint256 _price, uint256 _due)
func (_SalePlain *SalePlainFilterer) WatchUpdateSale(opts *bind.WatchOpts, sink chan<- *SalePlainUpdateSale, _contract []common.Address) (event.Subscription, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	logs, sub, err := _SalePlain.contract.WatchLogs(opts, "UpdateSale", _contractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SalePlainUpdateSale)
				if err := _SalePlain.contract.UnpackLog(event, "UpdateSale", log); err != nil {
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
