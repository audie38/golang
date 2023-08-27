package repository

import (
	"context"
	"database/sql"
	"golang_api/model/domain"
)

type CategoryRepository interface {
	Create(ctx context.Context, tx sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx sql.Tx, categoryId int64)  domain.Category
	FindAll(ctx context.Context, tx sql.Tx) []domain.Category
}