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

func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.usecase.GetAll()
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.OK(c, categories, "Berhasil ambil semua data!")
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req request.CategoryRequest

	if !utils.ValidateInput(c, &req) {
		return
	}

	category, err := h.usecase.Create(req)
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.Created(c, category, "Data berhasil ditambahkan!")
}

func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := h.usecase.GetByID(uint(id))

	if err != nil {
		response.NotFound(c, err, "Data tidak ditemukan!")
		return
	}

	response.OK(c, category, "Data berhasil diambil!")
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req request.CategoryRequest
	if !utils.ValidateInput(c, &req) {
		return
	}

	_, err := h.usecase.Update(uint(id), req)
	if err != nil {
		response.NotFound(c, err, "Data tidak ditemukan!")
		return
	}

	response.NoContent(c)
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.usecase.Delete(uint(id))
	if err != nil {
		response.NotFound(c, err, "Data tidak ditemukan!")
		return
	}

	response.NoContent(c)
}
