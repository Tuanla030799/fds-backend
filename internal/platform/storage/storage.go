package storage

import "mime/multipart"

type Storage interface {
	EnsureDirs(children ...string) error
	Save(file *multipart.FileHeader, folder string) (string, error)
	Delete(path string) error
}
