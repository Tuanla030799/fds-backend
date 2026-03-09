package refreshtoken

import "gorm.io/gorm"

type Repository interface {
	Create(*Token) error
	FindByHash(hash string) (*Token, error)
	Update(*Token) error
	RevokeAllByAdminID(adminID string) error
}

type GormRepository struct{ db *gorm.DB }

func NewGormRepository(db *gorm.DB) *GormRepository { return &GormRepository{db: db} }
func (r *GormRepository) Create(item *Token) error  { return r.db.Create(item).Error }
func (r *GormRepository) FindByHash(hash string) (*Token, error) {
	var item Token
	if err := r.db.Where("token_hash = ?", hash).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
func (r *GormRepository) Update(item *Token) error { return r.db.Save(item).Error }
func (r *GormRepository) RevokeAllByAdminID(adminID string) error {
	now := gorm.Expr("NOW()")
	return r.db.Model(&Token{}).Where("admin_id = ? AND revoked_at IS NULL", adminID).Update("revoked_at", now).Error
}
