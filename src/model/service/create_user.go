package service

import (
	"fmt"

	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/configurations/logger"
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/configurations/rest_err"
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,

) *rest_err.Erro {
	logger.Info("Init CreateUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetName())

	return nil
}
