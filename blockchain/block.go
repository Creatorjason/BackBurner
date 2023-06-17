package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"time"
)

type (
	Block struct {
		BlockHeader  *BlockHeader  `json:"block_header"`
		Transactions []Transaction `json:"transactions"`
		// Next         MoveTransaction
	}

	BlockHeader struct {
		MerkleRoot []byte    `json:"merkle_root"`
		PrevHash   []byte    `json:"prev_block_hash"`
		Hash       []byte    `json:"block_hash"`
		Timestamp  time.Time `json:"block_timestamp"`
		Height     int       `json:"block_height"`
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

func CreateGenesisBlock() *Block {
	// for testing purpose
	trx1 := CreateTransaction("Jason","Qoinpal", 100)
	trx2 := CreateTransaction("Kendrick","Dayo", 200)
	// testing...
	return CreateBlock([]Transaction{trx1, trx2}, nil, 0)
}

func GenesisBlock() *Block {
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

// implement MoveTransaction interface

// func (bl *Block) Execute(trx *Trxs) {
// 	if trx.AddedToBlock{

// 	}	
// }

// func (bl *Block) SetNext(next MoveTransaction) {

// }
