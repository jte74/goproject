package router

import (
	"training/goproject/interface/controller"
	"training/goproject/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
    e.HideBanner = true
    e.HidePort = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)
    e.POST("/auth", func(context echo.Context) error { return c.Auth.Auth(context) })
    e.GET("/", func(context echo.Context) error { return c.User.Home(context) })
    
    r := e.Group("/restricted")
    config := middleware.JWTConfig{
        Claims:     &utils.JwtCustomClaims{},
        SigningKey: []byte("secret"),
    }
    r.Use(middleware.JWTWithConfig(config))
    r.GET("", utils.Restricted)

	r.GET("/users", func(context echo.Context) error { return c.User.GetUsers(context) })

	// e.POST("/users", func(context echo.Context) error { return c.User.CreateUser(context) })

	return e
}

