package services

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	"scimta-be/model"
	"scimta-be/utils"
)

type UserServices struct {
	db *gorm.DB
}

func NewUserServices(db *gorm.DB) *UserServices {
	return &UserServices{db: db}
}

func (us *UserServices) GetByUsername(username string) (*model.User, error) {
	var m model.User
	log.Info().Msg("username is " + username)
	if err := us.db.Where(&model.User{Username: username}).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserServices) Create(u *model.User) error {
	return us.db.Create(u).Error
}

func (us *UserServices) GetUser(c echo.Context) (*model.User, error) {
	tokenStr := c.Request().Header.Get("Authorization")
	username, err := utils.VerifyJWT(tokenStr)
	if err != nil {
		return nil, err
	}
	var user = model.User{Username: username}
	if err := us.db.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
