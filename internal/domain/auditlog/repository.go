package auditlog

import "gorm.io/gorm"

type Repository interface{ Create(*AuditLog) error }

type GormRepository struct{ db *gorm.DB }

func NewGormRepository(db *gorm.DB) *GormRepository   { return &GormRepository{db: db} }
func (r *GormRepository) Create(item *AuditLog) error { return r.db.Create(item).Error }
