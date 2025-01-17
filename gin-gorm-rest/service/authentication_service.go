package service

import "vietanh/gin-gorm-rest/data/request"

type AuthenticationService interface {
	Login(users request.LoginRequest) (string, error)
	Register(users request.CreateUserRequest)
}
