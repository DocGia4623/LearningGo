package controller

import (
	"net/http"
	"vietanh/gin-gorm-rest/models"
	"vietanh/gin-gorm-rest/service"

	"github.com/gin-gonic/gin"
)

type PermissionController struct {
	PermissionService service.PermissionService
}

func NewPermissionController(permissionService service.PermissionService) PermissionController {
	return PermissionController{PermissionService: permissionService}
}

func (controller *PermissionController) GetPermission(c *gin.Context) {
	permissions := controller.PermissionService.FindAll()
	c.JSON(http.StatusOK, permissions)
}
func (controller *PermissionController) CreatePermission(c *gin.Context) {
	var permission models.Permission
	c.BindJSON(&permission)
	controller.PermissionService.CreatePermission(permission)
	c.JSON(http.StatusOK, permission)
}
func (controller *PermissionController) DeletePermission(c *gin.Context) {
	permission := c.Param("permission")
	controller.PermissionService.DeletePermission(permission)
	c.JSON(http.StatusOK, gin.H{"message": "Permission deleted"})
}
func (controller *PermissionController) CheckIfExist(c *gin.Context) {
	permission := c.Param("permission")
	result, _ := controller.PermissionService.CheckIfExist(permission)
	c.JSON(http.StatusOK, result)
}
