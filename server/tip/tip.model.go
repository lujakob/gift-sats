package tip

import (
	"gorm.io/gorm"

	"github.com/lujakob/gift-sats/user"
)

type Tip struct {
	gorm.Model
	Amount   int `gorm:"not nul"`
	Fee      int `gorm:"not nul"`
	Tipper   user.User
	TipperID uint `gorm:"not nul"`
}
