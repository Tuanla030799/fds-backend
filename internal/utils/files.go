package utils

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SaveUploadedFile(c *gin.Context, fileHeader *multipart.FileHeader, dir string) (string, error) {
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if ext == "" {
		ext = ".bin"
	}
	filename := fmt.Sprintf("%d-%s%s", time.Now().UnixNano(), uuid.NewString(), ext)
	fullPath := filepath.Join(dir, filename)
	if err := c.SaveUploadedFile(fileHeader, fullPath); err != nil {
		return "", err
	}
	return filename, nil
}
