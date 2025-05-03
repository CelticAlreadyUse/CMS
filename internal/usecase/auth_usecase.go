package usecase

import (
	"context"
	"errors"

	"github.com/CelticAlreadyUse/CMS/internal/config"
	"github.com/CelticAlreadyUse/CMS/internal/helper"
	"github.com/CelticAlreadyUse/CMS/internal/model"
	"github.com/sirupsen/logrus"
)

type authUsecase struct {
	userRepo model.UserAuthRepository
}

func NewAuthUsecase(userRepo model.UserAuthRepository) model.UserAuthUsecase {
	return &authUsecase{
		userRepo: userRepo,
	}
}
func (u *authUsecase) Login(ctx context.Context, req model.LoginRequest) (string, error) {
	logger := logrus.WithFields(logrus.Fields{
		"function": "Login",
		"username": req.Username,
	})
	
	// Validate input
	if req.Username == "" || req.Password == "" {
		logger.Warn("login attempt with empty credentials")
		return "", errors.New("invalid credentials")
	}
	
	user, err := u.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		logger.WithError(err).Error("failed to fetch user")
		return "", errors.New("authentication failed")
	}
	if user == nil {
		logger.Warn("login attempt for non-existent user")
		return "", errors.New("invalid credentials")
	}
	if !helper.CheckPasswword(req.Password, user.HashPassword) {
		logger.WithField("userID", user.ID).Warn("password mismatch")
		return "", errors.New("wrong password")
	}
	token, err := helper.GenerateToken(user.ID)
	if err != nil {
		logger.WithError(err).Error("failed to generate token")
		return "", errors.New("failed to generate authentication token")
	}
	config := model.ConfigJWT{
		SigningKey: config.JWTSigningKey(),
		ExpTime:    config.JWTExp().String(),
	}
	_, err = helper.ValidateToken(token, config)
	if err != nil {
		logger.WithError(err).Error("generated invalid token")
		return "", errors.New("authentication system error")
	}
	logger.WithField("userID", user.ID).Info("successful login")
	return token, nil
}
func (u *authUsecase)Register(ctx context.Context,req model.RegisterRequest)(string,error){
	logger := logrus.WithFields(logrus.Fields{
		"data" : req,
	})
	passwordHashed,err := helper.Hashpassword(req.Password)
	if err !=nil{
		logger.Error(err.Error())
		return "",err
	}
	newUserId,err := u.userRepo.Store(ctx,model.User{Username: req.Username,HashPassword: passwordHashed})
	if err !=nil{
		logger.Error(err.Error())
		return "",err
	}
	accessToken,err := helper.GenerateToken(newUserId)
	if err !=nil{
		logger.Error("Generate Token Failed")
		return "",err
	}
	return accessToken,nil
}
func (u *authUsecase)FindUserNameByID(ctx context.Context,id int64)(string,error){
	return u.FindUserNameByID(ctx,id)
}