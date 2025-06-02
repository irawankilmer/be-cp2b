package handler

import (
	"be-cp2b/internal/dto/request"
	"be-cp2b/internal/dto/response"
	"be-cp2b/internal/usecase"
	"be-cp2b/pkg/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CategoryHandler struct {
	usecase usecase.CategoryUsecase
}

func NewCategoryHandler(u usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{u}
}

// GetAllCategories godoc
// @Summary Get all categories
// @Tags Category
// @Security BearerAuth
// @Produce json
// @Success 200 {array} response.APIResponse
// @Success 500 {array} response.APIResponse
// @Router /api/category [get]
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.usecase.GetAll()
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	response.OK(c, categories, "Berhasil ambil semua data!")
}

// CreateCategory godoc
// @Security BearerAuth
// @Summary Create new category
// @Tags Category
// @Accept json
// @Produce json
// @Param request body request.CategoryRequest true "Category data"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /api/category [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req request.CategoryRequest

	if !utils.ValidateInput(c, &req) {
		return
	}

	category, err := h.usecase.Create(req)
	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	response.Created(c, category, "Data berhasil ditambahkan!")
}

// GetCategoryByID godoc
// @Summary Get category by ID
// @Tags Category
// @Security BearerAuth
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /api/category/{id} [get]
func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := h.usecase.GetByID(uint(id))

	if err != nil {
		response.NotFound(c, err.Error(), "Data tidak ditemukan!")
		return
	}

	response.OK(c, category, "Data berhasil diambil!")
}

// UpdateCategory godoc
// @Summary Update Category
// @Tags Category
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param data body request.CategoryRequest true "Category data"
// @Success 204 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /api/Category/{id} [put]
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req request.CategoryRequest
	if !utils.ValidateInput(c, &req) {
		return
	}

	_, err := h.usecase.Update(uint(id), req)
	if err != nil {
		response.NotFound(c, err.Error(), "Data tidak ditemukan!")
		return
	}

	response.NoContent(c)
}

// DeleteCategory godoc
// @Summary Delete category
// @Tags Category
// @Security BearerAuth
// @Produce json
// @Param id path int true "Category ID"
// @Success 204 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /api/category/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.usecase.Delete(uint(id))
	if err != nil {
		response.NotFound(c, err.Error(), "Data tidak ditemukan!")
		return
	}

	response.NoContent(c)
}
