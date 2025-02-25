package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"vietanh/gin-gorm-rest/config"
	"vietanh/gin-gorm-rest/constant"
	"vietanh/gin-gorm-rest/controller"
	"vietanh/gin-gorm-rest/repository"
	"vietanh/gin-gorm-rest/routes"
	"vietanh/gin-gorm-rest/service"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"

	_ "vietanh/gin-gorm-rest/docs"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

// @title Gin Gorm REST API
// @version 1.0
// @description This is a sample server for a Gin Gorm REST API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8081
// @BasePath /

// setupLogOutput sets up the log output to a file and stdout
func main() {

	// @host localhost:8081
	setupLogOutput()
	router := gin.New()
	log.Println("Starting server...")

	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	router.Use(gin.Recovery(), gindump.Dump())

	config.InitConfig(&appConfig)

	// validate := validator.New()

	//Init Repository
	userRepository := repository.NewUserRepositoryImpl(config.DB)
	permissionRepository := repository.NewPermissionRepositoryImpl(config.DB)
	roleRepository := repository.NewRoleRepositoryimpl(config.DB)
	//Init Service
	roleService := service.NewRoleServiceImpl(roleRepository)
	permissionService := service.NewPermissionServiceImpl(permissionRepository, roleRepository)
	//Init Service
	// authenticationService := service.NewAuthenticationServiceImpl(userRepository, validate)
	rabbitMQService := service.NewRabbitMQServiceImpl(config.RabbitMQConn)

	// 🟢 Khởi chạy consumer trong goroutine
	go StartConsumers(rabbitMQService)
	// //Init controller
	// authenticationController := controller.NewAuthenticationController(authenticationService)
	usersController := controller.NewUserController(userRepository)
	permissionController := controller.NewPermissionController(permissionService)

	routes.UserRoute(userRepository, permissionService, roleService, *usersController, router)
	routes.DeviceRoute(router)
	routes.AuthRoute(router)
	routes.PermissionRoute(permissionController, router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}

func StartConsumers(rabbitMQService service.RabbitMQService) {
	// Consumer cho refresh_token_events
	go rabbitMQService.ConsumeEvent("refresh_token_events", "direct", constant.TokenRefreshedKey, "refresh queue", func(message string) {
		var event struct {
			Token   string `json:"token"`
			Subject string `json:"subject"`
		}
		if err := json.Unmarshal([]byte(message), &event); err != nil {
			log.Printf("❌ Failed to parse refresh token event: %v", err)
			return
		}
		log.Printf("✅ Refresh Token Processed: %s for Subject: %s", event.Token, event.Subject)
	})
	log.Println("🔄 All RabbitMQ Consumers Started.")
}
