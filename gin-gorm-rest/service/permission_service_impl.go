package service

import (
	"vietanh/gin-gorm-rest/models"
	"vietanh/gin-gorm-rest/repository"
)

type PermissionServiceImpl struct {
	PermissionRepository repository.PermissionRepository
}

func NewPermissionService(permissionRepository repository.PermissionRepository) PermissionService {
	return &PermissionServiceImpl{PermissionRepository: permissionRepository}
}

func (a *PermissionServiceImpl) FindAll() []models.Permission {
	return a.PermissionRepository.FindAll()
}

func (a *PermissionServiceImpl) CheckIfExist(permission string) (*models.Permission, error) {
	result, err := a.PermissionRepository.CheckIfExist(permission)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (a *PermissionServiceImpl) CreatePermission(permission models.Permission) error {
	return a.PermissionRepository.CreatePermission(permission)
}
func (a *PermissionServiceImpl) DeletePermission(permission string) {
	a.PermissionRepository.DeletePermission(permission)
}
