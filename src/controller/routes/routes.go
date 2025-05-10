package routes

import (
	"github.com/ale-neto/golang/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.POST("/users", userController.CreateUser)
	r.GET("/users/email/:email", userController.GetUsersByEmail)
	r.GET("/users/id/:id", userController.GetUsersById)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)
}
