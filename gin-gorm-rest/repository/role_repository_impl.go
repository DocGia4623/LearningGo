package repository

import (
	"errors"
	"vietanh/gin-gorm-rest/helper"
	"vietanh/gin-gorm-rest/models"

	"gorm.io/gorm"
)

type RoleRepositoryimpl struct {
	Db *gorm.DB
}

func NewRoleRepositoryimpl(Db *gorm.DB) RoleRepository {
	return &RoleRepositoryimpl{Db: Db}
}

// FindAll implements RoleRepository.
func (r *RoleRepositoryimpl) FindAll() []models.Role {
	var roles []models.Role
	result := r.Db.Find(&roles)
	helper.ErrorPanic(result.Error)
	return roles
}

// CreateRole implements RoleRepository.
func (r *RoleRepositoryimpl) CreateRole(role models.Role) error {
	result := r.Db.Create(&role)
	helper.ErrorPanic(result.Error)
	return result.Error
}

// CheckRoleExist implements RoleRepository.
func (r *RoleRepositoryimpl) CheckRoleExist(role string) (*models.Role, error) {
	var roles models.Role
	result := r.Db.FirstOrCreate(&roles, models.Role{Name: role})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		helper.ErrorPanic(result.Error)
		return nil, result.Error
	}
	return &roles, nil
}

func (r *RoleRepositoryimpl) CheckRolePermission(roleID, permissionID uint) (*models.RolePermission, error) {
	var rolePermission models.RolePermission
	err := r.Db.Where("role_id = ? AND permission_id =?", roleID, permissionID).First(&rolePermission).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &rolePermission, nil
}
func (r *RoleRepositoryimpl) CreateRolePermission(roleID, permissionID uint) error {
	rolePermission := models.RolePermission{
		RoleID:       roleID,
		PermissionID: permissionID,
	}
	result := r.Db.Create(&rolePermission)
	helper.ErrorPanic(result.Error)
	return result.Error
}

func (r *RoleRepositoryimpl) FindBelongToPermission(permission string) ([]models.Role, error) {
	var roles []models.Role
	r.Db = r.Db.Debug() // Bật chế độ Debug để xem câu truy vấn SQL
	result := r.Db.Joins("JOIN role_permissions ON roles.id = role_permissions.role_id").
		Joins("JOIN permissions ON permissions.id = role_permissions.permission_id").
		Where("permissions.name = ?", permission).
		Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}
