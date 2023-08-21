package golang_db_pg

import (
	"context"
	"fmt"
	"golang_db_pg/config"
	"golang_db_pg/entity"
	"golang_db_pg/repository"
	"testing"
)

func TestCommentInsert(t *testing.T){
	commentRepo := repository.NewCommentRepository(config.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email: "golangpostgre@localhost.com",
		CommentDesc: "Test Golang Repository Pattern with PostgreSQL",
	}

	result, err := commentRepo.Insert(ctx, comment)
	if err != nil{
		t.Fatal(err)
	}

	fmt.Println(result)
}

func TestCommentSelectAll(t *testing.T){
	commentRepo := repository.NewCommentRepository(config.GetConnection())
	ctx := context.Background()
	result, err := commentRepo.FindAll(ctx)
	if err != nil{
		t.Fatal(err)
	}
	
	for i := 0; i < len(result); i++{
		fmt.Println(result[i])
	}
}

func TestCommentSelectById(t *testing.T){
	commentRepo := repository.NewCommentRepository(config.GetConnection())
	ctx := context.Background()
	result, err := commentRepo.FindById(ctx, 2)
	if err != nil{
		t.Fatal(err)
	}
	
	fmt.Println(result)
}

func TestCommentUpdate(t *testing.T){
	commentRepo := repository.NewCommentRepository(config.GetConnection())
	ctx := context.Background()

	newComment := entity.Comment{
		CommentId: 3,
		Email: "updategolangpostgre@localhost.com",
		CommentDesc: "Test Golang Postgre Update stmnt",
	}

	result, err := commentRepo.Update(ctx, newComment)
	if err != nil{
		t.Fatal(err)
	}

	fmt.Println("Updated Comment: ", result)
}

func TestCommentDelete(t *testing.T){
	commentRepo := repository.NewCommentRepository(config.GetConnection())
	ctx := context.Background()
	err := commentRepo.Delete(ctx, 2)
	if err != nil{
		t.Fatal(err)
	}

	fmt.Println("Comment Deleted Successfully")
}