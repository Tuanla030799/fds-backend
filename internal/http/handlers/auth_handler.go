package handlers

import (
	"fmt"

	"fds-backend/internal/domain/admin"
	"fds-backend/internal/domain/auditlog"
	"fds-backend/internal/shared/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *admin.AuthService
	audit   *auditlog.Service
}

func NewAuthHandler(service *admin.AuthService, audit *auditlog.Service) *AuthHandler {
	return &AuthHandler{service: service, audit: audit}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var payload struct {
		Name, Email, Password string
		Role                  admin.Role
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		writeError(c, invalidPayloadError(err))
		return
	}
	result, err := h.service.Register(payload.Name, payload.Email, payload.Password, payload.Role)
	if err != nil {
		writeError(c, err)
		return
	}
	audit(c, h.audit, "register", "admin", anyToString(result.Admin["id"]), nil, result.Admin)
	response.Created(c, result, "registered admin")
}

func (h *AuthHandler) Login(c *gin.Context) {
	var payload struct{ Email, Password string }
	if err := c.ShouldBindJSON(&payload); err != nil {
		writeError(c, invalidPayloadError(err))
		return
	}
	result, err := h.service.Login(payload.Email, payload.Password)
	if err != nil {
		writeError(c, err)
		return
	}
	audit(c, h.audit, "login", "admin", anyToString(result.Admin["id"]), nil, result.Admin)
	response.OK(c, result, "logged in")
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	var payload struct {
		RefreshToken string `json:"refreshToken"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		writeError(c, invalidPayloadError(err))
		return
	}
	result, err := h.service.Refresh(payload.RefreshToken)
	if err != nil {
		writeError(c, err)
		return
	}
	response.OK(c, result, "refreshed token")
}

func (h *AuthHandler) Logout(c *gin.Context) {
	if err := h.service.Logout(c.GetString("admin_id")); err != nil {
		writeError(c, err)
		return
	}
	response.OK(c, gin.H{"ok": true}, "logged out")
}

func anyToString(v any) string { return fmt.Sprint(v) }
