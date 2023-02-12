package wallet

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model

	TipId          int    `gorm:"not null"`
	LnbitsWalletId int    `gorm:"not null"`
	LnbitsUserId   int    `gorm:"not null"`
	AdminKey       string `gorm:"not null"`
}
