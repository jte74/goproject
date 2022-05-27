package interactor

import (
	"training/goproject/domain/model"
	"training/goproject/usecase/presenter"
	"training/goproject/usecase/repository"
)

type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserInteractor interface {
	Get(u []*model.User) ([]*model.User, error)
}

func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
	return &userInteractor{r, p}
}

func (us *userInteractor) Get(u []*model.User) ([]*model.User, error) {
	u, err := us.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsers(u), nil
}
