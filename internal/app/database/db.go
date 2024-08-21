package database

import (
	"fmt"

	"github.com/beruang43221/book-list/internal/app/config"
	"github.com/beruang43221/book-list/internal/app/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectDB() (*gorm.DB, error) {
	// Ambil konfigurasi database dari file config.go
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}

	// Buat koneksi database
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	// Migrasi model
	db.Debug().AutoMigrate(model.Book{}, model.Category{})

	return db, nil
}
