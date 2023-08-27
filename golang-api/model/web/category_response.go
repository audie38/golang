package web

type CategoryResponse struct {
	CategoryId int64  `json:"id"`
	Name       string `json:"name"`
}