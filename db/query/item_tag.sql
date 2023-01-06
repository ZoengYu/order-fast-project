-- name: CreateMenuItemTag :one
INSERT INTO item_tag (
	item_id,
	item_tag
) VALUES (
	$1, $2
) RETURNING *;

-- name: GetMenuItemTag :one
SELECT * FROM item_tag
WHERE item_id = $1 AND item_tag = $2;

-- name: RemoveMenuItemTag :exec
DELETE FROM item_tag
WHERE item_id = $1 AND item_tag = $2;

-- name: ListMenuItemTag :many
SELECT item_tag.item_tag FROM item_tag
WHERE item_id = $1;
