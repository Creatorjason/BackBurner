package coins

import (
	"fmt"
	"log"

	db "github.com/qoinpalhq/HQ_CHAIN/kvStore"
	"github.com/qoinpalhq/HQ_CHAIN/types"
	"github.com/qoinpalhq/HQ_CHAIN/utils"
)

// Airdrop basically shares the an amount of coins amongst whitelisted wallet addresses

type Airdrop struct {
	WhiteList    []string
	MaxAddrCount int
	AddrCount    int
	Balances     map[string]uint
	// represents the amount of coins to be airdropped
	ToShare     uint
	IsExhausted bool
}

func NewAirDrop() *Airdrop {
	bl := make(map[string]uint)
	fmt.Println("Created new airdrop....")
	return &Airdrop{
		Balances:     bl,
		ToShare:      TOTAL_SUPPLY - 10000,
		MaxAddrCount: 10,
	}
}

func (a *Airdrop) AddWalletAddress(wallet_addr string) {
	isPresent := a.CheckIfWalletAddressIsWhitelisted(wallet_addr)
	if !isPresent {
		a.WhiteList = append(a.WhiteList, wallet_addr)
		a.AddrCount += 1
		log.Println(wallet_addr, "successfully whitelisted")
	}
	// persist to db
	log.Println("unable to add wallet address, already whitelisted")
}

func (a *Airdrop) SendCoinToWalletAddresses(db *db.DB) error {
	//  if the address count is at max yet
	if len(a.WhiteList) == a.MaxAddrCount && !a.IsExhausted{
		coinsPerAddress := a.ToShare / uint(len(a.WhiteList))
		for _, addr := range a.WhiteList {
			a.Balances[addr] = coinsPerAddress
		}

		// Decrement a.ToShare by the total coins distributed
		totalCoinsDistributed := coinsPerAddress * uint(len(a.WhiteList))
		a.ToShare -= totalCoinsDistributed
		if a.ToShare == 0{
			a.IsExhausted = true
		}
		for _, addr := range a.WhiteList {
			//  create new acount
			newAccount := types.NewUserAccount(addr, a.Balances[addr])
			// pesist to db
			err := db.Write([]byte(addr), utils.Serialize(newAccount))
			if err != nil {
				return fmt.Errorf("failed to write account to database: %w", err)
			}
		}
		return nil
	}
	return fmt.Errorf("cannot send coins now, whitelist not yet filled")
}

func (a *Airdrop) CheckIfWalletAddressIsWhitelisted(wallet_addr string) bool {
	for _, w := range a.WhiteList {
		if w == wallet_addr {
			return true
		}
	}
	return false
}
