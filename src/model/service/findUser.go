package service

import (
	err_rest "github.com/ale-neto/golang/src/config/err"
	"github.com/ale-neto/golang/src/model"
)

func (u *userDomainService) FindUserByIDService(id string) (model.UserDomainInterface, *err_rest.Err) {
	userDomain, err := u.userRepository.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	if userDomain == nil {
		return nil, err_rest.NewNotFoundErr("User not found")
	}

	return userDomain, nil
}
