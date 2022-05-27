// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: query.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (
  name, firstname, age
) VALUES (
  $1, $2, $3
)
RETURNING id, name, firstname, age
`

type CreateUserParams struct {
	Name      string
	Firstname string
	Age       int32
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error  {
	_,err := q.db.ExecContext(ctx, createUser, arg.Name, arg.Firstname, arg.Age)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, name, firstname, age FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.Id,
		&i.Name,
		&i.Firstname,
		&i.Age,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, name, firstname, age FROM users
ORDER BY name
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {

	rows, err := q.db.QueryContext(ctx, listUsers)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.Id,
			&i.Name,
			&i.Firstname,
			&i.Age,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	
	return items, nil
}
