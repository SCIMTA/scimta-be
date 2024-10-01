package router

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"scimta-be/responses"
)

type UserRouter struct {
}

func NewUserRouter(g *echo.Group) *UserRouter {
	userRouter := &UserRouter{}
	g.GET("/users", userRouter.GetUsers)

	return userRouter
}

// CurrentUser godoc
// @Summary Get the current user
// @Description Gets the currently logged-in user
// @ID current-user
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.UserResponse
// @Router /users [get]
func (u *UserRouter) GetUsers(c echo.Context) error {
	tempUser := new(responses.UserResponse)
	tempUser.User.ID = 1
	tempUser.User.Name = "John Doe"
	return c.JSON(http.StatusOK, tempUser)
}
