package blockchain

import (
	"log"
	"time"
)

type BlockChain interface {
	AddBlock(transactions []Transaction) block
	ValidateChain() bool
}

func NewBlockChain() BlockChain {
	c := new(blockchain)
	c.CreateGenesisBlock()
	return c
}

type blockchain struct {
	chain                   []block
	unconfirmedTransactions []Transaction
}

func (c *blockchain) CreateGenesisBlock() {
	genesis := block{}
	genesis.timeStamp = time.Now().Format("02/01/2006 15:04:05 PM")
	genesis.transactions = make([]Transaction, 0)
	genesis.previousHash = "0"
	genesis.hash = genesis.GenerateHash()
	c.chain = append(c.chain, genesis)
}

func (c *blockchain) AddBlock(transactions []Transaction) block {
	previousHash := c.chain[len(c.chain)-1].hash
	newBlock := block{}
	newBlock.timeStamp = time.Now().Format("02/01/2006 15:04:05 PM")
	newBlock.transactions = transactions
	newBlock.previousHash = previousHash
	newBlock.hash = newBlock.GenerateHash()
	c.chain = append(c.chain, newBlock)
	return newBlock
}

func (c *blockchain) ValidateChain() bool {
	for i := 1; i < len(c.chain); i++ {
		current := c.chain[i]
		previous := c.chain[i-1]
		if current.hash != current.GenerateHash() {
			log.Printf("The current hash (%v) of the block does not equal the generated hash of the block.", current.hash)
			return false
		}
		if current.previousHash != previous.GenerateHash() {
			log.Printf("The previous block's hash (%v) does not equal the previous hash value stored in the current block.", current.previousHash)
			return false
		}
	}
	return true
}
