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

// Login godoc
// @Summary Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body request.AuthRequest true "Login data"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 401 {object} response.APIResponse
// @Router /api/login [post]
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

	response.OK(c, token, "")
}

// Logout godoc
// @Summary Logout user
// @Tags Auth
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.APIResponse
// @Success 401 {object} response.APIResponse
// @Success 500 {object} response.APIResponse
// @Router /api/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	userIDAny, exists := c.Get("userID")
	if !exists {
		response.UnAuthorized(c, "User tidak ditemukan", "")
		return
	}

	userID := userIDAny.(uint)

	err := h.usecase.Logout(userID)
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	response.OK(c, "", "Logout berhasil")
}
