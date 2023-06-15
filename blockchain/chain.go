package blockchain

import (
	"context"
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"
	ev "github.com/qoinpalhq/HQ_CHAIN/events"
)

type (
	Blockchain struct {
		Chain       []*Block `json:"block_chain"`
	}
)

// TODO : Define error messages properly
// Create Channels

func InitializeChain(eventStream *ev.EventStream) *Blockchain {
	// start subscription service
	eventStream.SubscribeMessage(context.Background(), "mempool.full")
	return &Blockchain{
		Chain:       []*Block{CreateGenesisBlock()},
	}
	
}

func ContinueChain() *Blockchain {
	// To be implemented
	return nil
}
func (bl *Blockchain) AddBlock(block *Block) {
	bl.Chain = append(bl.Chain, block)

	fmt.Println("New Block added successfully to chain")
}

// using a singleton design pattern, get a single instance of mempool, that can be shared across components

// Listen on channel, if the mempool has sent an array of transactions, only if mempool is full
// chain subscribes to "mempool.full" topic

func (bl *Blockchain) Process(messages <-chan *message.Message) {
	for msg := range messages {
		fmt.Printf("received message: %s, payload: %s\n", msg.UUID, string(msg.Payload))
		// if msg != nil {
		prevBlock := bl.Chain[len(bl.Chain)-1]
		transactions := DeserializeTxArray(msg.Payload)
		newBlock := CreateBlock(transactions, prevBlock.BlockHeader.Hash, prevBlock.BlockHeader.Height+1)
		bl.AddBlock(newBlock)
		bl.PrintChain()
		msg.Ack()
		// } else {
		// fmt.Println("Yeah")
		// msg.Nack()
		// }
	}
}

func (bl *Blockchain) PrintChain() {
	fmt.Printf("%#v\n", bl.Chain)
}

// func (bl *Blockchain) ListenForMempool(empChan chan []Transaction, notifier chan bool) {
// 	for {
// 		fmt.Println("forever running")
// 		select {
// 		case poolFullSignal := <-notifier:
// 			fmt.Println(poolFullSignal)
// 			if poolFullSignal {
// 				transactions := <-empChan
// 				// fmt.Println("transactions from listen", transactions)

// 			}
// 		default:
// 			// 	// Sleep for a short duration to avoid busy-waiting
// 			time.Sleep(100 * time.Millisecond)
// 		}
// 	}
// }
