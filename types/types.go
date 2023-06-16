package types

import (
	"bytes"
	"encoding/gob"
	"log"

	bl "github.com/qoinpalhq/HQ_CHAIN/blockchain"
)

type (
	WalletOwner struct {
		Name string `json:"name"`
	}
	UserAccount struct {
		WalletAddr string `json:"user_wallet_address"`
		Balance    uint   `json:"account_balance"`
	}
	AirDrop struct {
		WalletAddr string `json:"wallet_address"`
	}
	// chain of responsibility
	// CoR struct holds all the components that will be involved in the
	// processing of an array of transaction
	CoR struct {
		Mempool    *bl.Mempool
		Block      *bl.Block
		Blockchain *bl.Blockchain
	}
)

func NewUserAccount(wallet_addr string, balance uint) *UserAccount {
	return &UserAccount{
		WalletAddr: wallet_addr,
		Balance:    balance,
	}
}


func NewCoR(mp *bl.Mempool, bk *bl.Block, bc *bl.Blockchain) *CoR{
	return &CoR{
		Mempool : mp,
		Block : bk,
		Blockchain : bc,
	}
}

// move later
func (ua *UserAccount) Serialize() []byte {
	buff := new(bytes.Buffer)
	err := gob.NewEncoder(buff).Encode(ua)
	if err != nil {
		log.Printf("failed to serialize user account: %v\n", err.Error())
	}
	return buff.Bytes()
}

func Deserialize(payload []byte) *UserAccount {
	ua := &UserAccount{}
	err := gob.NewDecoder(bytes.NewReader(payload)).Decode(ua)
	if err != nil {
		log.Printf("failed to deserialize payload: %v\n", err.Error())
	}
	return ua
}

//
