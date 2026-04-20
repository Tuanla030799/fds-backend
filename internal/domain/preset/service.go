package preset

import (
	"encoding/json"
	"strings"

	"fds-backend/internal/domain/fileasset"
	"fds-backend/internal/shared/apperrors"
	"fds-backend/internal/shared/pagination"
	"fds-backend/internal/shared/query"
	"fds-backend/internal/shared/tx"
	"fds-backend/internal/shared/validation"

	"gorm.io/gorm"
)

type Service struct {
	repo  Repository
	db    *gorm.DB
	files *fileasset.Service
}

func NewService(repo Repository, db *gorm.DB, files *fileasset.Service) *Service {
	return &Service{repo: repo, db: db, files: files}
}

type CreateInput struct {
	Name      string
	Status    Status
	Note      string
	Tags      []string
	ImageURL  string
	FileID    string
	ActorID   string
	SortOrder int
}

func (s *Service) Create(input CreateInput) (*Preset, error) {
	if err := validation.Required(input.Name, "name"); err != nil {
		return nil, err
	}
	if !input.Status.IsValid() {
		return nil, apperrors.BadRequest("invalid status")
	}
	if strings.TrimSpace(input.ImageURL) == "" && strings.TrimSpace(input.FileID) == "" {
		return nil, apperrors.BadRequest("image or fileId is required")
	}
	tagsJSON, _ := json.Marshal(input.Tags)
	item := &Preset{Name: strings.TrimSpace(input.Name), Status: input.Status, Note: strings.TrimSpace(input.Note), Tags: tagsJSON, SortOrder: input.SortOrder}
	err := tx.Within(s.db, func(txx *gorm.DB) error {
		if strings.TrimSpace(input.FileID) != "" {
			if s.files == nil {
				return apperrors.BadRequest("file service is unavailable")
			}
			file, err := s.files.ActivateForUseTx(txx, input.FileID, input.ActorID)
			if err != nil {
				return err
			}
			item.ImageURL = file.Path
		} else {
			item.ImageURL = strings.TrimSpace(input.ImageURL)
		}
		return s.repo.CreateTx(txx, item)
	})
	if err != nil {
		return nil, err
	}
	return item, nil
}
func (s *Service) List(filters query.Filters, sort query.Sort, params pagination.Params) ([]Preset, int64, error) {
	return s.repo.List(filters, sort, params)
}
func (s *Service) Delete(id string) error { return s.repo.Delete(id) }
func (s Status) IsValid() bool            { return s == StatusActive || s == StatusDraft || s == StatusArchived }
