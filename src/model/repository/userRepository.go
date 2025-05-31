package repository

import (
	err_rest "github.com/ale-neto/golang/src/config/err"
	"github.com/ale-neto/golang/src/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
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
	) (model.UserDomainInterface, *err_rest.Err)
	DeleteUser(
		id string,
	) *err_rest.Err
	FindUserByID(
		id string,
	) (model.UserDomainInterface, *err_rest.Err)
	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *err_rest.Err)
	UpdateUser(
		userId string,
		userDomain model.UserDomainInterface,
	) *err_rest.Err
}
