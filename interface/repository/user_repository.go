package repository

import (
	"context"
	"database/sql"
	"log"
	"training/goproject/config"
	"training/goproject/domain/model"
	"training/goproject/infrastructure/db/sqlc"
)

type userRepository struct {
	db *sql.DB
}


type UserRepository interface {
	FindAll() ([]*model.User, error)
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
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
		contractModels[i] = &model.User{Id: v.ID, Name: v.Name, Firstname: v.Name, Age: v.Age}
	}

	return contractModels, err
}


