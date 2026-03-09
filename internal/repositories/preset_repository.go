package repositories

import (
	"fds-backend/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PresetRepository struct{ db *gorm.DB }

func NewPresetRepository(db *gorm.DB) *PresetRepository { return &PresetRepository{db: db} }

func (r *PresetRepository) Create(item *models.Preset) error {
	return r.db.Create(item).Error
}

func (r *PresetRepository) List(status string) ([]models.Preset, error) {
	var items []models.Preset
	query := r.db.Order("sort_order asc, created_at desc")
	if status != "" {
		query = query.Where("status = ?", status)
	}
	return items, query.Find(&items).Error
}

func (r *PresetRepository) Delete(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return r.db.Delete(&models.Preset{}, "id = ?", uid).Error
}
