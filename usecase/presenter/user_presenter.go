package presenter

import "training/goproject/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
