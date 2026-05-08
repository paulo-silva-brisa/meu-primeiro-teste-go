package service

import (
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/configurations/logger"
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/configurations/rest_err"
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,

) (model.UserDomainInterface, *rest_err.Erro) {
	logger.Info("Init CreateUser model", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)

	if err != nil {
		logger.Error("Error trying tocall repository,", err, zap.String("journey", "createUser"))
		return nil, err
	}
	logger.Info("Create User service executed successfull", zap.String(
		"userId", userDomainRepository.GetID()),
		zap.String("journey", "createUser"))
	return userDomainRepository, nil

}
