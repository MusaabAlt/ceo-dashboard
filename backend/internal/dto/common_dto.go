package dto

type PaginationRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type PaginationResponse struct {
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
	Data       interface{} `json:"data"`
}
