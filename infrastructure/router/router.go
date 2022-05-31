package router

import (
	"encoding/json"
	"net/http"
	"time"
	"training/goproject/domain/model"
	"training/goproject/interface/controller"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pieterclaerhout/go-log"
	"github.com/swaggo/echo-swagger"
)


const secret = "secret"
 
type jwtCustomClaims struct {
    Name  string `json:"name"`
    UUID  string `json:"uuid"`
    Admin bool   `json:"admin"`
    jwt.StandardClaims
}

func restricted(c echo.Context) error {
    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(*jwtCustomClaims)
    log.InfoDump(claims, "claims")
    name := claims.Name
    return c.String(http.StatusOK, "Welcome "+name+"!")
}

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
    e.HideBanner = true
    e.HidePort = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)
    e.POST("/login", login)
    e.GET("/", func(context echo.Context) error { return c.Home(context) })
    
    r := e.Group("/restricted")
    config := middleware.JWTConfig{
        Claims:     &jwtCustomClaims{},
        SigningKey: []byte("secret"),
    }
    r.Use(middleware.JWTWithConfig(config))
    r.GET("", restricted)

	r.GET("/users", func(context echo.Context) error { return c.GetUsers(context) })

	// e.GET("/users", func(context echo.Context) error { return c.GetUsers(context) })
	// e.POST("/users", func(context echo.Context) error { return c.User.CreateUser(context) })

	return e
}

// login
// @Summary login
// @Description login
// @Tags login
// @Accept  json
// @Produce  json
// @Param users body model.User true "login user"  
// @Success 200 
// @Router /login [post]
func login(c echo.Context) error {
 
	var user model.User
	json.NewDecoder(c.Request().Body).Decode(&user)
    if user.Name != "pieter" || user.Password != "claerhout" {
        return echo.ErrUnauthorized
    }
 
    claims := &jwtCustomClaims{
        Name:  "Pieter Claerhout",
        UUID:  "9E98C454-C7AC-4330-B2EF-983765E00547",
        Admin: true,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
        },
    }
 
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
 
    t, err := token.SignedString([]byte(secret))
    if err != nil {
        return err
    }
 
    return c.JSON(http.StatusOK, map[string]string{
        "token": t,
    })
}