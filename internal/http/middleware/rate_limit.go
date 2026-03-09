package middleware

import (
	"fmt"
	"sync"
	"time"

	"fds-backend/internal/shared/apperrors"
	"fds-backend/internal/shared/response"

	"github.com/gin-gonic/gin"
)

type bucket struct {
	count   int
	resetAt time.Time
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
