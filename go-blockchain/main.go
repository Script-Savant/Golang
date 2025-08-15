package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"crypto/sha256"
)

// block struct - nonce, timestamp, transactions, previous hash
type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []string
}

// Create a new block and return it
func NewBlock(nonce int, previousHash [32]byte) *Block {
	b := new(Block)
	b.nonce = nonce
	b.previousHash = previousHash
	b.timestamp = time.Now().UnixNano()

	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp      %d\n", b.timestamp)
	fmt.Printf("nonce          %d\n", b.nonce)
	fmt.Printf("previous_hash  %x\n", b.previousHash)
	fmt.Printf("transactions   %s\n", b.transactions)
}

// hash the block
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

// marshall to json
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64    `json:"timestamp"`
		Nonce        int      `json:"nonce"`
		PreviousHash [32]byte `json:"previous_hash"`
		Transactions []string `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

// blockchain structure - pool of transactions, chain of blocks
type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

// create a new block chain and return it
func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

// create block and append it to block chain
func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

// get last block
func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
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

	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(5, previousHash)
	bc.Print()

	previousHash = bc.LastBlock().Hash()
	bc.CreateBlock(10, previousHash)
	bc.Print()
}
