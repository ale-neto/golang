package repository

import (
	err_rest "github.com/ale-neto/golang/src/config/err"
	"github.com/ale-neto/golang/src/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewUserRepository(dataBaseConnection *mongo.Database) UserRepository {
	return *userRepository{
		database,
	}
}

type userRepository struct {
	dataBaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *err_rest.Err)
}
