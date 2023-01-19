-- name: CreateAccount :one
INSERT INTO accounts (
    owner
) VALUES (
    $1
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: UpdateAccount :one
UPDATE accounts
SET owner = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;
