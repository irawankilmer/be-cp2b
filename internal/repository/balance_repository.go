package repository

import (
	"be-cp2b/internal/domain"
	"gorm.io/gorm"
)

type BalanceRepository interface {
	GetAll(limit, offset int) ([]domain.Balance, int64, error)
	Create(balance *domain.Balance) error
	GetByID(id uint) (*domain.Balance, error)
	GetByAccountID(id uint) (*domain.Balance, error)
	Tambah(balance *domain.Balance) error
	Kurang(balance *domain.Balance) error
}

type balanceRepository struct {
	db *gorm.DB
}

func NewBalanceRepository(db *gorm.DB) BalanceRepository {
	return &balanceRepository{db}
}

func (r *balanceRepository) GetAll(limit, offset int) ([]domain.Balance, int64, error) {
	var balances []domain.Balance
	var total int64

	r.db.Model(&domain.Account{}).Count(&total)
	err := r.db.
		Preload("Account").
		Limit(limit).
		Offset(offset).
		Find(&balances).Error

	return balances, total, err
}

func (r *balanceRepository) Create(balance *domain.Balance) error {
	return r.db.Create(balance).Error
}

func (r *balanceRepository) GetByID(id uint) (*domain.Balance, error) {
	var balance domain.Balance
	err := r.db.First(&balance, id).Error
	return &balance, err
}

func (r *balanceRepository) Tambah(balance *domain.Balance) error {
	return r.db.Save(balance).Error
}

func (r *balanceRepository) Kurang(balance *domain.Balance) error {
	return r.db.Save(balance).Error
}

func (r *balanceRepository) GetByAccountID(id uint) (*domain.Balance, error) {
	var balance domain.Balance
	err := r.db.Where("account_id = ?", id).First(&balance).Error
	return &balance, err
}
