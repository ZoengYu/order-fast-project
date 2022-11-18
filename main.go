package main

import (
	"context"
	"database/sql"
	"fmt"
)

const (
	HOST = "localhost"
	PORT = "5432"
	DATABASE = "postgres"
	USER = "runner"
	PASSWORD = "password"
	SSL = "disable"
)

func OpenDB(ctx context.Context) *sql.DB {
	driver := "postgres"
	dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        HOST, PORT, USER, PASSWORD, DATABASE, SSL)

	db, err := sql.Open(driver, dsn)
	if err != nil {
		panic("open database error")
	}
	return db
}
