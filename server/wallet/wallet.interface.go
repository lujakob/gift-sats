package wallet

type IWalletStore interface {
	Create(*Wallet) error
}
