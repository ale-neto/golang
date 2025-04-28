package controller

import (
	"fmt"

	"github.com/ale-neto/golang/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := err.NewBadRequestValidationErr(fmt.Sprintf("The request body is invalid, erro=%s", err))
		c.JSON(restErr.code, restErr)
		return
	}
}
