package repository

import (
	"context"

	"github.com/CelticAlreadyUse/CMS/internal/model"
	"gorm.io/gorm"
)

type customPageRepository struct {
	db *gorm.DB
}

func NewCustomPageRepository(db *gorm.DB) model.CustomPageRepository {
	return &customPageRepository{
		db: db,
	}
}

func (r *customPageRepository) Create(ctx context.Context, page *model.CustomPage) (*model.CustomPage, error) {
	if err := r.db.WithContext(ctx).Create(page).Error; err != nil {
		return nil, err
	}
	return page, nil
}

func (r *customPageRepository) GetAll(ctx context.Context) ([]model.CustomPage, error) {
	var pages []model.CustomPage
	if err := r.db.WithContext(ctx).Find(&pages).Error; err != nil {
		return nil, err
	}
	return pages, nil
}

func (r *customPageRepository) GetByID(ctx context.Context, id int64) (*model.CustomPage, error) {
	var page model.CustomPage
	if err := r.db.WithContext(ctx).First(&page, id).Error; err != nil {
		return nil, err
	}
	return &page, nil
}

func (r *customPageRepository) GetByURL(ctx context.Context, url string) (*model.CustomPage, error) {
	var page model.CustomPage
	if err := r.db.WithContext(ctx).Where("url = ?", url).First(&page).Error; err != nil {
		return nil, err
	}
	return &page, nil
}

func (r *customPageRepository) Update(ctx context.Context, id int64,userID int64, page *model.CustomPage) error {
	if err := r.db.WithContext(ctx).Model(&model.CustomPage{}).Where("id = ?", id).Updates(page).Error; err != nil {
		return err
	}
	return nil
}

func (r *customPageRepository) Delete(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Delete(&model.CustomPage{}, id).Error; err != nil {
		return err
	}
	return nil
}
