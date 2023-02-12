package wallet

import (
	"gorm.io/gorm"
)

type WalletStore struct {
	db *gorm.DB
}

func NewWalletStore(db *gorm.DB) *WalletStore {
	return &WalletStore{
		db: db,
	}
}

func (ws *WalletStore) Create(u *Wallet) (err error) {
	return ws.db.Create(u).Error
}
