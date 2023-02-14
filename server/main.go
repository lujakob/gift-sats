package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lujakob/gift-sats/config"
	"github.com/lujakob/gift-sats/db"
	"github.com/lujakob/gift-sats/stores"
)

func main() {
	config := config.GetConfig()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	d := db.New()
	db.AutoMigrate(d)

	us := stores.NewUserStore(d)
	ts := stores.NewTipStore(d)
	ws := stores.NewWalletStore(d)

	h := NewHandler(us, ts, ws)
	h.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", config.SERVER_HOST, config.SERVER_PORT)))

}
