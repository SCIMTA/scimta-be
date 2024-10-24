package model

import (
	"errors"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int    `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func (u *User) HashPassword(p string) (string, error) {
	if len(p) == 0 {
		return "", errors.New("password is empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	return string(h), err
}

func (u *User) CheckPassword(p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	if err != nil {
		log.Info().Msg("password is" + u.Password + " and input is " + u.Username)
		log.Error().Err(err).Msg("password does not match")
	}
	return err == nil
}
