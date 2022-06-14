package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"training/goproject/domain/model"
	"training/goproject/usecase/interactor"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	GetUsers(c Context) error
	CreateUser(c Context) error
	DeleteUser(c Context) error
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

// Create Users godoc
// @Summary Return Users
// @Description Create User
// @Tags id
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param users body model.User true "login user"
// @Success 200
// @Router /restricted/create-user [post]
func (uc *userController) CreateUser(c Context) error {
	fmt.Println("Endpoint Hit: CreateUser")
	var params model.User
	if err := c.Bind(&params); !errors.Is(err, nil) {
		return err
	}

	u, err := uc.userInteractor.CreateUser(&params)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u.Name)
}

// Delete Users godoc
// @Summary Delete Users
// @Description Delete User
// @Tags id
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param id path integer true "id"
// @Success 200
// @Router /restricted/delete-user/{id} [delete]
func (uc *userController) DeleteUser(c Context) error {
	fmt.Println("Endpoint Hit: DeleteUser")

	id, iderr := strconv.Atoi(c.QueryParam("id"))

	err := uc.userInteractor.DeleteUser(&id)

	if err != nil || iderr != nil {
		return err
	}
	return c.JSON(http.StatusOK, nil)
}

func (uc *userController) Home(c Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}
