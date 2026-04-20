package fileasset

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusUnactive Status = "UNACTIVE"
	StatusActive   Status = "ACTIVE"
)

type File struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Path      string     `gorm:"size:255;not null" json:"path"`
	Filename  string     `gorm:"size:255;not null" json:"filename"`
	MimeType  string     `gorm:"size:120;not null" json:"mimeType"`
	Size      int64      `gorm:"not null" json:"size"`
	Status    Status     `gorm:"size:20;not null;default:'UNACTIVE';index" json:"status"`
	CreatedBy *uuid.UUID `gorm:"type:uuid;index" json:"createdBy,omitempty"`
	CreatedAt time.Time  `json:"createdAt"`
}

func (File) TableName() string { return "files" }

func (s Status) IsValid() bool { return s == StatusUnactive || s == StatusActive }
