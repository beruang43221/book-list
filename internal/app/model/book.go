package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string    `gorm:"not null"`
	Author      string    `gorm:"not null"`
	Publication time.Time `json:"publication_date"`
	Publisher   string    `json:"publisher"`
	Pages       uint      `json:"pages"`
	CategoryID  uint      `gorm:"not null"`
	Category    Category  `gorm:"foreignKey:CategoryID"`
}
