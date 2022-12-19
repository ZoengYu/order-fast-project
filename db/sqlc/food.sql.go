// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: food.sql

package db

import (
	"context"
)

const createMenuFood = `-- name: CreateMenuFood :one
INSERT INTO food (
	menu_id,
	name,
	price
) VALUES (
	$1, $2, $3
) RETURNING id, menu_id, name, price
`

type CreateMenuFoodParams struct {
	MenuID int64  `json:"menu_id"`
	Name   string `json:"name"`
	Price  int32  `json:"price"`
}

func (q *Queries) CreateMenuFood(ctx context.Context, arg CreateMenuFoodParams) (Food, error) {
	row := q.db.QueryRowContext(ctx, createMenuFood, arg.MenuID, arg.Name, arg.Price)
	var i Food
	err := row.Scan(
		&i.ID,
		&i.MenuID,
		&i.Name,
		&i.Price,
	)
	return i, err
}

const getMenuFood = `-- name: GetMenuFood :one
SELECT id, menu_id, name, price FROM food
WHERE id = $1
`

func (q *Queries) GetMenuFood(ctx context.Context, id int64) (Food, error) {
	row := q.db.QueryRowContext(ctx, getMenuFood, id)
	var i Food
	err := row.Scan(
		&i.ID,
		&i.MenuID,
		&i.Name,
		&i.Price,
	)
	return i, err
}

const listMenuFood = `-- name: ListMenuFood :many
SELECT id, menu_id, name, price FROM food
WHERE menu_id = $1
`

func (q *Queries) ListMenuFood(ctx context.Context, menuID int64) ([]Food, error) {
	rows, err := q.db.QueryContext(ctx, listMenuFood, menuID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Food{}
	for rows.Next() {
		var i Food
		if err := rows.Scan(
			&i.ID,
			&i.MenuID,
			&i.Name,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
