package blockchain

import (
	"fmt"
	"time"
	_ "time"
)

type (
	Blockchain struct {
		Chain []*Block `json:"block_chain"`
	}
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

// Listen on channel, if the mempool has sent an array of transactions, only if mempool is full
func (bl *Blockchain) ListenForMempool(empChan chan []Transaction, notifier chan bool) {
	for {
		fmt.Println("forever running")
		select {
		case poolFullSignal := <-notifier:
			fmt.Println(poolFullSignal)
			if poolFullSignal {
				transactions := <-empChan
				fmt.Println("transactions from listen", transactions)
				prevBlock := bl.Chain[len(bl.Chain)-1]
				newBlock := CreateBlock(transactions, prevBlock.BlockHeader.Hash, prevBlock.BlockHeader.Height+1)
				bl.AddBlock(newBlock)
				fmt.Println("Block added successfully to chain")
			}
		default:
			// 	// Sleep for a short duration to avoid busy-waiting
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (bl *Blockchain) PrintChain() {
	fmt.Printf("%#v\n", bl.Chain[0].BlockHeader)
}
