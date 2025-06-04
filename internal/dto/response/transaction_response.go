package response

import (
	"be-cp2b/internal/dto"
	"time"
)

type TransactionResponse struct {
	ID              uint            `json:"ID"`
	Date            time.Time       `json:"Date"`
	Type            string          `json:"Type"`
	AccountID       uint            `json:"AccountID"`
	Account         dto.AccountDTO  `json:"Account"`
	CategoryID      uint            `json:"CategoryID"`
	Category        dto.CategoryDTO `json:"Category"`
	TargetAccountID *uint           `json:"TargetAccountID,omitempty"`
	TargetAccount   dto.AccountDTO  `json:"TargetAccount,omitempty"`
	Amount          float64         `json:"Amount"`
	Descriptions    string          `json:"Descriptions"`
	UserID          uint            `json:"UserID"`
	User            dto.UserDTO     `json:"User"`
}
