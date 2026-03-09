package validation

import (
	"mime/multipart"
	"net/mail"
	"path/filepath"
	"regexp"
	"strings"

	"fds-backend/internal/shared/apperrors"
)

var phoneRegexp = regexp.MustCompile(`^[0-9+()\-\s]{8,20}$`)

func Required(value, field string) error {
	if strings.TrimSpace(value) == "" {
		return apperrors.InvalidPayload(map[string]any{"field": field, "message": field + " is required"})
	}
	return nil
}
func Email(value string) error {
	if _, err := mail.ParseAddress(strings.TrimSpace(value)); err != nil {
		return apperrors.InvalidPayload(map[string]any{"field": "email", "message": "invalid email"})
	}
	return nil
}
func Password(value string, min int) error {
	if len(strings.TrimSpace(value)) < min {
		return apperrors.InvalidPayload(map[string]any{"field": "password", "message": "password is too short", "min": min})
	}
	return nil
}
func Phone(value string) error {
	if !phoneRegexp.MatchString(strings.TrimSpace(value)) {
		return apperrors.InvalidPayload(map[string]any{"field": "phone", "message": "invalid phone format"})
	}
	return nil
}
func Enum(value, field string, allowed []string) error {
	for _, a := range allowed {
		if value == a {
			return nil
		}
	}
	return apperrors.InvalidPayload(map[string]any{"field": field, "message": "invalid value", "allowed": allowed})
}
func UploadFile(file *multipart.FileHeader, allowedExts []string, maxBytes int64) error {
	if file == nil {
		return apperrors.BadRequest("file is required")
	}
	if file.Size > maxBytes {
		return apperrors.InvalidPayload(map[string]any{"field": "file", "message": "file too large", "maxBytes": maxBytes})
	}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	for _, allowed := range allowedExts {
		if ext == strings.ToLower(strings.TrimSpace(allowed)) {
			return nil
		}
	}
	return apperrors.InvalidPayload(map[string]any{"field": "file", "message": "unsupported file extension", "ext": ext})
}
