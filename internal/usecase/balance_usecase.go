package usecase

import (
	"be-cp2b/internal/domain"
	"be-cp2b/internal/dto/request"
	"be-cp2b/internal/repository"
)

type BalanceUsecase interface {
	GetAll() ([]domain.Balance, error)
	Create(req request.BalanceRequest, accountID uint) (*domain.Balance, error)
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
