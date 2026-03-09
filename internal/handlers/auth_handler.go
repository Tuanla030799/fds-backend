package handlers

import (
	"net/http"
	"strings"

	"fds-backend/internal/services"
	"fds-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{ service *services.AuthService }

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

type authPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var payload authPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid payload", nil)
		return
	}
	if strings.TrimSpace(payload.Name) == "" || strings.TrimSpace(payload.Email) == "" || len(payload.Password) < 6 {
		response.Error(c, http.StatusBadRequest, "name, email and password(min 6) are required", nil)
		return
	}
	result, err := h.service.Register(strings.TrimSpace(payload.Name), strings.TrimSpace(strings.ToLower(payload.Email)), payload.Password)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.JSON(c, http.StatusCreated, result, "registered successfully")
}

func (h *AuthHandler) Login(c *gin.Context) {
	var payload authPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid payload", nil)
		return
	}
	result, err := h.service.Login(strings.TrimSpace(strings.ToLower(payload.Email)), payload.Password)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error(), nil)
		return
	}
	response.JSON(c, http.StatusOK, result, "login successfully")
}
