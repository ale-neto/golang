package controller

import (
	"net/http"

	err_rest "github.com/ale-neto/golang/src/config/err"
	"github.com/ale-neto/golang/src/config/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

func (u *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init deleteUser controller",
		zap.String("journey", "deleteUser"),
	)

	id := c.Param("id")
	if _, err := bson.ObjectIDFromHex(id); err != nil {
		errRest := err_rest.NewBadRequestErr("Invalid id, must be a hex value")
		c.JSON(errRest.Code, errRest)
		return
	}

	err := u.service.DeleteUserService(id)
	if err != nil {
		logger.Error(
			"Error trying to call deleteUser service",
			err,
			zap.String("journey", "deleteUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"deleteUser controller executed successfully",
		zap.String("id", id),
		zap.String("journey", "deleteUser"))

	c.Status(http.StatusOK)
}
