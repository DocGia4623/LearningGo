package service

import (
	"fmt"
	"vietanh/gin-gorm-rest/config"
	"vietanh/gin-gorm-rest/models"
	"vietanh/gin-gorm-rest/repository"
	"vietanh/gin-gorm-rest/utils"
)

type RefreshTokenServiceImpl struct {
	RefreshTokenRepository repository.RefreshTokenRepository
	RabbitMQService        RabbitMQService
}

func NewRefreshTokenServiceImpl(refreshTokenRepository repository.RefreshTokenRepository, rabbitMQService RabbitMQService) RefreshTokenService {
	return &RefreshTokenServiceImpl{
		RefreshTokenRepository: refreshTokenRepository,
		RabbitMQService:        rabbitMQService,
	}
}

func (a *RefreshTokenServiceImpl) RefreshToken(token string, signedKey string) (string, string, error) {
	// Kiểm tra refresh token có trong database không
	_, err := a.RefreshTokenRepository.FindByToken(token)
	if err != nil {
		return "", "", fmt.Errorf("refresh token is invalid")
	}
	// Kiểm tra token có hợp lệ không
	config, _ := config.LoadConfig()
	sub, expRaw, err := utils.ValidateRefreshToken(token, config.RefreshTokenSecret)
	if err != nil {
		a.RefreshTokenRepository.Delete(token)
		return "", "", fmt.Errorf("refresh token is invalid")
	}

	// Xóa refresh token cũ khỏi database
	a.DeleteToken(token)

	// Chuyển `expRaw` từ `interface{}` về `int64`
	expFloat, ok := expRaw.(float64)
	if !ok {
		return "", "", fmt.Errorf("invalid token expiration format")
	}
	exp := int64(expFloat) // Chuyển thành kiểu int64 (UNIX timestamp)

	// Tạo access token mới
	accessToken, err := utils.GenerateToken(config.AccessTokenExpiresIn, sub, config.AccessTokenSecret)
	if err != nil {
		return "", "", fmt.Errorf("cannot generate access token")
	}
	// Tạo refresh token mới
	newRefreshToken, err := utils.GenerateRefreshToken(exp, sub, config.RefreshTokenSecret)
	if err != nil {
		return "", "", fmt.Errorf("cannot generate refresh token")
	}

	// Lưu refresh token mới vào database
	a.SaveToken(models.RefreshToken{
		Token: newRefreshToken,
	})

	// Gửi sự kiện vào RabbitMQ (refresh token đã được tạo mới)
	go a.RabbitMQService.SendEvent("refresh_token_events", fmt.Sprintf(`{"token": "%s", "subject": "%s"}`, newRefreshToken, sub))
	return accessToken, newRefreshToken, nil
}

func (a *RefreshTokenServiceImpl) SaveToken(refreshToken models.RefreshToken) {
	a.RefreshTokenRepository.Save(refreshToken)
}
func (a *RefreshTokenServiceImpl) FindToken(token string) (*models.RefreshToken, error) {
	return a.RefreshTokenRepository.FindByToken(token)
}
func (a *RefreshTokenServiceImpl) DeleteToken(token string) {
	a.RefreshTokenRepository.Delete(token)
}
