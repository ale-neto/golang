package repository

import (
	"context"
	"os"

	err_rest "github.com/ale-neto/golang/src/config/err"
	"github.com/ale-neto/golang/src/config/logger"
	"go.mongodb.org/mongo-driver/v2/bson"

	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(
	id string,
) *err_rest.Err {
	logger.Info("Init deleteUser repository",
		zap.String("journey", "deleteUser"))

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userIdHex, err := bson.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Error converting id to ObjectID",
			err,
			zap.String("journey", "deleteUser"))
	}

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error trying to delete user",
			err,
			zap.String("journey", "deleteUser"))
		return err_rest.NewInternalServerError(err.Error())
	}

	logger.Info(
		"deleteUser repository executed successfully",
		zap.String("id", id),
		zap.String("journey", "deleteUser"))

	return nil
}
