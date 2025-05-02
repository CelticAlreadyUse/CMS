package repository

import (
	"errors"

	model "github.com/CelticAlreadyUse/CMS/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {

}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}

	return &user, nil
}

func (r *userRepository) Create(user *model.User) error {
	result := r.db.Create(user)
	return result.Error
}
