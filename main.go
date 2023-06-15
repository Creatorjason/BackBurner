package main

import (
	"fmt"
	"sync"

	bl "github.com/qoinpalhq/HQ_CHAIN/blockchain"
)

func main() {
	empChan := make(chan []bl.Transaction, 1)
	notifier := make(chan bool, 1)
	mp := bl.GetMempool(empChan, notifier)
	var wg sync.WaitGroup
	fmt.Println("🤗")
	chain := bl.InitializeChain()
	// initialize mempool
	// create transactions
	trx1 := bl.CreateTransaction([]byte("Jason"), []byte("Qoinpal"), 100)
	trx2 := bl.CreateTransaction([]byte("Kendrick"), []byte("Dayo"), 200)
	trx3 := bl.CreateTransaction([]byte("Kendrick"), []byte("Dayo"), 200)
	wg.Add(1)
	go bl.AddTransactionToMempool(trx1, mp)
	wg.Add(1)
	go bl.AddTransactionToMempool(trx2, mp)
	wg.Add(1)
	go bl.AddTransactionToMempool(trx3, mp)
	wg.Add(1)
	go chain.ListenForMempool(empChan, notifier)
	// go func() {
	// defer wg.Done()
	// fmt.Println("AM a waiter")
	// }()

	chain.PrintChain()
	wg.Wait()

}
