package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lujakob/gift-sats/interfaces"
	"github.com/lujakob/gift-sats/models"
	"github.com/lujakob/gift-sats/utils"
)

type userListResponse struct {
	Users      []*userResponse `json:"users"`
	UsersCount int64           `json:"usersCount"`
}

func newUserListResponse(users []models.User, count int64) *userListResponse {
	r := new(userListResponse)
	ur := new(userResponse)
	for _, u := range users {

		ur.User.Username = u.Username
		ur.User.Email = u.Email

		r.Users = append(r.Users, ur)
	}
	r.UsersCount = count
	return r
}

type UserController struct {
	userStore interfaces.IUserStore
}

func NewUserController(userStore interfaces.IUserStore) *UserController {
	return &UserController{
		userStore: userStore,
	}
}

func (u *UserController) List(c echo.Context) error {
	list, count, error := u.userStore.GetAll()
	if error != nil {
		fmt.Println(error)
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	return c.JSON(http.StatusOK, newUserListResponse(list, count))
}
