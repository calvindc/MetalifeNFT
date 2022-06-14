// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package saleAuction

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

// SaleAuctionMetaData contains all meta data concerning the SaleAuction contract.
var SaleAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"master\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"CancelSale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"solder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"ConfirmSale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_due\",\"type\":\"uint256\"}],\"name\":\"CreateSale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"NewBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_due\",\"type\":\"uint256\"}],\"name\":\"UpdateSale\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"bidWithToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleId\",\"type\":\"uint256\"}],\"name\":\"bidWithValue\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleId\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleId\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleId\",\"type\":\"uint256\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_due\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"record\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getSaleId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleId\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getSaleInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleId\",\"type\":\"uint256\"}],\"name\":\"getSaleInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isOpen\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_due\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"masterAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_due\",\"type\":\"uint256\"}],\"name\":\"updateDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051611f60380380611f6083398101604081905261002f91610054565b600180546001600160a01b0319166001600160a01b0392909216919091179055610082565b600060208284031215610065578081fd5b81516001600160a01b038116811461007b578182fd5b9392505050565b611ecf806100916000396000f3fe6080604052600436106100a75760003560e01c806340e58ee51161006457806340e58ee51461019657806364b3b844146101b65780638ff05a2a146101d6578063d1d58b25146101f6578063d365a08e14610223578063faa21b1114610245576100a7565b8063087db5b8146100ac5780630f0a1b9e146100e75780632ad1f1db14610109578063370b008114610129578063379607f5146101565780633969e1ae14610176575b600080fd5b3480156100b857600080fd5b506100cc6100c7366004611990565b610258565b6040516100de96959493929190611b72565b60405180910390f35b3480156100f357600080fd5b50610107610102366004611a9c565b610283565b005b34801561011557600080fd5b50610107610124366004611ad3565b610465565b34801561013557600080fd5b506101496101443660046119e8565b6105b4565b6040516100de9190611d12565b34801561016257600080fd5b50610107610171366004611a6c565b610819565b34801561018257600080fd5b50610149610191366004611990565b610f54565b3480156101a257600080fd5b506101076101b1366004611a6c565b610f67565b3480156101c257600080fd5b506100cc6101d1366004611a6c565b611141565b3480156101e257600080fd5b506101076101f1366004611ad3565b61119a565b34801561020257600080fd5b50610216610211366004611a6c565b6115ac565b6040516100de9190611b67565b34801561022f57600080fd5b506102386115fe565b6040516100de9190611af4565b610107610253366004611a6c565b61160d565b60008060008060008061026e6101d18989610f54565b949d939c50919a509850965090945092505050565b60008061028f85611938565b6000878152602081905260409020549193506001600160601b031691506001600160a01b031633146102dc5760405162461bcd60e51b81526004016102d390611c48565b60405180910390fd5b600154604051630e70619560e41b81526001600160a01b039091169063e70619509061030c908790600401611af4565b60206040518083038186803b15801561032457600080fd5b505afa158015610338573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061035c9190611a4c565b6103785760405162461bcd60e51b81526004016102d390611cc6565b6000858152602081905260409020600301546001600160a01b0316156103b05760405162461bcd60e51b81526004016102d390611bd0565b600085815260208190526040908190206001810180546001600160a01b0319166001600160a01b038881169190911791829055600290920180546001600160801b0319166001600160801b0388811691909117918290559351868416947ffe8937da95814c69018f7a05dcec87068d26502ab6dabf6516ea314dc1dca32b9461045694889433949190921692821691600160801b900467ffffffffffffffff1690611d60565b60405180910390a25050505050565b60008061047184611938565b6000868152602081905260409020549193506001600160601b031691506001600160a01b031633146104b55760405162461bcd60e51b81526004016102d390611c48565b4283116104d45760405162461bcd60e51b81526004016102d390611c78565b6000848152602081905260409020600301546001600160a01b03161561050c5760405162461bcd60e51b81526004016102d390611bd0565b6000848152602081905260409081902060028101805467ffffffffffffffff60801b1916600160801b67ffffffffffffffff8881168202929092179283905560019093015493516001600160a01b03878116957ffe8937da95814c69018f7a05dcec87068d26502ab6dabf6516ea314dc1dca32b956105a6958995339594909316936001600160801b038216939290910490911690611d60565b60405180910390a250505050565b6040516331a9108f60e11b815260009030906001600160a01b03891690636352211e906105e5908a90600401611d12565b60206040518083038186803b1580156105fd57600080fd5b505afa158015610611573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106359190611974565b6001600160a01b03161461064857600080fd5b4282116106675760405162461bcd60e51b81526004016102d390611c78565b826001600160801b0316831461068f5760405162461bcd60e51b81526004016102d390611bf2565b6106998787611940565b6040805160a0810182526001600160a01b03888116825287811660208084019182526001600160801b0389811685870190815267ffffffffffffffff8a8116606088019081526000608089018181528b825295819052899020975188546001600160a01b031990811691891691909117895595516001808a0180548916928a1692909217909155925160028901805492516001600160801b0319909316919095161767ffffffffffffffff60801b1916600160801b919092160217909155905160039094018054909216938316939093179055905491516308d3fed360e01b815292935016906308d3fed3906107959084908990600401611d1b565b600060405180830381600087803b1580156107af57600080fd5b505af11580156107c3573d6000803e3d6000fd5b50505050866001600160a01b03167e8f298a385d495e86787887da324d1603491ea0ea157acffd6c5851d6ddfa2f8787878787604051610807959493929190611da1565b60405180910390a29695505050505050565b60008061082583611938565b6001600160601b03169150915061083b836115ac565b6108575760405162461bcd60e51b81526004016102d390611ca0565b60008381526020819052604080822060020154905163152a902d60e11b81526001600160801b03909116919081906001600160a01b03861690632a55205a906108a69087908790600401611def565b604080518083038186803b1580156108bd57600080fd5b505afa1580156108d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108f591906119bb565b91509150600061271084600160009054906101000a90046001600160a01b03166001600160a01b0316630e716d976040518163ffffffff1660e01b815260040160206040518083038186803b15801561094d57600080fd5b505afa158015610961573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109859190611a84565b61098f9190611e35565b6109999190611e15565b6000888152602081905260409020600101549091506001600160a01b0316610b0f57600160009054906101000a90046001600160a01b03166001600160a01b031663c415b95c6040518163ffffffff1660e01b815260040160206040518083038186803b158015610a0957600080fd5b505afa158015610a1d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a419190611974565b6001600160a01b03166108fc829081150290604051600060405180830381858888f19350505050158015610a79573d6000803e3d6000fd5b506040516001600160a01b0384169083156108fc029084906000818181858888f19350505050158015610ab0573d6000803e3d6000fd5b506000878152602081905260409020546001600160a01b03166108fc83610ad78488611e54565b610ae19190611e54565b6040518115909202916000818181858888f19350505050158015610b09573d6000803e3d6000fd5b50610d59565b60008781526020818152604091829020600190810154905483516331056e5760e21b815293516001600160a01b039283169463a9059cbb94939092169263c415b95c9260048082019391829003018186803b158015610b6d57600080fd5b505afa158015610b81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ba59190611974565b836040518363ffffffff1660e01b8152600401610bc3929190611b4e565b602060405180830381600087803b158015610bdd57600080fd5b505af1158015610bf1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c159190611a4c565b506000878152602081905260409081902060010154905163a9059cbb60e01b81526001600160a01b039091169063a9059cbb90610c589086908690600401611b4e565b602060405180830381600087803b158015610c7257600080fd5b505af1158015610c86573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610caa9190611a4c565b506000878152602081905260409020600181015490546001600160a01b039182169163a9059cbb911684610cde8589611e54565b610ce89190611e54565b6040518363ffffffff1660e01b8152600401610d05929190611b4e565b602060405180830381600087803b158015610d1f57600080fd5b505af1158015610d33573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d579190611a4c565b505b6040516323b872dd60e01b81526001600160a01b038716906323b872dd90610d8990309033908a90600401611b08565b600060405180830381600087803b158015610da357600080fd5b505af1158015610db7573d6000803e3d6000fd5b505060015460405163337f793960e21b81526001600160a01b03909116925063cdfde4e49150610ded908a903390600401611d1b565b600060405180830381600087803b158015610e0757600080fd5b505af1158015610e1b573d6000803e3d6000fd5b50506001805460008b81526020819052604090819020805493015490516346eb67ef60e01b81526001600160a01b0392831695506346eb67ef9450610e6f938d939081169233929116908b90600401611d32565b600060405180830381600087803b158015610e8957600080fd5b505af1158015610e9d573d6000803e3d6000fd5b50505060008881526020819052604090819020805460019091015491513393506001600160a01b03918216928a8316927f8484265a68fd7cea22f85d6a338d7534f34ae4ff7d204bfa9faaec9e9e696ba392610efe928c9216908b90611dd0565b60405180910390a4505050600093845250505060208190526040902080546001600160a01b03199081168255600182018054821690556002820180546001600160c01b0319169055600390910180549091169055565b6000610f608383611940565b9392505050565b600080610f7383611938565b6000858152602081905260409020549193506001600160601b031691506001600160a01b03163314610fb75760405162461bcd60e51b81526004016102d390611c48565b6000838152602081905260409020600301546001600160a01b031615610fef5760405162461bcd60e51b81526004016102d390611bd0565b6000838152602081905260409081902080546001600160a01b03199081168255600182018054821690556002820180546001600160c01b0319169055600390910180549091169055516323b872dd60e01b81526001600160a01b038316906323b872dd9061106590309033908690600401611b08565b600060405180830381600087803b15801561107f57600080fd5b505af1158015611093573d6000803e3d6000fd5b5050600154604051632084212160e11b81526001600160a01b039091169250634108424291506110c99086903390600401611d1b565b600060405180830381600087803b1580156110e357600080fd5b505af11580156110f7573d6000803e3d6000fd5b50505050816001600160a01b03167f8b0a052bd72fec557d5735593db719bcb565d7afec8d76ee165e317e8cc98409826040516111349190611d12565b60405180910390a2505050565b60009081526020819052604090206001810154600282015482546003909301544267ffffffffffffffff600160801b840416908111956001600160a01b03948516956001600160801b0390941694919382169290911690565b6000806111a684611938565b6000868152602081905260409020600201549193506001600160601b0316915042600160801b90910467ffffffffffffffff16116111f65760405162461bcd60e51b81526004016102d390611baa565b6000848152602081905260409020600101546001600160a01b031661122d5760405162461bcd60e51b81526004016102d390611ceb565b6000848152602081905260409020600301546001600160a01b03161561128b576000848152602081905260409020600201546001600160801b031683116112865760405162461bcd60e51b81526004016102d390611c1a565b6112c5565b6000848152602081905260409020600201546001600160801b03168310156112c55760405162461bcd60e51b81526004016102d390611c1a565b600084815260208190526040908190206001015490516323b872dd60e01b81526001600160a01b03909116906323b872dd9061130990339030908890600401611b08565b602060405180830381600087803b15801561132357600080fd5b505af1158015611337573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061135b9190611a4c565b506000848152602081905260409020600301546001600160a01b0316156114a1576000848152602081905260409081902060018101546003820154600290920154925163a9059cbb60e01b81526001600160a01b039182169363a9059cbb936113d49316916001600160801b0390911690600401611b2c565b602060405180830381600087803b1580156113ee57600080fd5b505af1158015611402573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114269190611a4c565b506001546000858152602081905260409081902060030154905163337f793960e21b81526001600160a01b039283169263cdfde4e49261146e92899290911690600401611d1b565b600060405180830381600087803b15801561148857600080fd5b505af115801561149c573d6000803e3d6000fd5b505050505b60008481526020819052604090819020600381018054336001600160a01b03199091168117909155600290910180546001600160801b0319166001600160801b03871617905560015491516317f7559b60e11b81526001600160a01b0390921691632feeab369161151791889190600401611d1b565b600060405180830381600087803b15801561153157600080fd5b505af1158015611545573d6000803e3d6000fd5b505050600085815260208190526040908190206001015490513392506001600160a01b03858116927f3a62da8721790771856162caef0b31ff0c28626c1a7bcd82ed8a6a2a6b0009a29261159e92879216908990611dd0565b60405180910390a350505050565b60008181526020819052604081206002015442600160801b90910467ffffffffffffffff16118015906115f857506000828152602081905260409020600301546001600160a01b031633145b92915050565b6001546001600160a01b031681565b60008061161983611938565b6000858152602081905260409020600201549193506001600160601b0316915042600160801b90910467ffffffffffffffff16116116695760405162461bcd60e51b81526004016102d390611baa565b6000838152602081905260409020600101546001600160a01b0316156116a15760405162461bcd60e51b81526004016102d390611ceb565b6000838152602081905260409020600301546001600160a01b0316156116ff576000838152602081905260409020600201546001600160801b031634116116fa5760405162461bcd60e51b81526004016102d390611c1a565b611739565b6000838152602081905260409020600201546001600160801b03163410156117395760405162461bcd60e51b81526004016102d390611c1a565b6000838152602081905260409020600301546001600160a01b03161561182e57600083815260208190526040808220600381015460029091015491516001600160a01b03909116926001600160801b0390921680156108fc0292909190818181858888f193505050501580156117b3573d6000803e3d6000fd5b506001546000848152602081905260409081902060030154905163337f793960e21b81526001600160a01b039283169263cdfde4e4926117fb92889290911690600401611d1b565b600060405180830381600087803b15801561181557600080fd5b505af1158015611829573d6000803e3d6000fd5b505050505b60008381526020819052604090819020600381018054336001600160a01b03199091168117909155600290910180546001600160801b031916346001600160801b031617905560015491516317f7559b60e11b81526001600160a01b0390921691632feeab36916118a491879190600401611d1b565b600060405180830381600087803b1580156118be57600080fd5b505af11580156118d2573d6000803e3d6000fd5b505050600084815260208190526040908190206001015490513392506001600160a01b03858116927f3a62da8721790771856162caef0b31ff0c28626c1a7bcd82ed8a6a2a6b0009a29261192b92879216903490611dd0565b60405180910390a3505050565b606081901c91565b600061195f6bffffffffffffffffffffffff19606085901b1682611dfd565b9050610f606001600160601b03831682611dfd565b600060208284031215611985578081fd5b8151610f6081611e81565b600080604083850312156119a2578081fd5b82356119ad81611e81565b946020939093013593505050565b600080604083850312156119cd578182fd5b82516119d881611e81565b6020939093015192949293505050565b60008060008060008060c08789031215611a00578182fd5b8635611a0b81611e81565b9550602087013594506040870135611a2281611e81565b93506060870135611a3281611e81565b9598949750929560808101359460a0909101359350915050565b600060208284031215611a5d578081fd5b81518015158114610f60578182fd5b600060208284031215611a7d578081fd5b5035919050565b600060208284031215611a95578081fd5b5051919050565b600080600060608486031215611ab0578283fd5b833592506020840135611ac281611e81565b929592945050506040919091013590565b60008060408385031215611ae5578182fd5b50508035926020909101359150565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b039290921682526001600160801b0316602082015260400190565b6001600160a01b03929092168252602082015260400190565b901515815260200190565b95151586526001600160a01b03948516602087015260408601939093526060850191909152821660808401521660a082015260c00190565b6020808252600c908201526b496e76616c69642073616c6560a01b604082015260600190565b602080825260089082015267486173206269642160c01b604082015260600190565b6020808252600e908201526d5072696365206f766572666c6f7760901b604082015260600190565b602080825260149082015273125b9cdd59999a58da595b9d081c185e5b595b9d60621b604082015260600190565b602080825260169082015275139bdd081bdddb995c881bdc881a5d195b481cdbdb1960521b604082015260600190565b6020808252600e908201526d54696d6520756e646572666c6f7760901b604082015260600190565b6020808252600c908201526b43616e6e6f7420636c61696d60a01b604082015260600190565b6020808252600b908201526a2bb937b733902a37b5b2b760a91b604082015260600190565b6020808252600d908201526c15dc9bdb99c81c185e5b595b9d609a1b604082015260600190565b90815260200190565b9182526001600160a01b0316602082015260400190565b9485526001600160a01b03938416602086015291831660408501529091166060830152608082015260a00190565b9485526001600160a01b0393841660208601529190921660408401526001600160801b03909116606083015267ffffffffffffffff16608082015260a00190565b9485526001600160a01b0393841660208601529190921660408401526060830191909152608082015260a00190565b9283526001600160a01b03919091166020830152604082015260600190565b918252602082015260400190565b60008219821115611e1057611e10611e6b565b500190565b600082611e3057634e487b7160e01b81526012600452602481fd5b500490565b6000816000190483118215151615611e4f57611e4f611e6b565b500290565b600082821015611e6657611e66611e6b565b500390565b634e487b7160e01b600052601160045260246000fd5b6001600160a01b0381168114611e9657600080fd5b5056fea2646970667358221220282f43b85e2832de33298aab9d196b10cd561aba51590ba87379aa7fd7ca9b3464736f6c63430008010033",
}

// SaleAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use SaleAuctionMetaData.ABI instead.
var SaleAuctionABI = SaleAuctionMetaData.ABI

// SaleAuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SaleAuctionMetaData.Bin instead.
var SaleAuctionBin = SaleAuctionMetaData.Bin

// DeploySaleAuction deploys a new Ethereum contract, binding an instance of SaleAuction to it.
func DeploySaleAuction(auth *bind.TransactOpts, backend bind.ContractBackend, master common.Address) (common.Address, *types.Transaction, *SaleAuction, error) {
	parsed, err := SaleAuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SaleAuctionBin), backend, master)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SaleAuction{SaleAuctionCaller: SaleAuctionCaller{contract: contract}, SaleAuctionTransactor: SaleAuctionTransactor{contract: contract}, SaleAuctionFilterer: SaleAuctionFilterer{contract: contract}}, nil
}

// SaleAuction is an auto generated Go binding around an Ethereum contract.
type SaleAuction struct {
	SaleAuctionCaller     // Read-only binding to the contract
	SaleAuctionTransactor // Write-only binding to the contract
	SaleAuctionFilterer   // Log filterer for contract events
}

// SaleAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type SaleAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SaleAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SaleAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SaleAuctionSession struct {
	Contract     *SaleAuction      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SaleAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SaleAuctionCallerSession struct {
	Contract *SaleAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SaleAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SaleAuctionTransactorSession struct {
	Contract     *SaleAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SaleAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type SaleAuctionRaw struct {
	Contract *SaleAuction // Generic contract binding to access the raw methods on
}

// SaleAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SaleAuctionCallerRaw struct {
	Contract *SaleAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// SaleAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SaleAuctionTransactorRaw struct {
	Contract *SaleAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSaleAuction creates a new instance of SaleAuction, bound to a specific deployed contract.
func NewSaleAuction(address common.Address, backend bind.ContractBackend) (*SaleAuction, error) {
	contract, err := bindSaleAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SaleAuction{SaleAuctionCaller: SaleAuctionCaller{contract: contract}, SaleAuctionTransactor: SaleAuctionTransactor{contract: contract}, SaleAuctionFilterer: SaleAuctionFilterer{contract: contract}}, nil
}

// NewSaleAuctionCaller creates a new read-only instance of SaleAuction, bound to a specific deployed contract.
func NewSaleAuctionCaller(address common.Address, caller bind.ContractCaller) (*SaleAuctionCaller, error) {
	contract, err := bindSaleAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SaleAuctionCaller{contract: contract}, nil
}

// NewSaleAuctionTransactor creates a new write-only instance of SaleAuction, bound to a specific deployed contract.
func NewSaleAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*SaleAuctionTransactor, error) {
	contract, err := bindSaleAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SaleAuctionTransactor{contract: contract}, nil
}

// NewSaleAuctionFilterer creates a new log filterer instance of SaleAuction, bound to a specific deployed contract.
func NewSaleAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*SaleAuctionFilterer, error) {
	contract, err := bindSaleAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SaleAuctionFilterer{contract: contract}, nil
}

