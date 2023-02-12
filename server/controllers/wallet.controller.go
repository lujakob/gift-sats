package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lujakob/gift-sats/interfaces"
	"github.com/lujakob/gift-sats/models"
	"github.com/lujakob/gift-sats/utils"
)

type walletResponse struct {
	Wallet struct {
		ID             uint   `json:"id"`
		TipId          uint   `json:"tip_id"`
		LnbitsWalletId string `json:"lnbits_wallet_id"`
		LnbitsUserId   string `json:"lnbits_user_id"`
	} `json:"tip"`
}

type walletListResponse struct {
	Wallets     []*walletResponse `json:"wallets"`
	WalletCount int64             `json:"walletCount"`
}

func newWalletListResponse(wallets []models.Wallet, count int64) *walletListResponse {
	r := new(walletListResponse)
	tr := new(walletResponse)

	for _, t := range wallets {
		tr.Wallet.ID = t.ID
		tr.Wallet.TipId = t.TipId
		tr.Wallet.LnbitsUserId = t.LnbitsUserId
		tr.Wallet.LnbitsWalletId = t.LnbitsWalletId

		r.Wallets = append(r.Wallets, tr)
	}

	r.WalletCount = count
	return r
}

type WalletController struct {
	walletStore interfaces.IWalletStore
}

func NewWalletController(walletStore interfaces.IWalletStore) *WalletController {
	return &WalletController{
		walletStore: walletStore,
	}
}

func (w *WalletController) List(c echo.Context) error {
	list, count, error := w.walletStore.GetAll()
	if error != nil {
		fmt.Println(error)
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	return c.JSON(http.StatusOK, newWalletListResponse(list, count))
}
