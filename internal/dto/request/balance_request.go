package request

type BalanceRequest struct {
	AccountID uint    `json:"account_id" binding:"required"`
	Balance   float64 `json:"balance" binding:"required"`
}
