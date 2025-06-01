package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type APIResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

func OK(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, APIResponse{
		Code:    http.StatusOK,
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func Created(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusCreated, APIResponse{
		Code:    http.StatusCreated,
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func NotFound(c *gin.Context, errors interface{}, message string) {
	c.JSON(http.StatusNotFound, APIResponse{
		Code:    http.StatusNotFound,
		Status:  "error",
		Message: message,
		Errors:  errors,
	})
}

func BadRequest(c *gin.Context, errors interface{}, message string) {
	c.JSON(http.StatusBadRequest, APIResponse{
		Code:    http.StatusBadRequest,
		Status:  "error",
		Message: message,
		Errors:  errors,
	})
}

func ServerError(c *gin.Context, errors interface{}) {
	c.JSON(http.StatusInternalServerError, APIResponse{
		Code:    http.StatusInternalServerError,
		Status:  "error",
		Message: "Kesalahan server...",
		Errors:  errors,
	})
}
