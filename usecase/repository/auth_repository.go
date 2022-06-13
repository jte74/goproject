package repository

import (
	"training/goproject/domain/model"

	"github.com/dgrijalva/jwt-go"
)

type AuthRepository interface {
	Auth(u *model.Auth) (*jwt.Token, error)
}
