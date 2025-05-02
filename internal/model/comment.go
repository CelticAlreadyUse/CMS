package model

import "time"

type Comment struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	NewsID    int64     `gorm:"index" json:"news_id" binding:"required"`
	UserID    int64     `json:"email" binding:"required,email"`
	Content   string    `gorm:"type:text" json:"content" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
