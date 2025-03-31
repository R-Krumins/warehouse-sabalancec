package main

import (
	"database/sql"
	"log"
	"warehouse/internal/database"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg := loadConfig()

	db, err := sql.Open("sqlite3", cfg.dbPath)
	if err != nil {
		log.Fatal(err)
	}

	queries := database.New(db)
	server := NewServer(db, queries)

	server.Run(cfg.port)
}
