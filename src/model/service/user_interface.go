package service

import (
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/configurations/rest_err"
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.Erro
	UpdateUser(string, model.UserDomainInterface) *rest_err.Erro
	FindUser(string) (*model.UserDomainInterface, *rest_err.Erro)
	DeleteUser(string) *rest_err.Erro
}
