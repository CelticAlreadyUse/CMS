package model

import (
	"time"

	"gorm.io/gorm"
)

type CustomPage struct {
	ID        int64          `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"size:255;not null" json:"title"`
	URL       string         `gorm:"size:255;unique;not null" json:"url"`
	Content   string         `gorm:"type:text" json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
