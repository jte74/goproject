package controller

type AppController struct {
	User interface{ UserController }
	Auth interface{ AuthController }
}