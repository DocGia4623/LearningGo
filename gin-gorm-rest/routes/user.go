package routes

import (
	"vietanh/gin-gorm-rest/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	userRoutes := router.Group("/user")
	{
		userRoutes.GET("/", controller.GetUsers)
		userRoutes.POST("/", controller.CreateUser)
		userRoutes.DELETE("/:id", controller.DeleteUser)
		userRoutes.PUT("/:id", controller.UpdateUser)
		userRoutes.GET("/:id", controller.GetUser)
	}
}
