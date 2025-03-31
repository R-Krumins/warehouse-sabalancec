-- name: GetProduct :many
SELECT * FROM products;

-- name: CreateProduct :one
INSERT INTO products (name, img, price) VALUES (?, ?, ?) RETURNING *;
