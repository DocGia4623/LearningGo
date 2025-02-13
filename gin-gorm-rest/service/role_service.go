package service

import "vietanh/gin-gorm-rest/models"

type RoleService interface {
	FindAll() []models.Role
	CreateRole(role models.Role) error
	CheckRoleExist(role string) (*models.Role, error)
}
