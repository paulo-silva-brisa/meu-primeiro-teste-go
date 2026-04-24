package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/configurations/logger"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {}
func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init findUserByEmail controller",
		zap.String("journey", "findUserByEmail"),
	)
}
