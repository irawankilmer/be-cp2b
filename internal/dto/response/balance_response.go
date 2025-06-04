package response

import "be-cp2b/internal/dto"

type BalanceResponse struct {
	ID        uint           `json:"ID"`
	AccountID uint           `json:"AccountID"`
	Account   dto.AccountDTO `json:"Account"`
	Balance   float64        `json:"Balance"`
}
