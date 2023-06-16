package utils

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/gob"
	"encoding/hex"
	"log"
	"math/big"
	"strings"
	// "crypto/sha256"
)

func PrivToHexString(p *ecdsa.PrivateKey) string {
	pK := p.D.Bytes()
	hexString := hex.EncodeToString(pK)
	return strings.ToUpper(hexString)
}
func FromHexStringToBytes(hexString string) []byte {
	bytesVal, err := hex.DecodeString(hexString)
	if err != nil {
		log.Fatal(err)
	}
	return bytesVal
}

func ECDSAfromHex(hexString string) *ecdsa.PrivateKey {
	privKey := new(ecdsa.PrivateKey)
	privKey.D, _ = new(big.Int).SetString(hexString, 16)
	privKey.PublicKey.Curve = elliptic.P256()
	privKey.PublicKey.X, privKey.PublicKey.Y = privKey.PublicKey.Curve.ScalarBaseMult(privKey.D.Bytes())
	return privKey

}
func ToHexString(data []byte) string {
	hexString := hex.EncodeToString(data)
	return strings.ToUpper(hexString)
}

func Serialize(data interface{}) []byte {
	buff := new(bytes.Buffer)
	err := gob.NewEncoder(buff).Encode(data)
	if err != nil {
		log.Printf("failed to serialize payload :%v\n", err.Error())
	}
	return buff.Bytes()
}
