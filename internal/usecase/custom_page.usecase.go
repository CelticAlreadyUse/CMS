package usecase

import (
	"context"

	"github.com/CelticAlreadyUse/CMS/internal/model"
)

type customPageUsecase struct {
	customPageRepo model.CustomPageRepository
}

func NewCustomPageUsecase(repo model.CustomPageRepository) model.CustomPageUsecase {
	return &customPageUsecase{customPageRepo: repo}
}

func (u *customPageUsecase) Create(ctx context.Context, userID int64, page *model.CustomPageRequest) (*model.CustomPage, error) {
	return u.customPageRepo.Create(ctx, &model.CustomPage{
		CustomUrl: page.CustomUrl,
		 Content: page.Content,
		 UserID: userID,
	})
}

func (u *customPageUsecase) GetAll(ctx context.Context) ([]model.CustomPage, error) {
	return u.customPageRepo.GetAll(ctx)
}

func (u *customPageUsecase) GetByID(ctx context.Context, id int64) (*model.CustomPage, error) {
	return u.customPageRepo.GetByID(ctx, id)
}

func (u *customPageUsecase) GetByURL(ctx context.Context, url string) (*model.CustomPage, error) {
	return u.customPageRepo.GetByURL(ctx, url)
}

func (u *customPageUsecase) Update(ctx context.Context, id int64, userId int64, page *model.CustomPageRequest) error {
	return u.customPageRepo.Update(ctx, id, userId, &model.CustomPage{
		CustomUrl: page.CustomUrl,
		Content:   page.Content,
	})
}

func (u *customPageUsecase) Delete(ctx context.Context, id int64) error {
	return u.customPageRepo.Delete(ctx, id)
}
