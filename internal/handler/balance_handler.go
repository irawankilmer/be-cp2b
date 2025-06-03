package handler

import (
	"be-cp2b/internal/dto/response"
	"be-cp2b/internal/usecase"
	"github.com/gin-gonic/gin"
)

type BalanceHandler struct {
	usecase usecase.BalanceUsecase
}

func NewBalanceHandler(u usecase.BalanceUsecase) *BalanceHandler {
	return &BalanceHandler{u}
}

// GetAllBalances godoc
// @Summary Get all balances
// @Tags Balance
// @Security BearerAuth
// @Produce json
// @Success 200 {array} response.APIResponse
// @Success 500 {array} response.APIResponse
// @Router /api/balance [get]
func (h *BalanceHandler) GetAllBalances(c *gin.Context) {
	balances, err := h.usecase.GetAll()
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	response.OK(c, balances, "Semua data berhasil diambil!")
}
