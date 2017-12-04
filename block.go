package main

import (
	"time"
)

// Block is the primary data type used in the blockchain
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
	// other relevant info, such as version
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// NewBlock generates a new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

/* Archived. The proof of work function Run() generates a hash for us.
// SetHash is a method on *Block that sets the blocks hash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}
*/
