package golang_db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	dbDriver := "mysql"
	connStr := "root:@tcp(localhost:3306)/golang_db"

	db, err := sql.Open(dbDriver, connStr)
	if err != nil{
		panic(err)
	}
	defer db.Close()
}

func TestExecSql(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlQuery := "INSERT INTO CUSTOMER(CUSTOMER_NAME) VALUES('Kurosaki')"
	_, err := db.ExecContext(ctx, sqlQuery)
	if err != nil{
		panic(err)
	}

	fmt.Println("Success Insert New Customer")
}