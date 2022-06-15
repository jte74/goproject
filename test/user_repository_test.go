package test

import (
	"testing"
	"training/goproject/config"
	"training/goproject/domain/model"
	"training/goproject/interface/repository"
	"training/goproject/testutil"
)

func Test_Auth(t *testing.T) {

	db, mock, _ := testutil.NewDBMock(t)
	r := repository.NewAuthRepository(db)

	auth := model.Auth{
		Name:     "Tellier",
		Password: "lala",
	}

	token, err := r.Auth(&auth)

	if token == nil {
		t.Errorf(err.Error())
		println(mock)
	}
}

func Test_GetUsers(t *testing.T) {

	config.ReadConfig()
	db, mock, _ := testutil.NewDBMock(t)
	r := repository.NewUserRepository(db)

	users, err := r.GetUsers()

	if users == nil {
		t.Errorf(err.Error())
		println(mock)
	}
}

func Test_GetUser(t *testing.T) {

	config.ReadConfig()
	db, mock, _ := testutil.NewDBMock(t)
	r := repository.NewUserRepository(db)

	users, err := r.GetUsers()

	for _, user := range users {

		var id *int
		id = &user.Id
		user, err = r.GetUser(id)

		if err != nil {
			t.Errorf(err.Error())
			println(mock)
		}
	}
}

func Test_CreateUser(t *testing.T) {

	config.ReadConfig()
	db, mock, _ := testutil.NewDBMock(t)
	r := repository.NewUserRepository(db)

	user := model.User{
		Name:      "Tellier",
		Password:  "lala",
		Firstname: "CreateTest",
	}

	_, err := r.CreateUser(&user)

	if err != nil {
		t.Errorf(err.Error())
		println(mock)
	}
}

func Test_DeleteUser(t *testing.T) {

	config.ReadConfig()
	db, mock, _ := testutil.NewDBMock(t)
	r := repository.NewUserRepository(db)

	user := model.User{
		Name:      "Tellier",
		Password:  "lala",
		Firstname: "DeleteTest",
	}

	userCreate, err := r.CreateUser(&user)
	_, errorDelete := r.DeleteUser(&userCreate.Id)

	if errorDelete != nil {
		t.Errorf(err.Error())
		println(mock)
	}
}
