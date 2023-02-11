package handler

import (
	"github.com/lujakob/gift-sats/user"
)

type Handler struct {
	userStore user.IUserStore
}

func NewHandler(us user.IUserStore) *Handler {
	return &Handler{
		userStore: us,
	}
}
