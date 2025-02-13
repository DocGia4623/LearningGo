package repository

import "vietanh/gin-gorm-rest/models"

type PermissionRepository interface {
	FindAll() []models.Permission
	CheckIfExist(permission string) (*models.Permission, error)
	CreatePermission(permission models.Permission) error
	DeletePermission(permission string)
}
