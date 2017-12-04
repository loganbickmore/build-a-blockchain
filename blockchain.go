package main

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	blocks []*Block
}

// AddBlock is a method on *Blockchain that adds a new block to the chain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain creates a new *Blockchain with genesis Block
func NewBlockChain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
