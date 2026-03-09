package services

import (
	"encoding/json"

	"fds-backend/internal/models"
	"fds-backend/internal/repositories"
)

type PresetService struct {
	repo *repositories.PresetRepository
}

func NewPresetService(repo *repositories.PresetRepository) *PresetService {
	return &PresetService{repo: repo}
}

func (s *PresetService) Create(name string, status models.PresetStatus, note string, tags []string, imageURL string, sortOrder int) (*models.Preset, error) {
	payload, err := json.Marshal(tags)
	if err != nil {
		return nil, err
	}
	item := &models.Preset{Name: name, Status: status, Note: note, Tags: payload, ImageURL: imageURL, SortOrder: sortOrder}
	return item, s.repo.Create(item)
}

func (s *PresetService) List(status string) ([]models.Preset, error) { return s.repo.List(status) }
func (s *PresetService) Delete(id string) error                      { return s.repo.Delete(id) }
