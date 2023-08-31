package dto

type Pagination struct {
	Total    int64 `json:"total"`
	Current  int   `json:"current"`
	PageSize int   `json:"page_size"`
}

type PaginationParam struct {
	Current  int `json:"current"`
	PageSize int `json:"page_size"`
}
