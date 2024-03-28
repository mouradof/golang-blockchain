package blockchain

import (
	"bytes"
	"crypto/sha256"
)

// Block structure
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// Blockchain structure holding a slice of Blocks
type Blockchain struct {
	blocks []*Block
}

// DeriveHash generates the hash of the block
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// CreateBlock creates a new block using the data and the hash of the previous block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// AddBlock adds a block to the blockchain
func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

// Genesis creates the genesis block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockchain initializes a new blockchain with the genesis block
func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}
