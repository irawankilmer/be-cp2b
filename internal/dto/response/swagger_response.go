package response

type TransactionListSwaggerResponse struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    []TransactionResponse `json:"data"`
}

type AccountListSwaggerResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    []AccountResponse `json:"data"`
}

type CategoryListSwaggerResponse struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    []CategoryResponse `json:"data"`
}

type BalanceListSwaggerResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    []BalanceResponse `json:"data"`
}
