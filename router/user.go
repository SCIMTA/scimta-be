package router

import (
	"net/http"
	"scimta-be/model"
	"scimta-be/requests"
	"scimta-be/responses"
	"scimta-be/services"

	"github.com/labstack/echo/v4"
)

type UserRouter struct {
	userService *services.UserServices
}

func NewUserRouter(sg *echo.Group, us *services.UserServices) *UserRouter {
	ur := &UserRouter{userService: us}

	guest := sg.Group("/auth")
	guest.POST("/register", ur.Register)

	user := sg.Group("/user")
	user.GET("", ur.GetUsers)

	return ur
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user
// @ID register
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body requests.UserRegisterRequest true "User info for registration"
// @Success 201 {object} responses.UserRegisterResponse
// @Router /auth/register [post]
func (ur *UserRouter) Register(c echo.Context) error {
	var user model.User
	req := &requests.UserRegisterRequest{}
	if err := req.Bind(c, &user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	if err := ur.userService.Create(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, responses.NewUserRegisterResponse(&user))
}

// CurrentUser godoc
// @Summary Get the current user
// @Description Gets the currently logged-in user
// @ID current-user
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.UserResponse
// @Router /user [get]
func (ur *UserRouter) GetUsers(c echo.Context) error {
	user := ur.userService.GetUsers(c)
	return user
}
