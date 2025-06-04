package mapper

import (
	"be-cp2b/internal/domain"
	"be-cp2b/internal/dto"
	"be-cp2b/internal/dto/response"
)

func MapBalanceToDTO(t domain.Balance) response.BalanceResponse {
	return response.BalanceResponse{
		ID:        t.ID,
		AccountID: t.AccountID,
		Account: dto.AccountDTO{
			ID:   t.Account.ID,
			Name: t.Account.Name,
		},
		Balance: t.Balance,
	}
}
