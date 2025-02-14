package service

import "vietanh/gin-gorm-rest/models"

type PermissionService interface {
	FindAll() []models.Permission
	CheckIfExist(permission string) (*models.Permission, error)
	CreatePermission(permission models.Permission) error
	DeletePermission(permission string)
	CreatePermissionWithRole(permission string, role []string) error
}
