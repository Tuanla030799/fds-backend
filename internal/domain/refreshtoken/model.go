package refreshtoken

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Token struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	AdminID   uuid.UUID      `gorm:"type:uuid;index;not null" json:"adminId"`
	TokenHash string         `gorm:"size:255;not null;uniqueIndex" json:"-"`
	ExpiresAt time.Time      `gorm:"index;not null" json:"expiresAt"`
	RevokedAt *time.Time     `gorm:"index" json:"revokedAt,omitempty"`
}
