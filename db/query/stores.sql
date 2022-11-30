-- name: CreateStore :one
INSERT INTO stores (
    store_name,
    store_address,
    store_phone,
    store_owner,
    store_manager
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetStore :one
SELECT * FROM stores
WHERE id = $1 LIMIT 1;

-- name: GetStoreByName :one
SELECT * FROM stores
WHERE store_name = $1 LIMIT 1;

-- name: UpdateStore :one
UPDATE stores
SET store_name = $2, store_address = $3, store_phone = $4, store_owner = $5, store_manager = $6
WHERE id = $1
RETURNING *;

-- name: DeleteStore :exec
DELETE FROM stores
WHERE id = $1;
