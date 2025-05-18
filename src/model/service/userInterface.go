package service

import (
	err_rest "github.com/ale-neto/golang/src/config/err"
	"github.com/ale-neto/golang/src/model"
	"github.com/ale-neto/golang/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	DeleteUser(string) *err_rest.Err
	CreateUser(model.UserDomainInterface) *err_rest.Err
	FindUser(string) (*model.UserDomainInterface, *err_rest.Err)
	UpdateUser(id string, user model.UserDomainInterface) *err_rest.Err
}
