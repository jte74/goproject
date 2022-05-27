package registry

import (
	"training/goproject/interface/controller"
	ip "training/goproject/interface/presenter"
	ir "training/goproject/interface/repository"
	"training/goproject/usecase/interactor"
	up "training/goproject/usecase/presenter"
	ur "training/goproject/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
