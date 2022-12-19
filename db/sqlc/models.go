// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"database/sql"
	"time"
)

type Food struct {
	ID     int64  `json:"id"`
	MenuID int64  `json:"menu_id"`
	Name   string `json:"name"`
	Price  int32  `json:"price"`
}

type FoodTag struct {
	ID      int64  `json:"id"`
	FoodID  int64  `json:"food_id"`
	FoodTag string `json:"food_tag"`
}

type Menu struct {
	ID        int64        `json:"id"`
	StoreID   int64        `json:"store_id"`
	MenuName  string       `json:"menu_name"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type Store struct {
	ID           int64     `json:"id"`
	StoreName    string    `json:"store_name"`
	StoreAddress string    `json:"store_address"`
	StorePhone   string    `json:"store_phone"`
	StoreOwner   string    `json:"store_owner"`
	StoreManager string    `json:"store_manager"`
	CreatedAt    time.Time `json:"created_at"`
}

type Table struct {
	ID        int64     `json:"id"`
	StoreID   int64     `json:"store_id"`
	TableID   int64     `json:"table_id"`
	TableName string    `json:"table_name"`
	CreatedAt time.Time `json:"created_at"`
}
