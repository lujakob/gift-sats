package interfaces

import "github.com/lujakob/gift-sats/models"

type IWalletStore interface {
	GetAll() ([]models.Wallet, int64, error)
	Create(*models.Wallet) error
}
