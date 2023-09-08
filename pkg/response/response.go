package response

import (
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
	if err, ok := r.Message.(error); ok {
		r.Message = err.Error()
	}

	return ctx.JSON(r.Code, r)
}
