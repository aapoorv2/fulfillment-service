package database

import (
	"fmt"
	"fulfillment/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	username = "username"
	password = "password"
	dbName   = "fulfillment"
	sslMode  = "disable"
)

func Connection() *gorm.DB {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, username, password, dbName, sslMode)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}

	log.Println("Connected to the database")

	err = db.AutoMigrate(&models.User{}, &models.Delivery{})
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	return db
}

func FindByUsername(db *gorm.DB, username string) (*models.User, error) {
	var user models.User
	result := db.Where("username = ?", username).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
