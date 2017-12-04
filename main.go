package main

import (
	"fmt"
	"strconv"
	//"../advent-of-code/utils"
)

func main() {
	bc := NewBlockChain()

	bc.AddBlock("Send 1 BTC to addr.A")
	bc.AddBlock("Send 2 BTC to addr.B")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

	}
}
