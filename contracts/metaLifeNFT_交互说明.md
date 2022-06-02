#交互说明
合约地址: https://github.com/ubltroll/MetaLifeProtocol/tree/main
文档版本: v0.1
合约已经测试通过，测试用例在test文件夹下
##合约部署
推荐使用solidity 0.8.1编译，启用大小优化，优化次数200次。编译完成后有三个合约需要部署，依次是主合约metaMaster，一口价交易合约salePlain，限时拍卖合约saleAuction。

部署时，先部署主合约metaMaster，然后以主合约metaMaster的地址作为参数部署一口价交易合约salePlain和限时拍卖合约saleAuction。

部署后必须进行初始化设置，参见参数设置一节。

##参数设置
主合约metaMaster需要使用metaMaster的owner账号，默认为部署metaMaster的地址。

###初始化设置
1. 关联交易合约（部署后必须进行）
接口：` metaMaster -> addMarketSale(address)`
权限：metaMaster管理员
参数：
    -address: 交易合约地址
参数是需要添加的交易合约，推荐先添加一口价交易合约salePlain，再添加限时拍卖合约saleAuction，这样0对应限价交易，1对应拍卖，下文将使用这样的对应策略。
2. 设置支持的交易代币（部署后必须进行）
接口：` metaMaster -> setTokenSupport(address, bool)`
权限：metaMaster管理员
参数：
    -address: 代币地址
    -bool: true为支持，false为取消
注意，零地址address(0)为SMT(链原生代币)，初始化合约时已经自动支持。
3. 设置平台交易手续费
接口：` metaMaster -> setFeeConfig(address, uint)`
权限：metaMaster管理员
参数：
    -address: 收取手续费地址
    -uint: 手续费，单位是BP基点，7.5%即为750(默认值)

###创建、管理NFT集合
1. 创建集合
接口：` metaMaster -> createCollection(string, string, string, uint, string, address, uint)`
权限：public
参数：
    -string: collection name
    -string: collection symbol
    -string: collection baseURI: 原来代码此处写死https://gateway.pinata.cloud/ipfs/  ？
    -uint: 最大供应量，不能超过uint96
    -string: collection的描述信息，可以是描述，也可以是logo图片地址
    -address: 接收版权费的地址
    -uint: 版权费率，单位是BP基点，2.5%即为250
创建后会返回一个地址，即为创建的collection地址，可以按用户查询collection接口获取其地址。
2. collection的创建者
接口：` metaMaster -> collectionOwner(address) view`
权限：public
参数：
    -address: collection地址
返回该collection的创建者
3. 按用户查询collection
接口：` metaMaster -> userOwnedCollections(uint) view`
权限：public
参数：
    -uint: 序号
返回对应用户创建的collection地址
4. 按用户查询collection-总数
接口：` metaMaster -> userOwnedCollectionNum(address) view`
权限：public
参数：
    -address: collection地址
返回对应用户创建的collection地址总数
5. 设置collection描述信息
接口：` metaMaster -> setCollectionMetaInfo(address, string)`
权限：对应collection的创建者
参数：
    -address: collection地址
    -string: collection的描述信息
6. 设置collection版权费
接口：` metaMaster -> setCollectionRoyalty(address, address, uint)`
权限：对应collection的创建者
参数：
    -address: collection地址
    -address: 接收版权费的地址
    -uint: 版权费率，单位是BP基点，2.5%即为250
返回对应用户创建的collection地址总数
7. Mint NFT
接口：` metaMaster -> mint(address, address, string)`
权限：对应collection的创建者
参数：
    -address: collection地址
    -address: 接收NFT地址
    -string: tokenURI，与baseURI共同组成URI地址
mint出来的NFT编号是从1开始递增
8. Sell NFT
接口：` metaMaster -> sell(address, uint, uint, address, uint, uint)`
权限：对应NFT的owner，如果是metaMaster创建的collection则无需授权
参数：
    -address: collection地址
    -uint: NFT编号
    -uint: 销售方式，0是限价1是拍卖
    -address: 定价代币，零地址则为SMT，必须是设置支持的代币
    -uint: 价格，限价为价格，拍卖是起价，不能大于uint128
    -uint: 截止时间戳，单位为秒，超过后无法购买/竞拍
9. Mint and Sell
接口：` metaMaster -> mintAndSell(address, string, uint, address, uint, uint)`
权限：对应NFT的owner，如果是metaMaster创建的collection则无需授权
参数：
    -address: collection地址
    -string: tokenURI，与baseURI共同组成URI地址
    -uint: 销售方式，0是限价1是拍卖
    -address: 定价代币，零地址则为SMT，必须是设置支持的代币
    -uint: 价格，限价为价格，拍卖是起价，不能大于uint128
    -uint: 截止时间戳，单位为秒，超过后无法购买/竞拍
一个交易内完成mint和sell
##合约交互
###NFTInfoStorage（主合约继承自NFTInfoStorage）
这里的查询返回的数据结构体定义如下：
struct salesInfo {
   address sales; 市场地址
   address collection; NFT合约
   uint96 token_id; NFT编号
   address token; 交易定价代币地址，零地址是SMT
   uint128 price; 市场是限价合约时是限价价格，市场是拍卖合约时是当前最高价格
   uint64 duetime; 市场是限价合约时是限价价格，市场是拍卖合约时是当前最高价格
   address seller; NFT卖家
}
注意solidity0.8.0以上原生支持返回自定义结构体。
这里的查询主要是为了前端批量查询做的简单支持，如果需要更详细的前端显示最好实现后端缓存信息，超出本文档范围。
1. 查询在售的单个NFT
接口：` metaMaster -> getSaleInfo(address, uint) view`
参数：
    -address: collection地址
    -uint: NFT编号
