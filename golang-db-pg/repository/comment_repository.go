package repository

import (
	"context"
	"golang_db_pg/entity"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
	FindById(ctx context.Context, id int64) (entity.Comment, error)
	Update(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	Delete(ctx context.Context, id int64)(error)
}