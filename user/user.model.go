package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not nul"`
}
