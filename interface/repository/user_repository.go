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

	conn := db.OpenDB()
	queries := db.New(conn)

	//TODO Verifier doublon

	dbuser := db.User{ID: u.Id, Name: u.Name, Firstname: u.Firstname, Age: u.Age, Password: u.Password}
	err := queries.CreateUser(context.Background(), dbuser)

	if err != nil {
		log.Println("Create User error:", err.Error())
	}

	return u, err
}

func (ur *userRepository) GetUser(id *int) (*model.User, error) {

	conn := db.OpenDB()
	queries := db.New(conn)

	user, err := queries.GetUser(context.Background(), *id)

	if err != nil {
		log.Fatal("User error:", err)
	}

	contractModels := &model.User{Id: user.ID, Name: user.Name, Firstname: user.Firstname, Age: user.Age}

	return contractModels, err
}

func (ur *userRepository) GetUsers() ([]*model.User, error) {

	conn := db.OpenDB()
	queries := db.New(conn)

	users, err := queries.GetUsers(context.Background())

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

	conn := db.OpenDB()
	queries := db.New(conn)

	err := queries.DeleteUser(context.Background(), *id)

	if err != nil {
		log.Fatal("DeleteUsers error:", err)
	}

	return err
}
