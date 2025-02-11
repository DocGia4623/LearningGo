package service

import (
	"context"
	"time"
	"vietanh/gin-gorm-rest/config"
	"vietanh/gin-gorm-rest/data/request"
	"vietanh/gin-gorm-rest/helper"
	"vietanh/gin-gorm-rest/models"
	"vietanh/gin-gorm-rest/repository"
	"vietanh/gin-gorm-rest/utils"

	"errors"

	"github.com/go-playground/validator/v10"
)

type AuthenticationServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewAuthenticationServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

// Login implements AuthenticationService.
func (a *AuthenticationServiceImpl) Login(users request.LoginRequest) (string, string, error) {
	// Find username in the database
	login_user, user_err := a.UserRepository.FindByUserName(users.UserName)
	if user_err != nil {
		return "", "", errors.New("invalid username or password")
	}
	if login_user == nil {
		return "", "", errors.New("invalid username or password") // Xử lý khi user không tồn tại
	}

	config, _ := config.LoadConfig()

	// Verify password
	verify_err := utils.VerifyPassword(login_user.Password, users.Password)
	if verify_err != nil {
		return "", "", errors.New("invalid username or password")
	}

	// Generate access token
	accessToken, err_token := utils.GenerateToken(config.AccessTokenExpiresIn, login_user.ID, config.AccessTokenSecret)
	helper.ErrorPanic(err_token)

	// Generate refresh token
	refreshToken, err_refresh := utils.GenerateToken(config.RefreshTokenExpiresIn, login_user.ID, config.RefreshTokenSecret)
	helper.ErrorPanic(err_refresh)

	return refreshToken, accessToken, nil
}

func (a *AuthenticationServiceImpl) Logout(ctx context.Context, refreshToken string, accessToken string) error {

	// Lưu token vào redis
	expiration := time.Hour
	err := config.RedisClient.Set(ctx, accessToken, "logout", expiration).Err()
	if err != nil {
		return err
	}
	// Xóa refresh token khỏi database
	RefreshTokenService := NewRefreshTokenServiceImpl(repository.NewRefreshTokenRepositoryImpl(config.DB))
	RefreshTokenService.DeleteToken(refreshToken)
	return nil
}

// Register implements AuthenticationService.
func (a *AuthenticationServiceImpl) Register(users request.CreateUserRequest) error {
	// Kiểm tra xem username đã tồn tại chưa
	user, user_err := a.UserRepository.FindByUserName(users.UserName)
	if user_err != nil {
		return user_err // Nếu có lỗi khi tìm kiếm user, trả về lỗi đó
	}
	if user != nil {
		// Nếu user đã tồn tại trong cơ sở dữ liệu
		return errors.New("Username already exists")
	}

	// Hash mật khẩu
	hashedPassword, err := utils.HashPassword(users.Password)
	if err != nil {
		return err // Nếu có lỗi khi hash mật khẩu, trả về lỗi
	}

	// Tạo user mới
	newUser := models.User{
		UserName: users.UserName,
		Password: hashedPassword,
		Email:    users.Email,
		FullName: users.FullName,
		Role:     users.Role,
	}

	// Lưu user vào database
	a.UserRepository.Save(newUser)
	return nil // Trả về nil nếu không có lỗi
}
