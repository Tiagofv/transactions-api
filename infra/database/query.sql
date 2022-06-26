-- name: GetTransaction :one
SELECT * FROM transactions
WHERE id = $1 LIMIT 1;

-- name: ListTransactions :many
SELECT * FROM transactions
ORDER BY id;

-- name: CreateTransaction :one
INSERT INTO transactions (
    account_id, operation_type_id, amount, event_date
) VALUES (
             $1, $2, $3, $4
         )
    RETURNING *;

-- name: CreateAccount :one
INSERT INTO accounts (
    document
) VALUES (
             $1
         )
    RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts WHERE id = $1 LIMIT 1;

-- name: GetAccountByDocument :one
SELECT * FROM accounts WHERE document = $1 LIMIT 1;

-- name: ListAccounts :one
SELECT * FROM accounts;

-- name: GetOperation :one
SELECT * FROM operation_types WHERE id = $1 LIMIT 1;