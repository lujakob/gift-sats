package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/lujakob/gifting-sats/db"
	"github.com/lujakob/gifting-sats/handler"
	"github.com/lujakob/gifting-sats/user"
	"github.com/lujakob/gifting-sats/utils"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App, db *gorm.DB) {
	us := user.NewUserStore(db)

	app.Get("/api/v1/users", func(c *fiber.Ctx) error {
		list, error := us.GetAll()
		if error != nil {
			return c.Status(http.StatusNotFound).JSON(utils.NotFound())
		}
		return c.Status(http.StatusOK).JSON(list)
	})

	app.Post("/api/v1/users", func(c *fiber.Ctx) error {
		newUser := new(user.User)

		if err := c.BodyParser(newUser); err != nil {
			return err
		}

		error := us.Create(newUser)
		if error != nil {
			fmt.Printf("%v", error)
			return c.Status(http.StatusNotFound).JSON(utils.NotFound())
		}
		return c.Status(http.StatusCreated).JSON(newUser)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}

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
