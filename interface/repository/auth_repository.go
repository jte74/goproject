package repository

import (
	"context"
	"database/sql"
	"log"
	"time"
	"training/goproject/domain/model"
	db "training/goproject/infrastructure/db/sqlc"
	"training/goproject/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type authRepository struct {
	db *sql.DB
}

type AuthRepository interface {
	Auth(u *model.Auth) (*jwt.Token, error)
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db}
}

func (ur *authRepository) Auth(u *model.Auth) (*jwt.Token, error) {

	claims := &utils.JwtCustomClaims{
		Name:  "",
		UUID:  "",
		Admin: false,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//TODO TO put in config file or get in database !
	// if u.Name != "pieter" || u.Password != "claerhout" {
	// 	return t, echo.ErrUnauthorized
	// }

	conn := db.OpenDB()
	queries := db.New(conn)

	user, err := queries.GetUserWithName(context.Background(), u.Name)

	if u.Password != user.Password {
		log.Println(err.Error())
		return t, echo.ErrUnauthorized
	}

	//TODO add more claims !
	claims = &utils.JwtCustomClaims{
		Name:  "Pieter Claerhout",
		UUID:  "9E98C454-C7AC-4330-B2EF-983765E00547",
		Admin: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token, nil
}
