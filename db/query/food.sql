-- name: AddMenuFood :one
INSERT INTO food (
	menu_id,
	food_name
) VALUES (
	$1, $2
) RETURNING *;

-- name: GetMenuFood :one
SELECT * FROM food
WHERE menu_id = $1 AND food_name = $2;
