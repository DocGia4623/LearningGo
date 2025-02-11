package service

import "vietanh/gin-gorm-rest/models"

type RefreshTokenService interface {
	SaveToken(models.RefreshToken)
	DeleteToken(token string)
	FindToken(token string) (*models.RefreshToken, error)
	RefreshToken(token string, signedKey string) (string, string, error)
}
