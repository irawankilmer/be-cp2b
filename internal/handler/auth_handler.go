package handler

import (
	"be-cp2b/internal/dto/request"
	"be-cp2b/internal/dto/response"
	"be-cp2b/internal/usecase"
	"be-cp2b/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	usecase usecase.AuthUsecase
}

func NewAuthHandler(u usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{u}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req request.AuthRequest
	if !utils.ValidateInput(c, &req) {
		return
	}

	token, err := h.usecase.Login(req)
	if err != nil {
		response.UnAuthorized(c, err.Error(), "Email atau password salah!")
		return
	}

	response.OK(c, token, "Selamat, Anda berhasil login!")
}
