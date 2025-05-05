package service

import (
	err_rest "github.com/ale-neto/golang/src/config/err"
	"github.com/ale-neto/golang/src/model"
)

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *err_rest.Err
	UpdateUser(string, model.UserDomainInterface) *err_rest.Err
	FindUser(string) (*userDomain, *err_rest.Err)
	DeleteUser(string) *err_rest.Err
}
