package repository

import (
	"context"
	"database/sql"
	"log"
	"training/goproject/domain/model"
	db "training/goproject/infrastructure/db/sqlc"
)

type userRepository struct {
	db *sql.DB
}

func DbInitialize() *db.Queries {
	return db.New(db.OpenDB())
}

type UserRepository interface {
	GetUsers() ([]*model.User, error)
	CreateUser(*model.User) (*model.User, error)
	DeleteUser(*int) error
	GetUser(id *int) (*model.User, error)
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(u *model.User) (*model.User, error) {

	//TODO Verifier doublon
	dbuser := db.User{ID: u.Id, Name: u.Name, Firstname: u.Firstname, Age: u.Age, Password: u.Password}
	err := DbInitialize().CreateUser(context.Background(), dbuser)

	if err != nil {
		log.Println("Create User error:", err.Error())
	}

	return u, err
}

func (ur *userRepository) GetUser(id *int) (*model.User, error) {

	user, err := DbInitialize().GetUser(context.Background(), *id)

	if err != nil {
		log.Fatal("User error:", err)
	}

	contractModels := &model.User{Id: user.ID, Name: user.Name, Firstname: user.Firstname, Age: user.Age}

	return contractModels, err
}

func (ur *userRepository) GetUsers() ([]*model.User, error) {

	users, err := DbInitialize().GetUsers(context.Background())

	if err != nil {
		log.Fatal("ListUsers error:", err)
	}

	contractModels := make([]*model.User, len(users))
	for i, v := range users {
		contractModels[i] = &model.User{Id: v.ID, Name: v.Name, Firstname: v.Firstname, Age: v.Age}
	}

	return contractModels, err
}

func (ur *userRepository) DeleteUser(id *int) error {

	err := DbInitialize().DeleteUser(context.Background(), *id)

	if err != nil {
		log.Fatal("DeleteUsers error:", err)
	}

	return err
}
