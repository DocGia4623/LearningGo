package controller

import (
	"context"
	"net/http"
	"vietanh/gin-gorm-rest/config"
	"vietanh/gin-gorm-rest/data/request"
	"vietanh/gin-gorm-rest/data/response"
	"vietanh/gin-gorm-rest/helper"
	"vietanh/gin-gorm-rest/models"
	"vietanh/gin-gorm-rest/service"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	AuthenticationService service.AuthenticationService
	RefreshTokenService   service.RefreshTokenService
}

func NewAuthenticationController(authenticationService service.AuthenticationService, refreshTokenService service.RefreshTokenService) *AuthenticationController {
	return &AuthenticationController{
		AuthenticationService: authenticationService,
		RefreshTokenService:   refreshTokenService,
	}
}
func (controller *AuthenticationController) RefreshToken(c *gin.Context) {
	// Lấy refresh token từ cookie
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusUnauthorized,
			Status:  "unauthorized",
			Message: "Refresh token is missing",
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}
	config, _ := config.LoadConfig()
	newAccessToken, newRefreshToken, err := controller.RefreshTokenService.RefreshToken(refreshToken, config.RefreshTokenSecret)
	c.SetCookie("refresh_token", newRefreshToken, 3600*24*7, "/", "", false, true)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "bad request",
			Message: "Invalid refresh token",
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}
	loginResponse := response.LoginResponse{
		TokenType:    "Bearer Token",
		RefreshToken: newRefreshToken,
		AccessToken:  newAccessToken,
	}
	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "ok",
		Message: "Refresh token success",
		Data:    loginResponse,
	}
	c.JSON(http.StatusOK, webResponse)
}

func (controller *AuthenticationController) Login(c *gin.Context) {
	LoginRequest := request.LoginRequest{}
	err := c.ShouldBindJSON(&LoginRequest)
	helper.ErrorPanic(err)

	refreshToken, accessToken, err := controller.AuthenticationService.Login(LoginRequest)
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "bad request",
			Message: "Invalid username or password",
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	// Lưu refresh token vào cookie
	c.SetCookie("refresh_token", refreshToken, 3600*24*7, "/", "", false, true)
	// Lưu refresh token vào database
	saveToken := models.RefreshToken{
		Token: refreshToken,
	}
	controller.RefreshTokenService.SaveToken(saveToken)
	loginResponse := response.LoginResponse{
		TokenType:    "Bearer Token",
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
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

	err1 := controller.AuthenticationService.Register(CreateUserRequest)

	var webResponse response.Response

	if err1 != nil {
		webResponse = response.Response{
			Code:    http.StatusBadRequest,
			Status:  "bad request",
			Message: "duplicate username",
		}
	} else {
		webResponse = response.Response{
			Code:    http.StatusOK,
			Status:  "ok",
			Message: "Register success",
			Data:    nil,
		}
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *AuthenticationController) Logout(c *gin.Context) {
	// Lấy token từ header Authorization
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Status:  "bad request",
			Message: "Token is required",
		})
		return
	}
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		helper.ErrorPanic(err)
	}
	// thêm token vào blacklist
	err = controller.AuthenticationService.Logout(context.Background(), refreshToken, token)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Status:  "bad request",
			Message: err.Error(),
		})
		return
	}
	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "ok",
		Message: "Logout success",
	}
	c.JSON(http.StatusOK, webResponse)
}
