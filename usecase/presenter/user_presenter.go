package presenter

import "training/goproject/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
	ResponseUser(u *model.User) *model.User
	ResponseCreateUsers(id int) string
}
