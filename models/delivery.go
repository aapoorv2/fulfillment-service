package models

type User struct {
	Id uint64
	Username string
	Password string
	Availability string
}
type Delivery struct {
	Id uint64
	Order_id uint64
	City string
	DeliveryAgent *User
	Status string
}