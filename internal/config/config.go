package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port string
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error while reading config: %s", err.Error())
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return &Config{
		port,
	}
}
