package middleware

import (
	"time"

	appLogger "fds-backend/internal/platform/logger"

	"github.com/gin-gonic/gin"
)

func AccessLogger(logger *appLogger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startedAt := time.Now()
		c.Next()
		logger.Info("http_request", "request_id", c.GetString("request_id"), "method", c.Request.Method, "path", c.Request.URL.Path, "status", c.Writer.Status(), "latency_ms", time.Since(startedAt).Milliseconds(), "admin_id", c.GetString("admin_id"), "client_ip", c.ClientIP())
	}
}
