package interfaces

import "github.com/lujakob/gift-sats/models"

type IUserStore interface {
	GetAll() ([]models.User, int64, error)
	GetByEmail(string) (*models.User, error)
	Create(*models.User) error
}
