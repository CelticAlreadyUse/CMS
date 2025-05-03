package repository

import (
	"context"
	"errors"
	"github.com/CelticAlreadyUse/CMS/internal/model"
	"github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthUserRepository(db *gorm.DB) model.UserAuthRepository {
	return &authRepository{
		db: db,
	}
}
func (r authRepository) Store(ctx context.Context, req model.User) (int64, error) {
	var mysqlErr *mysql.MySQLError
	err := r.db.Create(&req).Error
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		logrus.Error(err.Error())
		return 0, errors.New("username has already use")
	}
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	return req.ID, nil
}
func (r authRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var data model.User
	err := r.db.WithContext(ctx).
		Where("username = ?", username).
		First(&data).Error

	if err != nil {
		return nil, err
	}
	return &data, err

}
func (r *authRepository) FindUserNameByID(ctx context.Context, id int64) (string, error) {
	var user model.User

	if err := r.db.WithContext(ctx).Select("username").First(&user, id).Error; err != nil {
		return "", err
	}

	return user.Username, nil
}
