package interactor

import (
	"training/goproject/domain/model"
	"training/goproject/utils"
	"training/goproject/usecase/presenter"
	"training/goproject/usecase/repository"
)

type authInteractor struct {
	AuthRepository repository.AuthRepository
	AuthPresenter  presenter.AuthPresenter
}

type AuthInteractor interface {
	GetToken(u *model.User) (string, error)
}

func NewAuthInteractor(r repository.AuthRepository, p presenter.AuthPresenter) AuthInteractor {
	return &authInteractor{r, p}
}

func (us *authInteractor) GetToken(u *model.User) (string, error) {

	j, err := us.AuthRepository.Auth(u)

	if err != nil {
		return "nil", err
	}

	t, err := j.SignedString([]byte(utils.Secret))
	if err != nil {
		return t, err
	}

	return us.AuthPresenter.ResponseAuth(t), nil
}