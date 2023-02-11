package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/lujakob/gift-sats/db"
	"github.com/lujakob/gift-sats/handler"
	"github.com/lujakob/gift-sats/user"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())

	d := db.New()
	db.AutoMigrate(d)

	us := user.NewUserStore(d)

	h := handler.NewHandler(us)
	h.Register(app)

	err := app.Listen(":3100")

	if err != nil {
		fmt.Printf("%v", err)
	}
}
