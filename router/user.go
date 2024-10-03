package router

import (
	"net/http"
	"scimta-be/model"
	"scimta-be/requests"
	"scimta-be/responses"
	"scimta-be/router/middleware"
	"scimta-be/services"
	"scimta-be/utils"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type UserRouter struct {
	userService *services.UserServices
}

func NewUserRouter(sg *echo.Group, us *services.UserServices) *UserRouter {
	ur := &UserRouter{userService: us}

	guest := sg.Group("/auth")
	guest.POST("/register", ur.Register)
	guest.POST("/login", ur.Login)

	user := sg.Group("/user")
	user.Use(middleware.JWTWithConfig())
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

// Login godoc
// @Summary Login
// @Description Login for user
// @ID login
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body requests.UserLoginRequest true "User info for login"
// @Success 200 {object} responses.UserLoginResponse
// @Router /auth/login [post]
func (ur *UserRouter) Login(c echo.Context) error {
	req := &requests.UserLoginRequest{}
	if err := req.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	u, err := ur.userService.GetByUsername(req.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, utils.ErrAuthWrongCredentials)
	}
	log.Info().Msgf("User: %v", u)
	if !u.CheckPassword(req.Password) {
		return c.JSON(http.StatusUnauthorized, utils.ErrAuthWrongCredentials)
	}
	return c.JSON(http.StatusOK, responses.NewUserLoginResponse(u))
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
func (ur *UserRouter) GetUsers(c echo.Context) error {
	for key, values := range c.Request().Header {
		log.Info().Msg(key)
		for _, value := range values {
			log.Info().Msg(value)
		}
	}
	user := ur.userService.GetUsers(c)
	return user
}
