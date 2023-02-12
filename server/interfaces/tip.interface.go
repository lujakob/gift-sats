package interfaces

import "github.com/lujakob/gift-sats/models"

type ITipStore interface {
	GetAll() ([]models.Tip, int64, error)
	Create(*models.Tip) error
}
