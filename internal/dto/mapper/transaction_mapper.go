package mapper

import (
	"be-cp2b/internal/domain"
	"be-cp2b/internal/dto"
	"be-cp2b/internal/dto/response"
)

func MapTransactionToDTO(t domain.Transaction) response.TransactionResponse {
	var targetAccount dto.AccountDTO
	if t.TargetAccountID != nil {
		targetAccount = dto.AccountDTO{
			ID:   t.TargetAccount.ID,
			Name: t.TargetAccount.Name,
		}
	}

	return response.TransactionResponse{
		ID:        t.ID,
		Date:      t.Date,
		Type:      t.Type,
		AccountID: t.AccountID,
		Account: dto.AccountDTO{
			ID:   t.Account.ID,
			Name: t.Account.Name,
		},
		CategoryID: t.CategoryID,
		Category: dto.CategoryDTO{
			ID:   t.Category.ID,
			Name: t.Category.Name,
		},
		TargetAccountID: t.TargetAccountID,
		TargetAccount:   targetAccount,
		Amount:          t.Amount,
		Descriptions:    t.Descriptions,
		UserID:          t.UserID,
		User: dto.UserDTO{
			ID:   t.User.ID,
			Name: t.User.Name,
		},
	}
}
