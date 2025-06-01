package repository

import (
	"be-cp2b/internal/domain"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAll() ([]domain.Category, error)
	Create(category *domain.Category) error
	GetByID(id uint) (*domain.Category, error)
	Update(category *domain.Category) error
	Delete(category *domain.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetAll() ([]domain.Category, error) {
	var categories []domain.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) Create(category *domain.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) GetByID(id uint) (*domain.Category, error) {
	var category domain.Category
	err := r.db.First(&category, id).Error
	return &category, err
}

func (r *categoryRepository) Update(category *domain.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(category *domain.Category) error {
	return r.db.Delete(category).Error
}
