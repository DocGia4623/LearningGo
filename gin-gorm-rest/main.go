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
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
	router := gin.New()
	log.Println("Starting server...")
	router.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8080")
	// router.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Hello world")
	// })
}
