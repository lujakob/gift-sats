package handler

import (
	"github.com/lujakob/gift-sats/tip"
	"github.com/lujakob/gift-sats/user"
	"github.com/lujakob/gift-sats/wallet"
)

type Handler struct {
	userStore   user.IUserStore
	tipStore    tip.ITipStore
	walletStore wallet.IWalletStore
}

func NewHandler(us user.IUserStore, ts tip.ITipStore, ws wallet.IWalletStore) *Handler {
	return &Handler{
		userStore:   us,
		tipStore:    ts,
		walletStore: ws,
	}
}
