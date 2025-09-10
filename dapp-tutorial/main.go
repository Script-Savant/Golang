package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Geth - latest block", block.Number().Uint64())
}

func GenerateKeyPair() (privateKey *ecdsa.PrivateKey, address string, err error) {
	priv, err := crypto.GenerateKey()
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate key: %w", err)
	}

	privBytes := crypto.FromECDSA(priv)
	fmt.Printf("PRIVATE KEY (dev only) 0x%x\n", privBytes)

	pub := priv.Public()
	pubECDSA := pub.(*ecdsa.PublicKey)
	addr := crypto.PubkeyToAddress(*pubECDSA)
	addrHex := addr.Hex()

	fmt.Println("Address: ", addrHex)

	return priv, addrHex, nil
}
