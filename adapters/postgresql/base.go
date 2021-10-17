package postgresql

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Base contains common columns for all tables.
type Base struct {
	UUID      uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(*gorm.DB) {
	base.UUID = uuid.New()
	return
}
