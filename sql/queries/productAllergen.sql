-- name: CreateProductAllergen :one
INSERT INTO product_allergen (product_fk, allergen_fk) 
VALUES (?, ?) RETURNING *;

-- name: GetAllergensForProduct :many
SELECT a.id, a.name
FROM product_allergen pa
LEFT JOIN allergens a ON pa.allergen_fk = a.id
WHERE pa.product_fk = sqlc.arg(productId);


-- name: GetProductsForAllergen :many
SELECT * FROM product_allergen
WHERE allergen_fk = ?;
