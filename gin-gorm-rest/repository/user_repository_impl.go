package repository

import (
	"errors"
	"vietanh/gin-gorm-rest/data/request"
	"vietanh/gin-gorm-rest/helper"
	"vietanh/gin-gorm-rest/models"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(userId int) {
	var users models.User
	result := u.Db.Where("id = ?", userId).Delete(&users)
	helper.ErrorPanic(result.Error)
}

// FindAll implements UserRepository.
func (u *UserRepositoryImpl) FindAll() []models.User {
	var users []models.User
	result := u.Db.Find(&users)
	helper.ErrorPanic(result.Error)
	return users
}

// FindByID implements UserRepository.
func (u *UserRepositoryImpl) FindByID(userId int) (models.User, error) {
	var users models.User
	result := u.Db.Where("id = ?", userId).First(&users)
	if result != nil {
		return users, nil
	} else {
		return users, errors.New("user not found")
	}
}

// FindByUserName implements UserRepository.
func (u *UserRepositoryImpl) FindByUserName(userName string) (*models.User, error) {
	var user models.User
	result := u.Db.Where("user_name = ?", userName).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // Trả về nil nếu không tìm thấy bản ghi
		}
		helper.ErrorPanic(result.Error)
		return nil, result.Error // Trả về lỗi nếu có lỗi khác từ GORM
	}
	return &user, nil // Trả về con trỏ đến user nếu tìm thấy
}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(users models.User) {
	result := u.Db.Create(&users)
	helper.ErrorPanic(result.Error)
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(users models.User) {
	var UpdateUsers = request.UpdateUserRequest{
		Id:       users.ID,
		UserName: users.UserName,
		Password: users.Password,
		FullName: users.FullName,
		Email:    users.Email,
	}
	result := u.Db.Model(&users).Updates(UpdateUsers)
	helper.ErrorPanic(result.Error)
}

func (u *UserRepositoryImpl) FindIfUserHasRole(userID uint, roles []models.Role) bool {
	var count int64

	// Extract role names from the slice of roles
	var roleNames []string
	for _, role := range roles {
		roleNames = append(roleNames, role.Name)
	}

	result := u.Db.Table("users").
		Joins("JOIN user_roles ON users.id = user_roles.user_id").
		Joins("JOIN roles ON user_roles.role_id = roles.id").
		Where("users.id = ? AND roles.name IN ?", userID, roleNames).
		Count(&count)

	if result.Error != nil {
		helper.ErrorPanic(result.Error)
		return false
	}

	return count > 0
}
