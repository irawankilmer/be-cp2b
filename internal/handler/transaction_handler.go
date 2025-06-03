package handler

import (
	"be-cp2b/internal/dto/request"
	"be-cp2b/internal/dto/response"
	"be-cp2b/internal/usecase"
	"be-cp2b/pkg/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type TransactionHandler struct {
	usecase usecase.TransactionUsecase
}

func NewTransactionHandler(u usecase.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{u}
}

// GetAllTransactions godoc
// @Summary Get all transactions
// @Tags Transaction
// @Security BearerAuth
// @Produce json
// @Success 200 {array} response.APIResponse
// @Success 500 {array} response.APIResponse
// @Router /api/transaction [get]
func (h *TransactionHandler) GetAllTransactions(c *gin.Context) {
	transactions, err := h.usecase.GetAll()
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	response.OK(c, transactions, "Data berhasil diambil!")
}

// CreateTransaction godoc
// @Security BearerAuth
// @Summary Create new transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param request body request.TransactionRequest true "Transaction data"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/transaction [post]
func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var req request.TransactionRequest
	if !utils.ValidateInput(c, &req) {
		return
	}

	transaction, err := h.usecase.Create(req, c.GetUint("userID"))

	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	response.Created(c, transaction, "Data berhasil ditambahkan")
}

// GetTransactionByID godoc
// @Summary Get transaction by ID
// @Tags Transaction
// @Security BearerAuth
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /api/transaction/{id} [get]
func (h *TransactionHandler) GetTransactionByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction, err := h.usecase.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, err.Error(), "Data tidak ditemukan!")
		return
	}

	response.OK(c, transaction, "Data berhasil diambil!")
}

// UpdateTransaction godoc
// @Summary Update transaction
// @Tags Transaction
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Param data body request.TransactionRequest true "Transaction data"
// @Success 204 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /api/transaction/{id} [put]
func (h *TransactionHandler) UpdateTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req request.TransactionRequest
	if !utils.ValidateInput(c, &req) {
		return
	}

	var _, err = h.usecase.Update(req, uint(id))
	if err != nil {
		response.ServerError(c, err.Error())
	}

	response.NoContent(c)
}

// DeleteTransaction godoc
// @Summary Delete transaction
// @Tags Transaction
// @Security BearerAuth
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 204 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /api/transaction/{id} [delete]
func (h *TransactionHandler) DeleteTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.usecase.Delete(uint(id))
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	response.NoContent(c)
}
