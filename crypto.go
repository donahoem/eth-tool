package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getBalanceFromStringAddress(address string) string {
	// Maybe optimize by bubbling the client up to main.
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	return balance.String()
}

// Could possibly create a smart contract that modifies user balance when receiving a deposit.
func createWallet() (privkey string, address string) {
	// Generate privkey key.
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err.Error())
	}
	// Convert the privateKey to bytes.
	privateKeyBytes := crypto.FromECDSA(privateKey)
	// Convert the bytes to string and strip off the '0x'
	privkey = hexutil.Encode(privateKeyBytes)[2:]
	// Get the pubkey key
	publicKey := privateKey.Public()
	// Conver pubkey key to hex
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		log.Fatal(err.Error())
	}

	// Convert to ETH address
	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	// Make sure balance is zero
	balance := getBalance(address)
	if balance.Cmp(big.NewInt(0)) != 0 {
		err := errors.New(fmt.Sprintf("Address %s has non-zero balance", address))
		log.Fatal(err.Error())
	}

	return privkey, address
}

func getBalance(pubkey string) *big.Int {
	ethClient, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress(pubkey)
	balance, err := ethClient.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	return balance
}

func sendEth(privkey string, amount string, toAddressString string) string {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA(privkey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := new(big.Int)
	value, ok = value.SetString(amount, 10)
	if !ok {
		fmt.Println("SetString: error")
		log.Fatal(err)
	}

	gasLimit := uint64(21000) // in units

	gasPrice := big.NewInt(30000000000) // in wei

	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	toAddress := common.HexToAddress(toAddressString)

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("transaction id: %s", signedTx.Hash().Hex())
}
