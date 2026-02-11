package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	DB_SSLMODE  string
}

type Config struct {
	Db *DbConfig
}

func GetConfig() *Config {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error for loading file .env")
	}

	return &Config{
		&DbConfig{
			DB_HOST:     os.Getenv("DB_HOST"),
			DB_USER:     os.Getenv("DB_USER"),
			DB_PASSWORD: os.Getenv("DB_PASSWORD"),
			DB_NAME:     os.Getenv("DB_NAME"),
			DB_PORT:     os.Getenv("DB_PORT"),
			DB_SSLMODE:  os.Getenv("DB_SSLMODE"),
		},
	}
}
