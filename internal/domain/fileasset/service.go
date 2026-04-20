package fileasset

import (
	"errors"
	"mime/multipart"
	"strings"
	"time"

	"fds-backend/internal/platform/storage"
	"fds-backend/internal/shared/apperrors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	repo         Repository
	storage      storage.Storage
	cleanupAfter time.Duration
}

type UploadInput struct {
	File      *multipart.FileHeader
	Folder    string
	CreatedBy string
}

func NewService(repo Repository, st storage.Storage, cleanupAfter time.Duration) *Service {
	return &Service{repo: repo, storage: st, cleanupAfter: cleanupAfter}
}

func (s *Service) CreateUpload(input UploadInput) (*File, error) {
	path, err := s.storage.Save(input.File, input.Folder)
	if err != nil {
		return nil, err
	}
	item := &File{
		Path:     path,
		Filename: strings.TrimSpace(input.File.Filename),
		MimeType: strings.TrimSpace(input.File.Header.Get("Content-Type")),
		Size:     input.File.Size,
		Status:   StatusUnactive,
	}
	if item.MimeType == "" {
		item.MimeType = "application/octet-stream"
	}
	if input.CreatedBy != "" {
		adminID, err := uuid.Parse(input.CreatedBy)
		if err != nil {
			_ = s.storage.Delete(path)
			return nil, apperrors.BadRequest("invalid creator")
		}
		item.CreatedBy = &adminID
	}
	if err := s.repo.Create(item); err != nil {
		_ = s.storage.Delete(path)
		return nil, err
	}
	return item, nil
}

func (s *Service) ActivateForUseTx(txx *gorm.DB, fileID, actorID string) (*File, error) {
	item, err := NewGormRepository(txx).FindByIDForUpdate(txx, fileID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NotFound("file not found")
		}
		return nil, err
	}
	if item.Status == StatusActive {
		return item, nil
	}
	if item.Status != StatusUnactive {
		return nil, apperrors.BadRequest("invalid file status")
	}
	if item.CreatedBy != nil && item.CreatedBy.String() != actorID {
		return nil, apperrors.Forbidden("file does not belong to current user")
	}
	item.Status = StatusActive
	if err := txx.Save(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) CleanupExpiredInactive() error {
	items, err := s.repo.ListExpiredInactive(time.Now().Add(-s.cleanupAfter))
	if err != nil {
		return err
	}
	for _, item := range items {
		if err := s.storage.Delete(item.Path); err != nil {
			return err
		}
		if err := s.repo.DeleteByID(item.ID); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) StartCleanupJob() {
	ticker := time.NewTicker(time.Minute)
	go func() {
		for range ticker.C {
			_ = s.CleanupExpiredInactive()
		}
	}()
}
