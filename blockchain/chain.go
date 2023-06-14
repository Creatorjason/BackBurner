package blockchain

import (
	"fmt"
	"sync"
)

type (
	Blockchain struct {
		Chain []*Block `json:"block_chain"`
	}
)

var (
	singleMempoolInstance *Mempool
	once                  sync.Once
	notifier              chan string
	empChan               chan []Transaction
)

// Create Channels

func InitializeChain() *Blockchain {
	return &Blockchain{[]*Block{CreateGenesisBlock()}}
}

func ContinueChain() *Blockchain {
	// To be implemented
	return nil
}
func (bl *Blockchain) AddBlock(block *Block) {
	// prevBlock := bl.Chain[len(bl.Chain)-1]
	// // newBlock := CreateBlock(nil, prevBlock.BlockHeader.Hash, prevBlock.BlockHeader.Height+1)
	bl.Chain = append(bl.Chain, block)
}


// using a singleton design pattern, get a single instance of mempool, that can be shared across components

func GetMempool() *Mempool {
	if singleMempoolInstance == nil {
		notifier = make(chan string, 12)
		empChan = make(chan []Transaction, 12)
		mempool := NewMempool(empChan, notifier)
		once.Do(func() {
			singleMempoolInstance = mempool
			fmt.Println("✅ New Mempool created!")
		})
		fmt.Println("💡 Mempool already created")
	}
	return singleMempoolInstance
}

// Listen on channel, if the mempool has sent an array of transactions, only if mempool is full
func (bl *Blockchain) ListenForMempool() {
	for {
		poolFullSignal := <- notifier
		if poolFullSignal == "mempool full"{
			transactions := <- empChan
			prevBlock := bl.Chain[len(bl.Chain) - 1]
			newBlock := CreateBlock(transactions, prevBlock.BlockHeader.Hash, prevBlock.BlockHeader.Height + 1)
			bl.AddBlock(newBlock)
			fmt.Println("Block added successfully to chain")
			// Refactor later
		}
	}
}




func (bl *Blockchain) PrintChain(){
	fmt.Println(bl)
}
