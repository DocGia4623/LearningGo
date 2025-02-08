package repository

import "vietanh/gin-gorm-rest/models"

type UserRepository interface {
	Save(users models.User)
	Update(users models.User)
	Delete(userId int)
	FindByID(userId int) (models.User, error)
	FindAll() []models.User
	FindByUserName(userName string) (*models.User, error)
}
