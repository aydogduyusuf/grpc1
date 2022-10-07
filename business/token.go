package business


import (
	"context"
	"crypto/ecdsa"
	"fmt"

	//"go/types"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

func InitNetwork(networkRPC string) *ethclient.Client {
	var err error
	client, err = ethclient.Dial(networkRPC)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to blockchain network")
	return client
}

func NewTransactOpts(privateKey *ecdsa.PrivateKey) *bind.TransactOpts {
	//fetch networkID
	networkID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	txOps, err := bind.NewKeyedTransactorWithChainID(privateKey, networkID)
	if err != nil {
		log.Fatal(err)
	}
	return txOps
}

func SetTransactOpts(address common.Address, txOps *bind.TransactOpts, value *big.Int, gasLimit uint64) {
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatal(err)
	}
	txOps.Nonce = big.NewInt(int64(nonce))
	//txOps.Value = value
	txOps.GasLimit = gasLimit
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	txOps.GasPrice = gasPrice
}

func DeployToken(address common.Address, privateKey *ecdsa.PrivateKey, name string, symbol string, supply *big.Int) (common.Address, *Business) {
	txOps := NewTransactOpts(privateKey)
	SetTransactOpts(address, txOps, big.NewInt(0), uint64(21000))
	
	contractAddress, tx, instance, err := DeployBusiness(txOps, client, name, symbol, address, supply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex())

	return contractAddress, instance
}

func TransferContract(privateKey *ecdsa.PrivateKey, public common.Address, address common.Address, to common.Address, amount *big.Int) (common.Hash, error) {
	txOps := NewTransactOpts(privateKey)
	SetTransactOpts(public, txOps, amount, uint64(2100000))
	txOps.Context = nil
	blockInstance, err := NewBusiness(address, client)
	if err != nil {
		log.Fatal("blockchain instance error: ", err)
	}

	tx, err := blockInstance.Transfer(txOps, to, amount)

	return tx.Hash(), err
}