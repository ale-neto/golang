package routes

import (
	"github.com/ale-neto/golang/src/controller"
	"github.com/ale-neto/golang/src/model"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.POST("/users", model.VerifyTokenMiddleware, userController.CreateUser)
	r.GET("/users/by/id/:id", model.VerifyTokenMiddleware, userController.FindUserByID)
	r.GET("/users/by/email/:email", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.PUT("/users/:id", model.VerifyTokenMiddleware, userController.UpdateUser)
	r.DELETE("/users/:id", model.VerifyTokenMiddleware, userController.DeleteUser)
	r.POST("/login", userController.LoginUser)
}
