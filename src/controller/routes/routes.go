package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/users/:id", controller)

	r.GET("/users/:email", controller.getUsers)
	r.POST("/users", controller.createUser)
	// e.g., r.PUT("/users/:id", updateUser)
	// e.g., r.DELETE("/users/:id", deleteUser)
}
