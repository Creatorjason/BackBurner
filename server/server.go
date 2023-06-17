package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	bc "github.com/qoinpalhq/HQ_CHAIN/blockchain"
	coin "github.com/qoinpalhq/HQ_CHAIN/coins"
	kv "github.com/qoinpalhq/HQ_CHAIN/kvStore"
)

type Server struct {
	DB     *kv.DB
	Router *gin.Engine
	AirDrop *coin.Airdrop
	Cor *bc.CoR
}

func NewServer(db *kv.DB, router *gin.Engine, ad *coin.Airdrop, cor *bc.CoR) *Server {
	return &Server{
		DB: db,
		Router: router,
		AirDrop : ad,
		Cor : cor,
	}
}

func (s *Server) RunServer() {
	// register endpoints here
	s.Router.GET("/api/wallet", s.handleGetWalletDetails)
	s.Router.POST("/api/wallet", s.handleGenerateNewWallet)
	s.Router.POST("/api/send", s.handleSendCoins)
	s.Router.GET("/", s.handleViewBlockchain)
	s.Router.POST("/api/airdrop", s.handleReceiveAirdrop)
	// Work on this later, wrong enpoint handler
	s.Router.GET("/api/airdrop", s.handleGetBalanceOfWhitelistedAddresses)
	// Adding CORS
	corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{"https://master--hq-chain-ui.netlify.app/"}
	corsConfig.AllowAllOrigins = true
	// Options method for react js
	corsConfig.AddAllowMethods("OPTIONS")
	corsConfig.AllowCredentials = true


	s.Router.Use(cors.New(corsConfig))

	s.Router.Run()
}
