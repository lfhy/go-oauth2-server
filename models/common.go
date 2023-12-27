package models

import (
	"time"
)

// MyGormModel mimixks GormModel but uses uuid's for ID, generated in go
type MyGormModel struct {
	ID        string     `gorm:"column:ID;primary_key"`
	CreatedAt time.Time  `gorm:"column:CreatedAt"`
	UpdatedAt time.Time  `gorm:"column:UpdatedAt"`
	DeletedAt *time.Time `gorm:"column:DeletedAt"`
}

// TimestampModel ...
type TimestampModel struct {
	CreatedAt time.Time  `gorm:"column:CreatedAt"`
	UpdatedAt time.Time  `gorm:"column:UpdatedAt"`
	DeletedAt *time.Time `gorm:"column:DeletedAt"`
}

// EmailTokenModel is an abstract model which can be used for objects from which
// we derive redirect emails (email confirmation, password reset and such)
type EmailTokenModel struct {
	MyGormModel
	Reference   string     `gorm:"column:Reference,unique;not null"`
	EmailSent   bool       `gorm:"column:EmailSent,index;not null"`
	EmailSentAt *time.Time `gorm:"column:EmailSentAt"`
	ExpiresAt   time.Time  `gorm:"column:ExpiresAt,index;not null"`
}
