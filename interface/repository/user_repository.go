package repository

import (
	"context"
	"database/sql"
	"log"
	"training/goproject/config"
	"training/goproject/domain/model"
	db "training/goproject/infrastructure/db/sqlc"
)

type userRepository struct {
	db *sql.DB
}

type UserRepository interface {
	FindAll() ([]*model.User, error)
	CreateUser(*model.User) (*model.User, error)
	DeleteUser(*int) error
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(u *model.User) (*model.User, error) {
	config.ReadConfig()
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

func (ur *userRepository) FindAll() ([]*model.User, error) {
	//TODO A REVOIR (automapper + db)
	config.ReadConfig()
	conn := db.OpenDB()
	queries := db.New(conn)

	users, err := queries.ListUsers(context.Background())

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
	//TODO A REVOIR (automapper + db)
	config.ReadConfig()
	conn := db.OpenDB()
	queries := db.New(conn)

	var idtest int
	idtest = *id

	err := queries.DeleteUser(context.Background(), idtest)

	if err != nil {
		log.Fatal("DeleteUsers error:", err)
	}

	return err
}
