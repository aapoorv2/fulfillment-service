package models

type Delivery struct {
	ID            uint64  `gorm:"primaryKey"`
	OrderID       uint64  
	City          string
	DeliveryAgent *User    `gorm:"foreignKey:DeliveryAgentID"`
    DeliveryAgentID uint64
	Status string
}

const (
	AVAILABLE = "AVAILABLE"
	UNAVAILABLE = "UNAVAILABLE"
	DELIVERED = "DELIVERED"
	UNDELIVERED = "UNDELIVERED"
)