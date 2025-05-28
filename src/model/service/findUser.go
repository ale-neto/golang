package service

import (
	err_rest "github.com/ale-neto/golang/src/config/err"
	"github.com/ale-neto/golang/src/config/logger"
	"github.com/ale-neto/golang/src/model"
	"go.uber.org/zap"
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

func (u *userDomainService) FindUserByEmailService(
	email string,
) (model.UserDomainInterface, *err_rest.Err) {
	logger.Info("Init findUserByEmail services.",
		zap.String("journey", "findUserById"))

	return u.userRepository.FindUserByEmail(email)
}
