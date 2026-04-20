package fileasset

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	Create(item *File) error
	FindByID(id string) (*File, error)
	FindByIDForUpdate(tx *gorm.DB, id string) (*File, error)
	DeleteByID(id uuid.UUID) error
	ListExpiredInactive(before time.Time) ([]File, error)
}

type GormRepository struct{ db *gorm.DB }

func NewGormRepository(db *gorm.DB) *GormRepository { return &GormRepository{db: db} }

func (r *GormRepository) Create(item *File) error { return r.db.Create(item).Error }

func (r *GormRepository) FindByID(id string) (*File, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	var item File
	if err := r.db.First(&item, "id = ?", uid).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *GormRepository) FindByIDForUpdate(tx *gorm.DB, id string) (*File, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	var item File
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&item, "id = ?", uid).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *GormRepository) DeleteByID(id uuid.UUID) error {
	return r.db.Delete(&File{}, "id = ?", id).Error
}

func (r *GormRepository) ListExpiredInactive(before time.Time) ([]File, error) {
	var items []File
	err := r.db.Where("status = ? AND created_at <= ?", StatusUnactive, before).Find(&items).Error
	return items, err
}
