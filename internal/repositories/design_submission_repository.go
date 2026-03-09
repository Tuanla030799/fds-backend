package repositories

import (
	"fds-backend/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DesignSubmissionRepository struct{ db *gorm.DB }

func NewDesignSubmissionRepository(db *gorm.DB) *DesignSubmissionRepository {
	return &DesignSubmissionRepository{db: db}
}

func (r *DesignSubmissionRepository) Create(item *models.DesignSubmission) error {
	return r.db.Create(item).Error
}

func (r *DesignSubmissionRepository) List(status string) ([]models.DesignSubmission, error) {
	var items []models.DesignSubmission
	query := r.db.Order("created_at desc")
	if status != "" {
		query = query.Where("status = ?", status)
	}
	return items, query.Find(&items).Error
}

func (r *DesignSubmissionRepository) FindByID(id string) (*models.DesignSubmission, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	var item models.DesignSubmission
	if err := r.db.First(&item, "id = ?", uid).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *DesignSubmissionRepository) Update(item *models.DesignSubmission) error {
	return r.db.Save(item).Error
}

func (r *DesignSubmissionRepository) Delete(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return r.db.Delete(&models.DesignSubmission{}, "id = ?", uid).Error
}
