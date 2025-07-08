package service

import (
	"github.com/ale-neto/golang/src/configuration/logger"
	"github.com/ale-neto/golang/src/configuration/rest_err"
	"github.com/ale-neto/golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(
	id string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init updateUser model.",
		zap.String("journey", "updateUser"))

	err := ud.userRepository.UpdateUser(id, userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "updateUser"))
		return err
	}

	logger.Info(
		"updateUser service executed successfully",
		zap.String("id", id),
		zap.String("journey", "updateUser"))
	return nil
}
