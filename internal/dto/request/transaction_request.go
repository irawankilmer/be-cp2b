package request

type TransactionRequest struct {
	Date            string  `json:"date" binding:"required"`
	Type            string  `json:"type" binding:"required,oneof=pemasukan pengeluaran pindah"`
	AccountID       uint    `json:"account_id" binding:"required"`
	CategoryID      uint    `json:"category_id" binding:"required"`
	TargetAccountID *uint   `json:"target_account_id"`
	Amount          float64 `json:"amount" binding:"required"`
	Descriptions    string  `json:"descriptions"`
}
