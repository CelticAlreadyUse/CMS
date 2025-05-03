package usecase

import (
	"context"
	"github.com/CelticAlreadyUse/CMS/internal/model"
	"github.com/sirupsen/logrus"
)

type commentUsecase struct {
	commentRepository model.CommentsRepository
	userRepo          model.UserAuthRepository
}

func CommmentUsecase(commentRepo model.CommentsRepository, userRepo model.UserAuthRepository) model.CommentUsecases {
	return &commentUsecase{
		commentRepository: commentRepo,
		userRepo:          userRepo,
	}
}
func (u *commentUsecase) Create(ctx context.Context, req *model.CommentRequest) (*model.Comment, error) {
	logger := logrus.WithFields(logrus.Fields{
		"data": req,
	})
	comment, err := u.commentRepository.Store(ctx, &model.Comment{
		Name:    req.Username,
		NewsID:  req.NewsID,
		Comment: req.Comment,
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return comment, nil
}

func (u *commentUsecase) Delete(ctx context.Context, id int64) error {
	return u.commentRepository.Delete(ctx, id)
}
func (u *commentUsecase) FindAllCommentsByNewsID(ctx context.Context, storyID int64) ([]model.Comment, error) {
	return u.commentRepository.FindAllCommentsByNewsID(ctx, storyID)
}
