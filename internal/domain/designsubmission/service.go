package designsubmission

import (
	"errors"
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
	FullName string
	Address  string
	Phone    string
	Note     string
	ImageURL string
	FileID   string
	ActorID  string
}

func (s *Service) Create(input CreateInput) (*Submission, error) {
	if err := validation.Required(input.FullName, "fullName"); err != nil {
		return nil, err
	}
	if err := validation.Required(input.Address, "address"); err != nil {
		return nil, err
	}
	if err := validation.Phone(input.Phone); err != nil {
		return nil, err
	}
	if strings.TrimSpace(input.ImageURL) == "" && strings.TrimSpace(input.FileID) == "" {
		return nil, apperrors.BadRequest("image or fileId is required")
	}
	item := &Submission{
		FullName: strings.TrimSpace(input.FullName),
		Address:  strings.TrimSpace(input.Address),
		Phone:    strings.TrimSpace(input.Phone),
		Note:     strings.TrimSpace(input.Note),
		Status:   StatusPending,
	}
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
func (s *Service) List(filters query.Filters, sort query.Sort, params pagination.Params) ([]Submission, int64, error) {
	return s.repo.List(filters, sort, params)
}
func (s *Service) UpdateStatus(id string, status Status) (*Submission, error) {
	if !status.IsValid() {
		return nil, apperrors.BadRequest("invalid status")
	}
	item, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NotFound("submission not found")
		}
		return nil, err
	}
	item.Status = status
	return item, s.repo.Update(item)
}
func (s *Service) Delete(id string) error { return s.repo.Delete(id) }
func (s Status) IsValid() bool {
	return s == StatusPending || s == StatusConfirmed || s == StatusRejected || s == StatusCompleted
}
