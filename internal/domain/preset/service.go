package preset

import (
	"encoding/json"
	"strings"

	"fds-backend/internal/shared/apperrors"
	"fds-backend/internal/shared/pagination"
	"fds-backend/internal/shared/query"
	"fds-backend/internal/shared/validation"
)

type Service struct{ repo Repository }

func NewService(repo Repository) *Service { return &Service{repo: repo} }

type CreateInput struct {
	Name      string
	Status    Status
	Note      string
	Tags      []string
	ImageURL  string
	SortOrder int
}

func (s *Service) Create(input CreateInput) (*Preset, error) {
	if err := validation.Required(input.Name, "name"); err != nil {
		return nil, err
	}
	if err := validation.Required(input.ImageURL, "image"); err != nil {
		return nil, err
	}
	if !input.Status.IsValid() {
		return nil, apperrors.BadRequest("invalid status")
	}
	tagsJSON, _ := json.Marshal(input.Tags)
	item := &Preset{Name: strings.TrimSpace(input.Name), Status: input.Status, Note: strings.TrimSpace(input.Note), Tags: tagsJSON, ImageURL: strings.TrimSpace(input.ImageURL), SortOrder: input.SortOrder}
	return item, s.repo.Create(item)
}
func (s *Service) List(filters query.Filters, sort query.Sort, params pagination.Params) ([]Preset, int64, error) {
	return s.repo.List(filters, sort, params)
}
func (s *Service) Delete(id string) error { return s.repo.Delete(id) }
func (s Status) IsValid() bool            { return s == StatusActive || s == StatusDraft || s == StatusArchived }
