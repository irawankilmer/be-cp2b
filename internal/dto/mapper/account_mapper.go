package mapper

import (
	"be-cp2b/internal/domain"
	"be-cp2b/internal/dto"
	"be-cp2b/internal/dto/response"
)

func MapAccountToDTO(t domain.Account) response.AccountResponse {
	return response.AccountResponse{
		ID:           t.ID,
		Name:         t.Name,
		Descriptions: t.Descriptions,
		BalanceID:    t.Balance.ID,
		Balance: dto.BalanceDTO{
			ID:      t.Balance.ID,
			Balance: t.Balance.Balance,
		},
	}
}
