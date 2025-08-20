package main

import (
	"fmt"
	"go-blockchain/block"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	myBlockchainAddress := "alex"
	bc := block.NewBlockchain(myBlockchainAddress)
	bc.Print()

	bc.AddTransaction("A", "B", 1.0)
	bc.Mining()
	bc.Print()

	bc.AddTransaction("C", "D", 2.0)
	bc.AddTransaction("x", "Y", 3.0)
	bc.Mining()
	bc.Print()

	fmt.Printf("my %.1f\n", bc.CalculateTotalAmount("alex"))
	fmt.Printf("C %.1f\n", bc.CalculateTotalAmount("C"))
	fmt.Printf("D %.1f\n", bc.CalculateTotalAmount("D"))

}
