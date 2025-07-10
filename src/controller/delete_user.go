package controller

import (
	"net/http"

	"github.com/ale-neto/golang/src/configuration/logger"
	"github.com/ale-neto/golang/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// DeleteUser deletes a user with the specified ID.
// @Summary Delete User
// @Description Deletes a user based on the ID provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "ID of the user to be deleted"
// @Success 200
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /user/delete/{id} [delete]
func (u *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init deleteUser controller",
		zap.String("journey", "deleteUser"),
	)

	id := c.Param("id")
	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid id, must be a hex value")
		c.JSON(errRest.Code, errRest)
		return
	}

	err := u.service.DeleteUser(id)
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
