package handlers

import (
	"encoding/json"
	"strconv"
	"strings"

	"fds-backend/internal/config"
	"fds-backend/internal/domain/auditlog"
	"fds-backend/internal/domain/preset"
	"fds-backend/internal/platform/storage"
	"fds-backend/internal/shared/apperrors"
	"fds-backend/internal/shared/pagination"
	"fds-backend/internal/shared/query"
	"fds-backend/internal/shared/response"
	"fds-backend/internal/shared/validation"

	"github.com/gin-gonic/gin"
)

type PresetHandler struct {
	service *preset.Service
	storage storage.Storage
	audit   *auditlog.Service
	cfg     *config.Config
}

func NewPresetHandler(service *preset.Service, st storage.Storage, audit *auditlog.Service, cfg *config.Config) *PresetHandler {
	return &PresetHandler{service: service, storage: st, audit: audit, cfg: cfg}
}

type presetClientRow struct {
	ID                           any `json:"id"`
	Name, Status, Note, ImageURL string
	Tags                         []string `json:"tags"`
	SortOrder                    int      `json:"sortOrder"`
}

func (h *PresetHandler) PublicList(c *gin.Context) {
	params := pagination.FromRequest(c, h.cfg.Pagination)
	filters := query.ReadFilters(c)
	sort := query.ReadSort(c, "sortOrder", map[string]string{"sortOrder": "sort_order", "createdAt": "created_at", "name": "name"})
	items, total, err := h.service.List(filters, sort, params)
	if err != nil {
		writeError(c, apperrors.Wrap(err, 500, "PRESET_LIST_FAILED", "cannot load presets"))
		return
	}
	rows := make([]presetClientRow, 0, len(items))
	for _, item := range items {
		var tags []string
		_ = json.Unmarshal(item.Tags, &tags)
		rows = append(rows, presetClientRow{ID: item.ID, Name: item.Name, Status: string(item.Status), Tags: tags, Note: item.Note, ImageURL: item.ImageURL, SortOrder: item.SortOrder})
	}
	response.Paginated(c, rows, "ok", pagination.NewMeta(params, total), query.Meta(filters, sort))
}

func (h *PresetHandler) AdminCreate(c *gin.Context) {
	name := strings.TrimSpace(c.PostForm("name"))
	note := strings.TrimSpace(c.PostForm("note"))
	status := preset.Status(strings.TrimSpace(c.DefaultPostForm("status", string(preset.StatusActive))))
	fileID := strings.TrimSpace(c.PostForm("fileId"))
	tags := strings.Split(strings.TrimSpace(c.DefaultPostForm("tags", "")), ",")
	sortOrder, _ := strconv.Atoi(c.DefaultPostForm("sortOrder", "0"))
	cleanTags := make([]string, 0, len(tags))
	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		if tag != "" {
			cleanTags = append(cleanTags, tag)
		}
	}
	imageURL := ""
	if fileID == "" {
		file, err := c.FormFile("image")
		if err != nil {
			writeError(c, apperrors.BadRequest("image or fileId is required"))
			return
		}
		if err := validation.UploadFile(file, h.cfg.Upload.AllowedExts, h.cfg.Upload.MaxBytes); err != nil {
			writeError(c, err)
			return
		}
		imageURL, err = h.storage.Save(file, h.cfg.Upload.PresetsSubdir)
		if err != nil {
			writeError(c, apperrors.Wrap(err, 500, "UPLOAD_SAVE_FAILED", "cannot save preset image"))
			return
		}
	}
	item, err := h.service.Create(preset.CreateInput{Name: name, Status: status, Note: note, Tags: cleanTags, ImageURL: imageURL, FileID: fileID, ActorID: c.GetString("admin_id"), SortOrder: sortOrder})
	if err != nil {
		writeError(c, err)
		return
	}
	audit(c, h.audit, "create", "preset", item.ID.String(), nil, item)
	response.Created(c, item, "created preset")
}

func (h *PresetHandler) AdminDelete(c *gin.Context) {
	if err := h.service.Delete(c.Param("id")); err != nil {
		writeError(c, apperrors.Wrap(err, 400, "PRESET_DELETE_FAILED", "cannot delete preset"))
		return
	}
	audit(c, h.audit, "delete", "preset", c.Param("id"), nil, gin.H{"id": c.Param("id")})
	response.OK(c, gin.H{"id": c.Param("id")}, "deleted preset")
}
