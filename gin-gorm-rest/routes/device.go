package routes

import (
	"vietanh/gin-gorm-rest/controller"

	"github.com/gin-gonic/gin"
)

func DeviceRoute(router *gin.Engine) {
	DeviceRoute := router.Group("/device")
	{
		DeviceRoute.GET("/", controller.GetDevices)
		DeviceRoute.POST("/", controller.CreateDevice)
		DeviceRoute.DELETE("/:id", controller.DeleteDevice)
		DeviceRoute.PUT("/:id", controller.UpdateDevice)
		DeviceRoute.GET("/:id", controller.GetDevice)
	}

}
