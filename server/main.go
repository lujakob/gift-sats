package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/lujakob/gift-sats/config"
	"github.com/lujakob/gift-sats/db"
	"github.com/lujakob/gift-sats/handler"
	"github.com/lujakob/gift-sats/tip"
	"github.com/lujakob/gift-sats/user"
)

func main() {
	config := config.GetConfig()

	app := fiber.New()
	app.Use(recover.New())

	d := db.New(config.DB_DSN)
	db.AutoMigrate(d)

	us := user.NewUserStore(d)
	ts := tip.NewTipStore(d)

	h := handler.NewHandler(us, ts)
	h.Register(app)

	err := app.Listen(":3100")

	if err != nil {
		fmt.Printf("%v", err)
	}
}
