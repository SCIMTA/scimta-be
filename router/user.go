package router

import (
	"net/http"
	"scimta-be/responses"
	"scimta-be/router/middleware"
	"scimta-be/services"
	"scimta-be/utils"

	"github.com/labstack/echo/v4"
)

type UserRouter struct {
	userService *services.UserServices
}

func NewUserRouter(sg *echo.Group, us *services.UserServices) *UserRouter {
	ur := &UserRouter{userService: us}

	user := sg.Group("/user")
	user.Use(middleware.JWTWithConfig())
	user.GET("", ur.GetUser)

	return ur
}

// CurrentUser godoc
// @Summary Get the current user
// @Description Gets the currently logged-in user
// @ID current-user
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.UserResponse
// @Security ApiKeyAuth
// @Router /user [get]
func (ur *UserRouter) GetUser(c echo.Context) error {
	user, err := ur.userService.GetUser(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if user == nil {
		return c.JSON(http.StatusBadRequest, utils.ErrUserNotFound)
	}
	return c.JSON(http.StatusOK, responses.GetUserResponse(user))
}
