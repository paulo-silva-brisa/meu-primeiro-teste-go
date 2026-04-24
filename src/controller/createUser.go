package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/configurations/logger"
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/configurations/validation"
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/controller/model/request"
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/model"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zapcore.Field{
			Key:    "journey",
			String: "CreateUser",
		},
	)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validation user info", err)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)
	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))
	c.String(http.StatusOK, "")
}
