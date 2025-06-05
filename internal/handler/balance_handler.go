package handler

import (
	"be-cp2b/internal/dto/mapper"
	"be-cp2b/internal/dto/response"
	"be-cp2b/internal/usecase"
	"github.com/gin-gonic/gin"
	"strconv"
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
// @Success 200 {array} response.BalanceListSwaggerResponse
// @Success 500 {array} response.APIResponse
// @Router /api/balance [get]
func (h *BalanceHandler) GetAllBalances(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	balances, total, err := h.usecase.GetAll(limit, offset)
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	var responseData []response.BalanceResponse
	for _, t := range balances {
		responseData = append(responseData, mapper.MapBalanceToDTO(t))
	}

	meta := map[string]interface{}{
		"total": total,
		"page":  page,
		"limit": limit,
	}

	response.OK(c, responseData, meta)
}
