package config

import (
	"go_api/internal/model"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate the User model to keep the schema up-to-date
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
	}
	dbSQL.Close()
}
