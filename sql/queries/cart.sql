-- name: PatchCart :one
INSERT INTO cart (user_uuid, product_fk, quantity)
VALUES (?, ?, ?)
ON CONFLICT(user_uuid, product_fk) DO UPDATE SET
    quantity = cart.quantity + EXCLUDED.quantity
RETURNING *;

-- name: GetCartForUser :many
SELECT 
  c.id AS cart_item_id, 
  p.id AS product_id, 
  p.name, 
  p.img, 
  p.price AS price_per_unit, 
  c.quantity, 
  (p.price * c.quantity) AS sum_total
FROM 
  cart c
INNER JOIN 
  products p ON c.product_fk = p.id
WHERE
  c.user_uuid = sqlc.arg(userUuid);
