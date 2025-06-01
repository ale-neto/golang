package controller

import (
	"net/http"

	err_rest "github.com/ale-neto/golang/src/config/err"
	"github.com/ale-neto/golang/src/config/logger"
	"github.com/ale-neto/golang/src/config/validation"
	"github.com/ale-neto/golang/src/controller/model/request"
	"github.com/ale-neto/golang/src/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

func (u *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init updateUser controller",
		zap.String("journey", "updateUser"),
	)
	var UserUpdateRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&UserUpdateRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "updateUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	id := c.Param("id")
	if _, err := bson.ObjectIDFromHex(id); err != nil {
		errRest := err_rest.NewBadRequestErr("Invalid ID, must be a hex value")
		c.JSON(errRest.Code, errRest)
	}

	domain := model.NewUserUpdateDomain(
		UserUpdateRequest.Name,
		UserUpdateRequest.Age,
	)
	err := u.service.UpdateUserService(id, domain)
	if err != nil {
		logger.Error(
			"Error trying to call updateUser service",
			err,
			zap.String("journey", "updateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"updateUser controller executed successfully",
		zap.String("id", id),
		zap.String("journey", "updateUser"))

	c.Status(http.StatusOK)
}
