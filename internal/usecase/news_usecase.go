package usecase

import (
	"context"
	"errors"
	"strconv"

	"github.com/CelticAlreadyUse/CMS/internal/model"
)

type newsUsecase struct {
	newsRepo     model.NewsRepository
	categoryRepo model.CategoryRepository
}

func NewNewsUsecase(newsRepo model.NewsRepository, categoryRepo model.CategoryRepository) model.NewsUsecase {
	return &newsUsecase{
		newsRepo:     newsRepo,
		categoryRepo: categoryRepo,
	}
}
func (u *newsUsecase) Create(ctx context.Context, req *model.NewsRequest) (*model.News, error) {
	_, err := u.categoryRepo.GetByID(ctx, req.CategoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}
	news := &model.News{
		Title:      req.Title,
		Content:    req.Content,
		CategoryID: req.CategoryID,
	}
	newNews, err := u.newsRepo.Create(ctx, news)
	if err != nil {
		return nil, err
	}
	return newNews, nil
}

func (u *newsUsecase) GetAll(ctx context.Context) ([]model.News, error) {
	return u.newsRepo.GetAll(ctx)
}

func (u *newsUsecase) GetByID(ctx context.Context, id string) (*model.News, error) {
	int64ID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}
	return u.newsRepo.GetByID(ctx, int64(int64ID))
}

func (u *newsUsecase) Update(ctx context.Context, id string, req *model.News)  error {
	int64ID, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("invalid ID format")
	}
	_, err = u.categoryRepo.GetByID(ctx, req.CategoryID)
	if err != nil {
		return  errors.New("category not found")
	}
	err = u.newsRepo.Update(ctx, int64(int64ID), req)
	if err != nil {
		return  err
	}
	return nil
}

func (u *newsUsecase) Delete(ctx context.Context, id string) error {
	int64ID, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	return u.newsRepo.Delete(ctx, int64(int64ID))
}
