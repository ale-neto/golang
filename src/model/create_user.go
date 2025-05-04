package model

import (
	err_rest "github.com/ale-neto/golang/src/config/err"
	"github.com/ale-neto/golang/src/config/logger"
)

func (*UserDomain) CreateUser(user UserDomain) *err_rest.Err {
	logger.Info("CreateUser function called")
	return nil
}
