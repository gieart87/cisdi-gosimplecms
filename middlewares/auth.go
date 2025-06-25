package middlewares

import (
	"github.com/gin-gonic/gin"
	"gosimplecms/utils/jwt"
	"net/http"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":  http.StatusUnauthorized,
				"error": "Missing token",
			})
			return
		}

		claims, err := jwt.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":  http.StatusUnauthorized,
				"error": err.Error(),
			})
			return
		}

		// inject to context
		c.Set("UserID", claims.UserID)
		c.Set("UserRole", claims.Role)
		c.Next()
	}
}

func AllowRoleMiddleware(allowedRole ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleValue, exists := c.Get("UserRole")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":  http.StatusForbidden,
				"error": "Forbidden",
			})
			return
		}

		userRole := strings.ToLower(roleValue.(string))

		for _, r := range allowedRole {
			if userRole == strings.ToLower(r) {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{
			"code":  http.StatusForbidden,
			"error": "access denied for role: " + userRole,
		})
		c.Abort()
	}
}
