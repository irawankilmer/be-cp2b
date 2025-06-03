package repository

import (
	"be-cp2b/internal/domain"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetAll() ([]domain.Transaction, error)
	Create(transaction *domain.Transaction) error
	GetByID(id uint) (*domain.Transaction, error)
	Update(transaction *domain.Transaction) error
	Delete(transaction *domain.Transaction) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) Create(transaction *domain.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepository) GetAll() ([]domain.Transaction, error) {
	var transactions []domain.Transaction

	err := r.db.
		Preload("Account").
		Preload("Category").
		Preload("TargetAccount").
		Preload("User").
		Find(&transactions).Error

	return transactions, err
}

func (r *transactionRepository) GetByID(id uint) (*domain.Transaction, error) {
	var transaction domain.Transaction
	err := r.db.First(&transaction, id).Error
	return &transaction, err
}

func (r *transactionRepository) Update(transaction *domain.Transaction) error {
	return r.db.Save(transaction).Error
}

func (r *transactionRepository) Delete(transaction *domain.Transaction) error {
	return r.db.Delete(transaction).Error
}
