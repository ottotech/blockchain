package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type block struct {
	timeStamp    string
	transactions []Transaction
	previousHash string
	nonce        int
	hash         string
}

type Transaction struct {
	Amount   int
	Sender   string
	Receiver string
}

func (b *block) GenerateHash() string {
	blockHeader := b.timeStamp + b.StringTransactions() + b.previousHash + strconv.Itoa(b.nonce)
	h := sha256.New()
	h.Write([]byte(blockHeader))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (b *block) StringTransactions() string {
	sb, err := json.Marshal(b.transactions)
	if err != nil {
		log.Fatalln(err)
	}
	return string(sb)
}
