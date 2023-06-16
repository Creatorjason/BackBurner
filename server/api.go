package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qoinpalhq/HQ_CHAIN/types"
	"github.com/qoinpalhq/HQ_CHAIN/wallet"
	// "github.com/qoinpalhq/HQ_CHAIN/wallet"
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
	if s.AirDrop.AddWalletAddress(ad.WalletAddr, s.DB) {
		c.JSON(http.StatusOK, gin.H{
			"message": "wallet address has been whitelisted successfully",
		})
		fmt.Println(s.AirDrop.AddrCount)
		return 
	}else{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"wallet address already exist or wallet address is invalid",
		})
		return
	}
}

func (s *Server) handleGetBalanceOfAddresses(c *gin.Context){
	var usersAcct []*types.UserAccount
	// get all addresses from whitelist
	for _, addr := range s.AirDrop.WhiteList{
		// get user account data from db 
		userAcctByte, err := s.DB.Read([]byte(addr))
		fmt.Println("Reading...")
		if err != nil{
			log.Printf("unable to read data from db %v\n", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve user account data"})
			return
		}
		userAcct := types.Deserialize(userAcctByte)
		usersAcct = append(usersAcct, userAcct)
	}
	fmt.Println(s.AirDrop.WhiteList)
	c.JSON(http.StatusOK, gin.H{"users": usersAcct})
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
		return
	}
}
