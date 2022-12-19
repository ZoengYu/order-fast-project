-- name: CreateMenuFood :one
INSERT INTO food (
	menu_id,
	name,
	price
) VALUES (
	$1, $2, $3
) RETURNING *;

-- name: GetMenuFood :one
SELECT * FROM food
WHERE id = $1;
