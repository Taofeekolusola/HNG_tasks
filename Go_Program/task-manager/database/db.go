package database

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"task-manager/models"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=admin password=password dbname=taskmanager port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate models
	err = db.AutoMigrate(&models.User{}, &models.Task{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database connection established successfully")
	DB = db
}
