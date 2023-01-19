// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: stores.sql

package db

import (
	"context"
)

const createStore = `-- name: CreateStore :one
INSERT INTO stores (
    account_id,
    store_name,
    store_address,
    store_phone,
    store_manager
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, account_id, store_name, store_address, store_phone, store_manager, created_at
`

type CreateStoreParams struct {
	AccountID    int64  `json:"account_id"`
	StoreName    string `json:"store_name"`
	StoreAddress string `json:"store_address"`
	StorePhone   string `json:"store_phone"`
	StoreManager string `json:"store_manager"`
}

func (q *Queries) CreateStore(ctx context.Context, arg CreateStoreParams) (Store, error) {
	row := q.db.QueryRowContext(ctx, createStore,
		arg.AccountID,
		arg.StoreName,
		arg.StoreAddress,
		arg.StorePhone,
		arg.StoreManager,
	)
	var i Store
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.StoreName,
		&i.StoreAddress,
		&i.StorePhone,
		&i.StoreManager,
		&i.CreatedAt,
	)
	return i, err
}

const deleteStore = `-- name: DeleteStore :exec
DELETE FROM stores
WHERE id = $1
`

func (q *Queries) DeleteStore(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteStore, id)
	return err
}

const getStore = `-- name: GetStore :one
SELECT id, account_id, store_name, store_address, store_phone, store_manager, created_at FROM stores
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetStore(ctx context.Context, id int64) (Store, error) {
	row := q.db.QueryRowContext(ctx, getStore, id)
	var i Store
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.StoreName,
		&i.StoreAddress,
		&i.StorePhone,
		&i.StoreManager,
		&i.CreatedAt,
	)
	return i, err
}

const listStoresByName = `-- name: ListStoresByName :many
SELECT id, account_id, store_name, store_address, store_phone, store_manager, created_at FROM stores
WHERE store_name ~* $1
LIMIT $2
OFFSET $3
`

type ListStoresByNameParams struct {
	StoreName string `json:"store_name"`
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
}

func (q *Queries) ListStoresByName(ctx context.Context, arg ListStoresByNameParams) ([]Store, error) {
	rows, err := q.db.QueryContext(ctx, listStoresByName, arg.StoreName, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Store{}
	for rows.Next() {
		var i Store
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.StoreName,
			&i.StoreAddress,
			&i.StorePhone,
			&i.StoreManager,
			&i.CreatedAt,
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

const updateStore = `-- name: UpdateStore :one
UPDATE stores
SET account_id = $2, store_name = $3, store_address = $4, store_phone = $5, store_manager = $6
WHERE id = $1
RETURNING id, account_id, store_name, store_address, store_phone, store_manager, created_at
`

type UpdateStoreParams struct {
	ID           int64  `json:"id"`
	AccountID    int64  `json:"account_id"`
	StoreName    string `json:"store_name"`
	StoreAddress string `json:"store_address"`
	StorePhone   string `json:"store_phone"`
	StoreManager string `json:"store_manager"`
}

func (q *Queries) UpdateStore(ctx context.Context, arg UpdateStoreParams) (Store, error) {
	row := q.db.QueryRowContext(ctx, updateStore,
		arg.ID,
		arg.AccountID,
		arg.StoreName,
		arg.StoreAddress,
		arg.StorePhone,
		arg.StoreManager,
	)
	var i Store
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.StoreName,
		&i.StoreAddress,
		&i.StorePhone,
		&i.StoreManager,
		&i.CreatedAt,
	)
	return i, err
}
