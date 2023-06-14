package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"time"
)

type (
	Block struct {
		BlockHeader  *BlockHeader
		Transactions []Transaction
	}

	Blockchain struct {
		Blocks []*Block
	}

	BlockHeader struct {
		MerkleRoot []byte
		PrevHash   []byte
		Hash       []byte
		Timestamp  time.Time
		Height     int
	}
)

// func (bl *Block) DeriveBlockHash() []byte {

// 	return nil
// }
func CreateBlock(trx []Transaction, prevHash []byte, height int) *Block {
	mRoot := GetMerkleRoot(trx)
	bHeader := CreateBlockHeader(prevHash, mRoot, time.Now(), height)
	blockHash := sha256.Sum256(bHeader.SerializeBH())
	bHeader.Hash = blockHash[:]

	block := &Block{
		Transactions: trx,
		BlockHeader:  bHeader,
	}

	return block
}
func CreateBlockHeader(prevHash, merkleRoot []byte, time_stamp time.Time, height int) *BlockHeader {
	bHeader := &BlockHeader{
		MerkleRoot: merkleRoot,
		PrevHash:   prevHash,
		Timestamp:  time_stamp,
		Height:     height,
	}
	return bHeader
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

func (bh *BlockHeader) SerializeBH() []byte {
	buff := new(bytes.Buffer)
	err := gob.NewEncoder(buff).Encode(bh)
	if err != nil {
		panic(err)
	}
	return buff.Bytes()
}
