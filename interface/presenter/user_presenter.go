package presenter

import (
	"training/goproject/domain/model"
)

type userPresenter struct {
}

type UserPresenter interface {
	ResponseUsers(us []*model.User) []*model.User
	ResponseUser(us *model.User) *model.User
	ResponseCreateUsers(us *model.User) *model.User
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

func (up *userPresenter) ResponseCreateUsers(us *model.User) *model.User {
	us.Name = "Mr." + us.Name
	return us
}
