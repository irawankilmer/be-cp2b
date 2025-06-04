package response

import "be-cp2b/internal/dto"

type AccountResponse struct {
	ID           uint           `json:"ID"`
	Name         string         `json:"Name"`
	Descriptions string         `json:"Descriptions"`
	BalanceID    uint           `json:"BalanceID"`
	Balance      dto.BalanceDTO `json:"Balance"`
}
