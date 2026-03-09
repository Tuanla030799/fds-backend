package admin

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(admin *User) error
	FindByEmail(email string) (*User, error)
	FindByID(id string) (*User, error)
}

type GormRepository struct{ db *gorm.DB }

func NewGormRepository(db *gorm.DB) *GormRepository { return &GormRepository{db: db} }
func (r *GormRepository) Create(admin *User) error  { return r.db.Create(admin).Error }
func (r *GormRepository) FindByEmail(email string) (*User, error) {
	var admin User
	if err := r.db.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
func (r *GormRepository) FindByID(id string) (*User, error) {
	var admin User
	if err := r.db.First(&admin, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
