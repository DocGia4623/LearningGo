package routes

import (
	"log"
	"vietanh/gin-gorm-rest/config"
	"vietanh/gin-gorm-rest/controller"
	"vietanh/gin-gorm-rest/repository"
	"vietanh/gin-gorm-rest/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AuthRoute(router *gin.Engine) {
	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	db := config.Connect(&appConfig) // Ensure the DB connection is established
	userRepo := repository.NewUserRepositoryImpl(db)
	refreshTokenRepo := repository.NewRefreshTokenRepositoryImpl(db)
	validate := validator.New()

	authenticationService := service.NewAuthenticationServiceImpl(userRepo, validate)
	refreshTokenService := service.NewRefreshTokenServiceImpl(refreshTokenRepo)
	authController := controller.NewAuthenticationController(authenticationService, refreshTokenService)
	authRoutes := router.Group("/auth")
	{
		// @Summary Login
		// @Description Login
		// @Tags auth
		// @Accept  json
		// @Produce  json
		// @Param login body request.LoginRequest true "Login"
		// @Success 200 {object} response.LoginResponse
		// @Router /auth/login [post]
		authRoutes.POST("/login", authController.Login)

		// @Summary Register
		// @Description Register
		// @Tags auth
		// @Accept  json
		// @Produce  json
		// @Param register body request.CreateUserRequest true "Register"
		// @Success 200 {object} response.Response
		// @Router /auth/register [post]
		authRoutes.POST("/register", authController.Register)

		// @Summary Logout
		// @Description Logout
		// @Tags auth
		// @Accept  json
		// @Produce  json
		// @Success 200 {object} response.Response
		// @Router /auth/logout [post]
		authRoutes.POST("/logout", authController.Logout)
		authRoutes.POST("/refresh", authController.RefreshToken)
	}
}
