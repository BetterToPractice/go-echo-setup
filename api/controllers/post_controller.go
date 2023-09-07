package controllers

import (
	"github.com/BetterToPractice/go-echo-setup/api/services"
	"github.com/BetterToPractice/go-echo-setup/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PostController struct {
	postService services.PostService
}

func NewPostController(postService services.PostService) PostController {
	return PostController{
		postService: postService,
	}
}

// List godoc
//
//	@Summary		List several posts
//	@Description	get list several posts
//	@Tags			post
//	@Accept			json
//	@Produce		json
//	@Router			/posts [get]
func (c PostController) List(ctx echo.Context) error {
	params := new(models.PostQueryParams)
	if err := ctx.Bind(params); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	qr, err := c.postService.Query(params)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, qr)
}

// Detail godoc
//
//	@Summary		Detail a post
//	@Description	get detail a post
//	@Tags			post
//	@Accept			json
//	@Produce		json
//	@Router			/posts/{id}/ [get]
func (c PostController) Detail(ctx echo.Context) error {
	post, err := c.postService.Get(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}
	return ctx.JSON(http.StatusOK, post)
}
