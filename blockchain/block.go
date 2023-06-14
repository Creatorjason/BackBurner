package blockchain

import (
	"time"
)

type (
	Block struct {
		BlockHeader  *BlockHeader
		Transactions []*Transaction
	}

	Blockchain struct {
		Blocks []*Block
	}
	Transaction struct {
		ID     []byte
		TxIn   []TxInputs
		TxOut  []TxOutputs
	}
	BlockHeader struct {
		MerkleRoot []byte
		PrevHash   []byte
		Hash       []byte
		Timestamp  time.Time
		Height     int
	}
	TxInputs struct {
		TXID []byte
		Vout int
		// sig of sender
		Sig []byte
		// pub key/ address of receiver
		PubKey []byte
	}
	TxOutputs struct {
		Value int
		Address []byte
	}
)

func (bl *Block) DeriveBlockHash() []byte {
	return nil
}
func CreateBlock(trx []*Transaction, prevHash []byte, height int) *Block {
	return nil
}

func InitializeChain() *Blockchain {
	return nil
}

func CreateTransaction() *Transaction {
	return nil
}

func GenesisBlock() *Block {
	return nil
}

func (bl *Blockchain) AddBlockHeader() *BlockHeader {
	return nil
}
