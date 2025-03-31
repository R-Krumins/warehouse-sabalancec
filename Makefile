build:
	go build -o bin/warehouse ./cmd

run: build
	./bin/warehouse

gen-db:
	sqlc generate

migrate-up:
	goose up

migrate-down:
	goose down
