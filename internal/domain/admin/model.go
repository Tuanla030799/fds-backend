package admin

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	RoleSuperAdmin Role = "super_admin"
	RoleOperator   Role = "operator"
	RoleViewer     Role = "viewer"
)

type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Name         string         `gorm:"size:120;not null" json:"name"`
	Email        string         `gorm:"size:160;uniqueIndex;not null" json:"email"`
	PasswordHash string         `gorm:"size:255;not null" json:"-"`
	Role         Role           `gorm:"size:40;not null;default:'super_admin'" json:"role"`
}

func (r Role) IsValid() bool { return r == RoleSuperAdmin || r == RoleOperator || r == RoleViewer }
