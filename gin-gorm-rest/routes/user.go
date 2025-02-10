package routes

import (
	"vietanh/gin-gorm-rest/controller"
	"vietanh/gin-gorm-rest/middlewares"
	"vietanh/gin-gorm-rest/repository"

	"github.com/gin-gonic/gin"
)

// UserRoute defines the user routes
func UserRoute(userRepository repository.UserRepository, userController controller.UserController, router *gin.Engine) {
	userRoutes := router.Group("/user", middlewares.DeserializeUser(userRepository))
	{

		// @Summary Get all users
		// @Description Get all users
		// @Tags users
		// @Accept  json
		// @Produce  json
		// @Success 200 {array} models.User
		// @Router /user/ [get]
		userRoutes.GET("/", middlewares.AuthorizeRole(userRepository, "admin", "manager"), userController.GetUsers)

		// @Summary Create a user
		// @Description Create a new user
		// @Tags users
		// @Accept  json
		// @Produce  json
		// @Param user body models.User true "User"
		// @Success 200 {object} models.User
		// @Router /user/ [post]
		userRoutes.POST("/", middlewares.AuthorizeRole(userRepository, "admin", "manager"), controller.CreateUser)

		// @Summary Delete a user
		// @Description Delete a user by ID
		// @Tags users
		// @Accept  json
		// @Produce  json
		// @Param id path int true "User ID"
		// @Success 200 {object} models.User
		// @Router /user/{id} [delete]
		userRoutes.DELETE("/:id", middlewares.AuthorizeRole(userRepository, "admin", "manager"), controller.DeleteUser)

		// @Summary Update a user
		// @Description Update a user by ID
		// @Tags users
		// @Accept  json
		// @Produce  json
		// @Param id path int true "User ID"
		// @Param user body models.User true "User"
		// @Success 200 {object} models.User
		// @Router /user/{id} [put]
		userRoutes.PUT("/:id", middlewares.AuthorizeRole(userRepository, "admin", "manager", "user"), controller.UpdateUser)

		// @Summary Get a user
		// @Description Get a user by ID
		// @Tags users
		// @Accept  json
		// @Produce  json
		// @Param id path int true "User ID"
		// @Success 200 {object} models.User
		// @Router /user/{id} [get]
		userRoutes.GET("/:id", middlewares.AuthorizeRole(userRepository, "admin", "manager"), controller.GetUser)
	}
}
