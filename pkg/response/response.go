package response

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
)

type Response struct {
	Code    int         `json:"-"`
	Data    interface{} `json:"data,omitempty"`
	Message interface{} `json:"message"`
}

type ValidationError struct {
	Field   string `json:"field"`   // Field name that failed validation
	Message string `json:"message"` // Error message describing the validation failure
}

type ValidationErrors []ValidationError

func (r Response) JSONValidationError(dto interface{}, ctx echo.Context) error {
	if r.Code == 0 {
		r.Code = http.StatusBadRequest
	}

	if err, ok := r.Message.(validator.ValidationErrors); ok && err != nil {
		var validationErrors []ValidationError
		v := reflect.TypeOf(dto)

		for _, e := range err {
			field, _ := v.FieldByName(e.Field())
			validationErrors = append(validationErrors, ValidationError{
				Field:   field.Tag.Get("json"),
				Message: e.Tag(),
			})
		}

		r.Message = http.StatusText(r.Code)
		r.Data = validationErrors

		return ctx.JSON(r.Code, r)
	}

	return r.JSON(ctx)
}

func (r Response) JSON(ctx echo.Context) error {
	if r.Code == 0 {
		r.Code = http.StatusOK
	}
	if r.Message == "" || r.Message == nil {
		r.Message = http.StatusText(r.Code)
	}

	if err, ok := r.Message.(error); ok {
		r.Message = err.Error()
	}

	return ctx.JSON(r.Code, r)
}
