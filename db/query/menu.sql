-- name: CreateStoreMenu :one
INSERT INTO menu (
    store_id,
	menu_name,
	created_at
) VALUES (
    $1, $2, $3
) RETURNING *;


-- name: UpdateStoreMenu :one
UPDATE menu
SET menu_name = $2
WHERE id = $1
RETURNING *;

-- name: GetStoreMenu :one
SELECT * FROM menu
WHERE store_id = $1 AND menu_name = $2;

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

-- name: AddFoodTag :one
INSERT INTO food_tag (
	menu_food_id,
	food_tag
) VALUES (
	$1, $2
) RETURNING *;

-- name: GetFoodTag :one
SELECT * FROM food_tag
WHERE menu_food_id = $1 AND food_tag = $2;

-- name: RemoveFoodTag :exec
DELETE FROM food_tag
WHERE menu_food_id = $1 AND food_tag = $2;

-- name: ListMenuFoodTag :many
SELECT food_tag.food_tag FROM food_tag
WHERE menu_food_id = $1;
