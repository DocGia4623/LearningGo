package service

import (
	"fmt"
	"vietanh/gin-gorm-rest/models"
	"vietanh/gin-gorm-rest/repository"

	"gorm.io/gorm"
)

type PermissionServiceImpl struct {
	PermissionRepository repository.PermissionRepository
	RoleRepository       repository.RoleRepository
}

func NewPermissionServiceImpl(permissionRepository repository.PermissionRepository, roleRepository repository.RoleRepository) PermissionService {
	return &PermissionServiceImpl{
		PermissionRepository: permissionRepository,
		RoleRepository:       roleRepository}
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

func (a *PermissionServiceImpl) CreatePermissionWithRole(permissionName string, roleNames []string) error {

	// Tìm hoặc tạo Permission
	// Kiểm tra nếu permission đã tồn tại
	permission, err := a.PermissionRepository.CheckIfExist(permissionName)
	if err != nil && err != gorm.ErrRecordNotFound {
		return fmt.Errorf("lỗi khi tìm permission: %v", err)
	}

	// Nếu permission chưa tồn tại, tạo mới
	if permission == nil {
		newPermission := models.Permission{Name: permissionName}
		if err := a.CreatePermission(newPermission); err != nil {
			return fmt.Errorf("không thể tạo permission: %v", err)
		}
		// Lấy lại permission vừa tạo
		permission, err = a.PermissionRepository.CheckIfExist(permissionName)
		if err != nil {
			return fmt.Errorf("lỗi khi lấy lại permission: %v", err)
		}
	}
	// Duyệt qua danh sách roleNames
	for _, roleName := range roleNames {

		role, err := a.RoleRepository.CheckRoleExist(roleName)
		if err != nil && err != gorm.ErrRecordNotFound {
			return fmt.Errorf("lỗi khi tìm role: %v", err)
		}
		// Tìm hoặc tạo Role
		if role == nil {
			newRole := models.Role{Name: roleName}
			if err := a.RoleRepository.CreateRole(newRole); err != nil {
				return fmt.Errorf("không thể tạo role: %v", err)
			}
			// Lấy lại role vừa tạo
			role, err = a.RoleRepository.CheckRoleExist(roleName)
			if err != nil {
				return fmt.Errorf("lỗi khi lấy lại role: %v", err)
			}
		}
		// Kiểm tra nếu đã có RolePermission chưa
		rolePermission, err := a.RoleRepository.CheckRolePermission(role.ID, permission.ID)
		if err != nil && err != gorm.ErrRecordNotFound {
			return fmt.Errorf("lỗi khi kiểm tra role_permission: %v", err)
		}

		// Nếu chưa có, tạo mới liên kết
		if rolePermission == nil {
			if err := a.RoleRepository.CreateRolePermission(role.ID, permission.ID); err != nil {
				return fmt.Errorf("không thể gán permission vào role: %v", err)
			}
		}
	}
	return nil
}
