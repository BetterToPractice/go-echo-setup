package controllers

import (
	"github.com/BetterToPractice/go-echo-setup/api/dto"
	"github.com/BetterToPractice/go-echo-setup/api/policies"
	"github.com/BetterToPractice/go-echo-setup/api/services"
	"github.com/BetterToPractice/go-echo-setup/pkg/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PostController struct {
	postService services.PostService
	postPolicy  policies.PostPolicy
	authService services.AuthService
}

func NewPostController(postService services.PostService, postPolicy policies.PostPolicy, authService services.AuthService) PostController {
	return PostController{
		postService: postService,
		postPolicy:  postPolicy,
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
//	@Success		200  {object}  response.Response{data=dto.PostPaginationResponse}  "ok"
//	@Failure		400  {object}  response.Response{data=[]response.ValidationErrors}  "bad request"
func (c PostController) List(ctx echo.Context) error {
	params := new(dto.PostQueryParam)
	if err := ctx.Bind(params); err != nil {
		return response.BadRequest{Message: err, Req: dto.PostQueryParam{}}.JSON(ctx)
	}

	qr, err := c.postService.Query(params)
	if err != nil {
		return response.BadRequest{Message: err}.JSON(ctx)
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
//	@Success		200  {object}  response.Response{data=dto.PostResponse}  "ok"
//	@Failure		404  {object}  response.Response{}  "not found"
func (c PostController) Detail(ctx echo.Context) error {
	post, resp, err := c.postService.Get(ctx.Param("id"))
	if err != nil {
		return response.NotFound{Message: err}.JSON(ctx)
	}

	user, _ := c.authService.Authenticate(ctx)
	if can, err := c.postPolicy.CanViewDetail(user, post); !can {
		return response.PolicyResponse{Message: err}.JSON(ctx)
	}

	return response.Response{Data: resp}.JSON(ctx)
}

// Create godoc
//
//	@Summary		Create a post
//	@Description	Create a post
//	@Tags			post
//	@Accept			application/json
//	@Produce		application/json
//	@Security 		BearerAuth
//	@Param 			data body dto.PostRequest true "Post"
//	@Router			/posts [post]
//	@Success		201  {object}  response.Response{data=dto.PostResponse}  "created"
//	@Failure		400  {object}  response.Response{data=[]response.ValidationErrors}  "bad request"
//	@Failure		404  {object}  response.Response  "not found"
//	@Failure		401  {object}  response.Response  "unauthorized"
//	@Failure		403  {object}  response.Response  "forbidden"
func (c PostController) Create(ctx echo.Context) error {
	user, _ := c.authService.Authenticate(ctx)
	if can, err := c.postPolicy.CanCreate(user); !can {
		return response.PolicyResponse{Message: err}.JSON(ctx)
	}

	params := new(dto.PostRequest)
	if err := ctx.Bind(params); err != nil {
		return response.BadRequest{Req: dto.PostRequest{}, Message: err}.JSON(ctx)
	}

	postResponse, err := c.postService.Create(params, user)
	if err != nil {
		return response.BadRequest{Message: err}.JSON(ctx)
	}

	return response.Response{Code: http.StatusCreated, Data: postResponse}.JSON(ctx)
}

// Update godoc
//
//	@Summary		Update a post
//	@Description	Update a post
//	@Tags			post
//	@Accept			application/json
//	@Produce		application/json
//	@Security 		BearerAuth
//	@Param 			data body dto.PostRequest true "Post"
//	@Router			/posts/{id} [patch]
//	@Success		200  {object}  response.Response  "ok"
//	@Failure		400  {object}  response.Response{data=[]response.ValidationErrors}  "bad request"
//	@Failure		404  {object}  response.Response  "not found"
//	@Failure		401  {object}  response.Response  "unauthorized"
//	@Failure		403  {object}  response.Response  "forbidden"
func (c PostController) Update(ctx echo.Context) error {
	post, _, err := c.postService.Get(ctx.Param("id"))
	if err != nil {
		return response.NotFound{Message: err}.JSON(ctx)
	}

	user, _ := c.authService.Authenticate(ctx)
	if can, err := c.postPolicy.CanUpdate(user, post); !can {
		return response.PolicyResponse{Message: err}.JSON(ctx)
	}

	params := new(dto.PostUpdateRequest)
	if err := ctx.Bind(params); err != nil {
		return response.BadRequest{Req: dto.PostUpdateRequest{}, Message: err}.JSON(ctx)
	}

	resp, err := c.postService.Update(post, params)
	if err != nil {
		return response.BadRequest{Message: err}.JSON(ctx)
	}

	return response.Response{Data: resp}.JSON(ctx)
}

// Destroy godoc
//
//	@summary		Delete a post
//	@Description	perform delete a post
//	@Param 			id path string true "post id"
//	@Tags			post
//	@Accept			application/json
//	@Produce		application/json
//	@Security 		BearerAuth
//	@Router			/posts/{id} [delete]
//	@Success		204  {object}  nil  "no content"
//	@Failure		404  {object}  response.Response  "not found"
//	@Failure		401  {object}  response.Response  "unauthorized"
//	@Failure		403  {object}  response.Response  "forbidden"
func (c PostController) Destroy(ctx echo.Context) error {
	post, _, err := c.postService.Get(ctx.Param("id"))
	if err != nil {
		return response.NotFound{Message: err}.JSON(ctx)
	}

	user, _ := c.authService.Authenticate(ctx)
	if can, err := c.postPolicy.CanUpdate(user, post); !can {
		return response.PolicyResponse{Message: err}.JSON(ctx)
	}

	if err := c.postService.Delete(post); err != nil {
		return response.BadRequest{Message: err}.JSON(ctx)
	}

	return response.Response{Code: http.StatusNoContent}.JSON(ctx)
}
