package controller

import (
	"errors"
	"fmt"
	"net/http"
	"training/goproject/domain/model"
	"training/goproject/usecase/interactor"
)


type authController struct {
	authInteractor interactor.AuthInteractor
}

type AuthController interface {
	Auth(c Context) error
}

func NewAuthController(us interactor.AuthInteractor) AuthController {
	return &authController{us}
}

// login
// @Summary login
// @Description login
// @Tags login
// @Accept  json
// @Produce  json
// @Param users body model.User true "login user"  
// @Success 200 
// @Router /auth [post]
func (uc *authController) Auth(c Context) error {
    fmt.Println("Endpoint Hit: Auth")
	var params model.User
	if err := c.Bind(&params); !errors.Is(err, nil) {
		return err
	}
	u, err := uc.authInteractor.GetToken(&params)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}