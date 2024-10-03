package middleware

import (
	"net/http"
	"scimta-be/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type jwtExtractor func(echo.Context) (string, error)

var (
	ErrJWTMissing = echo.NewHTTPError(http.StatusUnauthorized, "missing or malformed jwt")
	ErrJWTInvalid = echo.NewHTTPError(http.StatusForbidden, "invalid or expired jwt")
	JWTInstance   echo.MiddlewareFunc
)

func JWTWithConfig() echo.MiddlewareFunc {
	// extractor := jwtFromHeader("Authorization", "Bearer")
	config := echojwt.Config{
		TokenLookup: "header:Authorization:Bearer",
		SigningKey:  []byte("secret"),
	}
	return echojwt.WithConfig(config)
}

func NewJWTClaims(user *model.User) *jwtClaims {
	return &jwtClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
}

func NewTokenWithClaims(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// TODO: Implement other secret key management
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

func Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func Restricted(c echo.Context) error {
	user := c.Get("username").(*jwt.Token)
	claims := user.Claims.(*jwtClaims)
	name := claims.Username
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

// tokenParser returns a `jwtExtractor` that extracts token from the request header.

// jwtFromHeader returns a `jwtExtractor` that extracts token from the request header.
func jwtFromHeader(header string, authScheme string) jwtExtractor {
	return func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(header)
		l := len(authScheme)
		if len(auth) > l+1 && auth[:l] == authScheme {
			return auth[l+1:], nil
		}
		return "", ErrJWTMissing
	}
}
