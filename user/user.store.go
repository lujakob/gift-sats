package user

import (
	"errors"

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

func (us *UserStore) GetAll() ([]User, int64, error) {
	var users []User
	var count int64
	if err := us.db.Find(&users).Count(&count).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	return users, count, nil
}

func (us *UserStore) GetByUsername(username string) (*User, error) {
	var m User
	if err := us.db.Where(&User{Username: username}).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) Create(u *User) (err error) {
	return us.db.Create(u).Error
}
