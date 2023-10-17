package response

import (
	"errors"
	appError "github.com/BetterToPractice/go-echo-setup/errors"
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

type BadRequest struct {
	Req     interface{}
	Message interface{} `json:"message"`
}

type NotFound struct {
	Message interface{} `json:"message"`
}

type PolicyResponse struct {
	Message error `json:"message"`
}

func (r PolicyResponse) JSON(ctx echo.Context) error {
	resp := Response{
		Code:    http.StatusUnauthorized,
		Message: r.Message,
	}
	if errors.Is(r.Message, appError.Forbidden) {
		resp.Code = http.StatusForbidden
	}
	return resp.JSON(ctx)
}

func (r BadRequest) JSON(ctx echo.Context) error {
	resp := Response{Code: http.StatusBadRequest, Message: r.Message}

	if err, ok := r.Message.(validator.ValidationErrors); ok && err != nil {
		var validationErrors []ValidationError
		v := reflect.TypeOf(r.Req)

		for _, e := range err {
			field, _ := v.FieldByName(e.Field())
			validationErrors = append(validationErrors, ValidationError{
				Field:   field.Tag.Get("json"),
				Message: e.Tag(),
			})
		}
		resp.Data = validationErrors
		resp.Message = http.StatusText(resp.Code)
		return resp.JSON(ctx)
	}

	return resp.JSON(ctx)
}

func (r NotFound) JSON(ctx echo.Context) error {
	resp := Response{Code: http.StatusNotFound, Message: r.Message}
	return resp.JSON(ctx)
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
		if errors.Is(err, appError.DatabaseInternalError) {
			r.Code = http.StatusInternalServerError
		}
		r.Message = err.Error()
	}

	return ctx.JSON(r.Code, r)
}
