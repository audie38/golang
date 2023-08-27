package web

type CategoryUpdateRequest struct {
	CategoryId int64  `validate:"required"`
	Name       string `validate:"required, max=200, min=1"`
}