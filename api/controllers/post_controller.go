package controllers

import (
	"github.com/BetterToPractice/go-echo-setup/api/services"
	"github.com/BetterToPractice/go-echo-setup/models"
	"github.com/BetterToPractice/go-echo-setup/models/dto"
	"github.com/BetterToPractice/go-echo-setup/pkg/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PostController struct {
	postService services.PostService
	authService services.AuthService
}

func NewPostController(postService services.PostService, authService services.AuthService) PostController {
	return PostController{
		postService: postService,
		authService: authService,
	}
}

// List godoc
//
//	@Summary		List several posts
//	@Description	get list several posts
//	@Tags			post
//	@Accept			application/json
//	@Produce		application/json
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
//	@Param 			id path string true "post id"
//	@Tags			post
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/posts/{id} [get]
func (c PostController) Detail(ctx echo.Context) error {
	post, err := c.postService.Get(ctx.Param("id"))
	if err != nil {
		return response.Response{Code: http.StatusNotFound, Message: err}.JSON(ctx)
	}
	return response.Response{Code: http.StatusOK, Data: post}.JSON(ctx)
}

// Create godoc
//
//	@Summary		Create a post
//	@Description	create a post
//	@Tags			post
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/posts [post]
func (c PostController) Create(ctx echo.Context) error {
	user, err := c.authService.Authorize(ctx)
	if err != nil {
		return response.Response{Code: http.StatusUnauthorized, Message: err}.JSON(ctx)
	}

	params := new(dto.PostRequest)
	if err := ctx.Bind(params); err != nil {
		return response.Response{Code: http.StatusBadRequest, Message: err}.JSON(ctx)
	}

	postResponse, err := c.postService.Create(params, user)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Message: err}.JSON(ctx)
	}

	return response.Response{Code: http.StatusOK, Data: postResponse}.JSON(ctx)
}

// Destroy godoc
//
// @summary			Delete a post
// @Description		perform delete a post
// @Param 			id path string true "post id"
// @Tags			post
// @Accept			application/json
// @Product			application/json
// @Router			/posts/{id} [delete]
func (c PostController) Destroy(ctx echo.Context) error {
	err := c.postService.Delete(ctx.Param("id"))
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Message: err}.JSON(ctx)
	}

	return response.Response{Code: http.StatusNoContent}.JSON(ctx)
}
