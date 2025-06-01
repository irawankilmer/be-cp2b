package repository

import (
	"be-cp2b/internal/domain"
	"gorm.io/gorm"
)

type BalanceRepository interface {
	Create(balance *domain.Balance) error
}

type balanceRepository struct {
	db *gorm.DB
}

func NewBalanceRepository(db *gorm.DB) BalanceRepository {
	return &balanceRepository{db}
}

func (r *balanceRepository) Create(balance *domain.Balance) error {
	return r.db.Create(balance).Error
}
