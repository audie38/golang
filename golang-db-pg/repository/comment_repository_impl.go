package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_db_pg/entity"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository{
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error){
	tx, err := repository.DB.Begin()
	if err != nil{
		return entity.Comment{}, err
	}

	sqlQuery := `INSERT INTO "COMMENT"(EMAIL, COMMENT_DESC) VALUES ($1, $2) RETURNING COMMENT_ID`
	insertedId := 0
	err = tx.QueryRow(sqlQuery, comment.Email, comment.CommentDesc).Scan(&insertedId)
	if err != nil{
		tx.Rollback()
		return entity.Comment{}, err
	}

	if insertedId != 0{
		comment.CommentId = int64(insertedId)
	}

	tx.Commit()
	return comment, nil
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error){
	sqlQuery := `SELECT COMMENT_ID, EMAIL, COMMENT_DESC FROM "COMMENT"`
	rows, err := repository.DB.QueryContext(ctx, sqlQuery)
	comments := []entity.Comment{}
	if err != nil{
		return comments, err
	}
	defer rows.Close()
	for rows.Next(){
		comment := entity.Comment{}
		rows.Scan(&comment.CommentId, &comment.Email, &comment.CommentDesc)
		comments = append(comments, comment)
	}
	return comments, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int64) (entity.Comment, error){
	sqlQuery := `SELECT COMMENT_ID, EMAIL, COMMENT_DESC FROM "COMMENT" WHERE COMMENT_ID = $1`
	rows, err := repository.DB.QueryContext(ctx, sqlQuery, id)
	if err != nil{
		return entity.Comment{}, err
	}
	defer rows.Close()
	comment := entity.Comment{}
	if rows.Next(){
		rows.Scan(&comment.CommentId, &comment.Email, &comment.CommentDesc)
	}

	if comment.CommentId == 0 {
		return entity.Comment{}, errors.New("Comment Not Found")
	}

	return comment, nil
}

func (repository *commentRepositoryImpl) Update(ctx context.Context, comment entity.Comment) (entity.Comment, error){
	tx, err := repository.DB.Begin()
	if err != nil{
		return entity.Comment{}, err
	}

	sqlQuery := `UPDATE "COMMENT" SET EMAIL = $1 , COMMENT_DESC = $2 WHERE COMMENT_ID = $3`
	result, err := tx.ExecContext(ctx, sqlQuery, comment.Email, comment.CommentDesc, comment.CommentId)
	if err != nil{
		tx.Rollback()
		return entity.Comment{}, err
	}
	numRowAffected, err := result.RowsAffected()
	if err != nil{
		tx.Rollback()
		return entity.Comment{}, err
	}

	if numRowAffected == 0{
		tx.Rollback()
		return entity.Comment{}, errors.New("Comment Not Found")
	}

	tx.Commit()
	return comment, nil
}

func (repository *commentRepositoryImpl) Delete(ctx context.Context, id int64)(error){
	tx, err := repository.DB.Begin()
	if err != nil{
		return err
	}

	sqlQuery := `DELETE FROM "COMMENT" WHERE COMMENT_ID = $1`
	result, err := tx.ExecContext(ctx, sqlQuery, id)
	if err != nil{
		tx.Rollback()
		return err
	}
	numRowsAffected, err := result.RowsAffected()
	if err != nil{
		return err
	}

	if numRowsAffected == 0{
		tx.Rollback()
		return errors.New("Invalid Comment, Comment Not Found")
	}

	tx.Commit()
	return nil
}
