package routes

import (
	"github.com/ale-neto/golang/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("/users/email/:email", userController.CreateUser)
	r.POST("/users", userController.GetUsersByEmail)
	r.POST("/users", userController.GetUsersById)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)
}
