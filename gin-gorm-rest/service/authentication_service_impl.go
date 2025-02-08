package service

import (
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
func (a *AuthenticationServiceImpl) Login(users request.LoginRequest) (string, error) {
	// Find username in the database
	new_user, user_err := a.UserRepository.FindByUserName(users.UserName)
	if user_err != nil {
		return "", errors.New("invalid username or password")
	}

	config, _ := config.LoadConfig()

	// Verify password
	verify_err := utils.VerifyPassword(new_user.Password, users.Password)
	if verify_err != nil {
		return "", errors.New("invalid username or password")
	}

	// Generate token
	token, err_token := utils.GenerateToken(config.TokenExpiresIn, new_user.ID, config.TokenSecret)
	helper.ErrorPanic(err_token)

	return token, nil
}

// Register implements AuthenticationService.
func (a *AuthenticationServiceImpl) Register(users request.CreateUserRequest) error {
	// Kiểm tra xem username đã tồn tại chưa
	user, user_err := a.UserRepository.FindByUserName(users.UserName)
	if user_err != nil && user_err.Error() != "user not found" {
		// Nếu có lỗi ngoài việc không tìm thấy user, trả về lỗi đó
		return user_err
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
	}

	// Lưu user vào database
	a.UserRepository.Save(newUser)
	return nil // Trả về nil nếu không có lỗi
}
