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
