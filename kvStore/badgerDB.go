package kvStore

import (
	"fmt"
	"log"

	badger "github.com/dgraph-io/badger/v3"
)

const DB_PATH = "./kvStore/store"

type DB struct {
	Db *badger.DB
}

func NewDB() *DB {
	// opts := badger.DefaultOptions("").WithInMemory(true)
	opts := badger.DefaultOptions(DB_PATH)
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatalf("unable to open badger db: %v\n", err.Error())
	}
	return &DB{
		Db: db,
	}
}

func (db *DB) Write(key, value []byte) error {
	err := db.Db.Update(func(txn *badger.Txn) error {
		err := txn.Set(key, value)
		return err
	})
	if err != nil {
		log.Printf("update transaction failed: %v\n", err.Error())
	}
	fmt.Printf("key: %x has been stored\n", key)
	return nil
}

func (db *DB) Read(key []byte) ([]byte, error) {
	var value []byte
	err := db.Db.View(func(txn *badger.Txn)error{
		item, err := txn.Get(key)
		if err != nil{
			log.Printf("unable to get key: %v\n", err.Error())
		}
		err = item.Value(func(val []byte) error{
			value = append([]byte{}, val...)
			return nil
		})
		return err
	})
	if err != nil{
		return nil, err
	}
	return value, nil
}
