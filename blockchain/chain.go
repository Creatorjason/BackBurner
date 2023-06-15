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

var (
	blockchain *Blockchain
	// works for now...
	eStream <-chan *message.Message
)

// TODO : Define error messages properly
// Create Channels

func InitializeChain(eventStream *ev.EventStream) *Blockchain {
	// start subscription service
	eStream = eventStream.SubscribeMessage(context.Background(), "mempool.full")
	blc := &Blockchain{
		Chain:       []*Block{CreateGenesisBlock()},
	}
	// global variable that is used by other functions for side effects
	blockchain = blc
	return blc
	
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

func (bl *Blockchain) ProcessDataFromMemPool(messages <-chan *message.Message, stop chan struct{}) {
	for {
		select{

		
		case msg := <- messages:
		fmt.Printf("received message: %s, payload: %s\n", msg.UUID, string(msg.Payload))
		// if msg != nil {
		prevBlock := bl.Chain[len(bl.Chain)-1]
		transactions := DeserializeTxArray(msg.Payload)
		fmt.Printf("%#x\n", transactions)
		newBlock := CreateBlock(transactions, prevBlock.BlockHeader.Hash, prevBlock.BlockHeader.Height+1)
		bl.AddBlock(newBlock)
		bl.PrintChain()
		msg.Ack()
		// } else {
		// fmt.Println("Yeah")
		// msg.Nack()
		// }
		// carries cancellation signal over channel
		case <- stop:
			return
	}
}
}
// A function that will be called from mempool, when a pool is emptied, to begin processing
func StartDataProcessing(stop chan struct{}){
	go blockchain.ProcessDataFromMemPool(eStream, stop)
} 
func (bl *Blockchain) PrintChain() {
	fmt.Printf("%#v\n", len(bl.Chain))
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
