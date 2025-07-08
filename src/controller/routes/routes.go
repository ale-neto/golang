package routes

import (
	"github.com/ale-neto/golang/src/controller"
	"github.com/ale-neto/golang/src/model"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface) {

	r.GET("/user/id/:id", model.VerifyTokenMiddleware, userController.FindUserByID)
	r.GET("/user/email/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/user/create", userController.CreateUser)
	r.PUT("/user/update/:id", model.VerifyTokenMiddleware, userController.UpdateUser)
	r.DELETE("/user/delete/:id", model.VerifyTokenMiddleware, userController.DeleteUser)

	r.POST("/login", userController.LoginUser)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
