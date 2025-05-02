package repository

import (
	"context"
	"errors"

	model "github.com/CelticAlreadyUse/CMS/internal/model"
	"github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) model.CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) Create(ctx context.Context, req *model.Category) (*model.Category, error) {
	category := model.Category{
		Name:      req.Name,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
	}
	var mysqlErr *mysql.MySQLError
	err := r.db.WithContext(ctx).Create(&category).Error
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		logrus.Error(err.Error())
		return nil, errors.New("category already exist")
	}
	if err != nil {
		return nil, err
	}

	return &model.Category{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil
}

func (r *categoryRepository) GetAll(ctx context.Context) ([]model.Category, error) {
	var catsMySQL []model.Category
	if err := r.db.WithContext(ctx).Find(&catsMySQL).Error; err != nil {
		return nil, err
	}
	categories := make([]model.Category, 0, len(catsMySQL))
	for _, c := range catsMySQL {
		categories = append(categories, model.Category{
			ID:        c.ID,
			Name:      c.Name,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
	}
	return categories, nil
}

func (r *categoryRepository) GetByID(ctx context.Context, id int64) (*model.Category, error) {
	var category model.Category
	if err := r.db.WithContext(ctx).First(&category, id).Error; err != nil {
		return nil, err
	}
	return &model.Category{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil
}

func (r *categoryRepository) Update(ctx context.Context, id int64, req *model.Category) (*model.Category, error) {
	var category model.Category
	if err := r.db.WithContext(ctx).First(&category, id).Error; err != nil {
		return nil, err
	}
	category.Name = req.Name
	category.UpdatedAt = req.UpdatedAt
	if err := r.db.WithContext(ctx).Save(&category).Error; err != nil {
		return nil, err
	}
	return &model.Category{
		ID:        id,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil
}

func (r *categoryRepository) Delete(ctx context.Context, id int64) error {
	res := r.db.WithContext(ctx).Delete(&model.Category{}, id)
	if res.Error != nil {
		logrus.Println("Error deleting:", res.Error)
		return errors.New("deleting news failed")
	} else if res.RowsAffected == 0 {
		logrus.Println("No record found to delete")
		return errors.New("no record found to delete")
	}
	return nil
}
