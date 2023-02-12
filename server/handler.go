package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lujakob/gift-sats/controllers"
	"github.com/lujakob/gift-sats/interfaces"
)

type Handler struct {
	userStore   interfaces.IUserStore
	tipStore    interfaces.ITipStore
	walletStore interfaces.IWalletStore
}

func NewHandler(us interfaces.IUserStore, ts interfaces.ITipStore, ws interfaces.IWalletStore) *Handler {
	return &Handler{
		userStore:   us,
		tipStore:    ts,
		walletStore: ws,
	}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	authController := controllers.NewAuthController(h.userStore)
	userController := controllers.NewUserController(h.userStore)
	tipController := controllers.NewTipController(h.tipStore, h.walletStore)
	walletController := controllers.NewWalletController(h.walletStore)

	v1 := e.Group("/api/v1")

	v1.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	auth := v1.Group("/auth")
	user := v1.Group("/users")
	tips := v1.Group("/tips")
	wallets := v1.Group("/wallets")

	auth.POST("/signup", authController.Signup)
	auth.POST("/signin", authController.Signin)

	user.GET("", userController.List)

	tips.GET("", tipController.List)
	tips.POST("", tipController.Create)

	wallets.GET("", walletController.List)
}
