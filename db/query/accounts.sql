-- name: GetAccounts :many
SELECT * FROM accounts;

-- name: CreateAccount :one
INSERT INTO accounts(first_name, last_name, balance, currency)
VALUES($1,$2,$3,$4) RETURNING *;