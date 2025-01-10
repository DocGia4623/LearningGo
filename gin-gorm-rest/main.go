package main

import (
	"vietanh/gin-gorm-rest/config"
	"vietanh/gin-gorm-rest/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8080")
	// router.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Hello world")
	// })
}
