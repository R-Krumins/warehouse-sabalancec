-- name: CreateUser :one
INSERT INTO users (uuid, name, email, address)
VALUES (?, ?, ?, ?) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE uuid = ?;
