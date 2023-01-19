-- name: CreateStore :one
INSERT INTO stores (
    owner,
    name,
    address,
    phone,
    manager
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetStore :one
SELECT * FROM stores
WHERE id = $1 LIMIT 1;

-- name: ListStoresByName :many
SELECT * FROM stores
WHERE name ~* $1
LIMIT $2
OFFSET $3;

-- name: UpdateStore :one
UPDATE stores
SET owner = $2, name = $3, address = $4, phone = $5, manager = $6
WHERE id = $1
RETURNING *;

-- name: DeleteStore :exec
DELETE FROM stores
WHERE id = $1;
