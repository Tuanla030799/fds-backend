package apperrors

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	Status  int
	Code    string
	Message string
	Details any
	Err     error
}

func (e *AppError) Error() string {
	if e == nil {
		return ""
	}
	if e.Message != "" {
		return e.Message
	}
	if e.Err != nil {
		return e.Err.Error()
	}
	return http.StatusText(e.Status)
}
func (e *AppError) Unwrap() error                     { return e.Err }
func (e *AppError) WithDetails(details any) *AppError { e.Details = details; return e }
func New(status int, code, message string) error {
	return &AppError{Status: status, Code: code, Message: message}
}
func Wrap(err error, status int, code, message string) error {
	if err == nil {
		return nil
	}
	return &AppError{Status: status, Code: code, Message: message, Err: err}
}
func As(err error) (*AppError, bool) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr, true
	}
	return nil, false
}
func Normalize(err error) *AppError {
	if err == nil {
		return nil
	}
	if appErr, ok := As(err); ok {
		if appErr.Status == 0 {
			appErr.Status = http.StatusInternalServerError
		}
		if appErr.Code == "" {
			appErr.Code = defaultCodeForStatus(appErr.Status)
		}
		if appErr.Message == "" {
			appErr.Message = http.StatusText(appErr.Status)
		}
		return appErr
	}
	return &AppError{Status: http.StatusInternalServerError, Code: defaultCodeForStatus(http.StatusInternalServerError), Message: "internal server error", Err: err, Details: map[string]any{"reason": fmt.Sprintf("%v", err)}}
}
func defaultCodeForStatus(status int) string {
	switch status {
	case http.StatusBadRequest:
		return "BAD_REQUEST"
	case http.StatusUnauthorized:
		return "UNAUTHORIZED"
	case http.StatusForbidden:
		return "FORBIDDEN"
	case http.StatusNotFound:
		return "NOT_FOUND"
	case http.StatusConflict:
		return "CONFLICT"
	case http.StatusTooManyRequests:
		return "TOO_MANY_REQUESTS"
	case http.StatusUnprocessableEntity:
		return "UNPROCESSABLE_ENTITY"
	default:
		return "INTERNAL_SERVER_ERROR"
	}
}
func BadRequest(message string) error { return New(http.StatusBadRequest, "BAD_REQUEST", message) }
func InvalidPayload(details any) error {
	return (&AppError{Status: http.StatusBadRequest, Code: "INVALID_PAYLOAD", Message: "invalid payload", Details: details})
}
func Unauthorized(message string) error { return New(http.StatusUnauthorized, "UNAUTHORIZED", message) }
func Forbidden(message string) error    { return New(http.StatusForbidden, "FORBIDDEN", message) }
func NotFound(message string) error     { return New(http.StatusNotFound, "NOT_FOUND", message) }
func Conflict(message string) error     { return New(http.StatusConflict, "CONFLICT", message) }
func TooManyRequests(message string) error {
	return New(http.StatusTooManyRequests, "TOO_MANY_REQUESTS", message)
}
