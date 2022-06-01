package utils

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/pieterclaerhout/go-log"
)

const Secret = "secret"

type JwtCustomClaims struct {
    Name  string `json:"name"`
    UUID  string `json:"uuid"`
    Admin bool   `json:"admin"`
    jwt.StandardClaims
}

func Restricted(c echo.Context) error {
    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(*JwtCustomClaims)
    log.InfoDump(claims, "claims")
    name := claims.Name
    return c.String(http.StatusOK, "Welcome "+name+"!")
}