package models

type Delivery struct {
	ID            uint64  `gorm:"primaryKey"`
	OrderID       uint64  
	City          string
    DeliveryAgentID uint64 `gorm:"foreignKey:DeliveryAgentID"`
	Status string
}

const (
	AVAILABLE = "AVAILABLE"
	UNAVAILABLE = "UNAVAILABLE"
	DELIVERED = "DELIVERED"
	UNDELIVERED = "UNDELIVERED"
)