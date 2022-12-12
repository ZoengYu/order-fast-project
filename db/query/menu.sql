-- name: CreateStoreMenu :one
INSERT INTO menu (
    store_id,
	menu_name
) VALUES (
    $1, $2
) RETURNING *;


-- name: UpdateStoreMenu :one
UPDATE menu
SET menu_name = $3
WHERE store_id = $1 AND id = $2
RETURNING *;

-- name: GetStoreMenu :one
SELECT * FROM menu
WHERE store_id = $1 AND id = $2;

-- name: DeleteMenu :exec
DELETE FROM menu
WHERE id = $1;

-- name: ListStoreMenu :many
SELECT * FROM menu
WHERE store_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;
