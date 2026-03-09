package auditlog

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AuditLog struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	AdminID     *uuid.UUID     `gorm:"type:uuid;index" json:"adminId,omitempty"`
	Action      string         `gorm:"size:100;not null;index" json:"action"`
	Resource    string         `gorm:"size:100;not null;index" json:"resource"`
	ResourceID  string         `gorm:"size:120;index" json:"resourceId"`
	RequestID   string         `gorm:"size:120;index" json:"requestId"`
	IP          string         `gorm:"size:80" json:"ip"`
	UserAgent   string         `gorm:"size:255" json:"userAgent"`
	BeforeState datatypes.JSON `gorm:"type:jsonb" json:"beforeState,omitempty"`
	AfterState  datatypes.JSON `gorm:"type:jsonb" json:"afterState,omitempty"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
