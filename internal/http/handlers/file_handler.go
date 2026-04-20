package handlers

import (
	"strings"

	"fds-backend/internal/config"
	"fds-backend/internal/domain/fileasset"
	"fds-backend/internal/shared/apperrors"
	"fds-backend/internal/shared/response"
	"fds-backend/internal/shared/validation"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	service *fileasset.Service
	cfg     *config.Config
}

func NewFileHandler(service *fileasset.Service, cfg *config.Config) *FileHandler {
	return &FileHandler{service: service, cfg: cfg}
}

func (h *FileHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		writeError(c, apperrors.BadRequest("file is required"))
		return
	}
	if err := validation.UploadFile(file, h.cfg.Upload.AllowedExts, h.cfg.Upload.MaxBytes); err != nil {
		writeError(c, err)
		return
	}
	folder := strings.TrimSpace(c.DefaultPostForm("folder", h.cfg.Upload.TempSubdir))
	if !h.allowedFolder(folder) {
		writeError(c, apperrors.BadRequest("invalid upload folder"))
		return
	}
	item, err := h.service.CreateUpload(fileasset.UploadInput{
		File:      file,
		Folder:    folder,
		CreatedBy: c.GetString("admin_id"),
	})
	if err != nil {
		writeError(c, err)
		return
	}
	response.Created(c, item, "uploaded file")
}

func (h *FileHandler) allowedFolder(folder string) bool {
	return folder == h.cfg.Upload.TempSubdir || folder == h.cfg.Upload.PresetsSubdir || folder == h.cfg.Upload.OrdersSubdir
}
