package responses

import (
	"scimta-be/model"
)

type BaseResponse struct {
	Message string `json:"message"`
}

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
func NewBaseResponse(message string) *BaseResponse {
	return &BaseResponse{Message: message}
}

func NewUserLoginResponse(username string, tokenStr string) *UserLoginResponse {
	r := new(UserLoginResponse)
	r.Username = username
	r.Token = tokenStr
	return r
}

type UserRegisterResponse struct {
	Username string `json:"username"`
}

func NewUserRegisterResponse(username string) *UserRegisterResponse {
	r := new(UserRegisterResponse)
	r.Username = username
	return r
}

func GetUserResponse(u *model.User) *UserResponse {
	r := new(UserResponse)
	r.User.ID = u.Id
	r.User.Username = u.Username
	return r
}
