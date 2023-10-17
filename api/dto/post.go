package dto

import (
	"github.com/BetterToPractice/go-echo-setup/models"
)

type PostQueryParam struct {
	PaginationParam
}

type PostRequest struct {
	Title string `json:"title" validate:"required"`
	Body  string `json:"body" validate:"required"`
}

type PostUpdateRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type PostUserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type PostResponse struct {
	ID    uint             `json:"id"`
	Title string           `json:"title"`
	Body  string           `json:"body"`
	User  PostUserResponse `json:"user"`
}

func (p *PostResponse) Serializer(post *models.Post) {
	p.ID = post.ID
	p.Title = post.Title
	p.Body = post.Body
	p.User = PostUserResponse{
		ID:       post.UserID,
		Username: post.User.Username,
	}
}

type PostPaginationResponse struct {
	List       []PostResponse `json:"list"`
	Pagination *Pagination    `json:"pagination"`
}

func (p *PostPaginationResponse) Serializer(posts *models.Posts) {
	var list []PostResponse
	for _, post := range *posts {
		p := PostResponse{}
		p.Serializer(&post)
		list = append(list, p)
	}
	p.List = list
}
