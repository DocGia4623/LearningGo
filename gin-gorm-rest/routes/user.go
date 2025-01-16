package routes

import (
	"vietanh/gin-gorm-rest/controller"

	"github.com/gin-gonic/gin"
)

// UserRoute defines the user routes
func UserRoute(router *gin.Engine) {
	userRoutes := router.Group("/user")
	{
		// @Summary Get all users
		// @Description Get all users
		// @Tags users
		// @Accept  json
		// @Produce  json
		// @Success 200 {array} models.User
		// @Router /user/ [get]
		userRoutes.GET("/", controller.GetUsers)

		// @Summary Create a user
		// @Description Create a new user
		// @Tags users
		// @Accept  json
		// @Produce  json
		// @Param user body models.User true "User"
		// @Success 200 {object} models.User
		// @Router /user/ [post]
		userRoutes.POST("/", controller.CreateUser)

		// @Summary Delete a user
		// @Description Delete a user by ID
		// @Tags users
		// @Accept  json
		// @Produce  json
		// @Param id path int true "User ID"
		// @Success 200 {object} models.User
		// @Router /user/{id} [delete]
		userRoutes.DELETE("/:id", controller.DeleteUser)

		// @Summary Update a user
		// @Description Update a user by ID
		// @Tags users
		// @Accept  json
		// @Produce  json
		// @Param id path int true "User ID"
		// @Param user body models.User true "User"
		// @Success 200 {object} models.User
		// @Router /user/{id} [put]
		userRoutes.PUT("/:id", controller.UpdateUser)

		// @Summary Get a user
		// @Description Get a user by ID
		// @Tags users
		// @Accept  json
		// @Produce  json
		// @Param id path int true "User ID"
		// @Success 200 {object} models.User
		// @Router /user/{id} [get]
		userRoutes.GET("/:id", controller.GetUser)
	}
}
