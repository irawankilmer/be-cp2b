package repository

import (
	"be-cp2b/internal/domain"
	"gorm.io/gorm"
)

type AccountRepository interface {
	GetAll() ([]domain.Account, error)
	Create(account *domain.Account) error
	GetByID(id uint) (*domain.Account, error)
	Update(account *domain.Account) error
	Delete(account *domain.Account) error
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db}
}

func (r *accountRepository) GetAll() ([]domain.Account, error) {
	var accounts []domain.Account
	err := r.db.
		Preload("Balance").
		Find(&accounts).Error
	return accounts, err
}

func (r *accountRepository) Create(account *domain.Account) error {
	return r.db.Create(account).Error
}

func (r *accountRepository) GetByID(id uint) (*domain.Account, error) {
	var account domain.Account
	err := r.db.
		Preload("Balance").
		First(&account, id).Error
	return &account, err
}

func (r *accountRepository) Update(account *domain.Account) error {
	return r.db.Save(account).Error
}

func (r *accountRepository) Delete(account *domain.Account) error {
	return r.db.Delete(account).Error
}
