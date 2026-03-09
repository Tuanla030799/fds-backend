package response

import (
	"net/http"

	"fds-backend/internal/shared/apperrors"
	"fds-backend/internal/shared/pagination"

	"github.com/gin-gonic/gin"
)

type ErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}
type Envelope struct {
	Success   bool           `json:"success"`
	Data      any            `json:"data,omitempty"`
	Message   string         `json:"message,omitempty"`
	Error     *ErrorBody     `json:"error,omitempty"`
	Meta      map[string]any `json:"meta,omitempty"`
	RequestID string         `json:"requestId,omitempty"`
}

func requestIDFrom(c *gin.Context) string { return c.GetString("request_id") }
func mergeMeta(base map[string]any, extra map[string]any) map[string]any {
	if base == nil && extra == nil {
		return nil
	}
	out := map[string]any{}
	for k, v := range base {
		out[k] = v
	}
	for k, v := range extra {
		out[k] = v
	}
	return out
}
func Success(c *gin.Context, status int, data any, message string, meta map[string]any) {
	c.JSON(status, Envelope{Success: true, Data: data, Message: message, Meta: meta, RequestID: requestIDFrom(c)})
}
func OK(c *gin.Context, data any, message string) { Success(c, http.StatusOK, data, message, nil) }
func OKWithMeta(c *gin.Context, data any, message string, meta map[string]any) {
	Success(c, http.StatusOK, data, message, meta)
}
func Paginated(c *gin.Context, data any, message string, p pagination.Meta, extraMeta map[string]any) {
	OKWithMeta(c, data, message, mergeMeta(map[string]any{"pagination": p}, extraMeta))
}
func Created(c *gin.Context, data any, message string) {
	Success(c, http.StatusCreated, data, message, nil)
}
func Error(c *gin.Context, err error) {
	appErr := apperrors.Normalize(err)
	c.JSON(appErr.Status, Envelope{Success: false, Message: appErr.Message, Error: &ErrorBody{Code: appErr.Code, Message: appErr.Message, Details: appErr.Details}, RequestID: requestIDFrom(c)})
}
func AbortWithError(c *gin.Context, err error) { Error(c, err); c.Abort() }
