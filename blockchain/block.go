package blockchain

// Block structure
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// Blockchain structure holding a slice of Blocks
type Blockchain struct {
	blocks []*Block
}

// GetBlocks returns the blockchain's slice of blocks
func (chain *Blockchain) GetBlocks() []*Block {
	return chain.blocks
}

// CreateBlock creates a new block using the data and the hash of the previous block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

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
