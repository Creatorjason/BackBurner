package server

import (
	"github.com/gin-gonic/gin"
	"github.com/qoinpalhq/HQ_CHAIN/types"
	"net/http"
	"github.com/qoinpalhq/HQ_CHAIN/wallet"
	// "github.com/qoinpalhq/HQ_CHAIN/kvStore"
)

// ENPOINTS
// /api/wallet
// /api/airdrop
// /api/send
// /api/chain
// /api/


func (s *Server) handleGetWalletDetails(c *gin.Context){
	//  get wallet address from user
	
	// 
}

func (s *Server) handleGenerateNewWallet(c *gin.Context){
	var (
		owner *types.WalletOwner
	)
	err := c.BindJSON(&owner)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid data type, wants data type of Owner",
		})
	}
	if owner.Name != ""{
		// generate new wallet
		newWallet := wallet.NewWallet()
		// store in db
		err := s.DB.Write([]byte(owner.Name), newWallet.SerializeWallet())
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":"failed to write wallet data to db",
			})
		}
		// respond to client with new wallet data
		c.JSON(http.StatusOK, newWallet)
	}

}

func (s *Server) handleReceiveAirdrop(c *gin.Context){

}

func (s *Server) handleSendCoins(c *gin.Context){

}

func (s *Server) handleViewBlockchain(c *gin.Context){

}