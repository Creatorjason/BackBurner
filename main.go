package main


import (
	"github.com/gin-gonic/gin"
	"github.com/qoinpalhq/HQ_CHAIN/server"
	"github.com/qoinpalhq/HQ_CHAIN/kvStore"
)

func main() {
	// testing wallet
	router := gin.Default()
	db := kvStore.NewDB()
	sv := server.NewServer(db, router)
	sv.RunServer()


}
