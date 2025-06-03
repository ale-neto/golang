package service

import (
	err_rest "github.com/ale-neto/golang/src/config/err"
	"github.com/ale-neto/golang/src/config/logger"
	"github.com/ale-neto/golang/src/model"
	"go.uber.org/zap"
)

func (u *userDomainService) CreateUserService(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *err_rest.Err) {

	logger.Info("Init createUser model.",
		zap.String("journey", "createUser"))

	user, _ := u.FindUserByEmailService(userDomain.GetEmail())
	if user != nil {
		return nil, err_rest.NewBadRequestErr("Email is already resgistered im another account")
	}

	userDomain.EncryptPassword()
	userDomainRepository, err := u.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info(
		"CreateUser service executed successfully",
		zap.String("id", userDomainRepository.GetID()),
		zap.String("journey", "createUser"))
	return userDomainRepository, nil
}
