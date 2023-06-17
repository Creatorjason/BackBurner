package main


import (
	"github.com/gin-gonic/gin"
	"github.com/qoinpalhq/HQ_CHAIN/server"
	bc "github.com/qoinpalhq/HQ_CHAIN/blockchain"
	coin "github.com/qoinpalhq/HQ_CHAIN/coins"
	"github.com/qoinpalhq/HQ_CHAIN/kvStore"
)

func main() {
	// testing wallet
	// chain := bc.InitializeChain()
	// mempool := bc.NewMempool()
	cor := bc.NewCoR()
	router := gin.Default()
	db := kvStore.NewDB()
	airDrop := coin.NewAirDrop()
	// new CoR and Trxs here 
	
	sv := server.NewServer(db, router, airDrop, cor)
	sv.RunServer()
}
