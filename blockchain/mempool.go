package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"
	ev "github.com/qoinpalhq/HQ_CHAIN/events"
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
	EventStream *ev.EventStream
}

var (
	singleMempoolInstance *Mempool
	once                  sync.Once
)

func NewMempool(eventStream *ev.EventStream) *Mempool {
	tempStore := make([]Transaction, 0)
	return &Mempool{
		TempStore: tempStore,
		MaxCap:    MAX_MEMPOOL_CAP,
		EventStream: eventStream,
	}
}
func GetSingleMempoolInstance(eventStream *ev.EventStream) *Mempool {
	if singleMempoolInstance == nil {
		once.Do(func() {
			singleMempoolInstance = NewMempool(eventStream)
			fmt.Println("✅ New Mempool created!")
		})
	} else {
		fmt.Println("💡 Mempool already created")
	}
	return singleMempoolInstance
}


// function is called when a new transaction is created
func (mp *Mempool) AddTransaction(trx Transaction) {
	// before adding check if length of tempStore is less than MAX_MEMPOOL_CAP
	if !mp.IsMempoolFull() {
		mp.TempStore = append(mp.TempStore, trx)
		fmt.Println("Transactions add to mempool")
	} else {
		// LOGIC to empty tempStore
		mp.EmptyMemPool()
		fmt.Println("Emptied mempool")
	}

}



// Mempool publishes to this topic "mempool.full"
func (mp *Mempool) EmptyMemPool() {
	// copy all the transactions in TempStore
	tempStoreCopy := make([]Transaction, len(mp.TempStore))
	copy(tempStoreCopy, mp.TempStore)
	// original TempStore
	mp.TempStore = nil
	// publish "mempool.full" topic
	payload := SerializeTrxArray(tempStoreCopy)
	mp.EventStream.PublishMessage(payload, "mempool.full")
	

}



// serializes the  transaction array mempool
func SerializeTrxArray(tempStoreCpy []Transaction) []byte {
	buff := new(bytes.Buffer)
	err := gob.NewEncoder(buff).Encode(&tempStoreCpy)
	if err != nil {
		panic(err)
	}
	return buff.Bytes()
}

func DeserializeTxArray(data []byte) []Transaction {
	var trxArr []Transaction
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&trxArr)
	if err != nil {
		panic(err)
	}
	return trxArr
}
func (mp *Mempool) IsMempoolFull() bool {
	return len(mp.TempStore) >= mp.MaxCap
}

func (mp *Mempool) RemoveTransaction() {
	// TO IMPLEMENT
}

func (mp *Mempool) SortTransactionByID() {
	// TO IMPLEMENT
}
