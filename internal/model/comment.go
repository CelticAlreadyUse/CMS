package model

import (
	"context"
	"time"
)

type CommentUsecases interface {
	Create(ctx context.Context, req *CommentRequest) (*Comment, error)
	FindAllCommentsByNewsID(ctx context.Context, newsID int64) ([]Comment, error)
	Delete(ctx context.Context, id int64) error
}
type CommentsRepository interface {
	Store(ctx context.Context, req *Comment) (*Comment, error)
	FindAllCommentsByNewsID(ctx context.Context, newsID int64) ([]Comment, error)
	Delete(ctx context.Context, id int64) error
}
type CommentRequest struct {
	Username string `json:"username"`
	NewsID   int64  `json:"news_id" biding:"required,notblank"`
	Comment  string `json:"comment" binding:"required,notblank"`
}
type Comment struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	NewsID    int64     `gorm:"index" json:"news_id" binding:"required"`
	Comment   string    `gorm:"type:text" json:"content" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}
