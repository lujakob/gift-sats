package wallet

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model

	TipId          uint   `gorm:"not null"`
	LnbitsWalletId string `gorm:"not null"`
	LnbitsUserId   string `gorm:"not null"`
	AdminKey       string `gorm:"not null"`
}
