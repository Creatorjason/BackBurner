package blockchain

// Chain of Responsibility
// I am using this design pattern, to move transaction across the mempool, block and blockchain

// var trxs *Trxs

// func init(){
// 	trxs = &Trxs{}
// }
type CoR struct {
	Mempool    *Mempool
	Blockchain *Blockchain
	Trxs       *Trxs
}

func NewCoR() *CoR {
	return &CoR{
		Mempool:    NewMempool(),
		Blockchain: InitializeChain(),
		Trxs:       NewTrxs(),
	}
}

type MoveTransaction interface {
	Execute(*Trxs)
	SetNext(MoveTransaction)
}

// short for transactions
// a wrapper passed around components of CoR
type Trxs struct {
	Transactions []Transaction
	// MempoolFull bool
	// AddedToBlock bool
	// AddedToChain bool
}

func NewTrxs() *Trxs {
	t := make([]Transaction, 0)
	return &Trxs{
		Transactions: t,
	}
}
