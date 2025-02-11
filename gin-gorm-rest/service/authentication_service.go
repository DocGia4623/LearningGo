package service

import (
	"context"
	"vietanh/gin-gorm-rest/data/request"
)

type AuthenticationService interface {
	Login(users request.LoginRequest) (string, string, error)
	Register(users request.CreateUserRequest) error
	Logout(ctx context.Context, refreshToken string, accessToken string) error
}
