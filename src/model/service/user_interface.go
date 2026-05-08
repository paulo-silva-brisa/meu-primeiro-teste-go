package service

import (
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/configurations/rest_err"
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/model"
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/model/repository" // Import novo
)

func NewUserDomainService(repo repository.UserRepository) UserDomainService {
	return &userDomainService{
		userRepository: repo,
	}
}

type userDomainService struct {
	userRepository repository.UserRepository
}
type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.Erro)
	UpdateUser(string, model.UserDomainInterface) *rest_err.Erro
	FindUser(string) (*model.UserDomainInterface, *rest_err.Erro)
	DeleteUser(string) *rest_err.Erro
}
