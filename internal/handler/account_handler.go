package handler

import (
	"be-cp2b/internal/dto/mapper"
	"be-cp2b/internal/dto/request"
	"be-cp2b/internal/dto/response"
	"be-cp2b/internal/usecase"
	"be-cp2b/pkg/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AccountHandler struct {
	usecase usecase.AccountUsecase
}

func NewAccountHandler(u usecase.AccountUsecase) *AccountHandler {
	return &AccountHandler{u}
}

// GetAllAccounts godoc
// @Summary Get all accounts
// @Tags Account
// @Security BearerAuth
// @Produce json
// @Success 200 {array} response.AccountListSwaggerResponse
// @Success 500 {array} response.APIResponse
// @Router /api/account [get]
func (h *AccountHandler) GetAllAccounts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	accounts, total, err := h.usecase.GetAll(limit, offset)

	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	var responseData []response.AccountResponse
	for _, t := range accounts {
		responseData = append(responseData, mapper.MapAccountToDTO(t))
	}

	meta := map[string]interface{}{
		"total": total,
		"page":  page,
		"limit": limit,
	}

	response.OK(c, responseData, meta)
}

// CreateAccount godoc
// @Security BearerAuth
// @Summary Create new account
// @Tags Account
// @Accept json
// @Produce json
// @Param request body request.AccountRequest true "Account data"
// @Success 201 {object} response.AccountListSwaggerResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/account [post]
func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var req request.AccountRequest

	if !utils.ValidateInput(c, &req) {
		return
	}

	account, err := h.usecase.Create(req)
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	responseData := mapper.MapAccountToDTO(*account)
	response.Created(c, responseData, "Data berhasil dibuat!")
}

// GetAccountByID godoc
// @Summary Get account by ID
// @Tags Account
// @Security BearerAuth
// @Produce json
// @Param id path int true "Account ID"
// @Success 200 {object} response.AccountListSwaggerResponse
// @Failure 404 {object} response.APIResponse
// @Router /api/account/{id} [get]
func (h *AccountHandler) GetAccountByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	account, err := h.usecase.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, err.Error(), "Data tidak ditemukan!")
		return
	}

	responseData := mapper.MapAccountToDTO(*account)
	response.OK(c, responseData, "Data ditemukan!")
}

// UpdateAccount godoc
// @Summary Update account
// @Tags Account
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Param data body request.AccountRequest true "Account data"
// @Success 204 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /api/account/{id} [put]
func (h *AccountHandler) UpdateAccount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req request.AccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error(), "Input tidak valid!")
		return
	}

	_, err := h.usecase.Update(uint(id), req)
	if err != nil {
		response.NotFound(c, err.Error(), "Data tidak ditemukan!")
		return
	}

	response.NoContent(c)
}

// DeleteAccount godoc
// @Summary Delete account
// @Tags Account
// @Security BearerAuth
// @Produce json
// @Param id path int true "Account ID"
// @Success 204 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /api/account/{id} [delete]
func (h *AccountHandler) DeleteAccount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.usecase.Delete(uint(id))
	if err != nil {
		response.NotFound(c, err.Error(), "Data tidak ditemukan!")
		return
	}

	response.NoContent(c)
}
