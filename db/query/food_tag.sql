-- name: CreateMenuFoodTag :one
INSERT INTO food_tag (
	food_id,
	food_tag
) VALUES (
	$1, $2
) RETURNING *;

-- name: GetMenuFoodTag :one
SELECT * FROM food_tag
WHERE food_id = $1 AND food_tag = $2;

-- name: RemoveMenuFoodTag :exec
DELETE FROM food_tag
WHERE food_id = $1 AND food_tag = $2;

-- name: ListMenuFoodTag :many
SELECT food_tag.food_tag FROM food_tag
WHERE food_id = $1;
