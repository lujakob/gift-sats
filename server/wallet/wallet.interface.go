package wallet

type IWalletStore interface {
	GetAll() ([]Wallet, int64, error)
	Create(*Wallet) error
}
