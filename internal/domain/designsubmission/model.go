package designsubmission

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Status string

const (
	StatusPending   Status = "pending_confirmation"
	StatusConfirmed Status = "confirmed"
	StatusRejected  Status = "rejected"
	StatusCompleted Status = "completed"
)

type Submission struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	FullName  string         `gorm:"size:150;not null;index" json:"fullName"`
	Address   string         `gorm:"type:text;not null" json:"address"`
	Phone     string         `gorm:"size:30;not null;index" json:"phone"`
	Note      string         `gorm:"type:text" json:"note"`
	ImageURL  string         `gorm:"size:255;not null" json:"imageUrl"`
	Status    Status         `gorm:"size:40;not null;default:'pending_confirmation';index" json:"status"`
}
