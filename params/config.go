package params

import (
	"math/big"
	"path/filepath"

	"crypto/ecdsa"
	"time"

	"github.com/MetalifeNFT/mutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/node"
)

const (
	// These are the multipliers for ether denominations.
	// Example: To get the wei value of an amount in 'douglas', use
	//
	//    new(big.Int).Mul(value, big.NewInt(params.Douglas))
	//
	Wei      = 1
	Ada      = 1e3
	Babbage  = 1e6
	Shannon  = 1e9
	Szabo    = 1e12
	Finney   = 1e15
	Ether    = 1e18
	Einstein = 1e21
	Douglas  = 1e42
)

//Config is configuration for Photon,
type Config struct {
	EthRPCEndPoint  string
	Host            string
	Port            int
	PrivateKey      *ecdsa.PrivateKey
	RevealTimeout   int
	SettleTimeout   int
	DataBasePath    string
	MsgTimeout      time.Duration
	UseRPC          bool
	UseConsole      bool
	APIHost         string
	APIPort         int
	RegistryAddress common.Address
	DataDir         string
	MyAddress       common.Address
	Debug           bool
	DebugCrash      bool
	IsMeshNetwork   bool
	HTTPUsername    string
	HTTPPassword    string
}

var NetworkRPC = "https://jsonapi1.smartmesh.cn"
var NetworkName = "Spectrum"
var NetworkSymbol = "SMT"

var RegistryAddressMetaMaster = mutils.EmptyAddress
var RegistryAddressSalePlain = mutils.EmptyAddress
var RegistryAddressSaleAuction = mutils.EmptyAddress

var ChainID = big.NewInt(20180430)

//DefaultKeyStoreDir keystore path of ethereum
func DefaultKeyStoreDir() string {
	return filepath.Join(node.DefaultDataDir(), "keystore")
}
