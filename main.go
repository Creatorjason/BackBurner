package main


import (
	"github.com/gin-gonic/gin"
	"github.com/qoinpalhq/HQ_CHAIN/server"
	coin "github.com/qoinpalhq/HQ_CHAIN/coins"
	"github.com/qoinpalhq/HQ_CHAIN/kvStore"
)

func main() {
	// testing wallet
	router := gin.Default()
	db := kvStore.NewDB()
	airDrop := coin.NewAirDrop()
	// new CoR and Trxs here 
	sv := server.NewServer(db, router, airDrop)
	sv.RunServer()
}
