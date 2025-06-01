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
		response.ErrorResponse(c, 500, "Gagal menampilkan data!", err)
		return
	}

	response.SuccessResponse(c, 200, "Semua data ditemukan!", accounts)
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {
	var req request.AccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, 400, "Validasi input dadal!", err)
		return
	}

	account, err := h.usecase.Create(req)
	if err != nil {
		response.ErrorResponse(c, 500, "Gagal input data!", err)
		return
	}

	response.SuccessResponse(c, 401, "Data berhasil dibuat!", account)
}

func (h *AccountHandler) GetAccountByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	account, err := h.usecase.GetByID(uint(id))
	if err != nil {
		response.ErrorResponse(c, 404, "Data tidak ditemukan!", err)
		return
	}

	response.SuccessResponse(c, 200, "Data ditemukan!", account)
}

func (h *AccountHandler) UpdateAccount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req request.AccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, 400, "Input tidak valid!", err)
		return
	}

	_, err := h.usecase.Update(uint(id), req)
	if err != nil {
		response.ErrorResponse(c, 404, "Data tidak ditemukan!", err)
		return
	}

	response.SuccessResponse(c, 204, "Data berhasil di update!", "")
}

func (h *AccountHandler) DeleteAccount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.usecase.Delete(uint(id))
	if err != nil {
		response.ErrorResponse(c, 404, "Data tidak ditemukan!", err)
		return
	}

	response.SuccessResponse(c, 204, "Data berhasil dihapus!", "")
}
