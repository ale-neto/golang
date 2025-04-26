package routes

import (
	"github.com/ale-neto/golang/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/users/id/:id", controller.GetUsersById)
	r.GET("/users/email/:email", controller.GetUsersByEmail)
	r.POST("/users", controller.CreateUser)
	r.PUT("/users/:id", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)
}
