package middleware

import (
	"log"
	"net/http"
	"os"
	"scimta-be/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
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

func loadSecret() []byte {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secret := os.Getenv("JWT_SECRET")
	return []byte(secret)
}

func JWTWithConfig() echo.MiddlewareFunc {
	// extractor := jwtFromHeader("Authorization", "Bearer")
	config := echojwt.Config{
		TokenLookup: "header:Authorization:Bearer ",
		SigningKey:  loadSecret(),
		ContextKey:  "username",
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
	t, err := token.SignedString(loadSecret())
	if err != nil {
		return "", err
	}
	return t, nil
}

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

func VerifyToken(tokenStr string) (*jwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return loadSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*jwtClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrJWTInvalid
}
