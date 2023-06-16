package server

import (
	"github.com/gin-gonic/gin"
	kv "github.com/qoinpalhq/HQ_CHAIN/kvStore"
	coin "github.com/qoinpalhq/HQ_CHAIN/coins"

)

type Server struct {
	DB     *kv.DB
	Router *gin.Engine
	AirDrop *coin.Airdrop
}

func NewServer(db *kv.DB, router *gin.Engine, ad *coin.Airdrop) *Server {
	return &Server{
		DB: db,
		Router: router,
		AirDrop : ad,
	}
}

func (s *Server) RunServer() {
	// register endpoints here
	s.Router.GET("/api/wallet", s.handleGetWalletDetails)
	s.Router.POST("/api/wallet", s.handleGenerateNewWallet)
	s.Router.GET("/api/send", s.handleSendCoins)
	s.Router.GET("/api/chain", s.handleViewBlockchain)
	s.Router.POST("/api/airdrop", s.handleReceiveAirdrop)
	s.Router.GET("/api/airdrop", s.handleGetBalanceOfAddresses)
	s.Router.Run()

}
