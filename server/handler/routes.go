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
}
