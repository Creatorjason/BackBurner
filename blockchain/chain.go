package blockchain

type (
	Blockchain struct {
		Chain []*Block `json:"block_chain"`
	}
)


func InitializeChain() *Blockchain {
	return &Blockchain{[]*Block{CreateGenesisBlock()}}
}

 

// Listen on channel, if the mempool has sent an array of transactions, only if mempool is full
func (bl *Blockchain) ListenForMempool(){

}