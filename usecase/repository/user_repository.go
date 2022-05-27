package repository

import "training/goproject/domain/model"

type UserRepository interface {
	FindAll() ([]*model.User, error)
	// Create(u *model.User) (*model.User, error)
}
