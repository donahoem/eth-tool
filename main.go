package main

import (
	"fmt"
	"os"
)

func main() {
	checkArgs()

	switch os.Args[1] {
	case "--create-wallet":
		priv, address := createWallet()
		fmt.Println("private key: " + priv)
		fmt.Println("address: " + address)
	case "--check-balance":
		balance := getBalanceFromStringAddress(os.Args[2])
		fmt.Println("wei balance: " + balance)
	case "--send-wei":
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
		if len(os.Args) != 3 || len(os.Args[2]) != 42 {
			fmt.Println("option '--check-balance' requires an address (42 characters starting with '0x') as an argument")
			printUsage()
			os.Exit(1)
		}
	case "--send-wei":
		if len(os.Args) != 5 || len(os.Args[4]) != 42 {
			fmt.Println("option '--send-wei' requires a sender private key, amount (in wei), and recipient address (42 characters starting with '0x') as arguments")
			printUsage()
			os.Exit(1)
		}
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("usage: eth-tool [--create-wallet] [--check-balance <address>] [--send-wei <privkey> <amount> <toAddress>]")
}
