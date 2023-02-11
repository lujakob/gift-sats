package handler

import (
	"github.com/lujakob/gift-sats/user"
)

type userResponse struct {
	User struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
	} `json:"user"`
}

func newUserResponse(u *user.User) *userResponse {
	r := new(userResponse)
	r.User.ID = u.ID
	r.User.Username = u.Username
	return r
}

type userListResponse struct {
	Users      []*userResponse `json:"users"`
	UsersCount int64           `json:"usersCount"`
}

func newUserListResponse(users []user.User, count int64) *userListResponse {
	r := new(userListResponse)
	ur := new(userResponse)
	for _, u := range users {

		ur.User.Username = u.Username

		r.Users = append(r.Users, ur)
	}
	r.UsersCount = count
	return r
}
