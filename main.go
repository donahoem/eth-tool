package main

import (
	"fmt"
	"os"
)

func main() {
	checkArgs()

	switch os.Args[1] {
	case "--create-wallet":
		priv, pub, address := createWallet()
		fmt.Println("priv: " + priv)
		fmt.Println("pub: " + pub)
		fmt.Println("address: " + address)
	case "--check-balance":
		balance := getBalanceFromStringAddress(os.Args[2])
		fmt.Println("wei balance: " + balance)
	case "--send-eth":
		txid := sendEth(os.Args[2], os.Args[3], os.Args[4])
		fmt.Println("transaction id: " + txid)
	default:
		printUsage()
	}
}

func checkArgs() {
	if len(os.Args) <= 1 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "--create-wallet":
		if len(os.Args) != 2 {
			fmt.Println("option '--create-wallet' does not take any extra arguments")
			printUsage()
			os.Exit(1)
		}
	case "--check-balance":
		if len(os.Args) != 3 {
			fmt.Println("option '--check-balance' requires an ethereum address as an argument")
			printUsage()
			os.Exit(1)
		}
	case "--send-eth":
		if len(os.Args) != 5 {
			fmt.Println("option '--send-eth' requires a private key, amount, and to address as arguments")
			printUsage()
			os.Exit(1)
		}
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("usage: eth-tool [--create-wallet] [--check-balance <address>] [--send-ether <privkey> <amount> <toAddress>]")
}
