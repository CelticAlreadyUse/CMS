package usecase

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/CelticAlreadyUse/CMS/internal/helper"
	"github.com/CelticAlreadyUse/CMS/internal/model"
	"github.com/sirupsen/logrus"
)

type newsUsecase struct {
	newsRepo     model.NewsRepository
	categoryRepo model.CategoryRepository
	commentsRepo model.CommentsRepository
	redisClient  model.RedisClient
	userRepo     model.UserAuthRepository
}

func NewNewsUsecase(userRepo model.UserAuthRepository, redisClient model.RedisClient, newsRepo model.NewsRepository, categoryRepo model.CategoryRepository, commentsRepo model.CommentsRepository) model.NewsUsecase {
	return &newsUsecase{
		newsRepo:     newsRepo,
		commentsRepo: commentsRepo,
		categoryRepo: categoryRepo,
		redisClient:  redisClient,
		userRepo:     userRepo,
	}
}
func (u *newsUsecase) Create(ctx context.Context, userID int64, req *model.NewsRequest) (*model.News, error) {
	_, err := u.categoryRepo.GetByID(ctx, req.CategoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}
	news := &model.News{
		UserID:     userID,
		Title:      req.Title,
		Content:    req.Content,
		CategoryID: req.CategoryID,
	}
	newNews, err := u.newsRepo.Create(ctx, news)
	if err != nil {
		return nil, err
	}
	go func() {
		err := u.redisClient.HDelByBucketKey(context.Background(), helper.NewsBucketKey)
		if err != nil {
			logrus.Errorf("failed to delete data from redis %v", err)
		}
	}()
	return newNews, nil
}

func (u *newsUsecase) GetAll(ctx context.Context) ([]model.News, error) {
	listNews, err := u.newsRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return listNews, nil
}
func (u *newsUsecase) GetByID(ctx context.Context, newsID string) (*model.NewsWithCommentsAndCategories, error) {
	intNewsID, err := strconv.Atoi(newsID)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}
	cacheKey := helper.NewNewsByIDCacheKey(int64(intNewsID))
	logrus.Info(cacheKey)
	var newsData *model.NewsWithCommentsAndCategories
	err = u.redisClient.HGet(ctx, helper.NewsBucketKey, cacheKey, &newsData)
	if err == nil && newsData != nil {
		logrus.Info("data get from redis cache")
		return newsData, nil
	}
	news, err := u.newsRepo.GetByID(ctx, int64(intNewsID))
	if err != nil {
		return nil, err
	}
	category, err := u.categoryRepo.GetByID(ctx, news.CategoryID)
	if err != nil {
		return nil, err
	}
	comments, err := u.commentsRepo.FindAllCommentsByNewsID(ctx, news.ID)
	if err != nil {
		return nil, err
	}
	user, err := u.userRepo.FindUserNameByID(ctx, news.UserID)
	if err != nil {
		return nil, err
	}
	newsData =  &model.NewsWithCommentsAndCategories{
		ID:       news.ID,
		Title:    news.Title,
		UserID:   news.UserID,
		Content:  news.Content,
		Comments: comments,
		UserName: model.User{ID: news.UserID, Username: user},
		Category: *category,
	}
	err = u.redisClient.HSet(ctx, helper.NewsBucketKey, cacheKey, newsData, 10*time.Minute)
	if err != nil {
		logrus.Info("failed set caching data")
	}
	return newsData, nil

}

func (u *newsUsecase) Update(ctx context.Context, userID int64, newsID string, req *model.NewsRequest) error {
	intNewsID, err := strconv.Atoi(newsID)
	if err != nil {
		return errors.New("invalid ID format")
	}
	_, err = u.categoryRepo.GetByID(ctx, req.CategoryID)
	if err != nil {
		return errors.New("category not found")
	}

	err = u.newsRepo.Update(ctx, userID, &model.News{
		ID:         int64(intNewsID),
		Title:      req.Title,
		Content:    req.Content,
		CategoryID: req.CategoryID})
	if err != nil {
		return err
	}
	go func() {
		err = u.redisClient.Del(context.Background(), helper.NewNewsByIDCacheKey(int64(intNewsID)))
		if err != nil {
			logrus.Errorf("failed when delete data from redis, error: %v", err)
		}
		err = u.redisClient.HDelByBucketKey(context.Background(), helper.NewsBucketKey)
		if err != nil {
			logrus.Errorf("failed when delete data from redis, error: %v", err)
		}
	}()
	return nil
}

func (u *newsUsecase) Delete(ctx context.Context, id string) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("invalid ID format")
	}
	err = u.newsRepo.Delete(ctx, int64(intID))
	if err != nil {
		return err
	}
	go func() {
		err = u.redisClient.Del(context.Background(), helper.NewNewsByIDCacheKey(int64(intID)))
		if err != nil {
			logrus.Errorf("failed when delete data from redis, error: %v", err)
		}
		err = u.redisClient.HDelByBucketKey(context.Background(), helper.NewsBucketKey)
		if err != nil {
			logrus.Errorf("failed when delete data from redis, error: %v", err)
		}
	}()
	return nil
}
