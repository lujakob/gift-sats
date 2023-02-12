package stores

import (
	"errors"

	"github.com/lujakob/gift-sats/models"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) GetAll() ([]models.User, int64, error) {
	var users []models.User
	var count int64
	if err := us.db.Find(&users).Count(&count).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	return users, count, nil
}

func (us *UserStore) GetByEmail(email string) (*models.User, error) {
	var m models.User
	if err := us.db.Where(&models.User{Email: email}).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) Create(u *models.User) (err error) {
	return us.db.Create(u).Error
}