返回salesInfo，如果是salesInfo(0)说明不在售或者nft不存在
2. 查询在售的单个NFT，按编码record
接口：` metaMaster -> getSaleInfouint) view`
参数：
    -uint: 编码的记录
同1，只是编码了collection地址和编号
3. 批量返回在售的NFT信息
接口：` metaMaster -> getSales(uint, uint) view`
参数：
    -uint: limit
    -uint: offset
返回salesInfo数组，记录按提交市场出售时间排序，新上架的在前
4. 返回在售的NFT数量
接口：` metaMaster -> getSalesCount() view`
返回uint
5. 批量返回collection中在售的NFT信息
接口：` metaMaster -> getSalesByCollection(address, uint, uint) view`
参数：
    -address: NFT合约
    -uint: limit
    -uint: offset
返回salesInfo数组，等于按tokenid批量查询，从offset查到offset+limit
6. 批量返回用户在售的NFT信息
接口：` metaMaster -> getSalesByUser(address, uint, uint) view`
参数：
    -address: 用户地址
    -uint: limit
    -uint: offset
返回salesInfo数组，记录按提交市场出售时间排序，新上架的在前
7. 返回用户在售的NFT数量
接口：` metaMaster -> getSalesCountByUser(address) view`
参数：
    -address: 用户地址
返回uint
8. 批量返回用户的竞拍中的NFT
接口：` metaMaster -> getUserBids(address, uint, uint) view`
参数：
    -address: 用户地址
    -uint: limit
    -uint: offset
返回salesInfo数组，记录按竞拍时间排序，新上架的在前
9. 返回用户的竞拍中的NFT数量
接口：` metaMaster -> getUserBidsCount(address) view`
参数：
    -address: 用户地址
返回uint
10. 批量返回用户购买的NFT信息
接口：` metaMaster -> getUserBuys(address, uint, uint) view`
参数：
    -address: 用户地址
    -uint: limit
    -uint: offset
返回salesInfo数组，记录按购买时间排序，新上架的在前
此处的salesInfo中price是购买价格，duetime是购买时间
11. 返回用户购买的NFT数量
接口：` metaMaster -> getUserBuysCount(address) view`
参数：
    -address: 用户地址
返回uint
###NFTcollection合约
metaMaster创建出来的collection合约是一个标准的ERC721合约，同时支持ERC721Enumerable扩展，即可以用tokenOfOwnerByIndex遍历用户名下持有的NFT。
NFTcollection合约的owner是主合约，创建者在主合约中记录。
同时还具有如下扩展：
1. 接收版权费的地址
接口：` address public royaltiesReceiver`
2. 版权费
接口：` uint public royaltiesPercentageInBips`
3. 集合信息
接口：` string public metaInfo`
可以存储一些自定义信息，比如描述、logo、banner等等
###交易合约
包括了限价交易和拍卖交易，两个合约遵循共同的一套接口标准（见interface/ISales）
注意创建交易是在主合约进行，当然可以自行将NFT转移到交易合约然后自行调用create函数来创建交易
与前端交互有关的接口如下，不言自明的参数不再赘述。
0. 获取打包的saleId
接口：` getSaleId(address _contract, uint _tokenId) external view returns (uint _saleId)`
也可以参考RecordPacker自行实现
1. 取消挂单，取回NFT
接口：` cancel(uint _saleId) external`
注意：竞拍合约中一旦有人出价(getSaleInfo查询得到的bidder不是零地址)，则不可取消。
2. 更新价格
接口：` updatePrice(uint _saleId, address _token, uint _price) external`
注意：竞拍合约中一旦有人出价(getSaleInfo查询得到的bidder不是零地址)，则不可更新。
3. 更新截止时间
接口：` updateDuration(uint _saleId, uint _due) external`
注意：竞拍合约中一旦有人出价(getSaleInfo查询得到的bidder不是零地址)，则不可更新。
4. 出价（以SMT付款）
接口：` bidWithValue(uint _saleId) external payable`
注意：对应是_saleId的计价代币必须是SMT（零地址）
出价的价格附带在交易的value中，必须大于等于当前价格（竞拍时必须大于）
限价交易中此操作将直接购买NFT，竞拍中将退回前一个竞拍者的出价
5. 出价（以erc20付款）
接口：` bidWithToken(uint _saleId, uint _amount) external payable`
注意：对应是_saleId的计价代币必须是ERC20（非零地址）
调用前注意检查ERC20合约对市场地址的授权，approve后才能调用
限价交易中此操作将直接购买NFT，竞拍中将退回前一个竞拍者的出价
6. 获取NFT
接口：` claim(uint _saleId) external`
竞拍中截止时间到后由最高出价者（记录的bidder）获取NFT，并结算价格
7. 检查是否能获取NFT
接口：` claimable(uint _saleId) external view returns (bool)`
8. 获取交易信息
接口：` getSaleInfo(uint _saleId) external view returns (bool _isOpen, address _token, uint _price, uint _due, address _seller, address _bidder);`
其中_isOpen为真的可以购买/竞价，方便前端判断。
