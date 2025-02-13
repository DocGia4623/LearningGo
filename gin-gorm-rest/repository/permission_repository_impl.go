package repository

import (
	"vietanh/gin-gorm-rest/helper"
	"vietanh/gin-gorm-rest/models"

	"gorm.io/gorm"
)

type PermissionRepositoryImpl struct {
	Db *gorm.DB
}

func NewPermissionRepositoryImpl(Db *gorm.DB) PermissionRepository {
	return &PermissionRepositoryImpl{Db: Db}
}

func (r *PermissionRepositoryImpl) FindAll() []models.Permission {
	var permissions []models.Permission
	result := r.Db.Find(&permissions)
	helper.ErrorPanic(result.Error)
	return permissions
}
func (r *PermissionRepositoryImpl) CheckIfExist(permission string) (*models.Permission, error) {
	var permissions models.Permission
	result := r.Db.FirstOrCreate(&permissions, models.Permission{Name: permission})
	if result.Error != nil {
		helper.ErrorPanic(result.Error)
		return nil, result.Error
	}
	return &permissions, nil
}
func (r *PermissionRepositoryImpl) CreatePermission(permission models.Permission) error {
	result := r.Db.Create(&permission)
	helper.ErrorPanic(result.Error)
	return result.Error
}
func (r *PermissionRepositoryImpl) DeletePermission(permission string) {
	var permissions models.Permission
	result := r.Db.Where("name = ?", permission).Delete(&permissions)
	helper.ErrorPanic(result.Error)
}
