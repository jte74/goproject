package controller

import (
	"errors"
	"fmt"
	"log"
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
	GetUser(c Context) error
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
// @Tags User
// @Accept  json
// @Produce  json
// @Security Bearer
// @Success 200
// @Router /restricted/users [get]
func (uc *userController) GetUsers(c Context) error {
	fmt.Println("Endpoint Hit: GetUsers")

	u, err := uc.userInteractor.GetUsers()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

// Return User godoc
// @Summary Return User
// @Description Return User
// @Tags User
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param id path integer true "id"
// @Success 200
// @Router /restricted/user-id/{id} [get]
func (uc *userController) GetUser(c Context) error {
	fmt.Println("Endpoint Hit: GetUser")

	id, iderr := strconv.Atoi(c.Param("id"))
	u, err := uc.userInteractor.GetUser(&id)

	if err != nil || iderr != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

// Create Users godoc
// @Summary Return name created
// @Description Create User
// @Tags User
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
// @Tags User
// @Accept  json
// @Produce  json
// @Security Bearer
// @Param id path integer true "id"
// @Success 200
// @Router /restricted/delete-user/{id} [delete]
func (uc *userController) DeleteUser(c Context) error {
	fmt.Println("Endpoint Hit: DeleteUser")

	id, iderr := strconv.Atoi(c.Param("id"))

	log.Println("TESTSETSTSETESTSE1566id", id)

	err := uc.userInteractor.DeleteUser(&id)

	if err != nil || iderr != nil {
		return err
	}
	return c.JSON(http.StatusOK, nil)
}

func (uc *userController) Home(c Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}
