package main

import (
	"app/internal/api"
	"app/internal/database"
	"app/internal/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func main() {
	cfg := database.NewConfig()
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Run migrations
	migrate(db)

	// Start the server
	api.StartServer(db)
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
