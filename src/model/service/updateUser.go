package service

import (
	err_rest "github.com/ale-neto/golang/src/config/err"
	"github.com/ale-neto/golang/src/config/logger"
	"github.com/ale-neto/golang/src/model"
	"go.uber.org/zap"
)

func (u *userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *err_rest.Err {
	logger.Info("Init updateUser model.",
		zap.String("journey", "updateUser"))

	err := u.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "updateUser"))
		return err
	}

	logger.Info(
		"updateUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))
	return nil
}
