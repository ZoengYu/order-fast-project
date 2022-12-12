-- name: CreateMenuFood :one
INSERT INTO food (
	menu_id,
	name
) VALUES (
	$1, $2
) RETURNING *;

-- name: GetMenuFood :one
SELECT * FROM food
WHERE menu_id = $1 AND name = $2;
