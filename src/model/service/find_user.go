package service

import (
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/configurations/logger"
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/configurations/rest_err"
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUser(string) (*model.UserDomainInterface, *rest_err.Erro) {
	logger.Info("Init FindUser model", zap.String("journey", "FindUser"))

	return nil, nil
}
