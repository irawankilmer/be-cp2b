package utils

import (
	"be-cp2b/internal/dto/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func ValidateInput(c *gin.Context, input interface{}) bool {
	if err := c.ShouldBindJSON(input); err != nil {
		var verr validator.ValidationErrors

		if ok := AsValidationErrors(err, &verr); ok {
			errorMessages := make(map[string]string)

			for _, e := range verr {
				field := FieldJSONName(input, e.Field())
				msg := ValidationMessage(field, e.Tag(), e.Param())
				errorMessages[field] = msg
			}

			response.BadRequest(c, gin.H{
				"errors": errorMessages,
				"old":    input,
			}, "Validasi gagal!")
			return false
		}

		response.BadRequest(c, "", "Format input tidak valid")
	}

	return true
}

func AsValidationErrors(err error, target *validator.ValidationErrors) bool {
	ve, ok := err.(validator.ValidationErrors)
	if !ok {
		return false
	}

	*target = ve
	return true
}

func ValidationMessage(field, tag, param string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("Kolom %s wajib diisi", field)
	case "email":
		return fmt.Sprintf("Kolom %s harus berupa email yang valid", field)
	case "min":
		return fmt.Sprintf("Kolom %s minimal %s karakter", field, param)
	default:
		return fmt.Sprintf("Kolom %s tidak valid", field)
	}
}

func FieldJSONName(input interface{}, field string) string {
	t := reflect.TypeOf(input)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if f, ok := t.FieldByName(field); ok {
		jsonTag := f.Tag.Get("json")
		return strings.Split(jsonTag, ",")[0]
	}

	return field
}
