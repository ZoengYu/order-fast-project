-- name: CreateMenuItem :one
INSERT INTO item (
	menu_id,
	name,
	price
) VALUES (
	$1, $2, $3
) RETURNING *;

-- name: GetItem :one
SELECT * FROM item
WHERE id = $1;

-- name: ListAllMenuItem :many
SELECT * FROM item
WHERE menu_id = $1;

-- name: ListMenuItem :many
SELECT * FROM item
WHERE menu_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: DeleteMenuItem :exec
DELETE FROM item
WHERE id = $1 AND menu_id = $2;

-- name: DeleteMenuItemAll :exec
DELETE FROM item
WHERE menu_id = $1;

-- name: UpdateMenuItem :one
UPDATE item
SET name = $2, price=$3
WHERE id = $1
RETURNING *;
