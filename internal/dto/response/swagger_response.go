package response

type TransactionListSwaggerResponse struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    []TransactionResponse `json:"data"`
}
