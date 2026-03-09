package repositories

import (
	"fds-backend/internal/models"

	"gorm.io/gorm"
)

type AdminRepository struct{ db *gorm.DB }

func NewAdminRepository(db *gorm.DB) *AdminRepository { return &AdminRepository{db: db} }

func (r *AdminRepository) Create(admin *models.AdminUser) error {
	return r.db.Create(admin).Error
}

func (r *AdminRepository) FindByEmail(email string) (*models.AdminUser, error) {
	var admin models.AdminUser
	if err := r.db.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
