package handlers

import (
	"net/http"
	"path/filepath"
	"strings"

	"fds-backend/internal/config"
	"fds-backend/internal/models"
	"fds-backend/internal/services"
	"fds-backend/internal/utils"
	"fds-backend/pkg/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DesignSubmissionHandler struct {
	cfg     *config.Config
	service *services.DesignSubmissionService
}

func NewDesignSubmissionHandler(cfg *config.Config, service *services.DesignSubmissionService) *DesignSubmissionHandler {
	return &DesignSubmissionHandler{cfg: cfg, service: service}
}

func (h *DesignSubmissionHandler) Create(c *gin.Context) {
	fullName := strings.TrimSpace(c.PostForm("fullName"))
	address := strings.TrimSpace(c.PostForm("address"))
	phone := strings.TrimSpace(c.PostForm("phone"))
	note := strings.TrimSpace(c.PostForm("note"))

	file, err := c.FormFile("image")
	if err != nil {
		response.Error(c, http.StatusBadRequest, "image is required", nil)
		return
	}
	if fullName == "" || address == "" || phone == "" {
		response.Error(c, http.StatusBadRequest, "fullName, address and phone are required", nil)
		return
	}

	filename, err := utils.SaveUploadedFile(c, file, filepath.Join(h.cfg.UploadsDir, "design-submissions"))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "cannot save uploaded image", nil)
		return
	}

	item := &models.DesignSubmission{
		FullName: fullName,
		Address:  address,
		Phone:    phone,
		Note:     note,
		ImageURL: "/uploads/design-submissions/" + filename,
		Status:   models.SubmissionPending,
	}
	if err := h.service.Create(item); err != nil {
		response.Error(c, http.StatusInternalServerError, "cannot create design submission", nil)
		return
	}
	response.JSON(c, http.StatusCreated, item, "created design submission")
}

func (h *DesignSubmissionHandler) AdminList(c *gin.Context) {
	items, err := h.service.List(strings.TrimSpace(c.Query("status")))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "cannot load submissions", nil)
		return
	}
	response.JSON(c, http.StatusOK, items, "ok")
}

func (h *DesignSubmissionHandler) AdminUpdateStatus(c *gin.Context) {
	var payload struct {
		Status models.DesignSubmissionStatus `json:"status"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid payload", nil)
		return
	}
	item, err := h.service.FindByID(c.Param("id"))
	if err != nil {
		status := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			status = http.StatusNotFound
		}
		response.Error(c, status, "submission not found", nil)
		return
	}
	item.Status = payload.Status
	if err := h.service.Update(item); err != nil {
		response.Error(c, http.StatusInternalServerError, "cannot update submission", nil)
		return
	}
	response.JSON(c, http.StatusOK, item, "updated submission status")
}

func (h *DesignSubmissionHandler) AdminDelete(c *gin.Context) {
	if err := h.service.Delete(c.Param("id")); err != nil {
		response.Error(c, http.StatusBadRequest, "cannot delete submission", nil)
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"id": c.Param("id")}, "deleted submission")
}
