package db

import "database/sql"

type DBService struct {
	*Queries
	db *sql.DB
}

func NewDBService(db *sql.DB) *DBService{
	return &DBService{
		db: db,
		Queries: New(db),
	}
}
