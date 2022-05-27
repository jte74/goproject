package registry

import (
	"database/sql"
	"training/goproject/interface/controller"
)

type registry struct {
	db *sql.DB
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *sql.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewUserController()
}