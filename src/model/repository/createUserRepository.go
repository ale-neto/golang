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
	value, err := userDomain.GetJSONValue()
	if err != nil {
		err_rest.NewInternalServerErr("Erro ao converter usuário para JSON", err.Error())
	}
	u.dataBaseConnection.Collection(collectionName).InsertOne(context.Background(), value)
}
