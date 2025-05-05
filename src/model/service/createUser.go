package service

import (
	err_rest "github.com/ale-neto/golang/src/config/err"
	"github.com/ale-neto/golang/src/config/logger"
	"github.com/ale-neto/golang/src/model"
	"go.uber.org/zap"
)

func (u *userDomainService) CreateUser(user model.UserDomainInterface) *err_rest.Err {
	logger.Info("CreateUser function called", zap.String("function", "CreateUser"))
	user.EncryptPassword()
	logger.Info("User created successfully", zap.String("name", user.GetName()))
	return nil
}
