package usecase

import (
	"be-cp2b/internal/domain"
	"be-cp2b/internal/dto/request"
	"be-cp2b/internal/repository"
	"errors"
	"time"
)

type TransactionUsecase interface {
	GetAll() ([]domain.Transaction, error)
	Create(req request.TransactionRequest, userID uint) (*domain.Transaction, error)
	GetByID(id uint) (*domain.Transaction, error)
	Update(req request.TransactionRequest, id uint) (*domain.Transaction, error)
	Delete(id uint) error
}

type transactionUsecase struct {
	repo           repository.TransactionRepository
	balanceUsecase BalanceUsecase
}

func NewTransactionUsecase(
	r repository.TransactionRepository,
	bu BalanceUsecase,
) TransactionUsecase {
	return &transactionUsecase{
		repo:           r,
		balanceUsecase: bu,
	}
}

func (u *transactionUsecase) GetAll() ([]domain.Transaction, error) {
	return u.repo.GetAll()
}

func (u *transactionUsecase) Create(req request.TransactionRequest, userID uint) (*domain.Transaction, error) {
	// parsing tanggal
	dateStr := req.Date
	parseDate, _ := time.Parse("2006-01-02", dateStr)
	transaction := domain.Transaction{
		Date:            parseDate,
		Type:            req.Type,
		AccountID:       req.AccountID,
		CategoryID:      req.CategoryID,
		TargetAccountID: req.TargetAccountID,
		Amount:          req.Amount,
		Descriptions:    req.Descriptions,
		UserID:          userID,
	}

	if err := u.repo.Create(&transaction); err != nil {
		return nil, err
	}

	switch req.Type {
	case "pemasukan":
		_, err := u.balanceUsecase.Tambah(req.AccountID, req.Amount)
		if err != nil {
			return nil, err
		}
	case "pengeluaran":
		_, err := u.balanceUsecase.Kurang(req.AccountID, req.Amount)
		if err != nil {
			return nil, err
		}
	case "pindah":
		_, err := u.balanceUsecase.Kurang(req.AccountID, req.Amount)
		if err != nil {
			return nil, err
		}

		if req.TargetAccountID == nil {
			return nil, errors.New("Target account id tidak boleh kosong untuk transaksi pindah")
		}
		_, eri := u.balanceUsecase.Tambah(*req.TargetAccountID, req.Amount)
		if eri != nil {
			return nil, eri
		}
	}

	return &transaction, nil
}

func (u *transactionUsecase) GetByID(id uint) (*domain.Transaction, error) {
	return u.repo.GetByID(id)
}

func (u *transactionUsecase) Update(req request.TransactionRequest, id uint) (*domain.Transaction, error) {
	transaction, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	//parsing tanggal
	dateStr := req.Date
	parseDate, _ := time.Parse("2006-01-02", dateStr)

	transaction.Date = parseDate
	transaction.Type = req.Type
	transaction.AccountID = uint(req.AccountID)
	transaction.CategoryID = uint(req.CategoryID)
	transaction.TargetAccountID = req.TargetAccountID
	transaction.Amount = req.Amount
	transaction.Descriptions = req.Descriptions
	_ = u.repo.Update(transaction)

	return transaction, err
}

func (u *transactionUsecase) Delete(id uint) error {
	transaction, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}

	return u.repo.Delete(transaction)
}
