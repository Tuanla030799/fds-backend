package middleware

import (
	"strings"

	"fds-backend/internal/domain/admin"
	"fds-backend/internal/platform/auth"
	"fds-backend/internal/shared/apperrors"
	"fds-backend/internal/shared/response"

	"github.com/gin-gonic/gin"
)

type Set struct{ JWT gin.HandlerFunc }

func NewJWTMiddleware(jwtManager *auth.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			response.AbortWithError(c, apperrors.Unauthorized("missing bearer token"))
			return
		}
		claims, err := jwtManager.Parse(strings.TrimPrefix(header, "Bearer "))
		if err != nil {
			response.AbortWithError(c, apperrors.Unauthorized("invalid token"))
			return
		}
		c.Set("jwtClaims", claims)
		if sub, ok := claims["sub"].(string); ok {
			c.Set("admin_id", sub)
		}
		if role, ok := claims["role"].(string); ok {
			c.Set("admin_role", role)
		}
		c.Next()
	}
}

func RequireRoles(roles ...admin.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		current := c.GetString("admin_role")
		for _, role := range roles {
			if current == string(role) {
				c.Next()
				return
			}
		}
		response.AbortWithError(c, apperrors.Forbidden("insufficient role permissions"))
	}
}
