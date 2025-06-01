package routes

import (
	"github.com/ale-neto/golang/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.POST("/users", userController.CreateUser)
	r.GET("/users/by/id/:id", userController.FindUserByID)
	r.GET("/users/by/email/:userEmail", userController.FindUserByEmail)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)
}
