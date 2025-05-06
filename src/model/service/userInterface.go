package service

import (
	err_rest "github.com/ale-neto/golang/src/config/err"
	"github.com/ale-neto/golang/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *err_rest.Err
	UpdateUser(string, model.UserDomainInterface) *err_rest.Err
	FindUser(string) (*model.UserDomainInterface, *err_rest.Err)
	DeleteUser(string) *err_rest.Err
}
