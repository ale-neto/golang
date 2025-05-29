package repository

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"

	err_rest "github.com/ale-neto/golang/src/config/err"
	"github.com/ale-neto/golang/src/config/logger"
	"github.com/ale-neto/golang/src/model"
	"github.com/ale-neto/golang/src/model/repository/entity/converter"
	"go.uber.org/zap"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *err_rest.Err) {
	logger.Info("Init createUser repository",
		zap.String("journey", "createUser"))

	collectionName := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collectionName)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error trying to create user",
			err,
			zap.String("journey", "createUser"))
		return nil, err_rest.NewInternalServerError(err.Error())
	}

	oid, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		logger.Error("Erro ao criar usuário", err, zap.String("journey", "createUser"))

		return nil, err_rest.NewInternalServerError("Error casting InsertedID to ObjectID")
	}

	value.ID = oid

	logger.Info("CreateUser repository executed successfully",
		zap.String("id", value.ID.Hex()),
		zap.String("journey", "createUser"))

	return converter.ConvertEntityToDomain(*value), nil
}
