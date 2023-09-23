package lib

import (
	"errors"
	"fmt"
	"github.com/BetterToPractice/go-echo-setup/pkg/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HttpHandler struct {
	Engine   *echo.Echo
	Validate *validator.Validate
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func NewHttpHandler(logger Logger) HttpHandler {
	engine := echo.New()
	engine.Binder = &BinderWithValidation{}

	httpHandler := HttpHandler{
		Engine: engine,
	}
	httpHandler.Engine.HTTPErrorHandler = func(err error, ctx echo.Context) {
		var (
			code    = http.StatusInternalServerError
			message interface{}
		)

		var he *echo.HTTPError
		ok := errors.As(err, &he)
		if ok {
			code = he.Code
			message = he.Message
			if he.Internal != nil {
				message = fmt.Errorf("%v - %v", message, he.Internal)
			}
		}

		if !ctx.Response().Committed {
			if ctx.Request().Method == http.MethodHead {
				err = ctx.NoContent(he.Code)
			} else {
				err = response.Response{
					Code:    code,
					Message: message,
				}.JSON(ctx)
			}

			if err != nil {
				logger.DesugarZap.Error(err.Error())
			}
		}
	}

	httpHandler.Engine.Validator = func() echo.Validator {
		v := validator.New()
		return &Validator{validator: v}
	}()

	return httpHandler
}

type BinderWithValidation struct{}

func (BinderWithValidation) Bind(i interface{}, ctx echo.Context) error {
	binder := &echo.DefaultBinder{}

	if err := binder.Bind(i, ctx); err != nil {
		return errors.New(err.(*echo.HTTPError).Message.(string))
	}

	if err := ctx.Validate(i); err != nil {
		return err
	}

	return nil
}
