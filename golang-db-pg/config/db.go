package config

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	dbDriver := "postgres"
	connStr := "postgres://postgres:root@localhost/golang_db?sslmode=disable"

	db, err := sql.Open(dbDriver, connStr)
	if err != nil{
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}