package presenter

type authPresenter struct {
}

type AuthPresenter interface {
	ResponseAuth(t string) string
}

func NewAuthPresenter() AuthPresenter {
	return &authPresenter{}
}

func (up *authPresenter) ResponseAuth(t string) string {
	return t
}
