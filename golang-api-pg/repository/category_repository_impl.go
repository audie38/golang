package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_api_pg/helper"
	"golang_api_pg/model/domain"
)

type CategoryRepositoryImpl struct{}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repo *CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := `INSERT INTO CATEGORY("NAME") VALUES ($1) RETURNING CATEGORY_ID`
	insertedId := 0
	err := tx.QueryRowContext(ctx, query, category.Name).Scan(&insertedId)
	helper.PanicIfError(err)

	category.CategoryId = int64(insertedId)
	return category
}

func (repo *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	query := `SELECT CATEGORY_ID, "NAME" FROM CATEGORY`
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)

	var categories []domain.Category
	defer rows.Close()
	for rows.Next(){
		category := domain.Category{}
		err := rows.Scan(&category.CategoryId, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}

func (repo *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int64) (domain.Category, error) {
	query := `SELECT CATEGORY_ID, "NAME" FROM CATEGORY WHERE CATEGORY_ID = $1`
	rows, err := tx.QueryContext(ctx, query, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next(){
		err := rows.Scan(&category.CategoryId, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	}

	return category, errors.New("Category Not Found")
}

func (repo *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := `UPDATE CATEGORY SET "NAME" = $1 WHERE CATEGORY_ID = $2`
	_, err := tx.ExecContext(ctx, query, category.Name, category.CategoryId)
	helper.PanicIfError(err)

	return category
}

func (repo *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, categoryId int64) {
	query := `DELETE FROM CATEGORY WHERE CATEGORY_ID = $1`
	_, err := tx.ExecContext(ctx, query, categoryId)
	helper.PanicIfError(err)
}
