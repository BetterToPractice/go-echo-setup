package response

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Code    int         `json:"-"`
	Data    interface{} `json:"data,omitempty"`
	Message interface{} `json:"message"`
}

func (r Response) JSON(ctx echo.Context) error {
	if r.Code == 0 {
		r.Code = http.StatusOK
	}
	if r.Message == "" || r.Message == nil {
		r.Message = http.StatusText(r.Code)
	}

	if err, ok := r.Message.(validator.ValidationErrors); ok {
		r.Message = extractValidationErrors(err)
	}
	if err, ok := r.Message.(error); ok {
		r.Message = err.Error()
	}

	return ctx.JSON(r.Code, r)
}

type ValidationError struct {
	Field   string // Field name that failed validation
	Message string // Error message describing the validation failure
}

type ValidationErrors []ValidationError

func extractValidationErrors(err validator.ValidationErrors) ValidationErrors {
	validationErrors := ValidationErrors{}
	for _, fieldError := range err {
		validationErrors = append(validationErrors, ValidationError{
			Field:   fieldError.StructField(),
			Message: fieldError.Tag(),
		})
	}
	return validationErrors
}
