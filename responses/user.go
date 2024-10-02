package responses

import (
	"scimta-be/model"
	"scimta-be/utils"
)

type UserResponse struct {
	User struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}
}

type UserLoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func NewUserLoginResponse(u *model.User) *UserLoginResponse {
	r := new(UserLoginResponse)
	r.Username = u.Username
	t, err := utils.GenerateJWT(u)
	if err != nil {
		return nil
	}
	r.Token = t
	return r
}

type UserRegisterResponse struct {
	Username string `json:"username"`
}

func NewUserRegisterResponse(u *model.User) *UserRegisterResponse {
	r := new(UserRegisterResponse)
	r.Username = u.Username
	return r
}
