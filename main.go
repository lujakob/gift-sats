package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/lujakob/gifting-sats/db"
	"github.com/lujakob/gifting-sats/handler"
	"github.com/lujakob/gifting-sats/user"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())

	d := db.New()
	db.AutoMigrate(d)

	us := user.NewUserStore(d)

	h := handler.NewHandler(us)
	h.Register(app)

	err := app.Listen(":3000")

	if err != nil {
		fmt.Printf("%v", err)
	}
}
