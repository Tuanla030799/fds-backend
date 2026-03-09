package designsubmission

import (
	"errors"
	"strings"

	"fds-backend/internal/shared/apperrors"
	"fds-backend/internal/shared/pagination"
	"fds-backend/internal/shared/query"
	"fds-backend/internal/shared/validation"

	"gorm.io/gorm"
)

type Service struct{ repo Repository }

func NewService(repo Repository) *Service { return &Service{repo: repo} }

type CreateInput struct{ FullName, Address, Phone, Note, ImageURL string }

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
	if err := validation.Required(input.ImageURL, "image"); err != nil {
		return nil, err
	}
	item := &Submission{FullName: strings.TrimSpace(input.FullName), Address: strings.TrimSpace(input.Address), Phone: strings.TrimSpace(input.Phone), Note: strings.TrimSpace(input.Note), ImageURL: strings.TrimSpace(input.ImageURL), Status: StatusPending}
	return item, s.repo.Create(item)
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
