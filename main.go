package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// Create a block custom data type for block features.
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// Blockchain data type consist of slice of block datatype
type BlockChain struct {
	blocks []*Block
}

// Generate block hash using sha256 algorithm
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// Create new block using previous block's hash
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// Add blcoks to the chain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

// Create genenis block which will not have any hash thats why you have an empty slice of bytes
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()
	chain.AddBlock("1st block")
	chain.AddBlock("2nd block")

	for _, block := range chain.blocks {
		fmt.Printf(string(block.PrevHash))
	}
}
