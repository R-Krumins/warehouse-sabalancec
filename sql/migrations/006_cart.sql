-- +goose Up
CREATE TABLE cart (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_uuid TEXT NOT NULL,
	product_fk INTEGER NOT NULL,
	quantity INTEGER NOT NULL,
	FOREIGN KEY (user_uuid) REFERENCES users(uuid) ON DELETE CASCADE,
	FOREIGN KEY (product_fk) REFERENCES products(id) ON DELETE CASCADE,
	UNIQUE (user_uuid, product_fk)
);

-- +goose Down
DROP TABLE cart;

