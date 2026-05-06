package model

import "github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/configurations/rest_err"

type userDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string
	GetID() string

	SetID(string)

	EncryptPassword()
	GenerateToken() (string, *rest_err.Erro)
}

func newUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}

}

func NewUserLoginDomain(
	email, password string,
) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: password,
	}
}

func NewUserUpdateDomain(
	name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		Name: name,
		Age:  age,
	}
}
