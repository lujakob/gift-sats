package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lujakob/gift-sats/tip"
	"github.com/lujakob/gift-sats/user"
)

type userSignupRequest struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required, email"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (r *userSignupRequest) bind(c *fiber.Ctx, u *user.User) error {
	//validate

	if err := c.BodyParser(r); err != nil {
		return err
	}
	//fmt.Printf("%v", *r)

	u.Username = r.User.Username
	u.Email = r.User.Email
	h, err := u.HashPassword(r.User.Password)
	if err != nil {
		return err
	}
	u.Password = h
	return nil
}

type userLoginRequest struct {
	User struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate: "required"`
	} `json:"user"`
}

func (r *userLoginRequest) bind(c *fiber.Ctx) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}

	//fmt.Printf("%v", *r)
	return nil
}

type tipCreateRequest struct {
	Tip struct {
		Amount   int  `json:"amount" validate:"required"`
		Fee      int  `json:"fee" validate:"required"`
		TipperID uint `json:"tipper_id" validate:"required"`
	} `json:"tip"`
}

func (r *tipCreateRequest) bind(c *fiber.Ctx, t *tip.Tip) error {
	//validate

	if err := c.BodyParser(r); err != nil {
		return err
	}
	//fmt.Printf("%v", *r)

	t.Amount = r.Tip.Amount
	t.Fee = r.Tip.Fee
	t.TipperID = r.Tip.TipperID

	return nil
}
