package kvStore

import (
	"log"

	badger "github.com/dgraph-io/badger/v3"
)

const DB_PATH = "./kvStore/store"

type DB struct {
	Db *badger.DB
}

func NewDB() *DB {
	opts := badger.DefaultOptions("").WithInMemory(true)
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
	return nil
}

func (db *DB) Read(key []byte) ([]byte, error) {
	return nil, nil
}
