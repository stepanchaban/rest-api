package db

import (
	"log"
	"my-api/internal/taskService"
	"my-api/internal/userService"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=yourpassword dbname=main port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	if err := db.AutoMigrate(&taskService.Task{}, &userService.User{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	return db, nil
}

