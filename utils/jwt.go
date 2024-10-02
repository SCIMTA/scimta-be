package utils

import (
	"scimta-be/model"
	"scimta-be/router/middleware"
)

var JWTSecret = []byte("secret")

func GenerateJWT(u *model.User) (string, error) {
	claims := middleware.NewJWTClaims(u)
	t, err := middleware.NewTokenWithClaims(claims)
	if err != nil {
		return "", err
	}
	return t, nil
}
