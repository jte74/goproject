// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"database/sql"
)

type User struct {
	ID          int
	Name        string
	Firstname   string
	Age         int32
	Password    string
	Token       sql.NullString
	Datecreated sql.NullTime
}
