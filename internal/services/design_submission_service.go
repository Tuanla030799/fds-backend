package services

import (
	"fds-backend/internal/models"
	"fds-backend/internal/repositories"
)

type DesignSubmissionService struct {
	repo *repositories.DesignSubmissionRepository
}

func NewDesignSubmissionService(repo *repositories.DesignSubmissionRepository) *DesignSubmissionService {
	return &DesignSubmissionService{repo: repo}
}

func (s *DesignSubmissionService) Create(item *models.DesignSubmission) error {
	return s.repo.Create(item)
}
func (s *DesignSubmissionService) List(status string) ([]models.DesignSubmission, error) {
	return s.repo.List(status)
}
func (s *DesignSubmissionService) FindByID(id string) (*models.DesignSubmission, error) {
	return s.repo.FindByID(id)
}
func (s *DesignSubmissionService) Update(item *models.DesignSubmission) error {
	return s.repo.Update(item)
}
func (s *DesignSubmissionService) Delete(id string) error { return s.repo.Delete(id) }
