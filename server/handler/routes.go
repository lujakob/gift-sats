package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lujakob/gift-sats/user"
	"github.com/lujakob/gift-sats/utils"
)

func (h *Handler) Register(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	users := v1.Group("/users")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	users.Get("", func(c *fiber.Ctx) error {
		list, count, error := h.userStore.GetAll()
		if error != nil {
			return c.Status(http.StatusNotFound).JSON(utils.NotFound())
		}

		return c.Status(http.StatusOK).JSON(newUserListResponse(list, count))
	})

	users.Post("", func(c *fiber.Ctx) error {
		newUser := new(user.User)

		if err := c.BodyParser(newUser); err != nil {
			return err
		}

		error := h.userStore.Create(newUser)
		if error != nil {
			fmt.Printf("%v", error)
			return c.Status(http.StatusNotFound).JSON(utils.NotFound())
		}
		return c.Status(http.StatusCreated).JSON(newUser)
	})
}
