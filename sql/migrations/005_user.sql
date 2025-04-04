-- +goose Up
CREATE TABLE users (
	uuid TEXT PRIMARY KEY,
	name TEXT NOT NULL,
	email TEXT NOT NULL,
	address TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;
