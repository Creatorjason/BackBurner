package blockchain


// Chain of Responsibility
// I am using this design pattern, to move transaction across the mempool, block and blockchain 

// var trxs *Trxs


// func init(){
// 	trxs = &Trxs{}
// }


type MoveTransaction interface{

	Execute(*Trxs)
	SetNext(MoveTransaction)
}

// short for transactions
type Trxs struct{
	Transactions []Transaction
	MempoolFull bool
	// AddedToBlock bool
	// AddedToChain bool
}