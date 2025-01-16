package main

import (
	"io"
	"log"
	"os"
	"vietanh/gin-gorm-rest/config"
	"vietanh/gin-gorm-rest/middlewares"
	"vietanh/gin-gorm-rest/routes"

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

	router.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	config.Connect()

	routes.UserRoute(router)
	routes.DeviceRoute(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
