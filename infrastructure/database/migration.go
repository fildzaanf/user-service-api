package database

import (
	"log"

	"user-service-api/internal/user/adapter/model"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	migrator := db.Migrator()

	if !migrator.HasTable(&model.User{}) {
		if err := db.AutoMigrate(&model.User{}); err != nil {
			log.Fatalf("failed to migrate tables: %v", err)
		}
		log.Println("table users successfully created")
	} else {
		log.Println("table users already exists, skipping migration")
	}
}
