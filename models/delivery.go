package models

type Delivery struct {
	Id uint64
	Order_id uint64
	City string
	DeliveryAgent *User
	Status string
}

const (
	AVAILABLE = "AVAILABLE"
	UNAVAILABLE = "UNAVAILABLE"
)