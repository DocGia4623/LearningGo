package repository

import "vietanh/gin-gorm-rest/models"

type RoleRepository interface {
	FindAll() []models.Role
	CreateRole(role models.Role) error
	CheckRoleExist(role string) (*models.Role, error)
	CheckRolePermission(roleID, permissionID uint) (*models.RolePermission, error)
	CreateRolePermission(roleID, permissionID uint) error
}
