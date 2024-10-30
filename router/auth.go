package router

import (
	"net/http"
	"scimta-be/requests"
	"scimta-be/responses"
	"scimta-be/router/middleware"
	"scimta-be/services"
	"scimta-be/utils"

	"github.com/labstack/echo/v4"
)


type AuthRouter struct {
	authService *services.AuthServices
}

func NewAuthRouter(sg *echo.Group, as *services.AuthServices) *AuthRouter {
	ar := &AuthRouter{authService: as}

	auth := sg.Group("/auth")
	auth.POST("/register", ar.Register)
	auth.POST("/login", ar.Login)

	auth.Use(middleware.JWTWithConfig())
	auth.POST("/change-password", ar.ChangePassword)

	return ar
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user
// @ID register
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body requests.UserRegisterRequest true "User info for registration"
// @Success 201 {object} responses.UserRegisterResponse
// @Router /auth/register [post]
func (ar *AuthRouter) Register(c echo.Context) error {
	req := &requests.UserRegisterRequest{}
	if err := req.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	if err := ar.authService.Register(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, responses.NewUserRegisterResponse(req.Username))
}

// Login godoc
// @Summary Login
// @Description Login for user
// @ID login
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body requests.UserLoginRequest true "User info for login"
// @Success 200 {object} responses.UserLoginResponse
// @Router /auth/login [post]
func (ar *AuthRouter) Login(c echo.Context) error {
	req := &requests.UserLoginRequest{}
	if err := req.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	tokenStr, err := ar.authService.Login(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if tokenStr == "" {
		return c.JSON(http.StatusUnauthorized, utils.ErrAuthWrongCredentials)
	}
	return c.JSON(http.StatusOK, responses.NewUserLoginResponse(req.Username, tokenStr))
}

// ChangePassword godoc
// @Summary Change password
// @Description 
// @ID change-password
// @Tags auth
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user body requests.UserChangePasswordRequest true "Change password"
// @Success 200
// @Router /auth/change-password [post]
func (ar *AuthRouter) ChangePassword(c echo.Context) error {
	req := &requests.UserChangePasswordRequest{}
	if err := req.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	if err := ar.authService.ChangePassword(c, req); err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, responses.NewBaseResponse("success"))
}
