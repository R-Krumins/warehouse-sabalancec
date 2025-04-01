-- +goose Up
CREATE TABLE allergens (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
	img TEXT NOT NULL,
    info TEXT NOT NULL
);

-- +goose Down
DROP TABLE allergens;
