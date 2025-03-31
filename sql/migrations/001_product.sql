-- +goose Up
CREATE TABLE products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    img TEXT NOT NULL,
    price REAL NOT NULL
);

-- +goose Down
DROP TABLE products;
