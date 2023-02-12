package models

import (
	"gorm.io/gorm"
)

type Tip struct {
	gorm.Model
	Amount   int `gorm:"not nul"`
	Fee      int `gorm:"not nul"`
	Tipper   User
	TipperID uint `gorm:"not nul"`
}
