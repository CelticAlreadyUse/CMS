package usecase

import (
	"context"
	"time"

	model "github.com/CelticAlreadyUse/CMS/internal/model"
)

type categoryUsecase struct {
	categoryRepo model.CategoryRepository
}

func NewCategoryUsecase(categoryRepo model.CategoryRepository) model.CategoryUsecases {
	return &categoryUsecase{
		categoryRepo: categoryRepo,
	}
}
func (u *categoryUsecase) Create(ctx context.Context, req *model.Category) (*model.Category, error) {
	now := time.Now()
	category := &model.Category{
		Name:      req.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}
	category, err := u.categoryRepo.Create(ctx, category)
	if err != nil {
		return nil, err
	}

	return category, nil
}
func (u *categoryUsecase) GetAll(ctx context.Context) ([]model.Category, error) {
	return u.categoryRepo.GetAll(ctx)
}
func (u *categoryUsecase) GetByID(ctx context.Context, id int64) (*model.Category, error) {
	return u.categoryRepo.GetByID(ctx, id)
}
func (u *categoryUsecase) Update(ctx context.Context, id int64, req *model.Category) (*model.Category, error) {
	category, err := u.categoryRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	category.Name = req.Name
	category.UpdatedAt = time.Now()
	categoryRepo, err := u.categoryRepo.Update(ctx, id, category)
	if err != nil {
		return nil, err
	}
	return categoryRepo, nil
}
func (u *categoryUsecase) Delete(ctx context.Context, id int64) error {
	err := u.categoryRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
