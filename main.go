package main

import (
	"InterviewQuiz/abi/temperaturefactory"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math"
	"math/big"
)

var temperaturefactoryaddress common.Address = 0x903d2ac7e2Ce844855a5e589ba6e9B4A29298C3B

func main() {
	SetTemperature(1.1)
	GetTemperature()
}

func SetTemperature(value float64) {
	//CheckValue
	convertvalue := value * math.Pow10(10)
	finallyvalue := int64(convertvalue)

	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("xxxxx")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	instance, err := temperaturefactory.NewTemperaturefactory(temperaturefactoryaddress, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.SetTemperature(auth, big.NewInt(finallyvalue))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

func GetTemperature() {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	instance, err := temperaturefactory.NewTemperaturefactory(temperaturefactoryaddress, client)
	if err != nil {
		log.Fatal(err)
	}

	version, err := instance.Temperature(nil)
	if err != nil {
		log.Fatal(err)
	}
	convervalue := version.Int64()
	finallyvalue := float64(convervalue) / math.Pow10(10)
	fmt.Println(finallyvalue)
}
