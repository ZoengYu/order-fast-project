-- name: CreateMenuFood :one
INSERT INTO food (
	menu_id,
	name,
	price
) VALUES (
	$1, $2, $3
) RETURNING *;

-- name: GetFood :one
SELECT * FROM food
WHERE id = $1;

-- name: ListAllMenuFood :many
SELECT * FROM food
WHERE menu_id = $1;

-- name: ListMenuFood :many
SELECT * FROM food
WHERE menu_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: DeleteMenuFood :exec
DELETE FROM food
WHERE id = $1 AND menu_id = $2;

-- name: DeleteMenuFoodAll :exec
DELETE FROM food
WHERE menu_id = $1;
