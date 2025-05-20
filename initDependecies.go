package main

import (
	"fmt"

	"github.com/ale-neto/golang/src/controller"
	"github.com/ale-neto/golang/src/model/repository"
	"github.com/ale-neto/golang/src/model/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	fmt.Println("Iniciando repositório...")
	repo := repository.NewUserRepository(database)

	fmt.Println("Iniciando serviço...")
	service := service.NewUserDomainService(repo)

	fmt.Println("Iniciando controller...")
	return controller.NewUserControllerInterface(service)
}
