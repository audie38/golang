package golang_db

import (
	"database/sql"
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