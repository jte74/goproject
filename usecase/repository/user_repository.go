package repository

import "training/goproject/domain/model"

type UserRepository interface {
	GetUsers() ([]*model.User, error)
	GetUser(id *int) (*model.User, error)
	CreateUser(u *model.User) (int, error)
	DeleteUser(id *int) (string, error)
}
