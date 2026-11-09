package database

import (
	"log"

	"github.com/alisstor/backend/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(cfg config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(cfg.DatabaseURL()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Database connected successfully")
	return db, nil
}
