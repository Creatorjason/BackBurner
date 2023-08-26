package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"strconv"

	mk "github.com/cbergoon/merkletree"
)

type (
	Transaction struct {
		ID           string
		SenderAddr   string
		ReceiverAddr string
		Amount       int
	}
	// Simplifying this for now, since am not dealing with UTXOs
	// TxInputs struct {
	// 	TXID []byte
	// 	Vout int
	// 	// sig of sender
	// 	Sig []byte
	// 	// pub key/ address of receiver
	// 	PubKey []byte
	// }
	// TxOutputs struct {
	// 	Value   int
	// 	Address []byte
	// }
)

// Implement the Content Interface, so that I can use github.com/cbergoon/merkletree

func (trx Transaction) CalculateHash() ([]byte, error) {
	hash := sha256.Sum256(trx.SerializeTrx())
	return hash[:], nil
}

func (trx Transaction) Equals(other mk.Content) (bool, error) {
	// two transactions are equal if there hashes are
	otherTrx, err := other.(Transaction).CalculateHash()
	if err != nil {
		return false, errors.New("value is not of type Transaction")
	}
	trxHash, _ := trx.CalculateHash()
	return string(trxHash) == string(otherTrx), nil
}

func (trx Transaction) SerializeTrx() []byte {
	buff := new(bytes.Buffer)
	err := gob.NewEncoder(buff).Encode(trx)
	if err != nil {
		panic(err)
	}
	return buff.Bytes()
}

func GetMerkleRoot(trxs []Transaction) []byte {
	var trxArray []mk.Content
	for _, trx := range trxs {
		trxArray = append(trxArray, trx)
	}
	mTree, err := mk.NewTree(trxArray)
	if err != nil {
		panic(err)
	}
	return mTree.MerkleRoot()
}

func CreateTransaction(from, to, desc string, amount int) Transaction {
	// Hash of transaction data
	trx := Transaction{
		SenderAddr:   from,
		ReceiverAddr: to,
		Amount:       amount,
		Desc: desc,
	}
	amountStr := strconv.Itoa(amount)
	b := bytes.Join([][]byte{[]byte(from), []byte(to), []byte(amountStr),[]byte(desc)}, []byte{})
	hash := sha256.Sum256(b)
	trx.ID = hex.EncodeToString(hash[:])
	return trx
}

// func AddTransactionToMempool(trx Transaction, mempool *Mempool) {
// 	mempool.AddTransaction(trx)
// 	// fmt.Println(mempool.TempStore)
// 	fmt.Printf("💠 transaction %v, added to mempool\n", trx.Amount)
// 	// fmt.Println("Am been called")
// }
