-- name: CreateStore :one
INSERT INTO stores (
    account_id,
    store_name,
    store_address,
    store_phone,
    store_manager
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetStore :one
SELECT * FROM stores
WHERE id = $1 LIMIT 1;

-- name: ListStoresByName :many
SELECT * FROM stores
WHERE store_name ~* $1
LIMIT $2
OFFSET $3;

-- name: UpdateStore :one
UPDATE stores
SET account_id = $2, store_name = $3, store_address = $4, store_phone = $5, store_manager = $6
WHERE id = $1
RETURNING *;

-- name: DeleteStore :exec
DELETE FROM stores
WHERE id = $1;
