-- name: GetAllergen :many
SELECT * FROM allergens;

-- name: CreateAllergen :one
INSERT INTO allergens (name, img, info) VALUES (?, ?, ?) RETURNING *;

-- name: GetAllergenById :one
SELECT * FROM allergens WHERE id = ?;
