-- name: ReadTransfer :one
SELECT * FROM transfers
WHERE id = $1;

-- name: ReadAccountTransfer :many
SELECT * FROM transfers
WHERE from_account_id = $1
OR to_account_id = $1
LIMIT $2 OFFSET $3;

-- name: CreateTransfer :one
INSERT INTO transfers
( from_account_id, to_account_id, amount )
VALUES ($1,$2,$3) RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;

-- name: UpdateTransfer :one
UPDATE transfers
SET amount = $2
WHERE id = $1
RETURNING *;