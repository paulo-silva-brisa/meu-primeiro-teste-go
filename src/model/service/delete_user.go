package service

import (
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/configurations/logger"
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/configurations/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(string) *rest_err.Erro {
	logger.Info("Init DeleteUser model", zap.String("journey", "DeleteUser"))
	return nil
}
