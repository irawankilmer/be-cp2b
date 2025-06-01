package handler

import (
	"be-cp2b/internal/dto/request"
	"be-cp2b/internal/dto/response"
	"be-cp2b/internal/usecase"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AccountHandler struct {
	usecase usecase.AccountUsecase
}

func NewAccountHandler(u usecase.AccountUsecase) *AccountHandler {
	return &AccountHandler{u}
}

func (h *AccountHandler) GetAllAccounts(c *gin.Context) {
	accounts, err := h.usecase.GetAll()
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.OK(c, accounts, "Semua data berhasil diambil!")
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var req request.AccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err, "Input tidak valid!")
		return
	}

	account, err := h.usecase.Create(req)
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.Created(c, account, "Data berhasil dibuat!")
}

func (h *AccountHandler) GetAccountByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	account, err := h.usecase.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, err, "Data tidak ditemukan!")
		return
	}

	response.OK(c, account, "Data ditemukan!")
}

func (h *AccountHandler) UpdateAccount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req request.AccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err, "Input tidak valid!")
		return
	}

	_, err := h.usecase.Update(uint(id), req)
	if err != nil {
		response.NotFound(c, err, "Data tidak ditemukan!")
		return
	}

	response.NoContent(c)
}

func (h *AccountHandler) DeleteAccount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.usecase.Delete(uint(id))
	if err != nil {
		response.NotFound(c, err, "Data tidak ditemukan!")
		return
	}

	response.NoContent(c)
}
