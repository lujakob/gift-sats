package stores

import (
	"errors"

	"github.com/lujakob/gift-sats/models"
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

func (us *WalletStore) GetAll() ([]models.Wallet, int64, error) {
	var wallets []models.Wallet
	var count int64
	if err := us.db.Find(&wallets).Count(&count).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	return wallets, count, nil
}

func (ws *WalletStore) Create(u *models.Wallet) (err error) {
	return ws.db.Create(u).Error
}
