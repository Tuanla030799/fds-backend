package storage

import (
	"fmt"
	"mime/multipart"
)

type S3Storage struct{}

func NewS3Storage() *S3Storage                           { return &S3Storage{} }
func (s *S3Storage) EnsureDirs(children ...string) error { return nil }
func (s *S3Storage) Save(file *multipart.FileHeader, folder string) (string, error) {
	return "", fmt.Errorf("s3 storage driver scaffolded but not wired to an SDK in this package yet")
}
