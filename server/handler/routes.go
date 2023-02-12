package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lujakob/gift-sats/config"
	"github.com/lujakob/gift-sats/tip"
	"github.com/lujakob/gift-sats/user"
	"github.com/lujakob/gift-sats/utils"
	"github.com/lujakob/gift-sats/wallet"
)

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

func (h *Handler) Register(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	users := v1.Group("/users")
	tips := v1.Group("/tips")
	auth := v1.Group("/auth")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	users.Get("", func(c *fiber.Ctx) error {
		list, count, error := h.userStore.GetAll()
		if error != nil {
			fmt.Println(error)
			return c.Status(http.StatusNotFound).JSON(utils.NotFound())
		}

		return c.Status(http.StatusOK).JSON(newUserListResponse(list, count))
	})

	auth.Post("/signup", func(c *fiber.Ctx) error {
		var u user.User
		req := &userSignupRequest{}

		if err := req.bind(c, &u); err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
		}
		if err := h.userStore.Create(&u); err != nil {
			fmt.Println(err)
			return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
		}

		return c.Status(http.StatusCreated).JSON(newUserResponse(&u))
	})

	auth.Post("/signin", func(c *fiber.Ctx) error {
		req := &userLoginRequest{}
		if err := req.bind(c); err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
		}
		u, err := h.userStore.GetByEmail(req.User.Email)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(utils.NewError(err))
		}
		if u == nil {
			return c.Status(http.StatusForbidden).JSON(utils.AccessForbidden())
		}
		if !u.CheckPassword(req.User.Password) {
			fmt.Printf("wrong password %v", err)
			return c.Status(http.StatusForbidden).JSON(utils.AccessForbidden())
		}
		return c.Status(http.StatusOK).JSON(newUserResponse(u))
	})

	tips.Get("/", func(c *fiber.Ctx) error {
		list, count, error := h.tipStore.GetAll()
		if error != nil {
			fmt.Println(error)
			return c.Status(http.StatusNotFound).JSON(utils.NotFound())
		}

		return c.Status(http.StatusOK).JSON(newTipListResponse(list, count))
	})

	tips.Post("/", func(c *fiber.Ctx) error {
		config := config.GetConfig()

		// @Todo throw error for missing configs
		// if !config.LNBITS_API_KEY || !config.LNBITS_URL || !config.LNBITS_USER_ID

		// @Todo: verify "Only tips with positive, whole amounts are allowed"

		// @Todo: calculateFee

		// first create and store tip
		var t tip.Tip
		req := &tipCreateRequest{}
		if err := req.bind(c, &t); err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
		}

		if err := h.tipStore.Create(&t); err != nil {
			fmt.Println(err)
		}

		fmt.Println(PrettyPrint(t))

		// fetch LnbitsWallet

		createLnbitsUserRequest := &CreateLnbitsUserRequest{
			AdminId:    config.LNBITS_USER_ID,
			UserName:   fmt.Sprint("tip-", t.ID),
			WalletName: fmt.Sprint("tip-", t.ID),
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

		newWallet := wallet.Wallet{
			TipId:          t.ID,
			LnbitsWalletId: lnbitsWallet.WalletId,
			LnbitsUserId:   lnbitsWallet.User,
			AdminKey:       lnbitsWallet.Admin,
		}

		if err := h.walletStore.Create(&newWallet); err != nil {
			fmt.Println(err)
			return c.Status(http.StatusUnprocessableEntity).JSON(utils.NewError(err))
		}

		// @Todo if creating or storing wallet fails, delete tip

		// @Todo recreateTipFundingInvoice

		// @Todo create Achievement

		return c.Status(http.StatusCreated).JSON(newTipResponse(&t))
	})
}
