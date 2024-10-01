package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
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
