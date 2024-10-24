package database

import (
	"auth/models"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.AuthUser{})
	if err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}
	log.Println("Database migration completed successfully")
}
