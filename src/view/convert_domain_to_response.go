package view

import (
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/controller/model/response"
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
