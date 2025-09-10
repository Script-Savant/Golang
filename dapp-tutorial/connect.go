package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"math/big"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// func main() {
// 	client, err := ethclient.Dial("http://127.0.0.1:8545")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer client.Close()

// 	address := common.HexToAddress("0x71562b71999873DB5b286dF957af199Ec94617F7")

// 	balance := getBalance(client, address)

// 	// get the current block number
// 	block, err := client.BlockByNumber(context.Background(), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connected to Geth!")
// 	fmt.Println("Latest block number: ", block.Number().Uint64())
// 	fmt.Println("Balance: ", balance)
// }

// func getBalance(client *ethclient.Client, address common.Address) *big.Int {
// 	balance, err := client.BalanceAt(context.Background(), address, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return  balance
// }