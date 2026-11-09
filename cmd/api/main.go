package main

import (
	"fmt"
	"log"

	"github.com/alisstor/backend/internal/config"
	"github.com/alisstor/backend/internal/database"
	"github.com/alisstor/backend/internal/router"
)

func main() {
	cfg := config.Load()

	if _, err := database.New(cfg); err != nil {
		log.Printf("database connection failed: %v\n", err)
		log.Println("Make sure PostgreSQL is running or update your .env settings")
	}

	r := router.New(cfg)
	log.Printf("Server running on port %d", cfg.ServerPort)
	if err := r.Run(fmt.Sprintf(":%d", cfg.ServerPort)); err != nil {
		log.Fatal(err)
	}
}
