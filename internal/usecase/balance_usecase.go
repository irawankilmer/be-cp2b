package usecase

import (
	"be-cp2b/internal/domain"
	"be-cp2b/internal/dto/request"
	"be-cp2b/internal/repository"
	"errors"
	"gorm.io/gorm"
)

type BalanceUsecase interface {
	GetAll() ([]domain.Balance, error)
	Create(req request.BalanceRequest, accountID uint) (*domain.Balance, error)
	Tambah(tx *gorm.DB, accountID uint, amount float64) (*domain.Balance, error)
	Kurang(tx *gorm.DB, accountID uint, amount float64) (*domain.Balance, error)
}

type balanceUsecase struct {
	repo repository.BalanceRepository
}

func NewBalanceUsecase(r repository.BalanceRepository) BalanceUsecase {
	return &balanceUsecase{r}
}

func (u *balanceUsecase) GetAll() ([]domain.Balance, error) {
	return u.repo.GetAll()
}

func (u *balanceUsecase) Create(req request.BalanceRequest, accountID uint) (*domain.Balance, error) {
	balance := domain.Balance{
		AccountID: accountID,
		Balance:   req.Balance,
	}

	err := u.repo.Create(&balance)
	return &balance, err
}

func (u *balanceUsecase) Tambah(tx *gorm.DB, id uint, amount float64) (*domain.Balance, error) {
	balance, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	balance.Balance += amount
	err = tx.Save(balance).Error

	return nil, err
}

func (u *balanceUsecase) Kurang(tx *gorm.DB, id uint, amount float64) (*domain.Balance, error) {
	balance, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if balance.Balance < amount {
		return nil, errors.New("Saldo tidak cukup")
	}

	balance.Balance -= amount
	err = tx.Save(balance).Error

	return nil, err
}
