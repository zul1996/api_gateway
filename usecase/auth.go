package usecase

type LoginInterface interface {
	Autentikasi(username string, password string) bool
}

type Login struct {
}

func (b *Login) Autentikasi(username string, password string) bool {
	return username == "admin" && password == "admin123"
}

func NewLogin(loginType string) LoginInterface {
	if loginType == "basic" {
		return &Login{}
	}
	return nil
}
