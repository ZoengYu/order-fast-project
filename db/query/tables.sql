-- name: CreateTable :one
INSERT INTO tables (
    store_id,
    table_id,
    table_name
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetTable :one
SELECT * FROM tables
WHERE store_id = $1 and table_id = $2;

-- name: ListTables :many
SELECT * FROM tables
WHERE store_id = $1
ORDER BY table_id;

-- name: UpdateTable :exec
UPDATE tables SET table_name = $3
where store_id = $1 AND table_id = $2;

-- name: DeleteTable :exec
DELETE FROM tables
WHERE id = $1;

-- name: DeleteTableByName :exec
DELETE FROM tables
WHERE store_id = $1 AND table_name = $2;
