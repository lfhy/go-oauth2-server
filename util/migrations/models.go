package migrations

import (
	"gorm.io/gorm"
)

// Migration represents a single database migration
type Migration struct {
	gorm.Model
	Name string `gorm:"size:255"`
}
