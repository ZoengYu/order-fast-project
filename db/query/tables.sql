-- name: CreateTable :one
INSERT INTO tables (
    store_id,
    table_id,
    table_name
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetStoreTable :one
SELECT * FROM tables
WHERE store_id = $1 and table_id = $2;

-- name: ListStoreTables :many
SELECT * FROM tables
WHERE store_id = $1
ORDER BY table_id
LIMIT $2
OFFSET $3;

-- name: UpdateStoreTable :exec
UPDATE tables SET table_name = $3
where store_id = $1 AND table_id = $2;

-- name: DeleteStoreTable :exec
DELETE FROM tables
WHERE id = $1;

-- name: DeleteStoreTableByName :exec
DELETE FROM tables
WHERE store_id = $1 AND table_name = $2;
