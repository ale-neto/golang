package controller

import (
	"github.com/ale-neto/golang/src/model/service"
	"github.com/gin-gonic/gin"
)

type UserControllerInterface struct {
	DeleteUser (*gin.Context)
	CreateUser (*gin.Context)
	GetUser    (*gin.Context)
	UpdateUser (*gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
