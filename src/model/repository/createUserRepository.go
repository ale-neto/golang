package repository

import (
	"context"
	"os"

	err_rest "github.com/ale-neto/golang/src/config/err"
	"github.com/ale-neto/golang/src/config/logger"
	"github.com/ale-neto/golang/src/model"
)

var (
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func (u *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *err_rest.Err) {
	logger.Info("CreateUser - Iniciando criação de usuário")
	collectionName := os.Getenv(MONGODB_USER_DB)

	colletion := u.dataBaseConnection.Collection(collectionName)

	value, err := userDomain.GetJSONValue()
	if err != nil {
		return nil, err_rest.NewInternalServerError("Erro ao converter usuário para JSON" + err.Error())
	}

	result, err := colletion.InsertOne(context.TODO(), value)
	if err != nil {
		return nil, err_rest.NewInternalServerError("Erro ao converter usuário para JSON" + err.Error())
	}

	userDomain.SetID(result.InsertedID.(string))

	return userDomain, nil
}
