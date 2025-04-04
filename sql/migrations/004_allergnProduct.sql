-- +goose Up
CREATE TABLE product_allergen (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	product_fk INTEGER NOT NULL,
	allergen_fk INTEGER NOT NULL,
	FOREIGN KEY (product_fk) REFERENCES products(id) ON DELETE CASCADE,
	FOREIGN KEY (allergen_fk) REFERENCES allergens(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE product_allergen;
