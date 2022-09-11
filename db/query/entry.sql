-- name: ReadEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
LIMIT $1 OFFSET $2;

-- name: UpdateEntries :one
UPDATE entries
SET amount = $2
WHERE id = $1 
RETURNING *;

-- name: DeleteEntries :exec
DELETE FROM entries
WHERE id = $1;

-- name: DeleteAllEntries :exec
DELETE FROM entries
WHERE account_id = $1;

-- name: CreateEntry :one
INSERT INTO entries
( account_id, amount )
VALUES ($1,$2) RETURNING *;