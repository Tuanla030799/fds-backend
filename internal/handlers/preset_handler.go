package handlers

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"fds-backend/internal/config"
	"fds-backend/internal/models"
	"fds-backend/internal/services"
	"fds-backend/internal/utils"
	"fds-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type PresetHandler struct {
	cfg     *config.Config
	service *services.PresetService
}

func NewPresetHandler(cfg *config.Config, service *services.PresetService) *PresetHandler {
	return &PresetHandler{cfg: cfg, service: service}
}

type presetClientRow struct {
	ID       any      `json:"id"`
	Name     string   `json:"name"`
	Status   string   `json:"status"`
	Tags     []string `json:"tags"`
	Note     string   `json:"note"`
	ImageURL string   `json:"imageUrl,omitempty"`
}

func (h *PresetHandler) PublicList(c *gin.Context) {
	items, err := h.service.List(strings.TrimSpace(c.Query("status")))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "cannot load presets", nil)
		return
	}
	rows := make([]presetClientRow, 0, len(items))
	for _, item := range items {
		var tags []string
		_ = json.Unmarshal(item.Tags, &tags)
		rows = append(rows, presetClientRow{ID: item.ID, Name: item.Name, Status: string(item.Status), Tags: tags, Note: item.Note, ImageURL: item.ImageURL})
	}
	response.JSON(c, http.StatusOK, rows, "ok")
}

func (h *PresetHandler) AdminCreate(c *gin.Context) {
	name := strings.TrimSpace(c.PostForm("name"))
	note := strings.TrimSpace(c.PostForm("note"))
	status := models.PresetStatus(strings.TrimSpace(c.DefaultPostForm("status", string(models.PresetActive))))
	tags := strings.Split(strings.TrimSpace(c.DefaultPostForm("tags", "")), ",")
	sortOrder, _ := strconv.Atoi(c.DefaultPostForm("sortOrder", "0"))
	cleanTags := make([]string, 0, len(tags))
	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		if tag != "" {
			cleanTags = append(cleanTags, tag)
		}
	}
	if name == "" {
		response.Error(c, http.StatusBadRequest, "name is required", nil)
		return
	}
	file, err := c.FormFile("image")
	if err != nil {
		response.Error(c, http.StatusBadRequest, "image is required", nil)
		return
	}
	filename, err := utils.SaveUploadedFile(c, file, filepath.Join(h.cfg.UploadsDir, "presets"))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "cannot save preset image", nil)
		return
	}
	item, err := h.service.Create(name, status, note, cleanTags, "/uploads/presets/"+filename, sortOrder)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "cannot create preset", nil)
		return
	}
	response.JSON(c, http.StatusCreated, item, "created preset")
}

func (h *PresetHandler) AdminDelete(c *gin.Context) {
	if err := h.service.Delete(c.Param("id")); err != nil {
		response.Error(c, http.StatusBadRequest, "cannot delete preset", nil)
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"id": c.Param("id")}, "deleted preset")
}
