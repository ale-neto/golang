package model

import (
	err_rest "github.com/ale-neto/golang/src/config/err"
	"github.com/ale-neto/golang/src/config/logger"
	"go.uber.org/zap"
)

func (u *userDomain) CreateUser(user userDomain) *err_rest.Err {
	logger.Info("CreateUser function called", zap.String("function", "CreateUser"))
	u.EncryptPassword()
	logger.Info("User created successfully", zap.String("name", user.Name))
	return nil
}
