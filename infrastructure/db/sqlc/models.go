// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"time"
)

type User struct {
	ID          int32
	Name        string
	Firstname   string
	Age         int32
	Password    string
	Token       string
	Datecreated time.Time
}
