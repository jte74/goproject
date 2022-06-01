package controller

import (
	"fmt"
	"net/http"
	"training/goproject/domain/model"
	"training/goproject/usecase/interactor"
)


type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	GetUsers(c Context) error
	Home(c Context) error
}

func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{us}
}

// Return All Users godoc
// @Summary Return All Users
// @Description Return All Users
// @Tags id
// @Accept  json
// @Produce  json
// @Security Bearer
// @Success 200 
// @Router /restricted/users [get]
func (uc *userController) GetUsers(c Context) error {
	fmt.Println("Endpoint Hit: GetUsers")
	var u []*model.User

	u, err := uc.userInteractor.Get(u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func (uc *userController) Home(c Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}