package handlers

import (
	"fds-backend/internal/domain/auditlog"
	"fds-backend/internal/shared/apperrors"
	"fds-backend/internal/shared/response"

	"github.com/gin-gonic/gin"
)

func writeError(c *gin.Context, err error) { response.Error(c, err) }
func invalidPayloadError(err error) error {
	if err == nil {
		return nil
	}
	return apperrors.InvalidPayload(map[string]any{"reason": err.Error()})
}
func audit(c *gin.Context, svc *auditlog.Service, action, resource, resourceID string, before, after any) {
	if svc == nil {
		return
	}
	_ = svc.Record(auditlog.RecordInput{AdminID: c.GetString("admin_id"), Action: action, Resource: resource, ResourceID: resourceID, RequestID: c.GetString("request_id"), IP: c.ClientIP(), UserAgent: c.Request.UserAgent(), Before: before, After: after})
}
