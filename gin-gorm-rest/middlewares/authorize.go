package middlewares

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"vietanh/gin-gorm-rest/config"
	"vietanh/gin-gorm-rest/helper"
	"vietanh/gin-gorm-rest/repository"
	"vietanh/gin-gorm-rest/utils"

	"github.com/gin-gonic/gin"
)

// AuthorizationMiddleware is a middleware to check if the user have role to access the resource
func AuthorizeRole(userRepo repository.UserRepository, allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get token
		var token string
		authHeader := ctx.GetHeader("Authorization")
		fields := strings.Fields(authHeader)
		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}
		if authHeader == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"status": "fail", "message": "Missing token"})
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
		currentUser, err := userRepo.FindByID(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "User not found"})
			return
		}

		// Kiểm tra quyền
		for _, role := range allowedRoles {
			if currentUser.Role == role {
				ctx.Next()
				return
			}
		}

		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Permission denied"})
	}
}
