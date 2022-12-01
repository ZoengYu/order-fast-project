-- name: CreateStoreMenu :one
INSERT INTO menu (
    store_id,
	menu_name
) VALUES (
    $1, $2
) RETURNING *;


-- name: UpdateStoreMenu :one
UPDATE menu
SET menu_name = $3
WHERE store_id = $1 AND id = $2
RETURNING *;

-- name: GetStoreMenu :one
SELECT * FROM menu
WHERE store_id = $1 AND id = $2;

-- name: AddMenuFoodTag :one
INSERT INTO food_tag (
	menu_food_id,
	food_tag
) VALUES (
	$1, $2
) RETURNING *;

-- name: DeleteMenu :exec
DELETE FROM menu
WHERE id = $1;
