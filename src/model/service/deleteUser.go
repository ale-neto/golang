package service

import (
	err_rest "github.com/ale-neto/golang/src/config/err"
	"github.com/ale-neto/golang/src/config/logger"
	"go.uber.org/zap"
)

func (u *userDomainService) DeleteUserService(id string) *err_rest.Err {
	logger.Info("Init deleteUser model.",
		zap.String("journey", "deleteUser"))

	err := u.userRepository.DeleteUser(id)
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
