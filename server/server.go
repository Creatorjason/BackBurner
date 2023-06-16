package server

import (
	"github.com/gin-gonic/gin"
	kv "github.com/qoinpalhq/HQ_CHAIN/kvStore"
)

type Server struct {
	DB     *kv.DB
	Router *gin.Engine
}

func NewServer(db *kv.DB, router *gin.Engine) *Server {
	return &Server{
		DB: db,
		Router: router,
	}
}

func (s *Server) RunServer() {
	// register endpoints here
	s.Router.GET("/api/wallet", s.handleGetWalletDetails)
	s.Router.POST("/api/wallet", s.handleGenerateNewWallet)
	s.Router.GET("/api/send", s.handleSendCoins)
	s.Router.GET("/api/chain", s.handleViewBlockchain)
	s.Router.POST("/api/airdrop", s.handleReceiveAirdrop)
	s.Router.Run()

}
