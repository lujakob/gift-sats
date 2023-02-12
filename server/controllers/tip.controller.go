package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lujakob/gift-sats/config"
	"github.com/lujakob/gift-sats/interfaces"
	"github.com/lujakob/gift-sats/models"
	"github.com/lujakob/gift-sats/utils"
)

type tipCreateRequest struct {
	Tip struct {
		Amount   int  `json:"amount" validate:"required"`
		Fee      int  `json:"fee" validate:"required"`
		TipperID uint `json:"tipper_id" validate:"required"`
	} `json:"tip"`
}

func (r *tipCreateRequest) bind(c echo.Context, t *models.Tip) error {
	//validate

	if err := c.Bind(r); err != nil {
		return err
	}
	//fmt.Printf("%v", *r)

	t.Amount = r.Tip.Amount
	t.Fee = r.Tip.Fee
	t.TipperID = r.Tip.TipperID

	return nil
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

type CreateLnbitsUserRequest struct {
	AdminId    string `json:"admin_id"`
	UserName   string `json:"user_name"`
	WalletName string `json:"wallet_name"`
}

type CreateLnbitsUserResponse struct {
	Wallets []struct {
		WalletId string `json:"id"`
		Admin    string `json:"admin"`
		User     string `json:"user"`
		AdminKey string `json:"adminkey"`
	} `json:"wallets"`
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

func newTipResponse(t *models.Tip) *tipResponse {
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

func newTipListResponse(tips []models.Tip, count int64) *tipListResponse {
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

type TipController struct {
	tipStore    interfaces.ITipStore
	walletStore interfaces.IWalletStore
}

func NewTipController(tipStore interfaces.ITipStore, walletStore interfaces.IWalletStore) *TipController {
	return &TipController{
		tipStore:    tipStore,
		walletStore: walletStore,
	}
}

func (t *TipController) List(c echo.Context) error {
	list, count, error := t.tipStore.GetAll()
	if error != nil {
		fmt.Println(error)
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	return c.JSON(http.StatusOK, newTipListResponse(list, count))
}

func (t *TipController) Create(c echo.Context) error {
	config := config.GetConfig()

	// @Todo throw error for missing configs
	// if !config.LNBITS_API_KEY || !config.LNBITS_URL || !config.LNBITS_USER_ID

	// @Todo: verify "Only tips with positive, whole amounts are allowed"

	// @Todo: calculateFee

	// first create and store tip
	var tip models.Tip
	req := &tipCreateRequest{}
	if err := req.bind(c, &tip); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	if err := t.tipStore.Create(&tip); err != nil {
		fmt.Println(err)
	}

	fmt.Println(PrettyPrint(t))

	// fetch LnbitsWallet

	createLnbitsUserRequest := &CreateLnbitsUserRequest{
		AdminId:    config.LNBITS_USER_ID,
		UserName:   fmt.Sprint("tip-", tip.ID),
		WalletName: fmt.Sprint("tip-", tip.ID),
	}

	jsonData, _ := json.Marshal(createLnbitsUserRequest)

	fmt.Println(string(jsonData))

	url := fmt.Sprint(config.LNBITS_URL, "/usermanager/api/v1/users")

	walletReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	walletReq.Header.Set("Content-Type", "application/json")
	walletReq.Header.Set("Accept", "application/json")
	walletReq.Header.Set("X-Api-Key", config.LNBITS_API_KEY)

	client := &http.Client{}
	response, error := client.Do(walletReq)
	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

	var result CreateLnbitsUserResponse
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	fmt.Println(PrettyPrint(result))

	lnbitsWallet := result.Wallets[0]

	// create and store wallet

	newWallet := models.Wallet{
		TipId:          tip.ID,
		LnbitsWalletId: lnbitsWallet.WalletId,
		LnbitsUserId:   lnbitsWallet.User,
		AdminKey:       lnbitsWallet.Admin,
	}

	if err := t.walletStore.Create(&newWallet); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	// @Todo if creating or storing wallet fails, delete tip

	// @Todo recreateTipFundingInvoice

	// @Todo create Achievement

	return c.JSON(http.StatusCreated, newTipResponse(&tip))
}
