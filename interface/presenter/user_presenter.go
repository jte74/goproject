package presenter

import (
	"fmt"
	"training/goproject/domain/model"
)

type userPresenter struct {
}

type UserPresenter interface {
	ResponseUsers(us []*model.User) []*model.User
	ResponseUser(us *model.User) *model.User
	ResponseCreateUsers(id int) string
}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) ResponseUsers(us []*model.User) []*model.User {
	for _, u := range us {
		u.Name = "Mr." + u.Name
	}
	return us
}

func (up *userPresenter) ResponseUser(user *model.User) *model.User {
	user.Name = "Mr." + user.Name
	return user
}

func (up *userPresenter) ResponseCreateUsers(id int) string {
	response := fmt.Sprintf("Id number : %d", id)
	return response
}
