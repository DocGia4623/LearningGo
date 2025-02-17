package service

import (
	"vietanh/gin-gorm-rest/models"
	"vietanh/gin-gorm-rest/repository"
)

type RoleServiceImpl struct {
	RoleRepository repository.RoleRepository
}

func NewRoleServiceImpl(roleRepository repository.RoleRepository) RoleService {
	return &RoleServiceImpl{RoleRepository: roleRepository}
}

func (service *RoleServiceImpl) FindAll() []models.Role {
	return service.RoleRepository.FindAll()
}

func (service *RoleServiceImpl) CreateRole(role models.Role) error {
	return service.RoleRepository.CreateRole(role)
}

func (service *RoleServiceImpl) CheckRoleExist(role string) (*models.Role, error) {
	result, err := service.RoleRepository.CheckRoleExist(role)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (service *RoleServiceImpl) FindBelongToPermission(permission string) ([]models.Role, error) {
	return service.RoleRepository.FindBelongToPermission(permission)
}
