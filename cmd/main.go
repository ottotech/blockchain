package main

import (
	"fmt"
	"github.com/ottotech/blockchain"
)

func main()  {
	blockChain := blockchain.NewBlockChain()
	transaction1 := blockchain.Transaction{Amount:100, Sender:"Will", Receiver:"Smith"}
	transaction2 := blockchain.Transaction{Amount:200, Sender:"Ana", Receiver:"Smith"}
	blockChain.AddBlock([]blockchain.Transaction{transaction1})
	blockChain.AddBlock([]blockchain.Transaction{transaction2})
	fmt.Printf("%+v\n", blockChain)
	fmt.Println(blockChain.ValidateChain())
}
