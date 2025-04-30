package controller

import (
	"fmt"
	"log"

	"github.com/ale-neto/golang/src/config/validation"
	"github.com/ale-neto/golang/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("The request body is invalid, erro=%s\n", err.Error())
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
}
