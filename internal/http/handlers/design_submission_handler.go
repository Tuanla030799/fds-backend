package handlers

import (
	"fds-backend/internal/config"
	"fds-backend/internal/domain/auditlog"
	"fds-backend/internal/domain/designsubmission"
	"fds-backend/internal/platform/storage"
	"fds-backend/internal/shared/apperrors"
	"fds-backend/internal/shared/pagination"
	"fds-backend/internal/shared/query"
	"fds-backend/internal/shared/response"
	"fds-backend/internal/shared/validation"

	"github.com/gin-gonic/gin"
)

type DesignSubmissionHandler struct {
	service *designsubmission.Service
	storage storage.Storage
	audit   *auditlog.Service
	cfg     *config.Config
}

func NewDesignSubmissionHandler(service *designsubmission.Service, st storage.Storage, audit *auditlog.Service, cfg *config.Config) *DesignSubmissionHandler {
	return &DesignSubmissionHandler{service: service, storage: st, audit: audit, cfg: cfg}
}

func (h *DesignSubmissionHandler) Create(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		writeError(c, apperrors.BadRequest("image is required"))
		return
	}
	if err := validation.UploadFile(file, h.cfg.Upload.AllowedExts, h.cfg.Upload.MaxBytes); err != nil {
		writeError(c, err)
		return
	}
	imageURL, err := h.storage.Save(file, h.cfg.Upload.OrdersSubdir)
	if err != nil {
		writeError(c, apperrors.Wrap(err, 500, "UPLOAD_SAVE_FAILED", "cannot save image"))
		return
	}
	item, err := h.service.Create(designsubmission.CreateInput{FullName: c.PostForm("fullName"), Address: c.PostForm("address"), Phone: c.PostForm("phone"), Note: c.PostForm("note"), ImageURL: imageURL})
	if err != nil {
		writeError(c, err)
		return
	}
	response.Created(c, item, "created design submission")
}

func (h *DesignSubmissionHandler) AdminList(c *gin.Context) {
	params := pagination.FromRequest(c, h.cfg.Pagination)
	filters := query.ReadFilters(c)
	sort := query.ReadSort(c, "createdAt", map[string]string{"createdAt": "created_at", "updatedAt": "updated_at", "status": "status"})
	items, total, err := h.service.List(filters, sort, params)
	if err != nil {
		writeError(c, apperrors.Wrap(err, 500, "SUBMISSION_LIST_FAILED", "cannot load submissions"))
		return
	}
	response.Paginated(c, items, "ok", pagination.NewMeta(params, total), query.Meta(filters, sort))
}

func (h *DesignSubmissionHandler) AdminUpdateStatus(c *gin.Context) {
	var payload struct {
		Status designsubmission.Status `json:"status"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		writeError(c, invalidPayloadError(err))
		return
	}
	item, err := h.service.UpdateStatus(c.Param("id"), payload.Status)
	if err != nil {
		writeError(c, err)
		return
	}
	audit(c, h.audit, "update_status", "design_submission", c.Param("id"), nil, item)
	response.OK(c, item, "updated submission status")
}

func (h *DesignSubmissionHandler) AdminDelete(c *gin.Context) {
	if err := h.service.Delete(c.Param("id")); err != nil {
		writeError(c, apperrors.Wrap(err, 400, "SUBMISSION_DELETE_FAILED", "cannot delete submission"))
		return
	}
	audit(c, h.audit, "delete", "design_submission", c.Param("id"), nil, gin.H{"id": c.Param("id")})
	response.OK(c, gin.H{"id": c.Param("id")}, "deleted submission")
}
