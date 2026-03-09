package handlers

import (
	"fds-backend/internal/shared/response"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler { return &HealthHandler{} }

func (h *HealthHandler) Check(c *gin.Context) {
	response.OK(c, gin.H{"ok": true}, "healthy")
}
