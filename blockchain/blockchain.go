package blockchain

// Blockchain structure holding a slice of Blocks
type Blockchain struct {
	blocks []*Block
}

// InitBlockchain initializes a new blockchain with the genesis block
func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

// AddBlock adds a block to the blockchain
func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}
