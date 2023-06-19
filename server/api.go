package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	bc "github.com/qoinpalhq/HQ_CHAIN/blockchain"
	"github.com/qoinpalhq/HQ_CHAIN/types"

	// "github.com/qoinpalhq/HQ_CHAIN/utils"
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
	// var (
		// ad types.Wallet
	// )
	addr := c.Param("addr")
	// fmt.Println(len(ad.WalletAddr))
	// if len(ad.WalletAddr) != 40{
		// c.JSON(http.StatusBadRequest, gin.H{
			// "message":"invalid wallet address, address length too short or too long",
		// })
	// }
	userAcctByte, err := s.DB.Read([]byte(addr))
	fmt.Println("Reading...")
	if err != nil {
		log.Printf("unable to read data from db %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to retrieve user account data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": types.Deserialize(userAcctByte)})

	//
}

func (s *Server) handleGenerateNewWallet(c *gin.Context) {

	owner := &types.WalletOwner{}

	err := c.BindJSON(owner)
	// fmt.Println(owner)
	// fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid data type, wants data type of WalletOwner",
		})
		return
	}
	if owner.Name != "" {
		// generate new wallet
		newWallet := wallet.NewWallet()
		// store in db
		writeErr := s.DB.Write([]byte(owner.Name), newWallet.SerializeWallet())
		if writeErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to write wallet data to db",
			})
		}
		// respond to client with new wallet data
		c.JSON(http.StatusOK, gin.H{
			"message":newWallet,
		})
	}else{
		c.JSON(http.StatusInternalServerError, gin.H{"message":"unable to create new wallet"})
	}

}

func (s *Server) handleReceiveAirdrop(c *gin.Context) {
	var (
		ad types.Wallet
	)
	err := c.BindJSON(&ad)
	fmt.Println(len(ad.WalletAddr))
	handleBadRequestDueToWrongDataType(err, "Wallet", c)
	if s.AirDrop.AddWalletAddress(ad.WalletAddr, s.DB) {
		c.JSON(http.StatusOK, gin.H{
			"message": "wallet address has been whitelisted successfully",
		})
		fmt.Println(s.AirDrop.AddrCount)
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "wallet address already exist or wallet address is invalid",
		})
		return
	}
}

func (s *Server) handleGetBalanceOfWhitelistedAddresses(c *gin.Context) {
	var usersAcct []*types.UserAccount
	// get all addresses from whitelist
	for _, addr := range s.AirDrop.WhiteList {
		// get user account data from db
		userAcctByte, err := s.DB.Read([]byte(addr))
		fmt.Println("Reading...")
		if err != nil {
			log.Printf("unable to read data from db %v\n", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to retrieve user account data"})
			return
		}
		userAcct := types.Deserialize(userAcctByte)
		usersAcct = append(usersAcct, userAcct)
	}
	fmt.Println(s.AirDrop.WhiteList)
	if usersAcct == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "message account list is empty, no address has been whitelisted yet",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": usersAcct})

}

// TODO : Bug-: Transactions double in new block
func (s *Server) handleSendCoins(c *gin.Context) {
	var (
		transaction types.Transaction
	)
	err := c.BindJSON(&transaction)
	handleBadRequestDueToWrongDataType(err, "Transaction", c)

	// test transaction creation
	trx := bc.CreateTransaction(transaction.Sender, transaction.Receiver, transaction.Amount)
	// check balance of user
	senderAcct, err := s.DB.Read([]byte(transaction.Sender))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to get user account from db",
		})
		return
	}

	senderAcctStruct := types.Deserialize(senderAcct)
	senderBal := senderAcctStruct.Balance
	if uint(transaction.Amount) > senderBal {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Insufficient balance, you can't send more that what you have",
		})
		return
	}
	// update sender balance
	senderBal -= uint(transaction.Amount)
	senderAcctStruct.Balance = senderBal
	// updatedSenderAcct := senderAcctStruct.Serialize()
	// write to db .... works for now ... Change later
	err = s.DB.Write([]byte(transaction.Sender), senderAcctStruct.Serialize())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("failed to update sender account: %v", err),
		})
		return
	}

	// update receiver's balance
	recvAcct, err := s.DB.Read([]byte(transaction.Receiver))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to get user account from db",
		})
		return
	}

	recvAcctStruct := types.Deserialize(recvAcct)
	recvAcctStruct.Balance += uint(transaction.Amount)

	err = s.DB.Write([]byte(transaction.Receiver), recvAcctStruct.Serialize())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("failed to update receiver account: %v", err),
		})
		return
	}

	// stored created transaction in mempool
	s.Cor.Mempool.AddTransactionToMempool(trx)
	s.Cor.Mempool.Execute(s.Cor.Trxs)
	s.Cor.Mempool.SetNext(s.Cor.Blockchain)

	c.JSON(http.StatusOK, gin.H{
		"message":fmt.Sprintf("successfully sent %v from %v to %v", transaction.Amount, transaction.Sender, transaction.Receiver),
	})
	// "4e6b3bd43d3e70ff7a258982bb090a2dc50a7d09"
	// "22211180384ae191718d5d725c2ecda3b130e32e"

}

func (s *Server) handleViewBlockchain(c *gin.Context) {
	if len(s.Cor.Blockchain.Chain) > 0{
		c.JSON(http.StatusOK, s.Cor.Blockchain)
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"message":"No Blockchain to display",
	})
}

func handleBadRequestDueToWrongDataType(err error, data_type string, c *gin.Context) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid data type, wants data type of" + data_type,
		})
	}
}
