package services

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateInput(context *gin.Context, input interface{}) error {
	if err := context.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return err
	}

	return nil
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less or equal than " + fe.Param()
	case "gte":
		return "Should be greater or equal than " + fe.Param()
	}
	return "Unknown error"
}
