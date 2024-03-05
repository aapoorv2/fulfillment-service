package models

type User struct {
	ID           uint64 `gorm:"primaryKey"`
	Username     string
	Password     string
	City         string
	Availability string
}