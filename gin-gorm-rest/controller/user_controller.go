package controller

import (
	"net/http"
	"vietanh/gin-gorm-rest/config"
	"vietanh/gin-gorm-rest/data/response"
	"vietanh/gin-gorm-rest/models"
	"vietanh/gin-gorm-rest/repository"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Router /user/ [get]
func (controller *UserController) GetUsers(c *gin.Context) {
	users := controller.UserRepository.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   users,
	}
	c.JSON(200, &webResponse)
}

type UserController struct {
	UserRepository repository.UserRepository
}

func NewUserController(userRepository repository.UserRepository) *UserController {
	return &UserController{
		UserRepository: userRepository,
	}
}

// CreateUser godoc
// @Summary Create a user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Router /user/ [post]
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Router /user/{id} [delete]
func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Router /user/{id} [put]
func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}

// GetUser godoc
// @Summary Get a user
// @Description Get a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Router /user/{id} [get]
func GetUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.JSON(200, &user)
}
