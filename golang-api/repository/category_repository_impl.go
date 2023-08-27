package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_api/helper"
	"golang_api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository{
	return &CategoryRepositoryImpl{}
}

func (c *CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "INSERT INTO CATEGORY(NAME) VALUES(?)"
	result, err := tx.ExecContext(ctx, query, category.Name)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.CategoryId = int64(id)
	return category
}

func (c *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "UPDATE CATEGORY(NAME) SET NAME = ? WHERE CATEGORY_ID = ?"
	_, err := tx.ExecContext(ctx, query, category.Name, category.CategoryId)
	helper.PanicIfError(err)

	return category
}

func (c *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	query := "DELETE FROM CATEGORY WHERE CATEGORY_ID = ?"
	_, err := tx.ExecContext(ctx, query, category.CategoryId)
	helper.PanicIfError(err)
}

func (c *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int64) (domain.Category, error) {
	query := "SELECT CATEGORY_ID, NAME FROM CATEGORY WHERE CATEGORY_ID = ?"
	rows, err := tx.QueryContext(ctx, query, categoryId)
	helper.PanicIfError(err)

	category := domain.Category{}

	if rows.Next(){
		err := rows.Scan(&category.CategoryId, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	}else{
		return category, errors.New("Category Not Found")
	}
}

func (c *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	query := "SELECT CATEGORY_ID, NAME FROM CATEGORY"

	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)

	var categories []domain.Category
	for rows.Next(){
		category := domain.Category{}
		err := rows.Scan(&category.CategoryId, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
