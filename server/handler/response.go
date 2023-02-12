package handler

import (
	"github.com/lujakob/gift-sats/tip"
	"github.com/lujakob/gift-sats/user"
	"github.com/lujakob/gift-sats/wallet"
)

type userResponse struct {
	User struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"user"`
}

func newUserResponse(u *user.User) *userResponse {
	r := new(userResponse)
	r.User.ID = u.ID
	r.User.Username = u.Username
	r.User.Email = u.Email
	return r
}

type userListResponse struct {
	Users      []*userResponse `json:"users"`
	UsersCount int64           `json:"usersCount"`
}

func newUserListResponse(users []user.User, count int64) *userListResponse {
	r := new(userListResponse)
	ur := new(userResponse)
	for _, u := range users {

		ur.User.Username = u.Username
		ur.User.Email = u.Email

		r.Users = append(r.Users, ur)
	}
	r.UsersCount = count
	return r
}

type tipResponse struct {
	Tip struct {
		ID     uint `json:"id"`
		Tipper struct {
			Username string `json:"username"`
			ID       uint   `json:"id"`
		} `json:"tipper"`
		Amount int `json:"amount"`
		Fee    int `json:"fee"`
	} `json:"tip"`
}

func newTipResponse(t *tip.Tip) *tipResponse {
	r := new(tipResponse)

	r.Tip.Amount = t.Amount
	r.Tip.Fee = t.Fee
	r.Tip.Tipper.Username = t.Tipper.Username
	r.Tip.Tipper.ID = t.Tipper.ID

	return r
}

type tipListResponse struct {
	Tips      []*tipResponse `json:"tips"`
	TipsCount int64          `json:"tipsCount"`
}

func newTipListResponse(tips []tip.Tip, count int64) *tipListResponse {
	r := new(tipListResponse)
	tr := new(tipResponse)

	for _, t := range tips {
		tr.Tip.ID = t.ID
		tr.Tip.Amount = t.Amount
		tr.Tip.Fee = t.Fee
		tr.Tip.Tipper.Username = t.Tipper.Username
		tr.Tip.Tipper.ID = t.Tipper.ID

		r.Tips = append(r.Tips, tr)
	}

	r.TipsCount = count
	return r
}

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

func newWalletListResponse(wallets []wallet.Wallet, count int64) *walletListResponse {
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
