package main

import (
	"github.com/ale-neto/golang/src/controller"
	"github.com/ale-neto/golang/src/model/repository"
	"github.com/ale-neto/golang/src/model/service"
)

func initDependencies() (controller.UserControllerInterface, error) {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

}
