package router

import (
	"training/goproject/interface/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(context echo.Context) error { return c.GetUsers(context) })
	// e.POST("/users", func(context echo.Context) error { return c.User.CreateUser(context) })
	e.GET("/", func(context echo.Context) error { return c.Home(context) })
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	return e
}
