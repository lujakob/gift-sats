package handler

import (
	"github.com/lujakob/gifting-sats/user"
)

type Handler struct {
	userStore user.IUserStore
}

func NewHandler(us user.IUserStore) *Handler {
	return &Handler{
		userStore: us,
	}
}
