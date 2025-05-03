package repository

import (
	"context"
	"errors"
	"time"

	"github.com/CelticAlreadyUse/CMS/internal/model"
	"github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewcommentRepository(db *gorm.DB) model.CommentsRepository {
	return &commentRepository{
		db: db,
	}
}
func (r *commentRepository) Store(ctx context.Context, req *model.Comment) (*model.Comment, error) {
	req.CreatedAt = time.Now()
	var mysqlErr *mysql.MySQLError
	err := r.db.Create(req).Error
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1452 {
		logrus.Error(err.Error())
		return nil, errors.New("news id not found")
	}
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	return req, nil

}
func (r *commentRepository) FindAllCommentsByNewsID(ctx context.Context, newsId int64) ([]model.Comment, error) {
	var comments []model.Comment
	if err := r.db.WithContext(ctx).
		Where("news_id = ?", newsId).
		Order("created_at ASC").
		Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
func (r *commentRepository) Delete(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Delete(&model.Comment{}, id).Error; err != nil {
		return err
	}
	return nil
}
