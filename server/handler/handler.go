package handler

import (
	"github.com/lujakob/gift-sats/tip"
	"github.com/lujakob/gift-sats/user"
)

type Handler struct {
	userStore user.IUserStore
	tipStore  tip.ITipStore
}

func NewHandler(us user.IUserStore, ts tip.ITipStore) *Handler {
	return &Handler{
		userStore: us,
		tipStore:  ts,
	}
}
