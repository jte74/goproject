package model

type User struct {
	Id        int32       `json:"id"`
	Name      string     `json:"name"`
	Firstname  string    `json:"firstname"`
	Age       int32     `json:"age"`
}

func (User) TableName() string { return "users" }