-- name: PatchCart :one
INSERT INTO cart (user_uuid, product_fk, quantity)
VALUES (?, ?, ?)
ON CONFLICT(user_uuid, product_fk) DO UPDATE SET
    quantity = cart.quantity + EXCLUDED.quantity
RETURNING *;

-- name: GetAllCartItems :many
SELECT * FROM cart
WHERE user_uuid = sql.arg(userId)
