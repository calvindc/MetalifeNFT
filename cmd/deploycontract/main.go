package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"

	"github.com/MetalifeNFT/contracts/abi/metaMaster"
	"github.com/MetalifeNFT/contracts/abi/saleAuction"
	"github.com/MetalifeNFT/contracts/abi/salePlain"
	"github.com/MetalifeNFT/mutils"
	"github.com/MetalifeNFT/params"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "address",
			Usage: "The ethereum address you would like Photon to use and for which a keystore file exists in your local system.",
		},
		utils.DirectoryFlag{
			Name:  "keystore-path",
			Usage: "If you have a non-standard path for the ethereum keystore directory provide it using this argument. ",
			Value: utils.DirectoryString(params.DefaultKeyStoreDir()),
		},
		cli.StringFlag{
			Name: "eth-rpc-endpoint",
			Usage: `"host:port" address of ethereum JSON-RPC server.\n'
	           'Also accepts a protocol prefix (ws:// or ipc channel) with optional port',`,
			Value: params.NetworkRPC,
		},
	}
	app.Action = mainCtx
	app.Name = "deploycontract"
	app.Version = "0.1"
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func mainCtx(ctx *cli.Context) error {
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	conn, err := ethclient.Dial(ctx.String("eth-rpc-endpoint"))
	if err != nil {
		log.Fatalf(fmt.Sprintf("Failed to connect to the Ethereum client: %v", err))
	}
	address := common.HexToAddress(ctx.String("address"))
	address, keybin, err := mutils.PromptAccount(address, ctx.String("keystore-path"), "")
	if err != nil {
		log.Fatalf(fmt.Sprintf("failed to unlock account %s", err))
	}
	fmt.Println("start to deploy ...")
	key, err := crypto.ToECDSA(keybin)
	if err != nil {
		log.Fatalf(fmt.Sprintf("failed to parse priv key %s", err))
	}
	deployContractMetaMaster(key, conn)
	deployContractSalePlain(key, conn)
	deployContractSaleAuction(key, conn)

	return nil
}

func deployContractMetaMaster(key *ecdsa.PrivateKey, conn *ethclient.Client) {
	auth, err := bind.NewKeyedTransactorWithChainID(key, params.ChainID)
	if err != nil {
		log.Fatalf("failed to NewKeyedTransactor [MetaMaster] %s", err)
	}
	contractAddress, tx, _, err := metaMaster.DeployMetaMaster(auth, conn)
	if err != nil {
		log.Fatalf("failed to deploy registry [MetaMaster] %s", err)
	}
	ctx := context.Background()
	_, err = bind.WaitDeployed(ctx, conn, tx)
	if err != nil {
		log.Fatalf("failed to deploy contract when mining :%v", err)
	}
	params.RegistryAddressMetaMaster = contractAddress
	fmt.Printf("deploy registry complete...\n")
	fmt.Printf("[MetaMaster] RegistryContractAddress=%s\n \n", contractAddress.String())
}

func deployContractSalePlain(key *ecdsa.PrivateKey, conn *ethclient.Client) {
	auth, err := bind.NewKeyedTransactorWithChainID(key, params.ChainID)
	if err != nil {
		log.Fatalf("failed to NewKeyedTransactor [SalePlain] %s", err)
	}
	contractAddress, tx, _, err := salePlain.DeploySalePlain(auth, conn, params.RegistryAddressMetaMaster)
	if err != nil {
		log.Fatalf("failed to deploy registry [SalePlain] %s", err)
	}
	ctx := context.Background()
	_, err = bind.WaitDeployed(ctx, conn, tx)
	if err != nil {
		log.Fatalf("failed to deploy contract when mining :%v", err)
	}
	params.RegistryAddressSalePlain = contractAddress
	fmt.Printf("deploy registry complete...\n")
	fmt.Printf("[SalePlain] RegistryContractAddress=%s\n \n", contractAddress.String())
}

func deployContractSaleAuction(key *ecdsa.PrivateKey, conn *ethclient.Client) {
	auth, err := bind.NewKeyedTransactorWithChainID(key, params.ChainID)
	if err != nil {
		log.Fatalf("failed to NewKeyedTransactor [SaleAuction] %s", err)
	}
	contractAddress, tx, _, err := saleAuction.DeploySaleAuction(auth, conn, params.RegistryAddressMetaMaster)
	if err != nil {
		log.Fatalf("failed to deploy registry [SaleAuction] %s", err)
	}
	ctx := context.Background()
	_, err = bind.WaitDeployed(ctx, conn, tx)
	if err != nil {
		log.Fatalf("failed to deploy contract when mining :%v", err)
	}
	params.RegistryAddressSaleAuction = contractAddress
	fmt.Printf("deploy registry complete...\n")
	fmt.Printf("[SaleAuction] RegistryContractAddress=%s\n \n", contractAddress.String())
}
