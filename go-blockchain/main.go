package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// block struct - nonce, timestamp, transactions, previous hash
type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

// Create a new block and return it
func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.nonce = nonce
	b.previousHash = previousHash
	b.timestamp = time.Now().UnixNano()

	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp      %d\n", b.timestamp)
	fmt.Printf("nonce          %d\n", b.nonce)
	fmt.Printf("previous_hash  %s\n", b.previousHash)
	fmt.Printf("transactions   %s\n", b.transactions)
}

// blockchain structure - pool of transactions, chain of blocks
type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

// create a new block chain and return it
func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "init hash")
	return bc
}

// create block and append it to block chain
func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n\n", strings.Repeat("*", 25))
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	bc := NewBlockchain()
	bc.Print()
	bc.CreateBlock(5, "hash 1")
	bc.Print()
	bc.CreateBlock(10, "hash 2")
	bc.Print()
}
