package coins
// THIS CODE IS BROKEN
import (
	"fmt"
	// "log"

	db "github.com/qoinpalhq/HQ_CHAIN/kvStore"
	"github.com/qoinpalhq/HQ_CHAIN/types"
)

// Airdrop basically shares the an amount of coins amongst whitelisted wallet addresses
const COINS_PER_ADDRESS = 500
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
		MaxAddrCount: 0,
	}
}

func (a *Airdrop) AddWalletAddress(wallet_addr string, db *db.DB) bool {
	// isPresent := a.CheckIfWalletAddressIsWhitelisted(wallet_addr)
	// if len(wallet_addr) != 40 {
		// log.Println("invalid wallet address, length is too long or too short")
		// return false
	// }
	// if !isPresent {
		// a.WhiteList = append(a.WhiteList, wallet_addr)
		// a.AddrCount += 1
		// log.Println(wallet_addr, "successfully whitelisted")
		// if a.AddrCount > a.MaxAddrCount{
		// }
		// return true
	// }
	// log.Println("unable to add wallet address, already whitelisted")
	if len(wallet_addr) == 40{
	a.SendCoinToWalletAddresses(wallet_addr,db)
	return true
	}
	return false
}

func (a *Airdrop) SendCoinToWalletAddresses(wallet_addr string, db *db.DB) error {
	//  if the address count is at max yet
	// if len(a.WhiteList) > a.MaxAddrCount && !a.IsExhausted {

		// for _, addr := range a.WhiteList {
		//  create new acount
		// newAccount := types.NewUserAccount(addr, a.Balances[addr])
		// pesist to db
		// err := db.Write([]byte(addr), newAccount.Serialize())
		// if err != nil {
		// return fmt.Errorf("failed to write account to database: %w", err)
		// }
		// }
		coinsPerAddress := COINS_PER_ADDRESS

		// for _, addr := range a.WhiteList {
		a.Balances[wallet_addr] = uint(coinsPerAddress)
		// }
		// Decrement a.ToShare by the total coins distributed
		totalCoinsDistributed := uint(coinsPerAddress) 
		a.ToShare -= totalCoinsDistributed
		// if a.ToShare == 0 {ß
			// a.IsExhausted = true
		// }
		newAccount := types.NewUserAccount(wallet_addr, a.Balances[wallet_addr])
		// pesist to db
		err := db.Write([]byte(wallet_addr), newAccount.Serialize())
		if err != nil {
			return fmt.Errorf("failed to write account to database: %w", err)
		}
		return nil
	// }
	// return fmt.Errorf("cannot send coins now, whitelist not yet filled")
}

func (a *Airdrop) CheckIfWalletAddressIsWhitelisted(wallet_addr string) bool {
	for _, w := range a.WhiteList {
		if w == wallet_addr {
			return true
		}
	}
	return false
}
