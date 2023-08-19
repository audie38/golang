package golang_db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

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

	sqlQuery := 
	"INSERT INTO CUSTOMER(CUSTOMER_NAME, EMAIL, BALANCE, RATING, BIRTH_DATE, MARRIED) VALUES('Ichigo', 'ichigo@localhost.com', 100000, 5.0, '2000-08-30', true)"
	_, err := db.ExecContext(ctx, sqlQuery)
	if err != nil{
		panic(err)
	}

	fmt.Println("Success Insert New Customer")
}

func TestSelectSql(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	sqlQuery := "SELECT CUSTOMER_ID, CUSTOMER_NAME, EMAIL, BALANCE, RATING, CREATED_AT, BIRTH_DATE, MARRIED FROM CUSTOMER"
	rows, err := db.QueryContext(ctx, sqlQuery)
	if err != nil{
		panic(err)
	}

	for rows.Next(){
		var(
			customerId int64
			customerName string
			email sql.NullString
			balance int32
			rating float64
			createdAt time.Time
			birthDate sql.NullTime
			married bool
		)
		err := rows.Scan(&customerId, &customerName, &email, &balance, &rating, &createdAt, &birthDate, &married)
		if err != nil{
			panic(err)
		}
		fmt.Println("======================")
		fmt.Println("CustomerId: ", customerId)
		fmt.Println("Customer Name: ", customerName)
		if email.Valid{
			fmt.Println("Email: ", email.String)
		}
		fmt.Println("Balance: ", balance)
		fmt.Println("Rating: ", rating)
		fmt.Println("Created At: ", createdAt)
		if birthDate.Valid{
			fmt.Println("Birth Date: ", birthDate.Time)
		}
		fmt.Println("Married: ", married)
	}
	fmt.Println("======================")

	defer rows.Close()
}