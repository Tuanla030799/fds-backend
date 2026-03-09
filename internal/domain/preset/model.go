package preset

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Status string

const (
	StatusActive   Status = "active"
	StatusDraft    Status = "draft"
	StatusArchived Status = "archived"
)

type Preset struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `gorm:"size:150;not null;index" json:"name"`
	Status    Status         `gorm:"size:30;not null;default:'active';index" json:"status"`
	Note      string         `gorm:"type:text" json:"note"`
	Tags      datatypes.JSON `gorm:"type:jsonb;not null;default:'[]'" json:"tags"`
	ImageURL  string         `gorm:"size:255;not null" json:"imageUrl"`
	SortOrder int            `gorm:"not null;default:0;index" json:"sortOrder"`
}
