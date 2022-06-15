package testutil

import (
	"database/sql"
	"testing"
	dbt "training/goproject/infrastructure/db/sqlc"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewDBMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock, error) {
	t.Helper()

	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	println(db)
	gdb := dbt.OpenDB()
	if err != nil {
		return nil, nil, err
	}

	return gdb, mock, nil
}
