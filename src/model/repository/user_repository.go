package repository

import (
	"github.com/ale-neto/golang/src/configuration/rest_err"
	"github.com/ale-neto/golang/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUser(
		id string,
		userDomain model.UserDomainInterface,
	) *rest_err.RestErr

	DeleteUser(
		id string,
	) *rest_err.RestErr

	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailAndPassword(
		email string,
		password string,
	) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)
}
