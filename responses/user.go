package responses

import "scimta-be/model"

type UserResponse struct {
	User struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}
}

type UserRegisterResponse struct {
	Username string `json:"username"`
}

func NewUserRegisterResponse(u *model.User) *UserRegisterResponse {
	r := new(UserRegisterResponse)
	r.Username = u.Username
	return r
}
