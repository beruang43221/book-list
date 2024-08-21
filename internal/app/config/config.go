package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	AppPort    string
}

func LoadConfig() (*Config, error) {
	// Memuat file .env
	err := godotenv.Load("../../.env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	// Buat struct Config
	cfg := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		AppPort:    os.Getenv("APP_PORT"),
	}

	// Validasi konfigurasi
	if cfg.DBHost == "" || cfg.DBPort == "" || cfg.DBUser == "" || cfg.DBPassword == "" || cfg.DBName == "" || cfg.AppPort == "" {
		return nil, fmt.Errorf("missing required configuration variables")
	}

	return cfg, nil
}
