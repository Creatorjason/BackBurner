package main

import (
	"fmt"
	"sync"

	bl "github.com/qoinpalhq/HQ_CHAIN/blockchain"
	// "context"
	// "time"

	ev "github.com/qoinpalhq/HQ_CHAIN/events"
)

func main() {
		var wg sync.WaitGroup
		var mp *bl.Mempool
		wg.Add(1)
		go func() {
			defer wg.Done()
			e := ev.NewEventStream()
			chain := bl.InitializeChain(e)
			mp = bl.GetSingleMempoolInstance(e)
			fmt.Println(chain, mp)
		}()
	
		wg.Wait() // Wait for chain initialization to complete
	
		trx1 := bl.CreateTransaction([]byte("Jason"), []byte("Qoinpal"), 100)
		trx2 := bl.CreateTransaction([]byte("Kamal"), []byte("Dayo"), 900)
		trx3 := bl.CreateTransaction([]byte("Kay"), []byte("Dayo1"), 300)
		// trx4 := bl.CreateTransaction([]byte("Kdrick"), []byte("Dayo2"), 2900)
		// trx5 := bl.CreateTransaction([]byte("endrick"), []byte("Dayo3"), 500)
		// trx6 := bl.CreateTransaction([]byte("Kendk"), []byte("Dayo4"), 400)
		// trx7:= bl.CreateTransaction([]byte("Kendk"), []byte("Dayo4"), 400)
	
		bl.AddTransactionToMempool(trx1, mp)
		bl.AddTransactionToMempool(trx2, mp)
		bl.AddTransactionToMempool(trx3, mp)
		// bl.AddTransactionToMempool(trx4, mp)
		// bl.AddTransactionToMempool(trx5, mp)
		// bl.AddTransactionToMempool(trx6, mp)
		// bl.AddTransactionToMempool(trx7, mp)
	
		// Continue with the rest of your code
	

	
}
