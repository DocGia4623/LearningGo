package controller

import (
	"net/http"
	"vietanh/gin-gorm-rest/data/request"
	"vietanh/gin-gorm-rest/data/response"
	"vietanh/gin-gorm-rest/helper"
	"vietanh/gin-gorm-rest/service"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	AuthenticationService service.AuthenticationService
}

func NewAuthenticationController(authenticationService service.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{
		AuthenticationService: authenticationService,
	}
}
func (controller *AuthenticationController) Login(c *gin.Context) {
	LoginRequest := request.LoginRequest{}
	err := c.ShouldBindJSON(&LoginRequest)
	helper.ErrorPanic(err)

	token, err := controller.AuthenticationService.Login(LoginRequest)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "bad request",
			Message: "Invalid username or password",
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}
	loginResponse := response.LoginResponse{
		TokenType: "Bearer Token",
		Token:     token,
	}
	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "ok",
		Message: "Login success",
		Data:    loginResponse,
	}
	c.JSON(http.StatusOK, webResponse)
}

func (controller *AuthenticationController) Register(c *gin.Context) {
	CreateUserRequest := request.CreateUserRequest{}
	err := c.ShouldBindJSON(&CreateUserRequest)
	helper.ErrorPanic(err)

	controller.AuthenticationService.Register(CreateUserRequest)
	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "ok",
		Message: "Register success",
		Data:    nil,
	}
	c.JSON(http.StatusOK, webResponse)
}
