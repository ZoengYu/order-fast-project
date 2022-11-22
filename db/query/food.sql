-- name: AddMenuFood :one
INSERT INTO menu_food (
	menu_id,
	food_name,
	custom_option
) VALUES (
	$1, $2, $3
) RETURNING *;

-- name: GetMenuFood :one
SELECT * FROM menu_food
WHERE menu_id = $1 AND food_name = $2;
