package model

import (
	"context"
	"time"
)

type NewsRepository interface {
	Create(ctx context.Context, news *News) (*News, error)
	GetAll(ctx context.Context) ([]News, error)
	GetByID(ctx context.Context, newsID int64) (*News, error)
	Update(ctx context.Context, userID int64, news *News) error
	Delete(ctx context.Context, id int64) error
}
type NewsUsecase interface {
	Create(ctx context.Context, userID int64, news *NewsRequest) (*News, error)
	GetAll(ctx context.Context) ([]News, error)
	GetByID(ctx context.Context, newsID string) (*NewsWithCommentsAndCategories, error)
	Update(ctx context.Context, userID int64, id string, news *NewsRequest) error
	Delete(ctx context.Context, id string) error
}

type News struct {
	ID         int64     `gorm:"primaryKey" json:"id"`
	UserID     int64     `gorm:"index" json:"user_id" binding:"required"`
	Title      string    `gorm:"type:varchar(255)" json:"title" binding:"required"`
	Content    string    `gorm:"type:mediumtext" json:"content" binding:"required"`
	CategoryID int64     `gorm:"index" json:"category_id" binding:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
type NewsWithCommentsAndCategories struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title"`
	UserID     int64     `json:"user_id"`
	UserName   User `json:"username"`
	Content    string    `json:"content"`
	Comments   []Comment `json:"comments,omitempty"`
	Category   Category  `json:"category,omitempty"`
	CategoryID int64     `json:"category_id,omitempty" binding:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
type NewsRequest struct {
	ID         int64  `json:"id"`
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	CategoryID int64  `json:"category_id" binding:"required"`
}
