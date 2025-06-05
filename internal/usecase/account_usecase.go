package usecase

import (
	"be-cp2b/internal/domain"
	"be-cp2b/internal/dto/request"
	"be-cp2b/internal/repository"
)

type AccountUsecase interface {
	GetAll(limit, offset int) ([]domain.Account, int64, error)
	Create(req request.AccountRequest) (*domain.Account, error)
	GetByID(id uint) (*domain.Account, error)
	Update(id uint, req request.AccountRequest) (*domain.Account, error)
	Delete(id uint) error
}

type accountUsecase struct {
	repo        repository.AccountRepository
	balanceRepo repository.BalanceRepository
}

func NewAccountUsecase(
	r repository.AccountRepository,
	b repository.BalanceRepository,
) AccountUsecase {
	return &accountUsecase{
		repo:        r,
		balanceRepo: b,
	}
}

func (u *accountUsecase) GetAll(limit, offset int) ([]domain.Account, int64, error) {
	return u.repo.GetAll(limit, offset)
}

func (u *accountUsecase) Create(req request.AccountRequest) (*domain.Account, error) {
	account := domain.Account{
		Name:         req.Name,
		Descriptions: req.Descriptions,
	}

	if err := u.repo.Create(&account); err != nil {
		return nil, err
	}

	balance := domain.Balance{
		AccountID: account.ID,
		Balance:   0,
	}

	_ = u.balanceRepo.Create(&balance)

	return &account, nil
}

func (u *accountUsecase) GetByID(id uint) (*domain.Account, error) {
	return u.repo.GetByID(id)
}

func (u *accountUsecase) Update(id uint, req request.AccountRequest) (*domain.Account, error) {
	account, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		account.Name = req.Name
	}

	if req.Descriptions != "" {
		account.Descriptions = req.Descriptions
	}

	err = u.repo.Update(account)
	return account, err
}

func (u *accountUsecase) Delete(id uint) error {
	account, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}

	return u.repo.Delete(account)
}
