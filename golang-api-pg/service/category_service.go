package service

import (
	"context"
	"golang_api_pg/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
	FindById(ctx context.Context, categoryId int64) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int64)
}