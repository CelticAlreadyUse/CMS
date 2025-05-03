package model

import (
	"context"
	"time"
)

type CustomPageRepository interface {
	Create(ctx context.Context, page *CustomPage) (*CustomPage, error)
	GetAll(ctx context.Context) ([]CustomPage, error)
	GetByID(ctx context.Context, id int64) (*CustomPage, error)
	GetByURL(ctx context.Context, url string) (*CustomPage, error)
	Update(ctx context.Context, pageID int64, userID int64, page *CustomPage) error
	Delete(ctx context.Context, id int64) error
}

type CustomPageUsecase interface {
	Create(ctx context.Context, user_id int64, page *CustomPageRequest) (*CustomPage, error)
	GetAll(ctx context.Context) ([]CustomPage, error)
	GetByID(ctx context.Context, id int64) (*CustomPage, error)
	GetByURL(ctx context.Context, url string) (*CustomPage, error)
	Update(ctx context.Context, pageID int64, userId int64,page *CustomPageRequest) error
	Delete(ctx context.Context, id int64) error
}

type CustomPageRequest struct {
	CustomUrl string `json:"custom_url" binding:"required,notblank"`
	Content   string `json:"content"`
}
type CustomPage struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	CustomUrl string    `gorm:"size:255;unique;not null" json:"custom_url"`
	Content   string    `gorm:"type:text" json:"content"`
	UserID    int64     `gorm:"index" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
