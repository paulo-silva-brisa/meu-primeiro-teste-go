package repository

import (
	"context"
	"os"

	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/configurations/logger"
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/configurations/rest_err"
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/model"
)

const (
	MONGO_USER_DB = "MONGODB_USER_DB"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.Erro) {

	logger.Info("Init createUser repository")
	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	value, err := userDomain.GetJSONValue()
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	userDomain.SetID(result.InsertedID.(string))

	return userDomain, nil
}
