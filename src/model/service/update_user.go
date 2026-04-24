package service

import (
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/configurations/rest_err"
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/model"
)

func (ud *userDomainService) UpdateUser(
	userID string,
	userDomain model.UserDomainInterface,
) *rest_err.Erro {
	return nil
}
