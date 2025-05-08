package controller

import (
	"net/http"

	"github.com/ale-neto/golang/src/config/logger"
	"github.com/ale-neto/golang/src/config/validation"
	"github.com/ale-neto/golang/src/controller/model/request"
	"github.com/ale-neto/golang/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (u *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("CreateUser function called",
		zap.String("function", "CreateUser"))

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("function", "CreateUser"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(userRequest.Name, userRequest.Password, userRequest.Email, userRequest.Age)

	if err := u.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User created successfully",
		zap.String("name", userRequest.Name))

	c.String(http.StatusOK, "User created successfully")

}
