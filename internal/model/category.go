package model

import (
	"context"
	"time"
)

type CategoryUsecases interface {
	Create(ctx context.Context, req *Category) (*Category, error)
	GetAll(ctx context.Context) ([]Category, error)
	GetByID(ctx context.Context, id int64) (*Category, error)
	Update(ctx context.Context, id int64, req *Category) (*Category, error)
	Delete(ctx context.Context, id int64) error
}

type CategoryRepository interface {
	Create(ctx context.Context, req *Category) (*Category, error)
	GetAll(ctx context.Context) ([]Category, error)
	GetByID(ctx context.Context, id int64) (*Category, error)
	Update(ctx context.Context, id int64, req *Category) (*Category, error)
	Delete(ctx context.Context, id int64) error
}

type Category struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"not null" json:"name" binding:"notblank,lowercase"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}
