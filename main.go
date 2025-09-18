package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash []byte
	Data []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

type Blockchain struct {
	blocks []*Block
}

func (chain *Blockchain) AddBlock(data string){
	prevBlock := chain.blocks[len(chain.blocks) - 1]
	new := CreateBlock(data,prevBlock.Hash) 
	chain.blocks = append(chain.blocks, new)		
}

func Genesis() *Block{
	return CreateBlock("Genesis", []byte{})
}

func InitBlockchain() *Blockchain{
	return &Blockchain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockchain() 

	chain.AddBlock("FirstBLock")
	chain.AddBlock("SecondBlock")
	chain.AddBlock("ThirdBlock")

	for _, block := range chain.blocks {
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}