package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type AdminUser struct {
	BaseModel
	Name         string `gorm:"size:120;not null" json:"name"`
	Email        string `gorm:"size:160;uniqueIndex;not null" json:"email"`
	PasswordHash string `gorm:"size:255;not null" json:"-"`
}

type DesignSubmissionStatus string

const (
	SubmissionPending   DesignSubmissionStatus = "pending_confirmation"
	SubmissionConfirmed DesignSubmissionStatus = "confirmed"
	SubmissionRejected  DesignSubmissionStatus = "rejected"
	SubmissionCompleted DesignSubmissionStatus = "completed"
)

type DesignSubmission struct {
	BaseModel
	FullName string                 `gorm:"size:150;not null" json:"fullName"`
	Address  string                 `gorm:"type:text;not null" json:"address"`
	Phone    string                 `gorm:"size:30;not null" json:"phone"`
	Note     string                 `gorm:"type:text" json:"note"`
	ImageURL string                 `gorm:"size:255;not null" json:"imageUrl"`
	Status   DesignSubmissionStatus `gorm:"size:40;not null;default:'pending_confirmation'" json:"status"`
}

type PresetStatus string

const (
	PresetActive   PresetStatus = "active"
	PresetDraft    PresetStatus = "draft"
	PresetArchived PresetStatus = "archived"
)

type Preset struct {
	BaseModel
	Name      string         `gorm:"size:150;not null" json:"name"`
	Status    PresetStatus   `gorm:"size:30;not null;default:'active'" json:"status"`
	Note      string         `gorm:"type:text" json:"note"`
	Tags      datatypes.JSON `gorm:"type:jsonb;not null;default:'[]'" json:"tags"`
	ImageURL  string         `gorm:"size:255;not null" json:"imageUrl"`
	SortOrder int            `gorm:"not null;default:0" json:"sortOrder"`
}
