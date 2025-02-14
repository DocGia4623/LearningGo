package controller

import (
	"net/http"
	"vietanh/gin-gorm-rest/data/request"
	"vietanh/gin-gorm-rest/data/response"
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

func (controller *PermissionController) CreatePermissionWithRole(c *gin.Context) {
	var req request.PermissionRoleRequest
	// Kiểm tra lỗi khi parse JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}
	// Gọi service để tạo permission với role
	if err := controller.PermissionService.CreatePermissionWithRole(req.Name, req.Roles); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Trả về response thành công
	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "ok",
		Message: "Permission created successfully",
		Data:    req,
	}
	c.JSON(http.StatusOK, webResponse)
}
