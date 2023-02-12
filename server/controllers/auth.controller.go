package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lujakob/gift-sats/interfaces"
	"github.com/lujakob/gift-sats/models"
	"github.com/lujakob/gift-sats/utils"
)

type userSignupRequest struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required, email"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (r *userSignupRequest) bind(c echo.Context, u *models.User) error {
	//validate

	if err := c.Bind(r); err != nil {
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
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (r *userLoginRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}

	//fmt.Printf("%v", *r)
	return nil
}

type userResponse struct {
	User struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"user"`
}

func newUserResponse(u *models.User) *userResponse {
	r := new(userResponse)
	r.User.ID = u.ID
	r.User.Username = u.Username
	r.User.Email = u.Email
	return r
}

type AuthController struct {
	userStore interfaces.IUserStore
}

func NewAuthController(userStore interfaces.IUserStore) *AuthController {
	return &AuthController{
		userStore: userStore,
	}
}

func (a *AuthController) Signup(c echo.Context) error {
	var u models.User
	req := &userSignupRequest{}

	if err := req.bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := a.userStore.Create(&u); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, newUserResponse(&u))
}

func (a *AuthController) Signin(c echo.Context) error {
	req := &userLoginRequest{}
	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	u, err := a.userStore.GetByEmail(req.User.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}
	if !u.CheckPassword(req.User.Password) {
		fmt.Printf("wrong password %v", err)
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}
	return c.JSON(http.StatusOK, newUserResponse(u))
}
