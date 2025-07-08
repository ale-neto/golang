package service

import (
	"github.com/ale-neto/golang/src/configuration/logger"
	"github.com/ale-neto/golang/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(
	id string) *rest_err.RestErr {

	logger.Info("Init deleteUser model.",
		zap.String("journey", "deleteUser"))

	err := ud.userRepository.DeleteUser(id)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "deleteUser"))
		return err
	}

	logger.Info(
		"deleteUser service executed successfully",
		zap.String("id", id),
		zap.String("journey", "deleteUser"))
	return nil
}
