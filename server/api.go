package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qoinpalhq/HQ_CHAIN/types"
	"github.com/qoinpalhq/HQ_CHAIN/wallet"
	// "github.com/qoinpalhq/HQ_CHAIN/kvStore"
)

// ENPOINTS
// /api/wallet
// /api/airdrop
// /api/send
// /api/chain
// /api/

func (s *Server) handleGetWalletDetails(c *gin.Context) {
	//  get wallet address from user

	//
}

func (s *Server) handleGenerateNewWallet(c *gin.Context) {
	var (
		owner types.WalletOwner
	)
	err := c.BindJSON(&owner)
	handleBadRequestDueToWrongDataType(err, "Owner", c)
	if owner.Name != "" {
		// generate new wallet
		newWallet := wallet.NewWallet()
		// store in db
		err := s.DB.Write([]byte(owner.Name), newWallet.SerializeWallet())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to write wallet data to db",
			})
		}
		// respond to client with new wallet data
		c.JSON(http.StatusOK, newWallet)
	}

}

func (s *Server) handleReceiveAirdrop(c *gin.Context) {
	var (
		ad types.AirDrop
	)
	err := c.BindJSON(&ad)
	handleBadRequestDueToWrongDataType(err, "AirDrop", c)
	if s.AirDrop.AddWalletAddress(ad.WalletAddr) {
		c.JSON(http.StatusOK, gin.H{
			"message": "wallet address has been whitelisted successfully",
		})
		fmt.Println(len(s.AirDrop.WhiteList))
	}else{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"wallet address already exist or wallet address is invalid",
		})
	}
}

func (s *Server) handleSendCoins(c *gin.Context) {

}

func (s *Server) handleViewBlockchain(c *gin.Context) {

}

func handleBadRequestDueToWrongDataType(err error, data_type string, c *gin.Context) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid data type, wants data type of" + data_type,
		})
	}
}
