package repository

import (
	"context"
	"errors"
	"time"
	"github.com/CelticAlreadyUse/CMS/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type newsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) model.NewsRepository {
	return &newsRepository{
		db: db,
	}
}
func (r *newsRepository) Create(ctx context.Context, news *model.News) (*model.News, error) {
	now := time.Now()
	news.CreatedAt = now
	news.UpdatedAt = now
	if err := r.db.WithContext(ctx).Create(news).Error; err != nil {
		return nil, err
	}
	return news, nil
}

func (r *newsRepository) GetAll(ctx context.Context) ([]model.News, error) {
	var newsList []model.News
	if err := r.db.WithContext(ctx).
		Order("created_at DESC").
		Find(&newsList).Error; err != nil {
		return nil, err
	}
	logrus.Info("Sucessfully get user all news")
	return newsList, nil
}

func (r *newsRepository) GetByID(ctx context.Context, newsID int64) (*model.News, error) {
	var news model.News
	if err := r.db.WithContext(ctx).
		Where("id = ?", newsID).
		First(&news).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("news not found")
		}
		logrus.Error("get news id failed")
		return nil, err
	}
	return &news, nil
}
func (r *newsRepository) Update(ctx context.Context, userID int64, news *model.News) error {
	var existing model.News
	if err := r.db.WithContext(ctx).
		Select("id", "user_id").
		Where("id = ?", news.ID).
		First(&existing).Error; err != nil {
		return err
	}
	if existing.UserID != userID {
		return errors.New("you don't have permission to update this news")
	}
	news.UpdatedAt = time.Now()
	err := r.db.WithContext(ctx).
		Model(&model.News{}).
		Where("id = ?", news.ID).
		Updates(map[string]interface{}{
			"title":       news.Title,
			"content":     news.Content,
			"category_id": news.CategoryID,
			"updated_at":  news.UpdatedAt,
		}).Error

	if err != nil {
		logrus.Warnf("update news failed: %v", err)
		return err
	}
	return nil
}

func (r *newsRepository) Delete(ctx context.Context, id int64) error {
	res := r.db.WithContext(ctx).Delete(&model.News{}, id)
	if res.Error != nil {
		logrus.Println("Error deleting:", res.Error)
		return errors.New("deleting news failed")
	} else if res.RowsAffected == 0 {
		logrus.Println("No record found to delete")
		return errors.New("no record found to delete")
	}
	return nil
}
