package controller

import (
	"net/http"

	"github.com/ale-neto/golang/src/config/logger"
	"github.com/ale-neto/golang/src/config/validation"
	"github.com/ale-neto/golang/src/controller/model/request"
	"github.com/ale-neto/golang/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("CreateUser function called",
		zap.String("function", "CreateUser"))

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("function", "CreateUser"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	response := response.UserResponse{
		Name:  userRequest.Name,
		Email: userRequest.Email,
		Age:   userRequest.Age,
	}

	logger.Info("User created successfully",
		zap.String("name", userRequest.Name))

	c.JSON(http.StatusOK, response)

}
