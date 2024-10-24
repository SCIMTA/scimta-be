package services

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	"scimta-be/model"
	"scimta-be/responses"
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

// TODO: Implement this please
func (us *UserServices) GetUsers(c echo.Context) error {
	tempUser := new(responses.UserResponse)
	tempUser.User.ID = 1
	tempUser.User.Username = "john_doe"
	return c.JSON(http.StatusOK, tempUser)
}
