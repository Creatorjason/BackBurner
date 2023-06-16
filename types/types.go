package types

type (
	WalletOwner struct {
		Name string `json:"name"`
	}
	UserAccount struct {
		WalletAddr string `json:"user_wallet_address"`
		Balance    uint   `json:"account_balance"`
	}
	AirDrop struct{
		WalletAddr string `json:"wallet_address"`
	}
)


func NewUserAccount(wallet_addr string, balance uint) *UserAccount{
	return &UserAccount{
		WalletAddr : wallet_addr,
		Balance : balance,
	}
}