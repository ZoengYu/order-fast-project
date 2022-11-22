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

-- name: AddMenuFoodTag :one
INSERT INTO food_tag (
	menu_food_id,
	food_tag
) VALUES (
	$1, $2
) RETURNING *;
