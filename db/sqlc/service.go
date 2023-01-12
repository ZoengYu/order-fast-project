package db

import "database/sql"

type DBService interface {
	Querier
}

// DBQuery connect to the real postgresql to execute SQL queries and transaction
type DBQuery struct {
	*Queries
	db *sql.DB
}

func NewDBService(db *sql.DB) DBService {
	return &DBQuery{
		db:      db,
		Queries: New(db),
	}
}
