package handler

import (
	"be-cp2b/internal/dto/mapper"
	"be-cp2b/internal/dto/request"
	"be-cp2b/internal/dto/response"
	"be-cp2b/internal/usecase"
	"be-cp2b/pkg/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
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
// @Success 200 {array} response.TransactionListSwaggerResponse
// @Success 500 {array} response.APIResponse
// @Router /api/transaction [get]
func (h *TransactionHandler) GetAllTransactions(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	transactions, total, err := h.usecase.GetAll(limit, offset)
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	var responseData []response.TransactionResponse
	for _, t := range transactions {
		responseData = append(responseData, mapper.MapTransactionToDTO(t))
	}

	meta := map[string]interface{}{
		"total": total,
		"page":  page,
		"limit": limit,
	}

	response.OK(c, responseData, meta)
}

// CreateTransaction godoc
// @Security BearerAuth
// @Summary Create new transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param request body request.TransactionRequest true "Transaction data"
// @Success 201 {object} response.TransactionListSwaggerResponse
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

	responseData := mapper.MapTransactionToDTO(*transaction)
	response.Created(c, responseData, "Data berhasil ditambahkan")
}

// GetTransactionByID godoc
// @Summary Get transaction by ID
// @Tags Transaction
// @Security BearerAuth
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} response.TransactionListSwaggerResponse
// @Failure 404 {object} response.APIResponse
// @Router /api/transaction/{id} [get]
func (h *TransactionHandler) GetTransactionByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction, err := h.usecase.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, err.Error(), "Data tidak ditemukan!")
		return
	}

	responseData := mapper.MapTransactionToDTO(*transaction)
	response.OK(c, responseData, "Data berhasil diambil!")
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

// GetReportDaily godoc
// @Summary Get report daily
// @Tags Reports
// @Security BearerAuth
// @Produce json
// @Param date path string false "Date (format: dd-mm-yyyy)"
// @Success 200 {array} response.TransactionListSwaggerResponse
// @Success 500 {array} response.APIResponse
// @Router /api/report/daily/{date} [get]
func (h *TransactionHandler) GetReportDaily(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	dateParam := c.Param("date")
	var date time.Time
	var err error
	if dateParam == "" {
		date = time.Now()
	} else {
		date, err = time.Parse("02-01-2006", dateParam)
		if err != nil {
			response.BadRequest(c, err.Error(), "Format tanggal harus dd-mm-yyy, contoh: 02-12-2025")
			return
		}
	}

	transactions, total, err := h.usecase.DailyReport(date, limit, offset)
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	var responseData []response.TransactionResponse
	for _, t := range transactions {
		responseData = append(responseData, mapper.MapTransactionToDTO(t))
	}

	meta := map[string]interface{}{
		"total": total,
		"page":  page,
		"limit": limit,
	}

	response.OK(c, responseData, meta)
}
