package usecase

import (
	"be-cp2b/internal/domain"
	"be-cp2b/internal/dto/request"
	"be-cp2b/internal/repository"
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
	repo repository.TransactionRepository
}

func NewTransactionUsecase(r repository.TransactionRepository) TransactionUsecase {
	return &transactionUsecase{r}
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
