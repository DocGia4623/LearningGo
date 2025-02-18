package middlewares

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"vietanh/gin-gorm-rest/config"
	"vietanh/gin-gorm-rest/helper"
	"vietanh/gin-gorm-rest/repository"
	"vietanh/gin-gorm-rest/service"
	"vietanh/gin-gorm-rest/utils"

	"github.com/gin-gonic/gin"
)

// AuthorizationMiddleware is a middleware to check if the user have role to access the resource
func AuthorizeRole(userRepo repository.UserRepository, permissionService service.PermissionService, roleService service.RoleService, permission string) gin.HandlerFunc {
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
		_, err = userRepo.FindByID(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "User not found"})
			return
		}

		permission, err := permissionService.FindIfExist(permission)
		if err != nil || permission == nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Permission not found"})
			return
		}

		roles, err := roleService.FindBelongToPermission(permission.Name)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "error find role"})
			return
		}
		if len(roles) == 0 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "empty role"})
			return
		}
		idUser := uint(id)
		// Check if user has role
		check := userRepo.FindIfUserHasRole(idUser, roles)
		if check {
			ctx.Next()
			return
		}

		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Permission denied"})
	}
}
