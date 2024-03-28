package main

import (
	"fmt"

	"github.com/mouradof/golang-blockchain/blockchain"
)

// main function to demonstrate adding blocks to the blockchain
func main() {
	chain := blockchain.InitBlockchain()

	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.GetBlocks() {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		fmt.Println()
	}
}
