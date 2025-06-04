package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoleGuard membatasi akses endpoint berdasarkan role user dari JWT
func RoleGuard(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status_code": http.StatusForbidden,
				"message":     "Forbidden: role not found",
				"data":        nil,
			})
			return
		}

		roleStr, ok := role.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status_code": http.StatusForbidden,
				"message":     "Forbidden: invalid role type",
				"data":        nil,
			})
			return
		}

		for _, allowed := range allowedRoles {
			if roleStr == allowed {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"status_code": http.StatusForbidden,
			"message":     "Forbidden: insufficient permissions",
			"data":        nil,
		})
	}
}
