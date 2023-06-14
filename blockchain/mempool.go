package blockchain


// In this implementation
// I add transactions to the mempool, until a particular amount of transaction
// are present, then a block is automatically created and the transactions are moved
//  from the mempool to the block.

// Mempool will be in RAM for now
// TODO: Implement a way to either backtrack and gather previous transactions stored in mempool, if RAM is comprised
// TODO: or a method to replicate across multiple RAM or persist on DISK >>>> {Thinking out loud 💨💨}
type Mempool struct{
	TempStore []Transaction 
}

func NewMempool() *Mempool{
	return nil
}

func (mp *Mempool) AddTransaction(){

}

func (mp *Mempool) RemoveTransaction(){

}

func (mp *Mempool) SortTransactionByID(){
	// TO Implement
}

func (mp *Mempool) IsMempoolFull() bool{
	return false
}

func (mp *Mempool) EmptyMemPool() {

}

func (mp *Mempool) NotifyChain(){
	
}

// func (mp Mempoo