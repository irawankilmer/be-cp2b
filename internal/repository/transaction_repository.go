package repository

import (
	"be-cp2b/internal/domain"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetAll(limit, offset int) ([]domain.Transaction, int64, error)
	Create(transaction *domain.Transaction) error
	GetByID(id uint) (*domain.Transaction, error)
	Update(transaction *domain.Transaction) error
	Delete(transaction *domain.Transaction) error
	GetDB() *gorm.DB
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

func (r *transactionRepository) GetAll(limit, offset int) ([]domain.Transaction, int64, error) {
	var transactions []domain.Transaction
	var total int64

	r.db.Model(&domain.Transaction{}).Count(&total)

	err := r.db.
		Preload("Account").
		Preload("Category").
		Preload("TargetAccount").
		Preload("User").
		Limit(limit).
		Offset(offset).
		Find(&transactions).Error

	return transactions, total, err
}

func (r *transactionRepository) GetByID(id uint) (*domain.Transaction, error) {
	var transaction domain.Transaction
	err := r.db.
		Preload("Account").
		Preload("Category").
		Preload("TargetAccount").
		Preload("User").
		First(&transaction, id).Error
	return &transaction, err
}

func (r *transactionRepository) Update(transaction *domain.Transaction) error {
	return r.db.Save(transaction).Error
}

func (r *transactionRepository) Delete(transaction *domain.Transaction) error {
	return r.db.Delete(transaction).Error
}

func (r *transactionRepository) GetDB() *gorm.DB {
	return r.db
}
