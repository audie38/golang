package repository

import (
	"context"
	"fmt"
	"golang_db"
	"golang_db/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T){
	commentRepo := NewCommentRepository(golang_db.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email: "golangdao@localhost.com",
		CommentDesc: "Golang Repository Pattern",
	}

	result, err := commentRepo.Insert(ctx, comment)
	if err != nil{
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentSelectById(t *testing.T){
	commentRepo := NewCommentRepository(golang_db.GetConnection())
	ctx := context.Background()
	result, err := commentRepo.FindById(ctx, 31)
	if err != nil{
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentSelectAll(t *testing.T){
	commentRepo := NewCommentRepository(golang_db.GetConnection())
	ctx := context.Background()
	result, err := commentRepo.FindAll(ctx)
	if err != nil{
		panic(err)
	}

	for i := 0; i < len(result); i++{
		fmt.Println(result[i])
	}
}

func TestCommentUpdate(t *testing.T){
	commentRepo := NewCommentRepository(golang_db.GetConnection())
	ctx := context.Background()
	updateExistingComment := entity.Comment{
		CommentId: 31,
		Email: "updategolangdao@localhost.com",
		CommentDesc: "Update lagi Golang Repository Pattern",
	}
	result, err := commentRepo.Update(ctx, updateExistingComment)
	if err != nil{
		t.Fatal(err)
	}

	fmt.Println(result)
}

func TestCommentDelete(t *testing.T){
	commentRepo := NewCommentRepository(golang_db.GetConnection())
	ctx := context.Background()
	result, err := commentRepo.Delete(ctx, 31)
	if err != nil{
		t.Fatal(err)
	}

	fmt.Print(result)
}