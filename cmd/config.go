package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	port       string
	dbPath     string
	authApiKey string
	jwtSecret  []byte
}

func loadConfig() Config {
	godotenv.Load()

	return Config{
		port:       getEnv("PORT"),
		dbPath:     getEnv("DB_PATH"),
		authApiKey: getEnv("AUTH_API_KEY"),
		jwtSecret:  []byte(getEnv("JWT_SECRET")),
	}
}

func getEnv(name string) string {
	env := os.Getenv(name)
	if env == "" {
		log.Fatalf("No %s enviroment variable defined!", name)
	}
	return env
}
