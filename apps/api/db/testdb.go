package db

import (
	"log"
	"yoga/api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SetupTestDB initializes an in-memory SQLite DB and runs migrations
func SetupTestDB() *gorm.DB {
	testDB, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}

	err = testDB.AutoMigrate(
		&models.User{},
		&models.Class{},
		&models.Instructor{},
		&models.Location{},
		// Add any other models you have
	)
	if err != nil {
		log.Fatalf("Failed to migrate test database: %v", err)
	}

	return testDB
}
