package designsubmission

import (
	"strings"

	"fds-backend/internal/shared/pagination"
	"fds-backend/internal/shared/query"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(item *Submission) error
	List(filters query.Filters, sort query.Sort, params pagination.Params) ([]Submission, int64, error)
	FindByID(id string) (*Submission, error)
	Update(item *Submission) error
	Delete(id string) error
}

type GormRepository struct{ db *gorm.DB }

func NewGormRepository(db *gorm.DB) *GormRepository     { return &GormRepository{db: db} }
func (r *GormRepository) Create(item *Submission) error { return r.db.Create(item).Error }
func (r *GormRepository) List(filters query.Filters, sort query.Sort, params pagination.Params) ([]Submission, int64, error) {
	var items []Submission
	var total int64
	q := r.db.Model(&Submission{})
	if filters.Status != "" {
		q = q.Where("status = ?", filters.Status)
	}
	if filters.Keyword != "" {
		kw := "%" + strings.ToLower(filters.Keyword) + "%"
		q = q.Where("LOWER(full_name) LIKE ? OR LOWER(phone) LIKE ?", kw, kw)
	}
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := q.Order(sort.Field + " " + sort.Order).Offset(params.Offset).Limit(params.Limit).Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
func (r *GormRepository) FindByID(id string) (*Submission, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	var item Submission
	if err := r.db.First(&item, "id = ?", uid).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
func (r *GormRepository) Update(item *Submission) error { return r.db.Save(item).Error }
func (r *GormRepository) Delete(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return r.db.Delete(&Submission{}, "id = ?", uid).Error
}
