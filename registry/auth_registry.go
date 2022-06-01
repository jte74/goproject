package registry

import (
	"training/goproject/interface/controller"
	ip "training/goproject/interface/presenter"
	ir "training/goproject/interface/repository"
	"training/goproject/usecase/interactor"
	up "training/goproject/usecase/presenter"
	ur "training/goproject/usecase/repository"
)

func (r *registry) NewAuthController() controller.AuthController {
	return controller.NewAuthController(r.NewAuthInteractor())
}

func (r *registry) NewAuthInteractor() interactor.AuthInteractor {
	return interactor.NewAuthInteractor(r.NewAuthRepository(), r.NewAuthPresenter())
}

func (r *registry) NewAuthRepository() ur.AuthRepository {
	return ir.NewAuthRepository(r.db)
}

func (r *registry) NewAuthPresenter() up.AuthPresenter {
	return ip.NewAuthPresenter()
}
