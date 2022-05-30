package model

type User struct {
	// Id
	// in: int32
	Id        int32       `json:"id"`
	// Name
	// in: string
	Name      string     `json:"name"`
	// Firstname
	// in: string
	Firstname  string    `json:"firstname"`
	// Age
	// in: int32
	Age       int32     `json:"age"`
}

func (User) TableName() string { return "users" }