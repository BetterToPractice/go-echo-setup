package controllers

import (
	"github.com/BetterToPractice/go-echo-setup/api/services"
	"github.com/BetterToPractice/go-echo-setup/models"
	"github.com/BetterToPractice/go-echo-setup/pkg/response"
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
		return response.Response{Code: http.StatusBadRequest, Message: err}.JSON(ctx)
	}

	qr, err := c.postService.Query(params)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Message: err}.JSON(ctx)
	}

	return response.Response{Code: http.StatusOK, Data: qr}.JSON(ctx)
}

// Detail godoc
//
//	@Summary		Detail a post
//	@Description	get detail a post
//	@Param 			id path int true "post id"
//	@Tags			post
//	@Accept			json
//	@Produce		json
//	@Router			/posts/{id}/ [get]
func (c PostController) Detail(ctx echo.Context) error {
	post, err := c.postService.Get(ctx.Param("id"))
	if err != nil {
		return response.Response{Code: http.StatusNotFound, Message: err}.JSON(ctx)
	}
	return response.Response{Code: http.StatusOK, Data: post}.JSON(ctx)
}

// Destroy godoc
//
// @summary			Delete a post
// @Description		perform delete a post
// @Param 			id path int true "post id"
// @Tags			post
// @Accept			json
// @Product			json
// @Router			/posts/{id}/ [delete]
func (c PostController) Destroy(ctx echo.Context) error {
	err := c.postService.Delete(ctx.Param("id"))
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Message: err}.JSON(ctx)
	}

	return response.Response{Code: http.StatusOK}.JSON(ctx)
}
