package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_db/entity"
	"strconv"
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
		return comment, err
	}
	
	sqlQuery := "INSERT INTO `COMMENT`(EMAIL, COMMENT_DESC) VALUES(?, ?)"
	result, err := tx.ExecContext(ctx, sqlQuery, comment.Email, comment.CommentDesc)
	if err != nil{
		tx.Rollback()
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil{
		tx.Rollback()
		return comment, err
	}
	comment.CommentId = int64(id)
	tx.Commit()
	return comment, nil
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error){
	sqlQuery := "SELECT COMMENT_ID, EMAIL, COMMENT_DESC FROM `COMMENT`"
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
	sqlQuery := "SELECT COMMENT_ID, EMAIL, COMMENT_DESC FROM `COMMENT` WHERE COMMENT_ID = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, sqlQuery, id)
	comment := entity.Comment{}
	if err != nil{
		return comment, err
	}
	defer rows.Close()
	if rows.Next(){
		rows.Scan(&comment.CommentId, &comment.Email, &comment.CommentDesc)
		return comment, nil
	}else{
		return comment, errors.New("Comment with Id: " + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repository *commentRepositoryImpl) Update(ctx context.Context, comment entity.Comment) (entity.Comment, error){
	tx, err := repository.DB.Begin()
	if err != nil{
		return comment, err
	}
	
	sqlQuery := "UPDATE `COMMENT`SET EMAIL = ? , COMMENT_DESC = ? WHERE COMMENT_ID = ?"
	result, err := tx.ExecContext(ctx, sqlQuery, comment.Email, comment.CommentDesc, comment.CommentId)
	if err != nil{
		tx.Rollback()
		return entity.Comment{}, err
	}
	
	existingComment, err := result.RowsAffected()
	if err != nil{
		tx.Rollback()
		return entity.Comment{}, err
	}

	if existingComment > 0{
		tx.Commit()
		return comment, nil
	}else{
		return entity.Comment{}, errors.New("Comment with id "+strconv.Itoa(int(comment.CommentId))+" Not Found")
	}	
}

func (repository *commentRepositoryImpl) Delete(ctx context.Context, id int64) (entity.Comment, error){
	tx, err := repository.DB.Begin()
	if err != nil{
		return entity.Comment{}, err
	}
	
	sqlQuery := "DELETE FROM `COMMENT` WHERE COMMENT_ID = ?"
	result, err := tx.ExecContext(ctx, sqlQuery, id)
	if err != nil{
		tx.Rollback()
		return entity.Comment{}, err
	}
	
	existingComment, err := result.RowsAffected()
	if err != nil{
		tx.Rollback()
		return entity.Comment{}, err
	}

	if existingComment > 0{
		tx.Commit()
		return entity.Comment{}, nil
	}else{
		return entity.Comment{}, errors.New("Invalid Comment")
	}

}