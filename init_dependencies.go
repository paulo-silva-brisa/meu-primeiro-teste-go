package main

import (
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/controller"
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/model/repository"
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/model/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