// bindSaleAuction binds a generic wrapper to an already deployed contract.
func bindSaleAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SaleAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SaleAuction *SaleAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SaleAuction.Contract.SaleAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SaleAuction *SaleAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaleAuction.Contract.SaleAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SaleAuction *SaleAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SaleAuction.Contract.SaleAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SaleAuction *SaleAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SaleAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SaleAuction *SaleAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaleAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SaleAuction *SaleAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SaleAuction.Contract.contract.Transact(opts, method, params...)
}

// Claimable is a free data retrieval call binding the contract method 0xd1d58b25.
//
// Solidity: function claimable(uint256 _saleId) view returns(bool)
func (_SaleAuction *SaleAuctionCaller) Claimable(opts *bind.CallOpts, _saleId *big.Int) (bool, error) {
	var out []interface{}
	err := _SaleAuction.contract.Call(opts, &out, "claimable", _saleId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0xd1d58b25.
//
// Solidity: function claimable(uint256 _saleId) view returns(bool)
func (_SaleAuction *SaleAuctionSession) Claimable(_saleId *big.Int) (bool, error) {
	return _SaleAuction.Contract.Claimable(&_SaleAuction.CallOpts, _saleId)
}

// Claimable is a free data retrieval call binding the contract method 0xd1d58b25.
//
// Solidity: function claimable(uint256 _saleId) view returns(bool)
func (_SaleAuction *SaleAuctionCallerSession) Claimable(_saleId *big.Int) (bool, error) {
	return _SaleAuction.Contract.Claimable(&_SaleAuction.CallOpts, _saleId)
}

// GetSaleId is a free data retrieval call binding the contract method 0x3969e1ae.
//
// Solidity: function getSaleId(address _contract, uint256 _tokenId) pure returns(uint256 _saleId)
func (_SaleAuction *SaleAuctionCaller) GetSaleId(opts *bind.CallOpts, _contract common.Address, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SaleAuction.contract.Call(opts, &out, "getSaleId", _contract, _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSaleId is a free data retrieval call binding the contract method 0x3969e1ae.
//
// Solidity: function getSaleId(address _contract, uint256 _tokenId) pure returns(uint256 _saleId)
func (_SaleAuction *SaleAuctionSession) GetSaleId(_contract common.Address, _tokenId *big.Int) (*big.Int, error) {
	return _SaleAuction.Contract.GetSaleId(&_SaleAuction.CallOpts, _contract, _tokenId)
}

// GetSaleId is a free data retrieval call binding the contract method 0x3969e1ae.
//
// Solidity: function getSaleId(address _contract, uint256 _tokenId) pure returns(uint256 _saleId)
func (_SaleAuction *SaleAuctionCallerSession) GetSaleId(_contract common.Address, _tokenId *big.Int) (*big.Int, error) {
	return _SaleAuction.Contract.GetSaleId(&_SaleAuction.CallOpts, _contract, _tokenId)
}

// GetSaleInfo is a free data retrieval call binding the contract method 0x087db5b8.
//
// Solidity: function getSaleInfo(address _contract, uint256 _tokenId) view returns(bool, address, uint256, uint256, address, address)
func (_SaleAuction *SaleAuctionCaller) GetSaleInfo(opts *bind.CallOpts, _contract common.Address, _tokenId *big.Int) (bool, common.Address, *big.Int, *big.Int, common.Address, common.Address, error) {
	var out []interface{}
	err := _SaleAuction.contract.Call(opts, &out, "getSaleInfo", _contract, _tokenId)

	if err != nil {
		return *new(bool), *new(common.Address), *new(*big.Int), *new(*big.Int), *new(common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	out5 := *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return out0, out1, out2, out3, out4, out5, err

}

// GetSaleInfo is a free data retrieval call binding the contract method 0x087db5b8.
//
// Solidity: function getSaleInfo(address _contract, uint256 _tokenId) view returns(bool, address, uint256, uint256, address, address)
func (_SaleAuction *SaleAuctionSession) GetSaleInfo(_contract common.Address, _tokenId *big.Int) (bool, common.Address, *big.Int, *big.Int, common.Address, common.Address, error) {
	return _SaleAuction.Contract.GetSaleInfo(&_SaleAuction.CallOpts, _contract, _tokenId)
}

// GetSaleInfo is a free data retrieval call binding the contract method 0x087db5b8.
//
// Solidity: function getSaleInfo(address _contract, uint256 _tokenId) view returns(bool, address, uint256, uint256, address, address)
func (_SaleAuction *SaleAuctionCallerSession) GetSaleInfo(_contract common.Address, _tokenId *big.Int) (bool, common.Address, *big.Int, *big.Int, common.Address, common.Address, error) {
	return _SaleAuction.Contract.GetSaleInfo(&_SaleAuction.CallOpts, _contract, _tokenId)
}

// GetSaleInfo0 is a free data retrieval call binding the contract method 0x64b3b844.
//
// Solidity: function getSaleInfo(uint256 _saleId) view returns(bool _isOpen, address _token, uint256 _price, uint256 _due, address _seller, address _bidder)
func (_SaleAuction *SaleAuctionCaller) GetSaleInfo0(opts *bind.CallOpts, _saleId *big.Int) (struct {
	IsOpen bool
	Token  common.Address
	Price  *big.Int
	Due    *big.Int
	Seller common.Address
	Bidder common.Address
}, error) {
	var out []interface{}
	err := _SaleAuction.contract.Call(opts, &out, "getSaleInfo0", _saleId)

	outstruct := new(struct {
		IsOpen bool
		Token  common.Address
		Price  *big.Int
		Due    *big.Int
		Seller common.Address
		Bidder common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsOpen = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Token = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Due = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Bidder = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetSaleInfo0 is a free data retrieval call binding the contract method 0x64b3b844.
//
// Solidity: function getSaleInfo(uint256 _saleId) view returns(bool _isOpen, address _token, uint256 _price, uint256 _due, address _seller, address _bidder)
func (_SaleAuction *SaleAuctionSession) GetSaleInfo0(_saleId *big.Int) (struct {
	IsOpen bool
	Token  common.Address
	Price  *big.Int
	Due    *big.Int
	Seller common.Address
	Bidder common.Address
}, error) {
	return _SaleAuction.Contract.GetSaleInfo0(&_SaleAuction.CallOpts, _saleId)
}

// GetSaleInfo0 is a free data retrieval call binding the contract method 0x64b3b844.
//
// Solidity: function getSaleInfo(uint256 _saleId) view returns(bool _isOpen, address _token, uint256 _price, uint256 _due, address _seller, address _bidder)
func (_SaleAuction *SaleAuctionCallerSession) GetSaleInfo0(_saleId *big.Int) (struct {
	IsOpen bool
	Token  common.Address
	Price  *big.Int
	Due    *big.Int
	Seller common.Address
	Bidder common.Address
}, error) {
	return _SaleAuction.Contract.GetSaleInfo0(&_SaleAuction.CallOpts, _saleId)
}

// MasterAddress is a free data retrieval call binding the contract method 0xd365a08e.
//
// Solidity: function masterAddress() view returns(address)
func (_SaleAuction *SaleAuctionCaller) MasterAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SaleAuction.contract.Call(opts, &out, "masterAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterAddress is a free data retrieval call binding the contract method 0xd365a08e.
//
// Solidity: function masterAddress() view returns(address)
func (_SaleAuction *SaleAuctionSession) MasterAddress() (common.Address, error) {
	return _SaleAuction.Contract.MasterAddress(&_SaleAuction.CallOpts)
}

// MasterAddress is a free data retrieval call binding the contract method 0xd365a08e.
//
// Solidity: function masterAddress() view returns(address)
func (_SaleAuction *SaleAuctionCallerSession) MasterAddress() (common.Address, error) {
	return _SaleAuction.Contract.MasterAddress(&_SaleAuction.CallOpts)
}

// BidWithToken is a paid mutator transaction binding the contract method 0x8ff05a2a.
//
// Solidity: function bidWithToken(uint256 _saleId, uint256 _amount) returns()
func (_SaleAuction *SaleAuctionTransactor) BidWithToken(opts *bind.TransactOpts, _saleId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _SaleAuction.contract.Transact(opts, "bidWithToken", _saleId, _amount)
}

// BidWithToken is a paid mutator transaction binding the contract method 0x8ff05a2a.
//
// Solidity: function bidWithToken(uint256 _saleId, uint256 _amount) returns()
func (_SaleAuction *SaleAuctionSession) BidWithToken(_saleId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _SaleAuction.Contract.BidWithToken(&_SaleAuction.TransactOpts, _saleId, _amount)
}

// BidWithToken is a paid mutator transaction binding the contract method 0x8ff05a2a.
//
// Solidity: function bidWithToken(uint256 _saleId, uint256 _amount) returns()
func (_SaleAuction *SaleAuctionTransactorSession) BidWithToken(_saleId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _SaleAuction.Contract.BidWithToken(&_SaleAuction.TransactOpts, _saleId, _amount)
}

// BidWithValue is a paid mutator transaction binding the contract method 0xfaa21b11.
//
// Solidity: function bidWithValue(uint256 _saleId) payable returns()
func (_SaleAuction *SaleAuctionTransactor) BidWithValue(opts *bind.TransactOpts, _saleId *big.Int) (*types.Transaction, error) {
	return _SaleAuction.contract.Transact(opts, "bidWithValue", _saleId)
}

// BidWithValue is a paid mutator transaction binding the contract method 0xfaa21b11.
//
// Solidity: function bidWithValue(uint256 _saleId) payable returns()
func (_SaleAuction *SaleAuctionSession) BidWithValue(_saleId *big.Int) (*types.Transaction, error) {
	return _SaleAuction.Contract.BidWithValue(&_SaleAuction.TransactOpts, _saleId)
}

// BidWithValue is a paid mutator transaction binding the contract method 0xfaa21b11.
//
// Solidity: function bidWithValue(uint256 _saleId) payable returns()
func (_SaleAuction *SaleAuctionTransactorSession) BidWithValue(_saleId *big.Int) (*types.Transaction, error) {
	return _SaleAuction.Contract.BidWithValue(&_SaleAuction.TransactOpts, _saleId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 _saleId) returns()
func (_SaleAuction *SaleAuctionTransactor) Cancel(opts *bind.TransactOpts, _saleId *big.Int) (*types.Transaction, error) {
	return _SaleAuction.contract.Transact(opts, "cancel", _saleId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 _saleId) returns()
func (_SaleAuction *SaleAuctionSession) Cancel(_saleId *big.Int) (*types.Transaction, error) {
	return _SaleAuction.Contract.Cancel(&_SaleAuction.TransactOpts, _saleId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 _saleId) returns()
func (_SaleAuction *SaleAuctionTransactorSession) Cancel(_saleId *big.Int) (*types.Transaction, error) {
	return _SaleAuction.Contract.Cancel(&_SaleAuction.TransactOpts, _saleId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _saleId) returns()
func (_SaleAuction *SaleAuctionTransactor) Claim(opts *bind.TransactOpts, _saleId *big.Int) (*types.Transaction, error) {
	return _SaleAuction.contract.Transact(opts, "claim", _saleId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _saleId) returns()
func (_SaleAuction *SaleAuctionSession) Claim(_saleId *big.Int) (*types.Transaction, error) {
	return _SaleAuction.Contract.Claim(&_SaleAuction.TransactOpts, _saleId)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _saleId) returns()
func (_SaleAuction *SaleAuctionTransactorSession) Claim(_saleId *big.Int) (*types.Transaction, error) {
	return _SaleAuction.Contract.Claim(&_SaleAuction.TransactOpts, _saleId)
}

// Create is a paid mutator transaction binding the contract method 0x370b0081.
//
// Solidity: function create(address _contract, uint256 _tokenId, address seller, address _token, uint256 _price, uint256 _due) returns(uint256 record)
func (_SaleAuction *SaleAuctionTransactor) Create(opts *bind.TransactOpts, _contract common.Address, _tokenId *big.Int, seller common.Address, _token common.Address, _price *big.Int, _due *big.Int) (*types.Transaction, error) {
	return _SaleAuction.contract.Transact(opts, "create", _contract, _tokenId, seller, _token, _price, _due)
}

// Create is a paid mutator transaction binding the contract method 0x370b0081.
//
// Solidity: function create(address _contract, uint256 _tokenId, address seller, address _token, uint256 _price, uint256 _due) returns(uint256 record)
func (_SaleAuction *SaleAuctionSession) Create(_contract common.Address, _tokenId *big.Int, seller common.Address, _token common.Address, _price *big.Int, _due *big.Int) (*types.Transaction, error) {
	return _SaleAuction.Contract.Create(&_SaleAuction.TransactOpts, _contract, _tokenId, seller, _token, _price, _due)
}

// Create is a paid mutator transaction binding the contract method 0x370b0081.
//
// Solidity: function create(address _contract, uint256 _tokenId, address seller, address _token, uint256 _price, uint256 _due) returns(uint256 record)
func (_SaleAuction *SaleAuctionTransactorSession) Create(_contract common.Address, _tokenId *big.Int, seller common.Address, _token common.Address, _price *big.Int, _due *big.Int) (*types.Transaction, error) {
	return _SaleAuction.Contract.Create(&_SaleAuction.TransactOpts, _contract, _tokenId, seller, _token, _price, _due)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0x2ad1f1db.
//
// Solidity: function updateDuration(uint256 _saleId, uint256 _due) returns()
func (_SaleAuction *SaleAuctionTransactor) UpdateDuration(opts *bind.TransactOpts, _saleId *big.Int, _due *big.Int) (*types.Transaction, error) {
	return _SaleAuction.contract.Transact(opts, "updateDuration", _saleId, _due)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0x2ad1f1db.
//
// Solidity: function updateDuration(uint256 _saleId, uint256 _due) returns()
func (_SaleAuction *SaleAuctionSession) UpdateDuration(_saleId *big.Int, _due *big.Int) (*types.Transaction, error) {
	return _SaleAuction.Contract.UpdateDuration(&_SaleAuction.TransactOpts, _saleId, _due)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0x2ad1f1db.
//
// Solidity: function updateDuration(uint256 _saleId, uint256 _due) returns()
func (_SaleAuction *SaleAuctionTransactorSession) UpdateDuration(_saleId *big.Int, _due *big.Int) (*types.Transaction, error) {
	return _SaleAuction.Contract.UpdateDuration(&_SaleAuction.TransactOpts, _saleId, _due)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x0f0a1b9e.
//
// Solidity: function updatePrice(uint256 _saleId, address _token, uint256 _price) returns()
func (_SaleAuction *SaleAuctionTransactor) UpdatePrice(opts *bind.TransactOpts, _saleId *big.Int, _token common.Address, _price *big.Int) (*types.Transaction, error) {
	return _SaleAuction.contract.Transact(opts, "updatePrice", _saleId, _token, _price)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x0f0a1b9e.
//
// Solidity: function updatePrice(uint256 _saleId, address _token, uint256 _price) returns()
func (_SaleAuction *SaleAuctionSession) UpdatePrice(_saleId *big.Int, _token common.Address, _price *big.Int) (*types.Transaction, error) {
	return _SaleAuction.Contract.UpdatePrice(&_SaleAuction.TransactOpts, _saleId, _token, _price)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x0f0a1b9e.
//
// Solidity: function updatePrice(uint256 _saleId, address _token, uint256 _price) returns()
func (_SaleAuction *SaleAuctionTransactorSession) UpdatePrice(_saleId *big.Int, _token common.Address, _price *big.Int) (*types.Transaction, error) {
	return _SaleAuction.Contract.UpdatePrice(&_SaleAuction.TransactOpts, _saleId, _token, _price)
}

// SaleAuctionCancelSaleIterator is returned from FilterCancelSale and is used to iterate over the raw logs and unpacked data for CancelSale events raised by the SaleAuction contract.
type SaleAuctionCancelSaleIterator struct {
	Event *SaleAuctionCancelSale // Event containing the contract specifics and raw log

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
func (it *SaleAuctionCancelSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleAuctionCancelSale)
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
		it.Event = new(SaleAuctionCancelSale)
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
func (it *SaleAuctionCancelSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleAuctionCancelSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleAuctionCancelSale represents a CancelSale event raised by the SaleAuction contract.
type SaleAuctionCancelSale struct {
	Contract common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCancelSale is a free log retrieval operation binding the contract event 0x8b0a052bd72fec557d5735593db719bcb565d7afec8d76ee165e317e8cc98409.
//
// Solidity: event CancelSale(address indexed _contract, uint256 _tokenId)
func (_SaleAuction *SaleAuctionFilterer) FilterCancelSale(opts *bind.FilterOpts, _contract []common.Address) (*SaleAuctionCancelSaleIterator, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	logs, sub, err := _SaleAuction.contract.FilterLogs(opts, "CancelSale", _contractRule)
	if err != nil {
		return nil, err
	}
	return &SaleAuctionCancelSaleIterator{contract: _SaleAuction.contract, event: "CancelSale", logs: logs, sub: sub}, nil
}

// WatchCancelSale is a free log subscription operation binding the contract event 0x8b0a052bd72fec557d5735593db719bcb565d7afec8d76ee165e317e8cc98409.
//
// Solidity: event CancelSale(address indexed _contract, uint256 _tokenId)
func (_SaleAuction *SaleAuctionFilterer) WatchCancelSale(opts *bind.WatchOpts, sink chan<- *SaleAuctionCancelSale, _contract []common.Address) (event.Subscription, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	logs, sub, err := _SaleAuction.contract.WatchLogs(opts, "CancelSale", _contractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleAuctionCancelSale)
				if err := _SaleAuction.contract.UnpackLog(event, "CancelSale", log); err != nil {
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

// ParseCancelSale is a log parse operation binding the contract event 0x8b0a052bd72fec557d5735593db719bcb565d7afec8d76ee165e317e8cc98409.
//
// Solidity: event CancelSale(address indexed _contract, uint256 _tokenId)
func (_SaleAuction *SaleAuctionFilterer) ParseCancelSale(log types.Log) (*SaleAuctionCancelSale, error) {
	event := new(SaleAuctionCancelSale)
	if err := _SaleAuction.contract.UnpackLog(event, "CancelSale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SaleAuctionConfirmSaleIterator is returned from FilterConfirmSale and is used to iterate over the raw logs and unpacked data for ConfirmSale events raised by the SaleAuction contract.
type SaleAuctionConfirmSaleIterator struct {
	Event *SaleAuctionConfirmSale // Event containing the contract specifics and raw log

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
func (it *SaleAuctionConfirmSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleAuctionConfirmSale)
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
		it.Event = new(SaleAuctionConfirmSale)
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
func (it *SaleAuctionConfirmSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleAuctionConfirmSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleAuctionConfirmSale represents a ConfirmSale event raised by the SaleAuction contract.
type SaleAuctionConfirmSale struct {
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
func (_SaleAuction *SaleAuctionFilterer) FilterConfirmSale(opts *bind.FilterOpts, _contract []common.Address, solder []common.Address, buyer []common.Address) (*SaleAuctionConfirmSaleIterator, error) {

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

	logs, sub, err := _SaleAuction.contract.FilterLogs(opts, "ConfirmSale", _contractRule, solderRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &SaleAuctionConfirmSaleIterator{contract: _SaleAuction.contract, event: "ConfirmSale", logs: logs, sub: sub}, nil
}

// WatchConfirmSale is a free log subscription operation binding the contract event 0x8484265a68fd7cea22f85d6a338d7534f34ae4ff7d204bfa9faaec9e9e696ba3.
//
// Solidity: event ConfirmSale(address indexed _contract, uint256 _tokenId, address indexed solder, address indexed buyer, address token, uint256 _price)
func (_SaleAuction *SaleAuctionFilterer) WatchConfirmSale(opts *bind.WatchOpts, sink chan<- *SaleAuctionConfirmSale, _contract []common.Address, solder []common.Address, buyer []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SaleAuction.contract.WatchLogs(opts, "ConfirmSale", _contractRule, solderRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleAuctionConfirmSale)
				if err := _SaleAuction.contract.UnpackLog(event, "ConfirmSale", log); err != nil {
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

// ParseConfirmSale is a log parse operation binding the contract event 0x8484265a68fd7cea22f85d6a338d7534f34ae4ff7d204bfa9faaec9e9e696ba3.
//
// Solidity: event ConfirmSale(address indexed _contract, uint256 _tokenId, address indexed solder, address indexed buyer, address token, uint256 _price)
func (_SaleAuction *SaleAuctionFilterer) ParseConfirmSale(log types.Log) (*SaleAuctionConfirmSale, error) {
	event := new(SaleAuctionConfirmSale)
	if err := _SaleAuction.contract.UnpackLog(event, "ConfirmSale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SaleAuctionCreateSaleIterator is returned from FilterCreateSale and is used to iterate over the raw logs and unpacked data for CreateSale events raised by the SaleAuction contract.
type SaleAuctionCreateSaleIterator struct {
	Event *SaleAuctionCreateSale // Event containing the contract specifics and raw log

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
func (it *SaleAuctionCreateSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleAuctionCreateSale)
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
		it.Event = new(SaleAuctionCreateSale)
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
func (it *SaleAuctionCreateSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleAuctionCreateSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleAuctionCreateSale represents a CreateSale event raised by the SaleAuction contract.
type SaleAuctionCreateSale struct {
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
func (_SaleAuction *SaleAuctionFilterer) FilterCreateSale(opts *bind.FilterOpts, _contract []common.Address) (*SaleAuctionCreateSaleIterator, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	logs, sub, err := _SaleAuction.contract.FilterLogs(opts, "CreateSale", _contractRule)
	if err != nil {
		return nil, err
	}
	return &SaleAuctionCreateSaleIterator{contract: _SaleAuction.contract, event: "CreateSale", logs: logs, sub: sub}, nil
}

// WatchCreateSale is a free log subscription operation binding the contract event 0x008f298a385d495e86787887da324d1603491ea0ea157acffd6c5851d6ddfa2f.
//
// Solidity: event CreateSale(address indexed _contract, uint256 _tokenId, address seller, address token, uint256 _price, uint256 _due)
func (_SaleAuction *SaleAuctionFilterer) WatchCreateSale(opts *bind.WatchOpts, sink chan<- *SaleAuctionCreateSale, _contract []common.Address) (event.Subscription, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	logs, sub, err := _SaleAuction.contract.WatchLogs(opts, "CreateSale", _contractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleAuctionCreateSale)
				if err := _SaleAuction.contract.UnpackLog(event, "CreateSale", log); err != nil {
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

// ParseCreateSale is a log parse operation binding the contract event 0x008f298a385d495e86787887da324d1603491ea0ea157acffd6c5851d6ddfa2f.
//
// Solidity: event CreateSale(address indexed _contract, uint256 _tokenId, address seller, address token, uint256 _price, uint256 _due)
func (_SaleAuction *SaleAuctionFilterer) ParseCreateSale(log types.Log) (*SaleAuctionCreateSale, error) {
	event := new(SaleAuctionCreateSale)
	if err := _SaleAuction.contract.UnpackLog(event, "CreateSale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SaleAuctionNewBidIterator is returned from FilterNewBid and is used to iterate over the raw logs and unpacked data for NewBid events raised by the SaleAuction contract.
type SaleAuctionNewBidIterator struct {
	Event *SaleAuctionNewBid // Event containing the contract specifics and raw log

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
func (it *SaleAuctionNewBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleAuctionNewBid)
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
		it.Event = new(SaleAuctionNewBid)
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
func (it *SaleAuctionNewBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleAuctionNewBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleAuctionNewBid represents a NewBid event raised by the SaleAuction contract.
type SaleAuctionNewBid struct {
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
func (_SaleAuction *SaleAuctionFilterer) FilterNewBid(opts *bind.FilterOpts, _contract []common.Address, bidder []common.Address) (*SaleAuctionNewBidIterator, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _SaleAuction.contract.FilterLogs(opts, "NewBid", _contractRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &SaleAuctionNewBidIterator{contract: _SaleAuction.contract, event: "NewBid", logs: logs, sub: sub}, nil
}

// WatchNewBid is a free log subscription operation binding the contract event 0x3a62da8721790771856162caef0b31ff0c28626c1a7bcd82ed8a6a2a6b0009a2.
//
// Solidity: event NewBid(address indexed _contract, uint256 _tokenId, address indexed bidder, address token, uint256 _price)
func (_SaleAuction *SaleAuctionFilterer) WatchNewBid(opts *bind.WatchOpts, sink chan<- *SaleAuctionNewBid, _contract []common.Address, bidder []common.Address) (event.Subscription, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _SaleAuction.contract.WatchLogs(opts, "NewBid", _contractRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleAuctionNewBid)
				if err := _SaleAuction.contract.UnpackLog(event, "NewBid", log); err != nil {
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

// ParseNewBid is a log parse operation binding the contract event 0x3a62da8721790771856162caef0b31ff0c28626c1a7bcd82ed8a6a2a6b0009a2.
//
// Solidity: event NewBid(address indexed _contract, uint256 _tokenId, address indexed bidder, address token, uint256 _price)
func (_SaleAuction *SaleAuctionFilterer) ParseNewBid(log types.Log) (*SaleAuctionNewBid, error) {
	event := new(SaleAuctionNewBid)
	if err := _SaleAuction.contract.UnpackLog(event, "NewBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SaleAuctionUpdateSaleIterator is returned from FilterUpdateSale and is used to iterate over the raw logs and unpacked data for UpdateSale events raised by the SaleAuction contract.
type SaleAuctionUpdateSaleIterator struct {
	Event *SaleAuctionUpdateSale // Event containing the contract specifics and raw log

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
func (it *SaleAuctionUpdateSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleAuctionUpdateSale)
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
		it.Event = new(SaleAuctionUpdateSale)
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
func (it *SaleAuctionUpdateSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleAuctionUpdateSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleAuctionUpdateSale represents a UpdateSale event raised by the SaleAuction contract.
type SaleAuctionUpdateSale struct {
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
func (_SaleAuction *SaleAuctionFilterer) FilterUpdateSale(opts *bind.FilterOpts, _contract []common.Address) (*SaleAuctionUpdateSaleIterator, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	logs, sub, err := _SaleAuction.contract.FilterLogs(opts, "UpdateSale", _contractRule)
	if err != nil {
		return nil, err
	}
	return &SaleAuctionUpdateSaleIterator{contract: _SaleAuction.contract, event: "UpdateSale", logs: logs, sub: sub}, nil
}

// WatchUpdateSale is a free log subscription operation binding the contract event 0xfe8937da95814c69018f7a05dcec87068d26502ab6dabf6516ea314dc1dca32b.
//
// Solidity: event UpdateSale(address indexed _contract, uint256 _tokenId, address seller, address token, uint256 _price, uint256 _due)
func (_SaleAuction *SaleAuctionFilterer) WatchUpdateSale(opts *bind.WatchOpts, sink chan<- *SaleAuctionUpdateSale, _contract []common.Address) (event.Subscription, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	logs, sub, err := _SaleAuction.contract.WatchLogs(opts, "UpdateSale", _contractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleAuctionUpdateSale)
				if err := _SaleAuction.contract.UnpackLog(event, "UpdateSale", log); err != nil {
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

// ParseUpdateSale is a log parse operation binding the contract event 0xfe8937da95814c69018f7a05dcec87068d26502ab6dabf6516ea314dc1dca32b.
//
// Solidity: event UpdateSale(address indexed _contract, uint256 _tokenId, address seller, address token, uint256 _price, uint256 _due)
func (_SaleAuction *SaleAuctionFilterer) ParseUpdateSale(log types.Log) (*SaleAuctionUpdateSale, error) {
	event := new(SaleAuctionUpdateSale)
	if err := _SaleAuction.contract.UnpackLog(event, "UpdateSale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
