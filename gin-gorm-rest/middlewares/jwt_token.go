package middlewares

import (
	"fmt"
	"strconv"
	"strings"
	"vietanh/gin-gorm-rest/config"
	"vietanh/gin-gorm-rest/helper"
	"vietanh/gin-gorm-rest/repository"
	"vietanh/gin-gorm-rest/utils"

	"github.com/gin-gonic/gin"
)

func DeserializeUser(userRepository repository.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		authorization := ctx.GetHeader("Authorization")
		fields := strings.Fields(authorization)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}
		if token == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"status": "fail", "message": "You are not authorized to access this resource"})
			return
		}
		config, _ := config.LoadConfig()
		sub, err := utils.ValidateToken(token, config.AccessTokenSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		id, err_id := strconv.Atoi(fmt.Sprint(sub))
		helper.ErrorPanic(err_id)
		result, err := userRepository.FindByID(id)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"status": "fail", "message": "User not found"})
			return
		}

		ctx.Set("CurrentUser", result.UserName)
		ctx.Next()

	}

}
