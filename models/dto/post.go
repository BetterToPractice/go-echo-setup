package dto

type PostRequest struct {
	Title string `json:"title" validate:"required"`
	Body  string `json:"body" validate:"required"`
}

type PostUpdateRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type PostResponse struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
