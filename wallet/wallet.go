package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"log"
	"golang.org/x/crypto/ripemd160"
	"github.com/qoinpalhq/HQ_CHAIN/utils"
)

// !!! This is a very basic wallet JBOK, only meant for experimental purposes
// !!! DO NOT SEND REAL CRYPTOCURRENCY TO THIS WALLET ADDRESS

type Wallet struct {
	Sk   string `json:"private_key"`
	Pk   string           `json:"public_key"`
	Addr string           `json:"wallet_address"`
}

// WALLETS don't store cryptocurrencies 


func NewWallet() *Wallet {
	sk, pk := newKeyPair()
	return &Wallet{
		Sk:   utils.PrivToHexString(&sk),
		Pk:   utils.ToHexString(pk),
		Addr: generateAddress(pk),
	}
}
func newKeyPair() (ecdsa.PrivateKey, []byte) {
	// define curve
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Fatalf("failed to generate new private key: %v\n", err.Error())
	}
	pubKey := append(privateKey.X.Bytes(), privateKey.Y.Bytes()...)
	return *privateKey, pubKey

}

func generateAddress(pk []byte) string {
	hash256 := sha256.Sum256(pk)
	hash160 := ripemd160.New()
	_, err := hash160.Write(hash256[:])
	if err != nil {
		log.Printf("failed to generate wallet address: %v\n", err.Error())
	}
	addrByte := hash160.Sum(nil)
	address := hex.EncodeToString(addrByte)
	return address
}

func (w *Wallet) SerializeWallet() []byte{
	buff := new(bytes.Buffer)
	err := gob.NewEncoder(buff).Encode(w)
	if err != nil{
		log.Printf("failed to serialize wallet data: %v\n", err.Error())
	}
	return buff.Bytes()
}

