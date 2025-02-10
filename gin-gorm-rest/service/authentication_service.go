package service

import (
	"context"
	"vietanh/gin-gorm-rest/data/request"
)

type AuthenticationService interface {
	Login(users request.LoginRequest) (string, error)
	Register(users request.CreateUserRequest) error
	Logout(ctx context.Context, token string) error
}
