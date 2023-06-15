package blockchain

import (
	"fmt"
	"sync"
)

// In this implementation
// I add transactions to the mempool, until a particular amount of transaction
// are present, then a block is automatically created and the transactions are moved
//  from the mempool to the block.

// Mempool will be in RAM for now
// TODO: Implement a way to either backtrack and gather previous transactions stored in mempool, if RAM is comprised
// TODO: or a method to replicate across multiple RAM or persist on DISK >>>> {Thinking out loud 💨💨}
const MAX_MEMPOOL_CAP = 2

type Mempool struct {
	MaxCap    int
	TempStore []Transaction
	// Empty channel used ny Mempool to empty temp store to chain
	EmpChan chan []Transaction
	// Notifier, simply tells chain of the available of a transaction array
	Notifier chan bool
}

var (
	singleMempoolInstance *Mempool
	once                  sync.Once
	// notifier              chan string
	// empChan               chan []Transaction
)

func NewMempool(empChan chan []Transaction, notifier chan bool) *Mempool {
	tempStore := make([]Transaction, 0)
	return &Mempool{
		TempStore: tempStore,
		EmpChan:   empChan,
		Notifier:  notifier,
		MaxCap:    MAX_MEMPOOL_CAP,
	}
}

// function is called when a new transaction is created
func (mp *Mempool) AddTransaction(trx Transaction) {
	// before adding check if length of tempStore is less than MAX_MEMPOOL_CAP
	
	if mp.IsMempoolFull() {
		mp.TempStore = append(mp.TempStore, trx)
		fmt.Println("Transactions add to mempool")
	} else {
		// LOGIC to empty tempStore
		mp.EmptyMemPool()
		fmt.Println("Emptied mempool")
	}

}

func (mp *Mempool) RemoveTransaction() {
	// TO IMPLEMENT
}

func (mp *Mempool) SortTransactionByID() {
	// TO IMPLEMENT
}

func (mp *Mempool) IsMempoolFull() bool {
	return len(mp.TempStore) < mp.MaxCap
}

func (mp *Mempool) EmptyMemPool() {
	// copy all the transactions in TempStore
	tempStoreCopy := make([]Transaction, len(mp.TempStore))
	copy(tempStoreCopy, mp.TempStore)
	// original TempStore
	mp.TempStore = nil
	// set through channel
	mp.EmpChan <- tempStoreCopy
	// notify chain
	mp.Notifier <- true
}

func GetMempool(empChan chan []Transaction, notifier chan bool) *Mempool {
	if singleMempoolInstance == nil {
		once.Do(func() {
			singleMempoolInstance = NewMempool(empChan, notifier)
			fmt.Println("✅ New Mempool created!")
		})
	} else {
		fmt.Println("💡 Mempool already created")
	}
	return singleMempoolInstance
}
