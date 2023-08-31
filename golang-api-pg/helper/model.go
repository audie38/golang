package helper

import (
	"golang_api_pg/model/domain"
	"golang_api_pg/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse{
	return web.CategoryResponse{
		CategoryId: category.CategoryId,
		Name: category.Name,
	}
}

func ToListCategoryResponse(categories []domain.Category) []web.CategoryResponse{
	var categoryResponses []web.CategoryResponse
	for i:= 0; i < len(categories); i++{
		categoryResponses = append(categoryResponses, ToCategoryResponse(categories[i]))
	}
	return categoryResponses
}