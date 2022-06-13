package model

type User struct {
	// Id
	// in: int32
	Id int32 `json:"id,omitempty"`
	// Name
	// in: string
	Name string `json:"name,omitempty"`
	// Firstname
	// in: string
	Firstname string `json:"firstname,omitempty"`
	// Age
	// in: int32
	Age int32 `json:"age,omitempty"`
	// Password
	// in: string
	Password string `json:"password,omitempty"`
	// Token
	// in: string
	Token string `json:"token,omitempty"`
	// DateCreated
	// in: date
	DateCreated string `json:"datecreated,omitempty"`
}

func (User) TableName() string { return "users" }
