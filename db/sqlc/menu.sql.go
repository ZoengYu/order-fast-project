// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: menu.sql

package db

import (
	"context"
	"time"

	"github.com/lib/pq"
)

const addFoodTag = `-- name: AddFoodTag :one
INSERT INTO food_tag (
	menu_food_id,
	food_tag
) VALUES (
	$1, $2
) RETURNING id, menu_food_id, food_tag
`

type AddFoodTagParams struct {
	MenuFoodID int64  `json:"menu_food_id"`
	FoodTag    string `json:"food_tag"`
}

func (q *Queries) AddFoodTag(ctx context.Context, arg AddFoodTagParams) (FoodTag, error) {
	row := q.db.QueryRowContext(ctx, addFoodTag, arg.MenuFoodID, arg.FoodTag)
	var i FoodTag
	err := row.Scan(&i.ID, &i.MenuFoodID, &i.FoodTag)
	return i, err
}

const addMenuFood = `-- name: AddMenuFood :one
INSERT INTO menu_food (
	menu_id,
	food_name,
	custom_option
) VALUES (
	$1, $2, $3
) RETURNING id, menu_id, food_name, custom_option
`

type AddMenuFoodParams struct {
	MenuID       int64    `json:"menu_id"`
	FoodName     string   `json:"food_name"`
	CustomOption []string `json:"custom_option"`
}

func (q *Queries) AddMenuFood(ctx context.Context, arg AddMenuFoodParams) (MenuFood, error) {
	row := q.db.QueryRowContext(ctx, addMenuFood, arg.MenuID, arg.FoodName, pq.Array(arg.CustomOption))
	var i MenuFood
	err := row.Scan(
		&i.ID,
		&i.MenuID,
		&i.FoodName,
		pq.Array(&i.CustomOption),
	)
	return i, err
}

const createStoreMenu = `-- name: CreateStoreMenu :one
INSERT INTO menu (
    store_id,
	menu_name,
	created_at
) VALUES (
    $1, $2, $3
) RETURNING id, store_id, menu_name, created_at, updated_at
`

type CreateStoreMenuParams struct {
	StoreID   int64     `json:"store_id"`
	MenuName  string    `json:"menu_name"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) CreateStoreMenu(ctx context.Context, arg CreateStoreMenuParams) (Menu, error) {
	row := q.db.QueryRowContext(ctx, createStoreMenu, arg.StoreID, arg.MenuName, arg.CreatedAt)
	var i Menu
	err := row.Scan(
		&i.ID,
		&i.StoreID,
		&i.MenuName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getFoodTag = `-- name: GetFoodTag :one
SELECT id, menu_food_id, food_tag FROM food_tag
WHERE menu_food_id = $1 AND food_tag = $2
`

type GetFoodTagParams struct {
	MenuFoodID int64  `json:"menu_food_id"`
	FoodTag    string `json:"food_tag"`
}

func (q *Queries) GetFoodTag(ctx context.Context, arg GetFoodTagParams) (FoodTag, error) {
	row := q.db.QueryRowContext(ctx, getFoodTag, arg.MenuFoodID, arg.FoodTag)
	var i FoodTag
	err := row.Scan(&i.ID, &i.MenuFoodID, &i.FoodTag)
	return i, err
}

const getMenuFood = `-- name: GetMenuFood :one
SELECT id, menu_id, food_name, custom_option FROM menu_food
WHERE menu_id = $1 AND food_name = $2
`

type GetMenuFoodParams struct {
	MenuID   int64  `json:"menu_id"`
	FoodName string `json:"food_name"`
}

func (q *Queries) GetMenuFood(ctx context.Context, arg GetMenuFoodParams) (MenuFood, error) {
	row := q.db.QueryRowContext(ctx, getMenuFood, arg.MenuID, arg.FoodName)
	var i MenuFood
	err := row.Scan(
		&i.ID,
		&i.MenuID,
		&i.FoodName,
		pq.Array(&i.CustomOption),
	)
	return i, err
}

const getStoreMenu = `-- name: GetStoreMenu :one
SELECT id, store_id, menu_name, created_at, updated_at FROM menu
WHERE store_id = $1 AND menu_name = $2
`

type GetStoreMenuParams struct {
	StoreID  int64  `json:"store_id"`
	MenuName string `json:"menu_name"`
}

func (q *Queries) GetStoreMenu(ctx context.Context, arg GetStoreMenuParams) (Menu, error) {
	row := q.db.QueryRowContext(ctx, getStoreMenu, arg.StoreID, arg.MenuName)
	var i Menu
	err := row.Scan(
		&i.ID,
		&i.StoreID,
		&i.MenuName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listMenuFoodTag = `-- name: ListMenuFoodTag :many
SELECT food_tag.food_tag FROM food_tag
WHERE menu_food_id = $1
`

func (q *Queries) ListMenuFoodTag(ctx context.Context, menuFoodID int64) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listMenuFoodTag, menuFoodID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var food_tag string
		if err := rows.Scan(&food_tag); err != nil {
			return nil, err
		}
		items = append(items, food_tag)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeFoodTag = `-- name: RemoveFoodTag :exec
DELETE FROM food_tag
WHERE menu_food_id = $1 AND food_tag = $2
`

type RemoveFoodTagParams struct {
	MenuFoodID int64  `json:"menu_food_id"`
	FoodTag    string `json:"food_tag"`
}

func (q *Queries) RemoveFoodTag(ctx context.Context, arg RemoveFoodTagParams) error {
	_, err := q.db.ExecContext(ctx, removeFoodTag, arg.MenuFoodID, arg.FoodTag)
	return err
}

const updateStoreMenu = `-- name: UpdateStoreMenu :one
UPDATE menu
SET menu_name = $2
WHERE id = $1
RETURNING id, store_id, menu_name, created_at, updated_at
`

type UpdateStoreMenuParams struct {
	ID       int64  `json:"id"`
	MenuName string `json:"menu_name"`
}

func (q *Queries) UpdateStoreMenu(ctx context.Context, arg UpdateStoreMenuParams) (Menu, error) {
	row := q.db.QueryRowContext(ctx, updateStoreMenu, arg.ID, arg.MenuName)
	var i Menu
	err := row.Scan(
		&i.ID,
		&i.StoreID,
		&i.MenuName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
