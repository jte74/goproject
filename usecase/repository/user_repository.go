package repository

import "training/goproject/domain/model"

type UserRepository interface {
	FindAll() ([]*model.User, error)
	CreateUser(u *model.User) (*model.User, error)
	DeleteUser(id *int) error
}
