package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	port   string
	dbPath string
}

func loadConfig() Config {
	godotenv.Load()

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatal("No DB_PATH environment variable defined")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("No PORT environment variable defined")
	}

	return Config{port, dbPath}
}
