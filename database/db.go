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

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Error migrating User table: %v", err)
	}

	err = db.AutoMigrate(&models.Delivery{})
	if err != nil {
		log.Fatalf("Error migrating Delivery table: %v", err)
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

func FindAvailableDeliveryAgentByCity(db *gorm.DB, city string) (*models.User, error) {
	var user models.User
	result := db.Where("city = ? AND availability = ?", city, models.AVAILABLE).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	user.Availability = models.UNAVAILABLE
	db.Save(&user)
	return &user, nil
}

func FetchDeliveriesForAnAgent(db *gorm.DB, userID uint64) ([]models.Delivery) {
    var deliveries []models.Delivery
    db.Preload("DeliveryAgent").Where("delivery_agent_id = ?", userID).Find(&deliveries)
    return deliveries
}