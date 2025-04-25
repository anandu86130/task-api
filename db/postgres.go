package db

import (
	"fmt"
	"log"
	"task-api/config"
	"task-api/models"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Task{})
	if err != nil {
		log.Fatal("Failed to auto-migrate database: ", err)
	}

	return db
}

func GetDB() *gorm.DB {
	return db
}