package main

import (
	"fmt"
	"os"
)

func checkArgs(args []string) {
	if len(args) <= 1 {
		fmt.Println("invalid numargs")
		os.Exit(1)
	}
}

func main() {
	checkArgs(os.Args)

	switch os.Args[1] {
	case "--create-wallet":
		priv, pub, address, err := createWallet()
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
		fmt.Println("priv: " + priv)
		fmt.Println("pub: " + pub)
		fmt.Println("address: " + address)
	case "--check-balance":
		balance := getBalanceFromStringAddress(os.Args[2])
		fmt.Println("wei balance: " + balance)
	case "--send-eth":
		txid := sendEth(os.Args[2], os.Args[3], os.Args[4])
		fmt.Println("transaction id: " + txid)
	}
}
