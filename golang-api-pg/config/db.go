package config

import (
	"database/sql"
	"golang_api_pg/helper"
	"time"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB{
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost/golang_db?sslmode=disable")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}