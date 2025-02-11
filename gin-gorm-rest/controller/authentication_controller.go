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

// RefreshToken godoc
// @Summary Refresh access token
// @Description Generate a new access token using a refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /auth/refresh [post]
func (controller *AuthenticationController) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{Code: http.StatusUnauthorized, Status: "unauthorized", Message: "Refresh token is missing"})
		return
	}
	config, _ := config.LoadConfig()
	newAccessToken, newRefreshToken, err := controller.RefreshTokenService.RefreshToken(refreshToken, config.RefreshTokenSecret)
	c.SetCookie("refresh_token", newRefreshToken, 3600*24*7, "/", "", false, true)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{Code: http.StatusBadRequest, Status: "bad request", Message: "Invalid refresh token"})
		return
	}
	c.JSON(http.StatusOK, response.Response{Code: http.StatusOK, Status: "ok", Message: "Refresh token success", Data: response.LoginResponse{TokenType: "Bearer Token", RefreshToken: newRefreshToken, AccessToken: newAccessToken}})
}

// Login godoc
// @Summary Authenticate user
// @Description Authenticate a user and return access & refresh tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.LoginRequest true "Login Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /auth/login [post]
func (controller *AuthenticationController) Login(c *gin.Context) {
	LoginRequest := request.LoginRequest{}
	err := c.ShouldBindJSON(&LoginRequest)
	helper.ErrorPanic(err)

	refreshToken, accessToken, err := controller.AuthenticationService.Login(LoginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{Code: http.StatusBadRequest, Status: "bad request", Message: "Invalid username or password"})
		return
	}
	c.SetCookie("refresh_token", refreshToken, 3600*24*7, "/", "", false, true)
	controller.RefreshTokenService.SaveToken(models.RefreshToken{Token: refreshToken})
	c.JSON(http.StatusOK, response.Response{Code: http.StatusOK, Status: "ok", Message: "Login success", Data: response.LoginResponse{TokenType: "Bearer Token", RefreshToken: refreshToken, AccessToken: accessToken}})
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.CreateUserRequest true "Register Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /auth/register [post]
func (controller *AuthenticationController) Register(c *gin.Context) {
	CreateUserRequest := request.CreateUserRequest{}
	err := c.ShouldBindJSON(&CreateUserRequest)
	helper.ErrorPanic(err)

	err1 := controller.AuthenticationService.Register(CreateUserRequest)

	var webResponse response.Response
	if err1 != nil {
		webResponse = response.Response{Code: http.StatusBadRequest, Status: "bad request", Message: "duplicate username"}
	} else {
		webResponse = response.Response{Code: http.StatusOK, Status: "ok", Message: "Register success"}
	}
	c.JSON(http.StatusOK, webResponse)
}

// Logout godoc
// @Summary Logout user
// @Description Invalidate access and refresh tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Access Token"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /auth/logout [post]
func (controller *AuthenticationController) Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, response.Response{Code: http.StatusBadRequest, Status: "bad request", Message: "Token is required"})
		return
	}
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		helper.ErrorPanic(err)
	}
	err = controller.AuthenticationService.Logout(context.Background(), refreshToken, token)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{Code: http.StatusBadRequest, Status: "bad request", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{Code: http.StatusOK, Status: "ok", Message: "Logout success"})
}
