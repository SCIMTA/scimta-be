package services

import (
	"errors"
	"scimta-be/model"
	"scimta-be/requests"
	"scimta-be/utils"

	"github.com/labstack/echo/v4"
)

type AuthServices struct{
	userService *UserServices
}

func NewAuthServices(us *UserServices) *AuthServices {
	return &AuthServices{userService: us}
}

func (as *AuthServices) Register(req *requests.UserRegisterRequest) error {
	var user model.User
	hash, err := user.HashPassword(req.Password)
	if err != nil {
		return err
	}
	user.Username = req.Username
	user.Password = hash
	if err := as.userService.Create(&user); err != nil {
		return err
	}
	return nil
}

func (as *AuthServices) Login(req *requests.UserLoginRequest) (string, error) {
	u, err := as.userService.GetByUsername(req.Username)
	if err != nil {
		return "", err
	}
	if u == nil || !u.CheckPassword(req.Password) {
		return "", nil
	}
	tokenStr, err := utils.GenerateJWT(u)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (as *AuthServices) ChangePassword(c echo.Context, req *requests.UserChangePasswordRequest) error {
	user, err := as.userService.GetUser(c)
	if err != nil {
		return err
	}
	if user == nil || !user.CheckPassword(req.OldPassword) {
		return errors.New(utils.ErrPasswordNotMatch)
	}
	hash, err := user.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}
	
	user.Password = hash
	if err := as.userService.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}
