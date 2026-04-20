package middleware

import (
	"fmt"
	"sync"
	"time"

	"fds-backend/internal/domain/admin"
	"fds-backend/internal/shared/apperrors"
	"fds-backend/internal/shared/response"

	"github.com/gin-gonic/gin"
)

type bucket struct {
	count   int
	resetAt time.Time
	mu      sync.Mutex
}

var rateLimitStore sync.Map

func RateLimit(scope string, perMinute int) gin.HandlerFunc {
	if perMinute <= 0 {
		perMinute = 60
	}
	return func(c *gin.Context) {
		key := fmt.Sprintf("%s:%s", scope, c.ClientIP())
		now := time.Now()
		value, _ := rateLimitStore.LoadOrStore(key, &bucket{count: 0, resetAt: now.Add(time.Minute)})
		b := value.(*bucket)
		b.mu.Lock()
		defer b.mu.Unlock()
		if now.After(b.resetAt) {
			b.count = 0
			b.resetAt = now.Add(time.Minute)
		}
		b.count++
		if b.count > perMinute {
			response.AbortWithError(c, apperrors.TooManyRequests("rate limit exceeded"))
			return
		}
		c.Next()
	}
}

func UploadRateLimit(scope string, perMinute int) gin.HandlerFunc {
	if perMinute <= 0 {
		perMinute = 20
	}
	return func(c *gin.Context) {
		if c.GetString("admin_role") == string(admin.RoleSuperAdmin) {
			c.Next()
			return
		}
		identifier := c.GetString("admin_id")
		if identifier == "" {
			identifier = c.ClientIP()
		}
		key := fmt.Sprintf("%s:%s", scope, identifier)
		now := time.Now()
		value, _ := rateLimitStore.LoadOrStore(key, &bucket{count: 0, resetAt: now.Add(time.Minute)})
		b := value.(*bucket)
		b.mu.Lock()
		defer b.mu.Unlock()
		if now.After(b.resetAt) {
			b.count = 0
			b.resetAt = now.Add(time.Minute)
		}
		b.count++
		if b.count > perMinute {
			response.AbortWithError(c, apperrors.TooManyRequests("upload rate limit exceeded"))
			return
		}
		c.Next()
	}
}
