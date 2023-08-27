package helper

import (
	"golang_api/model/domain"
	"golang_api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse{
	return web.CategoryResponse{
		CategoryId : category.CategoryId,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse{
	var categoryResponses []web.CategoryResponse
	for _, category := range categories{
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}