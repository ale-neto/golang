package main

import (
	"github.com/ale-neto/golang/src/controller"
	"github.com/ale-neto/golang/src/model/repository"
	"github.com/ale-neto/golang/src/model/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
