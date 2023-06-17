package blockchain

import (
	"log"
)

// All transactions are sent to the mempool
// When it full it is emptied into a block that will added to chain

const MAX_POOL_CAP = 1

type Mempool struct {
	TempStore []Transaction
	Next      MoveTransaction
}

func NewMempool() *Mempool {
	pool := make([]Transaction, 0)
	return &Mempool{
		TempStore: pool,
	}
}

// TODO : After testing change backed to transaction type
func (mp *Mempool) AddTransactionToMempool(trx Transaction) {
	if !mp.isMempoolFull() {
		mp.TempStore = append(mp.TempStore, trx)
		log.Println("Added transaction to mempool....")
		return
	}
	log.Println("Mempool full, unable to add new transaction")
}

func (mp *Mempool) isMempoolFull() bool {
	return len(mp.TempStore) > MAX_POOL_CAP
}

//  empty contents of mempool into block
func (mp *Mempool) EmptyMempool() []Transaction {
	tempStoreCopy := make([]Transaction, len(mp.TempStore))

	copy(tempStoreCopy, mp.TempStore)

	mp.TempStore = nil

	return tempStoreCopy

}

// implement MoveTransaction interface

func (mp *Mempool) Execute(trx *Trxs) {
	// trx.MempoolFull = mp.isMempoolFull()
	if mp.isMempoolFull() {
		trxs := mp.EmptyMempool()
		trx.Transactions = append(trx.Transactions, trxs...)
		mp.Next.Execute(trx)
		// empty transactions here
		trx.Transactions = nil
		log.Println("Mempool is full, moved transactions to block....")
		return
	}
	log.Println("Can't move transaction to block, mempool is not yet full....")
}

func (mp *Mempool) SetNext(next MoveTransaction) {
	mp.Next = next
}
