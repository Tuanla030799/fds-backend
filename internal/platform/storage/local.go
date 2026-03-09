package storage

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

type LocalStorage struct{ root, publicBase string }

func NewLocalStorage(root, publicBase string) *LocalStorage {
	return &LocalStorage{root: root, publicBase: publicBase}
}
func (s *LocalStorage) EnsureDirs(children ...string) error {
	for _, child := range children {
		if err := os.MkdirAll(filepath.Join(s.root, child), 0o755); err != nil {
			return err
		}
	}
	return nil
}
func sanitizeExt(name string) string {
	ext := strings.ToLower(filepath.Ext(name))
	switch ext {
	case ".png", ".jpg", ".jpeg", ".webp":
		return ext
	default:
		return ".bin"
	}
}
func (s *LocalStorage) Save(file *multipart.FileHeader, folder string) (string, error) {
	ext := sanitizeExt(file.Filename)
	filename := fmt.Sprintf("%d-%s%s", time.Now().UnixNano(), uuid.NewString(), ext)
	targetDir := filepath.Join(s.root, folder)
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return "", err
	}
	target := filepath.Join(targetDir, filename)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	dst, err := os.Create(target)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	if _, err := dst.ReadFrom(src); err != nil {
		return "", err
	}
	return strings.TrimRight(s.publicBase, "/") + "/" + folder + "/" + filename, nil
}
