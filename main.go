package main

import (
	"fmt"
	bl "github.com/qoinpalhq/HQ_CHAIN/blockchain"
)

func main(){
	fmt.Println("🤗")
	chain := bl.InitializeChain()
	// create transactions
	trx1 := bl.CreateTransaction([]byte("Jason"),[]byte("Qoinpal"),100)
	// trx2 := bl.CreateTransaction([]byte("Kendrick"), []byte("Dayo"), 200)
	bl.AddTransactionToMempool(trx1)
	// bl.AddTransactionToMempool(trx2)
	chain.PrintChain()

}