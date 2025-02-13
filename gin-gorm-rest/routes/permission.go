package routes

import (
	"vietanh/gin-gorm-rest/controller"

	"github.com/gin-gonic/gin"
)

func PermissionRoute(PermissionController controller.PermissionController, Router *gin.Engine) {
	permission := Router.Group("/permission")
	{
		permission.GET("/", PermissionController.GetPermission)
		permission.POST("/", PermissionController.CreatePermission)
		permission.POST("/check", PermissionController.CheckIfExist)
		permission.DELETE("/:id", PermissionController.DeletePermission)
	}
}
