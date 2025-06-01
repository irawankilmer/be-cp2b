package usecase

import (
	"be-cp2b/internal/domain"
	"be-cp2b/internal/dto/request"
	"be-cp2b/internal/repository"
)

type CategoryUsecase interface {
	GetAll() ([]domain.Category, error)
	Create(req request.CategoryRequest) (*domain.Category, error)
	GetByID(id uint) (*domain.Category, error)
	Update(id uint, req request.CategoryRequest) (*domain.Category, error)
	Delete(id uint) error
}

type categoryUsecase struct {
	repo repository.CategoryRepository
}

func NewCategoryUsecase(r repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{r}
}

func (u *categoryUsecase) GetAll() ([]domain.Category, error) {
	return u.repo.GetAll()
}

func (u *categoryUsecase) Create(req request.CategoryRequest) (*domain.Category, error) {
	category := domain.Category{
		Name:         req.Name,
		Type:         req.Type,
		Descriptions: req.Descriptions,
	}

	err := u.repo.Create(&category)
	return &category, err
}

func (u *categoryUsecase) GetByID(id uint) (*domain.Category, error) {
	return u.repo.GetByID(id)
}

func (u *categoryUsecase) Update(id uint, req request.CategoryRequest) (*domain.Category, error) {
	category, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		category.Name = req.Name
	}

	if req.Type != "" {
		category.Type = req.Type
	}

	if req.Descriptions != "" {
		category.Descriptions = req.Descriptions
	}

	err = u.repo.Update(category)
	return category, err
}

func (u *categoryUsecase) Delete(id uint) error {
	category, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}
	
	return u.repo.Delete(category)
}
