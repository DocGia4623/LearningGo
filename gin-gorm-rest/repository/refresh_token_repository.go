package repository

import "vietanh/gin-gorm-rest/models"

type RefreshTokenRepository interface {
	Save(refreshToken models.RefreshToken)
	FindByToken(token string) (*models.RefreshToken, error)
	Delete(token string)
}
