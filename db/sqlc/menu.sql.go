// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: menu.sql

package db

import (
	"context"
)

const createStoreMenu = `-- name: CreateStoreMenu :one
INSERT INTO menu (
    store_id,
	menu_name
) VALUES (
    $1, $2
) RETURNING id, store_id, menu_name, created_at, updated_at
`

type CreateStoreMenuParams struct {
	StoreID  int64  `json:"store_id"`
	MenuName string `json:"menu_name"`
}

func (q *Queries) CreateStoreMenu(ctx context.Context, arg CreateStoreMenuParams) (Menu, error) {
	row := q.db.QueryRowContext(ctx, createStoreMenu, arg.StoreID, arg.MenuName)
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

const deleteMenu = `-- name: DeleteMenu :exec
DELETE FROM menu
WHERE id = $1
`

func (q *Queries) DeleteMenu(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteMenu, id)
	return err
}

const getMenu = `-- name: GetMenu :one
SELECT id, store_id, menu_name, created_at, updated_at FROM menu
WHERE id = $1
`

func (q *Queries) GetMenu(ctx context.Context, id int64) (Menu, error) {
	row := q.db.QueryRowContext(ctx, getMenu, id)
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

const getStoreMenu = `-- name: GetStoreMenu :one
SELECT id, store_id, menu_name, created_at, updated_at FROM menu
WHERE store_id = $1 AND id = $2
`

type GetStoreMenuParams struct {
	StoreID int64 `json:"store_id"`
	ID      int64 `json:"id"`
}

func (q *Queries) GetStoreMenu(ctx context.Context, arg GetStoreMenuParams) (Menu, error) {
	row := q.db.QueryRowContext(ctx, getStoreMenu, arg.StoreID, arg.ID)
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

const listStoreMenu = `-- name: ListStoreMenu :many
SELECT id, store_id, menu_name, created_at, updated_at FROM menu
WHERE store_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListStoreMenuParams struct {
	StoreID int64 `json:"store_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) ListStoreMenu(ctx context.Context, arg ListStoreMenuParams) ([]Menu, error) {
	rows, err := q.db.QueryContext(ctx, listStoreMenu, arg.StoreID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Menu{}
	for rows.Next() {
		var i Menu
		if err := rows.Scan(
			&i.ID,
			&i.StoreID,
			&i.MenuName,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateStoreMenu = `-- name: UpdateStoreMenu :one
UPDATE menu
SET menu_name = $3
WHERE store_id = $1 AND id = $2
RETURNING id, store_id, menu_name, created_at, updated_at
`

type UpdateStoreMenuParams struct {
	StoreID  int64  `json:"store_id"`
	ID       int64  `json:"id"`
	MenuName string `json:"menu_name"`
}

func (q *Queries) UpdateStoreMenu(ctx context.Context, arg UpdateStoreMenuParams) (Menu, error) {
	row := q.db.QueryRowContext(ctx, updateStoreMenu, arg.StoreID, arg.ID, arg.MenuName)
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
