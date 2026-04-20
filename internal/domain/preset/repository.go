package preset

import (
	"strings"

	"fds-backend/internal/shared/pagination"
	"fds-backend/internal/shared/query"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(item *Preset) error
	CreateTx(tx *gorm.DB, item *Preset) error
	List(filters query.Filters, sort query.Sort, params pagination.Params) ([]Preset, int64, error)
	Delete(id string) error
}

type GormRepository struct{ db *gorm.DB }

func NewGormRepository(db *gorm.DB) *GormRepository { return &GormRepository{db: db} }
func (r *GormRepository) Create(item *Preset) error { return r.db.Create(item).Error }
func (r *GormRepository) CreateTx(tx *gorm.DB, item *Preset) error {
	return tx.Create(item).Error
}
func (r *GormRepository) List(filters query.Filters, sort query.Sort, params pagination.Params) ([]Preset, int64, error) {
	var items []Preset
	var total int64
	q := r.db.Model(&Preset{})
	if filters.Status != "" {
		q = q.Where("status = ?", filters.Status)
	}
	if filters.Keyword != "" {
		kw := "%" + strings.ToLower(filters.Keyword) + "%"
		q = q.Where("LOWER(name) LIKE ? OR LOWER(note) LIKE ?", kw, kw)
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := q.Order(sort.Field + " " + sort.Order).Offset(params.Offset).Limit(params.Limit).Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
func (r *GormRepository) Delete(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return r.db.Delete(&Preset{}, "id = ?", uid).Error
}
